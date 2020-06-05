package service

type ModuleRequest struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	APIPaths []APIPath `json:"api_path"`
	Pattern  string    `json:"pattern"`
	Label    string    `json:"label"`
}

type APIPath struct {
	Path string `json:"path"`
}
