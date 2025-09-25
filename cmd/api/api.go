package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Albert-tru/ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string  //服务器监听地址，如":8080"
	db   *sql.DB //数据库连接对象
}

// 创建服务器实例
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// 运行服务器
func (s *APIServer) Run() error {

	// 创建一个新的路由器 （路由器是用来管理“请求路径和处理函数的映射关系”的）
	router := mux.NewRouter()
	// 创建带前缀的子路由器
	subrouter := router.PathPrefix("/api/v1").Subrouter() //只处理以 /api/v1 开头的请求【api版本化】

	userStore := user.NewStore(s.db) //创建用户存储对象，传入数据库连接

	// 创建专门处理用户相关接口的 handler，并注册路由
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter) //把用户相关的路由注册到子路由器上

	//	启动服务器前，打印一条日志
	log.Println("listening on", s.addr)

	// 启动 HTTP 服务，开始监听端口
	return http.ListenAndServe(s.addr, router)
}
