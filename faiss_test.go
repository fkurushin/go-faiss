package faiss

import (
	"strconv"
	"testing"
)

func TestReadIndex(t *testing.T) {
	_, err := ReadIndex("/Users/fedorkurusin/Desktop/presets.index", IOFlagReadOnly)
	if err != nil {
		t.Errorf("can't read index: %v", err)
	}
}

func TestIndexFactory(t *testing.T) {
	d := 100
	k := 1000
	_, err := IndexFactory(d, "IVF"+strconv.Itoa(k)+",Flat", MetricInnerProduct)
	if err != nil {
		t.Errorf("failed to create FAISS index: %v", err)
	}
}

func TestTrain(t *testing.T) {
	d := 100
	k := 1000
	_, err := IndexFactory(d, "IVF"+strconv.Itoa(k)+",Flat", MetricInnerProduct)
	if err != nil {
		t.Errorf("failed to create FAISS index: %v", err)
	}
}
