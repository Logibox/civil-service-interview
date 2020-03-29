package numorstr

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
)

type NumberOrString struct {
	n json.Number
}

func (nos *NumberOrString) Validate(formats strfmt.Registry) error {
	return nil
}

func (nos *NumberOrString) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &(nos.n))
}

func (nos *NumberOrString) MarshalJSON() ([]byte, error) {
	f, err := nos.Float64()
	if err != nil {
		return nil, err
	}
	return json.Marshal(f)
}

func (nos *NumberOrString) Float64() (float64, error) {
	return nos.n.Float64()
}
