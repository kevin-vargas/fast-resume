package service

import (
	"bff-server/internal/entity"
	"bff-server/internal/server"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	GetChannels(context.Context) ([]entity.Channel, error)
	Summarize(context.Context, string) (string, error)
}

type service struct {
	baseSynthesizerServer string
	r                     Repository
}

func (s *service) GetChannels(ctx context.Context) (*entity.Channels, error) {
	channels, err := s.r.GetChannels(ctx)
	if err != nil {
		return nil, err
	}
	for i, channel := range channels {
		if errs, ok := validate(channel); !ok {
			channel.Err = errs
			channels[i] = channel
		}
	}
	return &entity.Channels{
		Data: channels,
	}, nil
}

func (s *service) SummarizeChannel(ctx context.Context, id string) (*entity.Data[string], error) {
	summarize, err := s.r.Summarize(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Data[string]{
		Data: summarize,
	}, nil
}

func (s *service) SummarizeIAChannel(ctx context.Context, id string) (*entity.Data[string], error) {
	data, err := s.SummarizeChannel(ctx, id)
	if err != nil {
		return nil, err
	}
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(s.baseSynthesizerServer+"/synthesize", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return nil, err
	}
	var d entity.Data[string]
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func New(r Repository, b string) server.Service {
	return &service{
		baseSynthesizerServer: b,
		r:                     r,
	}
}

func validate(req any) ([]string, bool) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("fieldName"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := v.Struct(req); err != nil {
		fieldError, oneofError, minError, defaultError := make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0)

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				fieldError = append(fieldError, err.Field())
			case "oneof":
				oneofError = append(oneofError, err.Field())
			case "min", "len", "lte", "max":
				minError = append(minError, err.Field())
			default:
				defaultError = append(defaultError, err.Field())
			}
		}
		errs := []string{}
		if len(fieldError) > 0 {
			errs = append(errs, fmt.Sprintf("Los siguientes campos son requeridos: %v", strings.Join(fieldError, ", ")))
		}
		if len(oneofError) > 0 {
			errs = append(errs, fmt.Sprintf("Los siguientes campos no hacen match con el enumerado: %v", strings.Join(oneofError, ", ")))
		}
		if len(minError) > 0 {
			errs = append(errs, fmt.Sprintf("Los siguientes campos no cumplen con la longitud de caracteres requeridos: %v", strings.Join(minError, ", ")))
		}
		if len(defaultError) > 0 {
			errs = append(errs, fmt.Sprintf("Los siguientes campos son inválidos: %v", strings.Join(defaultError, ", ")))
		}
		return errs, false
	}
	return nil, true
}
