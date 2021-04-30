package main

import "fmt"

type IntSet struct {
  words []unit64
}

// ��ʾ�Ƿ���ڷǸ���x
func (s *Intset) Has(x int) bool {
  word, bit := x/64, unit(x%64)
  return word < len(s.words) && s.words[word]&(1<<bit) !=0
}
