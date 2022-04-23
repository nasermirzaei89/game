package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/actors"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
	"github.com/pkg/errors"
	"math/rand"
)

func Run() error {
	project1 := engine.NewProject()

	project1.Register(new(actors.Draw), new(actors.Transport), new(actors.Wrap), new(actors.MovementControl))

	w, h := ebiten.WindowSize()

	for i := 0; i < 10; i++ {
		project1.Add(
			&Enemy{
				Position: components.Position{
					X: float64(rand.Intn(w)),
					Y: float64(rand.Intn(h)),
				},
				Velocity: components.Velocity{
					Horizontal: float64(rand.Intn(10)) - 5,
					Vertical:   float64(rand.Intn(10)) - 5,
				},
				Character: 'x',
				Wrapable:  components.Wrapable{},
			},
		)
	}

	project1.Add(
		&Player{
			Position: components.Position{
				X: float64(rand.Intn(w)),
				Y: float64(rand.Intn(h)),
			},
			MovementControl: components.MovementControl{
				Up:    ebiten.KeyUp,
				Down:  ebiten.KeyDown,
				Left:  ebiten.KeyLeft,
				Right: ebiten.KeyRight,
			},
			Character: 'o',
			Wrapable:  components.Wrapable{},
		},
	)

	if err := project1.Run(); err != nil {
		return errors.Wrap(err, "error on run project1")
	}

	return nil
}
