package main

import (
	"fmt"
)

var notes []string
var loaded bool = false

func addNote() {
	var note string
	fmt.Print("Masukkan note: ")
	fmt.Scanln(&note)
	notes = append(notes, note)
}

func removeNote() {
	var noteID int
	fmt.Print("Masukkan id note yang ingin dihapus: ")
	fmt.Scanf("%d", &noteID)
	if !(1 <= noteID && noteID <= cap(notes)) {
		fmt.Println("Index tidak diketahui")
	} else {
		fmt.Printf("Menghapus indeks %d...\n", noteID)
		notes[noteID-1] = notes[len(notes)-1]
		notes[len(notes)-1] = ""
		notes = notes[:len(notes)-1]
		fmt.Println("Berhasil menghapus!")
	}
}

func listNotes() {
	fmt.Println("List notes:")
	for i, note := range notes {
		fmt.Printf("%d\t%s\n", i+1, note)
	}
	if len(notes) == 0 {
		fmt.Println("Tidak ada.")
	}
}

func userSelection() bool {
	fmt.Print("Pilih operasi t[a]mbah ku[r]angi [l]ihat k[e]luar: ")
	var selection string
	fmt.Scanln(&selection)
	switch selection {
	case string("a"):
		addNote()
	case string("r"):
		removeNote()
	case string("l"):
		listNotes()
	case string("e"):
		fmt.Println("Keluar...")
		return false
	default:
		fmt.Println("Tidak valid, ulangi lagi")
	}
	return true
}

func main() {
	for userSelection() {
	}
}
