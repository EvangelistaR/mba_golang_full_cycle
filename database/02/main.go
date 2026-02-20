package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "4k Screen ultrawide",
	// 	Price: 250.00,
	// })

	// products := []Product{
	// 	{Name: "Magic mouse", Price: 99},
	// 	{Name: "Magic keyboard", Price: 89},
	// }
	// db.Create(&products)

	var product Product

	// db.First(&product, 1)
	// fmt.Println(product)
	db.First(&product, "name = ?", "Magic mouse")

	fmt.Println(product)

	var products []Product

	db.Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}

}
