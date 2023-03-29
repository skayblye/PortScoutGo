package main

import (
	"fmt"
	"net"
	"sync"
)

func ScanPort(ip string, port int, wg *sync.WaitGroup) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		// El puerto está cerrado o no se pudo conectar
		wg.Done() // indicar que se completó un escaneo de puerto
		return
	}
	// El puerto está abierto
	fmt.Printf("El puerto %d está abierto en la dirección IP %s\n", port, ip)
	conn.Close()
	wg.Done() // indicar que se completó un escaneo de puerto
}

func main() {
	ip := "127.0.0.1"                     // La dirección IP que deseas escanear
	var wg sync.WaitGroup                 // Crear la espera grupal
	for port := 1; port <= 1024; port++ { // Escanea los primeros 1024 puertos
		wg.Add(1) // Agregar un escaneo de puerto a la espera grupal
		go ScanPort(ip, port, &wg)
	}
	wg.Wait() // Esperar a que se completen todos los escaneos de puerto
	fmt.Println("Todos los escaneos de puerto han finalizado.")
}

// hola
