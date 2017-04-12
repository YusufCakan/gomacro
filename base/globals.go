/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * globals.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"sort"
	"strings"

	"github.com/cosmos72/gomacro/imports"
)

type PackageRef struct {
	imports.Package
	Name, Path string
}

type Options uint
type WhichMacroExpand uint

const (
	OptTrapPanic Options = 1 << iota
	OptPanicStackTrace
	OptMacroExpandOnly // do not compile or execute code, only parse and macroexpand it
	OptCollectDeclarations
	OptCollectStatements
	OptShowPrompt
	OptShowParse
	OptShowMacroExpand
	OptShowCompile
	OptShowEval
	OptShowEvalType
	OptShowTime
	OptDebugMacroExpand
	OptDebugQuasiquote
	OptDebugCallStack
	OptDebugPanicRecover
)

const (
	CMacroExpand1 WhichMacroExpand = iota
	CMacroExpand
	CMacroExpandCodewalk
)

var optNames = map[Options]string{
	OptTrapPanic:           "TrapPanic",
	OptPanicStackTrace:     "OptPanicStackTrace",
	OptMacroExpandOnly:     "MacroExpandOnly",
	OptCollectDeclarations: "Declarations",
	OptCollectStatements:   "Statements",
	OptShowPrompt:          "Prompt.Show",
	OptShowParse:           "Parse.Show",
	OptShowMacroExpand:     "MacroExpand.Show",
	OptShowCompile:         "Compile.Show",
	OptShowEval:            "Eval.Show",
	OptShowEvalType:        "Type.Eval.Show",
	OptShowTime:            "Time.Show",
	OptDebugMacroExpand:    "?MacroExpand",
	OptDebugQuasiquote:     "?Quasiquote",
	OptDebugCallStack:      "?CallStack",
	OptDebugPanicRecover:   "?PanicRecover",
}

var optValues = map[string]Options{}

func init() {
	for k, v := range optNames {
		optValues[v] = k
	}
}

func (o Options) String() string {
	names := make([]string, 0)
	for k, v := range optNames {
		if k&o != 0 {
			names = append(names, v)
		}
	}
	sort.Strings(names)
	return strings.Join(names, " ")
}

func ParseOptions(str string) Options {
	var opts Options
	for _, name := range strings.Split(str, " ") {
		if opt, ok := optValues[name]; ok {
			opts ^= opt
		} else if len(name) != 0 {
			for k, v := range optNames {
				if strings.HasPrefix(v, name) {
					opts ^= k
				}
			}
		}
	}
	return opts
}

func (m WhichMacroExpand) String() string {
	switch m {
	case CMacroExpand1:
		return "MacroExpand1"
	case CMacroExpandCodewalk:
		return "MacroExpandCodewalk"
	default:
		return "MacroExpand"
	}
}