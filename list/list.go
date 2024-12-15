// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lists provides an abstract List interface.
//
// In computer science, a list or sequence is an abstract data type that represents an ordered sequence of values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a finite sequence; the (potentially) infinite analog of a list is a stream.  Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package list

import ()

// List interface that all lists implement
type List[T comparable] interface {
	Append(items ...T)
	Prepend(items ...T)
	Get(index int) (T, bool)
	Remove(index int) bool
	Contains(item T) bool
	GetAllNode() []T
	GetSize() int
	IsEmpty() bool
	Sort(compareFunction Comparator[T])
	Swap(item1, item2 int)
	Insert(index int, items ...T) bool
	UpdateNodeValue(index int, item T) bool
	Set(index int, value T)
  String() string
  Clear()
}
