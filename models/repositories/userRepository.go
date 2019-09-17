package repositories

import (
	"chatRoom/models/entitys"
	"chatRoom/util/mysql"
	"database/sql"
	"github.com/goinggo/mapstructure"
)

type userRepository struct {
	builder *mysql.SqlBuilder
}

const (
	TABLE = "user"
)

type UserRepository interface {
	Create(map[string]interface{}) (*entitys.UserEntity,error)
	GetUserById(int64) (*entitys.UserEntity,error)
	GetUserByOpenid(string) (*entitys.UserEntity,error)
	UpdateUserById(int64,map[string]interface{}) error
	All() ([]*entitys.UserEntity,error)
}

/**
	初始化数据仓库
 */
func NewUserRepository() UserRepository{
	return &userRepository{
		builder:mysql.NewSqlBuilderWithTable(TABLE),
	}
}

/**
	根据userId取得用户
 */
func (this *userRepository) GetUserById(idUser int64) (*entitys.UserEntity,error){
	var user entitys.UserEntity
	row, err := this.builder.Where("idUser","eq",idUser).First()
	if err != nil {
		return nil,err
	}
	err = row.Scan(&user.IdUser, &user.WechatAvatar, &user.WechatAvatar, &user.Openid, &user.CreatedAt, &user.LoginAt, &user.Sex)
	if err != nil {
		return nil,err
	}
	return &user,nil
}

/**
	获取所有用户
*/
func (this *userRepository) All() ([]*entitys.UserEntity,error){

	users := make([]*entitys.UserEntity,0)
	rows, err := this.builder.All()
	if err != nil {
		return nil,err
	}

	for rows.Next() {
		user := new(entitys.UserEntity)
		err = rows.Scan(&user.IdUser, &user.WechatAvatar, &user.WechatAvatar, &user.Openid, &user.CreatedAt, &user.LoginAt, &user.Sex)
		if err != nil {
			return nil,err
		}
		users = append(users,user)
	}

	return users,nil
}

/**
	创建用户
*/
func (this *userRepository) Create(data map[string]interface{}) (*entitys.UserEntity,error){
	lastId, err := this.builder.Insert(data)
	if err != nil {
		return nil,err
	}
	data["idUser"] = lastId
	var user entitys.UserEntity
	if err := mapstructure.Decode(data, &user); err != nil {
		return nil,err
	}
	return &user,nil
}

/**
	根据openid获取用户信息
 */
func (this *userRepository)GetUserByOpenid(openid string) (*entitys.UserEntity,error){
	var user entitys.UserEntity
	row, err := this.builder.Where("openid","eq",openid).First()
	if err != nil {
		return nil,err
	}
	err = row.Scan(&user.IdUser, &user.WechatName, &user.WechatAvatar, &user.Openid, &user.CreatedAt, &user.LoginAt, &user.Sex)
	if err == sql.ErrNoRows {
		return nil,nil
	} else if err == nil {
		return &user,nil
	} else {
		return nil,err
	}

}

/**
	根据id更新用户信息
 */
func (this *userRepository) UpdateUserById(idUser int64,data map[string]interface{}) error {
	_,err := this.builder.Where("idUser", "eq", idUser).Update(data)
	return err
}
