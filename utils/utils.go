package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mrinjamul/go-secret/models"
)

// RandomChars generates a random string of length n
// URL with length 5, will give 62⁵ = ~916 Million URLs
// URL with length 6, will give 62⁶ = ~56 Billion URLs
// URL with length 7, will give 62⁷ = ~3500 Billion URLs
func RandomChars(length int) string {
	var chars []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("wrong charset length")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// GenerateShortURL generates a short url
func GenerateHash() string {
	// generate unique short url using url
	return RandomChars(5)
}

// MakeRequest makes a request to a url
func MakeRequest(url string, method string, body []byte) ([]byte, error) {
	responseBody := bytes.NewBuffer(body)
	// Create a new request using http
	req, err := http.NewRequest(method, url, responseBody)
	if err != nil {
		return nil, err
	}
	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetMessage(url string) (models.Message, error) {
	body, err := MakeRequest(url, "GET", nil)
	if err != nil {
		return models.Message{}, err
	}
	var message models.Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		return models.Message{}, err
	}
	return message, nil
}

func AddMessage(username, msg, url string) (models.Message, error) {
	body, err := json.Marshal(models.Message{UserName: username, Message: msg})
	if err != nil {
		return models.Message{}, err
	}
	body, err = MakeRequest(url, "POST", body)
	if err != nil {
		return models.Message{}, err
	}
	var message models.Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		return models.Message{}, err
	}
	return message, nil
}
