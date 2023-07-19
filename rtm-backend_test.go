package main

import (
	"fmt"
	"os"
	"rtm/controller"
	"testing"
)

// var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

// func TestPrintenv(t *testing.T) {
// 	fmt.Println(MariaStringAkademik)
// }

// func TestInsertdata(t *testing.T) {
// 	id := "cc2"
// 	nama := "rofiganteng"
// 	email := "cth@gmail.com"
// 	hp := "000000"
// 	controller.InsertData()
// }
func TestInsertTagihanRegistrasi(t *testing.T) {
	dbname := "tagihan"
	tagihanregis := TagihanRegistrasi{
		ID:      primitive.NewObjectID(),
		NamaMahasiswa: "Nizar",
 		NIM:           "nizar@gmail.com",
 		Semester:      "dua",
 		BiayaRegistrasi:      "10000000",
 	}
 	insertedID := InsertTagihanRegistrasi(dbname, tagihanregis)
 	if insertedID == nil {
 		t.Error("Failed to insert tagihan")
 	}
 }
 func TestInsertTagihanSPP(t *testing.T) {
 	dbname := "tagihan"
 	tagihanspp:= TagihanSPP{
		ID:      primitive.NewObjectID(),
 		NamaMahasiswa: "Nizar",
 		NIM:           "nizar@gmail.com",
 		Semester:      "dua",
 		BiayaSPP:      "10000000",
 	}
 	insertedID := InsertTagihanSPP(dbname, tagihanspp)
 	if insertedID == nil {
 		t.Error("Failed to insert tagihan")
 	}
 }

 func TestGetTagihanRegistrasi(t *testing.T) {
	stats := "biaya_registrasi"
	data := GetTagihanRegistrasi(stats)
	fmt.Println(data)
}

func TestGetTagihanSPP(t *testing.T) {
	stats := "biaya_spp"
	data := GetTagihanSPP(stats)
	fmt.Println(data)
}
