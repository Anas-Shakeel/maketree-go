package models

// NodeType represents whether a node is a file or directory.
type NodeType int

const (
	DirNode NodeType = iota
	FileNode
)

// Node struct to store parsed tree
type Node struct {
	Type     NodeType // Type of the node
	Name     string   // Name of the entry/node
	Level    int      // Indentation level
	Parent   *Node    // Parent Node
	Children []*Node  // Childrens of this node
}

// IsDir retruns true if node is a directory, false otherwise
func (n *Node) IsDir() bool {
	return n.Type == DirNode
}

// IsFile retruns true if node is a file, false otherwiseF
func (n *Node) IsFile() bool {
	return n.Type == FileNode
}

// AddChild adds a child into the children slice of a directory node.
func (n *Node) AddChild(child *Node) {
	if n.Type == DirNode {
		n.Children = append(n.Children, child)
	}
}

// HasChildren returns true if node n has any children.
func (n *Node) HasChildren() bool {
	if n.Type == FileNode {
		return false
	}

	if len(n.Children) <= 0 {
		return false
	}

	return true
}
