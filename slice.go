// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package slice

// Slice function
func Slice(a []int, pstart *int, pend *int, pstep *int) (res []int) {
	if pstart == nil && pend == nil && pstep == nil {
		return a
	}

	step := defaultInt(pstep, 1)
	if step == 0 {
		return []int{}
	}

	lower, upper := bounds(pstart, pend, step, len(a))
	if step > 0 {
		for i := lower; i < upper; i += step {
			res = append(res, a[i])
		}
	} else {
		for i := upper; lower < i; i += step {
			res = append(res, a[i])
		}
	}
	return
}

func bounds(pstart, pend *int, step, length int) (lower, upper int) {
	normalize := normalizer(length)
	if step >= 0 {
		start := normalize(defaultInt(pstart, 0))
		end := normalize(defaultInt(pend, length))

		lower = minInt(maxInt(start, 0), length)
		upper = minInt(maxInt(end, 0), length)
	} else {
		start := normalize(defaultInt(pstart, length-1))
		end := normalize(defaultInt(pend, -length-1))

		upper = minInt(maxInt(start, -1), length-1)
		lower = minInt(maxInt(end, -1), length-1)
	}
	return
}

func defaultInt(v *int, def int) int {
	if v != nil {
		return *v
	}
	return def
}

type normalizeFunc func(int) int

func normalizer(length int) normalizeFunc {
	return func(v int) int {
		if v >= 0 {
			return v
		}
		return length + v
	}
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
