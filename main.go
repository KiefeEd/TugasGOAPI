package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"mahasiswawebapi/models"
	"mahasiswawebapi/controller"
)



func main() {

	r := gin.Default()
	db := models.SetupModels()
	r.Use(func(c * gin.Context) {
		
		c.Set("db", db)
	c.Next()
	})

	api := r.Group("/api")
	// api.POST("v1/mahasiswa", MahasiswaCreate)
	api.GET("v1/mahasiswa", controller.Tampil)
	api.POST("v1/mahasiswa", controller.Tambah)
	api.PUT("v1/mahasiswa/:nim", controller.Ubah)
	api.DELETE("v1/mahasiswa", controller.Hapus)
	r.Run()
	// note: add controller tambah too
}
// type Mahasiswa struct {
// 	ID int `json:"id" binding:"required"`
// 	Nama string `json:"nama" binding:"required"`
// 	Prodi string `json:"prodi" binding:"required"`
// 	Fakultas string `json:"fakultas" binding:"required"`
// 	NIM int `json:"nim" binding:"required"`
// 	Angkatan int `json:"tahun_angkatan" binding:"required"`
// }

// func MahasiswaCreate(c * gin.Context) {
// 	var MahasiswaInput Mahasiswa
// 	err := c.ShouldBindJSON(&MahasiswaInput)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(200, gin.H {
// 		"ID Mahasiswa": MahasiswaInput.ID,
// 		"Nama": MahasiswaInput.Nama,
// 		"Prodi": MahasiswaInput.Prodi,
// 		"Fakultas": MahasiswaInput.Fakultas,
// 		"NIM": MahasiswaInput.NIM,
// 		"Angkatan": MahasiswaInput.Angkatan,
// 	})
// }
