package main

import (
	"fmt"
	"go-web/model"
	"go-web/controller"
	//"go-web/util/token"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    "net/http"
    "github.com/gin-contrib/cors"


)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect PG Database")
	}
	fmt.Println("PG db Connection Successful")
    modelHandler := model.ModelHandler{Db:db}
    err = modelHandler.Init()
	if err != nil {
		panic("Failed to Init PG Database")
	}

	/* Gin Basic Usage ping/pong */

	r := gin.Default()

    //r.Use(CORSMiddleware())
    r.Use(cors.New(cors.Config{
     AllowOrigins:     []string{"*"},
     AllowMethods:     []string{"POST", "GET","DELETE","PUT"},
     AllowHeaders:     []string{"Origin","Authorization"},
     ExposeHeaders:    []string{"Content-Length"},
     AllowCredentials: true,
    }))


    protected := r.Group("/admin")
    protected.Use(controller.JwtAuthMiddleware());
    protected.GET("/user", modelHandler.FetchUser);

	r.POST("/upload", modelHandler.UploadFiles)

	r.POST("/user", modelHandler.CreateUser)
	r.POST("/login", modelHandler.Login)
    r.GET("/user", modelHandler.FetchUser)
    r.DELETE("/user", modelHandler.DeleteUser)
    r.PUT("/user", modelHandler.UpdateUser)
    r.GET("/user/list",func(c *gin.Context) {
         data, e := modelHandler.FetchAllUser()
        if e != nil {
            //c.HTML(http.StatusInternalServerError,"user/create.tmpl", gin.H{
            c.HTML(http.StatusInternalServerError,"userList", gin.H{
            "error" : "Fetch User Failed",
            })
        } else {
            //c.HTML(http.StatusOK,"user/list.tmpl", gin.H{
            c.HTML(http.StatusOK,"userList", gin.H{
                "data": data,
            })
        }
    })

	r.POST("/post", modelHandler.CreatePost)
	r.GET("/post", modelHandler.FetchPost)

    r.POST("/file", modelHandler.SaveFileHandler)
    //Lets use HTML

    //Lets Start to Write new Templates Through html/template module

    r.Static("/assets","./assets")

    r.LoadHTMLGlob("templates/**/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "main", gin.H{
            "title": "Main website",
        })
    })

    /* Guess this is not used.
    r.GET("/user/create",func(c *gin.Context) {
            c.HTML(http.StatusOK,"userCreate",gin.H{});
    })
    */

	r.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")

}
