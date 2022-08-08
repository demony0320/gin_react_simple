package pgmodel

import (
    "gorm.io/gorm"
    "encoding/json"
    "fmt"
    "net/http"
	"github.com/gin-gonic/gin"
)

type post struct {
    gorm.Model

    Category string `binding:"required" gorm:"not null;default:null"`//this needs limit?
    Subject string `binding:"required" gorm:"not null;default:null"`
    Content string `binding:"required" gorm:"not null;default:null"`
    Hits int
    //Consider Image later.

    //Other Types 
    //User Dependency
    Writer string  `binding:"required" gorm:"not null;default:null"`
    //User   user `binding:"required" gorm:"foreignkey:Userid;references:writer;constraint:OnUpdate:CASCADE"`

    //Not Implemented yet. 
    //Reacts react[]
    //Comments comment[]
}

//type comment {
//    Post post
//    Comment comment //if exists, this is reply of comment
//    Context string
//    Reacts react[]
//}
//
//type react {
//    Type boolean  //true->like, false->hate
//    By   user
//}

func (p post) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "category": p.Category,
        "subject": p.Subject,
        "content": p.Content,
        "hits": p.Hits,
        "writer" : p.Writer,
    })
}
func (mh *ModelHandler) CreatePost(c *gin.Context){
    p := post{}
    if err:=c.BindJSON(&p); err != nil{
         c.JSON(http.StatusInternalServerError, gin.H {
            "message" : err.Error(),
        })
        return 
    }
    result := mh.Db.Create(&p)
    if result.Error != nil {
      c.JSON(http.StatusInternalServerError, gin.H {
          "message" : result.Error.Error(),
      })
    return
    } 
    fmt.Println("Created Post: ", p)

    c.JSON(http.StatusOK, gin.H{
        "message": "Create User Successful",
        "Subject" : p.Subject,
        "Category" : p.Category,
    })
}

