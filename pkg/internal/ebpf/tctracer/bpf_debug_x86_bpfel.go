// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package tctracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_debugHttpInfoT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_debugConnectionInfoT
	_               [2]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [192]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
	TaskTid uint32
	_       [4]byte
}

type bpf_debugPidConnectionInfoT struct {
	Conn bpf_debugConnectionInfoT
	Pid  uint32
}

type bpf_debugTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid   uint32
	Valid uint8
	_     [3]byte
}

// loadBpf_debug returns the embedded CollectionSpec for bpf_debug.
func loadBpf_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_debug: %w", err)
	}

	return spec, err
}

// loadBpf_debugObjects loads bpf_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_debugObjects
//	*bpf_debugPrograms
//	*bpf_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugSpecs struct {
	bpf_debugProgramSpecs
	bpf_debugMapSpecs
}

// bpf_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugProgramSpecs struct {
	AppEgress  *ebpf.ProgramSpec `ebpf:"app_egress"`
	AppIngress *ebpf.ProgramSpec `ebpf:"app_ingress"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	DebugEvents         *ebpf.MapSpec `ebpf:"debug_events"`
	IncomingTraceMap    *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	OngoingHttp         *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttpFallback *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	OutgoingTraceMap    *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	TraceMap            *ebpf.MapSpec `ebpf:"trace_map"`
}

// bpf_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugObjects struct {
	bpf_debugPrograms
	bpf_debugMaps
}

func (o *bpf_debugObjects) Close() error {
	return _Bpf_debugClose(
		&o.bpf_debugPrograms,
		&o.bpf_debugMaps,
	)
}

// bpf_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugMaps struct {
	DebugEvents         *ebpf.Map `ebpf:"debug_events"`
	IncomingTraceMap    *ebpf.Map `ebpf:"incoming_trace_map"`
	OngoingHttp         *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttpFallback *ebpf.Map `ebpf:"ongoing_http_fallback"`
	OutgoingTraceMap    *ebpf.Map `ebpf:"outgoing_trace_map"`
	TraceMap            *ebpf.Map `ebpf:"trace_map"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.DebugEvents,
		m.IncomingTraceMap,
		m.OngoingHttp,
		m.OngoingHttpFallback,
		m.OutgoingTraceMap,
		m.TraceMap,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	AppEgress  *ebpf.Program `ebpf:"app_egress"`
	AppIngress *ebpf.Program `ebpf:"app_ingress"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.AppEgress,
		p.AppIngress,
	)
}

func _Bpf_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_debug_x86_bpfel.o
var _Bpf_debugBytes []byte