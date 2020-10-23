// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package slice

import (
	"math"
	"strconv"
	"testing"
)

func intSliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func pInt(v int) *int {
	return &v
}

func TestBounds(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5}
	testCases := []struct {
		start     int
		end       int
		step      int
		wantLower int
		wantUpper int
	}{
		{start: 0, end: 2, step: 1, wantLower: 0, wantUpper: 2},
		{start: 1, end: 1, step: 1, wantLower: 1, wantUpper: 1},
		{start: 10, end: 12, step: 1, wantLower: len(input), wantUpper: len(input)},
		{start: -2, end: 4, step: 1, wantLower: 4, wantUpper: 4},
		{start: -3, end: -1, step: 1, wantLower: 3, wantUpper: 5},

		{start: 0, end: 2, step: -1, wantLower: 2, wantUpper: 0},
		{start: 1, end: 1, step: -1, wantLower: 1, wantUpper: 1},
		{start: 10, end: 12, step: -1, wantLower: len(input) - 1, wantUpper: len(input) - 1},
		{start: -2, end: 4, step: -1, wantLower: 4, wantUpper: 4},
		{start: -3, end: -1, step: -1, wantLower: 5, wantUpper: 3},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotLower, gotUpper := bounds(&tc.start, &tc.end, tc.step, len(input))
			if gotLower != tc.wantLower || gotUpper != tc.wantUpper {
				t.Errorf("unexpected result, want: (%d, %d), got: (%d, %d)", tc.wantLower, tc.wantUpper, gotLower, gotUpper)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5}
	testCases := []struct {
		got  []int
		want []int
	}{
		// No delimiters
		{Slice(input, nil, nil, nil), input},
		// Out-of-bounds
		{Slice(input, pInt(6), pInt(7), nil), []int{}},
		// Positive delimiters
		{Slice(input, nil, pInt(2), nil), []int{0, 1}},
		{Slice(input, pInt(0), pInt(2), nil), []int{0, 1}},
		{Slice(input, nil, pInt(10), nil), []int{0, 1, 2, 3, 4, 5}},
		{Slice(input, nil, pInt(math.MaxInt32), nil), []int{0, 1, 2, 3, 4, 5}},
		{Slice(input, pInt(-math.MaxInt32), nil, nil), []int{0, 1, 2, 3, 4, 5}},
		// Negative delmiters
		{Slice(input, nil, pInt(-1), nil), []int{0, 1, 2, 3, 4}},
		{Slice(input, pInt(-3), pInt(-2), nil), []int{3}},
		{Slice(input, pInt(-4), pInt(-4), nil), []int{}},
		{Slice(input, nil, pInt(-math.MaxInt32), nil), []int{}},

		// Wider step
		{Slice(input, nil, nil, pInt(0)), []int{}},
		{Slice(input, nil, nil, pInt(2)), []int{0, 2, 4}},
		{Slice(input, pInt(1), nil, pInt(2)), []int{1, 3, 5}},
		{Slice(input, nil, nil, pInt(3)), []int{0, 3}},

		// Negative
		{Slice(input, nil, nil, pInt(-1)), []int{5, 4, 3, 2, 1, 0}},
		{Slice(input, nil, nil, pInt(-2)), []int{5, 3, 1}},
		{Slice(input, pInt(4), pInt(2), pInt(-1)), []int{4, 3}},
		{Slice(input, pInt(-1), pInt(-4), pInt(-1)), []int{5, 4, 3}},
		// {Slice(input, pInt(4), pInt(2), pInt(-1)), []int{5, 4, 3, 2, 1, 0}},
		// {Slice(input, pInt(4), pInt(2), pInt(-1)), []int{5, 4, 3, 2, 1, 0}},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !intSliceEquals(tc.got, tc.want) {
				t.Errorf("unexpected result, want: %v, got: %v", tc.want, tc.got)
			}
		})
	}
}
