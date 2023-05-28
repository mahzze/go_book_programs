package distance

type Inches float32
type Feet float32
type Yard float32
type Mile float32

func (ml Mile) MiToMm() Milimeter  { return Milimeter(ml * 1600000) }
func (ml Mile) MiToCm() Centimeter { return Centimeter(ml * 1600000) }
func (ml Mile) MiToM() Meter       { return Meter(ml * 1600) }
func (ml Mile) MiToKm() Milimeter  { return Milimeter(ml * 1.6) }
