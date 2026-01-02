package main

import (
	"fmt"
	"testing"
)

func Test_rotateTheDial(t *testing.T) {
	tests := []struct {
		d                               DialRotation
		wantNewDialPosition             int
		wantNTimesRotationBecomes0SoFar int
		wantErr                         bool
	}{
		// left rotation
		{
			d: DialRotation{
				dialPosition:                3,
				direction:                   'L',
				nStepsToRotate:              1000,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             3,
			wantNTimesRotationBecomes0SoFar: 10,
			wantErr:                         false,
		},
		// right rotation
		{
			d: DialRotation{
				dialPosition:                3,
				direction:                   'R',
				nStepsToRotate:              1000,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             3,
			wantNTimesRotationBecomes0SoFar: 10,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                0,
				direction:                   'R',
				nStepsToRotate:              1,
				nTimesRotationBecomes0SoFar: 1,
			},
			wantNewDialPosition:             1,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                99,
				direction:                   'R',
				nStepsToRotate:              1,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                99,
				direction:                   'R',
				nStepsToRotate:              10,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             9,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                99,
				direction:                   'R',
				nStepsToRotate:              1,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                0,
				direction:                   'R',
				nStepsToRotate:              1,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             1,
			wantNTimesRotationBecomes0SoFar: 0,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                0,
				direction:                   'R',
				nStepsToRotate:              100,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                0,
				direction:                   'R',
				nStepsToRotate:              1000,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 10,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                1,
				direction:                   'R',
				nStepsToRotate:              99,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                1,
				direction:                   'R',
				nStepsToRotate:              99,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
		{
			d: DialRotation{
				dialPosition:                90,
				direction:                   'R',
				nStepsToRotate:              20,
				nTimesRotationBecomes0SoFar: 0,
			},
			wantNewDialPosition:             0,
			wantNTimesRotationBecomes0SoFar: 1,
			wantErr:                         false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			newDialPosition, nTimesRotationBecomes0SoFar, err := rotateTheDial(tt.d)
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
			if nTimesRotationBecomes0SoFar != tt.wantNTimesRotationBecomes0SoFar {
				t.Fatalf("expected %v, but got %v", tt.wantNTimesRotationBecomes0SoFar, nTimesRotationBecomes0SoFar)
			}
		})
	}
}
