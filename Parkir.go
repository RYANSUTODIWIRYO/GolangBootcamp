package main

import (
	"fmt"
	"time"
)

func main() {
	menu := "\n=== Menu Parkir ===\n" +					//String menu
		"1.  Masuk\n" +
		"2.  Keluar\n" +
		"99. Keluar Program"

	kendaraan := "\n=== Jenis Kendaraan ===\n" +		//String kendaraan
		"1. Mobil\n" +
		"2. Motor"
	
	customer := make(map[int]time.Time)					//Map customer
	id := 210200										//Inisialisasi id
	var pil int32
	for pil != 99 {
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)									//Pilih menu
		
		switch pil {
		case 1:
			customer[id] = time.Now()
			fmt.Println("\nId Customer\t:", id)
			id++
		case 2:
			var ken int32
			var tarif int64
			var idKendaraan int
						
			for (ken != 1) && (ken != 2) {
				fmt.Println(kendaraan)
				fmt.Print("Pilihan\t: ")
				fmt.Scan(&ken)
				switch ken {
				case 1:
					tarif = 3000
				case 2:
					tarif = 2000
				default:
					fmt.Println("Pilihan Salah!")
				}	
			}

			fmt.Print("Id Kendaraan\t: ")
			fmt.Scan(&idKendaraan)

			jamMasuk := customer[idKendaraan]
			jamKeluar := time.Now()

			Keluar(jamMasuk, jamKeluar, tarif)
		case 99:
			fmt.Println("Berhasil Keluar")
		default:
			fmt.Println("Pilihan salah")
		}
	}
}

func Keluar(jamMasuk time.Time, jamKeluar time.Time, tarif int64){
	selisih := jamKeluar.Sub(jamMasuk)
	totalTarif := (int64 (selisih.Seconds())) * tarif

	fmt.Println("\nJam Masuk\t:", jamMasuk.Format("2006-01-02 15:04:05"))
	fmt.Println("Jam Keluar\t:", jamKeluar.Format("2006-01-02 15:04:05"))
	fmt.Println("Selisih\t\t:", int64 (selisih.Seconds()))
	fmt.Println("Total Tarif\t:", totalTarif)
}