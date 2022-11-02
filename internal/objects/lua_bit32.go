package objects

import (
	lua "github.com/yuin/gopher-lua"
	"math"
)

const bitCount = 32

func registerLuaBit32(l *lua.LState) int {
	mod := l.RegisterModule("Bit32", bitFuncs).(*lua.LTable)
	l.Push(mod)
	return 1
}

var bitFuncs = map[string]lua.LGFunction{
	"arshift": bitArshift,
	"band":    bitBand,
	"bnot":    bitBnot,
	"bor":     bitBor,
	"bxor":    bitBxor,
	"btest":   bitBtest,
	"extract": bitExtract,
	"lrotate": bitLrotate,
	"lshift":  bitLshift,
	"replace": bitReplace,
	"rrotate": bitRrotate,
	"rshift":  bitRshift,
}

func checkUnsignedNumber(ls *lua.LState, n int) uint {
	v := ls.Get(n)
	if lv, ok := v.(lua.LNumber); ok {
		const supUnsigned = float64(^uint32(0)) + 1
		return uint(float64(lv) - math.Floor(float64(lv)/supUnsigned)*supUnsigned)
	}
	ls.TypeError(n, lua.LTNumber)
	return 0
}

func bitArshift(l *lua.LState) int {
	if r, i := checkUnsignedNumber(l, 1), int(l.CheckNumber(2)); i < 0 || 0 == (r&(1<<(bitCount-1))) {
		return shift(l, r, -i)
	} else {
		if i >= bitCount {
			r = math.MaxUint32
		} else {
			r = trim((r >> uint(i)) | ^(math.MaxUint32 >> uint(i)))
		}
		l.Push(lua.LNumber(r))
	}
	return 1
}

func bitBand(l *lua.LState) int {
	l.Push(lua.LNumber(andHelper(l)))
	return 1
}

func bitBnot(l *lua.LState) int {
	l.Push(lua.LNumber(trim(^checkUnsignedNumber(l, 1))))
	return 1
}

func bitBor(l *lua.LState) int {
	l.Push(lua.LNumber(bitOp(l, 0, func(a, b uint) uint { return a | b })))
	return 1
}

func bitBxor(l *lua.LState) int {
	l.Push(lua.LNumber(bitOp(l, 0, func(a, b uint) uint { return a ^ b })))
	return 1
}

func bitBtest(l *lua.LState) int {
	l.Push(lua.LBool(andHelper(l) != 0))
	return 1
}

func bitExtract(l *lua.LState) int {
	r := checkUnsignedNumber(l, 1)
	f, w := fieldArguments(l, 2)
	l.Push(lua.LNumber((r >> f) & mask(w)))
	return 1
}

func bitLrotate(l *lua.LState) int {
	return rotate(l, int(l.CheckNumber(2)))
}

func bitLshift(l *lua.LState) int {
	return shift(l, checkUnsignedNumber(l, 1), int(l.CheckNumber(2)))
}

func bitReplace(l *lua.LState) int {
	r, v := checkUnsignedNumber(l, 1), checkUnsignedNumber(l, 2)
	f, w := fieldArguments(l, 3)
	m := mask(w)
	v &= m
	l.Push(lua.LNumber((r & ^(m << f)) | (v << f)))
	return 1
}

func bitRrotate(l *lua.LState) int {
	return rotate(l, -int(l.CheckNumber(2)))
}

func bitRshift(l *lua.LState) int {
	return shift(l, checkUnsignedNumber(l, 1), -int(l.CheckNumber(2)))
}

func trim(x uint) uint { return x & math.MaxUint32 }
func mask(n uint) uint { return ^(math.MaxUint32 << n) }

func shift(l *lua.LState, r uint, i int) int {
	if i < 0 {
		if i, r = -i, trim(r); i >= bitCount {
			r = 0
		} else {
			r >>= uint(i)
		}
	} else {
		if i >= bitCount {
			r = 0
		} else {
			r <<= uint(i)
		}
		r = trim(r)
	}
	l.Push(lua.LNumber(r))
	return 1
}

func rotate(l *lua.LState, i int) int {
	r := trim(checkUnsignedNumber(l, 1))
	if i &= bitCount - 1; i != 0 {
		r = trim((r << uint(i)) | (r >> uint(bitCount-i)))
	}
	l.Push(lua.LNumber(r))
	return 1
}

func bitOp(l *lua.LState, init uint, f func(a, b uint) uint) uint {
	r := init
	for i, n := 1, l.GetTop(); i <= n; i++ {
		r = f(r, checkUnsignedNumber(l, i))
	}
	return trim(r)
}

func andHelper(l *lua.LState) uint {
	x := bitOp(l, ^uint(0), func(a, b uint) uint { return a & b })
	return x
}

func fieldArguments(l *lua.LState, fieldIndex int) (uint, uint) {
	f, w := l.CheckNumber(fieldIndex), lua.LNumber(l.OptInt(fieldIndex+1, 1))
	ArgumentCheck(l, 0 <= f, fieldIndex, "field cannot be negative")
	ArgumentCheck(l, 0 < w, fieldIndex+1, "width must be positive")
	if f+w > bitCount {
		l.RaiseError("trying to access non-existent bits")
	}
	return uint(f), uint(w)
}

// ArgumentCheck checks whether cond is true. If not, raises an error with a standard message.
func ArgumentCheck(l *lua.LState, cond bool, index int, extraMessage string) {
	if !cond {
		l.ArgError(index, extraMessage)
	}
}
