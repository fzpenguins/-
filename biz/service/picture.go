package service

import "context"

type PictureService struct {
	ctx context.Context
}

func NewPictureService(ctx context.Context) *PictureService {
	return &PictureService{ctx: ctx}
}
