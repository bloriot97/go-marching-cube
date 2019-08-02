package map_pkg

import (
	"github.com/aquilax/go-perlin"
	"github.com/g3n/engine/util/application"
)

const (
	mapSizeX   = 10
	mapSizeY   = 2
	mapSizeZ   = 10
	chunkSizeX = 32
	chunkSizeY = 32
	chunkSizeZ = 32
)

const (
	alpha       = 2.
	beta        = 2.
	n           = 3
	seed  int64 = 10013
)

type Map struct {
	chunks [mapSizeX][mapSizeY][mapSizeZ]*Chunk
}


var p = perlin.NewPerlin(alpha, beta, n, seed)

func NewMap() (*Map, error){
	m := &Map{
		chunks: [mapSizeX][mapSizeY][mapSizeZ]*Chunk{},
	}

	m.genMap()

	return m, nil
}

func (m *Map) genMap(){
	for x := 0; x < mapSizeX; x ++{
		for y := 0; y < mapSizeY; y ++{
			for z := 0; z < mapSizeZ; z ++{
				m.chunks[x][y][z], _ = newChunk(x,y,z)
			}
		}
	}
}

func (m *Map) AddMesh(app *application.Application) {
	for x := 0; x < mapSizeX; x ++{
		for y := 0; y < mapSizeY; y ++{
			for z := 0; z < mapSizeZ; z ++{
				m.chunks[x][y][z].mesh.TranslateX(float32((x - mapSizeX / 2.) * chunkSizeX  ) )
				m.chunks[x][y][z].mesh.TranslateY(float32((y - mapSizeY / 2.) * chunkSizeY  ) )
				m.chunks[x][y][z].mesh.TranslateZ(float32((z - mapSizeZ / 2.) * chunkSizeZ  ) )
				app.Scene().Add(m.chunks[x][y][z].mesh)
			}
		}
	}
}