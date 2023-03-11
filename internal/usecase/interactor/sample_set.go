package interactor

import (
	"context"
	"gomulticache/internal/domain/cache"
	"gomulticache/internal/domain/model"
	"gomulticache/internal/domain/model/sample"
	"gomulticache/internal/usecase/port"

	"github.com/google/uuid"
)

var _ port.SampleSetUsecase = (*SampleSet)(nil)

type SampleSet struct {
	cache cache.Cache[model.Sample]
}

func NewSampleSet(
	cache cache.Cache[model.Sample],
) *SampleSet {
	return &SampleSet{
		cache: cache,
	}
}

func (ss *SampleSet) Execute(ctx context.Context, input port.SampleSetInput) (port.SampleSetOutput, error) {
	id := uuid.NewString()

	smp := model.Sample{
		ID:   sample.ID(id),
		Name: input.Name,
	}

	if err := ss.cache.Set(ctx, id, smp); err != nil {
		return port.SampleSetOutput{}, err
	}

	return port.SampleSetOutput{
		Sample: smp,
	}, nil
}
