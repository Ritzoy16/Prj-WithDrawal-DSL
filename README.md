# Prj-WithDrawal-DSL
**Alumno:** Chinchayhuara Pantoja, Mauricio Miguel
**Curso:** Lenguaje de Dominio Especifico

Este proyecto es una aplicación web desarrollada en **Go (Golang)** que simula el sistema de un cajero automático bancario. Fue construido aplicando la metodología **BDD (Behavior-Driven Development)** y validando los escenarios especificados en **Gherkin**. 

La aplicación implementa dos roles distintos (Usuario y Administrador), una interfaz gráfica moderna con pantallas de carga, notificaciones flotantes, y un sistema de auditoría persistente en formato JSON.

## Estructura del Proyecto

```text
prj_withdrawal/
│
├── main.go                 # Punto de entrada y enrutador CLI (user/admin)
├── go.mod                  # Módulo de Go
├── datos_banco.json        # Base de datos local (se genera automáticamente al correr por primera vez)
│
├── headers/                # Capa de Dominio (Reglas de Negocio)
│   ├── class.go            # Entidad Cuenta, Transacciones y lógica JSON
│   └── test_bdd.go         # Pruebas de escenarios BDD (CLI)
│
├── handlers/               # Capa de Controladores (API REST)
│   ├── api.go              # Endpoints HTTP para retiros, abonos y datos
│   └── database.go         # Lógica de lectura/escritura del archivo JSON
│
└── templates/              # Capa de Presentación (Frontend)
    ├── index.html          # Vista del Usuario
    └── admin.html          # Vista del Administrador
```

## Requisitos Previos

- Tener Go instalado

## Cómo ejecutar

Abrir la terminal, navegar hasta la carpeta raíz del proyecto (prj_withdrawal) y ejecutar los siguientes comandos según el rol que quieras probar. El programa abrirá el navegador web automáticamente.

1. Ejecutar como USUARIO

Iniciar el servidor del cajero automático y abrir `http://localhost:8080`
```bash
go run main.go user
```

2. Ejecutar como ADMINISTRADOR

Abrir una segunda pestaña en tu terminal (dejar la del usuario corriendo) e iniciar el panel de administración en `http://localhost:8081`
```bash
go run main.go admin
```
- Pruebas de Integración (Opcional)

Si se desea ver las pruebas por consola (TDD/BDD) que validen la lógica de negocio sin usar la interfaz web, se puede ejecutar:
```bash
go run headers/test_bdd.go
```

- ¿Cómo manipular los datos?

Si se desea poder cambiar los valores de los datos, se puede modificar los datos desde `datos_banco.json`. Si se desea borrar el historial y restablecer los datos a sus valores iniciales, basta con eliminar el archivo `datos_banco.json`. Como este programa es solo una simulación, no hay limitación en base a eso.