package repo

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"strings"
)

type listReq struct {
	Radius    float64 `json:"Radius"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	ShowAll   int     `json:"ShowALL"`
}

type respOfCar struct {
	CarNo          string  `json:"CarNo"`
	CarType        string  `json:"CarType"`
	CarTypeName    string  `json:"CarTypeName"`
	CarOfArea      string  `json:"CarOfArea"`
	ProjectName    string  `json:"ProjectName"`
	Rental         float64 `json:"Rental"`
	Mileage        float64 `json:"Mileage"`
	Insurance      int     `json:"Insurance"`
	InsurancePrice int     `json:"InsurancePrice"`
	ShowSpecial    int     `json:"ShowSpecial"`
	SpecialInfo    string  `json:"SpecialInfo"`
	Latitude       float64 `json:"Latitude"`
	Longitude      float64 `json:"Longitude"`
	Operator       string  `json:"Operator"`
	OperatorScore  float64 `json:"OperatorScore"`
	CarTypePic     string  `json:"CarTypePic"`
	Seat           int     `json:"Seat"`
	ProjID         string  `json:"ProjID"`
}

type listResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	Errormessage string `json:"ErrorMessage"`
	Data         struct {
		AnyRentObj []*respOfCar `json:"AnyRentObj"`
	} `json:"Data"`
}

func newCarFromResp(obj *respOfCar) *car.Info {
	return &car.Info{
		ID:          strings.ReplaceAll(obj.CarNo, " ", ""),
		Type:        obj.CarType,
		TypeName:    obj.CarTypeName,
		Area:        obj.CarOfArea,
		ProjectName: obj.ProjectName,
		ProjectID:   obj.ProjID,
		Latitude:    obj.Latitude,
		Longitude:   obj.Longitude,
		Seat:        obj.Seat,
		Distance:    0,
	}
}
