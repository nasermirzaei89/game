package game

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/game/internal/game2/engine"
	"github.com/pkg/errors"
)

func Run() error {
	g := engine.NewGame("Game 2")

	err := g.LoadImage("bat-brain", "./assets/bat-brain.png")
	if err != nil {
		return errors.Wrap(err, "error on load image")
	}

	err = g.LoadSubImage("bat-brain-0", "bat-brain", image.Rect(0, 0, 40, 40))
	if err != nil {
		return errors.Wrap(err, "error on load sub image")
	}

	err = g.LoadSubImage("bat-brain-1", "bat-brain", image.Rect(40, 0, 80, 40))
	if err != nil {
		return errors.Wrap(err, "error on load sub image")
	}

	err = g.LoadSubImage("bat-brain-2", "bat-brain", image.Rect(80, 0, 120, 40))
	if err != nil {
		return errors.Wrap(err, "error on load sub image")
	}

	err = g.LoadImage("orbinaut", "./assets/orbinaut.png")
	if err != nil {
		return errors.Wrap(err, "error on load image")
	}

	err = g.LoadSubImage("orbinaut-0", "orbinaut", image.Rect(0, 0, 48, 48))
	if err != nil {
		return errors.Wrap(err, "error on load sub image")
	}

	err = g.LoadSubImage("orbinaut-1", "orbinaut", image.Rect(0, 48, 48, 96))
	if err != nil {
		return errors.Wrap(err, "error on load sub image")
	}

	err = g.LoadSubImage("orbinaut-2", "orbinaut", image.Rect(0, 96, 48, 144))
	if err != nil {
		return errors.Wrap(err, "error on load sub image")
	}

	for i := 0; i < 64; i++ {
		g.AddObject(NewOrbinaut(g))
	}

	g.AddObject(NewDrawer())

	g.AddObject(NewBatBrain(g))

	g.OnStart(func(ev *engine.EventGameStart) {
		log.Println(ev.Name())
	})

	g.OnEnd(func(ev *engine.EventGameEnd) {
		log.Println(ev.Name())
	})

	g.OnDraw(func(ev *engine.EventDraw) {
		ebitenutil.DebugPrint(ev.Screen, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
	})

	if err := g.Run(); err != nil {
		return errors.Wrap(err, "error on run game")
	}

	return nil
}
