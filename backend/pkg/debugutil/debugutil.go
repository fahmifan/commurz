package debugutil

import (
	"encoding/json"
	"fmt"
)

// PrettyJSON returns a pretty-printed JSON string of v.
func PrettyJSON(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

// PrintlnDebug add "DEBUG >>> " prefix to the output.
func PrintlnDebug(v ...any) {
	vv := append([]any{"DEBUG >>>"}, v...)
	fmt.Println(vv...)
}
