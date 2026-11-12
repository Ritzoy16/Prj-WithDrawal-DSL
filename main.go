package main

import (
	"fmt"
	"net/http"
	"prj_withdrawal/handlers"
)

func main() {
	// 1. Rutas para las acciones del banco
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/retirar", handlers.RetirarHandler)
	http.HandleFunc("/abonar", handlers.AbonarHandler)

	// 2. Ruta para ver el JSON de auditoría
	http.HandleFunc("/api/historial", handlers.JSONHandler)

	// 3. Servir archivos estáticos (CSS, imágenes)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Servidor Web corriendo en http://localhost:8080")
	// Iniciar el servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}
