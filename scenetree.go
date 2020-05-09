package project

import (
	"strconv"
	"strings"
)

type sceneTree struct {
	Root       *scene
	SceneTable map[string]*scene
	idCounter  int
}

func newBlankSceneTree() sceneTree {
	root := scene{"", "", "", make([]*scene, 0)}
	st := make(map[string]*scene)
	st["root"] = &root
	return sceneTree{&root, st, 1}
}

func (s *sceneTree) newScene(title, text string) scene {
	return scene{title, s.generateID(), text, make([]*scene, 0)}
}

func (s *sceneTree) AddScene(parent string, newScene *scene) {
	destination := s.SceneTable[parent]
	destination.Children = append(destination.Children, newScene)
	s.SceneTable[newScene.ID] = newScene
}

//CountScenes uses a depth first tree traversal to count all scenes in the current scenetree.
func (s *sceneTree) countScenes() int {
	return countChildren(s.Root)
}

func countChildren(s *scene) int {
	result := len(s.Children)
	for _, scene := range s.Children {
		result += countChildren(scene)
	}
	return result
}

//CountLeaves does a depth first tree traversal, and counts all scenes which are "leaves", meaning those without any further children.
func countLeaves(s *scene) int {
	result := 0
	for _, scene := range s.Children {
		if len(scene.Children) == 0 {
			result++
		} else {
			result += countLeaves(scene)
		}
	}
	return result
}

//CountWords uses a depth first tree traversal, adds up all words in leaves, defined by spltting the leaf's text by the " " character.
func countWords(s *scene) int {
	sum := 0
	if s.isLeaf() {
		sum = len(strings.Split(s.Text, " "))
	} else {
		for _, scene := range s.Children {
			sum += countWords(scene)
		}
	}
	return sum
}

func (s *sceneTree) generateID() string {
	result := s.idCounter
	s.idCounter++
	return strconv.Itoa(result)
}
