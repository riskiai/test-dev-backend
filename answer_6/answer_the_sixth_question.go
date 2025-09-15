package main

import (
    "fmt"
    "strings"
)

// Definisi struct Hewan
type Hewan struct {
    Jenis         string
    Ras           string
    Nama          string
    Karakteristik string
}

// Fungsi untuk cek palindrome
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

// Fungsi untuk mencari hewan dengan nama palindrome
func cariPalindrome(hewanPeliharaan []Hewan) []Hewan {
    var hasil []Hewan
    for _, h := range hewanPeliharaan {
        if isPalindrome(h.Nama) {
            hasil = append(hasil, h)
        }
    }
    return hasil
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

    // Cari hewan dengan nama palindrome
    palindromes := cariPalindrome(hewanPeliharaan)

    fmt.Println("Hewan peliharaan Esa dengan nama palindrome:")
    for _, h := range palindromes {
        fmt.Printf("Nama: %s | Jenis: %s | Ras: %s | Panjang Nama: %d\n",
            h.Nama, h.Jenis, h.Ras, len(h.Nama))
    }
}
