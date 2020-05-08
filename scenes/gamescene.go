package scenes

import (
	"default/GitHub/puzzle/global"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// GameScene scene of game
type GameScene struct {
	bgimg   *ebiten.Image
	subImgs [global.PuzzleColumns * global.PuzzleRows]*ebiten.Image
	board   [global.PuzzleColumns][global.PuzzleRows]int
}

//Startup initialize GameScene
func (g *GameScene) Startup() {
	var err error
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	width := global.ScreenWidth / global.PuzzleColumns
	height := global.ScreenHeight / global.PuzzleRows

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; i < global.PuzzleRows; j++ {

			g.subImgs[j*global.PuzzleColumns+i] = g.bgimg.SubImage(image.Rect(i*width, j*height, i*width+width, j*height+height)).(*ebiten.Image)

		}
	}
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; i < global.PuzzleRows; j++ {

			g.board[i][j] = j*global.PuzzleColumns + i

		}
	}

}

// Update GameScene
func (g *GameScene) Update(screen *ebiten.Image) error {
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; i < global.PuzzleRows; j++ {

			screen.DrawImage(g.subImgs[g.board[i][j]], nil)

		}
	}

	return nil
}
