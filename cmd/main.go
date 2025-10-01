package main

import (
	"database/sql"
	"log"

	"github.com/Albert-tru/ecom/cmd/api"
	"github.com/Albert-tru/ecom/config"
	"github.com/Albert-tru/ecom/db"
	"github.com/go-sql-driver/mysql"
)

// 程序入口
func main() {
	// 添加调试输出
	log.Printf("数据库配置: 用户=%s, 地址=%s, 数据库名=%s",
		config.Envs.DBUser, config.Envs.DBAddress, config.Envs.DBName)

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  config.Envs.DBNet,
		Addr:                 "127.0.0.1:3306",
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorge(db)

	//创建并运行API服务器
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorge(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("无法连接到数据库:", err)
	}

	log.Println("成功连接到数据库")
}
