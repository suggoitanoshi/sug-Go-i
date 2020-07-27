package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	// menerima init, lalu variabel yang akan di switch
	switch os := runtime.GOOS; os {
	// tidak semua kasus dijalankan, hanya 1 lalu keluar.
	case "darwin":
		fmt.Println("OS X.")
		// perintah break otomatis dimasukkan
	case "linux":
		fmt.Println("Linux.")
		// kasus yang tidak tertangani masuk sini
	default:
		fmt.Printf("%s.\n", os)
	}
	// dapat digunakan tanpa kondisi
	// sebagai satu cara menulis if-else secara rapi
	i := 69
	switch {
	case i == 420:
		fmt.Println("Nice")
	case i == 69:
		fmt.Println("Noice")
	default:
		fmt.Println("Not Nice")
	}
}
