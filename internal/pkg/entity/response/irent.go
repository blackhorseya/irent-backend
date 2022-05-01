package response

// IRentResponse declare irent api return response format
type IRentResponse struct {
	Result       string      `json:"Result"`
	ErrorCode    string      `json:"ErrorCode"`
	NeedRelogin  int         `json:"NeedRelogin"`
	NeedUpgrade  int         `json:"NeedUpgrade"`
	Errormessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
}
