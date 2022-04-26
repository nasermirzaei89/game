package game

import (
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
	"github.com/nasermirzaei89/game/internal/game1/engine"
	systems2 "github.com/nasermirzaei89/game/internal/game1/systems"
	"image"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/errors"
)

func Run() error {
	project1 := engine.NewProject()

	project1.Register(
		new(systems2.Animator),
		new(systems2.Drawer),
		new(systems2.MovementController),
		new(systems2.Transporter),
		new(systems2.Wrapper),
	)

	batBrainImg, _, err := ebitenutil.NewImageFromFile("./assets/bat-brain.png")
	if err != nil {
		return errors.Wrap(err, "error on new image from file")
	}

	orbinautImg, _, err := ebitenutil.NewImageFromFile("./assets/orbinaut.png")
	if err != nil {
		return errors.Wrap(err, "error on new image from file")
	}

	batBrainAnimation := components2.Animation{
		Frames: []*components2.Sprite{
			{
				Image:   batBrainImg.SubImage(image.Rect(0, 0, 40, 40)).(*ebiten.Image),
				OffsetX: 20,
				OffsetY: 20,
			},
			{
				Image:   batBrainImg.SubImage(image.Rect(40, 0, 80, 40)).(*ebiten.Image),
				OffsetX: 20,
				OffsetY: 20,
			},
			{
				Image:   batBrainImg.SubImage(image.Rect(80, 0, 120, 40)).(*ebiten.Image),
				OffsetX: 20,
				OffsetY: 20,
			},
			{
				Image:   batBrainImg.SubImage(image.Rect(40, 0, 80, 40)).(*ebiten.Image),
				OffsetX: 20,
				OffsetY: 20,
			},
		},
		FrameRate: 16,
		FrameTime: 0,
	}

	orbinautAnimation := components2.Animation{
		Frames: []*components2.Sprite{
			{
				Image:   orbinautImg.SubImage(image.Rect(0, 0, 48, 48)).(*ebiten.Image),
				OffsetX: 24,
				OffsetY: 24,
			},
			{
				Image:   orbinautImg.SubImage(image.Rect(0, 48, 48, 96)).(*ebiten.Image),
				OffsetX: 24,
				OffsetY: 24,
			},
			{
				Image:   orbinautImg.SubImage(image.Rect(0, 96, 48, 144)).(*ebiten.Image),
				OffsetX: 24,
				OffsetY: 24,
			},
		},
		FrameRate: 16,
		FrameTime: 0,
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 32; i++ {
		project1.Add(
			&Orbinaut{
				Position: components2.Position{
					X: float64(rnd.Intn(project1.ScreenWidth)),
					Y: float64(rnd.Intn(project1.ScreenHeight)),
				},
				Velocity: components2.Velocity{
					Horizontal: float64(rnd.Intn(10)) - 5,
					Vertical:   float64(rnd.Intn(10)) - 5,
				},
				Animation: orbinautAnimation,
				Wrapable: components2.Wrapable{
					Horizontal: true,
					Vertical:   true,
				},
			},
		)
	}

	project1.Add(
		&BatBrain{
			Position: components2.Position{
				X: float64(rnd.Intn(project1.ScreenWidth)),
				Y: float64(rnd.Intn(project1.ScreenHeight)),
			},
			MovementControl: components2.MovementControl{
				Up:    ebiten.KeyUp,
				Down:  ebiten.KeyDown,
				Left:  ebiten.KeyLeft,
				Right: ebiten.KeyRight,
			},
			Animation: batBrainAnimation,
			Wrapable: components2.Wrapable{
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
