package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err =
		gorm.Open(sqlite.Open("./TekCloudResponsi.db"), &gorm.Config{})
	if err != nil {
		panic("Gagal Conect Ke Database")
	}
}

type (
	propinsi struct {
		ID    int     `json:"ID"`
		NAMA_PROPINSI  string  `json:"NAMA_PROPINSI"`
	}
)

func fetchAllBarang(c *gin.Context) {
	var model []propinsi

	db.Find(&model)

	if len(model) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
}

func main() {

	router := gin.Default()
	v1 := router.Group("/api/propinsi")
	{
		v1.GET("", fetchAllBarang)
	}
	router.Run(":8000")
}