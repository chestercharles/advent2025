package dial


type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func (ddl DoublyLinkedList) Append (value int) {
	newNode := &Node{Value: value}
	if (ddl.Len == 0 ) {
		ddl.Head = newNode
		ddl.Tail = newNode
		ddl.Len = 1
	}

	ddl.Tail.Next = newNode
	ddl.Tail = newNode
	ddl.Len = ddl.Len + 1
}

type Dial struct {
	pos *Node
}

func NewDial() *Dial {
	head := &Node{Value: 50}
	tail := head

	for i := 50; i < 99; i++ {
		tail.Next = &Node{Value: tail.Value + 1, Prev: tail}
		tail = tail.Next
	}

	for i := 0; i< 50; i++ {
		tail.Next = &Node{Value: i, Prev: tail}
		tail = tail.Next
	}

	tail.Next = head
	head.Prev = tail

	newDial := &Dial{pos: head}

    return newDial
}

func (d *Dial) TurnRight(value int) (int) {
	zeroes := 0

	for i := 0; i < value; i++ {
		d.pos = d.pos.Next

		if (d.pos.Value == 0) {
			zeroes = zeroes + 1
		}
	}

	return zeroes
}

func (d *Dial) TurnLeft(value int) (int) {
	zeroes := 0

	for i := 0; i< value; i++ {
		d.pos = d.pos.Prev

		if (d.pos.Value == 0) {
			zeroes = zeroes + 1
		}
	}

	return zeroes
}

func (d *Dial) Read() int {
	return d.pos.Value
}