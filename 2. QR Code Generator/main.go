package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/skip2/go-qrcode"
)

func main() {
	url := flag.String("url", "", "URL")
	filename := flag.String("filename", "", "File name")
	format := flag.String("format", "png", "File extension")
	fileSize := flag.Int("fileSize", 256, "File size")
	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		return
	}
	if *filename == "" {
		*filename = randomFilename()
	}
	if *fileSize <= 0 {
		fmt.Println("File size must be positive")
		return
	}
	if *format != "" {
		if *format != "png" && *format != "jpg" && *format != "webp" {
			fmt.Println("Format must be \"png\" \"jpg\" \"webp\"")
			return
		}
	}

	err := qrcode.WriteFile(
		*url,
		qrcode.Medium,
		*fileSize,
		fmt.Sprintf("%s.%s", *filename, *format),
	)
	if err != nil {
		fmt.Printf("failed to generate QR code: %v\n", err)
		return
	}

	fmt.Println("QR Code generated successfully")
}

// randomFilename generates a random filename.
func randomFilename() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
