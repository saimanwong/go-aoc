package day07

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type myFI struct {
	Name        string
	Type        string // file or dir
	Size        int    // if dir, is total size
	Children    []*myFI
	PrevDirPath string // fullpath
}

type virtualFS map[string]*myFI // name -> *myFI

type Problem struct {
	fs virtualFS
}

func (p *Problem) SetInput(input []string) {
	vfs := virtualFS{}
	pwd := ""
	for _, line := range input {
		spl := strings.Split(line, " ")
		if spl[0] == "$" { // cmd to exec
			if spl[1] == "cd" { // change dir
				if spl[2] == "/" { // root, only appears once
					vfs["/"] = &myFI{Name: "/", Type: "dir", Children: []*myFI{}}
					pwd = "/"
					continue
				}
				if spl[2] == ".." { // up
					pwd = vfs[pwd].PrevDirPath
					continue
				}

				// jump to dir
				dirName := filepath.Join(pwd, spl[2])
				if _, ok := vfs[dirName]; !ok {
					log.Fatalln("dir does not exist", dirName)
				}
				vfs[pwd].Children = append(vfs[pwd].Children, vfs[dirName])
				pwd = dirName
				continue
			}
			if spl[1] == "ls" {
				continue
			}
		}

		// create dir
		if spl[0] == "dir" {
			dirName := filepath.Join(pwd, spl[1])
			if _, ok := vfs[dirName]; !ok {
				vfs[dirName] = &myFI{}
			}
			vfs[dirName].Name = spl[1]
			vfs[dirName].Type = "dir"
			vfs[dirName].PrevDirPath = pwd
			continue
		}

		// create file
		size := toolbox.ToInt(spl[0])
		fname := filepath.Join(pwd, spl[1])
		if _, ok := vfs[fname]; !ok {
			vfs[fname] = &myFI{}
		}
		vfs[fname].Name = spl[1]
		vfs[fname].Type = "file"
		vfs[fname].Size = size
		vfs[fname].PrevDirPath = pwd
		vfs[pwd].Children = append(vfs[pwd].Children, vfs[fname])

		// populate directory size
		curr := vfs[fname]
		for curr.PrevDirPath != "" {
			prev := vfs[curr.PrevDirPath]
			prev.Size += size
			curr = prev
		}
	}
	p.fs = vfs
}

func (p *Problem) Run() {
	sum1 := 0
	for _, d := range p.fs {
		if d.Type != "dir" {
			continue
		}
		if d.Size > 100000 {
			continue
		}
		sum1 += d.Size
	}
	fmt.Println("Part 1:", sum1)

	const (
		capacity   = 70000000
		updateSize = 30000000
	)
	unused := capacity - p.fs["/"].Size
	p2 := capacity
	for _, d := range p.fs {
		if d.Type != "dir" || d.Name == "/" {
			continue
		}
		if updateSize > unused+d.Size { // not enough space
			continue
		}
		if d.Size > p2 {
			continue
		}
		p2 = d.Size
	}
	fmt.Println("Part 2:", p2)
}

func (p *Problem) debug(path string) {
	b, err := json.MarshalIndent(p.fs[path], "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
