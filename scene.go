package project

import (
	"encoding/json"
	"fmt"
)

type scene struct {
	Title    string
	ID       string
	Text     string
	Children []*scene
}

func (s *scene) isLeaf() bool {
	if len(s.Children) == 0 {
		return true
	}
	return false
}

func (s *scene) summarize() string {
	return fmt.Sprintf("This scene is called: %s, with ID: %s", s.Title, s.Text)
}

func (s *scene) asJson() string {
	result, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(result)
}

func (s *scene) addChildScene(scenedata *scene) {
	s.Children = append(s.Children, scenedata)
}

func (s *scene) updateScene(scenedata scene) {
	c := s.Children
	i := s.ID

	fmt.Println(s.summarize())
	*s = scenedata
	s.Children = c
	s.ID = i
	fmt.Println(s.summarize())
}
