package service

// AddMetaTagRequest is request model for addMetaTag method
type AddMetaTagRequest struct {
	RuleID  int64  `json:"rule_id"`
	Locale  string `json:"locale"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// AddTitleTagRequest is request model for addTitleTag method
type AddTitleTagRequest struct {
	RuleID int64  `json:"rule_id"`
	Locale string `json:"locale"`
	Title  string `json:"title"`
}

// AddCanonicalTagRequest is request model for addCanonicalTag method
type AddCanonicalTagRequest struct {
	Canonical string `json:"canonical"`
	Locale    string `json:"locale"`
	RuleID    int64  `json:"rule_id"`
	Href      string `json:"href"`
}

// AddScriptTagRequest is request model for addScriptTag method
type AddScriptTagRequest struct {
	Type         string `json:"type"`
	Locale       string `json:"locale"`
	RuleID       int64  `json:"rule_id"`
	DataSourceID int64  `json:"datasource_id"`
}
