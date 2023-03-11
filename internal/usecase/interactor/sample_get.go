package interactor

import (
	"context"
	"gomulticache/internal/domain/cache"
	"gomulticache/internal/domain/model"
	"gomulticache/internal/usecase/port"
)

var _ port.SampleGetUsecase = (*SampleGet)(nil)

type SampleGet struct {
	cache cache.Cache[model.Sample]
}

func NewSampleGet(
	cache cache.Cache[model.Sample],
) *SampleGet {
	return &SampleGet{
		cache: cache,
	}
}

func (sg *SampleGet) Execute(ctx context.Context, input port.SampleGetInput) (port.SampleGetOutput, error) {
	val, err := sg.cache.Get(ctx, input.ID.String())
	if err != nil {
		return port.SampleGetOutput{}, err
	}

	return port.SampleGetOutput{
		Sample: val,
	}, nil
}
