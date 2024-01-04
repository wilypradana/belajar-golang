package main

import "fmt"

func endApp()  {
	fmt.Println("Selesai memanggil function")
	message := recover()
	fmt.Println("TERJADI PANIC", message)
}

func runApplication(error bool)  {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("aplikasi dijalankan")
}

func main() {
	runApplication(true)
	fmt.Println("eko")
}