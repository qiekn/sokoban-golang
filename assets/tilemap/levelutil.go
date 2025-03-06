package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

////////////////////////////////////////////////////////////////////////
//                      Original JSON Structure                       //
////////////////////////////////////////////////////////////////////////

type TileMap struct {
	CompressionLevel int       `json:"compressionlevel"`
	Height           int       `json:"height"`
	Infinite         bool      `json:"infinite"`
	Layers           []Layer   `json:"layers"`
	NextLayerID      int       `json:"nextlayerid"`
	NextObjectID     int       `json:"nextobjectid"`
	Orientation      string    `json:"orientation"`
	RenderOrder      string    `json:"renderorder"`
	TiledVersion     string    `json:"tiledversion"`
	TileHeight       int       `json:"tileheight"`
	Tilesets         []Tileset `json:"tilesets"`
	TileWidth        int       `json:"tilewidth"`
	Type             string    `json:"type"`
	Version          string    `json:"version"`
	Width            int       `json:"width"`
}

type Layer struct {
	Data    []int   `json:"data"`
	Height  int     `json:"height"`
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Opacity float64 `json:"opacity"`
	Type    string  `json:"type"`
	Visible bool    `json:"visible"`
	Width   int     `json:"width"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
}

type Tileset struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

////////////////////////////////////////////////////////////////////////
//                              Process                               //
////////////////////////////////////////////////////////////////////////

func readJSON(filepath string) (*TileMap, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tileMap TileMap
	err = json.Unmarshal(data, &tileMap)
	if err != nil {
		return nil, err
	}
	return &tileMap, nil
}

func writeJSON(filepath string, tileMap *TileMap) error {
	data, err := json.MarshalIndent(tileMap, "", "  ")
	if err != nil {
		return nil
	}
	return os.WriteFile(filepath, data, 0644)
}

func trimLayer(layer *Layer) {
	grid := make([][]int, layer.Height)
	for i := range layer.Height {
		grid[i] = layer.Data[i*layer.Width : (i+1)*layer.Width]
	}

	// 查找边界
	top, bottom, left, right := 0, layer.Height-1, 0, layer.Width-1
	found := false

	// 找到 top（第一个非 0 行）
	for top <= bottom {
		for _, val := range grid[top] {
			if val != 0 {
				found = true
				break
			}
		}
		if found {
			break
		}
		top++
	}
	found = false

	// 找到 bottom（最后一个非 0 行）
	for bottom >= top {
		for _, val := range grid[bottom] {
			if val != 0 {
				found = true
				break
			}
		}
		if found {
			break
		}
		bottom--
	}
	found = false

	// 找到 left（第一个非 0 列）
	for left <= right {
		for i := top; i <= bottom; i++ {
			if grid[i][left] != 0 {
				found = true
				break
			}
		}
		if found {
			break
		}
		left++
	}
	found = false

	// 找到 right（最后一个非 0 列）
	for right >= left {
		for i := top; i <= bottom; i++ {
			if grid[i][right] != 0 {
				found = true
				break
			}
		}
		if found {
			break
		}
		right--
	}

	// 截取新的 data
	newWidth := right - left + 1
	newHeight := bottom - top + 1
	newData := []int{}

	for i := top; i <= bottom; i++ {
		newData = append(newData, grid[i][left:right+1]...)
	}

	// 更新图层
	layer.Data = newData
	layer.Width = newWidth
	layer.Height = newHeight
}

////////////////////////////////////////////////////////////////////////
//                                Main                                //
////////////////////////////////////////////////////////////////////////

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: leveltuil <command>")
		fmt.Println("\t- trim")
		fmt.Println("\t- override")
		return
	}

	command := args[1]

	input := "./assets/tilemap"
	output := "./levels"

	if command == "override" {
		files, err := os.ReadDir(output)
		if err != nil {
			fmt.Println("ReadDir faied\n", err)
			return
		}
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
				sourcePath := filepath.Join(output, file.Name())
				destPath := filepath.Join(input, file.Name())

				// move files
				err := os.Rename(sourcePath, destPath)
				if err != nil {
					fmt.Println("Move file failed: ", file.Name(), "\n", err)
					return
				} else {
					fmt.Println("Override file: ", file.Name())
				}
			}
		}
	}

	if command == "trim" {
		files, err := os.ReadDir(input)
		if err != nil {
			fmt.Println("ReadDir faied\n", err)
			return
		}
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
				sourcePath := filepath.Join(input, file.Name())
				destPath := filepath.Join(output, file.Name())
				fmt.Println("Processing file: ", sourcePath)

				// read json
				tileMap, err := readJSON(sourcePath)
				if err != nil {
					fmt.Println("readJSON failed: ", sourcePath, "\n", err)
					return
				}

				// process
				if len(tileMap.Layers) > 0 {
					trimLayer(&tileMap.Layers[0])
					tileMap.Width = tileMap.Layers[0].Width
					tileMap.Height = tileMap.Layers[0].Height
				}

				// write back
				err = writeJSON(destPath, tileMap)
				if err != nil {
					fmt.Println("writeJSON filed: ", destPath, "\n", err)
					return
				}
			}
		}
	} // end trim
	fmt.Println("Mission Success!")
}
