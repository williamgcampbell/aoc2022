package _7

import (
	_ "embed"
	"fmt"
	"github.com/williamgcampbell/aoc2022/internal"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"log"
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

	fs := &FileSystem{}
	for _, line := range lines {
		fs.ReplayLine(line)
	}

	if part1 {
		var totalSize int
		for _, dir := range fs.Directories {
			size := dir.GetSize()
			if size <= 100000 {
				totalSize += size
			}
		}
		return fmt.Sprintf("%d", totalSize)
	}

	unusedSpace := 70000000 - fs.RootDirectory.GetSize()
	spaceNeeded := 30000000 - unusedSpace
	var minSizeNeeded int
	for _, dir := range fs.Directories {
		size := dir.GetSize()
		fmt.Printf("Directory: %s Size %d\n", dir.Path, size)
		if size >= spaceNeeded && (size < minSizeNeeded || minSizeNeeded == 0) {
			minSizeNeeded = size
		}
	}
	return fmt.Sprintf("%d", minSizeNeeded)
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

type FileSystem struct {
	CurrentDirectory *Directory
	RootDirectory    *Directory
	Directories      []*Directory
}

func (fs *FileSystem) ReplayLine(line string) {
	if strings.HasPrefix(line, "$ ls") {
		return
	}

	if strings.HasPrefix(line, "$ cd ") {
		cd := line[5:]
		if cd == ".." {
			fs.CurrentDirectory = fs.CurrentDirectory.Parent
		} else if cd == "/" {
			// special case for this folder
			newDir := NewDirectory(nil, cd)
			fs.Directories = append(fs.Directories, newDir)
			fs.CurrentDirectory = newDir
			fs.RootDirectory = newDir
		} else {
			if d, ok := fs.CurrentDirectory.GetChildDirectory(cd); ok {
				fs.CurrentDirectory = d
			} else {
				// This should never happen.
				// Directories will always be added via the ls command before running cd into them
				log.Fatalf("Could not find directory %s", cd)
			}
		}
	} else if strings.HasPrefix(line, "dir ") {
		dirName := line[4:]
		if _, ok := fs.CurrentDirectory.GetChildDirectory(dirName); !ok {
			dir := NewDirectory(fs.CurrentDirectory, dirName)
			fs.Directories = append(fs.Directories, dir)
			fs.CurrentDirectory.AddDirectory(dir)
		}
	} else {
		lsFile := strings.Split(line, " ")
		file := &File{
			Name: lsFile[1],
			Size: internal.MustAtoI(lsFile[0]),
		}
		fs.CurrentDirectory.AddFile(file)
	}
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

func (d *Directory) AddDirectory(dir *Directory) {
	d.Directories = append(d.Directories, dir)
}

func (d *Directory) GetChildDirectory(name string) (*Directory, bool) {
	for _, dir := range d.Directories {
		if dir.Name == name {
			return dir, true
		}
	}
	return nil, false
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
