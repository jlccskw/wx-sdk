package login

import (
	"encoding/json"
	"errors"
	"net/url"
	"wx-sdk/common"
	"wx-sdk/utils"
)

type (
	// WxConfig 微信配置类
	WxConfig struct {
		AppID  string `json:"appid"`  // 微信APPID
		Secret string `json:"secret"` // 微信Secret
	}

	// WXBizDataCrypt 小程序解密密钥信息
	WXBizDataCrypt struct {
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionid"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
)

// WexLogin 微信小程序登录 直接登录获取用户信息
func (m *WxConfig) WexLogin(code string) (data *WXBizDataCrypt, err error) {
	data, err = m.GetJsCode2Session(code)
	if err != nil {
		return nil, err
	}
	return
}

// GetJsCode2Session 获取
func (m *WxConfig) GetJsCode2Session(code string) (wXBizDataCrypt *WXBizDataCrypt, err error) {

	if code == "" {
		return wXBizDataCrypt, errors.New("GetJsCode2Session error: code is null")
	}

	params := url.Values{
		"js_code":    []string{code},
		"grant_type": []string{"authorization_code"},
	}

	t, err := utils.Struct2Map(m)
	if err != nil {
		return wXBizDataCrypt, err
	}

	for k, v := range t {
		params.Set(k, v)
	}
	body, err := utils.NewRequest("GET", common.JsCode2SessionURL, []byte(params.Encode()))
	if err != nil {
		return wXBizDataCrypt, err
	}
	err = json.Unmarshal(body, &wXBizDataCrypt)
	if err != nil {
		return wXBizDataCrypt, err
	}

	if wXBizDataCrypt.ErrMsg != "" {
		return wXBizDataCrypt, errors.New(wXBizDataCrypt.ErrMsg)
	}

	return
}
