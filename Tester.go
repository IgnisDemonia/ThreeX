package main

/* Решение кубического уравнения при помощи метода Кардано*/
import (
	"fmt"
	"math"
)

type CubicSolution struct {
	X1 complex128
	X2 complex128
	X3 complex128
}

func SolveCubic(a, b, c, d float64) CubicSolution {
	p := (3*a*c - b*b) / (3 * a * a)
	q := (27*a*a*d - 9*a*b*c + 2*b*b*b) / (27 * a * a * a)

	D := q*q/4 + p*p*p/27

	var solutions CubicSolution

	if D > 0 {
		sqrtD := math.Sqrt(D)
		u := math.Cbrt(-q/2 + sqrtD)
		v := math.Cbrt(-q/2 - sqrtD)
		solutions.X1 = complex(u+v, 0)
	} else if D == 0 {
		solutions.X1 = complex(-1.5*q, 0)
		solutions.X2 = complex(0.75*q/a, 0)
	} else {
		r := math.Sqrt(-p * p * p / 27)
		theta := math.Acos(-q / (2 * math.Pow(r, 3)))

		solutions.X1 = complex(2*math.Cbrt(r)*math.Cos(theta/3)-b/(3*a), 0)
		solutions.X2 = complex(2*math.Cbrt(r)*math.Cos((theta+2*math.Pi)/3)-b/(3*a), 0)
		solutions.X3 = complex(2*math.Cbrt(r)*math.Cos((theta+4*math.Pi)/3)-b/(3*a), 0)
	}

	return solutions
}

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

	solutions := SolveCubic(a, b, c, d)

	fmt.Printf("X1: %.2f\nX2: %.2f\nX3: %.2f\n", real(solutions.X1), real(solutions.X2), real(solutions.X3))
}
