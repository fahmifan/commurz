package pkgutil

import (
	"encoding/json"
	"fmt"

	"github.com/fahmifan/ulids"
	"github.com/oklog/ulid/v2"
	"github.com/samber/lo"
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

func MustParseULID(s string) ulids.ULID {
	return ulids.ULID{ULID: ulid.MustParse(s)}
}

func StringULIDs(ids []ulids.ULID) []string {
	return lo.Map(ids, func(id ulids.ULID, index int) string {
		return id.String()
	})
}
