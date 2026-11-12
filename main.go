package main

import (
	"flag"
	"fmt"
	"os"

	"prj_withdrawal/headers"
)

func main() {
	// Definir banderas (flags) para los argumentos de línea de comandos
	nombre := flag.String("nombre", "", "Nombre del titular de la cuenta")
	saldo := flag.Float64("saldo", 0.0, "Saldo inicial de la cuenta")
	retiro := flag.Float64("retiro", 0.0, "Cantidad a retirar")
	// Bandera para abonos
	abono := flag.Float64("abono", 0.0, "Cantidad a abonar en efectivo")
	// Bandera especial para forzar el escenario de test
	escenario := flag.String("test", "", "Tipo de test a ejecutar: 'exito', 'fallo', 'error'")

	flag.Parse()

	// Si no se pasan argumentos, ejecutamos los escenarios por defecto del examen
	if len(os.Args) == 1 {
		fmt.Println("=== EJECUTANDO PRUEBAS POR DEFECTO (Gherkin v1.0.0) ===")
		headers.TestEscenario1Exito("Alice", 234.56, 200.00)
		headers.TestEscenario2Fallo("Hamza", 198.76, 200.00)
		//Nueva prueba por defecto:
		headers.TestEscenarioAbonoExito("Carlos", 60.00, 100.00)
		os.Exit(0)
	}

	// Validar que se hayan ingresado los datos necesarios
	if *nombre == "" {
		fmt.Println("Error: Debes proporcionar al menos un -nombre. Usa -h para ayuda.")
		os.Exit(1)
	}

	// Enrutamiento de escenarios
	switch *escenario {
	case "exito":
		headers.TestEscenario1Exito(*nombre, *saldo, *retiro)
	case "fallo":
		headers.TestEscenario2Fallo(*nombre, *saldo, *retiro)
	case "error":
		headers.TestEscenario3Error(*nombre, *saldo, *retiro)
	case "abono":
		headers.TestEscenarioAbonoExito(*nombre, *saldo, *abono)
	case "json":
		//Ejecuta una transaccion y muestra solo el JSON
		cuenta := headers.NewCuenta(*nombre, *saldo)
		cuenta.Retirar(*retiro) //Hace un intento
		fmt.Println(cuenta.ObtenerHistorialJSON())
	default:
		// Si no se especifica escenario, deducimos qué pasará según los datos
		fmt.Println("=== EJECUTANDO TEST DINÁMICO AUTOMÁTICO ===")
		// NUEVO: Si se envía un abono mayor a 0, ejecutamos el escenario de abono
		if *abono > 0 {
			headers.TestEscenarioAbonoExito(*nombre, *saldo, *abono)
		} else if *retiro < 0 {
			headers.TestEscenario3Error(*nombre, *saldo, *retiro)
		} else if *retiro > *saldo {
			headers.TestEscenario2Fallo(*nombre, *saldo, *retiro)
		} else if *retiro > 0 {
			headers.TestEscenario1Exito(*nombre, *saldo, *retiro)
		} else {
			fmt.Println("Por favor, especifica un -retiro o -abono mayor a 0.")
		}

	os.Exit(0)
	}
}
