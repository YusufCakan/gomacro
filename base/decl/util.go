/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * util.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package decl

import (
	"sort"
)

// keep only items satisfying pred(item).
// destructively modifies list.
func filter_if_inplace(list []string, pred func(string) bool) []string {
	out := 0
	for _, e := range list {
		if pred(e) {
			list[out] = e
			out++
		}
	}
	return list[:out]
}

// remove all strings equal to 'str' from list
// destructively modifies list.
func remove_item_inplace(str string, list []string) []string {
	out := 0
	for _, e := range list {
		if e != str {
			list[out] = e
			out++
		}
	}
	return list[:out]
}

// make a copy of list
func dup(list []string) []string {
	if len(list) == 0 {
		return nil
	}
	ret := make([]string, len(list))
	copy(ret, list)
	return ret
}

// append, sort and remove duplicates from lists
func sort_unique(list []string) []string {
	list = dup(list)
	if len(list) == 0 {
		return list
	}
	sort.Strings(list)

	prev := list[0]
	out := 1

	// remove duplicates
	for _, e := range list[1:] {
		if e == prev {
			continue
		}
		prev = e
		list[out] = e
		out++
	}
	return list[:out]
}

// sort by kind, then by name
func sort_decls(list []*Decl) {
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		return a.Kind < b.Kind || (a.Kind == b.Kind && a.Name < b.Name)
	})
}