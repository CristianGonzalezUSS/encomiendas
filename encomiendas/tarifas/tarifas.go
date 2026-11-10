
package tarifas

import "fmt"

// TipoServicio representa la modalidad de envío.
type TipoServicio int

// TODO 1: constantes de los tipos de servicio.
const (
	Estandar TipoServicio = iota
	Express
	Prioritario
)

// TODO 2: constantes de las tarifas.
const (
	TarifaBase float64 = 8.00
	PrecioKg   float64 = 2.50
	PrecioKm   float64 = 0.15
	PesoMaximo float64 = 50.00
	IGV        float64 = 0.18
)

// CalcularSubtotal devuelve el costo sin IGV o un error.
func CalcularSubtotal(pesoKg float64, distanciaKm int, servicio TipoServicio) (float64, error) {

	// TODO 3: validar peso y distancia.
	if pesoKg <= 0 || pesoKg > PesoMaximo {
		return 0, fmt.Errorf(
			"peso no válido: %.2f kg (debe estar entre 0 y 50 kg)",
			pesoKg,
		)
	}

	if distanciaKm <= 0 {
		return 0, fmt.Errorf(
			"distancia no válida: %d km (debe ser mayor que 0)",
			distanciaKm,
		)
	}

	// TODO 4: obtener el recargo según el servicio.
	var recargo float64

	switch servicio {
	case Estandar:
		recargo = 1.0
	case Express:
		recargo = 1.3
	case Prioritario:
		recargo = 1.6
	default:
		return 0, fmt.Errorf("tipo de servicio no válido")
	}

	// Calcular subtotal.
	subtotal := (TarifaBase +
		pesoKg*PrecioKg +
		float64(distanciaKm)*PrecioKm) * recargo

	return subtotal, nil
}

// TODO 5: aplicar IGV usando retornos con nombre.
func AplicarIGV(subtotal float64) (igv, total float64) {
	igv = subtotal * IGV
	total = subtotal + igv

	return
}
