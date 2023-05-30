package distance

type Milimeter float64
type Centimeter float32
type Meter float32
type Kilometer float32

func (mm Milimeter) MmToCm() Centimeter { return Centimeter(mm / 10) }
func (mm Milimeter) MmToM() Meter       { return Meter(mm / 1000) }
func (mm Milimeter) MmToKm() Kilometer  { return Kilometer(mm / 1000000) }
func (mm Milimeter) MmToMi() Mile       { return Mile(mm / (1600000)) }
func (mm Milimeter) MmToIn() Inches     { return Inches(mm / 25.4) }
func (mm Milimeter) MmToFt() Feet       { return Feet(mm / 304.8) }
func (mm Milimeter) MmToYd() Yard       { return Yard(mm / 914.4) }

func (cm Centimeter) CmToMm() Milimeter { return Milimeter(cm * 10) }
func (cm Centimeter) CmToM() Meter      { return Meter(cm / 100) }
func (cm Centimeter) CmToKm() Kilometer { return Kilometer(cm / 100000) }
func (cm Centimeter) CmToMi() Mile      { return Mile(cm / (160000)) }
func (cm Centimeter) CmToIn() Inches    { return Inches(cm / 2.54) }
func (cm Centimeter) CmToFt() Feet      { return Feet(cm / 30.48) }
func (cm Centimeter) CmToYd() Yard      { return Yard(cm / 91.44) }

func (m Meter) MToMm() Milimeter  { return Milimeter(m * 1000) }
func (m Meter) MToCm() Centimeter { return Centimeter(m * 100) }
func (m Meter) MToKm() Kilometer  { return Kilometer(m / 1000) }
func (m Meter) MToMi() Mile       { return Mile(m / (1600)) }
func (m Meter) MToIn() Inches     { return Inches(m / 0.0254) }
func (m Meter) MToFt() Feet       { return Feet(m / 0.3048) }
func (m Meter) MToYd() Yard       { return Yard(m / 0.9144) }

func (km Kilometer) KmToMm() Milimeter  { return Milimeter(km * 1000000) }
func (km Kilometer) KmToCm() Centimeter { return Centimeter(km * 100000) }
func (km Kilometer) KmToM() Meter       { return Meter(km * 1000) }
func (km Kilometer) KmToMi() Mile       { return Mile(km / 1.6) }
func (km Kilometer) KmToIn() Inches     { return Inches(km / 0.0000254) }
func (km Kilometer) KmToFt() Feet       { return Feet(km / 0.0003048) }
func (km Kilometer) KmToYd() Yard       { return Yard(km / 0.0009144) }
