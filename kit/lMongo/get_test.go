package lMongo

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	db := Get(&Config{})
	fmt.Println(db)
}
