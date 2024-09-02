package main

import (
	"log"
	"strings"
	"time"
)

func Merge_string_alternative(s1, s2 string) string {
	start := time.Now()
	l1 := len(s1)
	l2 := len(s2)

	i := 0
	j := 0

	var sb strings.Builder

	for i < l1 && j < l2 {
		sb.WriteByte(s1[i])
		i++
		sb.WriteByte(s2[j])
		j++
	}

	for i < l1 {
		sb.WriteByte(s1[i])
		j++
	}
	for j < l2 {
		sb.WriteByte(s2[j])
		j++
	}

	log.Printf("Runtime: %v", time.Since(start))
	return sb.String()
}
