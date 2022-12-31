package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Handlers   []gin.HandlerFunc
	MainEngine = MakeNewEngine()
	RS         = NewServers(MainEngine)
	Router     = RS.Group("")

	Ports = ":3040"

	//模拟资源
	vdata = []map[string]string{{"name": "infinsh war III", "id": "en00029", "times": "180min", "path": "video.asdaliyun.com/oss/usubda"},
		{"name": "infinsh war II", "id": "en00028", "times": "170min", "path": "video.asdaliyun.com/oss/usubde"}}
)

// 限制在写入磁盘之前多少内存占用 内存使用
func MakeNewEngine() *gin.Engine {
	me := gin.New()
	me.MaxMultipartMemory = 2000 << 20

	return me
}

// ServerGroup  服务和路由分组用到
type ServerGroup struct {
	*gin.Engine
}

// 作为Server的构造器  GroupServer的构造器 最后用作 链式调用
func NewServers(e *gin.Engine) *ServerGroup {
	s := &ServerGroup{Engine: MakeNewEngine()}
	if e != nil {
		s.Engine = e
	}
	return s
}

// Routers 内置服务路由，用于健康检查的ping 和 内置记录器 Recovery将 recover任何panic，如果有panic 将写入500
func (that *ServerGroup) Routers() {

	that.Use(gin.Logger())
	that.Use(gin.Recovery())

	that.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}

// 路由注册
func SetupRouters(sg *gin.RouterGroup) {

	//组名称
	lives := sg.Group("/resource")
	lives.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success", "data": "welcome!", "code": 200,
		})
	})
	//查询列表
	lives.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success", "data": vdata, "code": 200,
		})
	})

	//具体信息
	lives.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var data map[string]string
		for _, v := range vdata {
			if id == v["id"] {
				data = v
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success", "data": data, "code": 200,
		})
	})

}

// 注册路由
func NewRouter() http.Handler {
	RS.Routers()
	SetupRouters(Router)

	return RS
}

// 创建服务 HTTPServer，并处理可能的超时，绑定路由到服务引擎
func NewHttpServer(handler http.Handler) *http.Server {

	return &http.Server{
		Addr:           Ports,
		Handler:        handler,
		ReadTimeout:    500 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1000 << 20,
	}
}

// 启动服务
func Start() {
	var Servers = NewHttpServer(NewRouter())
	log.Fatalf(Servers.ListenAndServe().Error())
}

func main() {
	Start()
}
