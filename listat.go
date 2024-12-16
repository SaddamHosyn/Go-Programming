package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(nodePointer *NodeL, pos int) *NodeL {
	index := 0
	count := 0
	head := nodePointer

	for head != nil {
		index++
		head = head.Next
	}
	if pos <= index {
		for nodePointer != nil {
			if count == pos {
				return nodePointer
			}
			count++
			nodePointer = nodePointer.Next
		}
	}
	return nil
}
