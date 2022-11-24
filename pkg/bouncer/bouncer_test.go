package bouncer

import (
	"bytes"
	"testing"
)

func TestBouncerHappyPath(t *testing.T) {
	b := bytes.NewBufferString("")

	bouncer := New()
	bouncer.Out = b
	bouncer.Err = b
}
