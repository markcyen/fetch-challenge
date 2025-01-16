package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUniqueID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}
