// chapter 2, section 6, exercise 2.2
// to make it a bit more challenging, the package "flag" is forbidden
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mahzze/multiconv/conversions/distance"
)

func main() {
	var value float32
	var unit string

	var helpStr = `usage:
    multiconv -u=<unit> <value>
    
<value> has to be a numeric value. The result will only show up to the second
floating point value.
The value of <unit> must be one unit of the following categories:
-- distance measures
    mm (milimeters)
    cm (centimeters)
    m  (metres)
    km (kilometers)
    in (inches)
    ft (feet)
    yd (yards)
    mi (miles)

-- heat measures
    c (celsius)
    f (fahrenheit)
    k (kelvin)

-- weight measures
    mg (miligrams)
    g  (grams)
    kg (kilograms)
    t  (tons)
    lb (pounds)

After providing the value and the unit, the program will print out the same
value converted in it's categories. e.g 'multiconv -273.15 -u=c' will return:
    -273.15ºC
    -459.67ºF
    0K`

	// checks if the user has given all necessary values
	if len(os.Args) == 1 {
		fmt.Println(helpStr)
		os.Exit(0)
	}
	for _, arg := range os.Args[1:] {
		// checks if the user requested help
		if arg == "-h" || arg == "--h" || arg == "-help" || arg == "--help" {
			fmt.Println(helpStr)
			os.Exit(0)
		}
		// checks value of -unit (or -u) and adds it to a variable,
		// if it fails, checks the value passed to be converted
		if len(strings.Split(arg, "=")[0]) < len(arg) {
			unit = strings.Split(arg, "=")[1]
		} else {
			convertedFloat, err := strconv.ParseFloat(arg, 32)
			if err != nil {
				fmt.Printf("The value must contain only numbers\n")
				os.Exit(1)
			}
			value = float32(convertedFloat)
		}
	}
	switch unit {

	case "mm":
		mm := distance.Milimeter(value)
		fmt.Printf(`input:%.2f milimeters, measure of distance
%g cm
%g m
%g km
%g mi
`, mm, mm.MmToCm(), mm.MmToM(), mm.MmToKm(), mm.MmToMi())

	case "cm":
		cm := distance.Centimeter(value)
		fmt.Printf(`input:%.2f centimeters, measure of distance
%g mm
%g m
%g km
%g mi
`, cm, cm.CmToMm(), cm.CmToM(), cm.CmToKm(), cm.CmToMi())

	case "m":
		m := distance.Meter(value)
		fmt.Printf(`input:%.2f meters, measure of distance
%g mm
%g cm
%g km
%g mi
`, m, m.MToMm(), m.MToCm(), m.MToKm(), m.MToMi())

	case "km":
		km := distance.Kilometer(value)
		fmt.Printf(`input:%.2f kilometers, measure of distance
%g mm
%g cm
%g m
%g mi
`, km, km.KmToMm(), km.KmToCm(), km.KmToM(), km.KmToMi())

	case "mi":
		mi := distance.Mile(value)
		fmt.Printf(`input:%.2f miles, measure of distance
%g mm
%g cm
%g m
%g mi
`, mi, mi.MiToMm(), mi.MiToCm(), mi.MiToM(), mi.MiToKm())

	default:
		fmt.Printf("%s is not a valid unit, to see the valid options run the program with the -h or \n--help flags, or run it with no flags and no arguments\n", unit)
	}

}
