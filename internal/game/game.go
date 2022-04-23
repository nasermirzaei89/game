package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
	"github.com/nasermirzaei89/td/internal/systems"
	"github.com/pkg/errors"
)

func Run() error {
	project1 := engine.NewProject()

	project1.Register(
		new(systems.Animator),
		new(systems.Drawer),
		new(systems.MovementController),
		new(systems.Transporter),
		new(systems.Wrapper),
	)

	batBrainImg, _, err := ebitenutil.NewImageFromFile("./assets/bat-brain.png")
	if err != nil {
		return errors.Wrap(err, "error on new image from file")
	}

	orbinautImg, _, err := ebitenutil.NewImageFromFile("./assets/orbinaut.png")
	if err != nil {
		return errors.Wrap(err, "error on new image from file")
	}

	batBrainAnimation := components.Animation{
		Frames: []*components.Sprite{
			{
				Source:  batBrainImg,
				X0:      0,
				X1:      40,
				Y0:      0,
				Y1:      40,
				OffsetX: 20,
				OffsetY: 20,
			},
			{
				Source:  batBrainImg,
				X0:      40,
				X1:      80,
				Y0:      0,
				Y1:      40,
				OffsetX: 20,
				OffsetY: 20,
			},
			{
				Source:  batBrainImg,
				X0:      80,
				X1:      120,
				Y0:      0,
				Y1:      40,
				OffsetX: 20,
				OffsetY: 20,
			},
			{
				Source:  batBrainImg,
				X0:      40,
				X1:      80,
				Y0:      0,
				Y1:      40,
				OffsetX: 20,
				OffsetY: 20,
			},
		},
		FrameRate: 16,
		FrameTime: 0,
	}

	orbinautAnimation := components.Animation{
		Frames: []*components.Sprite{
			{
				Source:  orbinautImg,
				X0:      0,
				X1:      48,
				Y0:      0,
				Y1:      48,
				OffsetX: 24,
				OffsetY: 24,
			},
			{
				Source:  orbinautImg,
				X0:      0,
				X1:      48,
				Y0:      48,
				Y1:      96,
				OffsetX: 24,
				OffsetY: 24,
			},
			{
				Source:  orbinautImg,
				X0:      0,
				X1:      48,
				Y0:      96,
				Y1:      144,
				OffsetX: 24,
				OffsetY: 24,
			},
		},
		FrameRate: 16,
		FrameTime: 0,
	}

	for i := 0; i < 10; i++ {
		project1.Add(
			&Orbinaut{
				Position: components.Position{
					X: float64(rand.Intn(project1.ScreenWidth)),
					Y: float64(rand.Intn(project1.ScreenHeight)),
				},
				Velocity: components.Velocity{
					Horizontal: float64(rand.Intn(10)) - 5,
					Vertical:   float64(rand.Intn(10)) - 5,
				},
				Animation: orbinautAnimation,
				Wrapable: components.Wrapable{
					Horizontal: true,
					Vertical:   true,
				},
			},
		)
	}

	project1.Add(
		&BatBrain{
			Position: components.Position{
				X: float64(rand.Intn(project1.ScreenWidth)),
				Y: float64(rand.Intn(project1.ScreenHeight)),
			},
			MovementControl: components.MovementControl{
				Up:    ebiten.KeyUp,
				Down:  ebiten.KeyDown,
				Left:  ebiten.KeyLeft,
				Right: ebiten.KeyRight,
			},
			Animation: batBrainAnimation,
			Wrapable: components.Wrapable{
				Horizontal: true,
				Vertical:   true,
			},
		},
	)

	if err := project1.Run(); err != nil {
		return errors.Wrap(err, "error on run project1")
	}

	return nil
}
