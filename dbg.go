package dbg

import (
	"fmt"
	"encoding/json"
)

func PrintJson(v interface{})  {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("[error: " + err.Error() + "]")
	}
	fmt.Println(string(j))
}
