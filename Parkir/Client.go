package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

var baseURL = "http://localhost:8080"

type Customer struct {
	Id_kendaraan 	string `json:"id_kendaraan"`
	Jenis_kendaraan string `json:"jenis_kendaraan"`
	Jam_masuk		string `json:"jam_masuk"`
	Jam_keluar		string `json:"jam_keluar"`
	Durasi			int    `json:"durasi"`
	Total_tarif		int64  `json:"total_tarif"`
}

func main() {
	menu := "\n=== Menu Parkir ===\n" +						//String menu
		"1.  Masuk\n" +
		"2.  Keluar\n" +
		"99. Keluar Program"

	kendaraan := "\n=== Jenis Kendaraan ===\n" +			//String kendaraan
		"1. Mobil\n" +
		"2. Motor"
	
	var pil int
	for pil != 99 {
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)										//Pilih menu
		
		switch pil {
		case 1:
			id_kendaraan, _ := masuk()
			fmt.Println("Id Kendaraan\t:",id_kendaraan)
		case 2:
			var jenis_kendaraan, plat, id_kendaraan string

			fmt.Println(kendaraan)
			fmt.Print("Pilih (mobil/motor): ")
			fmt.Scan(&jenis_kendaraan)

			fmt.Print("Plat Nomor\t: ")
			fmt.Scan(&plat)

			fmt.Print("Id Kendaraan\t: ")
			fmt.Scan(&id_kendaraan)

			jam_masuk, jam_keluar, durasi, total_tarif, _ := keluar(id_kendaraan, jenis_kendaraan)
			
			fmt.Println("\n==== Struk Parkir ====")
			fmt.Println("ID Kendaraan\t:", id_kendaraan)
			fmt.Println("Plat Nomor\t:", plat)
			fmt.Println("Jam Masuk\t:", jam_masuk)
			fmt.Println("Jam Keluar\t:", jam_keluar)
			fmt.Println("Durasi\t\t:", durasi)
			fmt.Println("Total Tarif\t:", int(total_tarif))
		case 99:
			fmt.Println("Berhasil Keluar")
		default:
			fmt.Println("Pilihan salah")
		}
	}
}

func masuk() (string, error){	
	var err error
	var client = &http.Client{}
	var data string
	var customer Customer

	request, err := http.NewRequest("GET", baseURL+"/masuk", nil)	//Membuat http request
	if err != nil {
		return "", err
	}

	response, err := client.Do(request)						//Eksekusi http request
	if err != nil {
		return "", err
	}

	defer response.Body.Close()								//Menutup response body
															//saat seluruh proses selesai
	if response.StatusCode == http.StatusOK {				//Funsi untuk menerima response body
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		data = string(bodyBytes)							//Konversi ke string
	}		
	byteData := []byte(data)								//Konversi string ke byte
	json.Unmarshal(byteData, &customer)						//Konversi byte ke struct

	return customer.Id_kendaraan, nil
}

func keluar(Id_kendaraan string, Jenis_kendaraan string) (string, string, int, int, error){	
	var err error
	var client = &http.Client{}
	var data string
	var customer Customer

	customer.Id_kendaraan = Id_kendaraan
	customer.Jenis_kendaraan = Jenis_kendaraan

	buf := new(bytes.Buffer)								//Membuat bytes encoder
	json.NewEncoder(buf).Encode(customer)					//Konversi dan memasukan struct ke buf
	request, err := http.NewRequest("POST", baseURL+"/keluar", buf) //Membuat http request dan request body
	if err != nil {
		return "","",0, 0, err
	}

	response, err := client.Do(request)						//Ekseskusi http request
	if err != nil {
		return "","",0, 0, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)		//Menerima response body
		if err != nil {
			return "","",0, 0, err
		}
		data = string(bodyBytes)							//Konversi ke string
	} else {
		fmt.Println("erroorr")
	}

	byteData := []byte(data)								//Konversi json string ke byte
	json.Unmarshal(byteData, &customer)						//Konversi json byte ke struct

	return string (customer.Jam_masuk),
	string (customer.Jam_keluar),
	int (customer.Durasi),
	int (customer.Total_tarif),
	nil
}
