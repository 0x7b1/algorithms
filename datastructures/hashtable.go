package main

import (
	"fmt"
	"hash/fnv"
)

type kv struct {
	Key, Value string
}

type Table struct {
	m     int
	table [][]kv
}

func NewTable(m int) *Table {
	return &Table{
		m:     m,
		table: make([][]kv, m),
	}
}

func (t *Table) hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32()) % t.m
}

func (t *Table) Insert(key, value string) {
	i := t.hash(key)

	for j, kv := range t.table[i] {
		if key == kv.Key {
			t.table[i][j].Value = value
			return
		}
	}

	t.table[i] = append(t.table[i], kv{
		Key:   key,
		Value: value,
	})
}

func (t *Table) Lookup(key string) (string, bool) {
	i := t.hash(key)
	for _, kv := range t.table[i] {
		if kv.Key == key {
			return kv.Value, true
		}
	}

	return "", false
}

func main() {
	t := NewTable(5)
	t.Insert("10", "20")
	t.Insert("10", "21")
	t.Insert("10", "22")
	t.Insert("11", "33")

	fmt.Println(t.Lookup("10"))
	fmt.Println(t.Lookup("12"))
}
