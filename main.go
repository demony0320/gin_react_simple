package main

import (
	"fmt"
	"go-web/pgmodel"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"  
    "net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect PG Database")
	}
	fmt.Println("PG db Connection Successful")
    modelHandler := pgmodel.ModelHandler{Db:db}
    err = modelHandler.Init()
	if err != nil {
		panic("Failed to Init PG Database")
	}

	/* Gin Basic Usage ping/pong */

	r := gin.Default()

    r.GET("/user", modelHandler.FetchUser)
	r.POST("/user", modelHandler.CreateUser)
    r.DELETE("/user", modelHandler.DeleteUser)
    r.PUT("/user", modelHandler.UpdateUser)

	r.POST("/post", modelHandler.CreatePost)
    //Lets use HTML

    //Lets Start to Write new Templates Through html/template module

    r.Static("/assets","./assets")

    r.LoadHTMLGlob("templates/**/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "main", gin.H{
            "title": "Main website",
        })
    })
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
    r.GET("/user/create",func(c *gin.Context) {
            c.HTML(http.StatusOK,"userCreate",gin.H{});
    })
	r.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")

}
