package pgmodel

import (
	"fmt"
    "encoding/json"
    "net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type ModelHandler struct{
    Db *gorm.DB
}

type user struct {
	gorm.Model
    Userid   string  `gorm:"uniqueIndex"`
	Name     string 
    Level    int 
    Password string 
    Posts []post  `gorm:"foreignKey:Writer;references:Userid"`
}


func (mh *ModelHandler) Init() error{
    err := mh.Db.AutoMigrate(&user{})
    if err != nil{
        panic("DB User Migration Failed")
    }
    err = mh.Db.AutoMigrate(&post{})
    if err != nil{
        panic("DB Post Migration Failed")
    }
    return err
}

func (mh *ModelHandler) FetchAllUser() ([]user,error){
        var users []user
        result := mh.Db.Find(&users)
        return users, result.Error
}

func (u user) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "userid": u.Userid,
        "name": u.Name,
        "level": u.Level,
    })
}
func (mh *ModelHandler) CreateUser(c *gin.Context){
    u := user{}
      if err:=c.BindJSON(&u); err != nil{
         c.JSON(http.StatusInternalServerError, gin.H {
            "message" : err.Error(),
        })
         return 
      }
      fmt.Println("Request: ", u)
     // u := user{Userid: json.Userid,Name: json.Name,Level:json.Level,Password: json.Password}
      result := mh.Db.Create(&u)
      if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "message" : result.Error.Error(),
        })
        return
      } 
      fmt.Println("Created User : ", u)

      c.JSON(http.StatusOK, gin.H{
          "message": "Create User Successful",
          "Userid" : u.Userid,
          "Name" : u.Name,
      })
}

func (mh *ModelHandler) FetchUser(c *gin.Context){
    userid := c.Query("userid")
    if userid != ""{
        //Userid is specified.
        var user user
        result := mh.Db.First(&user,"userid = ?",userid)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : result.Error.Error(),
            })
        } else {
            c.JSON(http.StatusOK, user)
        }
    } else {
        // Fetch All Users
        var users []user
        result := mh.Db.Find(&users)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : result.Error.Error(),
            })
        } else {
            c.JSON(http.StatusOK, users)
        }
    }
}

func (mh *ModelHandler) UpdateUser(c *gin.Context){
    json := user{}
    if err:=c.BindJSON(&json); err != nil{
        c.JSON(http.StatusInternalServerError, gin.H {
        "message" : err.Error(),
        })
        return 
    }
    fetchUser := user{} 
    result := mh.Db.First(&fetchUser, "userid=?", json.Userid)
    if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : result.Error.Error(),
            })
            return
        }

    result = mh.Db.Model(&fetchUser).Updates(&json)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "message" : result.Error.Error(),
        })
    } else {
        c.JSON(http.StatusOK, json)
    }
}

func (mh *ModelHandler) DeleteUser(c *gin.Context){
    json := user{}
    if err:=c.BindJSON(&json); err != nil{
        c.JSON(http.StatusInternalServerError, gin.H {
        "message" : err.Error(),
        })
        return 
    }

    fetchUser := user{} 
    result := mh.Db.First(&fetchUser, "userid=?", json.Userid)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "message" : result.Error.Error(),
        })
        return
     }

    result = mh.Db.Delete(&fetchUser)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "message" : result.Error.Error(),
        })
    } else {
        c.JSON(http.StatusOK, json)
    }
}
