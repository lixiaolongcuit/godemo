package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	//自定创建表
	db.AutoMigrate(&Product{})

	//插入数据
	product := &Product{Code: "test", Price: 1000}
	if err := db.Create(product).Error; err != nil {
		log.Fataln("create error:", err)
	}
	log.Println("add product id:", product.ID)
}
