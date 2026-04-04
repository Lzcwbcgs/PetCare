package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type KnowledgeUploadReq struct {
	g.Meta   `path:"/knowledge/upload" method:"post" mime:"multipart/form-data" tags:"医疗知识库管理" summary:"上传知识库文档"`
	File     *ghttp.UploadFile `json:"file" type:"file" v:"required#请上传知识库文件" dc:"知识库文件，支持txt/md/pdf"`
	Category *string           `json:"category" v:"max-length:50#分类长度不能超过50" dc:"分类，如guide/disease/drug"`
	Title    *string           `json:"title" v:"max-length:200#标题长度不能超过200" dc:"文档标题"`
}

type KnowledgeUploadRes struct {
	KnowledgeID int64  `json:"knowledge_id" dc:"知识文档ID"`
	FileName    string `json:"file_name" dc:"文件名"`
	Status      string `json:"status" dc:"状态"`
}

type KnowledgeStatusReq struct {
	g.Meta      `path:"/knowledge/status" method:"get" tags:"医疗知识库管理" summary:"查看知识库向量化进度"`
	KnowledgeID *int64 `json:"knowledge_id" p:"knowledge_id" v:"min:1#knowledge_id不合法" dc:"指定知识文档ID"`
}

type KnowledgeStatusItem struct {
	KnowledgeID  int64  `json:"knowledge_id" dc:"知识文档ID"`
	FileName     string `json:"file_name" dc:"文件名"`
	Status       string `json:"status" dc:"状态"`
	Progress     int    `json:"progress" dc:"进度，0-100"`
	ChunkCount   int    `json:"chunk_count" dc:"切片总数"`
	VectorCount  int    `json:"vector_count" dc:"向量数量"`
	ErrorMessage string `json:"error_message" dc:"错误信息"`
	UpdatedAt    string `json:"updated_at" dc:"更新时间"`
}

type KnowledgeStatusRes struct {
	List []KnowledgeStatusItem `json:"list" dc:"文档状态列表"`
}
