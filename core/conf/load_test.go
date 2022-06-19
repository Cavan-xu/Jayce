package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_FromJson(t *testing.T) {
	content := `{
	"a": "foo",
	"b": 888,
	"c": ["asd", "asd"],
	"d": {
		"e": 9
	}
}
`
	var s struct {
		A string   `json:"a"`
		B int      `json:"b"`
		C []string `json:"c"`
		d struct {
			E int `json:"e"`
		}
	}

	err := loadFromJson([]byte(content), &s)
	assert.Nil(t, err)
	assert.Equal(t, "foo", s.A)
	assert.Equal(t, []string{"asd", "asd"}, s.C)
}
