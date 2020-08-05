package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var customer map[string]time.Time										
var id int	
var kendaraan string										

func main() {
	customer = make(map[string]time.Time)				//Map customer
	id = 210200										//Inisialisasi id

	//1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Selamat datang") //Langsung membuat function
	})

	//Masuk
	http.HandleFunc("/masuk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id_kendaraan := Masuk()
		data := "{\"id_kendaraan\":\""+ id_kendaraan +"\"}"				
		var jsonData = []byte(data)

		if r.Method == "GET" {
			w.Write(jsonData)
			return
		}

		http.Error(w, "Method not allowed", http.StatusBadRequest)
	})

	//Keluar
	http.HandleFunc("/keluar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, _ := ioutil.ReadAll(r.Body)			//Menerima request body
		keyVal := make(map[string]string)			//Membuat map	
		json.Unmarshal(body, &keyVal)				//Konversi lalu Memasukan body ke map
		
		id_kendaraan := keyVal["id_kendaraan"]		//Mengambil value by key dari map
		jenis_kendaraan := keyVal["jenis_kendaraan"]		

		//fmt.Println("id kendaraan :",id_kendaraan, "\njenis_kendaraan :", jenis_kendaraan)

		jam_masuk, jam_keluar, durasi, total_tarif := Keluar(id_kendaraan, jenis_kendaraan)

		data := "{\"jam_masuk\":\""+ jam_masuk +
			"\",\"jam_keluar\":\""+ jam_keluar +
			"\",\"durasi\":"+ strconv.Itoa(durasi) +
			",\"total_tarif\":"+ strconv.Itoa(int(total_tarif)) +"}"

		fmt.Println("data :", data)

		var jsonData = []byte(data)		

		if r.Method == "POST" {
			w.Write(jsonData)
			return
		}

		http.Error(w, "Method not allowed", http.StatusBadRequest)
	})

	fmt.Println("Server is running....")
	http.ListenAndServe(":8080", nil)			//Server stand by
}

func Masuk() string{
	idKendaraan := "RSS" + strconv.Itoa(id)		//Id kendaraan masuk
	customer[idKendaraan] = time.Now()
	//fmt.Println("\nId Customer\t:", idKendaraan)
	id++
	return idKendaraan
}

func Keluar(idKen string, jenisKen string) (string, string, int, int64){
	var (
		jamMasuk time.Time
		tarif, tarifAwal, tarifMaksimal int64
		)

	value, ok := customer[idKen];				//Cek Id Kendaraan
	if ok {
		jamMasuk = value
	} else {
		// fmt.Println("Id Kendaraan tidak ditemukan")
		// goto inputIdKen
	}

	jamKeluar := time.Now()						//Jam Keluar
	
	switch jenisKen {							//Cek tarif sesuai jenis
		case "mobil":							//kendaraan
			tarif = 3000
			tarifAwal = 5000
			tarifMaksimal = 1000000
		case "motor":
			tarif = 2000
			tarifAwal = 3000
			tarifMaksimal = 500000
	}

	durasi := jamKeluar.Sub(jamMasuk)
	totalTarif := tarifAwal + ((int64 (durasi.Seconds())) - 1) * tarif

	if totalTarif > tarifMaksimal{						//Cek tarif maksimal
		totalTarif = tarifMaksimal
	}

	fmt.Println("\nJam Masuk\t:", jamMasuk.Format("2006-01-02 15:04:05"))
	fmt.Println("Jam Keluar\t:", jamKeluar.Format("2006-01-02 15:04:05"))
	fmt.Println("Durasi Parkir\t:", int (durasi.Seconds()))
	fmt.Println("Total Tarif\t:", totalTarif)

	return string (jamMasuk.Format("2006-01-02 15:04:05")),
		string (jamKeluar.Format("2006-01-02 15:04:05")),
		int (durasi.Seconds()),
		totalTarif
}