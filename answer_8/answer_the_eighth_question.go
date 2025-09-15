package main

import (
    "fmt"
    "strings"
)

// Fungsi untuk cek apakah dua string adalah anagram
func isAnagram(s1, s2 string) bool {
    s1 = strings.ToLower(s1)
    s2 = strings.ToLower(s2)

    if len(s1) != len(s2) {
        return false
    }

    count1 := make(map[rune]int)
    count2 := make(map[rune]int)

    for _, c := range s1 {
        count1[c]++
    }
    for _, c := range s2 {
        count2[c]++
    }

    // bandingkan jumlah huruf
    for k, v := range count1 {
        if count2[k] != v {
            return false
        }
    }

    return true
}

func main() {
    // Inisiasi sendiri string untuk dites
    str1 := "listen"
    str2 := "silent"

    if isAnagram(str1, str2) {
        fmt.Printf("%q dan %q adalah anagram.\n", str1, str2)
    } else {
        fmt.Printf("%q dan %q bukan anagram.\n", str1, str2)
    }
}
