package model

import (
	"fmt"
    "encoding/json"
    "net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
    "io/ioutil"
	"go-web/util/token"
)

type user struct {
	gorm.Model
    Userid   string  `gorm:"uniqueIndex"`
	Name     string
    Level    int
    Password string
    Posts []post  `gorm:"foreignKey:Writer;references:Userid"`
}

func (u user) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "userid": u.Userid,
        "name": u.Name,
        "level": u.Level,
    })
}

func (mh *ModelHandler) FetchAllUser() ([]user,error){
        var users []user
        result := mh.Db.Find(&users)
        return users, result.Error
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
        fmt.Println("Result : ", result)
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
// User default functions end.

func (mh *ModelHandler) Login(c *gin.Context){
        //fmt.Println("At the First of Login ", c.PostForm("userid"))
        req, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            fmt.Println(err.Error())
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : "Incorrect type of Request Body",
            })
            return
        }

        fmt.Printf("ctx.Request.body: %v", string(req))

        var data map[string]interface{}
        json.Unmarshal([]byte(req), &data)
        var userid string
        var pwd string
        var errMessage string
        userid,_= data["userid"].(string)
        pwd,_= data["password"].(string)
        if len(userid) <= 0 || len(pwd) <= 0 {
            fmt.Println("id or password is empty")
                errMessage="Id or Password is Empty"
                /*
                c.JSON(http.StatusInternalServerError, gin.H {
                    "message" : "Id or Password is Empty",
                })
*/
        }
        var usr user
        result := mh.Db.First(&usr,"userid = ?",userid)
        if result.Error != nil {
            errMessage = userid + " not exist in db " + result.Error.Error()
            /*
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : userid + " not Exist , " + result.Error.Error(),
            })
            */
        } else {
            if  usr.Password != pwd {
                errMessage = "password is incorrect"
                /*
                c.JSON(http.StatusInternalServerError, gin.H {
                        "message" : "password is incorrect",
                        })
                */
            } else {
                //check token is valid
                if err != nil {
                    errMessage = err.Error()
                } else {
                    token,err := token.GenerateToken(usr.ID)
                    if err != nil {
                        errMessage = "Cannot Generate Token"
                    } else {
                        c.JSON(http.StatusOK, gin.H {
                                "token" : token,
                                "message" : "Login Successful",
                                })
                        return
                    }
                }
            }
        }
        if errMessage != "" {
                c.JSON(http.StatusInternalServerError, gin.H {
                        "message" : errMessage,
                        })
                return
        }
}


