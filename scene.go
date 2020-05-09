package project

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
