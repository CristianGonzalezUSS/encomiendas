
package main

import (
	"fmt"

	"encomiendas/tarifas"
)

// Registrar un envío.
func registrarEnvio(codigo string, peso float64, km int, servicio tarifas.TipoServicio) {

	fmt.Printf("\n[Inicio] Envío %s\n", codigo)

	// TODO 6: defer imprime el mensaje al finalizar.
	defer fmt.Printf("[Fin] Envío %s procesado\n", codigo)

	subtotal, err := tarifas.CalcularSubtotal(peso, km, servicio)

	if err != nil {
		fmt.Println(" Error:", err)
		return
	}

	igv, total := tarifas.AplicarIGV(subtotal)

	fmt.Printf(
		" Subtotal: S/ %.2f | IGV: S/ %.2f | Total: S/ %.2f\n",
		subtotal, igv, total,
	)
}

func main() {

	// TODO 7: declarar constante y variable.
	const empresa = "Encomiendas USS"
	totalEnvios := 3

	fmt.Printf("%s - envíos a registrar: %d\n", empresa, totalEnvios)

	// TODO 8: registrar los tres envíos.
	registrarEnvio("ENC-001", 3.5, 210, tarifas.Estandar)

	registrarEnvio("ENC-002", 12, 770, tarifas.Express)

	registrarEnvio("ENC-003", 60, 100, tarifas.Prioritario)
}
