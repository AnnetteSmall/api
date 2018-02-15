package train

import (
	"testing"
)

func expectTrainSpeed(t *testing.T, tr *Train, s int) {
	actual := tr.Speed()
	if actual != s {
		t.Fatal("expected", s, "got", actual)
	}
}

func TestTrain(t *testing.T) {
	tr := NewTrain()
	tr.Go()
	expectTrainSpeed(t, tr, 20)
	tr.Stop()
	expectTrainSpeed(t, tr, 0)
}
