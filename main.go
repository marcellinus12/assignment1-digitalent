package main

import (
	"fmt"
	"os"
	"strconv"
)

type murid struct {
	no        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Parameter tidak ada")
		return
	}

	// Mengonversi argumen ke integer
	no, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	biodata(no)
}

func biodata(no int) {
	var data = []murid{
		{no: 1, nama: "Marcel", alamat: "Surabaya", pekerjaan: "Developer", alasan: "karena menantang untuk dipelajari"},
		{no: 2, nama: "Ilham", alamat: "Sidoarjo", pekerjaan: "Frontend", alasan: "gampang dipelajari"},
		{no: 3, nama: "Baron", alamat: "Bojonegoro", pekerjaan: "Data Science", alasan: "lebih banyak digunakan didunia kerja"},
	}

	for _, a := range data {
		if no == a.no {
			fmt.Println("Nama:", a.nama)
			fmt.Println("Alamat:", a.alamat)
			fmt.Println("Pekerjaan:", a.pekerjaan)
			fmt.Println("Alasan:", a.alasan)
		}
	}
}
