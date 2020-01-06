package service

// AddMetaTagRequest is request model for addMetaTag method
type AddMetaTagRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// AddTitleTagRequest is request model for addTitleTag method
type AddTitleTagRequest struct {
	RuleID   int64  `json:"rule_id"`
	LocaleID int64  `json:"locale_id"`
	Title    string `json:"title"`
}

// AddCanonicalTagRequest is request model for addCanonicalTag method
type AddCanonicalTagRequest struct {
	Canonical string `json:"canonical"`
	RuleID    int64  `json:"rule_id"`
}

// AddScriptTagRequest is request model for addScriptTag method
type AddScriptTagRequest struct {
	Type         string `json:"type"`
	RuleID       int64  `json:"rule_id"`
	DataSourceID int64  `json:"datasource_id"`
}
