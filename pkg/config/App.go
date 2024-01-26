package config

import(
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  )
  
  var db *gorm.DB
 
  func ConnectDb(){
    d,err:=gorm.Open("mysql","sql6679718:RKYAdxSc5P@tcp(sql6.freesqldatabase.com:3306)/sql6679718?parseTime=true")
    if err!=nil{
      panic(err)
    }
    db=d
  }
  func GetDb() *gorm.DB{
    return db
  }