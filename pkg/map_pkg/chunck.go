package map_pkg

import (
	"fmt"
	"github.com/bloriot97/go-marching-cube/pkg/math_helper"
	"github.com/bloriot97/go-marching-cube/pkg/mesh"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Chunk struct {
	blocks            [chunkSizeX+1][chunkSizeY+1][chunkSizeZ+1]bool
	posX, posY, posZ  int
	mesh              *graphic.Mesh
}

func newChunk(posX, posY, posZ int) (*Chunk, error){
	chunk := &Chunk{
		blocks: [chunkSizeX+1][chunkSizeY+1][chunkSizeZ+1]bool{},
		posX: posX,
		posY: posY,
		posZ: posZ,

	}
	if err := chunk.genBlocks(); err != nil{
		return nil, err
	}
	if err := chunk.genMesh(); err != nil{
		return nil, err
	}

	return chunk, nil
}

func (c *Chunk) genBlocks() error{
	fmt.Sprintf("qsd")
	for x := 0; x < chunkSizeX+1; x++ {
		for y := 0; y < chunkSizeY+1; y++ {
			for z := 0; z < chunkSizeZ+1; z++ {
				pos := math_helper.Vector64{
					X: float64(x+c.posX*chunkSizeX) / 50.,
					Y: float64(y+c.posY*chunkSizeY) / 25.,
					Z: float64(z+c.posZ*chunkSizeZ) / 50.,
				}
				noiseRes := p.Noise3D(
					pos.X,
					pos.Y ,
					pos.Z)
				c.blocks[x][y][z] =  noiseRes < 0.5 - pos.Y/float64(chunkSizeY * mapSizeY / 25.) * 1
				fmt.Sprintf("qsd %s", c.blocks[x][y][z])
			}
		}
	}
	return nil
}

func (c *Chunk) genMesh() error{
	geom := geometry.NewGeometry()
	buffer := math32.NewArrayF32(0, 0)
	normalBuffer := math32.NewArrayF32(0, 0)
	// Create a blue torus and add it to the scene
	for x := 0.; x < chunkSizeX ; x++ {
		for y := 0.; y < chunkSizeY ; y++ {
			for z := 0.; z < chunkSizeZ ; z++ {
				cubeType := 0
				if c.blocks[int(x)][int(y)][int(z)] {
					cubeType += 1
				}
				if c.blocks[int(x+1)][int(y)][int(z)]  {
					cubeType += 2
				}
				if c.blocks[int(x+1)][int(y)][int(z+1)] {
					cubeType += 4
				}
				if c.blocks[int(x)][int(y)][int(z+1)] {
					cubeType += 8
				}
				if c.blocks[int(x)][int(y+1)][int(z)]  {
					cubeType += 16
				}
				if c.blocks[int(x+1)][int(y+1)][int(z)] {
					cubeType += 32
				}
				if c.blocks[int(x+1)][int(y+1)][int(z+1)] {
					cubeType += 64
				}
				if c.blocks[int(x)][int(y+1)][int(z+1)]{
					cubeType += 128
				}
				if cubeType != 0{
					x_pos := float32(x - chunkSizeX/2 + 0.5)
					y_pos := float32(y - chunkSizeY/2 + 0.5)
					z_pos := float32(z - chunkSizeZ/2 + 0.5)
					mesh.AddMesh(byte(cubeType), x_pos, y_pos, z_pos, &buffer, &normalBuffer)
				}

			}
		}
	}
	geom.AddVBO(gls.NewVBO(buffer).
		AddAttrib(gls.VertexPosition),
	)
	geom.AddVBO(gls.NewVBO(normalBuffer).
		AddAttrib(gls.VertexNormal),
	)
	mat := material.NewPhong(math32.NewColor("Brown"))
	mesh := graphic.NewMesh(geom, mat)

	c.mesh = mesh

	return nil
}