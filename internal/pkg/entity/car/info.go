package car

import (
	"github.com/blackhorseya/irent/pb"
)

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

// NewCarResponse return *pb.Car
func NewCarResponse(from *Info) *pb.Car {
	return &pb.Car{
		Id:          from.ID,
		CarType:     from.Type,
		CarTypeName: from.TypeName,
		CarOfArea:   from.Area,
		ProjectName: from.ProjectName,
		ProjectId:   from.ProjectID,
		Latitude:    from.Latitude,
		Longitude:   from.Longitude,
		Seat:        int64(from.Seat),
		Distance:    from.Distance,
	}
}
