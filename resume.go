// Package resume is a resume management tool
package resume

import (
	"encoding/json"
	"os"
)

const (
	DefaultFileName = "resume.json"
)

func New(filename string) (Resume, error) {
	r := new(resume)
	err := r.Load(filename)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type Resume interface {
	json.Marshaler
	json.Unmarshaler
}

type resume struct{}

func (r *resume) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *resume) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *resume) Load(filename string) error {
	if filename == "" {
		filename = DefaultFileName
	}
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return r.UnmarshalJSON(b)
}

func (r *resume) Save(filename string) error {
	if filename == "" {
		filename = DefaultFileName
	}

	data, err := r.MarshalJSON()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
