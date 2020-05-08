package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/cococookie113/puzzle/global"
	"github.com/cococookie113/puzzle/scenemanager"
	"github.com/cococookie113/puzzle/scenes"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "puzzle")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
