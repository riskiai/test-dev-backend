package main

import "fmt"

// Definisi struct Hewan
type Hewan struct {
    Jenis         string
    Ras           string
    Nama          string
    Karakteristik string
    Kesayangan    bool
}

// Function untuk menampilkan daftar hewan
func printHewan(hewanPeliharaan []Hewan) {
    for _, h := range hewanPeliharaan {
        status := ""
        if h.Kesayangan {
            status = "(Kesayangan)"
        }
        fmt.Printf("Nama: %s | Jenis: %s | Ras: %s | Karakteristik: %s %s\n",
            h.Nama, h.Jenis, h.Ras, h.Karakteristik, status)
    }
}

// Function untuk menambah hewan baru
func addHewan(hewanPeliharaan []Hewan, jenis, ras, nama, karakteristik string, kesayangan bool) []Hewan {
    baru := Hewan{
        Jenis:         jenis,
        Ras:           ras,
        Nama:          nama,
        Karakteristik: karakteristik,
        Kesayangan:    kesayangan,
    }
    return append(hewanPeliharaan, baru)
}

func main() {
    // Data awal hewan peliharaan Esa
    hewanPeliharaan := []Hewan{
        {"Anjing", "Golden Retriever", "Otto", "Energik dan senang bermain bola", true},
        {"Anjing", "Siberian Husky", "Max", "Bulu lebat dan mata biru", true},
        {"Anjing", "Beagle", "Bob", "Ceria dan aktif bermain di taman", false},
        {"Kucing", "Persia", "Luna", "Anggun dan manja", true},
        {"Kucing", "British Short Hair", "Milo", "Cerdas dan aktif", true},
        {"Ikan", "Koi", "Nana", "Indah dan menarik", false},
        {"Ikan", "Mas", "Goldie", "Berwarna cerah", false},
    }

    fmt.Println("Daftar Hewan Peliharaan Esa (sebelum ditambah):")
    printHewan(hewanPeliharaan)

    // Tambah Badak Jawa bernama Rino
    hewanPeliharaan = addHewan(
        hewanPeliharaan,
        "Badak", "Badak Jawa", "Rino", "Pekerja keras", true,
    )

    fmt.Println("\nDaftar Hewan Peliharaan Esa (setelah ditambah Rino):")
    printHewan(hewanPeliharaan)
}
