package project

import (
	"strconv"
	"strings"
)

type sceneTree struct {
	Root      *scene
	idCounter int
}

func newBlankSceneTree() sceneTree {
	root := scene{"", "", "", make([]*scene, 0)}
	return sceneTree{&root, 1}
}

func (s *sceneTree) newScene(title, text string) scene {
	return scene{title, s.generateID(), text, make([]*scene, 0)}
}

func (s *sceneTree) addScene(parent string, newScene *scene) {
	target := s.findSceneByID(parent)
	target.addScene(newScene)
}

func (s *sceneTree) findSceneByID(search string) *scene {
	queue := make([]*scene, 0)
	queue = append(queue, s.Root)
	for len(queue) > 0 {
		nextup := queue[0]
		queue = queue[1:]
		if nextup.ID == search {
			return nextup
		}
		if len(nextup.Children) > 0 {
			for _, child := range nextup.Children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

func (s *sceneTree) updateSceneByID(target string, scenedata scene) {
	queue := make([]*scene, 0)
	queue = append(queue, s.Root)
	for len(queue) > 0 {
		nextup := queue[0]
		queue = queue[1:]
		if nextup.ID == target {
			nextup = &scenedata
		}
		if len(nextup.Children) > 0 {
			for _, child := range nextup.Children {
				queue = append(queue, child)
			}
		}
	}
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
