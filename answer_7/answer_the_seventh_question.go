package main

import "fmt"

// Function untuk mencari bilangan genap dan jumlahnya
func jumlahGenap(arr []int) ([]int, int) {
    var genap []int
    total := 0
    for _, v := range arr {
        if v%2 == 0 {
            genap = append(genap, v)
            total += v
        }
    }
    return genap, total
}

func main() {
    // Data array
    angka := []int{15, 18, 3, 9, 6, 2, 12, 14}

    // Panggil fungsi
    bilGenap, total := jumlahGenap(angka)

    fmt.Println("Bilangan genap:", bilGenap)
    fmt.Println("Jumlah bilangan genap:", total)
}
