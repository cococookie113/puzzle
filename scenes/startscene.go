package scenes

import (
	"image/color"
	"log"

	"github.com/cococookie113/puzzle/font"
	"github.com/cococookie113/puzzle/global"
	"github.com/cococookie113/puzzle/scenemanager"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) Startup() {
	var err error
	s.startImg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)

	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText,
		global.ScreenWidth/2-width/2, global.ScreenHeight/2, 2, color.Black)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
