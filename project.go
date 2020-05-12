package project

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//Project contains metadata about the project, and the sceneTree which holds all project scenes
type Project struct {
	Author    string
	Title     string
	SceneTree sceneTree
}

//NewBlankProject creates a new project with as many blank, zero valued initials as possible
func NewBlankProject() *Project {
	scenetree := newBlankSceneTree()
	p := Project{"", "", scenetree}
	return &p
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

//SaveProject saves a project to the filename given in the parameters in the running user's Documents directory.
func (p *Project) SaveProject(filename string) error {
	saveData, err := json.Marshal(p)
	cwd, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	savepath := path.Join(cwd, "Documents")
	savepath = path.Join(savepath, filename)
	return ioutil.WriteFile(savepath, []byte(saveData), 0600)
}

func loadProject(data []byte) (*Project, error) {
	loadedProject := Project{}
	json.Unmarshal(data, &loadedProject)
	return &loadedProject, nil
}

//LoadProjectFile takes a file name, returns a Project object
func LoadProjectFile(filename string) (*Project, error) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return NewBlankProject(), err
	}
	return loadProject(d)
}

func (p *Project) AsJson() string {
	result, _ := json.Marshal(p)
	output := string(result)
	return output
}

//getScene takes a sceneid string and returns the scene object it refers to.
func (p *Project) getScene(id string) *scene {
	result := p.SceneTree.SceneTable[id]
	return result
}

func (p *Project) GetSceneExport(id string) string {
	return p.getScene(id).asJson()
}

//UpdateScene takes a sceneid string and a scene object, and overwrites the scene with that id with the scene parameter
func (p *Project) updateScene(id string, scenedata scene) {
	p.SceneTree.SceneTable[id] = &scenedata
}

func (p *Project) UpdateSceneFromJSON(id string, sceneJSON string) {

	newScene := scene{}
	json.Unmarshal([]byte(sceneJSON), &newScene)
	p.updateScene(id, newScene)
}