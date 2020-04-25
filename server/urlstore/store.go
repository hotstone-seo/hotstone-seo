package urlstore

// Copyright 2019 hotstone-seo. Modifications: 1) 'id' of node 2) Delete node by ID
// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// OG: https://github.com/go-ozzo/ozzo-routing/blob/master/store.go

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var (
	// ParameterSize is size of paramter
	ParameterSize = 512
)

// Store is a radix tree that supports storing data with parametric keys and retrieving them back with concrete keys.
// When retrieving a data item with a concrete key, the matching parameter names and values will be returned as well.
// A parametric key is a string containing tokens in the format of "<name>", "<name:pattern>", or "<:pattern>".
// Each token represents a single parameter.
type Store interface {
	Add(id int64, key string, data interface{}) int
	Get(path string) (data interface{}, param *Parameter)
	Delete(id int64) bool
	String() string
	Count() int
}

type storeImpl struct {
	root  *node // the root node of the radix tree
	count int   // the number of data nodes in the tree
}

// NewStore return new instance of Store
func NewStore() Store {
	return &storeImpl{
		root: &node{
			static:      true,
			id:          -1,
			children:    make([]*node, 256),
			childrenMap: map[byte]*node{},
			pchildren:   make([]*node, 0),
			pindex:      -1,
			pnames:      []string{},
		},
	}
}

func (s *storeImpl) Count() int {
	return s.count
}

// Add adds a new data item with the given parametric key.
// The number of parameters in the key is returned.
func (s *storeImpl) Add(id int64, key string, data interface{}) int {
	s.count++
	return s.root.add(id, key, data, s.count)
}

// Get returns the data item matching the given concrete key.
// If the data item was added to the store with a parametric key before, the matching
// parameter names and values will be returned as well.
func (s *storeImpl) Get(path string) (interface{}, *Parameter) {
	pvalues := make([]string, ParameterSize)
	data, pnames, _ := s.root.get(path, pvalues)
	return data, NewParameter(pnames, pvalues)
}

// func (s *storeImpl) GetAsMap(path string, size int) (data interface{}, pnames []string, pmap map[string]string) {
// 	data, pnames, pvalues := s.GetAsSlice(path, size)
// 	pmap = make(map[string]string)
// 	for i, name := range pnames {
// 		pmap[name] = pvalues[i]
// 	}
// 	return
// }

// Delete deletes the data item matching the given ID. It returns existness of deleted item.
func (s *storeImpl) Delete(id int64) bool {
	found, _, _, _ := s.root.delete(id)
	if found {
		s.count--
	}

	return found
}

// String dumps the radix tree kept in the store as a string.
func (s *storeImpl) String() string {
	return s.root.print(0)
}

// node represents a radix trie node
type node struct {
	static bool // whether the node is a static node or param node

	id   int64       // ID of this node in the DB
	key  string      // the key identifying this node
	data interface{} // the data associated with this node. nil if not a data node.

	order    int // the order at which the data was added. used to be pick the first one when matching multiple
	minOrder int // minimum order among all the child nodes and this node

	childrenMap map[byte]*node // child static nodes, tracked in map for speed reason
	children    []*node        // child static nodes, indexed by the first byte of each child key
	pchildren   []*node        // child param nodes

	regex  *regexp.Regexp // regular expression for a param node containing regular expression key
	pindex int            // the parameter index, meaningful only for param node
	pnames []string       // the parameter names collected from the root till this node
}

func (n *node) String() string {
	var regexStr string
	if n.regex != nil {
		regexStr = n.regex.String()
	}
	var s strings.Builder
	fmt.Fprintf(&s, "{")
	fmt.Fprintf(&s,
		"static:%t id:%d key:%s data:%+v order:%d minOrder:%d pindex:%d pnames:%+v regex:%s ",
		n.static, n.id, n.key, n.data, n.order, n.minOrder, n.pindex, n.pnames, regexStr)

	fmt.Fprintf(&s, "children:[")
	for i, c := range n.children {
		if c != nil {
			fmt.Fprintf(&s, "%s:%s ", string(i), c.key)
		}
	}
	fmt.Fprintf(&s, "] ")

	fmt.Fprintf(&s, "pchildren:[")
	for i, c := range n.pchildren {
		if c != nil {
			fmt.Fprintf(&s, "%s:%s ", string(i), c.key)
		}
	}
	fmt.Fprintf(&s, "]")

	fmt.Fprintf(&s, "}")

	return s.String()
}

// add adds a new data item to the tree rooted at the current node.
// The number of parameters in the key is returned.
func (n *node) add(id int64, key string, data interface{}, order int) int {
	matched := 0

	// find the common prefix
	for ; matched < len(key) && matched < len(n.key); matched++ {
		if key[matched] != n.key[matched] {
			break
		}
	}

	// fmt.Printf("[%s] add - key: %s - matched:%d - len(n.key):%d - len(key):%d \n", n.key, key, matched, len(n.key), len(key))

	// ME: n.key as base chars. It will compare key with n.key. If common key < n.key: split.

	// ME: if common chars == n.key
	if matched == len(n.key) {
		if matched == len(key) {
			// the node key is the same as the key: make the current node as data node
			// if the node is already a data node, ignore the new data since we only care the first matched node
			if n.data == nil {
				n.id = id
				n.data = data
				n.order = order
			}
			// n.id = id
			return n.pindex + 1
		}

		// the node key is a prefix of the key: create a child node
		newKey := key[matched:]
		// fmt.Printf("[%s] newKey: %s\n", n.key, newKey)

		// try adding to a static child
		if child := n.children[newKey[0]]; child != nil {
			// fmt.Printf("[%s] static children[%c]: %s - newKey: %s\n", n.key, newKey[0], child.key, newKey)
			if pn := child.add(id, newKey, data, order); pn >= 0 {
				return pn
			}
		}
		// try adding to a param child
		for _, child := range n.pchildren {
			// fmt.Printf("[%s] param child: %s\n", n.key, child.key)
			if pn := child.add(id, newKey, data, order); pn >= 0 {
				return pn
			}
		}

		// fmt.Printf("[%s] before addChild\n", n.key)
		return n.addChild(id, newKey, data, order)
	}

	if matched == 0 || !n.static {
		// no common prefix, or partial common prefix with a non-static node: should skip this node
		return -1
	}

	// fmt.Printf("[%s] before add (split)\n", n.key)

	// the node key shares a partial prefix with the key: split the node key
	n1 := &node{
		static:      true,
		id:          n.id,
		key:         n.key[matched:],
		data:        n.data,
		order:       n.order,
		minOrder:    n.minOrder,
		pchildren:   n.pchildren,
		children:    n.children,
		childrenMap: n.childrenMap,
		pindex:      n.pindex,
		pnames:      n.pnames,
	}

	// if matched == len(key) {
	// 	n.id = id
	// }

	n.id = -1
	n.key = key[0:matched]
	n.data = nil
	n.pchildren = make([]*node, 0)
	n.children = make([]*node, 256)
	n.childrenMap = map[byte]*node{}
	n.children[n1.key[0]] = n1
	n.childrenMap[n1.key[0]] = n1

	return n.add(id, key, data, order)
}

// addChild creates static and param nodes to store the given data
func (n *node) addChild(id int64, key string, data interface{}, order int) int {
	param := findFirstParam(key)
	// find the first occurrence of a param token

	if param == nil {
		// no param token: done adding the child
		child := &node{
			static:      true,
			id:          id,
			key:         key,
			minOrder:    order,
			data:        data,
			order:       order,
			children:    make([]*node, 256),
			childrenMap: map[byte]*node{},
			pchildren:   make([]*node, 0),
			pindex:      n.pindex,
			pnames:      n.pnames,
		}
		n.children[key[0]] = child
		n.childrenMap[key[0]] = child

		return child.pindex + 1
	}

	if param.StringBefore != "" {
		// param token occurs after a static string
		child := &node{
			static:      true,
			id:          -1,
			key:         param.StringBefore,
			minOrder:    order,
			children:    make([]*node, 256),
			childrenMap: map[byte]*node{},
			pchildren:   make([]*node, 0),
			pindex:      n.pindex,
			pnames:      n.pnames,
		}
		n.children[key[0]] = child
		n.childrenMap[key[0]] = child

		n = child
	}

	// add param node
	child := &node{
		static:      false,
		id:          -1,
		key:         param.Raw,
		minOrder:    order,
		children:    make([]*node, 256),
		childrenMap: map[byte]*node{},
		pchildren:   make([]*node, 0),
		pindex:      n.pindex,
		pnames:      n.pnames,
	}

	if param.Pattern != "" {
		// the param token contains a regular expression
		child.regex = regexp.MustCompile("^" + param.Pattern)
	}
	pnames := make([]string, len(n.pnames)+1)
	copy(pnames, n.pnames)
	pnames[len(n.pnames)] = param.Name
	child.pnames = pnames
	child.pindex = len(pnames) - 1
	n.pchildren = append(n.pchildren, child)

	if param.AtLastPos {
		// the param token is at the end of the key
		child.id = id
		child.data = data
		child.order = order
		return child.pindex + 1
	}

	// process the rest of the key
	return child.addChild(id, param.StringAfter, data, order)
}

func printDebug(key string, data interface{}, pnames, pvalues []string) {
	// fmt.Printf("[%s][data:%v][pnames:%+v][pvalues:%+v]\n\n", key, data, pnames, pvalues)
}

// get returns the data item with the key matching the tree rooted at the current node
func (n *node) get(key string, pvalues []string) (data interface{}, pnames []string, order int) {
	order = math.MaxInt32

repeat:
	// fmt.Printf("[n]: %s\n", n)
	// fmt.Printf("[key]: %+v\n", key)

	if n.static {
		// check if the node key is a prefix of the given key
		// a slightly optimized version of strings.HasPrefix
		nkl := len(n.key)
		if nkl > len(key) {
			printDebug(key, data, pnames, pvalues)
			return
		}
		for i := nkl - 1; i >= 0; i-- {
			if n.key[i] != key[i] {
				printDebug(key, data, pnames, pvalues)
				return
			}
		}
		key = key[nkl:]
	} else if n.regex != nil {
		// param node with regular expression
		if n.regex.String() == "^.*" {
			pvalues[n.pindex] = key
			key = ""
		} else if match := n.regex.FindStringIndex(key); match != nil {
			pvalues[n.pindex] = key[0:match[1]]
			key = key[match[1]:]
		} else {
			printDebug(key, data, pnames, pvalues)
			return
		}
	} else {
		// param node matching non-"/" characters
		i, kl := 0, len(key)
		for ; i < kl; i++ {
			if key[i] == '/' || key[i] == '-' {
				pvalues[n.pindex] = key[0:i]
				key = key[i:]
				break
			}
		}
		if i == kl {
			pvalues[n.pindex] = key
			key = ""
		}
	}

	if len(key) > 0 {
		// find a static child that can match the rest of the key
		if child := n.children[key[0]]; child != nil {
			if len(n.pchildren) == 0 {
				// use goto to avoid recursion when no param children
				n = child
				goto repeat
			}
			data, pnames, order = child.get(key, pvalues)
		}
	} else if n.data != nil {
		// do not return yet: a param node may match an empty string with smaller order
		data, pnames, order = n.data, n.pnames, n.order
	}

	// try matching param children
	tvalues := pvalues
	allocated := false
	for _, child := range n.pchildren {
		if child.minOrder >= order {
			continue
		}
		if data != nil && !allocated {
			tvalues = make([]string, len(pvalues))
			allocated = true
		}
		if d, p, s := child.get(key, tvalues); d != nil && s < order {
			if allocated {
				for i := child.pindex; i < len(p); i++ {
					pvalues[i] = tvalues[i]
				}
			}
			data, pnames, order = d, p, s
		}
	}

	printDebug(key, data, pnames, pvalues)
	return
}

func (n *node) print(level int) string {
	r := fmt.Sprintf("%v{id: %d, static: %t, key: %v, regex: %v, data: %v, order: %v, minOrder: %v, pindex: %v, pnames: %v}\n", strings.Repeat(" ", level<<2), n.id, n.static, n.key, n.regex, n.data, n.order, n.minOrder, n.pindex, n.pnames)
	for _, child := range n.children {
		if child != nil {
			r += child.print(level + 1)
		}
	}
	for _, child := range n.pchildren {
		r += child.print(level + 1)
	}
	return r
}

func (n *node) delete(id int64) (found, foundInThisNode bool, numChildStatic, numChildParam int) {
	if id == n.id {
		n.id = -1
		n.data = nil
		n.regex = nil

		return true, true, len(n.childrenMap), len(n.pchildren)
	}

	// fmt.Printf("ID: %d key:%s children: %d childrenMap: %d\n", n.id, n.key, len(n.children), len(n.childrenMap))

	// Delete child static
	for indexByte, child := range n.childrenMap {
		// fmt.Printf("ID: %d key:%s - CHILD: %d key:%s\n", n.id, n.key, child.id, child.key)
		found, foundInTheChild, numChildStatic, numChildParam := child.delete(id)
		if foundInTheChild || found {
			if numChildStatic == 0 && numChildParam == 0 {
				n.children[indexByte] = nil
				delete(n.childrenMap, indexByte)
			}

			return true, false, len(n.childrenMap), len(n.pchildren)
		}

	}

	// Delete child param
	found, foundInTheChildParam := false, false
	idxDeleted, numChildParam := -1, -1
	for i, child := range n.pchildren {

		found, foundInTheChildParam, _, numChildParam = child.delete(id)

		if foundInTheChildParam || found {
			if foundInTheChildParam {
				idxDeleted = i
			}

			break
		}
	}

	if foundInTheChildParam || found {
		// fmt.Printf(">>> ID TARGET DELETE:%d CURRENT ID:%d CURRENT KEY:%s IDX DELETED:%d numChildParam:%d\n", id, n.id, n.key, idxDeleted, numChildParam)

		if idxDeleted != -1 && numChildStatic == 0 && numChildParam == 0 {

			// delete item in n.pchildren (slice)
			copy(n.pchildren[idxDeleted:], n.pchildren[idxDeleted+1:]) // Shift a[i+1:] left one index.
			n.pchildren[len(n.pchildren)-1] = nil                      // Erase last element (write zero value).
			n.pchildren = n.pchildren[:len(n.pchildren)-1]
		}

		return true, false, len(n.childrenMap), len(n.pchildren)
	}

	// fmt.Printf("ID NOT FOUND:%d KEY: %s FOUNDINTHECHILD: %t\n", n.id, n.key, foundInTheChildParam)
	return false, false, len(n.childrenMap), len(n.pchildren)
}
