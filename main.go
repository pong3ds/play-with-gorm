package main

import (
	"log"

	"github.com/3dsinteractive/gorm"
	_ "github.com/3dsinteractive/mysql"
)

// Product is the product
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func simpleCRUD() {
	db, err := gorm.Open("mysql", "my_user:my_password@/my_database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connect to mysql error :", err.Error())
		return
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	log.Printf("Product: %+v\n", product)
	// Delete - delete product
	db.Delete(&product)
}

func main() {
	simpleCRUD()
}
