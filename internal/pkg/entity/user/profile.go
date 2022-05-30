package user

import (
	"github.com/blackhorseya/irent/pb"
)

// Profile declare user profile
type Profile struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

// NewProfileResponse return *pb.Profile
func NewProfileResponse(from *Profile) *pb.Profile {
	return &pb.Profile{
		Id:          from.ID,
		Name:        from.Name,
		AccessToken: from.AccessToken,
	}
}
