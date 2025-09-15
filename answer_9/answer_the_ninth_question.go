package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// Fungsi untuk memformat JSON dari file input ke file output
func formatJSON(inputPath, outputPath string) error {
    // Baca file input
    data, err := ioutil.ReadFile(inputPath)
    if err != nil {
        return fmt.Errorf("gagal membaca file: %v", err)
    }

    // Decode JSON ke generic map
    var obj interface{}
    if err := json.Unmarshal(data, &obj); err != nil {
        return fmt.Errorf("gagal decode JSON: %v", err)
    }

    // Encode ulang dengan indentasi agar rapi
    formatted, err := json.MarshalIndent(obj, "", "  ")
    if err != nil {
        return fmt.Errorf("gagal encode JSON: %v", err)
    }

    // Simpan ke file output
    if err := os.WriteFile(outputPath, formatted, 0644); err != nil {
        return fmt.Errorf("gagal menulis file: %v", err)
    }

    return nil
}

func main() {
    input := "assets/json/case.json"
    output := "assets/json/expectation.json"

    if err := formatJSON(input, output); err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Println("File berhasil diformat:", output)
}
