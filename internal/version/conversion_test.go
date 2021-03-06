package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterConvertAndSortAsc(t *testing.T) {
	tags := []string{"invalid", "1.0.0", "0.0.1", "123", "999.8.1"}

	versions, removedTags, err := FilterConvertAndSortAsc(tags)
	if err != nil {
		t.Errorf("Failure.")
	}

	assert.Equal(t, 1, len(removedTags))
	assert.Equal(t, "invalid", removedTags[0])

	assert.Equal(t, 4, len(versions))
	assert.Equal(t, "0.0.1", versions[0].String())
	assert.Equal(t, "1.0.0", versions[1].String())
	assert.Equal(t, "123.0.0", versions[2].String())
	assert.Equal(t, "999.8.1", versions[3].String())
}
