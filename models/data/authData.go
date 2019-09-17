package data

import (
	"chatRoom/models/entitys"
	"chatRoom/util"
	"chatRoom/util/redis"
	"time"
)

type AuthData interface {
	GetUserByToken(string) map[string]string
	GenerateToken(*entitys.UserEntity) string
}

type authData struct {

}

func NewAuthData() *authData {
	return new(authData)
}

/**
	根据token获取用户信息
*/
func (*authData)GetUserByToken(token string) map[string]string {
	if token == "" {
		return nil
	}
	userData,_ := redis.Connection.HGetAll(token).Result()
	return userData
}

/**
	颁发token
 */
func (*authData)GenerateToken(user *entitys.UserEntity) string {
	token := util.GetRandom(30) //生成30位token
	go prepareSessionData(token,user)
	return token
}

/**
	生成session数据
 */
func prepareSessionData(token string,user *entitys.UserEntity) {
	tx := redis.Connection.TxPipeline()
	_, _ = tx.HMSet(token, util.StructToMap(*user)).Result()
	_, _ = tx.Expire(token, time.Second*7200).Result()
	_, _ = tx.Exec()
}

