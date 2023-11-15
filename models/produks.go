package models

import (
	"github.com/jinzhu/gorm"
)

type Produk struct {
	Id         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	NamaProduk string `gorm:"size:100;not null" json:"namaProduk"`
	Keterangan string `gorm:"size:100;not null" json:"keterangan"`
	Harga      uint   `gorm:"not null" json:"harga"`
	Jumlah     uint   `gorm:"not null" json:"jumlah"`
}

func (p *Produk) CreatedProduk(db *gorm.DB) error {
	err := db.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *Produk) GetProduk(db *gorm.DB) ([]Produk, uint64, error) {
	count := 0
	produk := []Produk{}

	err := db.Find(&produk).
		Count(&count).
		Error

	if err != nil {
		return []Produk{}, 0, nil
	}

	return produk, uint64(count), nil
}

func (p *Produk) EditProduk(db *gorm.DB, id uint32) error {
	err := db.Model(&Produk{}).
		Where("id = ?", id).
		Update(&p).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (p *Produk) DeleteProduk(db *gorm.DB, id uint32) error {
	err := db.Model(&Produk{}).
		Where("id = ?", id).
		Delete(&p).
		Error

	if err != nil {
		return err
	}

	return nil
}
