package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	manager  *Manager
	player   *Entity
	board    [][]rune
	tileSize int
	images   *Images
}

func NewGame() *Game {
	images, err := loadImages()
	if err != nil {
		log.Fatal(err)
	}
	manager := NewManager()
	game := &Game{
		manager:  manager,
		tileSize: 16,
		images:   images,
	}
	game.initBoard()
	game.initEntities()
	return game
}

// '#' 墙壁, '.' 地板, 'B' 箱子, 'T' 目标, 'P' 玩家初始位置
func (g *Game) initBoard() {
	g.board = [][]rune{
		[]rune("#######"),
		[]rune("#.....#"),
		[]rune("#.BTP.#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
}

func (g *Game) initEntities() {
	for y, row := range g.board {
		for x, ch := range row {
			switch ch {
			case 'P':
				player := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'P'},
				}
				g.manager.AddEntity(player)
				g.player = player
				g.board[y][x] = '.'
			case 'B':
				box := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'B'},
				}
				g.manager.AddEntity(box)
				g.board[y][x] = '.'
			case 'T':
				target := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'T'},
				}
				g.manager.AddEntity(target)
				g.board[y][x] = '.'
			}
		}
	}
}

func (g *Game) Update() error {
	dx, dy := 0, 0
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		dy = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		dy = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		dx = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		dx = 1
	}

	if dx != 0 || dy != 0 {
		g.movePlayer(dx, dy)
	}

	return nil
}

func (g *Game) movePlayer(dx, dy int) {
	newX := g.player.Position.X + dx
	newY := g.player.Position.Y + dy

	// hit the wall
	if g.isWall(newX, newY) {
		return
	}

	box := g.getBoxAt(newX, newY)
	// hit the box
	if box != nil {
		// (assumed) new postion of box after being pushed
		boxNewX := box.Position.X + dx
		boxNewY := box.Position.Y + dy
		// can't push box
		if g.isWall(boxNewX, boxNewY) || g.getBoxAt(boxNewX, boxNewY) != nil {
			return
		}
		// can push box
		box.Position.X = boxNewX
		box.Position.Y = boxNewY
	}

	// move player
	g.player.Position.X = newX
	g.player.Position.Y = newY
}

func (g *Game) isWall(x, y int) bool {
	if y < 0 || y >= len(g.board) || x < 0 || x >= len(g.board[0]) {
		return true
	}
	return g.board[y][x] == '#'
}

func (g *Game) getBoxAt(x, y int) *Entity {
	for _, e := range g.manager.entities {
		if e.Renderable != nil && e.Renderable.Char == 'B' && e.Position.X == x && e.Position.Y == y {
			return e
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw background
	screen.Fill(color.RGBA{56, 63, 88, 255})

	// offset (used to center level)
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()

	mapPixelWidth := len(g.board[0]) * g.tileSize
	mapPixelHeight := len(g.board) * g.tileSize

	offsetX := (screenWidth - mapPixelWidth) / 2
	offsetY := (screenHeight - mapPixelHeight) / 2

	// draw wall
	for y, row := range g.board {
		for x, ch := range row {
			if ch == '#' {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x*g.tileSize)+float64(offsetX), float64(y*g.tileSize)+float64(offsetY))
				screen.DrawImage(g.images.Wall, opts)
			}
		}
	}
	// draw entities with reanderable component
	for _, e := range g.manager.GetEntitiesWithRenderable() {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(
			float64(e.Position.X*g.tileSize+offsetX),
			float64(e.Position.Y*g.tileSize+offsetY),
		)
		switch e.Renderable.Char {
		case 'P':
			screen.DrawImage(g.images.Player, opts)
		case 'B':
			screen.DrawImage(g.images.Box, opts)
		case 'T':
			screen.DrawImage(g.images.Point, opts)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return len(g.board[0]) * g.tileSize, len(g.board) * g.tileSize
	return 16 * 15, 16 * 10
}
