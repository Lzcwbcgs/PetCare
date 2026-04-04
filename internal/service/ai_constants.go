package service

const aiTimeLayout = "2006-01-02 15:04:05"

const (
	aiSenderTypeUser   = 1
	aiSenderTypeAI     = 2
	aiSenderTypeDoctor = 3

	aiSessionStatusActive = 1

	aiMessageTypeText = 1

	aiAnalysisTypeRisk = 3
	aiInputSourceChat  = 1
)

const (
	aiSessionColumnProviderName  = "provider_name"
	aiSessionColumnRagEnabled    = "rag_enabled"
	aiSessionColumnLastMessageAt = "last_message_at"
	aiSessionColumnRetrievalCnt  = "retrieval_count"

	aiMessageColumnProviderType   = "provider_type"
	aiMessageColumnProviderName   = "provider_name"
	aiMessageColumnFinishReason   = "finish_reason"
	aiMessageColumnPromptTokens   = "prompt_tokens"
	aiMessageColumnCompleteTokens = "completion_tokens"
	aiMessageColumnExtraMetadata  = "extra_metadata"

	aiAnalysisColumnSummaryTitle = "summary_title"
	aiAnalysisColumnRiskLevel    = "risk_level"
	aiAnalysisColumnRuleBased    = "rule_based_result"
	aiAnalysisColumnLlmBased     = "llm_based_result"
)
