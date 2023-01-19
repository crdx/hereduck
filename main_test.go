package hereduck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD(t *testing.T) {
	var doc string

	// Single line
	doc = D(`Hello world`)
	assert.Equal(t, doc, "Hello world")

	// Single line, broken up (tabs)
	doc = D(`
		Hello world
	`)
	assert.Equal(t, doc, "Hello world\n")

	// Multiline, heavy indent (tabs)
	doc = D(`
					Hello
					world
	`)
	assert.Equal(t, doc, "Hello\nworld\n")

	// Multiline, with a blank line (tabs)
	doc = D(`
					Hello

					world
	`)
	assert.Equal(t, doc, "Hello\n\nworld\n")

	// Multiline, variable indent (tabs and spaces)
	doc = D(`
					Hello
					    world
	`)
	assert.Equal(t, doc, "Hello\n    world\n")

	// Multiline, variable indent (spaces), trailing
	doc = D(`
                    Hello
                      world
                `)
	assert.Equal(t, doc, "Hello\n  world\n")

	// Multiline, variable indent (spaces), no trailing
	doc = D(`
                    Hello
                      world`)
	assert.Equal(t, doc, "Hello\n  world")
}

func TestDf(t *testing.T) {
	var doc string

	doc = Df(`
		Hello %s
	`, "world")
	assert.Equal(t, doc, "Hello world\n")
}
