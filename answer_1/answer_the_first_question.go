package main

import "fmt"

// Definisi struct Hewan
type Hewan struct {
    Jenis        string
    Ras          string
    Nama         string
    Karakteristik string
}

// Function untuk menampilkan daftar hewan
func tampilkanHewan(hewanPeliharaan []Hewan) {
    for _, h := range hewanPeliharaan {
        fmt.Printf("Nama: %s | Jenis: %s | Ras: %s | Karakteristik: %s\n",
            h.Nama, h.Jenis, h.Ras, h.Karakteristik)
    }
}

func main() {
    // Array (slice) hewan peliharaan Esa
    hewanPeliharaan := []Hewan{
        {"Anjing", "Golden Retriever", "Otto", "Energik dan senang bermain bola"},
        {"Anjing", "Siberian Husky", "Max", "Bulu lebat dan mata biru"},
        {"Anjing", "Beagle", "Bob", "Ceria dan aktif bermain di taman"},
        {"Kucing", "Persia", "Luna", "Anggun dan manja"},
        {"Kucing", "British Short Hair", "Milo", "Cerdas dan aktif"},
        {"Ikan", "Koi", "Nana", "Indah dan menarik"},
        {"Ikan", "Mas", "Goldie", "Berwarna cerah"},
    }

    fmt.Println("Daftar Hewan Peliharaan Esa:")
    tampilkanHewan(hewanPeliharaan)
}
