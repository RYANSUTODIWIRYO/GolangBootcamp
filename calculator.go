package main

import (
	"fmt"
	"math"
)

type Hitung struct {
	Angka1, Angka2 float64
}

func main()  {

	var pil int
	var menu string = "\n===Program Kalkulator===\n"+
		"1.  Penjumlahan\n" +
		"2.  Pengurangan\n" +
		"3.  Perkalian\n" +
		"4.  Pembagian\n" +
		"5.  Akar Kuadrat\n" +
		"6.  Kuadrat\n" +
		"7.  Luas Persegi\n" +
		"8.  Luas Lingkaran\n" +
		"9.  Volume Tabung\n" +
		"10. Volume Balok\n" +
		"11. Volume Prisma\n" +
		"99. Exit\n"

	for pil != 99 {
		fmt.Print(menu)
		fmt.Print("Masukan Pilihan (1-11) : ")
		fmt.Scan(&pil)
		// fmt.Printf("pil : %d\n", pil)

		var angka1, angka2 float64

		switch pil {
		case 1:			
			fmt.Println("\n===Penjumlahan===")
			fmt.Print("Angka pertama\t: ")
			fmt.Scan(&angka1)
			fmt.Print("Angka kedua\t: ")
			fmt.Scan(&angka2)

			perhitungan := Hitung{angka1, angka2}
			// fmt.Printf("Hasil %s + %s = %.2f", angka1, angka2, penjumlahan.Penjumlahan())
			fmt.Printf("Hasil\t\t: %.2f\n", perhitungan.Penjumlahan())
		case 2:			
			fmt.Println("\n===Pengurangan===")
			fmt.Print("Angka pertama\t: ")
			fmt.Scan(&angka1)
			fmt.Print("Angka kedua\t: ")
			fmt.Scan(&angka2)

			perhitungan := Hitung{angka1, angka2}
			fmt.Printf("Hasil\t\t: %.2f\n", perhitungan.Pengurangan())
		case 3:			
			fmt.Println("\n===Perkalian===")
			fmt.Print("Angka pertama\t: ")
			fmt.Scan(&angka1)
			fmt.Print("Angka kedua\t: ")
			fmt.Scan(&angka2)

			perhitungan := Hitung{angka1, angka2}
			fmt.Printf("Hasil\t\t: %.2f\n", perhitungan.Perkalian())
		case 4:			
			fmt.Println("\n===Pembagian===")
			fmt.Print("Angka pertama\t: ")
			fmt.Scan(&angka1)
			fmt.Print("Angka kedua\t: ")
			fmt.Scan(&angka2)

			perhitungan := Hitung{angka1, angka2}
			fmt.Printf("Hasil\t\t: %.2f\n", perhitungan.Pembagian())
		case 5:
			var angka float64
			fmt.Println("\n===Akar Kuadrat===")
			fmt.Print("Angka\t: ")
			fmt.Scan(&angka)
			fmt.Printf("Hasil\t: %.2f\n", math.Sqrt(angka))
		case 6:
			var angka float64
			fmt.Println("\n===Kuadrat===")
			fmt.Print("Angka\t: ")
			fmt.Scan(&angka)
			fmt.Printf("Hasil\t: %.2f\n", math.Pow(angka, 2))
		case 7:
			var panjang, lebar float64
			fmt.Println("\n===Luas Persegi===")
			fmt.Print("Panjang\t: ")
			fmt.Scan(&panjang)
			fmt.Print("Lebar\t: ")
			fmt.Scan(&lebar)
			fmt.Printf("Hasil\t: %.2f\n", LuasPersegi(panjang, lebar))
		case 8:
			var jari float64
			fmt.Println("\n===Luas Lingkaran===")
			fmt.Print("Jari-jari\t: ")
			fmt.Scan(&jari)
			fmt.Printf("Hasil\t\t: %.2f\n", LuasLingkaran(jari))
		case 9:
			var jari, tinggi float64
			fmt.Println("\n===Volume Tabung===")
			fmt.Print("Jari-jari\t: ")
			fmt.Scan(&jari)
			fmt.Print("Tinggi\t: ")
			fmt.Scan(&tinggi)
			fmt.Printf("Hasil\t\t: %.2f\n", VolumeTabung(jari, tinggi))
		case 10:
			var panjang, lebar, tinggi float64
			fmt.Println("\n===Volume Balok===")
			fmt.Print("Panjang\t: ")
			fmt.Scan(&panjang)
			fmt.Print("Lebar\t: ")
			fmt.Scan(&lebar)
			fmt.Print("Tinggi\t: ")
			fmt.Scan(&tinggi)
			fmt.Printf("Hasil\t\t: %.2f\n", VolumeBalok(panjang, lebar, tinggi))
		case 11:
			var panjang, lebar, tinggi float64
			fmt.Println("\n===Volume Prisma===")
			fmt.Print("Alas Segitiga\t: ")
			fmt.Scan(&panjang)
			fmt.Print("Tinggi Segitiga\t: ")
			fmt.Scan(&lebar)
			fmt.Print("Tinggi Prisma\t: ")
			fmt.Scan(&tinggi)
			fmt.Printf("Hasil\t\t: %.2f\n", VolumeBalok(panjang, lebar, tinggi))
		case 99:
			fmt.Printf("Berhasil Keluar")
		default:
			fmt.Printf("Pilihan salah")
		}
	}
}

func (h Hitung) Penjumlahan() float64 {
	return h.Angka1 + h.Angka2
}

func (h Hitung) Pengurangan() float64 {
	return h.Angka1 - h.Angka2
}

func (h Hitung) Perkalian() float64 {
	return h.Angka1 * h.Angka2
}

func (h Hitung) Pembagian() float64 {
	return h.Angka1 / h.Angka2
}

func LuasPersegi(panjang float64, lebar float64) float64 {
	return panjang * lebar
}

func LuasLingkaran(jari float64) float64 {
	return math.Pi * jari * jari
}

func VolumeTabung(jari float64, tinggi float64)  float64 {
	return math.Pi * jari * jari * tinggi
}

func VolumeBalok(panjang float64, lebar float64, tinggi float64)  float64 {
	return panjang * lebar * tinggi
}

func VolumePrisma(alasSegitiga float64, tinggiSegitiga float64, tinggiPrisma float64)  float64 {
	return alasSegitiga * tinggiSegitiga * tinggiPrisma
}

