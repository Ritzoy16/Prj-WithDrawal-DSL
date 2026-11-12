package headers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time" //Para registrar la fecha de la transacción
)

// Transaccion reperesenta el log de auditoría
type Transaccion struct {
	Tipo          string    `json:"tipo"`              // RETIRO, ABONO
	Monto         float64   `json:"monto"`             // Cantidad solicitada
	Estado        string    `json:"estado"`            // EXITOSA, RECHAZADA, ERROR
	SaldoAnterior float64   `json:"saldo_anterior"`    // Saldo antes de la operacion
	SaldoNuevo    float64   `json:"saldo_nuevo"`       // Saldo después de la operación
	Mensaje       string    `json:"mensaje_omitempty"` // Detalle del error o éxito
	Fecha         time.Time `json:"fecha"`             // Timestamp
}

// Cuenta representa la entidad de negocio
type Cuenta struct {
	saldo     float64
	nombre    string
	historial []Transaccion //NUEVO: Historial de Transacciones
}

func NewCuenta(nombre string, saldoInicial float64) *Cuenta {
	return &Cuenta{
		nombre:    nombre,
		saldo:     saldoInicial,
		historial: make([]Transaccion, 0), // Inicializar historial
	}
}

func (c *Cuenta) ObtenerSaldo() float64 {
	return c.saldo
}

func (c *Cuenta) Nombre() string {
	return c.nombre
}

/*
Retirar intenta retirar una cantidad.
Regla: No se puede retirar más de lo que hay en la cuenta.
*/
func (c *Cuenta) Retirar(cantidad float64) (bool, error) {
	saldoAnt := c.saldo
	//Creación de registro de la transacción
	tx := Transaccion{
		Tipo:          "RETIRO",
		Monto:         cantidad,
		SaldoAnterior: saldoAnt,
		SaldoNuevo:    saldoAnt,
		Fecha:         time.Now(),
	}
	if cantidad < 0 {
		tx.Estado = "ERROR"
		tx.Mensaje = "No se puede retirar cantidades negativas"
		c.historial = append(c.historial, tx)
		return false, errors.New(tx.Mensaje)
	}
	if cantidad > c.saldo {
		tx.Estado = "RECHAZADA"
		tx.Mensaje = "Saldo Insuficiente"
		c.historial = append(c.historial, tx)
		return false, nil // Escenario: Withdrawal declined
	}

	//Si se llega hasta quí, la transacción es exitosa
	c.saldo -= cantidad
	tx.Estado = "EXITOSA"
	tx.SaldoNuevo = c.saldo
	tx.Mensaje = "Retiro procesado correctamente"
	c.historial = append(c.historial, tx)
	return true, nil // Escenario: Successful withdrawal
}

func (c *Cuenta) String() string {
	return fmt.Sprintf("Cuenta{%s, saldo=%.2f}", c.nombre, c.saldo)
}

// Abonar intenta depositar una cantidad en la cuenta
// Regla: No se pueden abonar cantidades negativas o cero
func (c *Cuenta) Abonar(cantidad float64) (bool, error) {
	saldoAnt := c.saldo
	tx := Transaccion{
		Tipo:          "ABONO",
		Monto:         cantidad,
		SaldoAnterior: saldoAnt,
		SaldoNuevo:    saldoAnt,
		Fecha:         time.Now(),
	}

	if cantidad <= 0 {
		tx.Estado = "ERROR"
		tx.Mensaje = "La cantidad a abonar debe ser mayor a cero"
		c.historial = append(c.historial, tx)
		return false, errors.New(tx.Mensaje)
	}
	c.saldo += cantidad
	tx.Estado = "EXITOSA"
	tx.SaldoNuevo = c.saldo
	tx.Mensaje = "Abono procesado correctamente"
	c.historial = append(c.historial, tx)
	return true, nil //Escenario: Succesful deposit
}

// ObtenerHistorialJSON devuelve el historial de transacciones en formato JSON
func (c *Cuenta) ObtenerHistorialJSON() string {
	jsonData, err := json.MarshalIndent(c.historial, "", " ")
	if err != nil {
		return fmt.Sprintf("Error generando JSON: %v", err)
	}
	return string(jsonData)
}

// ObtenerHistorial devuelve el slice de transacciones (necesario para el admin y filtros)
func (c *Cuenta) ObtenerHistorial() []Transaccion {
	return c.historial
}
