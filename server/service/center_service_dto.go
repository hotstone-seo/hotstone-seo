package service

// MetaTagRequest is request model for MetaTag related method
type MetaTagRequest struct {
	ID      int64  `json:"id"`
	RuleID  int64  `json:"ruleID"`
	Locale  string `json:"locale"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// TitleTagRequest is request model for TitleTag related method
type TitleTagRequest struct {
	ID     int64  `json:"id"`
	RuleID int64  `json:"ruleID"`
	Locale string `json:"locale"`
	Title  string `json:"title"`
}

// CanonicalTagRequest is request model for CanonicalTag related method
type CanonicalTagRequest struct {
	ID     int64  `json:"id"`
	RuleID int64  `json:"ruleID"`
	Locale string `json:"locale"`
	Href   string `json:"href"`
}

// ScriptTagRequest is request model for ScriptTag related method
type ScriptTagRequest struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	RuleID int64  `json:"ruleID"`
	Locale string `json:"locale"`
	Source string `json:"source"`
}
