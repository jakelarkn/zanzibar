// Code generated by zanzibar
// @generated
// Checksum : BkOksA9FmMnjkULpufmE1Q==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop(in *jlexer.Lexer, out *SimpleService_SillyNoop_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "authErr":
			if in.IsNull() {
				in.Skip()
				out.AuthErr = nil
			} else {
				if out.AuthErr == nil {
					out.AuthErr = new(AuthErr)
				}
				easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBaz(in, &*out.AuthErr)
			}
		case "serverErr":
			if in.IsNull() {
				in.Skip()
				out.ServerErr = nil
			} else {
				if out.ServerErr == nil {
					out.ServerErr = new(base.ServerErr)
				}
				(*out.ServerErr).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop(out *jwriter.Writer, in SimpleService_SillyNoop_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AuthErr != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"authErr\":")
		if in.AuthErr == nil {
			out.RawString("null")
		} else {
			easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBaz(out, *in.AuthErr)
		}
	}
	if in.ServerErr != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"serverErr\":")
		if in.ServerErr == nil {
			out.RawString("null")
		} else {
			(*in.ServerErr).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_SillyNoop_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_SillyNoop_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_SillyNoop_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_SillyNoop_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop(l, v)
}
func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBaz(in *jlexer.Lexer, out *AuthErr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MessageSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBaz(out *jwriter.Writer, in AuthErr) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	out.RawByte('}')
}
func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop1(in *jlexer.Lexer, out *SimpleService_SillyNoop_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop1(out *jwriter.Writer, in SimpleService_SillyNoop_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_SillyNoop_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_SillyNoop_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_SillyNoop_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_SillyNoop_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServiceSillyNoop1(l, v)
}
