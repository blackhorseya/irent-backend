package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient"
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

func (i *impl) QueryArrears(ctx contextx.Contextx, from *user.Profile) (info *user.Arrears, err error) {
	uri, err := url.Parse(i.o.Endpoint + "/ArrearsQuery")
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(&queryArrearsReq{IDNO: from.ID})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *queryArrearsResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.ErrorMessage != "Success" {
		return nil, errors.New(data.ErrorMessage)
	}

	var records []*user.ArrearsRecord
	for _, ar := range data.Data.ArrearsInfos {
		r := &user.ArrearsRecord{
			OrderNo:     ar.OrderNo,
			TotalAmount: ar.TotalAmount,
		}

		records = append(records, r)
	}

	ret := &user.Arrears{
		Records:     records,
		TotalAmount: data.Data.TotalAmount,
	}

	return ret, nil
}
