package controller

import (
	"encoding/json"
	"gomulticache/internal/domain/model/sample"
	"gomulticache/internal/usecase/port"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const SamplePath = "samples"

type Sample struct {
	set port.SampleSetUsecase
	get port.SampleGetUsecase
	del port.SampleDelUsecase
}

func NewSample(
	set port.SampleSetUsecase,
	get port.SampleGetUsecase,
	del port.SampleDelUsecase,
) *Sample {
	return &Sample{
		set: set,
		get: get,
		del: del,
	}
}

func (s *Sample) Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)

	switch r.Method {
	case http.MethodGet:
		s.getHandle(w, r)
	case http.MethodPost:
		s.postHandle(w, r)
	case http.MethodDelete:
		s.deleteHandle(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Sample) getHandle(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/"+SamplePath)
	_, id := filepath.Split(sub)
	if id == "" {
		log.Printf("Failed to get sample: %s", r.URL.Path)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input := port.SampleGetInput{
		ID: sample.ID(id),
	}

	output, err := s.get.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Failed to get sample: %s", err)

		w.WriteHeader(http.StatusNotFound)

		return
	}

	type Res struct {
		Sample struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"sample"`
	}

	res := &Res{
		Sample: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   output.Sample.ID.String(),
			Name: output.Sample.Name.String(),
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (s *Sample) postHandle(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		Name string `json:"name"`
	}

	var req Req

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	input := port.SampleSetInput{
		Name: sample.Name(req.Name),
	}

	output, err := s.set.Execute(r.Context(), input)
	if err != nil {
		log.Printf("Failed to set sample: %s", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	type Res struct {
		Sample struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"sample"`
	}

	res := &Res{
		Sample: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   output.Sample.ID.String(),
			Name: output.Sample.Name.String(),
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (s *Sample) deleteHandle(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/"+SamplePath)
	_, id := filepath.Split(sub)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	input := port.SampleDelInput{
		ID: sample.ID(id),
	}

	if _, err := s.del.Execute(r.Context(), input); err != nil {
		log.Printf("Failed to del sample: %s", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("OK"))

	w.Header().Set("Content-Type", "application/json")
}
