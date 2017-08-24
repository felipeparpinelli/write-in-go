package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Could not initialize ttf: %v", err)
	}
	defer ttf.Quit()

	window, render, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer window.Destroy()

	if err := drawTitle(render); err != nil {
		return fmt.Errorf("Could not draw title: %v", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func drawTitle(render *sdl.Renderer) error {
	render.Clear()
	font, err := ttf.OpenFont("res/fonts/flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("Could not load font: %v", err)
	}
	defer font.Close()

	color := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	surface, err := font.RenderUTF8_Solid("Flappy Gopher", color)
	if err != nil {
		return fmt.Errorf("Could not render title: %v", err)
	}
	defer surface.Free()

	texture, err := render.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("Could not create texture: %v", err)
	}
	defer texture.Destroy()

	if err := render.Copy(texture, nil, nil); err != nil {
		return fmt.Errorf("Could not copy texture: %v", err)
	}
	render.Present()

	return nil
}
