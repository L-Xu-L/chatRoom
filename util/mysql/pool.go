package mysql
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	pool *sql.DB
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     =  3306
	DATABASE = "chatroom"
)

func NewPool() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB,err := sql.Open("mysql",dsn)
	if err != nil{
		log.Fatalf("Open mysql failed,err:%v\n",err)
		return
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数
	pool = DB
}