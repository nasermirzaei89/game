package engine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/errors"
)

type Project struct {
	objs   []Object
	actors []Actor
}

func NewProject() *Project {
	p := Project{
		objs:   make([]Object, 0),
		actors: make([]Actor, 0),
	}

	return &p
}

func (p *Project) Add(objs ...Object) {
	p.objs = append(p.objs, objs...)
}

func (p *Project) Register(actors ...Actor) {
	p.actors = append(p.actors, actors...)
}

func (p *Project) Run() error {
	if err := ebiten.RunGame(p); err != nil {
		return errors.Wrap(err, "error on run game")
	}

	return nil
}
func (p *Project) Update() error {
	for i := range p.objs {
		for j := range p.actors {
			err := p.actors[j].Update(p.objs[i])
			if err != nil {
				return errors.Wrap(err, "error on update object")
			}
		}
	}

	return nil
}

func (p *Project) Draw(screen *ebiten.Image) {
	for i := range p.objs {
		for j := range p.actors {
			p.actors[j].Draw(p.objs[i], screen)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.CurrentFPS()))
}

func (p *Project) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
