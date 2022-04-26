package components

type Wrapable struct {
	Horizontal bool
	Vertical   bool
}

func (w *Wrapable) GetWrapable() *Wrapable {
	return w
}
