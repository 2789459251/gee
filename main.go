package main

import (
	"gee"
	"net/http"
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
//func onlyForv2() gee.HandlerFunc {
//	return func(c *gee.Context) {
//		t := time.Now()
//		c.Fail(500, "Internal Server 错误\n")
//		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
//	}
//}
//func main() {
//	r := gee.New()
//	r.Use(gee.Logger())
//	r.GET("/", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "<h1>hello gee&nbsp</h1>")
//	})
//	v2 := r.Group("/v2")
//	v2.Use(onlyForv2())
//	{
//		v2.GET("/hello/:name", func(c *gee.Context) {
//			c.String(http.StatusOK, "hello %s,you are at %s", c.Param("name"), c.Req.RequestURI)
//		})
//	}
//
//	r.Run(":9999")
//}

//		template文件
//type student struct {
//	Name string
//	Age  int8
//}
//
//func FormatAsDate(t time.Time) string {
//	year, mouth, day := t.Date()
//	return fmt.Sprintf("%d-%02d-%02d", year, mouth, day)
//}
//func main() {
//	r := gee.New()
//	r.Use(gee.Logger())
//	r.SetFuncMap(template.FuncMap{
//		"FormatAsDate": FormatAsDate,
//	})
//	r.LoadHTMLGlob("templates/*")
//	r.Static("/assets", "./static")
//
//	stu1 := &student{Name: "zy", Age: 19}
//	stu2 := &student{Name: "sy", Age: 29}
//
//	r.GET("/", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "css.tmpl", nil)
//	})
//	r.GET("/students", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
//			"title":  "gee",
//			"stuArr": []*student{stu2, stu1},
//		})
//	})
//	r.GET("/date", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
//			"title": "gee",
//			"now":   time.Date(2024, 3, 7, 0, 0, 0, 0, time.UTC),
//		})
//	})
//	r.Run(":9999")
//}

//func main() {
//	r := gee.New()
//	r.Static("/assets", "./static")
//	r.Run(":9999")
//}

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello  geektutu")
	})
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, "hello %s", names[100])
	})
	r.Run(":9999")
}
