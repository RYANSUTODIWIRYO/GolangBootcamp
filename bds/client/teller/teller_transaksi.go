package teller

import (
	"fmt"
	"golang.org/x/net/context"

	ent "bds/entities"
	conf "bds/config"
	bank "bds/proto"
)

func SetorTunai(transaksi ent.Transaksi) (ent.Transaksi, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.Transaksi{}, err
	}
	defer conn.Close()

	//Memanggil funtcion SetorTunai() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.SetorTunai(context.Background(), &bank.Transaksi{
		IdUser: transaksi.Id_user,
		NoRekening: transaksi.No_rekening,
		Tanggal: transaksi.Tanggal,
		JenisTransaksi: transaksi.Jenis_transaksi,
		Nominal: transaksi.Nominal,
		Saldo: transaksi.Saldo,
		Berita: transaksi.Berita,
	})

	//Memasukan nilai yang didapat
	if response.Saldo > 0 {
		transaksi.Saldo = response.Saldo
		return transaksi, nil
	} else {
		return ent.Transaksi{}, err
	}
}

func TarikTunai(transaksi ent.Transaksi) (ent.Transaksi, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return ent.Transaksi{}, err
	}
	defer conn.Close()

	//Memanggil funtcion TarikTunai() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.TarikTunai(context.Background(), &bank.Transaksi{
		IdUser: transaksi.Id_user,
		NoRekening: transaksi.No_rekening,
		Tanggal: transaksi.Tanggal,
		JenisTransaksi: transaksi.Jenis_transaksi,
		Nominal: transaksi.Nominal,
		Saldo: transaksi.Saldo,
		Berita: transaksi.Berita,
	})

	//Memasukan nilai yang didapat
	if response.Saldo > 0 { //berhasil
		transaksi.Saldo = response.Saldo
		return transaksi, nil
	} else if response.Berita == "Saldo Tidak Cukup" { //berhasil, saldo tidak cukup
		return ent.Transaksi{Berita:response.Berita}, nil
	} else {
		return ent.Transaksi{}, err //gagal
	}
}

func CetakBuku(transaksi ent.Transaksi) ([]ent.Transaksi, error) {
	//Koneksi ke grpc
	conn, err := conf.KoneksiGrpc()
	if err != nil {
		fmt.Println(err)
		return []ent.Transaksi{}, err
	}
	defer conn.Close()

	//Memanggil funtcion CetakBuku() dari server
	s := bank.NewBankServiceClient(conn)
	response, err := s.CetakBuku(context.Background(), &bank.Transaksi{
		NoRekening: transaksi.No_rekening,
	})

	fmt.Println("priinnttt", response)

	for _ , value := range response.Transaksi{
    	fmt.Println(value.IdTransaksi)
		fmt.Println(value.NoRekening)
		fmt.Println(value.JenisTransaksi)
		fmt.Println(value.Tanggal)
		fmt.Println(value.Nominal)
		fmt.Println(value.Saldo)
		fmt.Println(value.Berita)
	}

	return []ent.Transaksi{}, err

	// //Memasukan nilai yang didapat
	// if response.Saldo > 0 {
	// 	transaksi.Saldo = response.Saldo
	// 	return transaksi, nil
	// } else if response.Berita == "Saldo Tidak Cukup" {
	// 	return ent.Transaksi{Berita:response.Berita}, nil
	// } else {
	// 	return ent.Transaksi{}, err
	// }

}

