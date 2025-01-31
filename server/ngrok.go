package server

import (
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

// startNgrokTunnel starts a TCP tunnel for the given port using ngrok
func startNgrokTunnel(port string) (string, error) {
	cmd := exec.Command("ngrok", "tcp", port)

	// Start ngrok in the background
	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("failed to start ngrok: %v", err)
	}

	// Wait for ngrok to initialize
	time.Sleep(2 * time.Second)

	// Fetch ngrok tunnel status
	statusCmd := exec.Command("curl", "-s", "http://localhost:4040/api/tunnels")
	output, err := statusCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to fetch ngrok status: %v\nOutput: %s", err, output)
	}

	// Extract the public URL
	re := regexp.MustCompile(`"public_url":"tcp://([^"]+)"`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return "", fmt.Errorf("could not determine ngrok tunnel URL")
	}

	return matches[1], nil
}
