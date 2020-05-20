package service

// RoleTypeRequest is request model for RoleType related method
type RoleTypeRequest struct {
	ID      int64        `json:"id"`
	Name    string       `json:"name"`
	Modules []ModuleItem `json:"modules"`
}

type ModuleItem struct {
	Module string `json:"name"`
}
