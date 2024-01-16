package main

import "os"

type pgnum uint64

type page struct {
	num  pgnum
	data []byte
}

type dal struct {
	file *os.File
	pageSize int
	*meta
	*freelist
}

type freelist struct {
	maxPage       pgnum
	releasedPages []pgnum
}

type Item struct {
	key []byte
	value []byte
}

type Node struct {
	* dal
	pageNum pgnum
	items []*Item
	childNodes []pgnum
}
