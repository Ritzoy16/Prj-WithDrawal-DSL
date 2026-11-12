package handlers

import (
	"encoding/json"
	"net/http"
	"prj_withdrawal/headers"
	"strconv"
)

// Instancia global para la cuenta
var MiCuenta = headers.NewCuenta("Mauricio", 500.00)

// Helper para responder en JSON
func responderJSON(w http.ResponseWriter, status int, datos interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(datos)
}

// SaldoHandler devuelve el saldo actual
func SaldoHandler(w http.ResponseWriter, r *http.Request) {
	responderJSON(w, http.StatusOK, map[string]interface{}{
		"nombre": MiCuenta.Nombre(),
		"saldo":  MiCuenta.ObtenerSaldo(),
	})
}

// RetirarHandler procesa el formulario de retiro
func RetirarHandler(w http.ResponseWriter, r *http.Request) {
	monto, err := strconv.ParseFloat(r.FormValue("monto"), 64)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "Monto inválido"})
		return
	}

	exito, err := MiCuenta.Retirar(monto)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if !exito {
		responderJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "error",
			"mensaje": "Retiro rechazado: Fondos insuficientes.",
		})
		return
	}

	responderJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"mensaje": "Retiro procesado con éxito.",
		"saldo":   MiCuenta.ObtenerSaldo(),
	})
}

// AbonarHandler procesa el formulario de abono
func AbonarHandler(w http.ResponseWriter, r *http.Request) {
	monto, err := strconv.ParseFloat(r.FormValue("monto"), 64)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "Monto inválido"})
		return
	}

	exito, err := MiCuenta.Abonar(monto)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if !exito {
		responderJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "error",
			"mensaje": "Abono rechazado.",
		})
		return
	}

	responderJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"mensaje": "Abono procesado con éxito.",
		"saldo":   MiCuenta.ObtenerSaldo(),
	})
}

// HistorialHandler devuelve el JSON crudo de las transacciones
func HistorialHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(MiCuenta.ObtenerHistorialJSON()))
}
