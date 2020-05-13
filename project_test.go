package project

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProject(t *testing.T) {
	p := NewBlankProject()
	s := p.SceneTree.newScene("Hi mom", "Hello mother")
	p.SceneTree.addScene("root", &s)
	fmt.Println(p.summarize())
	if p.SceneTree.countScenes() != 0 {
		t.Errorf("Expected %d scenes, found %d", 0, p.SceneTree.countScenes())
	}
	s1 := p.SceneTree.newScene("I am a new scene", "Hell")
	p.SceneTree.addScene(s.ID, &s1)
	s2 := p.SceneTree.newScene("I am a leaf", "leaf leaf leaf")
	p.SceneTree.addScene(s.ID, &s2)
	err := p.SaveProject("toast.project")
	d, err := ioutil.ReadFile("/Users/kyle/Documents/toast.project")
	if err != nil {
		t.Errorf("Error reading file")
	}
	newProject, err := loadProject(d)
	if p.summarize() != newProject.summarize() {
		t.Errorf("Saved and loaded project not the same as original")
	}
	x := NewBlankProject()
	x, err = LoadProjectFile("/Users/kyle/Documents/toast.project")

	if x.summarize() != p.summarize() {
		t.Errorf("Saved and loaded project not the same as original")
	}
	fmt.Println(x.SceneTree.Root.Children[1].ID)
	fmt.Println(x.summarize())
	summary := x.getScene("1")
	if summary != nil {
		fmt.Println(summary.summarize())
	} else {
		t.Error("Returned scene is nil")
	}

	// s3 := x.SceneTree.newScene("I love", "Antonia")
	// x.SceneTree.addScene(s2.ID, &s3)
	// jString, err := json.Marshal(s3)
	// x.UpdateSceneFromJSON("1", string(jString))
	// fmt.Println(x.getScene("1").summarize())
}
