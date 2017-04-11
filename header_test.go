package swagger

import (
	"testing"

	"reflect"

	"encoding/json"

	"github.com/Azure/go-autorest/autorest/to"
)

func TestHeaders_UnmarshalJSON(t *testing.T) {
	const marshaled = `{
    "X-Rate-Limit-Limit": {
        "description": "The number of allowed requests in the current period",
        "type": "integer"
    },
    "X-Rate-Limit-Remaining": {
        "description": "The number of remaining requests in the current period",
        "type": "integer"
    },
    "X-Rate-Limit-Reset": {
        "description": "The number of seconds left in the current period",
        "type": "integer"
    }
}`

	expected := Headers{
		"X-Rate-Limit-Limit": Header{
			Description: to.StringPtr("The number of allowed requests in the current period"),
			Type:        HeaderTypeInteger,
		},
		"X-Rate-Limit-Remaining": Header{
			Description: to.StringPtr("The number of remaining requests in the current period"),
			Type:        HeaderTypeInteger,
		},
		"X-Rate-Limit-Reset": Header{
			Description: to.StringPtr("The number of seconds left in the current period"),
			Type:        HeaderTypeInteger,
		},
	}

	result := make(Headers)

	if err := json.Unmarshal([]byte(marshaled), &result); err != nil {
		t.Error(err)
	}

	for key, value := range result {
		if fromExpected, ok := expected[key]; ok {
			if !reflect.DeepEqual(value, fromExpected) {
				t.Logf("got:\n%v\nwant:%v", value, fromExpected)
				t.Fail()
			}
		} else {
			t.Logf("Did not find expected key '%s'", key)
			t.Fail()
		}
	}
}

func TestHeaders_MarshalJSON(t *testing.T) {
	unmarshaled := Headers{
		"nameA": Header{
			Description: to.StringPtr("descriptionA"),
			Type:        HeaderTypeInteger,
		},
		"nameB": Header{
			Description: to.StringPtr("descriptionB"),
			Type:        HeaderTypeBoolean,
		},
	}

	const expected = `{"nameA":{"description":"descriptionA","type":"integer"},"nameB":{"description":"descriptionB","type":"boolean"}}`

	if result, err := json.Marshal(unmarshaled); err != nil {
		t.Error(err)
	} else if string(result) != expected {
		t.Logf("\ngot:\n%s\nwant:\n%s", string(result), expected)
		t.Fail()
	}
}
