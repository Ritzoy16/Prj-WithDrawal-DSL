package handlers

import (
	"fmt"
	"net/http"
	"prj_withdrawal/headers"
	"strconv"
)

// Instancia global de la cuenta para el ejemplo
var MiCuenta = headers.NewCuenta("Mauricio", 500.00)

func parseMonto(w http.ResponseWriter, r *http.Request) (float64, bool) {
	montoStr := r.FormValue("monto")
	monto, err := strconv.ParseFloat(montoStr, 64)
	if err != nil || monto <= 0 {
		http.Error(w, "Monto inválido", http.StatusBadRequest)
		return 0, false
	}
	return monto, true
}

// IndexHandler muestra la página principal
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// RetirarHandler procesa el formulario de retiro
func RetirarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	monto, ok := parseMonto(w, r)
	if !ok {
		return
	}

	exito, err := MiCuenta.Retirar(monto)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if !exito {
		fmt.Fprintf(w, "Retiro rechazado: Fondos insuficientes.")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// AbonarHandler procesa el formulario de abono
func AbonarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	monto, ok := parseMonto(w, r)
	if !ok {
		return
	}

	exito, err := MiCuenta.Abonar(monto)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if !exito {
		fmt.Fprintf(w, "Abono rechazado.")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// JSONHandler devuelve el historial en formato JSON
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(MiCuenta.ObtenerHistorialJSON()))
}
