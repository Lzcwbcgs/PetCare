package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"PetCare/internal/consts"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

const (
	knowledgeDocumentTable = "knowledge_document"
	knowledgeChunkTable    = "knowledge_chunk"
	knowledgeJobTable      = "knowledge_job"

	knowledgeStatusUploaded   = 1
	knowledgeStatusProcessing = 2
	knowledgeStatusSuccess    = 3
	knowledgeStatusFailed     = 4

	knowledgeChunkStatusPending = 1
	knowledgeChunkStatusSuccess = 2
	knowledgeChunkStatusFailed  = 3

	knowledgeJobTypeParseEmbedding = 1
	knowledgeJobStatusPending      = 1
	knowledgeJobStatusProcessing   = 2
	knowledgeJobStatusSuccess      = 3
	knowledgeJobStatusFailed       = 4
)

type (
	AdminKnowledgeUploadInput struct {
		OperatorID int64
		File       *ghttp.UploadFile
		Category   *string
		Title      *string
	}

	AdminKnowledgeUploadOutput struct {
		KnowledgeID int64
		FileName    string
		Status      string
	}

	AdminKnowledgeStatusInput struct {
		OperatorID  int64
		KnowledgeID *int64
	}

	AdminKnowledgeStatusItem struct {
		KnowledgeID    int64
		FileName       string
		Status         string
		Progress       int
		ChunkTotal     int
		EmbeddedChunks int
		VectorCount    int
		ErrorMessage   string
		UpdatedAt      string
	}

	AdminKnowledgeStatusOutput struct {
		Items []AdminKnowledgeStatusItem
	}
)

type IAdminKnowledge interface {
	Upload(ctx context.Context, in AdminKnowledgeUploadInput) (*AdminKnowledgeUploadOutput, error)
	Status(ctx context.Context, in AdminKnowledgeStatusInput) (*AdminKnowledgeStatusOutput, error)
}

var AdminKnowledge IAdminKnowledge = adminKnowledgeService{}

type adminKnowledgeService struct{}

type knowledgeDocumentDO struct {
	g.Meta            `orm:"table:knowledge_document, do:true"`
	Id                any
	DocumentNo        any
	Title             any
	DocumentType      any
	SourceType        any
	FileName          any
	FileUrl           any
	FileSize          any
	ContentText       any
	EmbeddingProvider any
	EmbeddingModel    any
	ChunkCount        any
	Status            any
	ErrorMessage      any
	CreatedBy         any
	CreatedAt         any
	UpdatedAt         any
}

type knowledgeChunkDO struct {
	g.Meta           `orm:"table:knowledge_chunk, do:true"`
	Id               any
	DocumentId       any
	ChunkNo          any
	ChunkIndex       any
	Content          any
	TokenCount       any
	VectorStoreType  any
	VectorCollection any
	VectorPointId    any
	Status           any
	CreatedAt        any
	UpdatedAt        any
}

type knowledgeJobDO struct {
	g.Meta       `orm:"table:knowledge_job, do:true"`
	Id           any
	JobNo        any
	DocumentId   any
	JobType      any
	Status       any
	Progress     any
	StartedAt    any
	FinishedAt   any
	ErrorMessage any
	CreatedAt    any
	UpdatedAt    any
}

type knowledgeChunkData struct {
	Index      int
	Content    string
	TokenCount int
}

func (s adminKnowledgeService) Upload(ctx context.Context, in AdminKnowledgeUploadInput) (*AdminKnowledgeUploadOutput, error) {
	if in.OperatorID <= 0 {
		return nil, consts.NewUnauthorizedError("")
	}
	if in.File == nil {
		return nil, consts.NewBadRequestError("please upload file")
	}

	fileType := normalizeKnowledgeFileType(path.Ext(in.File.Filename))
	if fileType == "" {
		return nil, consts.NewBadRequestError("only txt/md/pdf files are supported")
	}

	now := time.Now()
	saveDir := gfile.Join("resource/public/uploads/knowledge", now.Format("2006"), now.Format("01"), now.Format("02"))
	fileName, err := in.File.Save(saveDir, true)
	if err != nil {
		return nil, consts.WrapInternalError(err, "save file failed")
	}

	filePath := gfile.Join(saveDir, fileName)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, consts.WrapInternalError(err, "read uploaded file metadata failed")
	}

	title := strings.TrimSpace(derefString(in.Title))
	if title == "" {
		title = fileName
	}
	fileURL := "/uploads/knowledge/" + now.Format("2006/01/02") + "/" + fileName

	documentNo, err := generateKnowledgeDocumentNo(ctx, now)
	if err != nil {
		return nil, err
	}
	jobNo, err := generateKnowledgeJobNo(ctx, now)
	if err != nil {
		return nil, err
	}

	var documentID int64
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := tx.Model(knowledgeDocumentTable).Data(knowledgeDocumentDO{
			DocumentNo:   documentNo,
			Title:        title,
			DocumentType: fileType,
			SourceType:   1,
			FileName:     fileName,
			FileUrl:      fileURL,
			FileSize:     fileInfo.Size(),
			ChunkCount:   0,
			Status:       knowledgeStatusProcessing,
			CreatedBy:    in.OperatorID,
			CreatedAt:    now,
			UpdatedAt:    now,
		}).Insert()
		if err != nil {
			return consts.WrapInternalError(err, "create knowledge document failed")
		}
		documentID, err = result.LastInsertId()
		if err != nil {
			return consts.WrapInternalError(err, "read knowledge document id failed")
		}

		_, err = tx.Model(knowledgeJobTable).Data(knowledgeJobDO{
			JobNo:      jobNo,
			DocumentId: documentID,
			JobType:    knowledgeJobTypeParseEmbedding,
			Status:     knowledgeJobStatusPending,
			Progress:   0,
			CreatedAt:  now,
			UpdatedAt:  now,
		}).Insert()
		if err != nil {
			return consts.WrapInternalError(err, "create knowledge job failed")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	RecordOperationLogByRole(ctx, consts.RoleAdmin, in.OperatorID, "knowledge", "upload", fmt.Sprintf("upload knowledge document id=%d", documentID))

	category := strings.TrimSpace(derefString(in.Category))
	go s.processKnowledgeDocument(context.Background(), documentID, filePath, fileType, category)

	return &AdminKnowledgeUploadOutput{
		KnowledgeID: documentID,
		FileName:    fileName,
		Status:      "processing",
	}, nil
}

func (s adminKnowledgeService) Status(ctx context.Context, in AdminKnowledgeStatusInput) (*AdminKnowledgeStatusOutput, error) {
	model := g.DB().Model(knowledgeDocumentTable)
	if in.KnowledgeID != nil {
		model = model.Where("id", *in.KnowledgeID)
	} else {
		model = model.OrderDesc("id").Limit(20)
	}

	records, err := model.All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query knowledge status failed")
	}

	items := make([]AdminKnowledgeStatusItem, 0, len(records))
	for _, record := range records {
		documentID := record["id"].Int64()

		chunkTotal := record["chunk_count"].Int()
		if chunkTotal <= 0 {
			chunkTotal, _ = g.DB().Model(knowledgeChunkTable).Where("document_id", documentID).Count()
		}
		embeddedChunks, _ := g.DB().Model(knowledgeChunkTable).
			Where("document_id", documentID).
			Where("status", knowledgeChunkStatusSuccess).
			Count()

		statusValue := record["status"].Int()
		progress := calcKnowledgeProgress(statusValue, chunkTotal, embeddedChunks)

		items = append(items, AdminKnowledgeStatusItem{
			KnowledgeID:    documentID,
			FileName:       record["file_name"].String(),
			Status:         knowledgeStatusText(statusValue),
			Progress:       progress,
			ChunkTotal:     chunkTotal,
			EmbeddedChunks: embeddedChunks,
			VectorCount:    embeddedChunks,
			ErrorMessage:   record["error_message"].String(),
			UpdatedAt:      formatGTime(record["updated_at"].GTime()),
		})
	}

	RecordOperationLogByRole(ctx, consts.RoleAdmin, in.OperatorID, "knowledge", "status", "query knowledge status")
	return &AdminKnowledgeStatusOutput{Items: items}, nil
}

func (s adminKnowledgeService) processKnowledgeDocument(ctx context.Context, documentID int64, filePath string, fileType string, category string) {
	startAt := time.Now()
	_, _ = g.DB().Model(knowledgeJobTable).
		Where("document_id", documentID).
		Data(knowledgeJobDO{
			Status:    knowledgeJobStatusProcessing,
			Progress:  10,
			StartedAt: startAt,
			UpdatedAt: startAt,
		}).
		Update()

	contentText, err := parseKnowledgeContent(filePath, fileType)
	if err != nil {
		s.markKnowledgeFailed(ctx, documentID, err.Error())
		return
	}

	chunks := splitKnowledgeChunks(contentText, aiChunkSize(ctx), aiChunkOverlap(ctx))
	if len(chunks) == 0 {
		s.markKnowledgeFailed(ctx, documentID, "no text chunks generated")
		return
	}

	now := time.Now()
	_, _ = g.DB().Model(knowledgeDocumentTable).
		Where("id", documentID).
		Data(knowledgeDocumentDO{
			ContentText: contentText,
			ChunkCount:  len(chunks),
			UpdatedAt:   now,
		}).
		Update()

	insertedChunkIDs := make([]int64, 0, len(chunks))
	chunkTexts := make([]string, 0, len(chunks))
	chunkPointIDs := make([]string, 0, len(chunks))

	for _, chunk := range chunks {
		chunkNo := fmt.Sprintf("KCH%d%04d", documentID, chunk.Index+1)
		result, insertErr := g.DB().Model(knowledgeChunkTable).Data(knowledgeChunkDO{
			DocumentId: documentID,
			ChunkNo:    chunkNo,
			ChunkIndex: chunk.Index,
			Content:    chunk.Content,
			TokenCount: chunk.TokenCount,
			Status:     knowledgeChunkStatusPending,
			CreatedAt:  now,
			UpdatedAt:  now,
		}).Insert()
		if insertErr != nil {
			s.markKnowledgeFailed(ctx, documentID, "insert chunks failed")
			return
		}
		chunkID, _ := result.LastInsertId()
		insertedChunkIDs = append(insertedChunkIDs, chunkID)
		chunkTexts = append(chunkTexts, chunk.Content)
		chunkPointIDs = append(chunkPointIDs, chunkNo)
	}

	providerType := aiEmbeddingProvider(ctx)
	provider, err := aiProviderByType(ctx, providerType)
	if err != nil {
		s.markKnowledgeFailed(ctx, documentID, err.Error())
		return
	}
	embeddingModel := aiDefaultEmbeddingModel(ctx, providerType)
	vectors, err := provider.Embed(ctx, chunkTexts, embeddingModel)
	if err != nil {
		s.markKnowledgeFailed(ctx, documentID, "embedding failed: "+err.Error())
		return
	}
	if len(vectors) != len(chunkTexts) || len(vectors) == 0 {
		s.markKnowledgeFailed(ctx, documentID, "embedding result size mismatch")
		return
	}

	if aiVectorStoreType(ctx) != "qdrant" {
		s.markKnowledgeFailed(ctx, documentID, "only qdrant vector store is supported in current version")
		return
	}

	qdrant := newQdrantClient(ctx)
	if err = qdrant.ensureCollection(ctx, len(vectors[0])); err != nil {
		s.markKnowledgeFailed(ctx, documentID, "qdrant ensure collection failed: "+err.Error())
		return
	}

	points := make([]qdrantPoint, 0, len(vectors))
	for i := range vectors {
		points = append(points, qdrantPoint{
			ID:     chunkPointIDs[i],
			Vector: vectors[i],
			Payload: map[string]any{
				"document_id": documentID,
				"chunk_id":    insertedChunkIDs[i],
				"chunk_index": i,
				"content":     chunkTexts[i],
				"category":    category,
			},
		})
	}
	if err = qdrant.upsertPoints(ctx, points); err != nil {
		s.markKnowledgeFailed(ctx, documentID, "qdrant upsert failed: "+err.Error())
		return
	}

	for i, chunkID := range insertedChunkIDs {
		_, _ = g.DB().Model(knowledgeChunkTable).Where("id", chunkID).Data(knowledgeChunkDO{
			VectorStoreType:  "qdrant",
			VectorCollection: aiQdrantCollection(ctx),
			VectorPointId:    chunkPointIDs[i],
			Status:           knowledgeChunkStatusSuccess,
			UpdatedAt:        time.Now(),
		}).Update()
	}

	_, _ = g.DB().Model(knowledgeDocumentTable).
		Where("id", documentID).
		Data(knowledgeDocumentDO{
			EmbeddingProvider: provider.Name(),
			EmbeddingModel:    embeddingModel,
			Status:            knowledgeStatusSuccess,
			ErrorMessage:      "",
			UpdatedAt:         time.Now(),
		}).
		Update()

	_, _ = g.DB().Model(knowledgeJobTable).
		Where("document_id", documentID).
		Data(knowledgeJobDO{
			Status:     knowledgeJobStatusSuccess,
			Progress:   100,
			FinishedAt: time.Now(),
			UpdatedAt:  time.Now(),
		}).
		Update()
}

func (s adminKnowledgeService) markKnowledgeFailed(ctx context.Context, documentID int64, message string) {
	now := time.Now()
	_, _ = g.DB().Model(knowledgeDocumentTable).
		Where("id", documentID).
		Data(knowledgeDocumentDO{
			Status:       knowledgeStatusFailed,
			ErrorMessage: truncateMessage(message, 500),
			UpdatedAt:    now,
		}).
		Update()

	_, _ = g.DB().Model(knowledgeChunkTable).
		Where("document_id", documentID).
		Where("status", knowledgeChunkStatusPending).
		Data(knowledgeChunkDO{
			Status:    knowledgeChunkStatusFailed,
			UpdatedAt: now,
		}).
		Update()

	_, _ = g.DB().Model(knowledgeJobTable).
		Where("document_id", documentID).
		Data(knowledgeJobDO{
			Status:       knowledgeJobStatusFailed,
			Progress:     100,
			FinishedAt:   now,
			ErrorMessage: truncateMessage(message, 500),
			UpdatedAt:    now,
		}).
		Update()
}

func normalizeKnowledgeFileType(ext string) string {
	switch strings.ToLower(strings.TrimPrefix(strings.TrimSpace(ext), ".")) {
	case "txt":
		return "txt"
	case "md":
		return "md"
	case "pdf":
		return "pdf"
	default:
		return ""
	}
}

func parseKnowledgeContent(filePath string, fileType string) (string, error) {
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var content string
	switch fileType {
	case "txt", "md":
		content = string(contentBytes)
	case "pdf":
		content = extractTextFromPDF(contentBytes)
	default:
		return "", fmt.Errorf("unsupported file type")
	}

	content = normalizeContentText(content)
	if strings.TrimSpace(content) == "" {
		return "", fmt.Errorf("document has no extractable text")
	}
	return content, nil
}

func extractTextFromPDF(content []byte) string {
	raw := string(content)
	tjRegex := regexp.MustCompile(`\(([^()]*)\)\s*Tj`)
	tjMatches := tjRegex.FindAllStringSubmatch(raw, -1)

	builder := strings.Builder{}
	for _, match := range tjMatches {
		if len(match) < 2 {
			continue
		}
		builder.WriteString(unescapePDFText(match[1]))
		builder.WriteString("\n")
	}

	tjArrayRegex := regexp.MustCompile(`\[(.*?)\]\s*TJ`)
	arrayMatches := tjArrayRegex.FindAllStringSubmatch(raw, -1)
	innerRegex := regexp.MustCompile(`\(([^()]*)\)`)
	for _, arr := range arrayMatches {
		if len(arr) < 2 {
			continue
		}
		inners := innerRegex.FindAllStringSubmatch(arr[1], -1)
		for _, inner := range inners {
			if len(inner) < 2 {
				continue
			}
			builder.WriteString(unescapePDFText(inner[1]))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func unescapePDFText(text string) string {
	replacer := strings.NewReplacer(
		`\\n`, "\n",
		`\\r`, "\n",
		`\\t`, " ",
		`\\(`, "(",
		`\\)`, ")",
		`\\\\`, `\`,
	)
	return replacer.Replace(text)
}

func normalizeContentText(content string) string {
	content = strings.ReplaceAll(content, "\x00", "")
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")

	var builder strings.Builder
	for _, r := range content {
		if r == '\n' || r == '\t' || r == ' ' {
			builder.WriteRune(r)
			continue
		}
		if r < 32 {
			continue
		}
		builder.WriteRune(r)
	}
	return strings.TrimSpace(builder.String())
}

func splitKnowledgeChunks(content string, chunkSize int, chunkOverlap int) []knowledgeChunkData {
	runes := []rune(strings.TrimSpace(content))
	if len(runes) == 0 {
		return nil
	}
	if chunkSize <= 0 {
		chunkSize = 700
	}
	if chunkOverlap < 0 {
		chunkOverlap = 0
	}
	if chunkOverlap >= chunkSize {
		chunkOverlap = chunkSize / 5
	}

	step := chunkSize - chunkOverlap
	if step <= 0 {
		step = chunkSize
	}

	chunks := make([]knowledgeChunkData, 0)
	index := 0
	for start := 0; start < len(runes); start += step {
		end := start + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		contentPart := strings.TrimSpace(string(runes[start:end]))
		if contentPart == "" {
			continue
		}
		chunks = append(chunks, knowledgeChunkData{
			Index:      index,
			Content:    contentPart,
			TokenCount: len([]rune(contentPart)),
		})
		index++
		if end == len(runes) {
			break
		}
	}
	return chunks
}

func generateKnowledgeDocumentNo(ctx context.Context, now time.Time) (string, error) {
	count, err := countKnowledgeRowsByDay(ctx, knowledgeDocumentTable, now)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("KDO%s%04d", now.Format("20060102"), count+1), nil
}

func generateKnowledgeJobNo(ctx context.Context, now time.Time) (string, error) {
	count, err := countKnowledgeRowsByDay(ctx, knowledgeJobTable, now)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("KJB%s%04d", now.Format("20060102"), count+1), nil
}

func countKnowledgeRowsByDay(ctx context.Context, table string, now time.Time) (int, error) {
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	count, err := g.DB().Model(table).Ctx(ctx).
		WhereGTE("created_at", startOfDay).
		WhereLT("created_at", endOfDay).
		Count()
	if err != nil {
		return 0, consts.WrapInternalError(err, "query row count failed")
	}
	return count, nil
}

func calcKnowledgeProgress(status int, chunkTotal int, embeddedChunks int) int {
	switch status {
	case knowledgeStatusSuccess:
		return 100
	case knowledgeStatusFailed:
		if chunkTotal <= 0 {
			return 0
		}
		return embeddedChunks * 100 / chunkTotal
	default:
		if chunkTotal <= 0 {
			return 0
		}
		progress := embeddedChunks * 100 / chunkTotal
		if progress > 99 {
			progress = 99
		}
		return progress
	}
}

func knowledgeStatusText(status int) string {
	switch status {
	case knowledgeStatusUploaded:
		return "uploaded"
	case knowledgeStatusProcessing:
		return "processing"
	case knowledgeStatusSuccess:
		return "success"
	case knowledgeStatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

func truncateMessage(message string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	content := strings.TrimSpace(message)
	runes := []rune(content)
	if len(runes) <= maxLen {
		return content
	}
	return string(runes[:maxLen])
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func toJSONString(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}
	return string(data)
}
