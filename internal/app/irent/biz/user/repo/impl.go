package repo

import (
	"bytes"
	"encoding/json"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

// Options declare app's configuration
type Options struct {
	Endpoint   string `json:"endpoint" yaml:"endpoint"`
	AppVersion string `json:"appVersion" yaml:"appVersion"`
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

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *user.Profile, err error) {
	uri, err := url.Parse(i.o.Endpoint + "/Login")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(&loginReq{
		IDNO:       id,
		DeviceID:   uuid.New().String(),
		App:        "1",
		AppVersion: i.o.AppVersion,
		PWD:        password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *loginResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.ErrorMessage != "Success" {
		return nil, errors.New(data.ErrorMessage)
	}

	return newProfileFromResp(data), nil
}
