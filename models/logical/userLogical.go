package logical

import (
	"chatRoom/models/entitys"
	"chatRoom/util/wechatPublic"
	"log"
	"time"
)

func BuildUserInserData(code string) *entitys.UserEntity{
	userInfo := getWechatUserInfo(code)
	if userInfo != nil {
		return nil
	}
	return &entitys.UserEntity{
		WechatName:userInfo.Nickname,
		WechatAvatar:userInfo.HeadImgUrl,
		Openid:userInfo.Openid,
		Sex:userInfo.Sex,
		LoginAt:time.Now().Unix(),
		CreatedAt:time.Now().Unix(),
	}
}
/**
	获取微信侧用户信息
 */
func getWechatUserInfo(code string) *wechatPublic.UserInfo{
	//如果code不存在则重定向获取code
	oauth,err := wechatPublic.GetAccessToken(code)
	if err != nil {
		log.Printf("get error when access_token wechatPublic #%v",err)
		return nil
	}
	userInfo,err := wechatPublic.GetUserInfo(oauth.AccessToken,oauth.Openid)
	if err != nil {
		log.Printf("get error when userInfo wechatPublic #%v",err)
		return nil
	}
	return userInfo
}
