package models

type Mahasiswa struct {
	ID int `json:"id" binding:"required"`
	Nama string `json:"nama" binding:"required"`
	Prodi string `json:"prodi" binding:"required"`
	Fakultas string `json:"fakultas" binding:"required"`
	NIM int `json:"nim" binding:"required,min=6" gorm:"primary_key"`
	Angkatan int `json:"tahun_angkatan" binding:"required"`
}