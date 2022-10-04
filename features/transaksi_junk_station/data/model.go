package data

import (
	transaksijunkstation "rozhok/features/transaksi_junk_station"

	"gorm.io/gorm"
)

type TransaksiJunkStation struct {
	gorm.Model
	UserID     uint
	KodeTf     string
	GrandTotal int64

	TransaksiJunkStationDetail []TransaksiJunkStationDetail
}

type TransaksiJunkStationDetail struct {
	gorm.Model
	TransaksiJunkStationID uint
	KategoriID             uint
	Berat                  int
	Subtotal               int64

	KategoriRosok KategoriRosok
}

type KategoriRosok struct {
	gorm.Model
	NamaKategori string
	HargaMitra   int64
	HargaClient  int64
	Desc         string
}

type KeranjangRosok struct {
	gorm.Model
	ClientID        uint
	KategoriRosokID uint
	Berat           int
	Subtotal        int64
	KategoriRosok   KategoriRosok
}

func ToCore(transaksiModel TransaksiJunkStation) transaksijunkstation.Core {
	transaksiCore := transaksijunkstation.Core{
		GrandTotal: transaksiModel.GrandTotal,
		KodeTF:     transaksiModel.KodeTf,
		CreatedAt:  transaksiModel.CreatedAt.Format("2006-01-02"),
	}

	barangRosokCoreList := []transaksijunkstation.BarangRosok{}
	for _, barangrosok := range transaksiModel.TransaksiJunkStationDetail{
		barangRosokCoreList = append(barangRosokCoreList, transaksijunkstation.BarangRosok{
			ID: barangrosok.ID,
			Kategori: barangrosok.KategoriRosok.NamaKategori,
			Berat: uint(barangrosok.Berat),
			Subtotal: int64(barangrosok.Berat) * barangrosok.KategoriRosok.HargaMitra,
		})
	}

	transaksiCore.BarangRosok = barangRosokCoreList

	return transaksiCore
}
