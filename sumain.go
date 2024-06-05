package main

import "fmt"

type Component interface {
	Operation() string
}

type Leaf struct {
	Name string
}

func (l *Leaf) Operation() string {
	return l.Name
}

type Composite struct {
	Name     string
	Children []Component
}

func (c *Composite) Operation() string {
	result := c.Name + ": ["
	for i, child := range c.Children {
		if i > 0 {
			result += ", "
		}
		result += child.Operation()
	}
	result += "]"
	return result
}

func (c *Composite) Add(child Component) {
	c.Children = append(c.Children, child)
}

func main() {

	leaf1 := &Leaf{Name: "Leaf 1"}
	leaf2 := &Leaf{Name: "Leaf 2"}
	leaf3 := &Leaf{Name: "Leaf 3"}

	composite1 := &Composite{Name: "Composite 1"}
	composite1.Add(leaf1)
	composite1.Add(leaf2)

	composite2 := &Composite{Name: "Composite 2"}
	composite2.Add(composite1)
	composite2.Add(leaf3)

	fmt.Println(composite2.Operation())
}
