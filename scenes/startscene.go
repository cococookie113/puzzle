package scenes

import (
	"image/color"
	"log"

	"github.com/cococookie113/puzzle/font"
	"github.com/cococookie113/puzzle/scenemanager"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/yourname/yourgame/go/pkg/mod/github.com/cococookie113/12jangi@v0.0.0-20200506051726-18b80c2abaff/game/global"
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

	font.DrawTextWithShadow(screen, "Press Any Button", global.ScreenWidth/2, global.ScreenHeight/2, 3, color.Black)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}
	return nil
}
