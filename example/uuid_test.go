package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := ktools.UUID

	uuidStr := uuid.NewV1(false)
	fmt.Println("Get:", uuidStr)

	uuidStr = uuid.NewV1(true)
	fmt.Println("Get:", uuidStr)

	uuidStr = uuid.NewV4(false)
	fmt.Println("GetV4:", uuidStr)

	uuidStr = uuid.NewV4(true)
	fmt.Println("GetV4:", uuidStr)

	uuidStr = uuid.New()
	fmt.Println("Get:", uuidStr)
}
