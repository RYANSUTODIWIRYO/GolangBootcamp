package main

import (
	"fmt"
	"time"
)

func main()  {	
	//gajadi yeeee
	
	const tahun int = 2020
	bulan := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "Nobvember", "Desember"}		//Inisialisasi bulan array

	fmt.Printf("\nCalendar Tahun %d\n", tahun)
	
	for i := 0; i < 12  ; i++ {												//Looping print kalendar
		fmt.Printf("=====\t%s   =====\n", bulan[i])							//Bahasa Indonesia
		
		tanggal := Date(tahun, i + 2, 0)									//Mendapatkan date bulan yg akan diprint
		// fmt.Printf("=====\t%s  =====\n", tanggal.Month())				//Bahasa Inggris
		fmt.Printf("S   S   R   K   J   S   M\n")

		tanggalSebelum := Date(tahun, i + 1, 0)								//Mendapatkan date bulan sebelumnya
		hariSebelum := int(tanggalSebelum.Weekday())						//Mendapatkan hari terakhir bulan sebelumnya
																 			//sehingga tanggal dimulai setelah hariSebelum
		spasi := ""
		for i := 0; i < hariSebelum; i++ {									//Membuat jarak tanggal pertama
			spasi = spasi + "    "
		}
		fmt.Printf(spasi)	

		for j := 0; j < int(tanggal.Day()); j++ {							
			fmt.Printf("%d", j + 1)											//Print tanggal
			
			if (j + 1) > 9 {												//Print jarak antar tanggal
				fmt.Printf("  ")
			} else {
				fmt.Printf("   ")
			}
			
			hariSebelum++
			if hariSebelum == 7 {
				hariSebelum = 0
				fmt.Printf("\n")
			}
		}
		fmt.Printf("\n\n")
	}
}

func Date(year, month, day int) time.Time {									//Method untuk date dengan tiga paramater
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)	//Mengembalikan date
}
