package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/errors"
)

type Project struct {
	entities                  []Entity
	systems                   []System
	ScreenWidth, ScreenHeight int
}

func NewProject() *Project {
	p := Project{
		entities:     make([]Entity, 0),
		systems:      make([]System, 0),
		ScreenWidth:  800,
		ScreenHeight: 600,
	}

	return &p
}

func (p *Project) Add(entities ...Entity) {
	p.entities = append(p.entities, entities...)
}

func (p *Project) Register(systems ...System) {
	for i := range systems {
		systems[i].Register(p)
		p.systems = append(p.systems, systems[i])
	}
}

func (p *Project) Run() error {
	if err := ebiten.RunGame(p); err != nil {
		return errors.Wrap(err, "error on run game")
	}

	return nil
}
func (p *Project) Update() error {
	for i := range p.entities {
		for j := range p.systems {
			u, ok := p.systems[j].(Updatable)
			if !ok {
				continue
			}

			err := u.Update(p.entities[i])
			if err != nil {
				return errors.Wrap(err, "error on update object")
			}
		}
	}

	return nil
}

func (p *Project) Draw(screen *ebiten.Image) {
	for i := range p.entities {
		for j := range p.systems {
			d, ok := p.systems[j].(Drawable)
			if !ok {
				continue
			}

			d.Draw(p.entities[i], screen)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.CurrentFPS()))
}

func (p *Project) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	p.ScreenWidth = outsideWidth
	p.ScreenHeight = outsideHeight

	return outsideWidth, outsideHeight
}
