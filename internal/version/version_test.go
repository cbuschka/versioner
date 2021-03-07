package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersionWithoutPrefix(t *testing.T) {
	version, err := ParseVersion("1.0")
	if err != nil {
		t.Errorf("Failure.")
	}

	assert.Equal(t, "1.0.0", version.String())
	assert.Equal(t, "1.0", version.Original())
}

func TestParseVersionWithLeadingV(t *testing.T) {
	version, err := ParseVersion("v1")
	if err != nil {
		t.Errorf("Failure.")
	}

	assert.Equal(t, "1.0.0", version.String())
	assert.Equal(t, "v1", version.Original())
}

func TestNext(t *testing.T) {
	version, err := ParseVersion("v1")
	if err != nil {
		t.Errorf("Failure.")
	}

	next := version.Next()

	assert.Equal(t, "1.0.1", next.String())
	assert.Equal(t, "v1.0.1", next.Original())
}

func TestNextFromPrerelease(t *testing.T) {
	version, err := ParseVersion("v1-beta")
	if err != nil {
		t.Errorf("Failure.")
	}

	next := version.Next()

	assert.Equal(t, "1.0.0", next.String())
	assert.Equal(t, "v1.0.0", next.Original())
}
