package TEST

import (
	"fmt"
	"go-com/common/mypath"
	"testing"
)

func TestOLE(t *testing.T) {
	fmt.Println("hello")
	fmt.Println(mypath.GetCurrentPath())
}
