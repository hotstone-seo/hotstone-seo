package service

// AddMetaTagRequest is request model for addMetaTag method
type AddMetaTagRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// AddTitleTagRequest is request model for addTitleTag method
type AddTitleTagRequest struct {
	Title string `json:"title"`
}
