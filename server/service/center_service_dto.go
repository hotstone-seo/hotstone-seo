package service

// AddMetaTagRequest is request model for addMetaTag method
// TODO: remove "Add" from each of the struct name
type AddMetaTagRequest struct {
	RuleID  int64  `json:"ruleID"`
	Locale  string `json:"locale"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// AddTitleTagRequest is request model for addTitleTag method
type AddTitleTagRequest struct {
	RuleID int64  `json:"ruleID"`
	Locale string `json:"locale"`
	Title  string `json:"title"`
}

// AddCanonicalTagRequest is request model for addCanonicalTag method
type AddCanonicalTagRequest struct {
	RuleID int64  `json:"ruleID"`
	Locale string `json:"locale"`
	Href   string `json:"href"`
}

// AddScriptTagRequest is request model for addScriptTag method
type AddScriptTagRequest struct {
	Type   string `json:"type"`
	RuleID int64  `json:"ruleID"`
	Locale string `json:"locale"`
	Source string `json:"source"`
}

func NewTagRequest(tagType string) interface{} {
	switch tagType {
	case "meta":
		return &AddMetaTagRequest{}
	case "title":
		return &AddTitleTagRequest{}
	case "canonical":
		return &AddCanonicalTagRequest{}
	case "script":
		return &AddScriptTagRequest{}
	default:
		return nil
	}
}
