package conv

import (
	"fmt"
	"math"
)

type Celsius 	float64
type Fahrenheit float64
type Kelvin		float64
type Meter		float64
type Foot 		float64

//FToM converts Foots width to Meters
func FToM(f Foot) Meter {
	return Meter(math.Round(float64(f)/3.28084))
}
//MToF converts Meters width to Foots
func MToF(m Meter) Foot {
	return Foot(math.Round(float64(m) * 3.28084))
}
// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(math.Round(float64(c*9/5 + 32)))
}
// CToK conv. a Cel. temp. to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(math.Round(float64(c + 273.15)))
}
// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius(math.Round(float64(f - 32)*5/9))
}
// FToK conv. Fahren. to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(math.Round(float64(f + 459.67)* 5/9))
}
// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(math.Round(float64(k - 273.15)))
}
// KToF conv. Kelvin to Fahr.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(math.Round(float64(k * 9/5 - 459.67)))
}
func (f Foot) String()(string){
	return fmt.Sprintf("%g ft.", f)
}
func (m Meter) String()(string){
	return fmt.Sprintf("%g meters", m)
}
func (c Celsius) String()(string){
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String()(string){
	return fmt.Sprintf("%g°F", f)
}
func(k Kelvin) String()(string){
	return fmt.Sprintf("%g K", k)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC 	  Celsius = 0
	BoilingC	  Celsius = 100
	FreezingF	  Fahrenheit = 32
)

	
