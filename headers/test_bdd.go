package headers

import (
	"fmt"
	"math"
)

// imprimirSaldo imprime el saldo formateado a 2 decimales.
func imprimirSaldo(cuenta *Cuenta) {
	fmt.Printf("Saldo de %.2f €\n", cuenta.ObtenerSaldo())
}

// Validación de Escenarios Gherkin Dinámicos

// TestEscenario1Exito: Simula un retiro exitoso con datos dinamicos
func TestEscenario1Exito(nombre string, saldoInicial float64, cantidadRetirar float64) {
	fmt.Printf("--- Escenario 1: Successful withdrawal within balance (%s) ---\n", nombre)

	// Given
	cuenta := NewCuenta(nombre, saldoInicial)
	fmt.Printf("Given: %s tiene %.2f en su cuenta\n", nombre, saldoInicial)

	// When
	fmt.Printf("When: %s intenta retirar %.2f\n", nombre, cantidadRetirar)
	exito, err := cuenta.Retirar(cantidadRetirar)

	// Then
	if err != nil {
		fmt.Printf("Resultado: FALLO. Error inesperado: %v\n\n", err)
		return
	}

	if exito {
		fmt.Println("Resultado: ÉXITO. La retirada fue procesada.")
		imprimirSaldo(cuenta)
		saldoEsperado := saldoInicial - cantidadRetirar
		if math.Abs(cuenta.ObtenerSaldo()-saldoEsperado) < 0.001 {
			fmt.Printf("Validación: El saldo restante es correcto (%.2f).\n\n", saldoEsperado)
		} else {
			fmt.Println("ERROR: El saldo restante no coincide.\n")
		}
	} else {
		fmt.Println("Resultado: FALLO. La retirada fue rechazada (incorrecto).\n")
	}

	// Imprimir el JSON de auditoría al finalizar el escenario
	fmt.Println("--- LOG DE AUDITORÍA (JSON) ---")
	fmt.Println(cuenta.ObtenerHistorialJSON())
	fmt.Println() // Salto de línea extra para separar
}

// TestEscenario2Fallo: Simula un retiro rechazado con datos dinámicos
func TestEscenario2Fallo(nombre string, saldoInicial float64, cantidadRetirar float64) {
	fmt.Printf("--- Escenario 2: Declined withdrawal in excess of balance (%s) ---\n", nombre)

	// Given
	cuenta := NewCuenta(nombre, saldoInicial)
	fmt.Printf("Given: %s tiene %.2f en su cuenta\n", nombre, saldoInicial)

	// When
	fmt.Printf("When: %s intenta retirar %.2f\n", nombre, cantidadRetirar)
	exito, err := cuenta.Retirar(cantidadRetirar)

	// Then
	if err != nil {
		fmt.Printf("Resultado: FALLO. Error inesperado: %v\n\n", err)
		return
	}

	if !exito {
		fmt.Println("Resultado: ÉXITO. La retirada fue rechazada correctamente.")
		imprimirSaldo(cuenta)
		if math.Abs(cuenta.ObtenerSaldo()-saldoInicial) < 0.001 {
			fmt.Printf("Validación: El saldo no ha cambiado (%.2f).\n\n", saldoInicial)
		} else {
			fmt.Println("ERROR: El saldo se modificó incorrectamente.\n")
		}
	} else {
		fmt.Println("Resultado: FALLO. La retirada fue procesada (debería haber sido rechazada).\n")
	}

	// Imprimir el JSON de auditoría al finalizar el escenario
	fmt.Println("--- LOG DE AUDITORÍA (JSON) ---")
	fmt.Println(cuenta.ObtenerHistorialJSON())
	fmt.Println() // Salto de línea extra para separar
}

// TestEscenario3Error: Simula un error de validación (ej. cantidad negativa)
func TestEscenario3Error(nombre string, saldoInicial float64, cantidadRetirar float64) {
	fmt.Printf("--- Escenario 3: Validation Error (%s) ---\n", nombre)
	cuenta := NewCuenta(nombre, saldoInicial)

	exito, err := cuenta.Retirar(cantidadRetirar)
	if err != nil {
		fmt.Printf("Resultado: ÉXITO. El sistema atrapó el error correctamente: %v\n", err)
		imprimirSaldo(cuenta)
	} else if !exito {
		fmt.Println("Resultado: RECHAZADO. Fondos insuficientes (no es un error de validación).\n")
	} else {
		fmt.Println("Resultado: FALLO. Se procesó un retiro que debía dar error.\n")
	}

	// Imprimir el JSON de auditoría al finalizar el escenario
	fmt.Println("--- LOG DE AUDITORÍA (JSON) ---")
	fmt.Println(cuenta.ObtenerHistorialJSON())
	fmt.Println() // Salto de línea extra para separar
}

//TestEscenarioAbonoExito: Simula un abono exitoso en efectivo
func TestEscenarioAbonoExito(nombre string, saldoInicial float64, cantidadAbonar float64){
	fmt.Printf("--- Escenario US-01b: Succesful cash deposit (%s) --- \n", nombre)

	// Given
	cuenta := NewCuenta(nombre, saldoInicial)
	fmt.Printf("Given: %s tiene %.2f en su cuenta \n", nombre, saldoInicial)

	//When
	fmt.Printf("When: %s intenta abonar %.2f\n", nombre, cantidadAbonar)
	exito, err := cuenta.Abonar(cantidadAbonar)

	//Then
	if err != nil{
		fmt.Printf("Resultado: FALLO. Error inesperado: %v\n\n", err)
		return
	}

	if exito {
		fmt.Printf("Resultado: ÉXITO. El abono fue procesado")
		imprimirSaldo(cuenta)
		saldoEsperado := saldoInicial + cantidadAbonar
		if math.Abs(cuenta.ObtenerSaldo() - saldoEsperado) < 0.001 {
			fmt.Printf("Validación: El saldo incrementado es correcto (%.2f).\n\n", saldoEsperado)
		} else {
			fmt.Println("ERROR: El saldo restante no coincide.\n")
		}
	} else {
		fmt.Println("Resultado: FALLO. El abono fue rechazado (incorrecto).\n")
	}

	// Imprimir el JSON de auditoría al finalizar el escenario
	fmt.Println("--- LOG DE AUDITORÍA (JSON) ---")
	fmt.Println(cuenta.ObtenerHistorialJSON())
	fmt.Println() // Salto de línea extra para separar
}
