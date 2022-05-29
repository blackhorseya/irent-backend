package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/irent/internal/pkg/base/timex"
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
	"github.com/blackhorseya/irent/pb"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
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
	o *Options
}

// NewImpl serve caller to create an IRepo
func NewImpl(o *Options) IRepo {
	return &impl{o: o}
}

func (i *impl) QueryBookings(ctx contextx.Contextx, from *user.Profile) (orders []*pb.OrderInfo, err error) {
	url := fmt.Sprintf("%s/BookingQuery", i.o.Endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *QueryBookingsResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.ErrorMessage != "Success" {
		return nil, errors.New(data.ErrorMessage)
	}

	var ret []*pb.OrderInfo
	for _, o := range data.Data.OrderObj {
		info := &pb.OrderInfo{
			No:           o.OrderNo,
			CarId:        strings.ReplaceAll(o.CarNo, " ", ""),
			CarLatitude:  float32(o.CarLatitude),
			CarLongitude: float32(o.CarLongitude),
			StartAt:      timex.ParseString2Time(o.StartTime).UnixNano(),
			EndAt:        timex.ParseString2Time(o.StopTime).UnixNano(),
			StopPickAt:   timex.ParseString2Time(o.StopPickTime).UnixNano(),
		}

		ret = append(ret, info)
	}

	return ret, nil
}

func (i *impl) Book(ctx contextx.Contextx, id, projID string, from *user.Profile) (info *pb.Booking, err error) {
	url := fmt.Sprintf("%s/Booking", i.o.Endpoint)
	payload, _ := json.Marshal(&BookReq{ProjID: projID, EDate: "", SDate: "", CarNo: id})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *BookResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.ErrorMessage != "Success" {
		return nil, errors.New(data.ErrorMessage)
	}

	return &pb.Booking{
		No:         data.Data.OrderNo,
		LastPickAt: timex.ParseYYYYMMddHHmmss(data.Data.LastPickTime).UnixNano(),
	}, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, id string, from *user.Profile) (err error) {
	url := fmt.Sprintf("%s/BookingCancel", i.o.Endpoint)
	payload, _ := json.Marshal(&CancelBookingReq{OrderNo: id})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", from.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data *CancelBookingResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	if data.ErrorMessage != "Success" {
		return errors.New(data.ErrorMessage)
	}

	return nil
}
