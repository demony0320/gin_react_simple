package model

import (
    "gorm.io/gorm"
    "encoding/json"
    "fmt"
    "net/http"
	"github.com/gin-gonic/gin"

    //For FileHandling
    "path/filepath"
    "github.com/google/uuid" // To generate random file names

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

    //mh.UploadFiles(c) 
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

func (mh *ModelHandler) FetchPost(c *gin.Context){
    id := c.Query("id")
    if id != ""{
        //postid is specified.
        var post post
        result := mh.Db.First(&post,"id = ?",id)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : result.Error.Error(),
            })
        } else {
            c.JSON(http.StatusOK, post)
        }
    } else {
        // Fetch All posts
        var posts []post
        result := mh.Db.Find(&posts)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "message" : result.Error.Error(),
            })
        } else {
            c.JSON(http.StatusOK, posts)
        }
    }
}

func (mh *ModelHandler) SaveFileHandler(c *gin.Context) {
    file, err := c.FormFile("file")

    // The file cannot be received.
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No file is received",
        })
        return
    }

    // Retrieve file information
    extension := filepath.Ext(file.Filename)
    // Generate random file name for the new uploaded file so it doesn't override the old file with same name
    newFileName := uuid.New().String() + extension

    // The file is received, so let's save it
    if err := c.SaveUploadedFile(file, "/home/ubuntu/repos/gin_react_simple/files" + newFileName); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "Unable to save the file",
        })
        return
    }

    // File saved successfully. Return proper result
    c.JSON(http.StatusOK, gin.H{
        "message": "Your file has been successfully uploaded.",
    })
}

//https://stackoverflow.com/questions/27674200/golang-panic-runtime-error-invalid-memory-address-or-nil-pointer-dereference
func (mh *ModelHandler) UploadFiles(c *gin.Context){
        form,err := c.MultipartForm()
        if err != nil { 
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }

        fmt.Println(form)
        files := form.File["upload[]"]
        fmt.Println(len(files))

        for _, file := range files {
            fmt.Println(file.Filename)
            dest_file_name := "assets/upload/" + filepath.Base(file.Filename)
                // Upload the file to specific dst.
                // Need to consider how to keep uniqueness 
                if err := c.SaveUploadedFile(file, dest_file_name); err != nil {
                    c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
                    return
                }
        c.Next( ) //don't know what it should be
    }
}
