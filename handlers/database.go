package handlers

import (
	"encoding/json"
	"os"
	"prj_withdrawal/headers"
)

const ArchivoDB = "datos_banco.json"

// El mapa global que compartirán ambos servidores
var Banco = make(map[string]*headers.Cuenta)

// Guarda el estado actual de la memoria al archivo JSON
func GuardarDB() {
	data, err := json.MarshalIndent(Banco, "", "  ")
	if err == nil {
		os.WriteFile(ArchivoDB, data, 0644)
	}
}

// Lee el archivo JSON y lo carga a la memoria
func CargarDB() {
	data, err := os.ReadFile(ArchivoDB)
	if err == nil {
		json.Unmarshal(data, &Banco)
	} else {
		// Si el archivo no existe, creamos los datos iniciales
		Banco = map[string]*headers.Cuenta{
			"mauricio": headers.NewCuenta("Mauricio", 500.00),
			"alice":    headers.NewCuenta("Alice", 234.56),
			"hamza":    headers.NewCuenta("Hamza", 198.76),
		}
		GuardarDB()
	}
}
