package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

//分组测试
//func main() {
//	r := gee.New()
//	r.GET("/index", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello index page</h1>")
//	})
//	v1 := r.Group("/v1")
//	{
//		v1.GET("/", func(c *gee.Context) {
//			c.HTML(http.StatusOK, "<h1>hello gee</h1>")
//		})
//		//	/hello?name=zy
//		v1.GET("/hello", func(c *gee.Context) {
//			c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
//		})
//	}
//	v2 := r.Group("/v2")
//	{
//		v2.GET("/hello/:name", func(c *gee.Context) {
//
//			c.String(http.StatusOK, "hello %s ,you are at %s\n", c.Param("name"), c.Path)
//		})
//		v2.POST("/login", func(c *gee.Context) {
//			c.JSON(http.StatusOK, gee.H{
//				"username": c.PostForm("username"),
//				"password": c.PostForm("password"),
//			})
//		})
//	}
//
//	r.Run(":9999")
//}

//r.POST("/login", func(c *gee.Context) {
//	c.JSON(http.StatusOK, gee.H{
//		"username": c.PostForm("username"),
//		"password": c.PostForm("password"),
//	})
//})
//func indexHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "URL.PATH=%q\n", req.URL.Path)
//}
//func helloHander(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
//	}
//}

// 中间件
func onlyForv2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server 错误\n")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello gee&nbsp</h1>")
	})
	v2 := r.Group("/v2")
	v2.Use(onlyForv2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s,you are at %s", c.Param("name"), c.Req.RequestURI)
		})
	}

	r.Run(":9999")
}
