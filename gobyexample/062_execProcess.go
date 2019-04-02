// 62 exec process

package main

import (
	"os/exec"
)

func main() {
	binary, lookErr := exec.LookPath("git")
	if lookErr != nil {
		panic(lookErr)
	}
	//args := []string{"git", "add", "."}
	s := []string{"cmd.exe", "/C", "start", binary}
	cmd := exec.Command(s[0], s[1:]...)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
