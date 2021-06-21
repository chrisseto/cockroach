package trigram_test

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/util/trigram"
	"github.com/stretchr/testify/assert"
)

func TestTrigrams(t *testing.T) {
	testCases := []struct {
		In  string
		Out []string
	}{
		{
			In:  "f",
			Out: []string{"  f"},
		},
		{
			In:  "fo",
			Out: []string{"  f", " fo"},
		},
		{
			In:  "cat",
			Out: []string{"  c", " ca", "at ", "cat"},
		},
		{
			In:  "foo|bar",
			Out: []string{"  b", "  f", " ba", " fo", "ar ", "bar", "foo", "oo "},
		},
		{
			In:  "words words words",
			Out: []string{"  w", " wo", "ds ", "ord", "rds", "wor"},
		},
	}

	for _, tc := range testCases {
		actual := trigram.Trigrams(tc.In)
		assert.Equal(t, tc.Out, actual)
	}
}

func TestSimilarity(t *testing.T) {
	testCases := []struct {
		Word1 string
		Word2 string
		Out   float32
	}{
		{
			Word1: "",
			Word2: "foo",
			Out:   0.0,
		},
		{
			Word1: "foo",
			Word2: "",
			Out:   0.0,
		},
		{
			Word1: "word",
			Word2: "two words",
			Out:   0.36363637,
		},
	}

	for _, tc := range testCases {
		actual := trigram.Similarity(tc.Word1, tc.Word2)
		assert.Equal(t, tc.Out, actual)
	}
}

func TestWordSimilarity(t *testing.T) {
	testCases := []struct {
		Word1 string
		Word2 string
		Out   float32
	}{
		{
			Word1: "",
			Word2: "foo",
			Out:   0.0,
		},
		{
			Word1: "foo",
			Word2: "",
			Out:   0.0,
		},
		{
			Word1: "word",
			Word2: "two words",
			Out:   0.36363637,
		},
	}

	for _, tc := range testCases {
		actual := trigram.WordSimilarity(tc.Word1, tc.Word2)
		assert.Equal(t, tc.Out, actual)
	}
}
