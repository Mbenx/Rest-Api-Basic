package models

import "github.com/jinzhu/gorm"

// Karyawan is Model
type Karyawan struct {
	gorm.Model
	Nama, Alamat, Jabatan string
}

// TableName karyawan
func (Karyawan) TableName() string {
	return "karyawan"
}
