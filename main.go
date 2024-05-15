package main

import (
	"aerith/menu"
	"aerith/renderer"
	"aerith/screen"

	rl "github.com/gen2brain/raylib-go/raylib"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

// Module declarations

var showFPS bool = false

// var exitKey int32 = rl.KeyQ
var exitKey int32 = rl.KeyNull

func main() {
	defer midi.CloseDriver()

	menu.Init()
	renderer.Init()

	// Setup window
	rl.InitWindow(screen.Width, screen.Height, screen.GetTitle())
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0)

	rl.SetExitKey(exitKey)

	defer rl.CloseWindow()

	var dt float32

	for !rl.WindowShouldClose() {
		if screen.ShouldClose {
			break
		}
		// Call updates

		dt = rl.GetFrameTime()
		update(dt)

		rl.BeginDrawing()

		// Call draws
		draw()

		rl.EndDrawing()
	}
}

func update(dt float32) {
	if screen.Mode == "menu" {
		menu.Update(dt)
	} else if screen.Mode == "render" {
		renderer.Update(dt)
	}
}

func draw() {
	if screen.Mode == "menu" {
		menu.Draw()
	} else if screen.Mode == "render" {
		renderer.Draw()
	}
}
