package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/errors"
	"image"
	_ "image/png"
)

type ImageManager interface {
	LoadImage(name, filePath string) error
	UnloadImage(name string) error
	LoadSubImage(name, fromName string, bound image.Rectangle) error
	GetImage(name string) (*ebiten.Image, error)
	MustGetImage(name string) *ebiten.Image
}

type imageManager struct {
	images map[string]*ebiten.Image
}

func (em *imageManager) LoadImage(name, filePath string) error {
	img, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		return errors.Wrap(err, "error on new image from file")
	}

	if em.images == nil {
		em.images = make(map[string]*ebiten.Image)
	}

	em.images[name] = img

	return nil
}

func (em *imageManager) UnloadImage(name string) error {
	img, ok := em.images[name]
	if !ok {
		return errors.Errorf("image with name '%s' not found", name)
	}

	delete(em.images, name)

	img.Dispose()

	return nil
}

func (em *imageManager) LoadSubImage(name, fromName string, bound image.Rectangle) error {
	from, ok := em.images[fromName]
	if !ok {
		return errors.Errorf("image with name '%s' not found", fromName)
	}

	subImage := from.SubImage(bound)

	em.images[name] = ebiten.NewImageFromImage(subImage)

	return nil
}

func (em *imageManager) GetImage(name string) (*ebiten.Image, error) {
	if img, ok := em.images[name]; ok {
		return img, nil
	}

	return nil, errors.Errorf("image with name '%s' not found", name)
}
func (em *imageManager) MustGetImage(name string) *ebiten.Image {
	img, err := em.GetImage(name)
	if err != nil {
		panic(errors.Wrap(err, "error on get image"))
	}

	return img
}
