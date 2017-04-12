package swagger

import (
	"testing"
)

func Test_UnescapeJSONReferencePathElement(t *testing.T) {
	testCases := []struct {
		element string
		want    string
	}{
		{"standard", "standard"},
		{"", ""},
		{"hey~0", "hey~"},
		{"hey~1", "hey/"},
		{"~0~1", "~/"},
		{"~01", "~1"},
	}

	for _, tc := range testCases {
		t.Run(tc.element, func(t *testing.T) {
			if got, err := unescapeJSONReferencePathElement(tc.element); err != nil {
				t.Error(err)
			} else if got != tc.want {
				t.Logf("got: '%s' want: '%s'", got, tc.want)
				t.Fail()
			}
		})
	}
}

func Test_UnescapedJSONReferencePathElement_Errs(t *testing.T) {
	errCases := []func(*testing.T){
		func(subT *testing.T) {
			if _, err := unescapeJSONReferencePathElement("element1/element2"); err.Error() != "multiple elements were passed as a single one" {
				subT.Fail()
			}
		},
	}

	for _, tc := range errCases {
		t.Run("", tc)
	}
}

func Test_EscapedJSONReferencePathElement(t *testing.T) {
	testCases := []struct {
		element string
		want    string
	}{
		{"standard", "standard"},
		{"one~two", "one~0two"},
		{"~01", "~001"},
	}

	for _, tc := range testCases {
		t.Run(tc.element, func(t *testing.T) {
			got := escapeJSONReferencePathElement(tc.element)
			if got != tc.want {
				t.Logf("got: '%s' want: '%s'", got, tc.want)
				t.Fail()
			}
		})
	}
}
