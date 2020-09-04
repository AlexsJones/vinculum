package impl

import (
	"context"
	"github.com/AlexsJones/vinculum/pkg/proto"
)

type HandShakeImpl struct{}

func (HandShakeImpl) Commence(ctx context.Context, request *proto.HandShakeRequest) (*proto.HandShakeResponse, error) {
	panic("implement me")
}
