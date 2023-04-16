// Copyright 2023 The Tangled Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/itsubaki/q"
)

func main() {
	sim := func(a, b bool) {
		qsim := q.New()
		var q0, q1 q.Qubit
		if a {
			q0 = qsim.One()
		} else {
			q0 = qsim.Zero()
		}
		if b {
			q1 = qsim.One()
		} else {
			q1 = qsim.Zero()
		}

		qsim.H(q0).CNOT(q0, q1)

		fmt.Println(a, b)
		for _, s := range qsim.State() {
			fmt.Println(s)
		}
		fmt.Println()
	}
	sim(false, false)
	sim(false, true)
	sim(true, false)
	sim(true, true)
}
