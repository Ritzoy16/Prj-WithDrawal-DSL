package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"prj_withdrawal/handlers"
	"runtime"
)

// Función para abrir el navegador automáticamente según el SO
func abrirNavegador(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("sistema operativo no soportado")
	}
	if err != nil {
		fmt.Println("⚠️ No se pudo abrir el navegador automáticamente. Por favor, entra a:", url)
	}
}

func main() {
	// Validar argumentos de CLI
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [user|admin]")
		os.Exit(1)
	}
	rol := os.Args[1]

	if rol == "user" {
		fmt.Println("🚀 Iniciando Servidor para USUARIO en http://localhost:8080")

		http.HandleFunc("/api/saldo", handlers.SaldoHandler)
		http.HandleFunc("/api/retirar", handlers.RetirarHandler)
		http.HandleFunc("/api/abonar", handlers.AbonarHandler)
		http.HandleFunc("/api/historial", handlers.HistorialUsuarioHandler) // Ruta filtrada

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "templates/index.html")
		})

		abrirNavegador("http://localhost:8080")

	} else if rol == "admin" {
		fmt.Println("🚀 Iniciando Servidor para ADMINISTRADOR en http://localhost:8081")

		http.HandleFunc("/api/admin/data", handlers.AdminDataHandler)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "templates/admin.html")
		})

		abrirNavegador("http://localhost:8081")
		// El admin usa otro puerto para no chocar con el usuario
		http.ListenAndServe(":8081", nil)
		return
	} else {
		fmt.Println("Rol no reconocido. Usa 'user' o 'admin'.")
		os.Exit(1)
	}

	http.ListenAndServe(":8080", nil)
}
