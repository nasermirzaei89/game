package engine

import (
	"math"
)

type Animation struct {
	FrameIndex float64
	FrameRate  float64
	Frames     []*Sprite
}

func (animation *Animation) Update() {
	frameCount := len(animation.Frames)
	if frameCount == 0 {
		return
	}

	animation.FrameIndex += animation.FrameRate / 60
	animation.FrameIndex = math.Mod(animation.FrameIndex, float64(frameCount))
}

func (animation *Animation) CurrentFrame() *Sprite {
	index := math.Floor(animation.FrameIndex)

	return animation.Frames[int(index)]
}
