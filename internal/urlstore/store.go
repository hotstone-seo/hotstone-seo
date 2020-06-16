package urlstore

// Copyright 2019 hotstone-seo. Modifications: 1) 'id' of node 2) Delete node by ID
// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// OG: https://github.com/go-ozzo/ozzo-routing/blob/master/store.go

var (
	// ParameterSize is size of paramter
	ParameterSize = 512
)

// Store is a radix tree that supports storing data with parametric keys and retrieving them back with concrete keys.
// When retrieving a data item with a concrete key, the matching parameter names and values will be returned as well.
// A parametric key is a string containing tokens in the format of "<name>", "<name:pattern>", or "<:pattern>".
// Each token represents a single parameter.
// @mock
type Store interface {
	Add(id int64, key string, data interface{}) int
	Get(path string) (data interface{}, param *Parameter)
	Delete(id int64) bool
	String() string
	Count() int
	Reset()
}

type storeImpl struct {
	root  *node // the root node of the radix tree
	count int   // the number of data nodes in the tree
}

// NewStore return new instance of Store
// @ctor
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

func (s *storeImpl) Count() int {
	return s.count
}

func (s *storeImpl) Reset() {
	s.root.childrenMap = map[byte]*node{}
	s.root.children = make([]*node, 256)
	s.root.pchildren = make([]*node, 0)
	s.root.pindex = -1
	s.root.pnames = []string{}
	s.count = 0
}
