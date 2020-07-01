package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"fmt"
    "os"
    "path/filepath"
)

var img *ebiten.Image

func init() {
	var files []string

    root := "assets/characters/"
    filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    	if path[len(path)-3:len(path)] == "jpg"{
        	files = append(files, path)
    	}
        return nil
    })
    for _, file := range files {
        fmt.Println(file)
    }

    var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/characters/brad_pitt.jpg", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}


}

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)

	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}