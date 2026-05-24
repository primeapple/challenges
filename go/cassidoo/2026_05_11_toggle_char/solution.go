package toggleChar

import (
	"io"
	"os/exec"
	"strings"
)

func ToggleChar(input string) (string, error) {
	return executeInNvim(input, "+norm g~G")
}

func AlternateChar(input string) (string, error) {
	return executeInNvim(input, "+execute 'normal! qqgUllgull@qq' | execute 'normal! 999@q'")
}

func executeInNvim(input, command string) (string, error) {
	nvimCmdParts := []string{
		"-",
		"--clean",
		"--headless",
		"--cmd",
		"set nonumber",
		command,
		"+%print",
		"+qa!",
	}

	nvimCmd := exec.Command("nvim", nvimCmdParts...)
	r, w := io.Pipe()
	nvimCmd.Stdin = r

	go func() {
		defer w.Close()
		_, writeErr := w.Write([]byte(input))
		if writeErr != nil {
			// If we can't write, just close it; the error will be caught by nvimCmd.Output() or elsewhere if needed
		}
	}()

	b, err := nvimCmd.Output()
	return strings.TrimRight(string(b), "\n"), err
}
