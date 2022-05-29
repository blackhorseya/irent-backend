package car

// Info declare car information
type Info struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	TypeName    string  `json:"type_name"`
	Area        string  `json:"area"`
	ProjectName string  `json:"project_name"`
	ProjectID   string  `json:"project_id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Seat        int     `json:"seat"`
	Distance    float64 `json:"distance"`
}
