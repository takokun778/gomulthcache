package port

import (
	"gomulticache/internal/domain/model"
	"gomulticache/internal/domain/model/sample"
	"gomulticache/internal/usecase"
)

type SampleSetInput struct {
	usecase.Input
	sample.Name
}

type SampleSetOutput struct {
	usecase.Output
	model.Sample
}

type SampleSetUsecase interface {
	usecase.Usecase[SampleSetInput, SampleSetOutput]
}

type SampleGetInput struct {
	usecase.Input
	sample.ID
}

type SampleGetOutput struct {
	usecase.Output
	model.Sample
}

type SampleGetUsecase interface {
	usecase.Usecase[SampleGetInput, SampleGetOutput]
}

type SampleDelInput struct {
	usecase.Input
	sample.ID
}

type SampleDelOutput struct {
	usecase.Output
}

type SampleDelUsecase interface {
	usecase.Usecase[SampleDelInput, SampleDelOutput]
}
