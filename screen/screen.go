package screen

import rl "github.com/gen2brain/raylib-go/raylib"

var Mode string = "menu"
var Width int32 = 1280
var Height int32 = 720
var ShouldClose bool = false

var title string = "Aerith Renderer"

func GetTitle() string {
	return title
}

func SetTitle(t string) {
	title = t
	rl.SetWindowTitle(title)
}
