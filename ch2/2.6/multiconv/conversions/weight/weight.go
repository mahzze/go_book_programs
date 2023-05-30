package weight

type Pound float32
type Miligram float32
type Gram float32
type Kilo float32
type Ton float32

func (lb Pound) LbToMg() Miligram { return Miligram(lb * 454000) }
func (lb Pound) LbToG() Gram      { return Gram(lb * 454) }
func (lb Pound) LbToKg() Kilo     { return Kilo(lb * 0.454) }
func (lb Pound) LbToT() Ton       { return Ton(lb * 0.000454) }

func (mg Miligram) MgToLb() Pound { return Pound(mg / 454000) }
func (mg Miligram) MgToG() Gram   { return Gram(mg / 1000) }
func (mg Miligram) MgToKg() Kilo  { return Kilo(mg / 1000000) }
func (mg Miligram) MgToT() Ton    { return Ton(mg / 10e9) }

func (g Gram) GToLb() Pound    { return Pound(g / 454) }
func (g Gram) GToMg() Miligram { return Miligram(g * 1000) }
func (g Gram) GToKg() Kilo     { return Kilo(g / 1000) }
func (g Gram) GToT() Ton       { return Ton(g / 1000000) }

func (kg Kilo) KgToLb() Pound    { return Pound(kg / 0.454) }
func (kg Kilo) KgToMg() Miligram { return Miligram(kg * 1000000) }
func (kg Kilo) KgToG() Gram      { return Gram(kg * 1000) }
func (kg Kilo) KgToT() Ton       { return Ton(kg / 1000) }

func (t Ton) TToLb() Pound    { return Pound(t / 0.000454) }
func (t Ton) TToMg() Miligram { return Miligram(t * 10e9) }
func (t Ton) TToG() Gram      { return Gram(t * 1000000) }
func (t Ton) TToKg() Kilo     { return Kilo(t * 1000) }
