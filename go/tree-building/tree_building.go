package tree

import (
	"errors"
	"sort"
)

// Record store the flattened tree structure
type Record struct {
	ID, Parent int
}

// Node is a node in the tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree from a record
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		r1 := records[i]
		r2 := records[j]
		// if r1.Parent == r2.Parent {
		// 	return r1.ID < r2.ID
		// }
		// return r1.Parent < r2.Parent
		return r1.ID < r2.ID
	})
	nodes := make([]*Node, len(records))
	// pop the first item off records since it should be root.
	rootRecord, records := records[0], records[1:]
	if !isRoot(rootRecord) {
		return nil, errors.New("no root record")
	}
	nodes[0] = makeNode(rootRecord.ID)
	for i, record := range records {
		if !isValid(record) {
			return nil, errors.New("invalid record")
		}
		if i+1 != record.ID {
			return nil, errors.New("non-continuous IDs")
		}
		if i > 0 {
			p := records[i-1]
			if p.ID == record.ID && p.Parent == record.Parent {
				return nil, errors.New("duplicate record")
			}
		}
		parent := nodes[record.Parent]
		n := makeNode(record.ID)
		nodes[n.ID] = n
		parent.Children = append(parent.Children, n)
	}
	return nodes[0], nil
}

func isRoot(r Record) bool {
	return r.ID == 0 && r.Parent == 0
}

func isValid(r Record) bool {
	return r.ID > r.Parent
}

func makeNode(ID int) *Node {
	return &Node{
		ID: ID,
	}
}
