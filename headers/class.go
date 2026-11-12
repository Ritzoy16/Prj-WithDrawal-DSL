package headers

import (
	"encoding/json"
	"errors"
	"time"
)

type Transaccion struct {
	Tipo          string    `json:"tipo"`
	Monto         float64   `json:"monto"`
	Estado        string    `json:"estado"`
	SaldoAnterior float64   `json:"saldo_anterior"`
	SaldoNuevo    float64   `json:"saldo_nuevo"`
	Mensaje       string    `json:"mensaje,omitempty"`
	Fecha         time.Time `json:"fecha"`
}

// Campos en Mayúscula para que se guarden en JSON
type Cuenta struct {
	Saldo     float64       `json:"saldo"`
	Nombre    string        `json:"nombre"`
	Historial []Transaccion `json:"historial"`
}

func NewCuenta(nombre string, saldoInicial float64) *Cuenta {
	return &Cuenta{
		Nombre:    nombre,
		Saldo:     saldoInicial,
		Historial: make([]Transaccion, 0),
	}
}

func (c *Cuenta) ObtenerSaldo() float64 { return c.Saldo }
func (c *Cuenta) ObtenerNombre() string { return c.Nombre }

func (c *Cuenta) Retirar(cantidad float64) (bool, error) {
	saldoAnt := c.Saldo
	tx := Transaccion{Tipo: "RETIRO", Monto: cantidad, SaldoAnterior: saldoAnt, SaldoNuevo: saldoAnt, Fecha: time.Now()}

	if cantidad < 0 {
		tx.Estado = "ERROR"
		tx.Mensaje = "No se pueden retirar cantidades negativas"
		c.Historial = append(c.Historial, tx)
		return false, errors.New(tx.Mensaje)
	}
	if cantidad == 0 {
		tx.Estado = "ERROR"
		tx.Mensaje = "La cantidad a retirar debe ser mayor a cero"
		c.Historial = append(c.Historial, tx)
		return false, errors.New(tx.Mensaje)
	}
	if cantidad > c.Saldo {
		tx.Estado = "RECHAZADA"
		tx.Mensaje = "Saldo insuficiente"
		c.Historial = append(c.Historial, tx)
		return false, nil
	}

	c.Saldo -= cantidad
	tx.Estado = "EXITOSA"
	tx.SaldoNuevo = c.Saldo
	tx.Mensaje = "Retiro procesado correctamente"
	c.Historial = append(c.Historial, tx)
	return true, nil
}

func (c *Cuenta) Abonar(cantidad float64) (bool, error) {
	saldoAnt := c.Saldo
	tx := Transaccion{Tipo: "ABONO", Monto: cantidad, SaldoAnterior: saldoAnt, SaldoNuevo: saldoAnt, Fecha: time.Now()}

	if cantidad <= 0 {
		tx.Estado = "ERROR"
		tx.Mensaje = "La cantidad a abonar debe ser mayor a cero"
		c.Historial = append(c.Historial, tx)
		return false, errors.New(tx.Mensaje)
	}

	c.Saldo += cantidad
	tx.Estado = "EXITOSA"
	tx.SaldoNuevo = c.Saldo
	tx.Mensaje = "Abono procesado correctamente"
	c.Historial = append(c.Historial, tx)
	return true, nil
}

func (c *Cuenta) ObtenerHistorial() []Transaccion {
	return c.Historial
}

func (c *Cuenta) ObtenerHistorialJSON() string {
	jsonData, _ := json.MarshalIndent(c.Historial, "", "  ")
	return string(jsonData)
}
