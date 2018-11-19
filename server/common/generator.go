package common

import (
	"encoding/base64"
	"math/rand"
	"strconv"
)

// GenerateIPAddress generates a fake IP address
func GenerateIPAddress() (ipAddress string) {
	ipAddress = strconv.Itoa(rand.Intn(256)) + "." + strconv.Itoa(rand.Intn(256)) + "." + strconv.Itoa(rand.Intn(256)) + "." + strconv.Itoa(rand.Intn(256))
	return
}

// GenerateRandomHash generates a random string of length n
func GenerateRandomHash(n int) string {
	buff := make([]byte, n)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	// Base 64 can be longer than len
	return str[:n]
}
