package multiconv

type Milimeter float64
type Centimeter float32
type Meter float32
type Kilometer float32
type Inches float32
type Feet float32
type Yard float32
type Mile float32

func (mm Milimeter) MmToCm() Centimeter { return Centimeter(mm / 10) }
func (mm Milimeter) MmToM() Meter       { return Meter(mm / 1000) }
func (mm Milimeter) MmToKm() Kilometer  { return Kilometer(mm / 1000000) }
func (mm Milimeter) MmToMl() Mile       { return Mile(mm / (1600000)) }

func (cm Centimeter) CmToMm() Milimeter { return Milimeter(cm * 10) }
func (cm Centimeter) CmToM() Meter      { return Meter(cm / 100) }
func (cm Centimeter) CmToKm() Kilometer { return Kilometer(cm / 100000) }
func (cm Centimeter) CmToMl() Mile      { return Mile(cm / (160000)) }

func (m Meter) MToMM() Milimeter  { return Milimeter(m * 1000) }
func (m Meter) MToCm() Centimeter { return Centimeter(m * 100) }
func (m Meter) MToKm() Kilometer  { return Kilometer(m / 1000) }
func (m Meter) MToMl() Mile       { return Mile(m / (1600)) }

func (km Kilometer) KmToMM() Milimeter  { return Milimeter(km * 1000000) }
func (km Kilometer) KmToCm() Centimeter { return Centimeter(km * 100000) }
func (km Kilometer) KmToM() Meter       { return Meter(km * 1000) }
func (km Kilometer) KmToMl() Mile       { return Mile(km / 1.6) }

func (ml Mile) MlToMm() Milimeter  { return Milimeter(ml * 1600000) }
func (ml Mile) MlToCm() Centimeter { return Centimeter(ml * 1600000) }
func (ml Mile) MlToM() Meter       { return Meter(ml * 1600) }
func (ml Mile) MlToKm() Milimeter  { return Milimeter(ml * 1.6) }
