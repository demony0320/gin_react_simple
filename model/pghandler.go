package model

import (
    "fmt"
    "io/ioutil"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
    "encoding/json"
    "strings"
)

type ModelHandler struct{
    Db *gorm.DB
}

type DbInfo struct {
    DbType   string `json:"DbType"`
    DbHost   string `json:"DbHost"`
    DbUser   string `json:"DbUser"`
    DbPass   string `json:"DbPass"`
    DbPort   int    `json:"DbPort"`
    DbName   string `json:"DbName"`
    TimeZone string `json:"TimeZone"`
}

func ModelInit(jsonFile string) (error, ModelHandler) {
    //read our opened jsonFile as a byte array.
    //we unmarshal our byteArray which contains our
    //jsonFile's content into 'users' which we defined above

    byteValue, _ := ioutil.ReadFile(jsonFile)
    var dbInfo DbInfo
    json.Unmarshal(byteValue, &dbInfo)
    if strings.ToUpper(dbInfo.DbType) != "POSTGRESQL" {
        panic("Currently only POstgreSQL DB was supported")
    }
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", dbInfo.DbHost, dbInfo.DbUser, dbInfo.DbPass,dbInfo.DbName, dbInfo.DbPort, dbInfo.TimeZone)

    fmt.Println("Generated dsn : %s", dsn)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect PG Database")
	}

    mh := ModelHandler{Db:db}

    err = mh.Db.AutoMigrate(&user{})
    if err != nil{
        panic("DB User Migration Failed")
    }
    err = mh.Db.AutoMigrate(&post{})
    if err != nil{
        panic("DB Post Migration Failed")
    }
    return err, mh
}
/*
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
*/
