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
		return nil // should this be initialized?
	}

	lower, upper := bounds(pstart, pend, step, len(a))
	if step > 0 {
		for i := lower; i <= upper && i < len(a); i += step {
			res = append(res, a[i])
		}
	} else {
		// TODO: fix this
		panic("not implemented")
		for i := upper; lower < i && i < len(a); i += step {
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

		lower = minInt(length, maxInt(0, start))
		upper = minInt(length, maxInt(0, end))
	} else {
		start := normalize(defaultInt(pstart, length-1))
		end := normalize(defaultInt(pend, -(length - 1)))

		lower = minInt(length-1, maxInt(-1, end))
		upper = maxInt(length-1, maxInt(-1, start))
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
		if v < 0 {
			return length - 1 + v
		}
		return v
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
