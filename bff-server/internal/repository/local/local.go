package local

import (
	"bff-server/internal/entity"
	"bff-server/internal/service"
	"context"

	jsoninter "github.com/json-iterator/go"
)

type repository struct {
	j jsoninter.API
}

func (s *repository) GetChannels(context.Context) ([]entity.Channel, error) {
	return channels, nil
}

func (s *repository) Summarize(context.Context, string) (string, error) {
	return "summarize example", nil
}
func New() service.Repository {
	return &repository{
		j: jsoninter.ConfigCompatibleWithStandardLibrary,
	}
}
