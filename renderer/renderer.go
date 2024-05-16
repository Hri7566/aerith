package renderer

import (
	"aerith/screen"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var showFPS bool = false
var EnableFPS bool = false
var BackgroundColor rl.Color = rl.DarkGray

func Init() {
	showFPS = EnableFPS
}

func Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyF) {
		showFPS = !showFPS
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		screen.Mode = "menu"
	}
}

func Draw() {
	rl.ClearBackground(BackgroundColor)

	var whiteIndex int32 = 0
	var blackIndex int32 = 0

	var x float32 = 0
	var y float32 = 0

	keyWidth := float32(screen.Width) / 52
	keyHeight := float32(screen.Height) / 5

	blackKeyWidth := keyWidth * .75
	blackKeyHeight := (keyHeight / 12) * 8

	// Draw keys
	for _, key := range keys {
		if !key.Sharp {
			x = keyWidth * float32(whiteIndex)
			y = float32(screen.Height) - keyHeight - 1

			rl.DrawRectangle(int32(x), int32(y), int32(keyWidth), int32(keyHeight), rl.White)
			rl.DrawRectangleLines(int32(x), int32(y)-1, int32(keyWidth)+1, int32(keyHeight), rl.LightGray)

			if showFPS {
				rl.DrawText(key.Name, int32(x), int32(y)+int32(keyHeight-10), 10, rl.Blue)
			}

			whiteIndex++
		}
	}

	whiteIndex = 0
	blackIndex = 0

	for _, key := range keys {
		if !key.Sharp {
			x = keyWidth * float32(whiteIndex)
			y = float32(screen.Height) - keyHeight - 1

			whiteIndex++
		} else {
			x += keyWidth - (blackKeyWidth / 2)

			rl.DrawRectangle(int32(x), int32(y)-1, int32(blackKeyWidth), int32(blackKeyHeight), rl.Black)

			if showFPS {
				rl.DrawText(key.Name, int32(x), int32(y)+int32(blackKeyHeight-10), 10, rl.Blue)
			}

			blackIndex++
		}
	}

	if showFPS {
		rl.DrawFPS(0, screen.Height-26)
	}
}
