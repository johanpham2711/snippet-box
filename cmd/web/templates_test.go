package main

import (
	"testing"
	"time"

	"github.com/johanpham2711/snippet-box/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			// Use the new assert.Equal() helper to compare the expected and
			// actual values.
			assert.Equal(t, hd, tt.want)
		})
	}
}

func TestHumanDateWithSubTests(t *testing.T) {
	t.Run("UTC", func(t *testing.T) {
		hd := humanDate(time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC))

		assert.Equal(t, hd, "17 Mar 2024 at 10:15")
	})

	t.Run("Empty", func(t *testing.T) {
		hd := humanDate(time.Time{})

		assert.Equal(t, hd, "")
	})

	t.Run("CET", func(t *testing.T) {
		hd := humanDate(time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)))

		assert.Equal(t, hd, "17 Mar 2024 at 09:15")
	})
}
