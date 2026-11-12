package main

import (
	"fmt"
	"net/http"
	"prj_withdrawal/handlers"
)

func main() {
	// Rutas de la API (JSON)
	http.HandleFunc("/api/saldo", handlers.SaldoHandler)
	http.HandleFunc("/api/retirar", handlers.RetirarHandler)
	http.HandleFunc("/api/abonar", handlers.AbonarHandler)
	http.HandleFunc("/api/historial", handlers.HistorialHandler)

	// Ruta para la página principal (HTML)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	fmt.Println("Servidor Web corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
