// Code generated by thriftrw v1.3.0
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_EchoStringSet_Args struct {
	Arg map[string]struct{} `json:"arg,required"`
}

type _Set_String_ValueList map[string]struct{}

func (v _Set_String_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		w, err := wire.NewValueString(x), error(nil)
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_String_ValueList) Size() int {
	return len(v)
}

func (_Set_String_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Set_String_ValueList) Close() {
}

func (v *SecondService_EchoStringSet_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg == nil {
		return w, errors.New("field Arg of SecondService_EchoStringSet_Args is required")
	}
	w, err = wire.NewValueSet(_Set_String_ValueList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Set_String_Read(s wire.ValueList) (map[string]struct{}, error) {
	if s.ValueType() != wire.TBinary {
		return nil, nil
	}
	o := make(map[string]struct{}, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Close()
	return o, err
}

func (v *SecondService_EchoStringSet_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TSet {
				v.Arg, err = _Set_String_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of SecondService_EchoStringSet_Args is required")
	}
	return nil
}

func (v *SecondService_EchoStringSet_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("SecondService_EchoStringSet_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Set_String_Equals(lhs, rhs map[string]struct{}) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for x := range rhs {
		if _, ok := lhs[x]; !ok {
			return false
		}
	}
	return true
}

func (v *SecondService_EchoStringSet_Args) Equals(rhs *SecondService_EchoStringSet_Args) bool {
	if !_Set_String_Equals(v.Arg, rhs.Arg) {
		return false
	}
	return true
}

func (v *SecondService_EchoStringSet_Args) MethodName() string {
	return "echoStringSet"
}

func (v *SecondService_EchoStringSet_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_EchoStringSet_Helper = struct {
	Args           func(arg map[string]struct{}) *SecondService_EchoStringSet_Args
	IsException    func(error) bool
	WrapResponse   func(map[string]struct{}, error) (*SecondService_EchoStringSet_Result, error)
	UnwrapResponse func(*SecondService_EchoStringSet_Result) (map[string]struct{}, error)
}{}

func init() {
	SecondService_EchoStringSet_Helper.Args = func(arg map[string]struct{}) *SecondService_EchoStringSet_Args {
		return &SecondService_EchoStringSet_Args{Arg: arg}
	}
	SecondService_EchoStringSet_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_EchoStringSet_Helper.WrapResponse = func(success map[string]struct{}, err error) (*SecondService_EchoStringSet_Result, error) {
		if err == nil {
			return &SecondService_EchoStringSet_Result{Success: success}, nil
		}
		return nil, err
	}
	SecondService_EchoStringSet_Helper.UnwrapResponse = func(result *SecondService_EchoStringSet_Result) (success map[string]struct{}, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SecondService_EchoStringSet_Result struct {
	Success map[string]struct{} `json:"success"`
}

func (v *SecondService_EchoStringSet_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueSet(_Set_String_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoStringSet_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_EchoStringSet_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TSet {
				v.Success, err = _Set_String_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SecondService_EchoStringSet_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SecondService_EchoStringSet_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("SecondService_EchoStringSet_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_EchoStringSet_Result) Equals(rhs *SecondService_EchoStringSet_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Set_String_Equals(v.Success, rhs.Success))) {
		return false
	}
	return true
}

func (v *SecondService_EchoStringSet_Result) MethodName() string {
	return "echoStringSet"
}

func (v *SecondService_EchoStringSet_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
