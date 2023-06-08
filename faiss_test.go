package faiss

import (
	"testing"
)

func TestReadIndex(t *testing.T) {
	_, err := ReadIndex("/Users/fedorkurusin/Desktop/presets.index", IOFlagReadOnly)
	if err != nil {
		t.Errorf("can't read index: %v", err)
	}
}
