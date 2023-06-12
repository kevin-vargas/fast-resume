package slack

import (
	"bff-server/internal/entity"
	"bff-server/internal/middleware"
	"bff-server/internal/service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type repository struct {
	base string
}

func (s *repository) GetChannels(ctx context.Context) ([]entity.Channel, error) {
	t, ok := middleware.GetToken(ctx)
	if !ok {
		return nil, errors.New("invalid token")
	}
	request, err := http.NewRequest(http.MethodGet, s.base+"/users.conversations", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+t)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	var b Channels
	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return nil, err
	}
	if !b.Ok {
		return nil, fmt.Errorf("on get channel %s", b.Err)
	}
	channels := make([]entity.Channel, len(b.Channels))
	for i, ch := range b.Channels {
		channels[i] = entity.Channel{
			Title:       ch.Name,
			Description: ch.Purpose.Value,
			ID:          ch.ID,
		}
	}
	return channels, nil
}

func (s *repository) Summarize(ctx context.Context, id string) (string, error) {
	t, ok := middleware.GetToken(ctx)
	if !ok {
		return "", errors.New("invalid token")
	}
	request, err := http.NewRequest(http.MethodGet, s.base+"/conversations.history?channel="+id, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", "Bearer "+t)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	var b Conversations
	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return "", err
	}
	if !b.Ok {
		return "", fmt.Errorf("on get conversations history %s", b.Err)
	}
	var bf strings.Builder

	for i := len(b.Messages) - 1; i >= 0; i-- {
		ch := b.Messages[i]
		userName, err := s.getUserName(ctx, ch.User)
		if err != nil {
			return "", err
		}
		bf.WriteString(userName)
		bf.WriteString(": ")
		text, err := s.sanitize(ctx, ch.Text)
		if err != nil {
			return "", err
		}
		bf.WriteString(text)
		bf.WriteString("\n")
	}
	return bf.String(), nil
}

func (s *repository) sanitize(ctx context.Context, text string) (string, error) {
	regex := regexp.MustCompile("<@.*?>")
	result := regex.ReplaceAllStringFunc(text, func(str string) string {
		id := strings.ReplaceAll(str[2:], ">", "")
		userName, err := s.getUserName(ctx, id)
		if err != nil {
			return id
		}
		return userName
	})
	return result, nil
}

func (s *repository) getUserName(ctx context.Context, id string) (string, error) {
	t, ok := middleware.GetToken(ctx)
	if !ok {
		return "", errors.New("invalid token")
	}
	request, err := http.NewRequest(http.MethodGet, s.base+"/users.info?user="+id, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", "Bearer "+t)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	var b UserInfo
	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return "", err
	}
	if !b.Ok {
		return "", fmt.Errorf("on get userinfo %s", b.Err)
	}

	return b.User.Name, nil
}

func New(base string) service.Repository {
	return &repository{
		base: base,
	}
}
