package main

import (
    "fmt"
    "os"
    "strconv"
)

type Teman struct {
    Nama       string
    Alamat     string
    Pekerjaan  string
    AlasanGo   string
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Argumen tidak valid")
        return
    }

    absen, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Argumen harus berupa bilangan bulat")
        return
    }

    teman := []Teman{
        {
            Nama:       "Rizki",
            Alamat:     "Wonogiri",
            Pekerjaan:  "Software Engineer",
            AlasanGo:   "Cari bahasa pemrograman baru untuk dikuasai",
        },
        {
            Nama:       "Adib",
            Alamat:     "Batang",
            Pekerjaan:  "Web Developer",
            AlasanGo:   "Ingin mengolah data dengan mudah dan efisien",
        },
        {
            Nama:       "Fulan",
            Alamat:     "Yogyakarta",
            Pekerjaan:  "Mahasiswa",
            AlasanGo:   "Coba-coba",
        },
    }

    if absen > 0 && absen <= len(teman) {
        fmt.Println("Data Teman ke-", absen)
        fmt.Println("Nama:", teman[absen-1].Nama)
        fmt.Println("Alamat:", teman[absen-1].Alamat)
        fmt.Println("Pekerjaan:", teman[absen-1].Pekerjaan)
        fmt.Println("Alasan memilih kelas Golang:", teman[absen-1].AlasanGo)
    } else {
        fmt.Println("Absen tidak ditemukan")
    }
}
