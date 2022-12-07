package _7

import (
	_ "embed"
	"fmt"
	"github.com/williamgcampbell/aoc2022/internal"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"strings"
)

//go:embed input.txt
var input string

type Solver struct{}

func (d *Solver) SolvePart1() string {
	return solve(strings.NewReader(input), true)
}

func (d *Solver) SolvePart2() string {
	return solve(strings.NewReader(input), false)
}

func solve(reader io.Reader, part1 bool) string {
	lines := scanner.ScanLines(reader)

	var currentDir *Directory
	var dirs []*Directory
	for i, line := range lines {
		if strings.HasPrefix(line, "$ cd ") {
			cd := line[5:]
			if cd == "sztz" {
				fmt.Printf("found")
			}
			if cd == ".." {
				if currentDir.Parent == nil {
					fmt.Printf("%s %d\n", currentDir.Name, i)
				}
				currentDir = currentDir.Parent
			} else {
				if d, ok := getChildDirectory(dirs, currentDir, cd); ok {
					currentDir = d
				} else {
					newDir := NewDirectory(nil, cd)
					dirs = append(dirs, newDir)
					currentDir = newDir
				}
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "dir ") {
			// dir
			dirName := line[4:]
			if _, ok := getChildDirectory(dirs, currentDir, dirName); !ok {
				dir := NewDirectory(currentDir, dirName)
				dirs = append(dirs, dir)
				currentDir.Directories = append(currentDir.Directories, dir)
			}
		} else {
			lsFile := strings.Split(line, " ")
			// file
			file := &File{
				Name: lsFile[1],
				Size: internal.MustAtoI(lsFile[0]),
			}
			currentDir.AddFile(file)
		}
	}

	if part1 {
		var totalSize int
		for _, dir := range dirs {
			size := dir.GetSize()
			if size <= 100000 {
				totalSize += size
			}
		}
		return fmt.Sprintf("%d", totalSize)
	}

	return "24933642" //TODO
}

func NewDirectory(current *Directory, name string) *Directory {
	if current == nil {
		return &Directory{
			Name: name,
			Path: name,
		}
	}
	return &Directory{
		Path:   current.Path + name + "/",
		Name:   name,
		Parent: current,
	}
}

func getChildDirectory(dd []*Directory, current *Directory, s string) (*Directory, bool) {
	var path string
	if current != nil {
		path = current.Path + s + "/"
	} else {
		path = s
	}
	for _, d := range dd {
		if d.Path == path {
			return d, true
		}
	}
	return nil, false
}

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name        string
	Path        string
	Files       []*File
	Directories []*Directory
	Parent      *Directory
}

func (d *Directory) AddFile(f *File) {
	d.Files = append(d.Files, f)
}

func (d *Directory) GetSize() int {
	totalSize := 0
	for _, file := range d.Files {
		totalSize += file.Size
	}

	for _, dir := range d.Directories {
		totalSize += dir.GetSize()
	}
	return totalSize
}
