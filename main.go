package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ()

var (
	maxBuildings int = 100
)

type Cell struct {
	up, right, down, left     bool
	visited                   bool
	object                    rl.Rectangle
	new_visisted              bool
	nup, nright, ndown, nleft bool
}

type Land struct {
	length, height int
	cells          []Cell
	visited_cells  int
	all_path       []int
}

const (
	up = iota
	right
	down
	left
)

func main() {
	sortColor := []rl.Color{
		rl.Yellow,
		rl.Lime,
		rl.Green,
		rl.SkyBlue,
		rl.Blue,
		rl.DarkBlue,
		rl.DarkGreen,
		rl.Red,
		rl.Orange,
		rl.Pink,
		rl.Purple,
		rl.Violet,
	}
	random_pos := int(rl.GetRandomValue(0, int32(len(sortColor)-1)))
	aliveColor := sortColor[random_pos]
	deadColor := rl.White
	particleColor := rl.Black
	screenWidth := int32(1920)
	screenHeight := int32(1080)

	land := Land{}
	land.length = 40
	land.height = 20
	land.visited_cells = 0
	land.all_path = make([]int, 0)
	land.cells = make([]Cell, land.height*land.length)
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 2d camera")
	for i := 0; i < land.height; i++ {
		for j := 0; j < land.length; j++ {
			land.cells[j+i*land.length] = Cell{
				up:           true,
				right:        true,
				down:         true,
				left:         true,
				visited:      false,
				new_visisted: false,
				nup:          true,
				nright:       true,
				ndown:        true,
				nleft:        true,
			}

			land.cells[j+i*land.length].object = rl.NewRectangle(float32(j*int(screenWidth)/land.length),
				float32(i*int(screenHeight)/land.height),
				float32(int(screenWidth)/land.length),
				float32(screenHeight/int32(land.height)))
		}
	}
	land.Generate()

	rl.SetTargetFPS(120)
	size := len(land.all_path)
	pos := 0

	init_x := 0
	init_y := 0
	counter := 0
	for !rl.WindowShouldClose() {
		if counter%1000 == 0 {

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			if pos < size {

				land.cells[init_x+init_y*land.length].new_visisted = true

				switch land.all_path[pos] {

				case up:
					init_y -= 1
					break

				case right:
					init_x += 1
					break

				case down:
					init_y += 1
					break

				case left:
					init_x -= 1
					break

				}
				land.cells[init_x+init_y*land.length].Draw(particleColor)
				for i := range land.cells {
					if i == init_x+init_y*land.length {
						continue
					}
					if land.cells[i].new_visisted {
						land.cells[i].DrawOld(aliveColor)
					} else {

						land.cells[i].DrawDead(deadColor)
					}

				}
				land.Path()
				pos++

			} else {
				land = Land{}
				land.length = 40
				land.height = 20
				land.visited_cells = 0
				land.all_path = make([]int, 0)
				land.cells = make([]Cell, land.height*land.length)
				for i := 0; i < land.height; i++ {
					for j := 0; j < land.length; j++ {
						land.cells[j+i*land.length] = Cell{
							up:           true,
							right:        true,
							down:         true,
							left:         true,
							visited:      false,
							new_visisted: false,
							nup:          true,
							nright:       true,
							ndown:        true,
							nleft:        true,
						}
						land.cells[j+i*land.length].object =
							rl.NewRectangle(
								float32(j*int(screenWidth)/land.length),
								float32(i*int(screenHeight)/land.height),
								float32(int(screenWidth)/land.length),
								float32(screenHeight/int32(land.height)))

					}

				}
				pos = 0

				land.Generate()
				random_pos = int(rl.GetRandomValue(0, int32(len(sortColor)-1)))
				aliveColor = sortColor[random_pos]
				size = len(land.all_path)
				init_x = 0
				init_y = 0

			}

			rl.EndMode2D()

			rl.EndDrawing()
			counter = 0
		}

		counter++
	}

	rl.CloseWindow()
}
func (l *Land) Path() {
	for i := range l.cells {
		if l.cells[i].up {
			rl.DrawLine(
				int32(l.cells[i].object.X),
				int32(l.cells[i].object.Y),
				int32(l.cells[i].object.X+l.cells[i].object.Width),
				int32(l.cells[i].object.Y),
				rl.Black)
		}
		if l.cells[i].right {
			rl.DrawLine(
				int32(l.cells[i].object.X+l.cells[i].object.Width),
				int32(l.cells[i].object.Y),
				int32(l.cells[i].object.X+l.cells[i].object.Width),
				int32(l.cells[i].object.Y+l.cells[i].object.Height),
				rl.Black)

		}
		if l.cells[i].down {
			rl.DrawLine(
				int32(l.cells[i].object.X),
				int32(l.cells[i].object.Y+l.cells[i].object.Height),
				int32(l.cells[i].object.X+l.cells[i].object.Width),
				int32(l.cells[i].object.Y+l.cells[i].object.Height),
				rl.Black)

		}
		if l.cells[i].right {
			rl.DrawLine(
				int32(l.cells[i].object.X+l.cells[i].object.Width),
				int32(l.cells[i].object.Y),
				int32(l.cells[i].object.X+l.cells[i].object.Width),
				int32(l.cells[i].object.Y+l.cells[i].object.Height),
				rl.Black)

		}

	}

}

func (c *Cell) DrawDead(color rl.Color) {

	rl.DrawRectangleRec(c.object, color)
}
func (c *Cell) Draw(color rl.Color) {
	rl.DrawRectangleRec(c.object, color)
}
func (c *Cell) DrawOld(color rl.Color) {
	rl.DrawRectangleRec(c.object, color)
}

func (l *Land) Generate() {
	inity, initx := 0, 0
	path := make([]int, 0)
	l.cells[inity*l.length+initx].visited = true
	i, j := inity, initx
	l.visited_cells++

	for l.visited_cells < l.height*l.length {
		if l.have_unvisited_neighbour(i, j) {
			unvisited := l.unvisited_neighbour(i, j)
			pos := rand.Int() % len(unvisited)
			switch unvisited[pos] {
			case up:
				l.cells[(i-1)*l.length+j].down = false
				l.cells[(i-1)*l.length+j].visited = true
				l.cells[(i)*l.length+j].up = false
				i -= 1
				l.visited_cells++
				path = append(path, up)
				l.all_path = append(l.all_path, up)
				break

			case right:
				l.cells[(i)*l.length+j+1].left = false
				l.cells[(i)*l.length+j+1].visited = true
				l.cells[(i)*l.length+j].right = false
				j += 1
				l.visited_cells++
				path = append(path, right)
				l.all_path = append(l.all_path, right)
				break

			case down:
				l.cells[(i+1)*l.length+j].up = false
				l.cells[(i+1)*l.length+j].visited = true
				l.cells[(i)*l.length+j].down = false
				i += 1
				l.visited_cells++
				path = append(path, down)
				l.all_path = append(l.all_path, down)
				break

			case left:
				l.cells[(i)*l.length+j-1].right = false
				l.cells[(i)*l.length+j-1].visited = true
				l.cells[(i)*l.length+j].left = false
				j -= 1
				l.visited_cells++
				path = append(path, left)
				l.all_path = append(l.all_path, left)
				break

			}
		} else {
			switch path[len(path)-1] {
			case up:
				l.all_path = append(l.all_path, down)
				i += 1
				break

			case right:

				l.all_path = append(l.all_path, left)
				j -= 1
				break

			case down:

				l.all_path = append(l.all_path, up)
				i -= 1
				break

			case left:
				l.all_path = append(l.all_path, right)
				j += 1
				break
			}

			path = path[:len(path)-1]
		}
	}
}

func (l *Land) unvisited_neighbour(i, j int) []int {
	return_array := make([]int, 0)
	// Cima

	if l.is_acessable(i-1, j) {
		if !l.cells[(i-1)*l.length+j].visited {
			return_array = append(return_array, up)
		}
	}
	// Direita
	if l.is_acessable(i, j+1) {
		if !l.cells[(i)*l.length+j+1].visited {
			return_array = append(return_array, right)
		}

	}

	// Baixo
	if l.is_acessable(i+1, j) {
		if !l.cells[(i+1)*l.length+j].visited {
			return_array = append(return_array, down)
		}
	}

	// Esquerda
	if l.is_acessable(i, j-1) {
		if !l.cells[(i)*l.length+j-1].visited {
			return_array = append(return_array, left)
		}
	}

	return return_array
}

func (l *Land) have_unvisited_neighbour(i, j int) bool {
	// Cima
	if l.is_acessable(i-1, j) && !l.cells[(i-1)*l.length+j].visited {
		return true
	}

	// Direita
	if l.is_acessable(i, j+1) && !l.cells[(i)*l.length+j+1].visited {
		return true
	}

	// Baixo
	if l.is_acessable(i+1, j) && !l.cells[(i+1)*l.length+j].visited {
		return true
	}

	// Esquerda
	if l.is_acessable(i, j-1) && !l.cells[(i)*l.length+j-1].visited {
		return true
	}
	return false

}

func (l *Land) is_acessable(i, j int) bool {
	if i < 0 || i >= l.height || j < 0 || j >= l.length {
		return false

	}
	return true
}
