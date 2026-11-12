package handlers

import (
	"encoding/json"
	"net/http"
	"prj_withdrawal/headers"
	"strconv"
)

// "Base de datos" en memoria con múltiples usuarios
var Banco = map[string]*headers.Cuenta{
	"mauricio": headers.NewCuenta("Mauricio", 500.00),
	"alice":    headers.NewCuenta("Alice", 234.56),
	"hamza":    headers.NewCuenta("Hamza", 198.76),
}

// Usuario activo por defecto para la vista de usuario
const UsuarioActivo = "mauricio"

// Helper para responder en JSON
func responderJSON(w http.ResponseWriter, status int, datos interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(datos)
}

// SaldoHandler devuelve el saldo actual del usuario activo
func SaldoHandler(w http.ResponseWriter, r *http.Request) {
	cuenta := Banco[UsuarioActivo]
	responderJSON(w, http.StatusOK, map[string]interface{}{
		"nombre": cuenta.Nombre(),
		"saldo":  cuenta.ObtenerSaldo(),
	})
}

// RetirarHandler
func RetirarHandler(w http.ResponseWriter, r *http.Request) {
	monto, err := strconv.ParseFloat(r.FormValue("monto"), 64)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "Monto inválido"})
		return
	}

	cuenta := Banco[UsuarioActivo]
	exito, err := cuenta.Retirar(monto)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if !exito {
		responderJSON(w, http.StatusOK, map[string]interface{}{"status": "error", "mensaje": "Retiro rechazado: Fondos insuficientes."})
		return
	}

	responderJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "mensaje": "Retiro procesado con éxito.", "saldo": cuenta.ObtenerSaldo()})
}

// AbonarHandler
func AbonarHandler(w http.ResponseWriter, r *http.Request) {
	monto, err := strconv.ParseFloat(r.FormValue("monto"), 64)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "Monto inválido"})
		return
	}

	cuenta := Banco[UsuarioActivo]
	exito, err := cuenta.Abonar(monto)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if !exito {
		responderJSON(w, http.StatusOK, map[string]interface{}{"status": "error", "mensaje": "Abono rechazado."})
		return
	}

	responderJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "mensaje": "Abono procesado con éxito.", "saldo": cuenta.ObtenerSaldo()})
}

// HistorialUsuarioHandler (SOLO MUESTRA EXITOSOS)
func HistorialUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	cuenta := Banco[UsuarioActivo]
	todas := cuenta.ObtenerHistorial() // Necesitamos exponer esto en class.go

	var exitosas []headers.Transaccion
	for _, tx := range todas {
		if tx.Estado == "EXITOSA" {
			exitosas = append(exitosas, tx)
		}
	}

	if exitosas == nil {
		exitosas = make([]headers.Transaccion, 0)
	}
	responderJSON(w, http.StatusOK, exitosas)
}

// --- RUTAS DE ADMINISTRADOR ---

// AdminDataHandler devuelve TODA la info de todos los usuarios (Saldos e Historial completo)
func AdminDataHandler(w http.ResponseWriter, r *http.Request) {
	type UsuarioData struct {
		Nombre    string                `json:"nombre"`
		Saldo     float64               `json:"saldo"`
		Historial []headers.Transaccion `json:"historial"`
	}

	var reporte []UsuarioData
	for _, cuenta := range Banco {
		reporte = append(reporte, UsuarioData{
			Nombre:    cuenta.Nombre(),
			Saldo:     cuenta.ObtenerSaldo(),
			Historial: cuenta.ObtenerHistorial(), // Historial completo
		})
	}
	responderJSON(w, http.StatusOK, reporte)
}
