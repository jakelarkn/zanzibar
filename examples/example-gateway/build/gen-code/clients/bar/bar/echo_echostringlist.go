// Code generated by thriftrw v1.3.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Echo_EchoStringList_Args struct {
	Arg []string `json:"arg,required"`
}

func (v *Echo_EchoStringList_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg == nil {
		return w, errors.New("field Arg of Echo_EchoStringList_Args is required")
	}
	w, err = wire.NewValueList(_List_String_ValueList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Echo_EchoStringList_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Arg, err = _List_String_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of Echo_EchoStringList_Args is required")
	}
	return nil
}

func (v *Echo_EchoStringList_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("Echo_EchoStringList_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_EchoStringList_Args) Equals(rhs *Echo_EchoStringList_Args) bool {
	if !_List_String_Equals(v.Arg, rhs.Arg) {
		return false
	}
	return true
}

func (v *Echo_EchoStringList_Args) MethodName() string {
	return "echoStringList"
}

func (v *Echo_EchoStringList_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Echo_EchoStringList_Helper = struct {
	Args           func(arg []string) *Echo_EchoStringList_Args
	IsException    func(error) bool
	WrapResponse   func([]string, error) (*Echo_EchoStringList_Result, error)
	UnwrapResponse func(*Echo_EchoStringList_Result) ([]string, error)
}{}

func init() {
	Echo_EchoStringList_Helper.Args = func(arg []string) *Echo_EchoStringList_Args {
		return &Echo_EchoStringList_Args{Arg: arg}
	}
	Echo_EchoStringList_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Echo_EchoStringList_Helper.WrapResponse = func(success []string, err error) (*Echo_EchoStringList_Result, error) {
		if err == nil {
			return &Echo_EchoStringList_Result{Success: success}, nil
		}
		return nil, err
	}
	Echo_EchoStringList_Helper.UnwrapResponse = func(result *Echo_EchoStringList_Result) (success []string, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Echo_EchoStringList_Result struct {
	Success []string `json:"success"`
}

func (v *Echo_EchoStringList_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueList(_List_String_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_EchoStringList_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Echo_EchoStringList_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_String_Read(field.Value.GetList())
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
		return fmt.Errorf("Echo_EchoStringList_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Echo_EchoStringList_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("Echo_EchoStringList_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_EchoStringList_Result) Equals(rhs *Echo_EchoStringList_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _List_String_Equals(v.Success, rhs.Success))) {
		return false
	}
	return true
}

func (v *Echo_EchoStringList_Result) MethodName() string {
	return "echoStringList"
}

func (v *Echo_EchoStringList_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
