package booking

type reqID struct {
	ID string `uri:"id" binding:"required"`
}

type bookRequest struct {
	ID         string `json:"id"`
	ProjectID  string `json:"project_id"`
	UserID     string `json:"user_id"`
	Circularly bool   `json:"circularly"`
}
