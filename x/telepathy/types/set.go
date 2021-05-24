package types

import "sort"

type Set struct {
	m map[string]bool
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]bool)
	return s
}

func (s *Set) Add(value string) {
	s.m[value] = true
}

func (s *Set) Remove(value string) {
	delete(s.m, value)
}

func (s *Set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

type Struct struct {
	Fields map[string]*Comment
}

// KeyValue is a simple key/value representation of one field of a Struct.
type KeyValue struct {
	Key   string
	Value *Comment
}

// MarshalAmino transforms the Struct to an array of key/value.
func (m Struct) MarshalAmino() ([]KeyValue, error) {
	p := make([]KeyValue, len(m.Fields))
	fieldKeys := make([]string, len(m.Fields))
	i := 0
	for key := range m.Fields {
		fieldKeys[i] = key
		i++
	}
	sort.Stable(sort.StringSlice(fieldKeys))
	for i, key := range fieldKeys {
		p[i] = KeyValue{
			Key:   key,
			Value: m.Fields[key],
		}
	}
	return p, nil
}

// UnmarshalAmino transforms the key/value array to a Struct.
func (m *Struct) UnmarshalAmino(keyValues []KeyValue) error {
	m.Fields = make(map[string]*Comment, len(keyValues))
	for _, p := range keyValues {
		m.Fields[p.Key] = p.Value
	}
	return nil
}

type Fields map[string]*Comment
