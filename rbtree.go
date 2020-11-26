package rbtree

const (
	red   = 0
	black = 1
)

// CompareCmd
//     -1: less
//      0: equal
//      1: more
type CompareCmd func(interface{}, interface{}) int

// RbNode the red-black tree node
type RbNode struct {
	K, V interface{}

	color int

	p *RbNode

	left  *RbNode
	right *RbNode
}

// RbTree the structure of red-black tree
type RbTree struct {
	root *RbNode
	null *RbNode

	cmp CompareCmd

	unique bool
	size   int
}

// NewRbTree new
//  cmp:  method of key comparison
//  unique: whether to allow duplicate keys
func NewRbTree(cmp CompareCmd, unique bool) *RbTree {
	null := &RbNode{
		color: black,
	}
	return &RbTree{
		null:   null,
		root:   null,
		cmp:    cmp,
		unique: unique,
	}
}

//      |                            |
//      y                            x
//     / \   <--- left rotate ---   / \
//    x   γ  --- right rotate ---> α   y
//   / \                              / \
//  α   β                            β   γ
// /
func (rb *RbTree) leftRotate(x *RbNode) {
	y := x.right
	x.right = y.left
	if y.left != rb.null {
		y.left.p = x
	}

	y.p = x.p
	if x.p == rb.null {
		rb.root = y
	} else {
		if x == x.p.left {
			x.p.left = y
		} else {
			x.p.right = y
		}
	}

	y.left = x
	x.p = y
}

func (rb *RbTree) rightRotate(y *RbNode) {
	x := y.left
	y.left = x.right
	if x.right != rb.null {
		x.right.p = y
	}

	x.p = y.p
	if y.p == rb.null {
		rb.root = x
	} else {
		if y == y.p.left {
			y.p.left = x
		} else {
			y.p.right = x
		}
	}

	x.right = y
	y.p = x
}

// Insert insert
func (rb *RbTree) Insert(key, value interface{}) {
	if key == nil {
		return
	}

	// z is the insert node
	// x is the inserted position
	// y is the parent node of the insertion position
	y := rb.null
	x := rb.root

	z := &RbNode{
		K:     key,
		V:     value,
		color: red,
		left:  rb.null,
		right: rb.null,
	}

	for x != rb.null {
		y = x
		switch rb.cmp(z.K, x.K) {
		case -1:
			x = x.left
		case 0:
			// the key already exists in the red-black tree,
			// and the unique flag is true (equivalent to replacing the node)
			if rb.unique {
				x.V = z.V
				return
			}
			x = x.right
		case 1:
			x = x.right
		}
	}
	z.p = y

	// insert in the leaf node, as a general case
	if y == rb.null { // empty
		rb.root = z
	} else {
		switch rb.cmp(key, y.K) {
		case -1:
			y.left = z
		case 0, 1:
			y.right = z
		}
	}

	rb.size++
	rb.insertFixup(z)
}

func (rb *RbTree) insertFixup(z *RbNode) {
	// z is the node to start repair
	// y is the uncle of z
	var y *RbNode
	for z.p.color == red {
		if z.p == z.p.p.left {
			y = z.p.p.right
			if y.color == red { // case 1
				z.p.color = black
				y.color = black
				z.p.p.color = red
				z = z.p.p
			} else {
				if z == z.p.right { // case 2
					z = z.p
					rb.leftRotate(z)
				}
				z.p.color = black // case 3
				z.p.p.color = red
				rb.rightRotate(z.p.p)
			}
		} else {
			y = z.p.p.left
			if y.color == red { // case 4
				z.p.color = black
				y.color = black
				z.p.p.color = red
				z = z.p.p
			} else {
				if z == z.p.left { // case 5
					z = z.p
					rb.rightRotate(z)
				}
				z.p.color = black // case 6
				z.p.p.color = red
				rb.leftRotate(z.p.p)
			}
		}
	}
	rb.root.color = black
}

func (rb *RbTree) transplant(u, v *RbNode) {
	if u.p == rb.null {
		rb.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (rb *RbTree) deleteFixup(x *RbNode) {
	for x != rb.root && x.color != black {
		if x == x.p.left {
			w := x.p.right
			if w.color == red { // case 1
				w.color = black
				x.p.color = red
				rb.leftRotate(x.p)
				w = x.p.right
			}
			if w.left.color == black && w.right.color == black { // case 2
				w.color = red
				x = x.p
			} else if w.right.color == black { // case 3
				w.left.color = black
				w.color = red
				rb.rightRotate(w)
				w = x.p.right
			}
			w.color = x.p.color // case 4
			x.p.color = black
			w.right.color = black
			rb.leftRotate(x.p)
			x = rb.root
		} else {
			w := x.p.right
			if w.color == red { // case 5
				w.color = black
				x.p.color = red
				rb.rightRotate(x.p)
				w = x.p.left
			}
			if w.right.color == black && w.left.color == black { // case 6
				w.color = red
				x = x.p
			} else if w.left.color == black { // case 7
				w.right.color = black
				w.color = red
				rb.leftRotate(w)
				w = x.p.left
			}
			w.color = x.p.color // case 8
			x.p.color = black
			w.right.color = black
			rb.rightRotate(x.p)
			x = rb.root
		}
	}
	x.color = black
}

// DeleteByKey delete nodes by key.
// If unique is false, there may be multiple deleted nodes.
func (rb *RbTree) DeleteByKey(key interface{}) {
	deleteNodeList := rb.findByKey(key)
	for _, node := range deleteNodeList {
		rb.delete(node)
	}
}

// DeleteByNode delete the node directly.
func (rb *RbTree) DeleteByNode(z *RbNode) {
	rb.delete(z)
}

func (rb *RbTree) delete(z *RbNode) {
	var x *RbNode
	y := z
	yOrigninalColor := y.color
	if z.left == rb.null {
		x = z.right
		rb.transplant(z, z.right)
	} else if z.right == rb.null {
		x = z.left
		rb.transplant(z, z.left)
	} else {
		y = rb.Next(z) //successor node of z
		yOrigninalColor = y.color
		x = y.right
		if y.p == z {
			x.p = z
		} else {
			rb.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		rb.transplant(z, y)
		y.left = z.left
		y.left.p = y
		y.color = z.color
	}
	if yOrigninalColor == black {
		rb.deleteFixup(x)
	}
	rb.size--
}

// Find find nodes by key
func (rb *RbTree) Find(key interface{}) []*RbNode {
	return rb.findByKey(key)
}

func (rb *RbTree) findByKey(key interface{}) []*RbNode {
	z := rb.root
	for z != rb.null {
		exist := false
		switch rb.cmp(key, z.K) {
		case -1:
			z = z.left
		case 0:
			exist = true
		case 1:
			z = z.right
		}
		if exist {
			break
		}
	}
	if z != rb.null {
		var result []*RbNode
		result = append(result, z)
		next := z
		for {
			next = rb.Next(next)
			if next != rb.null && rb.cmp(key, next.K) == 0 {
				result = append(result, next)
			} else {
				break
			}
		}
		pre := z
		for {
			pre = rb.Prev(pre)
			if pre != rb.null && rb.cmp(key, pre.K) == 0 {
				result = append(result, pre)
			} else {
				break
			}
		}
		return result
	}
	return nil
}

// Size number of nodes
func (rb *RbTree) Size() int {
	return rb.size
}

// First returns the first node (in sort order) of the tree.
func (rb *RbTree) First() *RbNode {
	return rb.first(rb.root)
}

func (rb *RbTree) first(node *RbNode) *RbNode {
	var item *RbNode
	for item = node; item.left != rb.null; {
		item = item.left
	}
	return item
}

// Last returns the last node (in sort order) of the tree.
func (rb *RbTree) Last() *RbNode {
	return rb.last(rb.root)
}

func (rb *RbTree) last(node *RbNode) *RbNode {
	var item *RbNode
	for item = node; item.right != rb.null; {
		item = item.right
	}
	return item
}

// Next returns the next node (in sort order) of the tree.
func (rb *RbTree) Next(node *RbNode) *RbNode {
	if node.right != rb.null {
		return rb.first(node.right)
	}
	parent := rb.null
	for {
		if node.p == rb.null {
			break
		}
		parent = node.p
		if node == parent.left {
			break
		}
		node = parent
	}
	return parent
}

// Prev returns the previous node (in sort order) of the tree.
func (rb *RbTree) Prev(node *RbNode) *RbNode {
	if node.left != rb.null {
		return rb.last(node.left)
	}
	parent := rb.null
	for {
		if node.p == rb.null {
			break
		}
		parent = node.p
		if node == parent.right {
			break
		}
		node = parent
	}
	return parent
}
