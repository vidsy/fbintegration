package fbintegration

import (
	"fmt"

	facebookLib "github.com/huandu/facebook"
)

const videoType = "VIDEO"

type (
	// Creative comment pending
	Creative struct {
		ID         string `facebook:"id"`
		ObjectType string `facebook:"object_type"`
		PostID     string `facebook:"effective_object_story_id"`
	}
)

// NewCreativeFromResult comment pending
func NewCreativeFromResult(result facebookLib.Result) Creative {
	var creative Creative
	result.DecodeField("", &creative)
	return creative
}

// GenerateParams comments pending
func (c *Creative) GenerateParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=%s", c.ID, "object_id,object_type,effective_object_story_id"))
}

// IsVideo comment pending
func (c *Creative) IsVideo() bool {
	if c.ObjectType == videoType {
		return true
	}
	return false
}
