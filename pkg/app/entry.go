package main

import (
	"github.com/bloriot97/go-marching-cube/pkg/map_pkg"

	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

func main() {

	app, _ := application.Create(application.Options{
		Title:  "Hello G3N",
		Width:  800,
		Height: 600,
	})

	m, _ := map_pkg.NewMap()
	m.AddMesh(app)


	// Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 3000.0)
	pointLight.SetPosition(100, 50, 100)
	app.Scene().Add(pointLight)

	pointLight2 := light.NewPoint(&math32.Color{1, 1, 1}, 3000.0)
	pointLight2.SetPosition(-100, 50, -100)
	app.Scene().Add(pointLight2)

	pointLight3 := light.NewPoint(&math32.Color{1, 1, 1}, 3000.0)
	pointLight3.SetPosition(0, 50, 0)
	app.Scene().Add(pointLight3)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(0.5)
	app.Scene().Add(axis)

	app.CameraPersp().SetPosition(0, 0, 10)
	app.Run()
}
