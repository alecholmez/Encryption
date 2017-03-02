package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/richard-lyman/lithcrypt"
)

// Response object
type Response struct {
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

// EncryptHandler - will be called when a user pastes text in an input to encrypt
func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	var req struct {
		Text string `json:"text"`
		Key  string `json:"key"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
	}

	encrypted, err := lithcrypt.Encrypt([]byte(req.Key), []byte(req.Text))
	if err != nil {
		log.Println(err)
	}

	res := Response{
		Text:      byteSliceToBase64(encrypted),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	bytes, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Write(bytes)

}

// DecryptHandler - will be called when a user pastes text in an input to decrypt
func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	// Get the request body
	var req struct {
		Text string `json:"text"`
		Key  string `json:"key"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
	}

	b := decodeBase64([]byte(req.Text))
	original, err := lithcrypt.Decrypt([]byte(req.Key), []byte(b))
	if err != nil {
		log.Println("Failed to decrypt: ", err)
	}
	res := Response{
		Text:      string(original),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	bytes, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Write(bytes)
}

func byteSliceToBase64(b []byte) string {
	result := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(result, b)
	return string(result)
}

func decodeBase64(b []byte) string {
	result := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	l, err := base64.StdEncoding.Decode(result, b)
	if err != nil {
		log.Println(err)
	}
	return string(result[:l])
}

// ServeDocs returns a static html file to the ResponseWriter
// that contains the documentation for the microservice
func ServeDocs(w http.ResponseWriter, r *http.Request) {
	// Create the file server
	path := os.Getenv("DOCS")
	fs := http.FileServer(http.Dir(path))

	fs.ServeHTTP(w, r)
}
