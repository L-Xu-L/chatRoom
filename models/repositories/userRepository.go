package repositories

import (
	"chatRoom/models/entitys"
	"chatRoom/util/mysql"
	"github.com/goinggo/mapstructure"
)

type userRepository struct {
	builder *mysql.SqlBuilder
}

const (
	TABLE = "user"
)


// MovieRepository处理电影实体/模型的基本操作。
//它是一个可测试的接口，即一个内存电影库或连接到sql数据库。
type UserRepository interface {
	Create(map[string]interface{}) (*entitys.UserEntity,error)
	GetUserById(idUser int64) (*entitys.UserEntity,error)
	GetUserByOpenid(string) (*entitys.UserEntity,error)
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
	var userEntity entitys.UserEntity
	if err := mapstructure.Decode(data, &userEntity); err != nil {
		return nil,err
	}
	return &userEntity,nil
}

func (this *userRepository)GetUserByOpenid(openid string) (*entitys.UserEntity,error){
	var user entitys.UserEntity
	row, err := this.builder.Where("openid","eq",openid).First()
	if err != nil {
		return nil,err
	}
	err = row.Scan(&user.IdUser, &user.WechatAvatar, &user.WechatAvatar, &user.Openid, &user.CreatedAt, &user.LoginAt, &user.Sex)
	if err != nil {
		return nil,err
	}
	return &user,nil
}