package main

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	parkir "parkir_grpc/parkir"
)

func main() {	
	menu := "\n=== Menu Parkir ===\n" +						//String menu
		"1.  Masuk\n" +
		"2.  Keluar\n" +
		"99. Keluar Program"

	var pil int
	for pil != 99 {
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)										//Pilih menu
		
		switch pil {
		case 1:
			MasukParkir()
		case 2:
			KeluarParkir()
		case 99:
			fmt.Println("Berhasil Keluar")
		default:
			fmt.Println("Pilihan salah")
		}
	}
}

// func Connection(conn *grpc.ClientConn)  {
// 	//var conn *grpc.ClientConn
// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %s", err)
// 	} else {
// 		log.Printf("Server connected")
// 	}
// }

func MasukParkir()  {
	// var conn *grpc.ClientConn
	// Connection(conn)
	// defer conn.Close()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	} else {
		// log.Printf("Server connected")
	}

	defer conn.Close()

	park := parkir.NewParkirServiceClient(conn)

	response, err := park.MasukParkir(context.Background(), &parkir.Empty{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	// log.Printf("Id Kendaraan\t: %s", response.IdKendaraan)
	fmt.Printf("\n===== Karcis Parkir =====\n")
	fmt.Printf("Id Kendaraan\t: %s\n", response.IdKendaraan)
}

func KeluarParkir()  {
	// var conn *grpc.ClientConn
	// Connection(conn)
	// defer conn.Close()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	} else {
		// log.Printf("Server connected")
	}

	defer conn.Close()

	park := parkir.NewParkirServiceClient(conn)

	// jenisKendaraan := "jk";
    // platNomor := "pn";
	// idKendaraan := "ik";
	
	var jenisKendaraan, platNomor, idKendaraan string

	kendaraan := "\n=== Jenis Kendaraan ===\n" +			//String kendaraan
	"1. Mobil\n" +
	"2. Motor"
	fmt.Println(kendaraan)
	fmt.Print("Pilih (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)

	fmt.Print("Plat Nomor\t: ")
	fmt.Scan(&platNomor)

	fmt.Print("Id Kendaraan\t: ")
	fmt.Scan(&idKendaraan)

	response, err := park.KeluarParkir(context.Background(), &parkir.Keluar{
		JenisKendaraan: jenisKendaraan,
		PlatNomor: platNomor,
		IdKendaraan: idKendaraan})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	fmt.Printf("\n==== Struk Parkir ====\n")
	fmt.Printf("Id Kendaraan\t: %s\n", response.IdKendaraan)
	fmt.Printf("Jenis Kendaraan\t: %s\n", response.JenisKendaraan)
	fmt.Printf("Plat Nomor\t: %s\n", response.PlatNomor)
	fmt.Printf("Jam Masuk\t: %s\n", response.JamMasuk)
	fmt.Printf("Jam Keluar\t: %s\n", response.JamKeluar)
	fmt.Printf("Durasi Parkir\t: %d\n", response.Durasi)
	fmt.Printf("Total Tarif\t: %d\n", response.TotalTarif)
}