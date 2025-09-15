package main

import "fmt"

// Definisi struct Hewan
type Hewan struct {
    Jenis         string
    Ras           string
    Nama          string
    Karakteristik string
}

// Function untuk mengganti kucing Persia jadi Maine Coon
func gantiPersiaDenganMaineCoon(kucing Hewan) (Hewan, Hewan) {
    before := kucing
    if kucing.Jenis == "Kucing" && kucing.Ras == "Persia" {
        kucing.Ras = "Maine Coon"
    }
    after := kucing
    return before, after
}

func main() {
    // Data kucing Persia (Luna)
    luna := Hewan{"Kucing", "Persia", "Luna", "Anggun dan manja"}

    // Jalankan fungsi
    before, after := gantiPersiaDenganMaineCoon(luna)

    fmt.Println("Data Ras sebelum diganti:")
    fmt.Printf("Nama: %s | Jenis: %s | Ras: %s | Karakteristik: %s\n",
        before.Nama, before.Jenis, before.Ras, before.Karakteristik)

    fmt.Println("\nData Ras sesudah diganti:")
    fmt.Printf("Nama: %s | Jenis: %s | Ras: %s | Karakteristik: %s\n",
        after.Nama, after.Jenis, after.Ras, after.Karakteristik)
}
