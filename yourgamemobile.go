package yourgamemobile

import (
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/mobile"

    "github.com/Revanthk/Balance"
)

func init() {
    // yourgame.Game must implement mobile.Game (= ebiten.Game) interface.
    // For more details, see
    // * https://pkg.go.dev/github.com/hajimehoshi/ebiten?tab=doc#Game
    mobile.SetGame(&yourgame.Game{})
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}