package interactor

import (
	"context"
	"gomulticache/internal/domain/cache"
	"gomulticache/internal/domain/model"
	"gomulticache/internal/usecase/port"
)

var _ port.SampleDelUsecase = (*SampleDel)(nil)

type SampleDel struct {
	cache cache.Cache[model.Sample]
}

func NewSampleDel(
	cache cache.Cache[model.Sample],
) *SampleDel {
	return &SampleDel{
		cache: cache,
	}
}

func (sd *SampleDel) Execute(ctx context.Context, input port.SampleDelInput) (port.SampleDelOutput, error) {
	if err := sd.cache.Del(ctx, input.ID.String()); err != nil {
		return port.SampleDelOutput{}, err
	}

	return port.SampleDelOutput{}, nil
}
