package service

// MatchRuleRequest is request for match rule
type MatchRuleRequest struct {
	Path string `json:"path"`
}

type MatchRuleResponse struct {
	RuleID    int64             `json:"ruleID"`
	PathParam map[string]string `json:"pathParam"`
}

type RetrieveDataRequest struct {
	RuleID int64 `json:"ruleID"`
}
