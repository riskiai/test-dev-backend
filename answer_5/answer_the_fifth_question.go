package main

import "fmt"

// Definisi struct Hewan
type Hewan struct {
    Jenis         string
    Ras           string
    Nama          string
    Karakteristik string
}

// Function untuk menghitung jumlah hewan sesuai jenis
func hitungJumlahPerJenis(hewanPeliharaan []Hewan) map[string]int {
    jumlah := make(map[string]int)
    for _, h := range hewanPeliharaan {
        jumlah[h.Jenis]++
    }
    return jumlah
}

func main() {
    // Data hewan peliharaan Esa
    hewanPeliharaan := []Hewan{
        {"Anjing", "Golden Retriever", "Otto", "Energik dan senang bermain bola"},
        {"Anjing", "Siberian Husky", "Max", "Bulu lebat dan mata biru"},
        {"Anjing", "Beagle", "Bob", "Ceria dan aktif bermain di taman"},
        {"Kucing", "Persia", "Luna", "Anggun dan manja"},
        {"Kucing", "British Short Hair", "Milo", "Cerdas dan aktif"},
        {"Ikan", "Koi", "Nana", "Indah dan menarik"},
        {"Ikan", "Mas", "Goldie", "Berwarna cerah"},
        {"Badak", "Badak Jawa", "Rino", "Pekerja keras"},
    }

    // Hitung jumlah hewan per jenis
    hasil := hitungJumlahPerJenis(hewanPeliharaan)

    fmt.Println("Jumlah hewan peliharaan Esa berdasarkan jenis:")
    for jenis, count := range hasil {
        fmt.Printf("%s: %d\n", jenis, count)
    }
}
