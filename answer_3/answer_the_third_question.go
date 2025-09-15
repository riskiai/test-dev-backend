package main

import (
    "fmt"
    "sort"
)

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

// Function untuk ambil hanya hewan kesayangan
func getKesayangan(hewanPeliharaan []Hewan) []Hewan {
    var result []Hewan
    for _, h := range hewanPeliharaan {
        if h.Kesayangan {
            result = append(result, h)
        }
    }
    return result
}

// Function untuk sort ascending (berdasarkan Nama)
func sortKesayanganAsc(hewan []Hewan) {
    sort.Slice(hewan, func(i, j int) bool {
        return hewan[i].Nama < hewan[j].Nama
    })
}

// Function untuk sort descending (berdasarkan Nama)
func sortKesayanganDesc(hewan []Hewan) {
    sort.Slice(hewan, func(i, j int) bool {
        return hewan[i].Nama > hewan[j].Nama
    })
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
        {"Badak", "Badak Jawa", "Rino", "Pekerja keras", true},
    }

    // Ambil hewan kesayangan
    kesayangan := getKesayangan(hewanPeliharaan)

    fmt.Println("Hewan Kesayangan Esa (Ascending):")
    sortKesayanganAsc(kesayangan)
    printHewan(kesayangan)

    fmt.Println("\nHewan Kesayangan Esa (Descending):")
    sortKesayanganDesc(kesayangan)
    printHewan(kesayangan)
}
