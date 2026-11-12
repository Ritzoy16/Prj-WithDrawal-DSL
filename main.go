package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"prj_withdrawal/handlers"
	"runtime"
)

func abrirNavegador(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Println("Abre manualmente:", url)
	}
}

func main() {
	handlers.CargarDB()

	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [user|admin]")
		os.Exit(1)
	}
	rol := os.Args[1]

	if rol == "user" {
		fmt.Println("Servidor USUARIO en http://localhost:8080")
		http.HandleFunc("/api/saldo", handlers.SaldoHandler)
		http.HandleFunc("/api/retirar", handlers.RetirarHandler)
		http.HandleFunc("/api/abonar", handlers.AbonarHandler)
		http.HandleFunc("/api/historial", handlers.HistorialUsuarioHandler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "templates/index.html")
		})
		abrirNavegador("http://localhost:8080")
		http.ListenAndServe(":8080", nil)

	} else if rol == "admin" {
		fmt.Println("Servidor ADMIN en http://localhost:8081")
		http.HandleFunc("/api/admin/data", handlers.AdminDataHandler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "templates/admin.html")
		})
		abrirNavegador("http://localhost:8081")
		http.ListenAndServe(":8081", nil)
	}
}
