package nm

import (
	"strings"
	"testing"
)

var validLines = []string{
	`           16781312 U _mmap`,
	`           16781312 U _munmap`,
	`           16781312 U _notify_is_valid_token`,
	`           16781312 U _open`,
	` 10d8fe0          8 R $f64.4024000000000000`,
	` 10d8fe8          8 R $f64.403a000000000000`,
	`115d4a0        256 D time.utcLoc`,
	`1091ca0        192 T type:.eq.[2]interface {}`,
	`1060160        160 T type:.eq.[2]runtime.Frame`,
	`1001fa0        256 T type:.eq.[6]internal/cpu.option`,
	`       0          0 _ go.go`,
}

func TestParseLine(t *testing.T) {
	for _, line := range validLines {
		_, err := parseLine(line)
		t.Log(strings.Fields(line))
		if err != nil {
			t.Errorf("%q parsing failed: %v", line, err)
		}
	}
}
