// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * callnret1.go
 *
 *  Created on Apr 20, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

func callnret1(c *Call, maxdepth int) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	kret := expr.Type.Out(0).Kind()
	argfuns := c.MakeArgfuns()
	var call I
	switch kret {
	case r.Bool:
		call = func(env *Env) bool {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Bool()
		}
	case r.Int:
		call = func(env *Env) int {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int(ret0.Int())
		}
	case r.Int8:
		call = func(env *Env) int8 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int8(ret0.Int())
		}
	case r.Int16:
		call = func(env *Env) int16 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int16(ret0.Int())
		}
	case r.Int32:
		call = func(env *Env) int32 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return int32(ret0.Int())
		}
	case r.Int64:
		call = func(env *Env) int64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Int()
		}
	case r.Uint:
		call = func(env *Env) uint {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint(ret0.Uint())
		}
	case r.Uint8:
		call = func(env *Env) uint8 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint8(ret0.Uint())
		}
	case r.Uint16:
		call = func(env *Env) uint16 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint16(ret0.Uint())
		}

	case r.Uint32:
		call = func(env *Env) uint32 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uint32(ret0.Uint())
		}

	case r.Uint64:
		call = func(env *Env) uint64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Uint()
		}

	case r.Uintptr:
		call = func(env *Env) uintptr {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return uintptr(ret0.Uint())
		}

	case r.Float32:
		call = func(env *Env) float32 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return float32(ret0.Float())
		}

	case r.Float64:
		call = func(env *Env) float64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Float()
		}

	case r.Complex64:
		call = func(env *Env) complex64 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return complex64(ret0.Complex())
		}

	case r.Complex128:
		call = func(env *Env) complex128 {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.Complex()
		}

	case r.String:
		call = func(env *Env) string {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0.String()
		}

	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			ret0 := funv.Call(argv)[0]
			return ret0

		}

	}
	return call
}
