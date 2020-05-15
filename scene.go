package project

import (
	"encoding/json"
	"fmt"
)

//The scene struct has methods to operate on its own data to enable in place operations on a scene within the scene tree. TODO: Rework the ID property to be reassigned on tree structure transactions.
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

//This is an in place update of the targeted scene. It keeps the ID and the scene's children intact, updating the scene's title and text
func (s *scene) updateScene(scenedata scene) {

	//Save properties of original scene that should be kept
	c := s.Children
	i := s.ID

	*s = scenedata
	//Reassign saved properties
	s.Children = c
	s.ID = i
}

func (s *scene) addScene(data *scene) error {
	s.Children = append(s.Children, data)
	return nil
}
