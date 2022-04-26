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
		ScreenWidth:  1366,
		ScreenHeight: 768,
	}

	return &p
}

func (p *Project) Add(entities ...Entity) {
	p.entities = append(p.entities, entities...)
}

func (p *Project) Register(systems ...System) {
	for i := range systems {
		p.systems = append(p.systems, systems[i])

		v, ok := p.systems[i].(Registrable)
		if !ok {
			continue
		}

		v.Register(p)
	}
}

func (p *Project) Run() error {
	ebiten.SetWindowSize(p.ScreenWidth, p.ScreenHeight)

	if err := ebiten.RunGame(p); err != nil {
		return errors.Wrap(err, "error on run game")
	}

	return nil
}
func (p *Project) Update() error {
	for i := range p.entities {
		for j := range p.systems {
			v, ok := p.systems[j].(Updatable)
			if !ok {
				continue
			}

			err := v.Update(p.entities[i])
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
			v, ok := p.systems[j].(Drawable)
			if !ok {
				continue
			}

			v.Draw(p.entities[i], screen)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.CurrentFPS()))
}

func (p *Project) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	p.ScreenWidth = outsideWidth
	p.ScreenHeight = outsideHeight

	return outsideWidth, outsideHeight
}
