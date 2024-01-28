package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenameTransformFunc(filename)
	fmt.Println("new filename: ", newFilename)
	return nil
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	return hex.EncodeToString(hash[:])
}

func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + "_" + filename
	}
}

func main() {
	s := &Server{
		filenameTransformFunc: prefixFilename("JK"),
	}

	s.handleRequest("cool_pic.jpg")
}
