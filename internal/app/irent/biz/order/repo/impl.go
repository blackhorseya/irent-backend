package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/base/timex"
	"github.com/blackhorseya/irent/internal/pkg/entity/car"
	"github.com/blackhorseya/irent/internal/pkg/entity/order"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/internal/pkg/infra/transports/restclient"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"strings"
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

func (i *impl) QueryBookings(ctx contextx.Contextx, from *user.Profile) (orders []*order.Info, err error) {
	uri, err := url.Parse(i.o.Endpoint + "/BookingQuery")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), nil)
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

	var data *queryBookingsResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.ErrorMessage != "Success" {
		return nil, errors.New(data.ErrorMessage)
	}

	var ret []*order.Info
	for _, o := range data.Data.OrderObj {
		info := &order.Info{
			No: o.OrderNo,
			Car: &car.Info{
				ID:          strings.ReplaceAll(o.CarNo, " ", ""),
				Type:        "",
				TypeName:    "",
				Area:        "",
				ProjectName: "",
				ProjectID:   "",
				Latitude:    o.CarLatitude,
				Longitude:   o.CarLongitude,
				Seat:        0,
				Distance:    0,
			},
			StartedAt:  timex.ParseString2Time(o.StartTime),
			EndAt:      timex.ParseString2Time(o.StopTime),
			StopPickAt: timex.ParseString2Time(o.StopPickTime),
		}

		ret = append(ret, info)
	}

	return ret, nil
}

func (i *impl) Book(ctx contextx.Contextx, id, projID string, from *user.Profile) (info *order.Booking, err error) {
	uri, err := url.Parse(i.o.Endpoint + "/Booking")
	if err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(&bookReq{ProjID: projID, EDate: "", SDate: "", CarNo: id})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewBuffer(payload))
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

	var data *bookResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.ErrorMessage != "Success" {
		return nil, errors.New(data.ErrorMessage)
	}

	return &order.Booking{
		No:         data.Data.OrderNo,
		LastPickAt: timex.ParseYYYYMMddHHmmss(data.Data.LastPickTime),
		CarID:      id,
		ProjID:     projID,
	}, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, id string, from *user.Profile) (err error) {
	uri, err := url.Parse(i.o.Endpoint + "/BookingCancel")
	if err != nil {
		return nil
	}
	payload, _ := json.Marshal(&cancelBookingReq{OrderNo: id})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	resp, err := i.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data *cancelBookingResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	if data.ErrorMessage != "Success" {
		return errors.New(data.ErrorMessage)
	}

	return nil
}
