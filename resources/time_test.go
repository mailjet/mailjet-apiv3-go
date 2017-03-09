package resources

import (
	"bytes"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	dateTime, _ := time.Parse(time.RFC3339, "2016-10-14T12:42:05Z")
	dt := RFC3339DateTime{dateTime}

	want := []byte(`"2016-10-14T12:42:05Z"`)
	got, err := dt.MarshalJSON()

	if err != nil {
		t.Error("unexpected error:", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("expected %v", want)
		t.Errorf("     got %v", got)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	data := []byte(`"2016-10-14T12:42:05Z"`)
	var dt RFC3339DateTime

	err := dt.UnmarshalJSON(data)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	if dt.UnixNano() != 1476448925000000000 {
		t.Errorf("expected: %v", 1476448925000000000)
		t.Errorf("     got: %v", dt.UnixNano())
	}
}

func TestUnmarshalEmptyJSON(t *testing.T) {
	var data []byte
	var dt RFC3339DateTime

	err := dt.UnmarshalJSON(data)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	data = []byte{34, 34}
	err = dt.UnmarshalJSON(data)
	if err != nil {
		t.Error("unexpected error:", err)
	}
}

func TestUnmarshalBrokenJSON(t *testing.T) {
	data := []byte{1, 1}
	var dt RFC3339DateTime

	err := dt.UnmarshalJSON(data)
	if err == nil {
		t.Error("error expected")
	}
}
