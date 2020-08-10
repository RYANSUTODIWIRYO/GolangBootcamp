package main

import (
	"fmt"
	"log"
	"net"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"golang.org/x/net/context"

	ent "bds/entities"
	conf "bds/config"
	bank "bds/proto"
	serv "bds/service"
)

const (
	port = ":9000"
)

type server struct {
	bank.UnimplementedBankServiceServer
}

func main() {
	fmt.Println("Server is running...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bank.RegisterBankServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Login(ctx context.Context, user *bank.User) (*bank.User, error) {
	
	// Koneksi Database
	db, err := conf.KoneksiDB()
	if err != nil {
		panic(err)
		// return &bank.User{}, err
	}

	// Membuat struct koneksi
	con := serv.UserService{
		db,
	}
	
	idUser := user.IdUser
	password := user.Password

	// Memanggil function LoginUser() untuk login
	// fmt.Println(nama)
	// fmt.Println(password)
	response, err := con.LoginUser(idUser, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	u := bank.User{
		IdUser: 	response.Id_user,
		NamaUser: 	response.Nama_user,
		Password: 	response.Password,
		Cabang: 	response.Cabang,
		Role: 		response.Role,
	}
	return &u, nil
}

func (s *server) CariNasabahDetail(ctx context.Context, nasabah *bank.NasabahDetail) (*bank.NasabahDetail, error) {
	
	// Koneksi Database
	db, err := conf.KoneksiDB()
	if err != nil {
		// fmt.Println("errornya di siniiiii")
		panic(err)
		// return &bank.User{}, err
	}

	// Membuat struct koneksi
	con := serv.UserService{
		db,
	}
	
	rekTujuan := nasabah.NoRekening

	// Memanggil function LoginUser() untuk login
	response, err := con.CariNasabah(rekTujuan)
	if err != nil {
		panic(err)
	}

	// Membuat struct nasabah untuk dikembalikan
	//fmt.Println(nasabah)
	u := bank.NasabahDetail{
		Cif:		response.Cif,
		Nama:      	response.Nama,
		NoRekening: response.No_rekening,
		Saldo:      response.Saldo,
	}
	return &u, nil
}

func (s *server) SetorTunai(ctx context.Context, transaksi *bank.Transaksi) (*bank.Transaksi, error) {
	
	fmt.Println(transaksi)

	db, err := conf.KoneksiDB()
	if err != nil {
		panic(err)
	} else {
		con := serv.UserService{
			db,
		}
		// norek := transaksi.NoRekening
		// nominal := transaksi.Nominal
		// berita := transaksi.Berita

		// respons server
		// log.Printf(" client request : %v,%v,%v", noreq, nominal, berita)

		// call method stor tunai for check no rek exist or not
		nasabahDetail, err := con.CariNasabah(transaksi.NoRekening)
		fmt.Println(nasabahDetail)
		if err != nil {
			panic(err)
		}
		// if no req exist
		if nasabahDetail.No_rekening != 0 {

			// method add setor tunai to db called
			trx := ent.Transaksi {
				Id_user: transaksi.IdUser,
				No_rekening: transaksi.NoRekening,
				Tanggal: transaksi.Tanggal,
				Jenis_transaksi: transaksi.JenisTransaksi,
				Nominal: transaksi.Nominal,
				Saldo: transaksi.Saldo,
				Berita: transaksi.Berita,
			}

			stt, trx, err := con.SetorTunaiService(trx, nasabahDetail)
			if err != nil {
				panic(err)
			}
			// status := bank.Status{
			// 	Status: stt,
			// }
			if stt > 0 {
				fmt.Println("transaksi berhasil")
				transaksi.Saldo = trx.Saldo
				fmt.Println(transaksi)
				return transaksi, nil
			} else {
				fmt.Println("Transaksi gagal")
				return &bank.Transaksi{}, err
			}
		}
	}
	return &bank.Transaksi{}, nil
}

func (s *server) TarikTunai(ctx context.Context, transaksi *bank.Transaksi) (*bank.Transaksi, error) {
	
	fmt.Println(transaksi)

	db, err := conf.KoneksiDB()
	if err != nil {
		panic(err)
	} else {
		con := serv.UserService{
			db,
		}
		// norek := transaksi.NoRekening
		// nominal := transaksi.Nominal
		// berita := transaksi.Berita

		// respons server
		// log.Printf(" client request : %v,%v,%v", noreq, nominal, berita)

		// call method stor tunai for check no rek exist or not
		nasabahDetail, err := con.CariNasabah(transaksi.NoRekening)
		fmt.Println(nasabahDetail)
		if err != nil {
			panic(err)
		}
		// if no req exist
		if nasabahDetail.No_rekening != 0 {

			// method add tarik tunai to db called
			trx := ent.Transaksi {
				Id_user: transaksi.IdUser,
				No_rekening: transaksi.NoRekening,
				Tanggal: transaksi.Tanggal,
				Jenis_transaksi: transaksi.JenisTransaksi,
				Nominal: transaksi.Nominal,
				Saldo: transaksi.Saldo,
				Berita: transaksi.Berita,
			}

			stt, trx, err := con.TarikTunaiService(trx, nasabahDetail)
			if err != nil {
				panic(err)
			}
			// status := bank.Status{
			// 	Status: stt,
			// }
			if stt > 0 {
				fmt.Println("Transaksi berhasil")
				transaksi.Saldo = trx.Saldo
				fmt.Println(transaksi)
				return transaksi, nil
			} else if stt == -1{
				fmt.Println("Transaksi gagal, Saldo tidak cukup")
				return &bank.Transaksi{Berita:"Saldo Tidak Cukup"}, err
			} else {
				fmt.Println("Transaksi gagal")
				return &bank.Transaksi{}, err
			}
		}
	}
	return &bank.Transaksi{}, nil
}

func (s *server) CetakBuku(ctx context.Context, transaksi *bank.Transaksi) (*bank.ListTransaksi, error) {
	fmt.Println("masuuuuuuuuk", transaksi.NoRekening)
	var listTransaksi []*bank.Transaksi
	
	//Koneksi Ke Database
	db, err := conf.KoneksiDB()
	if err != nil {
		panic(err)
	} else {
		con := serv.UserService{
			db,
		}

		// call method CetakBuku for check no rek exist or not
		response, err := con.CetakBuku(int (transaksi.NoRekening))
		if err != nil {
			panic(err)
		}

		// Menampung nilai dari response
		for _, value := range response {
			trx := bank.Transaksi{
				IdTransaksi:    value.Id_transaksi,
				IdUser:         value.Id_user,
				NoRekening:     value.No_rekening,
				Tanggal:        value.Tanggal,
				Nominal:        value.Nominal,
				Saldo:          value.Saldo,
				JenisTransaksi:	value.Jenis_transaksi,
				Berita:         value.Berita,
			}
			listTransaksi = append(listTransaksi, &trx)
		}
	}

	fmt.Println(listTransaksi)

	return &bank.ListTransaksi{
		Transaksi: listTransaksi,
	}, nil
}