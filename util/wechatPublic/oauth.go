package wechatPublic

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"io/ioutil"
	"net/http"
)

const (
	appId = "wx941a4cc933c6f96a"
	appSerect = "88dee3d6edaad91ef6a3f152f584340c"
)

/**
	重定向获取code
 */
func GetCodeUrl(redirectUrl,state string) string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=%s#wechat_redirect",appId,redirectUrl,state)
}


type Oauth struct {
	AccessToken string
	ExpiresIn int16
	RefreshToken string
	Openid string
	Scope string
}


func GetAccessToken(code string) (*Oauth,error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",appId,appSerect,code)
	resp,err := http.Get(url)
	if err != nil {
		return nil,errors.New("failed to build get access_token request")
	}

	if resp == nil || resp.Body == nil {
		return nil,errors.New("failed to build get access_token request")
	}

	var responseData []byte
	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,errors.New("failed to translator responseData into []byte")
	}
	defer resp.Body.Close()

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(responseData,&responseMap)
	if err != nil {
		return nil,errors.New("failed to parse json data from responseData")
	}

	if _,exists := responseMap["errcode"]; exists {
		return nil,errors.New(responseMap["errmsg"].(string))
	}

	oAuth := &Oauth{
		AccessToken:responseMap["access_token"].(string),
		ExpiresIn:int16(responseMap["expires_in"].(float64)),
		RefreshToken:responseMap["refresh_token"].(string),
		Openid:responseMap["openid"].(string),
		Scope:responseMap["scope"].(string),
	}

	return oAuth,nil
}

type UserInfo struct {
	Openid string
	Nickname string
	Sex int
	Province string
	City string
	Country string
	HeadImgUrl string
}

func GetUserInfo(accessToken,openid string) (*UserInfo,error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",accessToken,openid)
	resp,err := http.Get(url)
	if err != nil {
		return nil,errors.New("failed to build get userInfo request")
	}

	if resp == nil || resp.Body == nil {
		return nil,errors.New("failed to build get userInfo request")
	}

	var responseData []byte
	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,errors.New("failed to translator responseData into []byte")
	}
	defer resp.Body.Close()

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(responseData,&responseMap)
	if err != nil {
		return nil,errors.New("failed to parse json data from responseData")
	}

	if _,exists := responseMap["errcode"]; exists {
		return nil,errors.New(responseMap["errmsg"].(string))
	}

	userInfo := &UserInfo{
		Openid:responseMap["openid"].(string),
		Nickname:responseMap["nickname"].(string),
		Sex:int(responseMap["sex"].(float64)),
		Province:responseMap["province"].(string),
		City:responseMap["city"].(string),
		Country:responseMap["country"].(string),
		HeadImgUrl:responseMap["headimgurl"].(string),
	}

	return userInfo,nil
}