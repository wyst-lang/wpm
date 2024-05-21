package main

import (
	"fmt"
	"strings"
)

type ProgressBar struct {
	total       int
	length      int
	last_suffix int
	enabled     bool
}

func (prb *ProgressBar) change(amount int, prefix, suffix string) {
	if prb.enabled {
		percent := float64(amount) / float64(prb.total)
		filledLength := int(float64(prb.length) * percent)
		fill := "█"
		end := "█"
		if amount == prb.total {
			end = "█"
		}
		for len(suffix) < prb.last_suffix {
			suffix += " "
		}
		bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (prb.length-filledLength))
		fmt.Printf("\r%s [%s] %s", prefix, bar, suffix)
		// if amount == prb.total {
		// 	fmt.Println()
		// }
		prb.last_suffix = len(suffix) + 3
	}
}

func (prb ProgressBar) clean() {
	if prb.enabled {
		fmt.Printf("\r" + strings.Repeat(" ", prb.last_suffix+prb.length+2) + "\r")
	}
}
