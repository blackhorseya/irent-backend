package repo

// QueryBookingsResp declare query bookings response struct
type QueryBookingsResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		OrderObj []struct {
			StationInfo struct {
				StationID           string        `json:"StationID"`
				StationName         string        `json:"StationName"`
				Tel                 string        `json:"Tel"`
				ADDR                string        `json:"ADDR"`
				Latitude            float64       `json:"Latitude"`
				Longitude           float64       `json:"Longitude"`
				Content             string        `json:"Content"`
				IsRent              interface{}   `json:"IsRent"`
				ContentForAPP       string        `json:"ContentForAPP"`
				IsRequiredForReturn int           `json:"IsRequiredForReturn"`
				StationPic          []interface{} `json:"StationPic"`
			} `json:"StationInfo"`
			Operator          string      `json:"Operator"`
			OperatorScore     float64     `json:"OperatorScore"`
			CarTypePic        string      `json:"CarTypePic"`
			CarNo             string      `json:"CarNo"`
			CarBrend          string      `json:"CarBrend"`
			CarTypeName       string      `json:"CarTypeName"`
			Seat              int         `json:"Seat"`
			ParkingSection    string      `json:"ParkingSection"`
			IsMotor           int         `json:"IsMotor"`
			CarOfArea         string      `json:"CarOfArea"`
			CarLatitude       float64     `json:"CarLatitude"`
			CarLongitude      float64     `json:"CarLongitude"`
			MotorPowerBaseObj interface{} `json:"MotorPowerBaseObj"`
			ProjType          int         `json:"ProjType"`
			ProjName          string      `json:"ProjName"`
			WorkdayPerHour    int         `json:"WorkdayPerHour"`
			HolidayPerHour    int         `json:"HolidayPerHour"`
			MaxPrice          int         `json:"MaxPrice"`
			MaxPriceH         int         `json:"MaxPriceH"`
			MotorBasePriceObj interface{} `json:"MotorBasePriceObj"`
			OrderStatus       int         `json:"OrderStatus"`
			OrderNo           string      `json:"OrderNo"`
			StartTime         string      `json:"StartTime"`
			PickTime          string      `json:"PickTime"`
			ReturnTime        string      `json:"ReturnTime"`
			StopPickTime      string      `json:"StopPickTime"`
			StopTime          string      `json:"StopTime"`
			OpenDoorDeadLine  string      `json:"OpenDoorDeadLine"`
			CarRentBill       int         `json:"CarRentBill"`
			MileagePerKM      float64     `json:"MileagePerKM"`
			MileageBill       int         `json:"MileageBill"`
			Insurance         int         `json:"Insurance"`
			InsurancePerHour  int         `json:"InsurancePerHour"`
			InsuranceBill     int         `json:"InsuranceBill"`
			TransDiscount     int         `json:"TransDiscount"`
			Bill              int         `json:"Bill"`
			DailyMaxHour      int         `json:"DailyMaxHour"`
			CARMGTSTATUS      int         `json:"CAR_MGT_STATUS"`
			AppStatus         int         `json:"AppStatus"`
		} `json:"OrderObj"`
	} `json:"Data"`
}

// CancelBookingReq declare cancel booking request struct
type CancelBookingReq struct {
	OrderNo string `json:"OrderNo"`
}

// CancelBookingResp declare cancel booking response struct
type CancelBookingResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
	} `json:"Data"`
}

// BookReq declare book a car request struct
type BookReq struct {
	ProjID string `json:"ProjID"`
	EDate  string `json:"EDate"`
	SDate  string `json:"SDate"`
	CarNo  string `json:"CarNo"`
}

// BookResp declare book a car response struct
type BookResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		OrderNo      string `json:"OrderNo"`
		LastPickTime string `json:"LastPickTime"`
	} `json:"Data"`
}
