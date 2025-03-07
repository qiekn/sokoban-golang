package managers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Level struct {
	Width  int
	Height int
	Gid    int
	Data   []int
}

type LevelManager struct {
	levels  []Level
	current int
}

// singleton
var (
	levelManagerInstance *LevelManager
	levelManagerOnce     sync.Once
)

func GetLevelManager() *LevelManager {
	levelManagerOnce.Do(func() {
		levelManagerInstance = &LevelManager{
			levels:  []Level{},
			current: 0,
		}
		levelManagerInstance.LoadLevelsFromJSON(
			"./assets/tilemap/1.json",
			"./assets/tilemap/2.json",
		)
	})
	return levelManagerInstance
}

func (lm *LevelManager) LoadLevelsFromJSON(filePaths ...string) {

	lm.levels = []Level{}

	for _, filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal("level manager: read level failed: ", filePath, err)
		}

		var tileMap struct {
			Layers []struct {
				Data   []int `json:"data"`
				Width  int   `json:"width"`
				Height int   `json:"height"`
			} `json:"layers"`
			Tilesets []struct {
				Firstgid int `json:"firstgid"`
			} `json:"tilesets"`
		}

		err = json.Unmarshal(data, &tileMap)
		if err != nil {
			log.Fatal("level manager: json unmarshal failed ", err)
		}

		if len(tileMap.Layers) > 0 {
			level := Level{
				Width:  tileMap.Layers[0].Width,
				Height: tileMap.Layers[0].Height,
				Gid:    tileMap.Tilesets[0].Firstgid,
				Data:   tileMap.Layers[0].Data,
			}
			lm.levels = append(lm.levels, level)
		}
		fmt.Println("level manager: load level ", filePath)
	}
	fmt.Println("level manager: level loaded finised")
}

func (lm *LevelManager) GetCurrentLevel() *Level {
	if int(lm.current) < len(lm.levels) {
		return &lm.levels[lm.current]
	}
	return nil
}

func (lm *LevelManager) SwitchToNext() {
	if lm.current+1 < len(lm.levels) {
		lm.current++
	} else {
		fmt.Println("已经到底了喵")
	}
}

func (lm *LevelManager) SwitchToPrev() {
	if lm.current-1 >= 0 {
		lm.current--
	} else {
		fmt.Println("已经到顶了喵")
	}
}
