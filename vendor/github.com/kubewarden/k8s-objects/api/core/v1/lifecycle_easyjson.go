// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *Lifecycle) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "postStart":
			if in.IsNull() {
				in.Skip()
				out.PostStart = nil
			} else {
				if out.PostStart == nil {
					out.PostStart = new(LifecycleHandler)
				}
				easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.PostStart)
			}
		case "preStop":
			if in.IsNull() {
				in.Skip()
				out.PreStop = nil
			} else {
				if out.PreStop == nil {
					out.PreStop = new(LifecycleHandler)
				}
				easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.PreStop)
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
func easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in Lifecycle) {
	out.RawByte('{')
	first := true
	_ = first
	if in.PostStart != nil {
		const prefix string = ",\"postStart\":"
		first = false
		out.RawString(prefix[1:])
		easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.PostStart)
	}
	if in.PreStop != nil {
		const prefix string = ",\"preStop\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.PreStop)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Lifecycle) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Lifecycle) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Lifecycle) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Lifecycle) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *LifecycleHandler) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "exec":
			if in.IsNull() {
				in.Skip()
				out.Exec = nil
			} else {
				if out.Exec == nil {
					out.Exec = new(ExecAction)
				}
				(*out.Exec).UnmarshalEasyJSON(in)
			}
		case "httpGet":
			if in.IsNull() {
				in.Skip()
				out.HTTPGet = nil
			} else {
				if out.HTTPGet == nil {
					out.HTTPGet = new(HTTPGetAction)
				}
				(*out.HTTPGet).UnmarshalEasyJSON(in)
			}
		case "tcpSocket":
			if in.IsNull() {
				in.Skip()
				out.TCPSocket = nil
			} else {
				if out.TCPSocket == nil {
					out.TCPSocket = new(TCPSocketAction)
				}
				easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.TCPSocket)
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
func easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in LifecycleHandler) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Exec != nil {
		const prefix string = ",\"exec\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Exec).MarshalEasyJSON(out)
	}
	if in.HTTPGet != nil {
		const prefix string = ",\"httpGet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.HTTPGet).MarshalEasyJSON(out)
	}
	if in.TCPSocket != nil {
		const prefix string = ",\"tcpSocket\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.TCPSocket)
	}
	out.RawByte('}')
}
func easyjson595d38b2DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *TCPSocketAction) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "host":
			out.Host = string(in.String())
		case "port":
			if in.IsNull() {
				in.Skip()
				out.Port = nil
			} else {
				if out.Port == nil {
					out.Port = new(intstr.IntOrString)
				}
				*out.Port = intstr.IntOrString(in.String())
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
func easyjson595d38b2EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in TCPSocketAction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Host != "" {
		const prefix string = ",\"host\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Port == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Port))
		}
	}
	out.RawByte('}')
}