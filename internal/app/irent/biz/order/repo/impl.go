package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/blackhorseya/irent/internal/app/irent/biz/order/repo/models"
	"github.com/blackhorseya/irent/internal/pkg/base/contextx"
	"github.com/blackhorseya/irent/internal/pkg/base/timex"
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

func (i *impl) QueryBookings(ctx contextx.Contextx, user *pb.Profile) (orders []*pb.OrderInfo, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/BookingQuery", i.o.Endpoint)
	req, err := http.NewRequestWithContext(timeout, http.MethodPost, url, nil)
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

	var res *models.QueryBookingsResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorMessage != "Success" {
		return nil, errors.New(res.ErrorMessage)
	}

	var ret []*pb.OrderInfo
	for _, o := range res.Data.OrderObj {
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

func (i *impl) Book(ctx contextx.Contextx, id, projID string, user *pb.Profile) (info *pb.Booking, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/Booking", i.o.Endpoint)
	payload, _ := json.Marshal(&models.BookReq{ProjID: projID, EDate: "", SDate: "", CarNo: id})
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

	var res *models.BookResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if res.ErrorMessage != "Success" {
		return nil, errors.New(res.ErrorMessage)
	}

	return &pb.Booking{
		No:         res.Data.OrderNo,
		LastPickAt: timex.ParseYYYYMMddHHmmss(res.Data.LastPickTime).UnixNano(),
	}, nil
}

func (i *impl) CancelBooking(ctx contextx.Contextx, id string, user *pb.Profile) (err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/BookingCancel", i.o.Endpoint)
	payload, _ := json.Marshal(&models.CancelBookingReq{OrderNo: id})
	req, err := http.NewRequestWithContext(timeout, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", user.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res *models.CancelBookingResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	if res.ErrorMessage != "Success" {
		return errors.New(res.ErrorMessage)
	}

	return nil
}
