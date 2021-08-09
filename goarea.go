package goarea

import "math"

// Pi é uma proporção númerica
const Pi = 3.1416

//Circ ...
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect ...
func Rect(base, altura float64) float64 {
	return base * altura
}

func _RetanguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
 