package model

import (
	"gorm.io/gorm"
)

type ModelHandler struct{
    Db *gorm.DB
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

