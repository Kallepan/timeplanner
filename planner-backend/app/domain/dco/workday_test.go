package dco

import (
	"errors"
	"testing"
)

func TestValidateWorkday(t *testing.T) {
	tests := []struct {
		name string
		req  UpdateWorkdayRequest
		want error
	}{
		{
			name: "valid request",
			req: UpdateWorkdayRequest{
				StartTime: "08:00",
				EndTime:   "16:00",
				Comment:   nil,
				Active:    nil,
			},
			want: nil,
		},
		{
			name: "still valid start time",
			req: UpdateWorkdayRequest{
				StartTime: "08:00:00",
				EndTime:   "16:00",
				Comment:   nil,
				Active:    nil,
			},
			want: nil,
		},
		{
			name: "still valid end time",
			req: UpdateWorkdayRequest{
				StartTime: "08:00",
				EndTime:   "16:00:00",
				Comment:   nil,
				Active:    nil,
			},
			want: nil,
		},
		{
			name: "invalid start and end time",
			req: UpdateWorkdayRequest{
				StartTime: "16:00",
				EndTime:   "08:00",
				Comment:   nil,
				Active:    nil,
			},
			want: errors.New("start time must be before end time"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.req.Validate()
			if got == nil && tt.want == nil {
				return
			}
			if got != nil && tt.want != nil {
				return
			}

			t.Errorf("Validate() = %v, want %v", got, tt.want)
		})
	}

}
