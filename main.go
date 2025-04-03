package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID   int
	Name string
}

var tasks []Task
var nextID = 1

func main() {
	for {
		fmt.Println("\n=== TO-DO LIST CLI ===")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Tugas")
		fmt.Println("3. Hapus Tugas")
		fmt.Println("4. Keluar")
		fmt.Println("Pilih Opsi: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask()
		case "2":
			listTasks()
		case "3":
			deleteTask()
		case "4":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini")
			return
		default:
			fmt.Println("pilihan tidak valid, coba lagi.")

		}
	}
}

func addTask() {
	fmt.Print("masukan nama tugas: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskName := scanner.Text()

	newTask := Task{ID: nextID, Name: taskName}
	tasks = append(tasks, newTask)
	nextID++

	fmt.Println("tugas berhasil ditambahkan!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("tidak ada tugas dalam daftar")
		return
	}
	fmt.Println("\nDaftar Tugas:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Name)
	}
}

func deleteTask() {
	if len(tasks) == 0 {
		fmt.Println("tidak ada tugas dalam daftar")
		return
	}
	fmt.Print("masukan id tugas yang akan dihapus")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idStr := scanner.Text()

	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Id tidak valid")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tugas berhasil dihapus.")
			return
		}
	}
	fmt.Println("Tugas dengan ID tersebut tidak ditemukan.")
}
