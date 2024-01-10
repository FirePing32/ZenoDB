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
}
