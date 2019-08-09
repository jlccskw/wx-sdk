package main

import (
	"fmt"
	"wx-sdk/login"
)

func main() {
	w := login.WxConfig{
		AppID:  "wxa0d453600377b58a",
		Secret: "46d7252c5e8b807ae415e165490dd4e7",
	}
	code := "071FY5U40Yl4bD1x9UW40ddKT40FY5U"
	loginData, err := w.WexLogin(code)
	fmt.Println("data:", loginData)
	//fmt.Println("data.Openid:", loginData.Openid)
	//fmt.Println("data.SessionKey:", loginData.SessionKey)
	fmt.Println("err:", err)

}
