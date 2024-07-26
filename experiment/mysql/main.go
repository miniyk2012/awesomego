package main

import (
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

func createDB() (*gorm.DB, error) {
	dsn := "yangkai.04:Yk@5078231@tcp(127.0.0.1:3306)/demo1?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

type UserWithJSON struct {
	gorm.Model
	ID        int64
	JsonValue datatypes.JSON
}

// TableName 设置User的表名为`custom_user`
func (UserWithJSON) TableName() string {
	return "test_json"
}

func queryJson(db *gorm.DB) {
	var users []UserWithJSON
	db.Find(&users, datatypes.JSONArrayQuery("json_value").Contains("z"))
	fmt.Printf("%+v\n", users)
}

func insertJSON(db *gorm.DB) {
	db.Create(&UserWithJSON{
		JsonValue: datatypes.JSON([]byte(`["ab", "b", "z"]`)),
	})
}

type User struct {
	ID        uint
	Name      string                `gorm:"uniqueIndex:udx_name;type:varchar(100)"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name;not null"`
}

func createTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&User{})
}
func main() {
	db, err := createDB()
	if err != nil {
		panic(err)
	}
	//queryJson(db)
	//insertJSON(db)
	//createTable(db)
	operateUser(db)
}

func operateUser(db *gorm.DB) {
	db.Debug().Where("name='yangkai1'").Delete(&User{})
	tx := db.Debug().Create(&User{Name: "yangkai1"})
	fmt.Printf("%v", tx.Error)
	//tx = db.Create(&User{Name: "yangkai2"})
	//fmt.Printf("%v", tx.Error)
}
