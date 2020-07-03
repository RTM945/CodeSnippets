package main

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Account struct {
	ID      int `gorm:"primary_key"`
	Balance int `gorm:"not_null"`
}

//不指定表名会默认为accounts
func (Account) TableName() string {
	return "account"
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	fmt.Println(db.HasTable(Account{}))
	db = db.Create(&Account{Balance: 1000})
	if db.Error != nil {
		t.Error(db.Error)
	}
	fmt.Println(db.RowsAffected)
	var accounts []Account
	db.Find(&accounts)
	fmt.Println(accounts)
}
