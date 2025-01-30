package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://script.google.com/macros/s/AKfycbxHDhTy_UIjL51FrT6E9dKqMb1rYmdy2ZnLrRpubTdhXMgdy-fCKeKY1eSvPJuw_0s/exec"

	data := map[string]string{
		"name":    "Alice",
		"message": "Hello from Go!",
	}

	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response:", resp.Status)
}
