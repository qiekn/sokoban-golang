package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	manager  *Manager
	player   *Entity  // 玩家实体
	board    [][]rune // 游戏地图：字符矩阵（'#'代表墙壁，'.'代表地板）
	tileSize int      // 每个瓦片的像素大小
	images   *Images
}

// NewGame 初始化游戏，构建地图并生成实体
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

// 初始化游戏地图（关卡），这里用简单的字符数组表示
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

// 根据地图数据创建实体，并将特定字符替换为地板（'.'）以便后续逻辑处理
func (g *Game) initEntities() {
	for y, row := range g.board {
		for x, ch := range row {
			switch ch {
			case 'P':
				// 创建玩家实体
				player := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'P'},
				}
				g.manager.AddEntity(player)
				g.player = player
				g.board[y][x] = '.' // 替换为地板
			case 'B':
				// 创建箱子实体
				box := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'B'},
				}
				g.manager.AddEntity(box)
				g.board[y][x] = '.'
			case 'T':
				// 创建目标实体（仅作渲染提示，也可以用于判断是否完成关卡）
				target := &Entity{
					Position:   &Position{X: x, Y: y},
					Renderable: &Renderable{Char: 'T'},
				}
				g.manager.AddEntity(target)
				g.board[y][x] = '.'
				// 墙壁直接用地图数据处理，不必生成实体
			}
		}
	}
}

// ---------------------
// 游戏更新与渲染逻辑
// ---------------------

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

// movePlayer 处理玩家移动和推动箱子的逻辑
func (g *Game) movePlayer(dx, dy int) {
	newX := g.player.Position.X + dx
	newY := g.player.Position.Y + dy

	// 检测是否碰到墙壁
	if g.isWall(newX, newY) {
		return
	}

	// 检查新位置是否有箱子
	box := g.getBoxAt(newX, newY)
	if box != nil {
		// 箱子被推挤后的新位置
		boxNewX := box.Position.X + dx
		boxNewY := box.Position.Y + dy
		// 如果箱子后面是墙或者另一个箱子，则不能推挤
		if g.isWall(boxNewX, boxNewY) || g.getBoxAt(boxNewX, boxNewY) != nil {
			return
		}
		// 推动箱子
		box.Position.X = boxNewX
		box.Position.Y = boxNewY
	}

	// 移动玩家
	g.player.Position.X = newX
	g.player.Position.Y = newY
}

// 判断给定坐标是否为墙壁
func (g *Game) isWall(x, y int) bool {
	if y < 0 || y >= len(g.board) || x < 0 || x >= len(g.board[0]) {
		return true
	}
	return g.board[y][x] == '#'
}

// 获取指定位置的箱子实体（如果存在）
func (g *Game) getBoxAt(x, y int) *Entity {
	for _, e := range g.manager.entities {
		if e.Renderable != nil && e.Renderable.Char == 'B' && e.Position.X == x && e.Position.Y == y {
			return e
		}
	}
	return nil
}

// Draw 负责将地图和实体绘制到屏幕上
func (g *Game) Draw(screen *ebiten.Image) {
	// draw background
	screen.Fill(color.RGBA{56, 63, 88, 255})

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
	// return g.tileSize * 15, g.tileSize * 15
	return 160, 120
}
