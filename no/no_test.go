package no

import (
	"testing"
	"time"

	"github.com/gohugoio/locales/currency"
)

func TestLocale(t *testing.T) {
	trans := New()
	expected := "no"

	if trans.Locale() != expected {
		t.Errorf("Expected '%s' Got '%s'", expected, trans.Locale())
	}
}

func TestFmtTimeFull(t *testing.T) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Errorf("Expected '<nil>' Got '%s'", err)
	}

	fixed := time.FixedZone("OTHER", -4)

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 5, 1, 0, loc),
			expected: "09:05:01 normaltid for den nordamerikanske østkysten",
		},
		{
			t:        time.Date(2016, 0o2, 0o3, 20, 5, 1, 0, fixed),
			expected: "20:05:01 OTHER",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeFull(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeLong(t *testing.T) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Errorf("Expected '<nil>' Got '%s'", err)
	}

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 5, 1, 0, loc),
			expected: "09:05:01 EST",
		},
		{
			t:        time.Date(2016, 0o2, 0o3, 20, 5, 1, 0, loc),
			expected: "20:05:01 EST",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeLong(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeMedium(t *testing.T) {
	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 5, 1, 0, time.UTC),
			expected: "09:05:01",
		},
		{
			t:        time.Date(2016, 0o2, 0o3, 20, 5, 1, 0, time.UTC),
			expected: "20:05:01",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeShort(t *testing.T) {
	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 5, 1, 0, time.UTC),
			expected: "09:05",
		},
		{
			t:        time.Date(2016, 0o2, 0o3, 20, 5, 1, 0, time.UTC),
			expected: "20:05",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeShort(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateFull(t *testing.T) {
	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 0, 1, 0, time.UTC),
			expected: "onsdag 3. februar 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateFull(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateLong(t *testing.T) {
	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 0, 1, 0, time.UTC),
			expected: "3. februar 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateLong(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateMedium(t *testing.T) {
	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 0, 1, 0, time.UTC),
			expected: "3. feb. 2016",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateShort(t *testing.T) {
	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 0o2, 0o3, 9, 0, 1, 0, time.UTC),
			expected: "03.02.2016",
		},
		{
			t:        time.Date(-500, 0o2, 0o3, 9, 0, 1, 0, time.UTC),
			expected: "03.02.500",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateShort(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtNumber(t *testing.T) {
	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			expected: "1 123 456,56",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtNumber(tt.num, tt.v)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtCurrency(t *testing.T) {
	tests := []struct {
		num      float64
		v        uint64
		currency currency.Type
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			currency: currency.USD,
			expected: "1 123 456,56 USD",
		},
		{
			num:      1123456.5643,
			v:        2,
			currency: currency.NOK,
			expected: "1 123 456,56 kr",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtCurrency(tt.num, tt.v, tt.currency)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtAccounting(t *testing.T) {
	tests := []struct {
		num      float64
		v        uint64
		currency currency.Type
		expected string
	}{
		{
			num:      550.5643,
			v:        2,
			currency: currency.USD,
			expected: "550,56 USD",
		},
		{
			num:      123.56,
			v:        2,
			currency: currency.NOK,
			expected: "123,56 kr",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtAccounting(tt.num, tt.v, tt.currency)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtPercent(t *testing.T) {
	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      15,
			v:        0,
			expected: "15 %",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtPercent(tt.num, tt.v)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}
