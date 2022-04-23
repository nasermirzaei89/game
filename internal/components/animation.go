package components

type Animation struct {
	Frames    []*Sprite
	FrameRate float64
	FrameTime float64
}

func (a *Animation) GetAnimation() *Animation {
	return a
}

func (a *Animation) GetSprite() *Sprite {
	return a.Frames[int(a.FrameTime)%len(a.Frames)]
}
