package main

import (
	"fmt"
	"strings"
)

func searchString(kalimat string, keyword string) {
	kataToLower := strings.ToLower(kalimat)
	keywordToLower := strings.ToLower(keyword)

	if strings.Contains(kataToLower, keywordToLower) {
		fmt.Println("kata ditemukan")
	} else {
		fmt.Println("kata tidak ditemukan")
	}
}
