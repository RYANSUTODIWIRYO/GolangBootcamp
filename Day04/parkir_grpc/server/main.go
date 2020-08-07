package main

import (
	// "context"
	"log"
	"net"
	"fmt"
	"time"
	"strconv"

	parkir "parkir_grpc/parkir"
	
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const (
	port = ":9000"
)

var (
	customer map[string]time.Time										
	id int	
	kendaraan string
)

type Server struct {
	parkir.UnimplementedParkirServiceServer
}

func (s *Server) MasukParkir(ctx context.Context, in *parkir.Empty) (*parkir.KarcisParkir, error) {
	log.Printf("Masuk Pak Ekoooooo")

	idKendaraan := "RSS" + strconv.Itoa(id)		//Id kendaraan masuk
	customer[idKendaraan] = time.Now()
	//fmt.Println("\nId Customer\t:", idKendaraan)
	id++

	return &parkir.KarcisParkir{IdKendaraan: idKendaraan}, nil
}

func (s *Server) KeluarParkir(ctx context.Context, in *parkir.Keluar) (*parkir.StrukParkir, error) {
	log.Printf("Keluar Pak Ekoooooo")

	var (
		jamMasuk time.Time
		tarif, tarifAwal, tarifMaksimal int64
		)

	value, ok := customer[in.IdKendaraan];				//Cek Id Kendaraan
	if ok {
		jamMasuk = value
	} else {
		// fmt.Println("Id Kendaraan tidak ditemukan")
		// goto inputIdKen
	}

	jamKeluar := time.Now()						//Jam Keluar
	
	switch in.JenisKendaraan {							//Cek tarif sesuai jenis
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

	// fmt.Println("\nJam Masuk\t:", jamMasuk.Format("2006-01-02 15:04:05"))
	// fmt.Println("Jam Keluar\t:", jamKeluar.Format("2006-01-02 15:04:05"))
	// fmt.Println("Durasi Parkir\t:", int (durasi.Seconds()))
	// fmt.Println("Total Tarif\t:", totalTarif)

	return &parkir.StrukParkir{
			IdKendaraan: in.IdKendaraan,
			JenisKendaraan: in.JenisKendaraan,
			PlatNomor : in.PlatNomor,
			JamMasuk : jamMasuk.Format("2006-01-02 15:04:05"),
			JamKeluar : jamKeluar.Format("2006-01-02 15:04:05"),
			Durasi : int64 (durasi.Seconds()),
			TotalTarif : int64 (totalTarif)},
		nil
}

func main() {
	fmt.Println("Server is running...")

	customer = make(map[string]time.Time)		//Map customer
	id = 210200									//Inisialisasi id

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	parkir.RegisterParkirServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}