package tests

import (
	"chatRoom/models/entitys"
	"chatRoom/util/mysql"
	"fmt"
	"log"
	"testing"
)

func TestMysqlBuilder(t *testing.T) {
	tt := new(entitys.UserEntity)
	builder := mysql.NewSqlBuilderWithTable("user")
	result, _ := builder.Where("idUser","=",1).All()
	_ = result.Scan(&tt.IdUser, &tt.WechatAvatar, &tt.WechatAvatar, &tt.Openid, &tt.CreatedAt, &tt.LoginAt, &tt.Sex)
	fmt.Println(tt)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
