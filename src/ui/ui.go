package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/proudynyu/pkgrun/src/file"
)

func execTerminal() func() {
	switch runtime.GOOS {
	case "linux":
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "-echo").Run()
		return func() {
			exec.Command("stty", "-F", "/dev/tty", "-cbreak", "echo").Run()
		}

	case "darwin":
		cmd := exec.Command("sh", "-c", "stty cbreak -echo < /dev/tty")
		cmd.Stdin = os.Stdin
		cmd.Run()
		return func() {
			cmd := exec.Command("sh", "-c", "stty -cbreak echo < /dev/tty")
			cmd.Stdin = os.Stdin
			cmd.Run()
		}

	case "windows":
		exec.Command("cmd", "/c", "mode con: cp select=437").Run()
		return func() { }

	default:
		return func() {}
	}
}

func Selector(selected int, keys []string, tty *os.File) {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Select a script (↑↓ to move, Enter to confirm):\n")

		for i, key := range keys {
			if i == selected {
				fmt.Fprintf(tty, " \033[32m|> %s\033[0m\n", key)
			} else {
				fmt.Printf(" %s \n", key)
			}
		}
}

func BuildInteractiveCmdChoose(pkgJson *file.PackageFormat) string {
	scripts := pkgJson.Scripts

	keys := []string{}
	for key := range scripts {
		if key != "" {
			keys = append(keys, key)
		}
	}
	if len(keys) == 0 {
		return ""
	}

	clean := execTerminal()
	defer clean()

	reader := bufio.NewReader(os.Stdin)
	selected := 0
	size := len(keys)

	loop:
	for {
		Selector(selected, keys, os.Stdin)

		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		switch b {
			case 27: 
				b2, _ := reader.ReadByte()
				if b2 != '[' && b2 != 'O' {
					continue loop
				}
				arrow, _ := reader.ReadByte()
				switch arrow {
				case 'A':
					if selected > 0 {
						selected--
					}
				case 'B':
					if selected < size - 1 {
						selected++
					}
				}
			case 13, 10:
				break loop
		}
	}

	return keys[selected]
}
