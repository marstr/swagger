package swagger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Responses struct {
	Default *Response
	Codes   map[int]Response
}

type Response struct {
	Description string `json:"description"`
}

func (r *Responses) UnmarshalJSON(text []byte) error {
	dec := json.NewDecoder(bytes.NewReader(text))

	if _, err := dec.Token(); err != nil {
		return err
	}

	for dec.More() {
		var statusCode int
		var response Response
		if statusToken, err := dec.Token(); err == nil {
			statusCode, err = strconv.Atoi(fmt.Sprintf("%v", statusToken))
		} else {
			return err
		}

		r.Codes = make(map[int]Response)

		if err := dec.Decode(&response); err == nil {
			r.Codes[statusCode] = response
		}
	}

	return nil
}
