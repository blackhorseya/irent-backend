package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/spf13/viper"
	"net/http"
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

func (i *impl) List(ctx contextx.Contextx) (cars []*car.Info, err error) {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
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

	var list *listResp
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	var ret []*car.Info
	for _, obj := range list.Data.AnyRentObj {
		ret = append(ret, newCarFromResp(obj))
	}

	return ret, nil
}
