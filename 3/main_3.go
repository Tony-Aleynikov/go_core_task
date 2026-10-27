package main

import "maps"

type StringIntMap struct {
	data map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	if s.data == nil {
		s.data = make(map[string]int)
	}
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() *StringIntMap {
	newMap := StringIntMap{data: make(map[string]int, len(s.data))}
	maps.Copy(newMap.data, s.data)
	return &newMap
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.data[key]
	return ok
}
