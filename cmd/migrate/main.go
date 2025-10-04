package main

//数据库迁移的入口

import (
        "log"
        "os"

        "github.com/Albert-tru/ecom/db"

        "github.com/Albert-tru/ecom/config"
        mysqlCfg "github.com/go-sql-driver/mysql"
        "github.com/golang-migrate/migrate/v4"
        "github.com/golang-migrate/migrate/v4/database/mysql"
        _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
        //读取数据库的连接配置
        db, err := db.NewMySQLStorage(mysqlCfg.Config{
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

        //创建一个新的数据库驱动实例
        driver, err := mysql.WithInstance(db, &mysql.Config{})
        if err != nil {
                log.Fatal(err)
        }

        //创建一个新的迁移实例
        m, err := migrate.NewWithDatabaseInstance(
                "file://cmd/migrate/migrations",
                "mysql",
                driver,
        )

        if err != nil {
                log.Fatal(err)
        }

        //根据命令行参数执行相应的迁移操作
        cmd := os.Args[(len(os.Args) - 1)]
        switch cmd {
        case "up":
                //执行向上迁移
                if err := m.Up(); err != nil && err != migrate.ErrNoChange {
                        log.Fatal(err)
                }
                log.Println("Successfully applied up migrations")
        case "down":
                //执行向下迁移
                if err := m.Down(); err != nil && err != migrate.ErrNoChange {
                        log.Fatal(err)
                }
                log.Println("Successfully applied down migrations")
        default:
                log.Fatalf("unknown command: %s", cmd)
        }

}
