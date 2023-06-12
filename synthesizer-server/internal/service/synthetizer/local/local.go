package local

import "synthesizer-server/internal/service"

type synthesizer struct {
}

func (s *synthesizer) Synthesize(str string) (string, error) {
	return "foo:bar", nil
}
func New() service.Synthetizer {
	return &synthesizer{}
}
