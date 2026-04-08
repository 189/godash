package godash

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKeys(t *testing.T) {
	is := assert.New(t)

	data := make(map[string]string, 0)
	data["Name"] = "Tom"
	data["Country"] = "USA"

	rst := GetKeys(data)
	sort.Slice(rst, func(i, j int) bool {
		return strings.Compare(rst[i], rst[j]) > 0
	})
	expected := []string{"Name", "Country"}

	if !is.Equal(rst, expected) {
		t.Fatalf("expect %v, but got %v", expected, rst)
	}
}

func TestGetValues(t *testing.T) {
	is := assert.New(t)

	data := make(map[string]string, 0)
	data["Name"] = "Tom"
	data["Sex"] = "Boy"

	rst := GetValues(data)
	expected := []string{"Tom", "Boy"}
	sort.Slice(rst, func(i, j int) bool {
		return strings.Compare(rst[i], rst[j]) > 0
	})

	if !is.Equal(rst, expected) {
		t.Fatalf("expect %v, but got %v", expected, rst)
	}
}
