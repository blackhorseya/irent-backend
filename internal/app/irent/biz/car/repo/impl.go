package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/spf13/viper"
)

// Options declare app's configuration
type Options struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, err
	}

	return o, nil
}

type impl struct {
	o *Options
}

// NewImpl serve caller to create an IRepo
func NewImpl(o *Options) IRepo {
	return &impl{o: o}
}

type listReq struct {
	Radius    float64 `json:"Radius"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	ShowAll   int     `json:"ShowALL"`
}

type listResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	Errormessage string `json:"ErrorMessage"`
	Data         struct {
		AnyRentObj []struct {
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
		} `json:"AnyRentObj"`
	} `json:"Data"`
}

func (i *impl) List(ctx contextx.Contextx) (cars []*pb.Car, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/AnyRent", i.o.Endpoint)
	payload, err := json.Marshal(&listReq{
		Radius:    1.5,
		Latitude:  0,
		Longitude: 0,
		ShowAll:   1,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(timeout, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var list *listResp
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}

	var ret []*pb.Car
	for _, obj := range list.Data.AnyRentObj {
		ret = append(ret, &pb.Car{
			Id:          strings.ReplaceAll(obj.CarNo, " ", ""),
			CarType:     obj.CarType,
			CarTypeName: obj.CarTypeName,
			CarOfArea:   obj.CarOfArea,
			ProjectName: obj.ProjectName,
			ProjectId:   obj.ProjID,
			Latitude:    obj.Latitude,
			Longitude:   obj.Longitude,
			Seat:        int64(obj.Seat),
		})
	}

	return ret, nil
}
