package service

// MatchRuleRequest is request for match rule
type MatchRuleRequest struct {
	Path string `json:"path"`
}

// MatchRuleResponse is response of match rule
type MatchRuleResponse struct {
	RuleID    int64             `json:"rule_id"`
	PathParam map[string]string `json:"path_param"`
}

// RetrieveDataRequest is request for retrieve data
type RetrieveDataRequest struct {
	DataSourceID int64             `json:"data_source_id"`
	PathParam    map[string]string `json:"path_param"`
}

// RetrieveDataResponse is response for retreieve data
type RetrieveDataResponse struct {
	Data []byte
}

// ProvideTagsRequest us reques for provide tags
type ProvideTagsRequest struct {
	RuleID    int64             `json:"rule_id"`
	PathParam map[string]string `json:"path_param"`
	Locale    string            `json:"locale"`
	Data      interface{}       `json:"data"`
}
