package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersion(t *testing.T) {
	version, err := ParseVersion("v1")
	if err != nil {
		t.Errorf("Failure.")
	}

	assert.Equal(t, "1.0.0", version.String())
	assert.Equal(t, "v1", version.Original())
}
