package main

import (
    "encoding/json"
    //"log"
    "fmt"
    "bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-web/model"
	//"go-web/util/geo"
	//"go-web/controller"
    //"github.com/gin-contrib/cors"
)



func TestGinPosts(t *testing.T) {
/*
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect PG Database")
	}
    modelHandler := model.ModelHandler{Db:db}
*/
    err, modelHandler := model.ModelInit("test_env.json")
	if err != nil {
		panic("failed to init model handler")
	}
	router := SetupRestRouter(&modelHandler)

    var jsonData []byte//json throughout this test 

    //Create User 
	w := httptest.NewRecorder()
    jsonData = []byte(`{
                "userid" : "test_userid",
                "name" : "test_user",
                "password" : "test_password"
	}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	fmt.Println("Create : ", w.Body.String())

    // Fetch User
	w = httptest.NewRecorder()

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	fmt.Println("Login : ",  w.Body.String())

    //type for login
    type LoginResponse struct {
        Message string `json:"message"`
        Token string `json:"token"`
    }
    // Login
    loginData := LoginResponse{}
    if err := json.Unmarshal([]byte(w.Body.String()), &loginData); err != nil {
        panic(err)
    }

    //Before Call User
	req, _ = http.NewRequest("GET", "/admin/user", nil)

    authString := "Bearer " + loginData.Token
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", authString)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	fmt.Println("Fetch : ",  w.Body.String())

    //delete 
    jsonData = []byte(`{
                "userid" : "test_userid"
	}`)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/user", bytes.NewBuffer(jsonData))
	//w := httptest.NewRecorder()
	//req, _ := http.NewRequest("DELETE", "/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	fmt.Println("Delete: ",  w.Body.String())

    // Get To Fetch All Post 
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/post", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

    //geo.GetLocation()
    // Post on Post
    /*
    var jsonData = []byte(`{
                "Category" : "test_category",
                "Subject" : "test_subject",
                "Content" : "test_content",
                "Writer": "jeffjang" 
	}`)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)
    fmt.Println(w)
	assert.Equal(t, 200, w.Code)
    */

    // Get To Fetch Specific Post  -- current create post does not return postid, so need to add
    /*
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/post", nil)
    q := req.URL.Query()
    q.Add
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
    */
//delete post for tset
}
