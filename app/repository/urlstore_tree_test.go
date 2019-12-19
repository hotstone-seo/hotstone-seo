// Copyright 2019 Tiket.Com. Modifications: 1) 'id' of node 2) Delete node by ID
// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type storeTestEntry struct {
	id        int
	key, data string
	params    int
}

func testStorePlay(t *testing.T) {
	tests := []struct {
		id       string
		entries  []storeTestEntry
		expected string
	}{
		{
			"playground",
			[]storeTestEntry{
				{1, "/gopher/bumper.png", "1", 0},
				{2, "/gopher/bumper192x108.png", "2", 0},
				{11, "/gopher/bumper.jpg", "1a", 0},
				{3, "/gopher/doc.png", "3", 0},
				{4, "/gopher/bumper320x180.png", "4", 0},
				{5, "/gopher/docpage.png", "5", 0},
				{6, "/gopher/doc.png", "6", 0},
				{7, "/gopher/doc", "7", 0},
			},
			`{key: , regex: <nil>, data: <nil>, order: 0, minOrder: 0, pindex: -1, pnames: []}
    {key: /gopher/, regex: <nil>, data: <nil>, order: 1, minOrder: 1, pindex: -1, pnames: []}
        {key: bumper, regex: <nil>, data: <nil>, order: 1, minOrder: 1, pindex: -1, pnames: []}
            {key: .png, regex: <nil>, data: 1, order: 1, minOrder: 1, pindex: -1, pnames: []}
            {key: 192x108.png, regex: <nil>, data: 2, order: 2, minOrder: 2, pindex: -1, pnames: []}
            {key: 320x180.png, regex: <nil>, data: 4, order: 4, minOrder: 4, pindex: -1, pnames: []}
        {key: doc, regex: <nil>, data: 7, order: 7, minOrder: 3, pindex: -1, pnames: []}
            {key: .png, regex: <nil>, data: 3, order: 3, minOrder: 3, pindex: -1, pnames: []}
            {key: page.png, regex: <nil>, data: 5, order: 5, minOrder: 5, pindex: -1, pnames: []}
`,
		},
	}
	for _, test := range tests {
		h := NewURLStoreTree()
		for _, entry := range test.entries {
			n := h.Add(entry.id, entry.key, entry.data)
			assert.Equal(t, entry.params, n, test.id+" > "+entry.key+" > param count =")
		}
		assert.Equal(t, test.expected, h.String(), test.id+" > store.String() =")
	}
}

func TestStorePlay2(t *testing.T) {
	var children []*string = make([]*string, 256)
	hello := "hello"
	children[0] = &hello

	null := "NIL"

	start := time.Now()
	for i, val := range children {
		value := val
		if val == nil {
			value = &null
		}
		fmt.Printf("i:%d val:%s\n", i, *value)
	}
	elapsed1 := time.Since(start)

	var children2 map[int]*string = map[int]*string{}
	children2[0] = &hello

	start = time.Now()
	for i, val := range children2 {
		value := val
		if val == nil {
			value = &null
		}
		fmt.Printf("i:%d val:%s\n", i, *value)
	}
	elapsed2 := time.Since(start)

	fmt.Printf("Children1 %s\n", elapsed1)
	fmt.Printf("Children2 %s\n", elapsed2)
}

func TestStoreAdd(t *testing.T) {
	tests := []struct {
		id       string
		entries  []storeTestEntry
		expected string
	}{
		{
			"all static",
			[]storeTestEntry{
				{1, "/gopher/bumper.png", "1", 0},
				{2, "/gopher/bumper192x108.png", "2", 0},
				{3, "/gopher/doc.png", "3", 0},
				{4, "/gopher/bumper320x180.png", "4", 0},
				{5, "/gopher/docpage.png", "5", 0},
				{6, "/gopher/doc.png", "6", 0},
				{7, "/gopher/doc", "7", 0},
			},
			`{id: -1, static: true, key: , regex: <nil>, data: <nil>, order: 0, minOrder: 0, pindex: -1, pnames: []}
    {id: -1, static: true, key: /gopher/, regex: <nil>, data: <nil>, order: 1, minOrder: 1, pindex: -1, pnames: []}
        {id: -1, static: true, key: bumper, regex: <nil>, data: <nil>, order: 1, minOrder: 1, pindex: -1, pnames: []}
            {id: 1, static: true, key: .png, regex: <nil>, data: 1, order: 1, minOrder: 1, pindex: -1, pnames: []}
            {id: 2, static: true, key: 192x108.png, regex: <nil>, data: 2, order: 2, minOrder: 2, pindex: -1, pnames: []}
            {id: 4, static: true, key: 320x180.png, regex: <nil>, data: 4, order: 4, minOrder: 4, pindex: -1, pnames: []}
        {id: 7, static: true, key: doc, regex: <nil>, data: 7, order: 7, minOrder: 3, pindex: -1, pnames: []}
            {id: 3, static: true, key: .png, regex: <nil>, data: 3, order: 3, minOrder: 3, pindex: -1, pnames: []}
            {id: 5, static: true, key: page.png, regex: <nil>, data: 5, order: 5, minOrder: 5, pindex: -1, pnames: []}
`,
		},
		{
			"parametric",
			[]storeTestEntry{
				{1, "/users/<id>", "11", 1},
				{2, "/users/<id>/profile", "12", 1},
				{3, "/users/<id>/<accnt:\\d+>/address", "13", 2},
				{4, "/users/<id>/age", "14", 1},
				{5, "/users/<id>/<accnt:\\d+>", "15", 2},
			},
			`{id: -1, static: true, key: , regex: <nil>, data: <nil>, order: 0, minOrder: 0, pindex: -1, pnames: []}
    {id: -1, static: true, key: /users/, regex: <nil>, data: <nil>, order: 0, minOrder: 1, pindex: -1, pnames: []}
        {id: 1, static: false, key: <id>, regex: <nil>, data: 11, order: 1, minOrder: 1, pindex: 0, pnames: [id]}
            {id: -1, static: true, key: /, regex: <nil>, data: <nil>, order: 2, minOrder: 2, pindex: 0, pnames: [id]}
                {id: 4, static: true, key: age, regex: <nil>, data: 14, order: 4, minOrder: 4, pindex: 0, pnames: [id]}
                {id: 2, static: true, key: profile, regex: <nil>, data: 12, order: 2, minOrder: 2, pindex: 0, pnames: [id]}
                {id: 5, static: false, key: <accnt:\d+>, regex: ^\d+, data: 15, order: 5, minOrder: 3, pindex: 1, pnames: [id accnt]}
                    {id: 3, static: true, key: /address, regex: <nil>, data: 13, order: 3, minOrder: 3, pindex: 1, pnames: [id accnt]}
`,
		},
		{
			"corner cases",
			[]storeTestEntry{
				{1, "/users/<id>/test/<name>", "101", 2},
				{2, "/users/abc/<id>/<name>", "102", 2},
				{3, "", "103", 0},
			},
			`{id: 3, static: true, key: , regex: <nil>, data: 103, order: 3, minOrder: 0, pindex: -1, pnames: []}
    {id: -1, static: true, key: /users/, regex: <nil>, data: <nil>, order: 0, minOrder: 1, pindex: -1, pnames: []}
        {id: -1, static: true, key: abc/, regex: <nil>, data: <nil>, order: 0, minOrder: 2, pindex: -1, pnames: []}
            {id: -1, static: false, key: <id>, regex: <nil>, data: <nil>, order: 0, minOrder: 2, pindex: 0, pnames: [id]}
                {id: -1, static: true, key: /, regex: <nil>, data: <nil>, order: 0, minOrder: 2, pindex: 0, pnames: [id]}
                    {id: 2, static: false, key: <name>, regex: <nil>, data: 102, order: 2, minOrder: 2, pindex: 1, pnames: [id name]}
        {id: -1, static: false, key: <id>, regex: <nil>, data: <nil>, order: 0, minOrder: 1, pindex: 0, pnames: [id]}
            {id: -1, static: true, key: /test/, regex: <nil>, data: <nil>, order: 0, minOrder: 1, pindex: 0, pnames: [id]}
                {id: 1, static: false, key: <name>, regex: <nil>, data: 101, order: 1, minOrder: 1, pindex: 1, pnames: [id name]}
`,
		},
	}
	for _, test := range tests {
		h := NewURLStoreTree()
		for _, entry := range test.entries {
			n := h.Add(entry.id, entry.key, entry.data)
			assert.Equal(t, entry.params, n, test.id+" > "+entry.key+" > param count =")
		}
		assert.Equal(t, test.expected, h.String(), test.id+" > store.String() =")
	}
}

func TestStoreGetAndDelete(t *testing.T) {
	pairs := []struct {
		id         int
		key, value string
	}{
		{1, "/gopher/bumper.png", "1"},
		{2, "/gopher/bumper192x108.png", "2"},
		{3, "/gopher/doc.png", "3"},
		{4, "/gopher/bumper320x180.png", "4"},
		{5, "/gopher/docpage.png", "5"},
		{6, "/gopher/doc.png", "6"},
		{7, "/gopher/doc", "7"},
		{8, "/users/<id>", "8"},
		{9, "/users/<id>/profile", "9"},
		{10, "/users/<id>/<accnt:\\d+>/address", "10"},
		{11, "/users/<id>/age", "11"},
		{12, "/users/<id>/<accnt:\\d+>", "12"},
		{13, "/users/<id>/test/<name>", "13"},
		{14, "/users/abc/<id>/<name>", "14"},
		{15, "", "15"},
		{16, "/all/<:.*>", "16"},
	}
	h := NewURLStoreTree()
	maxParams := 0
	for _, pair := range pairs {
		fmt.Printf("=== ID (by order): %d\n", pair.id)
		n := h.Add(pair.id, pair.key, pair.value)
		if n > maxParams {
			maxParams = n
		}
	}
	assert.Equal(t, 2, maxParams, "param count = ")

	tests := []struct {
		key    string
		value  interface{}
		params string
	}{
		{"/gopher/bumper.png", "1", ""},
		{"/gopher/bumper192x108.png", "2", ""},
		{"/gopher/doc.png", "3", ""},
		{"/gopher/bumper320x180.png", "4", ""},
		{"/gopher/docpage.png", "5", ""},
		{"/gopher/doc.png", "3", ""},
		{"/gopher/doc", "7", ""},
		{"/users/abc", "8", "id:abc,"},
		{"/users/abc/profile", "9", "id:abc,"},
		{"/users/abc/123/address", "10", "id:abc,accnt:123,"},
		{"/users/abcd/age", "11", "id:abcd,"},
		{"/users/abc/123", "12", "id:abc,accnt:123,"},
		{"/users/abc/test/123", "13", "id:abc,name:123,"},
		{"/users/abc/xyz/123", "14", "id:xyz,name:123,"},
		{"", "15", ""},
		{"/g", nil, ""},
		{"/all", nil, ""},
		{"/all/", "16", ":,"},
		{"/all/abc", "16", ":abc,"},
		{"/users/abc/xyz", nil, ""},
	}
	pvalues := make([]string, maxParams)
	for _, test := range tests {
		data, pnames := h.Get(test.key, pvalues)
		assert.Equal(t, test.value, data, "store.Get("+test.key+") =")
		params := ""
		if len(pnames) > 0 {
			for i, name := range pnames {
				params += fmt.Sprintf("%v:%v,", name, pvalues[i])
			}
		}
		assert.Equal(t, test.params, params, "store.Get("+test.key+").params =")
	}

	t.Logf("\nBEFORE DELETE:\n%s", h.String())
	assert.Equal(t, 16, h.Count())

	t.Run("delete static", func(t *testing.T) {
		deleted := h.Delete(7)
		assert.Equal(t, true, deleted)

		data, _ := h.Get("/gopher/doc", pvalues)
		assert.Equal(t, nil, data)

		// deleted = h.Delete(3)
		// assert.Equal(t, true, deleted)

		// data, _ = h.Get("/gopher/doc.png", pvalues)
		// assert.Equal(t, nil, data)

		// deleted = h.Delete(5)
		// assert.Equal(t, true, deleted)

		// data, _ = h.Get("/gopher/docpage.png", pvalues)
		// assert.Equal(t, nil, data)
	})

	t.Run("delete param", func(t *testing.T) {

		deleted := h.Delete(13)
		assert.Equal(t, true, deleted)

		data, _ := h.Get("/users/<id>/test/<name>", pvalues)
		assert.Equal(t, nil, data)

		deleted = h.Delete(9)
		assert.Equal(t, true, deleted)

		data, _ = h.Get("/users/44/profile", pvalues)
		assert.Equal(t, nil, data)

		// deleted = h.Delete(12)
		// assert.Equal(t, true, deleted)

		deleted = h.Delete(10)
		assert.Equal(t, true, deleted)

		// === WEIRD BUG: IF we delete 11, 12 is also deleted. ???
		deleted = h.Delete(11)
		assert.Equal(t, true, deleted)
		// === END WEIRD BUG

		data, _ = h.Get("/users/23/35", pvalues)
		t.Logf("### DATA: %+v\n", data)
		assert.NotEqual(t, nil, data)

		assert.Equal(t, 11, h.Count())

	})

	t.Logf("\nAFTER DELETE:\n%s", h.String())

	// deleted := h.Delete(7)
	// t.Logf("_found_deleted_: %t", deleted)

	// deleted = h.Delete(20)
	// t.Logf("_found_deleted_ 20: %t", deleted)
}

func TestStoreDelete(t *testing.T) {
	t.Run("GIVEN 4-level tree", func(t *testing.T) {

	})
}
