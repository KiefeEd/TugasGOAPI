package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mahasiswawebapi/models"
)

type MahasiswaInput struct {
	ID int `json:"id" binding:"required"`
	Nama string `json:"nama" binding:"required"`
	Prodi string `json:"prodi" binding:"required"`
	Fakultas string `json:"fakultas" binding:"required"`
	NIM int `json:"nim" binding:"required,min=6" gorm:"primary_key"`
	Angkatan int `json:"tahun_angkatan" binding:"required"`
}

func Tampil(c* gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(200, gin.H {
		"data": mhs,
	})
}
func Tambah(c* gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
			c.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
	}
	//proses data
	mhs := models.Mahasiswa{
		ID: dataInput.ID,
		Nama: dataInput.Nama,
		Prodi: dataInput.Prodi,
		Fakultas: dataInput.Fakultas,
		NIM: dataInput.NIM,
		Angkatan: dataInput.Angkatan,

	}
	db.Create(&mhs)
	c.JSON(200, gin.H{
		"Message": "Data telah diinput",
		"Data": mhs,
	})
	

	
}

func Ubah(c* gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	//cek apakah data ada atau tdk
	var mhs models.Mahasiswa
	var dataInput MahasiswaInput
	if err := db.Where("nim = ? ", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(400, gin.H{
			"Error":"Data mahasiswa tidak ditemukan",
		})
	}

	if err := c.ShouldBindJSON(&dataInput); err != nil {
			c.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
	}

	// db.Create(&mhs)
	db.Model(&mhs).Update(dataInput)
	c.JSON(200, gin.H{
		"Message": "Data telah diedit",
		"Data": mhs,
	})
	

	
}

func Hapus(c* gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	//cek apakah data ada atau tdk
	var mhs models.Mahasiswa
	if err := db.Where("nim = ? ", c.Query("nim")).First(&mhs).Error; err != nil {
		c.JSON(400, gin.H{
			"Error":"Data mahasiswa tidak ditemukan",
		})
	}

	// db.Create(&mhs)
	db.Delete(&mhs)
	c.JSON(200, gin.H{
		"Message": "Data telah dihapus",
		"Berhasil": true,
	})
	

	
}



