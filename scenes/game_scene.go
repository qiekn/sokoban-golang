package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	. "github.com/qiekn/components"
	. "github.com/qiekn/entities"
)

type GameScene struct {
	isloaded      bool
	entityManager *EntityManager
	player        *Entity
	board         [][]rune
	tileSize      int
	textures      *Textures
}

func NewGameScene() *GameScene {
	gameScene := &GameScene{
		entityManager: NewEntityManager(),
		tileSize:      16,
		textures:      loadImages(),
		isloaded:      false,
	}
	gameScene.initBoard()
	gameScene.initEntities()
	return gameScene
}

func (g *GameScene) IsLoaded() bool { return g.isloaded }

func (g *GameScene) OnEnter() {
	fmt.Println("Game Scene Enter")
}

func (g *GameScene) OnExit() {
	g.isloaded = false
	fmt.Println("Game Scene Exit")
}

func (g *GameScene) Start() {
	g.isloaded = true
}

func (g *GameScene) Update() {
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
}

func (g *GameScene) Draw(screen *ebiten.Image) {
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
				screen.DrawImage(g.textures.Wall, opts)
			}
		}
	}
	// draw entities with reanderable component
	for _, e := range g.entityManager.GetEntitiesWithRenderable() {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(
			float64(e.Position.X*g.tileSize+offsetX),
			float64(e.Position.Y*g.tileSize+offsetY),
		)
		switch e.Renderable.Char {
		case 'P':
			screen.DrawImage(g.textures.Player, opts)
		case 'B':
			screen.DrawImage(g.textures.Box, opts)
		case 'T':
			screen.DrawImage(g.textures.Point, opts)
		}
	}
}

func (g *GameScene) UpdateSceneId() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ExitSceneId
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return PauseSceneId
	}
	return GameSceneId
}

var _ Scene = (*GameScene)(nil)

////////////////////////////////////////////////////////////////////////
//                            New Contine                             //
////////////////////////////////////////////////////////////////////////

// '#' 墙壁, '.' 地板, 'B' 箱子, 'T' 目标, 'P' 玩家初始位置
func (g *GameScene) initBoard() {
	g.board = [][]rune{
		[]rune("#######"),
		[]rune("#.....#"),
		[]rune("#.BTP.#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
}

func (g *GameScene) initEntities() {
	for y, row := range g.board {
		for x, ch := range row {
			switch ch {
			case 'P':
				player := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'P'},
				}
				g.entityManager.AddEntity(player)
				g.player = player
				g.board[y][x] = '.'
			case 'B':
				box := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'B'},
				}
				g.entityManager.AddEntity(box)
				g.board[y][x] = '.'
			case 'T':
				target := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'T'},
				}
				g.entityManager.AddEntity(target)
				g.board[y][x] = '.'
			}
		}
	}
}

func (g *GameScene) movePlayer(dx, dy int) {
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

func (g *GameScene) isWall(x, y int) bool {
	if y < 0 || y >= len(g.board) || x < 0 || x >= len(g.board[0]) {
		return true
	}
	return g.board[y][x] == '#'
}

func (g *GameScene) getBoxAt(x, y int) *Entity {
	for _, e := range g.entityManager.GetAllEntities() {
		if e.Renderable != nil && e.Renderable.Char == 'B' && e.Position.X == x && e.Position.Y == y {
			return e
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////
//                           Load Textures                            //
////////////////////////////////////////////////////////////////////////

type Textures struct {
	Player *ebiten.Image
	Wall   *ebiten.Image
	Box    *ebiten.Image
	Point  *ebiten.Image
}

func loadImages() *Textures {
	player, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}
	wall, _, err := ebitenutil.NewImageFromFile("assets/images/wall.png")
	if err != nil {
		log.Fatal(err)
	}
	box, _, err := ebitenutil.NewImageFromFile("assets/images/box.png")
	if err != nil {
		log.Fatal(err)
	}
	point, _, err := ebitenutil.NewImageFromFile("assets/images/point.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Textures{
		Player: player,
		Wall:   wall,
		Box:    box,
		Point:  point,
	}
}
