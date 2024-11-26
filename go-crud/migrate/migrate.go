package main

import (
	"github.com/fachi-r/go-crud/database"
	"github.com/fachi-r/go-crud/models"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	// database.DB.AutoMigrate(&models.Student{}, &models.HELSB{}, &models.Receipts{}, &models.Guardian{})
	database.DB.AutoMigrate(&models.Student{}, &models.Guardian{})
	// receipts := []int{
	// 	9848116,
	// 	3482454,
	// 	4950452,
	// 	4754198,
	// 	1224566,
	// 	4123071,
	// 	7524276,
	// 	2775063,
	// 	1801019,
	// 	3871548,
	// }

	// for _, receipt := range receipts {
	// 	newReceipt := models.Receipts{
	// 		Number: uint(receipt),
	// 	}
	// 	database.DB.Create(&newReceipt)
	// }

	// 8981358894
	// 6314886543
}
