package http

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func BasicGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get models: %v\n", err)

		return []byte{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)

		return []byte{}, err
	}

	return data, nil
}

func BaiscPost(url string, contentType string, stringData string) ([]byte, error) {
	postData := strings.NewReader(stringData)
	res, err := http.Post(url, contentType, postData)
	if err != nil {
		fmt.Printf("Failed to send post request: %v\n", err)

		return []byte{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)

		return []byte{}, err
	}

	return data, nil
}

func StreamPost(url, contentType, stringData string, output chan<- []byte) {
	defer close(output)

	postData := strings.NewReader(stringData)
	res, err := http.Post(url, contentType, postData)
	if err != nil {
		fmt.Printf("Failed to send post request: %v\n", err)

		return
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)

	for scanner.Scan() {
		line := scanner.Text()
		output <- []byte(line)
	}

	if err := scanner.Err(); err != nil {
		return
	}

	return
}
