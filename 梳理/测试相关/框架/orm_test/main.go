package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Uai struct {
	gorm.Model
	Email string
	Password string
	Type int
	Salt string
}

type Upi struct {
	gorm.Model
	PhotoUrl string
	Name string
	Sex int
	ContactPhone string
	ContactEmail string
	Birthday int
}

func main() {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			"localhost",
			5432,
			"postgres",
			"test",
			"123456",
		),
	)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.DropTable(&Uai{})
	db.DropTable(&Upi{})

	db.CreateTable(&Uai{})
	db.CreateTable(&Upi{})
	// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// 创建
	//for{
	//	db.Create(&Product{Code: "L1212", Price: 1000})
	//	time.Sleep(5*time.Second)
	//}
	//
	//// 读取
	//var product Product
	//db.First(&product, 1) // 查询id为1的product
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	//db.Delete(&product)
}
