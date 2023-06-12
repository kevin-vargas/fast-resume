package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"synthesizer-server/internal/service"
)

type synthesizer struct {
	base string
	t    string
}

const (
	model      = "text-davinci-003"
	max_tokens = 200
	promp      = "Summarize the following text: "
)

func (s *synthesizer) Synthesize(str string) (string, error) {
	req := Request{
		Model:     model,
		MaxTokens: max_tokens,
		Prompt:    promp + str,
	}
	raw, err := json.Marshal(&req)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest(http.MethodPost, s.base+"/v1/completions", bytes.NewBuffer(raw))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+s.t)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	var r Response
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return "", err
	}
	if len(r.Choices) > 0 {
		return r.Choices[0].Text, nil
	}
	return "", errors.New("on generate response")
}
func New(t string, base string) service.Synthetizer {
	return &synthesizer{
		base: base,
		t:    t,
	}
}
