package project

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProject(t *testing.T) {
	p := NewBlankProject()
	s := p.SceneTree.newScene("Hi mom", "Hello mother")
	p.SceneTree.AddScene("root", &s)
	if p.SceneTree.countScenes() != 1 {
		t.Errorf("Expected %d scenes, found %d", p.SceneTree.countScenes(), 1)
	}
	s1 := p.SceneTree.newScene("I am a new scene", "Hell")
	p.SceneTree.AddScene(s.ID, &s1)
	s2 := p.SceneTree.newScene("I am a leaf", "leaf leaf leaf")
	p.SceneTree.AddScene(s.ID, &s2)
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

	summary := x.getScene("1")
	fmt.Println(summary.summarize())
	s3 := x.SceneTree.newScene("I love", "Antonia")
	x.SceneTree.AddScene(s2.ID, &s3)
	x.UpdateScene("1", s3)
	fmt.Println(x.getScene("1").summarize())
}
