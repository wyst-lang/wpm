package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/term"
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

func ReadPassword(buff *string) error {
	safeChar := "*"
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	for true {
		b := make([]byte, 1)
		_, err = os.Stdin.Read(b)
		if err != nil {
			return err
		}
		chr := string(b[0])
		if chr == "\r" || chr == "\n" {
			fmt.Printf("\n\r")
			return nil
		} else if chr == "\b" || b[0] == 127 {
			if len(*buff) > 0 {
				fmt.Print("\033[1D \033[1D")
				*buff = string(*buff)[:len(*buff)-1]
			}
		} else {
			*buff += chr
			if safeChar == "off" {
				fmt.Print(chr)
			} else {
				fmt.Print(safeChar)
			}
		}

	}
	return nil
}

func clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
