package object

import (
	"fmt"
	"strings"

	"github.com/c2micro/mlan/pkg/engine/utils"
	"github.com/c2micro/mlan/pkg/parser"
)

// список, содержащий другие объекты
type List struct {
	Impl
	value   []Object
	methods map[string]*BuiltinFunc
}

func NewList(v []Object) *List {
	l := &List{
		value: v,
	}
	l.fillMethods()
	return l
}

func (o *List) fillMethods() {
	o.methods = make(map[string]*BuiltinFunc)
	o.methods["len"] = NewBuiltinFunc("len", o.MethodLen)
	o.methods["reverse"] = NewBuiltinFunc("reverse", o.MethodReverse)
	o.methods["pop"] = NewBuiltinFunc("pop", o.MethodPop)
}

func (o *List) TypeName() string {
	return "list"
}

func (o *List) String() string {
	var items []string
	for _, item := range o.value {
		items = append(items, item.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(items, ", "))
}

func (o *List) GetValue() any {
	return o.value
}

func (o *List) BinaryOp(op int, rhs Object) (Object, error) {
	switch op {
	case parser.MlanLexerOr:
		return o.LogicalOr(rhs)
	case parser.MlanLexerAnd:
		return o.LogicalAnd(rhs)
	case parser.MlanLexerEq:
		return o.Equal(rhs)
	case parser.MlanLexerNeq:
		return o.NotEqual(rhs)
	case parser.MlanLexerAdd:
		return o.Add(rhs)
	case parser.MlanLexerAssSum:
		return o.Add(rhs)
	case parser.MlanLexerMultiply:
		return o.Mul(rhs)
	}
	return nil, ErrInvalidOp
}

func (o *List) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	}
	return nil, ErrInvalidOp
}

func (o *List) IndexGet(rs Object) (Object, error) {
	idx, ok := rs.(*Int)
	if !ok {
		return nil, ErrInvalidIndexType
	}
	if idx.value < 0 || int(idx.value) >= len(o.value) {
		return nil, ErrIndexOutOfRange
	}
	return o.value[idx.value], nil
}

func (o *List) IndexSet(idx Object, rs Object) error {
	idxInt, ok := idx.(*Int)
	if !ok {
		return ErrInvalidIndexType
	}
	if idxInt.value < 0 || int(idxInt.value) >= len(o.value) {
		return ErrIndexOutOfRange
	}
	o.value[idxInt.value] = rs
	return nil
}

func (o *List) CanIterate() bool {
	return true
}

func (o *List) Iterate() Iterator {
	return &ListIterator{
		v: o.value,
		i: 0,
		l: len(o.value),
	}
}

func (o *List) LogicalNot() (Object, error) {
	if len(o.value) == 0 {
		return NewBool(true), nil
	}
	return NewBool(false), nil
}

func (o *List) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *List) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *List) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(false), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(false), nil
	case *Int:
		return NewBool(false), nil
	case *List:
		if len(o.value) != len(rs.(*List).value) {
			return NewBool(false), nil
		}
		for i, v := range o.value {
			val, err := v.Equal(rs.(*List).value[i])
			if err != nil {
				return nil, err
			}
			if !val.GetValue().(bool) {
				return NewBool(false), nil
			}
		}
		return NewBool(true), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, ErrInvalidOp
}

func (o *List) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(true), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(true), nil
	case *Int:
		return NewBool(true), nil
	case *List:
		if len(o.value) != len(rs.(*List).value) {
			return NewBool(true), nil
		}
		for i, v := range o.value {
			val, err := v.Equal(rs.(*List).value[i])
			if err != nil {
				return nil, err
			}
			if !val.GetValue().(bool) {
				return NewBool(true), nil
			}
		}
		return NewBool(false), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, ErrInvalidOp
}

func (o *List) Add(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		o.value = append(o.value, rs)
		return o, nil
	case *Dict:
		o.value = append(o.value, rs)
		return o, nil
	case *Float:
		o.value = append(o.value, rs)
		return o, nil
	case *Int:
		o.value = append(o.value, rs)
		return o, nil
	case *List:
		o.value = append(o.value, rs)
		return o, nil
	case *Null:
		o.value = append(o.value, rs)
		return o, nil
	case *Str:
		o.value = append(o.value, rs)
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *List) Mul(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		var list []Object
		var i int64
		for i = 0; i < utils.BoolToInt(rs.(*Bool).value); i++ {
			list = append(list, o.value...)
		}
		return NewList(list), nil
	case *Int:
		var list []Object
		var i int64
		for i = 0; i < rs.(*Int).value; i++ {
			list = append(list, o.value...)
		}
		return NewList(list), nil
	}
	return nil, ErrInvalidOp
}

func (o *List) MethodCall(name string, args ...Object) (Object, error) {
	m := o.methods[name]
	if m == nil {
		return nil, ErrUnknownMethod
	}
	return m.Call(args...)
}

func (o *List) MethodLen(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	return NewInt(int64(len(o.value))), nil
}

func (o *List) MethodReverse(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	var temp []Object
	for i := len(o.value) - 1; i >= 0; i-- {
		temp = append(temp, o.value[i])
	}
	o.value = temp
	return NewNull(), nil
}

func (o *List) MethodPop(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 arguments, got %d", len(args))
	}
	idx, ok := args[0].(*Int)
	if !ok {
		return nil, fmt.Errorf("expecting int as 1st argument, got '%s'", args[1].TypeName())
	}
	if idx.GetValue().(int64) < 0 || int(idx.GetValue().(int64)) >= len(o.value) {
		return nil, ErrIndexOutOfRange
	}
	o.value = append(o.value[:idx.GetValue().(int64)], o.value[idx.GetValue().(int64)+1:]...)
	return NewNull(), nil
}
