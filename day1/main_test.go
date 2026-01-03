package main

import (
	"fmt"
	"testing"
)

func Test_rotateTheDial(t *testing.T) {
	tests := []struct {
		d                                DialRotation
		wantNewDialPosition              int
		wantNTimesCrossed0InThisRotation int
		wantErr                          bool
	}{
		{
			d: DialRotation{
				dialPosition:   3,
				direction:      'R',
				nStepsToRotate: 1000,
			},
			wantNewDialPosition:              3,
			wantNTimesCrossed0InThisRotation: 10,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   99,
				direction:      'R',
				nStepsToRotate: 1,
			},
			wantNewDialPosition:              0,
			wantNTimesCrossed0InThisRotation: 1,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   99,
				direction:      'R',
				nStepsToRotate: 10,
			},
			wantNewDialPosition:              9,
			wantNTimesCrossed0InThisRotation: 1,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   99,
				direction:      'R',
				nStepsToRotate: 1,
			},
			wantNewDialPosition:              0,
			wantNTimesCrossed0InThisRotation: 1,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   0,
				direction:      'R',
				nStepsToRotate: 1,
			},
			wantNewDialPosition:              1,
			wantNTimesCrossed0InThisRotation: 0,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   0,
				direction:      'R',
				nStepsToRotate: 100,
			},
			wantNewDialPosition:              0,
			wantNTimesCrossed0InThisRotation: 1,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   0,
				direction:      'R',
				nStepsToRotate: 1000,
			},
			wantNewDialPosition:              0,
			wantNTimesCrossed0InThisRotation: 10,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   1,
				direction:      'R',
				nStepsToRotate: 99,
			},
			wantNewDialPosition:              0,
			wantNTimesCrossed0InThisRotation: 1,
			wantErr:                          false,
		},
		{
			d: DialRotation{
				dialPosition:   1,
				direction:      'R',
				nStepsToRotate: 99,
			},
			wantNewDialPosition:              0,
			wantNTimesCrossed0InThisRotation: 1,
			wantErr:                          false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			newDialPosition, nTimesCrossed0InThisRotation, err := rotateTheDial(tt.d)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("rotateTheDial() failed: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("rotateTheDial() succeeded unexpectedly")
			}
			if newDialPosition != tt.wantNewDialPosition {
				t.Fatalf("expected %v, but got %v", tt.wantNewDialPosition, newDialPosition)
			}
			if nTimesCrossed0InThisRotation != tt.wantNTimesCrossed0InThisRotation {
				t.Fatalf("expected %v, but got %v", tt.wantNTimesCrossed0InThisRotation, nTimesCrossed0InThisRotation)
			}
		})
	}
}
