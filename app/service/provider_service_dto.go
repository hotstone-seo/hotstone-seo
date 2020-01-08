package service

// MatchRuleRequest is request for match rule
type MatchRuleRequest struct {
	Path string `json:"path"`
}

type MatchRuleResponse struct {
	RuleID    int64             `json:"rule_id"`
	PathParam map[string]string `json:"path_param"`
}

type RetrieveDataRequest struct {
	DataSourceID int64 `json:"data_source_id"`
}

type ProvideTagsRequest struct {
	RuleID   int64       `json:"rule_id"`
	LocaleID int64       `json:"locale"`
	Data     interface{} `json:"data"`
}
