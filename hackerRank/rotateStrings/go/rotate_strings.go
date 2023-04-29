package _go

import "math"

func rotateString(s string, toLeft, toRight int) string {
	if len(s) == 1 || toRight == toLeft {
		return s
	}
	rotationAux := toRight - toLeft
	if rotationAux > 0 {
		if rotationAux > len(s) {
			rotationAux = rotationAux % len(s)
		}
		return rotateToRight(s, rotationAux)
	} else {
		rotationAuxAbs := int(math.Abs(float64(rotationAux)))
		if rotationAux > len(s) {
			rotationAux = rotationAux % len(s)
		}
		return rotateToLeft(s, rotationAuxAbs)
	}
}

func rotateToLeft(s string, aux int) string {
	return s[0:len(s)-aux] + s[:len(s)-aux]
}

func rotateToRight(s string, aux int) string {
	return s[len(s)-aux:] + s[0:len(s)-aux]
}
