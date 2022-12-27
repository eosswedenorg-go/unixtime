package unixtime

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestTime_Time(t *testing.T) {
	tests := []struct {
		name string
		ts   Time
		want time.Time
	}{
		{"Epoc", Time(0), time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{"Two Seconds after epoch", Time(2000), time.Date(1970, time.January, 1, 0, 0, 2, 0, time.UTC)},
		{"Date1", Time(1644612684432), time.Date(2022, time.February, 11, 20, 51, 24, 0, time.UTC).Add(time.Millisecond * 432)},
		{"Date2", Time(1831324037241), time.Date(2028, time.January, 12, 21, 7, 17, 0, time.UTC).Add(time.Millisecond * 241)},
		{"Date3", Time(1272908563433), time.Date(2010, time.May, 3, 17, 42, 43, 0, time.UTC).Add(time.Millisecond * 433)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_FromTime(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  Time
	}{
		{"Epoc", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC), Time(0)},
		{"Two Seconds after epoch", time.Date(1970, time.January, 1, 0, 0, 2, 0, time.UTC), Time(2000)},
		{"Date1", time.Date(2022, time.February, 11, 20, 51, 24, 0, time.UTC).Add(time.Millisecond * 432), Time(1644612684432)},
		{"Date2", time.Date(2028, time.January, 12, 21, 7, 17, 0, time.UTC).Add(time.Millisecond * 241), Time(1831324037241)},
		{"Date3", time.Date(2010, time.May, 3, 17, 42, 43, 0, time.UTC).Add(time.Millisecond * 433), Time(1272908563433)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ts Time
			ts.FromTime(tt.input)
			if !reflect.DeepEqual(ts, tt.want) {
				t.Errorf("Time.Time() = %v, want %v", ts, tt.want)
			}
		})
	}
}

func TestTime_MarshalJson(t *testing.T) {
	tests := []struct {
		name      string
		input     Time
		expectErr bool
		expected  []byte
	}{
		{"number", Time(1074932802), false, []byte("1074932802")},
		{"number milliseconds", Time(1800718379432), false, []byte("1800718379432")},
		{"zero", Time(0), false, []byte("0")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			var res []byte

			if res, err = json.Marshal(tt.input); (err != nil) != tt.expectErr {
				s := "did not expect"
				if tt.expectErr {
					s = "expected"
				}
				t.Errorf("Time.MarshalJSON(%v) %s error but got: %v", tt.input, s, err)
				return
			}

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Time.MarshalJSON(%v) encoded value = %s, expected: %s", tt.input, string(res), string(tt.expected))
			}
		})
	}
}

func TestTime_UnmarshalJson(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expectErr bool
		expected  Time
	}{
		{"number", []byte("1074932802"), false, Time(1074932802)},
		{"number milliseconds", []byte("1800718379432"), false, Time(1800718379432)},
		{"string", []byte("\"1476870484\""), false, Time(1476870484)},
		{"string milliseconds", []byte("\"1440894197834\""), false, Time(1440894197834)},
		{"null string", []byte("null"), false, Time(0)},
		{"random", []byte{0x1, 0xff, 0x3c}, true, Time(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ts Time

			if err := ts.UnmarshalJSON(tt.input); (err != nil) != tt.expectErr {
				s := "did not expect"
				if tt.expectErr {
					s = "expected"
				}
				t.Errorf("Time.UnmarshalJSON(%s) %s error but got: %v", string(tt.input), s, err)
				return
			}

			if ts != tt.expected {
				t.Errorf("Time.UnmarshalJSON(%s) parsed value = %v, expected: %v", string(tt.input), ts, tt.expected)
			}
		})
	}
}
