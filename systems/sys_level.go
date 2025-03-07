package systems

import (
	"fmt"

	"github.com/qiekn/constants"
	"github.com/qiekn/entities"
	"github.com/qiekn/managers"
)

func InitCurrentLevel() {
	managers.GetEntityManager().Clear()
	layers := managers.GetLevelManager().GetCurrentLevelLayers()
	if layers == nil {
		fmt.Println("sys_level: layers is nil")
		return
	}
	for _, layer := range *layers {
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
}

func SwitchToNextLevel() {
	managers.GetLevelManager().SwitchToNext()
	InitCurrentLevel()
}

func SwitchToPrevLevel() {
	managers.GetLevelManager().SwitchToPrev()
	InitCurrentLevel()
}
