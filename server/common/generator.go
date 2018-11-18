package common

import (
	"math/rand"
	"strconv"
)

// GenerateIPAddress generates a fake IP address
func GenerateIPAddress() (ipAddress string) {
	ipAddress = strconv.Itoa(rand.Intn(256)) + "." + strconv.Itoa(rand.Intn(256)) + "." + strconv.Itoa(rand.Intn(256)) + "." + strconv.Itoa(rand.Intn(256))
	return
}
