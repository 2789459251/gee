package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello index page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>hello gee</h1>")
		})
		//	/hello?name=zy
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {

			c.String(http.StatusOK, "hello %s ,you are at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}

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
