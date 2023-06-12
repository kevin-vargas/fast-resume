package service

import (
	"crypto/sha256"
	"encoding/base64"
	"synthesizer-server/internal/entity"
	"synthesizer-server/internal/server"
)

type Cache[T any] interface {
	Set(string, T) error
	Get(string) (T, bool, error)
}

type Synthetizer interface {
	Synthesize(string) (string, error)
}
type DataCache Cache[entity.Data[string]]
type synthesizer struct {
	st Synthetizer
	dc DataCache
}

func (s *synthesizer) GetSynthesize(strInput string) (*entity.Data[string], error) {
	h := sha256.New()
	h.Write([]byte(strInput))
	hash := base64.RawStdEncoding.EncodeToString(h.Sum(nil))
	if str, ok, err := s.dc.Get(hash); err == nil {
		if ok {
			return &str, nil
		}
	}
	strOutput, err := s.st.Synthesize(strInput)
	if err != nil {
		return nil, err
	}
	out := entity.Data[string]{
		Data: strOutput,
	}
	if err = s.dc.Set(hash, out); err != nil {
		return nil, err
	}
	return &out, nil
}

func New(d DataCache, st Synthetizer) server.Service {
	return &synthesizer{
		st: st,
		dc: d,
	}
}
