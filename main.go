package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const FPS = 60

var window *sdl.Window
var renderer *sdl.Renderer
var running = true

func main() {
	// Initialize Video, Keyboard and Mouse Subsystems
	sdl.Init(sdl.INIT_VIDEO)

	// Set hardware accelaration, render is synchronized with refresh rate of the
	// screen, and window to fullscreen
	flags := uint32(sdl.RENDERER_ACCELERATED |
		sdl.RENDERER_PRESENTVSYNC |
		sdl.WINDOW_FULLSCREEN_DESKTOP)

	var err error

	// Create a fullscreen window without a definied position or size
	window, err = sdl.CreateWindow("Parallels", 0, 0, 0, 0, flags)

	// If a window can't be created, we crash
	if err != nil {
		log.Fatal(err)
	}

	// Create a renderer for the window. -1 is set to give us the first video
	// driver available. No flags (0) for the render is passed.
	renderer, err = sdl.CreateRenderer(window, -1, 0)

	// If a rendered can't be created, we crash
	if err != nil {
		log.Fatal(err)
	}

	delay := uint32(1000 / FPS)

	for running {
		start := sdl.GetTicks()

		processInput()
		draw()

		last := sdl.GetTicks() - start

		if last < delay {
			sdl.Delay(delay - last)
		}
	}

}

func processInput() {
	event := sdl.PollEvent()
	if event == nil {
		return
	}

	switch event := event.(type) {
	case *sdl.QuitEvent:
		quit()
	case *sdl.MouseButtonEvent:
		if event.Type == sdl.MOUSEBUTTONDOWN {
		}
	}

	return
}

func draw() {
	// Set renderer to white color (RGBA)
	renderer.SetDrawColor(0, 0, 0, 255)

	// Clear renderer to draw color
	renderer.Clear()

	// Display render at the window
	renderer.Present()
}

func quit() {
	// Destroy global renderer
	renderer.Destroy()

	// Close global window
	window.Destroy()

	// Close subsystems
	sdl.Quit()

	running = false
}
