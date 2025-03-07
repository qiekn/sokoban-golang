package managers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/qiekn/constants"
)

type LevelManager struct {
	levels  []level
	current int
}

type level struct {
	Layers []layer
	Height int
	Width  int
}

type layer struct {
	Id     int
	Name   string
	Width  int
	Height int
	Gid    int
	Data   []int
}

// singleton
var (
	levelManagerInstance *LevelManager
	levelManagerOnce     sync.Once
)

func GetLevelManager() *LevelManager {
	levelManagerOnce.Do(func() {
		levelManagerInstance = &LevelManager{
			levels:  []level{},
			current: 0,
		}
		filePaths := levelManagerInstance.initLevelFilePaths()
		levelManagerInstance.loadLevelsFromJSON(filePaths...)
	})
	return levelManagerInstance
}

func (lm *LevelManager) initLevelFilePaths() []string {
	var filePaths []string
	for i := range constants.LevelCounts {
		filePath := fmt.Sprintf("./assets/tilemap/%02d.json", i+1)
		if _, err := os.Stat(filePath); err == nil { // level file exists
			filePaths = append(filePaths, filePath)
		} else if !os.IsNotExist(err) { // error
			log.Printf("Error checking %s: %v", filePath, err)
		}
	}
	sort.Strings(filePaths)
	return filePaths
}

func (lm *LevelManager) loadLevelsFromJSON(filePaths ...string) {

	lm.levels = []level{}

	for _, filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal("level manager: read level failed: ", filePath, err)
		}

		var levelJSON struct {
			Height int `json:"height"`
			Width  int `json:"width"`
			Layers []struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Data   []int  `json:"data"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"layers"`
			Tilesets []struct {
				Firstgid int `json:"firstgid"`
			} `json:"tilesets"`
		}

		err = json.Unmarshal(data, &levelJSON)
		if err != nil {
			log.Fatal("level manager: json unmarshal failed ", err)
		}

		var level level
		level.Width = levelJSON.Width
		level.Height = levelJSON.Height
		for _, Layer := range levelJSON.Layers {
			layer := layer{
				Id:     Layer.Id,
				Name:   Layer.Name,
				Width:  Layer.Width,
				Height: Layer.Height,
				Gid:    levelJSON.Tilesets[0].Firstgid,
				Data:   Layer.Data,
			}
			level.Layers = append(level.Layers, layer)
		}
		lm.levels = append(lm.levels, level)

		fmt.Println("level manager: load level ", filePath)
	} // end filePaths for-loop
	fmt.Println("level manager: level loaded finised")
}

func (lm *LevelManager) GetCurrentLevel() *level {
	if int(lm.current) < len(lm.levels) {
		return &lm.levels[lm.current]
	}
	return nil
}

func (lm *LevelManager) GetCurrentLevelLayers() *[]layer {
	level := lm.GetCurrentLevel()
	return &level.Layers
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
