package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"fmt"
    "os"
    "path/filepath"
)

var img *ebiten.Image
var files []string
var frame int
var counter int
const DELAY int = 7
const screenWidth = 800
const screenHeight = 700

func init() {
	

    root := "assets/characters/"
    filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    	if path[len(path)-3:len(path)] == "jpg"{
        	files = append(files, path)
    	}
        return nil
    })

    var err error
			img, _, err = ebitenutil.NewImageFromFile(files[0], ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}

    counter = 0
    frame = 8


}

// // In returns true if (x, y) is in the sprite, and false otherwise.
// func In(x, y int, s *ebiten.Image) bool {
// 	// Check the actual color (alpha) value at the specified position
// 	// so that the result of In becomes natural to users.
// 	//
// 	// Note that this is not a good manner to use At for logic
// 	// since color from At might include some errors on some machines.
// 	// As this is not so important logic, it's ok to use it so far.

// 	w, h := s.Size()

// 	if (s.x + w) > x && x > s.x && (s.y + h) > y && y > s.y {
// 		fmt.Printf("true")
// 		return true
// 	}

// 	return false

// 	//return s.image.At(x-s.x, y-s.y).(color.RGBA).A > 0
// }

type Game struct{
	startX int
	startY int
	dX int
	dY int
	lastLocX int
	lastLocY int
}

func (g *Game) Update(screen *ebiten.Image) error {


	frame++
	if(frame > DELAY) {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			counter--
			if(counter < 0){
				counter = len(files) - 1
			}

			var err error
			img, _, err = ebitenutil.NewImageFromFile(files[counter], ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			frame = 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			counter++
			if(counter == len(files)){
				counter = 0
			}
			var err error
			img, _, err = ebitenutil.NewImageFromFile(files[counter], ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			frame = 0
		}		
	}

	if g.lastLocX == 0 {

		w, h := img.Size()

	    g.lastLocX = screenWidth/2 - (w/4)
	    g.lastLocY = screenHeight/2 - (h/4)
	}

	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.startX = mx
		g.startY = my

	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {

		g.dX = mx - g.startX
		g.dY = my - g.startY
		g.lastLocX = g.lastLocX + g.dX
		g.lastLocY = g.lastLocY + g.dY
		g.startX = 0
		g.startY = 0
	}


	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		var dX = float64(mx - g.startX)
		var dY = float64(my - g.startY)
		op.GeoM.Translate(float64(g.lastLocX) + dX, float64(g.lastLocY) + dY)
	} else{
		op.GeoM.Translate(float64(g.lastLocX), float64(g.lastLocY))
	}

	screen.DrawImage(img, op)

	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("(%d, %d)", mx, my)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{startX: 0, startY: 0, dX: 0, dY: 0,}); err != nil {
		log.Fatal(err)
	}
}