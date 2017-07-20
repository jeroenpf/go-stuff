package hex

// GetGradient returns point in gradient
func GetGradient(startHex int, endHex int, step float64) int {

	aR := float64(startHex >> 16 & 0xFF)
	aG := float64(startHex >> 8 & 0xFF)
	aB := float64(startHex & 0xFF)

	bR := float64(endHex >> 16 & 0xFF)
	bG := float64(endHex >> 8 & 0xFF)
	bB := float64(endHex & 0xFF)

	// For each channel determine the size of a step
	sR := (bR - aR) / 255.0
	sG := (bG - aG) / 255.0
	sB := (bB - aB) / 255.0

	oR := int((aR + (sR * step)))
	oG := int((aG + (sG * step)))
	oB := int((aB + (sB * step)))

	return oR<<16 | oG<<8 | oB
}
