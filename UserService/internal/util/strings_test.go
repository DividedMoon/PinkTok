package utils

import (
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestCameCaseToUnderscore(t *testing.T) {
	underscore := CameCaseToUnderscore("FollowCount")
	assert.Assert(t, underscore == "follow_count")
}
