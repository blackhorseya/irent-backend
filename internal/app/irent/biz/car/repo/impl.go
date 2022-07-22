package repo

import (
	"bytes"
	"encoding/json"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
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
	o      *Options
	client restclient.HTTPClient
}

// NewImpl serve caller to create an IRepo
func NewImpl(o *Options, client restclient.HTTPClient) IRepo {
	return &impl{
		o:      o,
		client: client,
	}
}

func (i *impl) List(ctx contextx.Contextx) (cars []*car.Info, err error) {
	uri, err := url.Parse(i.o.Endpoint + "/AnyRent")
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(&listReq{
		Radius:    1.5,
		Latitude:  0,
		Longitude: 0,
		ShowAll:   1,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *listResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Errormessage != "Success" {
		return nil, errors.New(data.Errormessage)
	}

	var ret []*car.Info
	for _, obj := range data.Data.AnyRentObj {
		ret = append(ret, newCarFromResp(obj))
	}

	return ret, nil
}
