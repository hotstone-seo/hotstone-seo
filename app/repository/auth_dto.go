package repository

type GoogleOauth2UserInfoResp map[string]interface{}

type TokenReq struct {
	Holder string `json:"holder,omitempty"`
}

type TokenResp struct {
	Token string `json:"token,omitempty"`
}
