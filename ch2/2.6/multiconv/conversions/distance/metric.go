package distance

type Milimeter float64
type Centimeter float32
type Meter float32
type Kilometer float32

func (mm Milimeter) MmToCm() Centimeter { return Centimeter(mm / 10) }
func (mm Milimeter) MmToM() Meter       { return Meter(mm / 1000) }
func (mm Milimeter) MmToKm() Kilometer  { return Kilometer(mm / 1000000) }
func (mm Milimeter) MmToMi() Mile       { return Mile(mm / (1600000)) }

func (cm Centimeter) CmToMm() Milimeter { return Milimeter(cm * 10) }
func (cm Centimeter) CmToM() Meter      { return Meter(cm / 100) }
func (cm Centimeter) CmToKm() Kilometer { return Kilometer(cm / 100000) }
func (cm Centimeter) CmToMi() Mile      { return Mile(cm / (160000)) }

func (m Meter) MToMm() Milimeter  { return Milimeter(m * 1000) }
func (m Meter) MToCm() Centimeter { return Centimeter(m * 100) }
func (m Meter) MToKm() Kilometer  { return Kilometer(m / 1000) }
func (m Meter) MToMi() Mile       { return Mile(m / (1600)) }

func (km Kilometer) KmToMm() Milimeter  { return Milimeter(km * 1000000) }
func (km Kilometer) KmToCm() Centimeter { return Centimeter(km * 100000) }
func (km Kilometer) KmToM() Meter       { return Meter(km * 1000) }
func (km Kilometer) KmToMi() Mile       { return Mile(km / 1.6) }
