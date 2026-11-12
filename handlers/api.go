package handlers

import (
	"encoding/json"
	"net/http"
	"prj_withdrawal/headers"
	"strconv"
)

const UsuarioActivo = "mauricio"

func responderJSON(w http.ResponseWriter, status int, datos interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(datos)
}

func SaldoHandler(w http.ResponseWriter, r *http.Request) {
	CargarDB() // <-- Leer datos frescos
	cuenta := Banco[UsuarioActivo]
	responderJSON(w, http.StatusOK, map[string]interface{}{
		"nombre": cuenta.ObtenerNombre(),
		"saldo":  cuenta.ObtenerSaldo(),
	})
}

func RetirarHandler(w http.ResponseWriter, r *http.Request) {
	CargarDB() // <-- Leer datos
	monto, err := strconv.ParseFloat(r.FormValue("monto"), 64)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "Monto inválido"})
		return
	}

	cuenta := Banco[UsuarioActivo]
	exito, err := cuenta.Retirar(monto)
	GuardarDB() // <-- Guardar cambios SIEMPRE

	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if !exito {
		responderJSON(w, http.StatusOK, map[string]interface{}{"status": "error", "mensaje": "Retiro rechazado: Fondos insuficientes."})
		return
	}
	responderJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "mensaje": "Retiro exitoso.", "saldo": cuenta.ObtenerSaldo()})
}

func AbonarHandler(w http.ResponseWriter, r *http.Request) {
	CargarDB()
	monto, err := strconv.ParseFloat(r.FormValue("monto"), 64)
	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": "Monto inválido"})
		return
	}

	cuenta := Banco[UsuarioActivo]
	exito, err := cuenta.Abonar(monto)
	GuardarDB() // <-- Guardar cambios SIEMPRE

	if err != nil {
		responderJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if !exito {
		responderJSON(w, http.StatusOK, map[string]interface{}{"status": "error", "mensaje": "Abono rechazado."})
		return
	}
	responderJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "mensaje": "Abono exitoso.", "saldo": cuenta.ObtenerSaldo()})
}

func HistorialUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	CargarDB()
	cuenta := Banco[UsuarioActivo]
	todas := cuenta.ObtenerHistorial()

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

func AdminDataHandler(w http.ResponseWriter, r *http.Request) {
	CargarDB() // <-- El admin siempre lee del archivo, viendo los cambios en tiempo real
	type UsuarioData struct {
		Nombre    string                `json:"nombre"`
		Saldo     float64               `json:"saldo"`
		Historial []headers.Transaccion `json:"historial"`
	}

	var reporte []UsuarioData
	for _, cuenta := range Banco {
		reporte = append(reporte, UsuarioData{
			Nombre:    cuenta.ObtenerNombre(),
			Saldo:     cuenta.ObtenerSaldo(),
			Historial: cuenta.ObtenerHistorial(),
		})
	}
	responderJSON(w, http.StatusOK, reporte)
}
