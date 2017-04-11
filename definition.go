package swagger

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Definitions map[string]Schema

func (defs *Definitions) UnmarshalJSON(text []byte) error {
	dec := json.NewDecoder(bytes.NewReader(text))

	if _, err := dec.Token(); err != nil {
		return err
	}

	for dec.More() {
		var defName string
		var current Schema

		if token, err := dec.Token(); err == nil {
			defName = fmt.Sprintf("%v", token)
		} else {
			return err
		}

		if err := dec.Decode(&current); err == nil {
			(*defs)[defName] = current
		} else {
			return err
		}
	}

	return nil
}

func (defs Definitions) MarshalJSON() ([]byte, error) {
	marshaled := &bytes.Buffer{}
	enc := json.NewEncoder(marshaled)
	marshaled.WriteRune('{')

	backspace := func() {
		marshaled.Truncate(marshaled.Len() - 1)
	}

	for k, v := range defs {
		enc.Encode(k)
		backspace()
		marshaled.WriteRune(':')
		enc.Encode(v)
		backspace()
		marshaled.WriteRune(',')
	}
	backspace()
	marshaled.WriteRune('}')

	return marshaled.Bytes(), nil
}
