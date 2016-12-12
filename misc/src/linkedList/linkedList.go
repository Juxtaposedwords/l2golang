package linkedList
import (
	"fmt"
	"math/rand"
)
type node struct{
	label string
	next *node
}
type nodeList []*node
