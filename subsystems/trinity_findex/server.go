package main

import (
	"context"

	"github.com/clownpriest/trinity/api/trinity"
)

type FIndexServer struct {
}

func (s *FIndexServer) GetDocMap(ctx context.Context, req *trinity.DocMapRequest) (*trinity.ForwardMap, error) {
	return &trinity.ForwardMap{}, nil
}
