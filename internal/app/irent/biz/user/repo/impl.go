package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/blackhorseya/irent/internal/app/irent/biz/user/repo/models"
	"github.com/blackhorseya/irent/internal/pkg/base/contextx"
	"github.com/blackhorseya/irent/pb"
	"github.com/google/uuid"
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

func (i *impl) Login(ctx contextx.Contextx, id, password string) (info *pb.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/Login", i.o.Endpoint)
	payload, err := json.Marshal(&models.LoginReq{
		IDNO:       id,
		DeviceID:   uuid.New().String(),
		App:        "1",
		AppVersion: i.o.AppVersion,
		PWD:        password,
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

	var res *models.LoginResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorMessage != "Success" {
		return nil, errors.New(res.ErrorMessage)
	}

	return &pb.Profile{
		Id:          res.Data.UserData.MEMIDNO,
		Name:        res.Data.UserData.MEMCNAME,
		AccessToken: res.Data.Token.AccessToken,
	}, nil
}
