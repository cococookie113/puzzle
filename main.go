package main

import (
	"log"

	"github.com/cococookie113/puzzle/game/global"
	"github.com/cococookie113/puzzle/game/scenemanager"
	"github.com/cococookie113/puzzle/game/scenes"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "puzzle")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}