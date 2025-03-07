package systems

import (
	"fmt"

	"github.com/qiekn/constants"
	"github.com/qiekn/entities"
	"github.com/qiekn/managers"
)

func InitCurrentLevel() {
	managers.GetEntityManager().Clear()
	layer := managers.GetLevelManager().GetCurrentLevel()
	if layer == nil {
		fmt.Println("layer is nil")
		return
	}
	m := layer.Height
	n := layer.Width
	for y := range m {
		for x := range n {
			switch layer.Data[y*n+x] - layer.Gid {
			case constants.Tile_Wall:
				entities.NewWall(x, y)
			case constants.Tile_Target:
				entities.NewTarget(x, y)
			case constants.Tile_Player:
				entities.NewPlayer(x, y)
			case constants.Tile_Box:
				entities.NewBox(x, y)
			}
		}
	}
}

func SwitchToNextLevel() {
	managers.GetLevelManager().SwitchToNext()
	InitCurrentLevel()
}

func SwitchToPrevLevel() {
	managers.GetLevelManager().SwitchToPrev()
	InitCurrentLevel()
}
