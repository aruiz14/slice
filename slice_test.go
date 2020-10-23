// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package slice

import (
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
		{start: -2, end: 4, step: 1, wantLower: 3, wantUpper: 4},
		{start: -3, end: -1, step: 1, wantLower: 2, wantUpper: 4},

		// TODO: test negative step
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
		{Slice(input, nil, nil, nil), input},
		{Slice(input, pInt(6), pInt(7), nil), []int{}},
		{Slice(input, pInt(0), pInt(2), nil), []int{0, 1, 2}},
		{Slice(input, pInt(0), pInt(10), nil), []int{0, 1, 2, 3, 4, 5}},
		{Slice(input, pInt(0), pInt(-1), nil), []int{0, 1, 2, 3, 4}},
		{Slice(input, pInt(-3), pInt(-2), nil), []int{2, 3}},
		{Slice(input, pInt(-4), pInt(-4), nil), []int{1}},
		{Slice(input, pInt(2), pInt(5), pInt(2)), []int{2, 4}},

		// TODO: add negative step tests
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !intSliceEquals(tc.got, tc.want) {
				t.Errorf("unexpected result, want: %v, got: %v", tc.want, tc.got)
			}
		})
	}
}
