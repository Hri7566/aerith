package menu

type MenuTabUpdateCallback func(dt float32)
type MenuTabDrawCallback func()

type MenuTab struct {
	Name   string
	Draw   MenuTabDrawCallback
	Update MenuTabUpdateCallback
}

var tabs []MenuTab
