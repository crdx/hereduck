package hereduck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD(t *testing.T) {
	var doc string

	// Single line
	doc = D(`Hello world`)
	assert.Equal(t, "Hello world", doc)

	// Single line, broken up (tabs)
	doc = D(`
		Hello world
	`)
	assert.Equal(t, "Hello world\n", doc)

	// Multiline, heavy indent (tabs)
	doc = D(`
					Hello
					world
	`)
	assert.Equal(t, "Hello\nworld\n", doc)

	// Multiline, with a blank line (tabs)
	doc = D(`
					Hello

					world
	`)
	assert.Equal(t, "Hello\n\nworld\n", doc)

	// Multiline, variable indent (tabs and spaces)
	doc = D(`
					Hello
					    world
	`)
	assert.Equal(t, "Hello\n    world\n", doc)

	// Multiline, variable indent (spaces), trailing
	doc = D(`
                    Hello
                      world
                `)
	assert.Equal(t, "Hello\n  world\n", doc)

	// Multiline, variable indent (spaces), no trailing
	doc = D(`
                    Hello
                      world`)
	assert.Equal(t, "Hello\n  world", doc)
}

func TestDf(t *testing.T) {
	var doc string

	doc = Df(`
		Hello %s
	`, "world")
	assert.Equal(t, "Hello world\n", doc)
}
