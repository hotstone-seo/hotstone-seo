package service

// MatchRuleRequest is request for match rule
type MatchRuleRequest struct {
	Path string `json:"path"`
}

type RetrieveDataRequest struct {
	RuleID int64 `json:"ruleID"`
}
