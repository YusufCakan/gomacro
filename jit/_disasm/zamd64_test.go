/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * zamd64_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"math/rand"
	"testing"

	. "github.com/cosmos72/gomacro/jit/amd64"
	pkgasm "github.com/cosmos72/gomacro/jit/asm"
)

func Var(index uint16) Mem {
	return MakeMem(int32(index)*8, RSI, Int64)
}

func VarK(index uint16, k Kind) Mem {
	return MakeMem(int32(index)*8, RSI, k)
}

func InitAmd64(asm *Asm) *Asm {
	asm.InitArch(Amd64{})
	asm.RegIncUse(RSI)
	asm.Load(MakeMem(8, RSP, Uint64), MakeReg(RSI, Uint64))
	return asm
}

func TestAmd64Mov(t *testing.T) {

	m := Var(0)
	var asm Asm
	for id := RLo; id <= RHi; id++ {
		InitAmd64(&asm)
		if asm.RegIsUsed(id) {
			continue
		}
		r := MakeReg(id, Int64)
		c := ConstInt64(int64(rand.Uint64()))
		asm.Mov(c, r).Mov(r, m).Epilogue()

		PrintDisasm(t, asm.Code())
	}
}

func TestAmd64Unary(t *testing.T) {
	var asm Asm

	v1, v2, v3 := Var(0), Var(1), Var(2)

	for id := RLo; id <= RHi; id++ {
		asm.InitArch(Amd64{})
		if asm.RegIsUsed(id) {
			continue
		}
		r := MakeReg(id, Int64)
		asm.Asm(MOV, v1, r, //
			NEG1, r, //
			NOT1, r, //
			INC, r, //
			ADD2, v2, r, //
			NOT1, r, //
			NEG1, r, //
			INC, r, //
			MOV, r, v3, //
		)

		PrintDisasm(t, asm.Code())
	}
}

func TestAmd64Sum(t *testing.T) {
	var asm Asm

	Total, I := Var(1), Var(2)
	asm.InitArch(Amd64{}).Asm( //
		MOV, ConstInt64(0xFF), I,
		ADD2, ConstInt64(2), I,
		ADD2, I, Total)

	PrintDisasm(t, asm.Code())
}

func TestAmd64Mul(t *testing.T) {
	var asm Asm

	for _, k := range []Kind{Int8, Int16, Int32, Int64} {
		I, J, K := VarK(0, k), VarK(1, k), VarK(2, k)
		InitAmd64(&asm)
		asm.Asm( //
			MUL2, MakeConst(9, k), I,
			MUL2, MakeConst(16, k), I,
			MUL2, MakeConst(0x7F, k), I,
			MUL3, MakeConst(0x11, k), I, J,
			MUL3, I, J, K,
		)

		PrintDisasm(t, asm.Code())
	}
}

// broken
func _TestAmd64Div(t *testing.T) {
	var asm Asm

	for _, k := range []Kind{Int8, Int16, Int32, Int64} {
		I, J, K := VarK(0, k), VarK(1, k), VarK(2, k)

		InitAmd64(&asm)

		asm.Asm( //
			DIV3, I, J, K,
		)

		if pkgasm.SUPPORTED && pkgasm.ARCH_ID == AMD64 {
			var f func(*int64)
			asm.Func(&f)

			PrintDisasm(t, asm.Code())

			var a, b int64 = 17, 3
			ints := [3]int64{a, b, 0}
			f(&ints[0])
			c := a / b
			if ints[2] != c {
				t.Errorf("DIV3 returned %v, expecting %v", ints[2], c)
			}
		} else {
			PrintDisasm(t, asm.Code())
		}
	}
}

func TestAmd64Cast(t *testing.T) {
	N := [...]Mem{
		VarK(0, Uint64),
		VarK(1, Uint8), VarK(2, Uint16), VarK(3, Uint32),
		VarK(4, Int8), VarK(5, Int16), VarK(6, Int32),
	}
	V := [...]Mem{
		VarK(8, Uint64),
		VarK(9, Uint64), VarK(10, Uint64), VarK(11, Uint64),
		VarK(12, Uint64), VarK(13, Uint64), VarK(14, Uint64),
	}
	var asm Asm
	asm.InitArch(Amd64{})
	asm.Asm(
		NOP,
		CAST, N[1], V[1],
		CAST, N[2], V[2],
		CAST, N[3], V[3],
		CAST, N[4], V[4],
		CAST, N[5], V[5],
		CAST, N[6], V[6],
		NOP,
		CAST, V[1], N[1],
		CAST, V[2], N[2],
		CAST, V[3], N[3],
		CAST, V[4], N[4],
		CAST, V[5], N[5],
		CAST, V[6], N[6],
		RET,
	)

	PrintDisasm(t, asm.Code())
}

func TestAmd64Lea(t *testing.T) {
	N := Var(0)
	M := Var(1)

	var asm Asm
	r0 := asm.InitArch(Amd64{}).RegAlloc(N.Kind())
	r1 := asm.RegAlloc(N.Kind())
	asm.Asm(
		MUL2, ConstInt64(9), N,
		LEA2, N, r0,
		LEA2, M, r0,
		LEA4, M, r0, ConstInt64(2), r1,
	)
	asm.RegFree(r0)

	PrintDisasm(t, asm.Code())
}

func TestAmd64Shift(t *testing.T) {
	N := Var(0)
	M := Var(1)

	var asm Asm
	asm.InitArch(Amd64{})
	asm.RegIncUse(RCX)
	r := MakeReg(RCX, Uint8)
	asm.Asm(
		SHL2, ConstUint64(0), M, // nop
		SHL2, ConstUint64(1), M,
		SHL2, r, N,
		SHR2, ConstUint64(3), M,
		SHR2, r, N,
	)
	asm.RegDecUse(RCX)

	PrintDisasm(t, asm.Code())
}

func TestAmd64SoftReg(t *testing.T) {
	var asm Asm
	asm.InitArch(Amd64{})

	var a, b, c SoftRegId = 0, 1, 2
	asm.Asm(
		ALLOC, a, Uint64,
		ALLOC, b, Uint64,
		ALLOC, c, Uint64,
		MOV, ConstUint64(1), a,
		MOV, ConstUint64(2), b,
		ADD3, a, b, c,
		FREE, a, Uint64,
		FREE, b, Uint64,
		FREE, c, Uint64,
	).Epilogue()
	PrintDisasm(t, asm.Code())
}
