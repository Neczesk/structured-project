package project

import (
	"fmt"
)

//Project contains metadata about the project, and the sceneTree which holds all project scenes
type Project struct {
	author    string
	title     string
	sceneTree sceneTree
}

//NewBlankProject creates a new project with as many blank, zero valued initials as possible
func NewBlankProject() (Project, error) {
	scenetree := newBlankSceneTree()
	p := Project{"", "", scenetree}
	return p, nil
}

/*
	Make a string of the format: "This project has x scenes, y of which are leaves, with a wordcount of : z"
*/
func (p *Project) summarize() string {
	sceneCount := p.SceneTree.countScenes()
	leafCount := countLeaves(p.SceneTree.Root)
	wordCount := countWords(p.SceneTree.Root)
	return fmt.Sprintf("This project has %d scenes, %d of which are leaves, with a wordcount of %d\n", sceneCount, leafCount, wordCount)
}
