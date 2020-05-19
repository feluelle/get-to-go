package exercises

import (
	"fmt"
)

// Sqrt http://127.0.0.1:3999/moretypes/18
func Sqrt(x float64) float64 {
	z := 1.0
	for prev := z; ; prev = z {
		z -= (z*z - x) / (2 * z)
		if prev == z {
			return z
		}
		fmt.Println(z)
	}
}
