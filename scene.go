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

func (s *scene) addScene(data *scene) error {
	s.Children = append(s.Children, data)
	return nil
}