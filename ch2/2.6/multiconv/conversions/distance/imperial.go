package distance

type Inches float32
type Feet float32
type Yard float32
type Mile float32

func (in Inches) InToMm() Milimeter  { return Milimeter(in * 25.4) }
func (in Inches) InToCm() Centimeter { return Centimeter(in * 2.54) }
func (in Inches) InToM() Meter       { return Meter(in * 0.0254) }
func (in Inches) InToKm() Kilometer  { return Kilometer(in * 0.0000254) }
func (in Inches) InToMi() Mile       { return Mile(in / 63360) }
func (in Inches) InToFt() Feet       { return Feet(in / 12) }
func (in Inches) InToYd() Yard       { return Yard(in / 36) }

func (ft Feet) FtToMm() Milimeter  { return Milimeter(ft * 304.8) }
func (ft Feet) FtToCm() Centimeter { return Centimeter(ft * 30.48) }
func (ft Feet) FtToM() Meter       { return Meter(ft * 0.3048) }
func (ft Feet) FtToKm() Kilometer  { return Kilometer(ft * 0.0003048) }
func (ft Feet) FtToMi() Mile       { return Mile(ft / 5280) }
func (ft Feet) FtToIn() Inches     { return Inches(ft * 12) }
func (ft Feet) FtToYd() Yard       { return Yard(ft / 3) }

func (yd Yard) YdToMm() Milimeter  { return Milimeter(yd * 914.4) }
func (yd Yard) YdToCm() Centimeter { return Centimeter(yd * 91.44) }
func (yd Yard) YdToM() Meter       { return Meter(yd * 0.9144) }
func (yd Yard) YdToKm() Kilometer  { return Kilometer(yd * 0.0009144) }
func (yd Yard) YdToMi() Mile       { return Mile(yd / 1760) }
func (yd Yard) YdToIn() Inches     { return Inches(yd * 36) }
func (yd Yard) YdToFt() Feet       { return Feet(yd * 3) }

func (ml Mile) MiToMm() Milimeter  { return Milimeter(ml * 1600000) }
func (ml Mile) MiToCm() Centimeter { return Centimeter(ml * 1600000) }
func (ml Mile) MiToM() Meter       { return Meter(ml * 1600) }
func (ml Mile) MiToKm() Milimeter  { return Milimeter(ml * 1.6) }
func (ml Mile) MiToIn() Inches     { return Inches(ml * 63360) }
func (ml Mile) MiToFt() Feet       { return Feet(5280) }
func (ml Mile) MiToYd() Yard       { return Yard(1760) }
