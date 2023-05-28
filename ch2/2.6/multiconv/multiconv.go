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
		fmt.Printf(`input:%g milimeters, measure of distance
------ Metric Scale -------
%g cm
%g m
%g km
----- Imperial Scale ------
%g in
%g ft
%g yd
%g mi
`, mm, mm.MmToCm(), mm.MmToM(), mm.MmToKm(), mm.MmToIn(), mm.MmToFt(), mm.MmToYd(), mm.MmToMi())

	case "cm":
		cm := distance.Centimeter(value)
		fmt.Printf(`input:%g centimeters, measure of distance
------ Metric Scale -------
%g mm
%g m
%g km
----- Imperial Scale ------
%g in
%g ft
%g yd
%g mi
`, cm, cm.CmToMm(), cm.CmToM(), cm.CmToKm(), cm.CmToIn(), cm.CmToFt(), cm.CmToYd(), cm.CmToMi())

	case "m":
		m := distance.Meter(value)
		fmt.Printf(`input:%g meters, measure of distance
------ Metric Scale -------
%g mm
%g cm
%g km
----- Imperial Scale ------
%g in
%g ft
%g yd
%g mi
`, m, m.MToMm(), m.MToCm(), m.MToKm(), m.MToIn(), m.MToFt(), m.MToYd(), m.MToMi())

	case "km":
		km := distance.Kilometer(value)
		fmt.Printf(`input:%g kilometers, measure of distance
------ Metric Scale -------
%g mm
%g cm
%g m
----- Imperial Scale ------
%g in
%g ft
%g yd
%g mi
`, km, km.KmToMm(), km.KmToCm(), km.KmToM(), km.KmToIn(), km.KmToFt(), km.KmToYd(), km.KmToMi())

	case "mi":
		mi := distance.Mile(value)
		fmt.Printf(`input:%g miles, measure of distance
------ Metric Scale -------
%g mm
%g cm
%g m
%g km
----- Imperial Scale ------
%g in
%g ft
%g yd
`, mi, mi.MiToMm(), mi.MiToCm(), mi.MiToM(), mi.MiToKm(), mi.MiToIn(), mi.MiToFt(), mi.MiToYd())

	case "in":
		in := distance.Inches(value)
		fmt.Printf(`input:%g inches, measure of distance
------ Metric Scale -------
%g mm
%g cm
%g m
%g km
----- Imperial Scale ------
%g ft
%g yd
%g mi
`, in, in.InToMm(), in.InToCm(), in.InToM(), in.InToKm(), in.InToFt(), in.InToYd(), in.InToMi())

	case "ft":
		ft := distance.Feet(value)
		fmt.Printf(`input:%g feet, measure of distance
------ Metric Scale -------
%g mm
%g cm
%g m
%g km
----- Imperial Scale ------
%g in
%g yd
%g mi 
`, ft, ft.FtToMm(), ft.FtToCm(), ft.FtToM(), ft.FtToKm(), ft.FtToIn(), ft.FtToYd(), ft.FtToMi())

	case "yd":
		yd := distance.Yard(value)
		fmt.Printf(`input:%g yards, measure of distance
------ Metric Scale -------
%g mm
%g cm
%g m
%g km
----- Imperial Scale ------
%g in
%g ft
%g mi 
`, yd, yd.YdToMm(), yd.YdToCm(), yd.YdToM(), yd.YdToKm(), yd.YdToIn(), yd.YdToFt(), yd.YdToMi())

	default:
		fmt.Printf("%s is not a valid unit, to see the valid options run the program with the -h or \n--help flags, or run it with no flags and no arguments\n", unit)
	}
}
