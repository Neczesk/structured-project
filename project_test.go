package project

import (
	"fmt"
	"testing"
)

func TestProject(t *testing.T) {
	p, err := NewBlankProject()
	if err != nil {
		t.Errorf("Error creating project")
	}
	s := p.SceneTree.newScene("Hi mom", "Hello mother")
	p.SceneTree.AddScene("root", &s)
	if p.SceneTree.countScenes() != 1 {
		t.Errorf("Expected %d scenes, found %d", p.SceneTree.CountScenes(), 1)
	}
	s1 := p.SceneTree.newScene("I am a new scene", "Hell")
	p.SceneTree.AddScene(s.ID, &s1)
	s2 := p.SceneTree.newScene("I am a leaf", "leaf leaf leaf")
	p.SceneTree.AddScene(s.ID, &s2)
	fmt.Print(p.summarize())
}
