package service

// MetaTagRequest is request model for MetaTag related method
type MetaTagRequest struct {
	ID      int64  `json:"id"`
	RuleID  int64  `json:"rule_id"`
	Locale  string `json:"locale"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// TitleTagRequest is request model for TitleTag related method
type TitleTagRequest struct {
	ID     int64  `json:"id"`
	RuleID int64  `json:"rule_id"`
	Locale string `json:"locale"`
	Title  string `json:"title"`
}

// CanonicalTagRequest is request model for CanonicalTag related method
type CanonicalTagRequest struct {
	ID     int64  `json:"id"`
	RuleID int64  `json:"rule_id"`
	Locale string `json:"locale"`
	Href   string `json:"href"`
}

// ScriptTagRequest is request model for ScriptTag related method
type ScriptTagRequest struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	RuleID int64  `json:"rule_id"`
	Locale string `json:"locale"`
	Source string `json:"source"`
}

type FAQPageRequest struct {
	ID     int64 `json:"id"`
	RuleID int64 `json:"rule_id"`
	FAQs   []FAQ `json:"faqs"`
}

type FAQ struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type BreadcrumbListRequest struct {
	ID       int64            `json:"id"`
	RuleID   int64            `json:"rule_id"`
	ListItem []BreadcrumbItem `json:"list_item"`
}

type BreadcrumbItem struct {
	Name string `json:"name"`
	Item string `json:"item"`
}
