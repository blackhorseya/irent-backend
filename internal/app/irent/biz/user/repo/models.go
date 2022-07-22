package repo

import (
	"github.com/blackhorseya/irent/internal/pkg/entity/user"
)

type loginReq struct {
	IDNO       string `json:"IDNO"`
	DeviceID   string `json:"DeviceID"`
	App        string `json:"app"`
	AppVersion string `json:"appVersion"`
	PWD        string `json:"PWD"`
}

type loginResp struct {
	Result       string `json:"Result"`
	ErrorCode    string `json:"ErrorCode"`
	NeedRelogin  int    `json:"NeedRelogin"`
	NeedUpgrade  int    `json:"NeedUpgrade"`
	ErrorMessage string `json:"ErrorMessage"`
	Data         struct {
		Token struct {
			AccessToken      string `json:"Access_token"`
			RefrashToken     string `json:"Refrash_token"`
			RxpiresIn        int    `json:"Rxpires_in"`
			RefrashRxpiresIn int    `json:"Refrash_Rxpires_in"`
		} `json:"Token"`
		UserData struct {
			MEMIDNO        string `json:"MEMIDNO"`
			MEMCNAME       string `json:"MEMCNAME"`
			MEMTEL         string `json:"MEMTEL"`
			MEMHTEL        string `json:"MEMHTEL"`
			MEMBIRTH       string `json:"MEMBIRTH"`
			MEMAREAID      int    `json:"MEMAREAID"`
			MEMADDR        string `json:"MEMADDR"`
			MEMEMAIL       string `json:"MEMEMAIL"`
			MEMCOMTEL      string `json:"MEMCOMTEL"`
			MEMCONTRACT    string `json:"MEMCONTRACT"`
			MEMCONTEL      string `json:"MEMCONTEL"`
			MEMMSG         string `json:"MEMMSG"`
			CARDNO         string `json:"CARDNO"`
			UNIMNO         string `json:"UNIMNO"`
			MEMSENDCD      int    `json:"MEMSENDCD"`
			CARRIERID      string `json:"CARRIERID"`
			NPOBAN         string `json:"NPOBAN"`
			HasCheckMobile int    `json:"HasCheckMobile"`
			NeedChangePWD  int    `json:"NeedChangePWD"`
			HasBindSocial  int    `json:"HasBindSocial"`
			HasVaildEMail  int    `json:"HasVaildEMail"`
			Audit          int    `json:"Audit"`
			IrFlag         int    `json:"IrFlag"`
			PayMode        int    `json:"PayMode"`
			RentType       int    `json:"RentType"`
			IDPic          int    `json:"ID_pic"`
			DDPic          int    `json:"DD_pic"`
			MOTORPic       int    `json:"MOTOR_pic"`
			AAPic          int    `json:"AA_pic"`
			F01Pic         int    `json:"F01_pic"`
			SignturePic    int    `json:"Signture_pic"`
			SigntureCode   string `json:"SigntureCode"`
			MEMRFNBR       string `json:"MEMRFNBR"`
			SIGNATURE      string `json:"SIGNATURE"`
		} `json:"UserData"`
	} `json:"Data"`
}

func newProfileFromResp(resp *loginResp) *user.Profile {
	return &user.Profile{
		ID:          resp.Data.UserData.MEMIDNO,
		Name:        resp.Data.UserData.MEMCNAME,
		AccessToken: resp.Data.Token.AccessToken,
	}
}
