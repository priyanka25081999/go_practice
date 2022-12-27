package main

import "fmt"

type Cache struct {
	queue Queue
	hmap  Hash
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

type Hash map[string]*Node

var SIZE = 5

func NewCache() Cache {
	return Cache{queue: newQueue(), hmap: Hash{}}
}

func newQueue() Queue {
	head := &Node{}
	tail := &Node{}

	// initially we don't have any element, they are pointing to each other
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) check(str string) {
	node := &Node{}

	// Check if the value is present in the queue or not
	// if not present then add the value at the beginning
	// if present, then remove the value from the queue and add it to the beginning

	// here, we are checking in the hashmap
	if nodeVal, ok := c.hmap[str]; ok {
		// node presents
		node = c.Remove(nodeVal)
	} else {
		// not present, hence add it
		node = &Node{Data: str}
	}

	c.Add(node)
	c.hmap[str] = node
}

func (c *Cache) Remove(node *Node) *Node {
	left := node.Left
	right := node.Right

	// remove the middle element of linkedlist logic
	left.Right = right
	right.Left = left

	c.queue.Length -= 1
	delete(c.hmap, node.Data)

	return node
}

func (c *Cache) Add(node *Node) {
	// add it to the beginning
	// update the current head, for that take the current head value into temp
	temp := c.queue.Head.Right
	c.queue.Head.Right = node
	node.Left = c.queue.Head
	node.Right = temp
	temp.Left = node

	// increase the queue length by 1
	c.queue.Length += 1
	// check the overflow condition
	if c.queue.Length > SIZE {
		// then remove the last element of the queue
		c.Remove(c.queue.Tail.Left)
	}
}

func (c *Cache) display() {
	c.queue.display()
}

func (q *Queue) display() {
	node := q.Head.Right
	fmt.Printf("%d : [", q.Length)

	for i := 0; i < q.Length; i++ {
		fmt.Printf("{Node: %s}", node.Data)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}

		// moving to right nodes
		node = node.Right
	}

	fmt.Println("]")
}

func main() {
	fmt.Println("Start Cache!!!")

	// initiate the new cache
	cache := NewCache()
	for _, word := range []string{"cat", "dog", "tree", "mountains", "tree", "hat", "mat", "tree"} {
		cache.check(word)
		cache.display()
	}
}
