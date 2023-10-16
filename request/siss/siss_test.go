package siss

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"testing"
)

func TestRegisterImage(t *testing.T) {
	file, err := os.Open("./sample.jpeg") // Replace with the actual image file path
	if err != nil {
		t.Errorf("Open file error: %v", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		t.Errorf("Image decode error: %v", err)
		return
	}

	// Convert the image to a byte slice
	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, img, nil) // Replace 'jpeg' with your image format
	if err != nil {
		t.Errorf("Image encode error: %v", err)
		return
	}

	// Get the byte slice
	imageBytes := buffer.Bytes()

	id, err := RegisterImage(imageBytes)
	if err != nil {
		t.Errorf("request failed : %v", err)
		return
	}

	fmt.Println("Success! id: ", id)
}
