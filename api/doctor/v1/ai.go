package v1

import "github.com/gogf/gf/v2/frame/g"

type AISessionListReq struct {
	g.Meta   `path:"/sessions" method:"get" tags:"医生AI查看" summary:"获取AI会话列表"`
	Page     *int   `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int   `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	PetID    *int64 `json:"pet_id" p:"pet_id" v:"min:1#pet_id不合法" dc:"宠物ID筛选"`
}

type AISessionListItem struct {
	ID        int64  `json:"id" dc:"会话ID"`
	SessionNo string `json:"session_no" dc:"会话编号"`
	PetID     int64  `json:"pet_id" dc:"宠物ID"`
	PetName   string `json:"pet_name" dc:"宠物名称"`
	ModelName string `json:"model_name" dc:"模型名称"`
	Status    int    `json:"status" dc:"状态"`
	CreatedAt string `json:"created_at" dc:"创建时间"`
}

type AISessionListRes struct {
	List       []AISessionListItem `json:"list" dc:"会话列表"`
	Pagination Pagination          `json:"pagination" dc:"分页信息"`
}

type AISessionMessageListReq struct {
	g.Meta    `path:"/sessions/{session_id}/messages" method:"get" tags:"医生AI查看" summary:"获取AI会话消息记录"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#会话ID不能为空|会话ID不合法" dc:"会话ID"`
	Page      *int  `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize  *int  `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
}

type AISessionMessageItem struct {
	ID             int64  `json:"id" dc:"消息ID"`
	SenderType     int    `json:"sender_type" dc:"发送者类型"`
	SenderID       *int64 `json:"sender_id" dc:"发送者ID"`
	MessageContent string `json:"message_content" dc:"消息内容"`
	CreatedAt      string `json:"created_at" dc:"创建时间"`
}

type AISessionMessageListRes struct {
	List []AISessionMessageItem `json:"list" dc:"消息列表"`
}

type AISessionAnalysisListReq struct {
	g.Meta    `path:"/sessions/{session_id}/analysis-records" method:"get" tags:"医生AI查看" summary:"获取AI分析结果"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#会话ID不能为空|会话ID不合法" dc:"会话ID"`
}

type AISessionAnalysisItem struct {
	ID               int64  `json:"id" dc:"分析ID"`
	AnalysisType     int    `json:"analysis_type" dc:"分析类型"`
	InputSource      int    `json:"input_source" dc:"输入来源"`
	AnalysisResult   string `json:"analysis_result" dc:"分析结果"`
	RuleBasedResult  string `json:"rule_based_result" dc:"规则分析结果"`
	LlmBasedResult   string `json:"llm_based_result" dc:"模型分析结果"`
	RiskLevel        *int   `json:"risk_level" dc:"风险等级"`
	ReviewedByDoctor int    `json:"reviewed_by_doctor" dc:"是否已审核"`
	CreatedAt        string `json:"created_at" dc:"创建时间"`
}

type AISessionAnalysisListRes struct {
	List []AISessionAnalysisItem `json:"list" dc:"分析记录列表"`
}
