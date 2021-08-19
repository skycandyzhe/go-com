package stringUtil

import (
	"fmt"
	"testing"
)

func TestStrConnect(t *testing.T) {
	fmt.Println(Str_LinkBySpecialChar('_', "___", "asdasda"))
	fmt.Println(Str_LinkBySpecialChar('_', "Dasda", "dvd23$#@%@"))
	fmt.Println(Str_LinkBySpecialChar('_', "", ""))
	fmt.Println(Str_LinkBySpecialChar('_', "1231"))
	fmt.Println(Str_LinkBySpecialChar('_', "___"))
	fmt.Println(Str_LinkBySpecialChar('_', "___", "bbb"))
	fmt.Println(Str_LinkBySpecialChar('_', "___", "bbb", "das", "dasda", "+_+", "__", "__", "dadsa"))
	fmt.Println(Str_Link("___", "bbb", "das", "dasda", "+_+", "__", "__", "dadsa"))
	// assert.Equal(1, 1)

}
