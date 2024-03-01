package config

import(
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  )
  
  var db *gorm.DB
 
  func ConnectDb(){
    d,err:=gorm.Open("mysql","sql6687891:cRk6pD7LRK@tcp(sql6.freesqldatabase.com:3306)/sql6687891?parseTime=true")
    if err!=nil{
      panic(err)
    }
    db=d
  }
  func GetDb() *gorm.DB{
    return db
  }