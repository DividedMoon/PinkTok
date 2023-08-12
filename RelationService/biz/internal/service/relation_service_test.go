package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"relation_service/biz/model/client"
	"testing"
)

func TestSubmitFollowRelationAction(t *testing.T) {
	var (
		ctx = context.Background()
		req = &client.RelationActionReq{
			UserId:     1,
			ToUserId:   1001,
			ActionType: 1,
		}
	)

	err := SubmitFollowRelationAction(ctx, req)
	assert.Assert(t, err == nil)
}
