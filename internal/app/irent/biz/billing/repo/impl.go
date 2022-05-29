package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/app/irent/biz/billing/repo/models"
	"github.com/blackhorseya/irent/pb"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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
	o *Options
}

// NewImpl serve caller to create an IRepo
func NewImpl(o *Options) IRepo {
	return &impl{o: o}
}

func (i *impl) QueryArrears(ctx contextx.Contextx, user *pb.Profile) (info *pb.Arrears, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/ArrearsQuery", i.o.Endpoint)
	payload, err := json.Marshal(&models.QueryArrearsReq{IDNO: user.Id})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(timeout, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", user.AccessToken))

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

	var res *models.QueryArrearsResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorMessage != "Success" {
		return nil, errors.New(res.ErrorMessage)
	}

	var records []*pb.ArrearsRecord
	for _, ar := range res.Data.ArrearsInfos {
		r := &pb.ArrearsRecord{
			OrderNo:     ar.OrderNo,
			TotalAmount: int32(ar.TotalAmount),
		}

		records = append(records, r)
	}

	ret := &pb.Arrears{
		Records:     records,
		TotalAmount: int32(res.Data.TotalAmount),
	}

	return ret, nil
}
