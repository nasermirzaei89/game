package components

type Velocity struct {
	Horizontal, Vertical float64
}

func (v *Velocity) GetVelocity() *Velocity {
	return v
}
