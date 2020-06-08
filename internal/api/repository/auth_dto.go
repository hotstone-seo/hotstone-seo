package repository

type TokenReq struct {
	Holder    string `json:"holder,omitempty"`
	SetCookie bool   `json:"set_cookie,omitempty"`
}

type TokenResp struct {
	Token string `json:"token,omitempty"`
}
