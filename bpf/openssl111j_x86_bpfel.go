// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type Openssl111jConnEvtT struct {
	ConnInfo Openssl111jConnInfoT
	ConnType Openssl111jConnTypeT
	_        [4]byte
	Ts       uint64
}

type Openssl111jConnIdS_t struct {
	TgidFd  uint64
	NoTrace bool
	_       [7]byte
}

type Openssl111jConnInfoT struct {
	ConnId struct {
		Upid struct {
			Pid            uint32
			_              [4]byte
			StartTimeTicks uint64
		}
		Fd   int32
		_    [4]byte
		Tsid uint64
	}
	ReadBytes     uint64
	WriteBytes    uint64
	SslReadBytes  uint64
	SslWriteBytes uint64
	Laddr         struct {
		In6 struct {
			Sin6Family   uint16
			Sin6Port     uint16
			Sin6Flowinfo uint32
			Sin6Addr     struct{ In6U struct{ U6Addr8 [16]uint8 } }
			Sin6ScopeId  uint32
		}
	}
	Raddr struct {
		In6 struct {
			Sin6Family   uint16
			Sin6Port     uint16
			Sin6Flowinfo uint32
			Sin6Addr     struct{ In6U struct{ U6Addr8 [16]uint8 } }
			Sin6ScopeId  uint32
		}
	}
	Protocol            Openssl111jTrafficProtocolT
	Role                Openssl111jEndpointRoleT
	PrevCount           uint64
	PrevBuf             [4]int8
	PrependLengthHeader bool
	NoTrace             bool
	Ssl                 bool
	_                   [1]byte
}

type Openssl111jConnTypeT uint32

const (
	Openssl111jConnTypeTKConnect       Openssl111jConnTypeT = 0
	Openssl111jConnTypeTKClose         Openssl111jConnTypeT = 1
	Openssl111jConnTypeTKProtocolInfer Openssl111jConnTypeT = 2
)

type Openssl111jControlValueIndexT uint32

const (
	Openssl111jControlValueIndexTKTargetTGIDIndex   Openssl111jControlValueIndexT = 0
	Openssl111jControlValueIndexTKStirlingTGIDIndex Openssl111jControlValueIndexT = 1
	Openssl111jControlValueIndexTKEnabledXdpIndex   Openssl111jControlValueIndexT = 2
	Openssl111jControlValueIndexTKNumControlValues  Openssl111jControlValueIndexT = 3
)

type Openssl111jEndpointRoleT uint32

const (
	Openssl111jEndpointRoleTKRoleClient  Openssl111jEndpointRoleT = 1
	Openssl111jEndpointRoleTKRoleServer  Openssl111jEndpointRoleT = 2
	Openssl111jEndpointRoleTKRoleUnknown Openssl111jEndpointRoleT = 4
)

type Openssl111jKernEvt struct {
	FuncName [16]int8
	Ts       uint64
	Seq      uint64
	Len      uint32
	Flags    uint8
	_        [3]byte
	ConnIdS  Openssl111jConnIdS_t
	IsSample int32
	Step     Openssl111jStepT
}

type Openssl111jKernEvtData struct {
	Ke      Openssl111jKernEvt
	BufSize uint32
	Msg     [30720]int8
	_       [4]byte
}

type Openssl111jSockKey struct {
	Sip   [2]uint64
	Dip   [2]uint64
	Sport uint16
	Dport uint16
	_     [4]byte
}

type Openssl111jStepT uint32

const (
	Openssl111jStepTStart       Openssl111jStepT = 0
	Openssl111jStepTSSL_OUT     Openssl111jStepT = 1
	Openssl111jStepTSYSCALL_OUT Openssl111jStepT = 2
	Openssl111jStepTTCP_OUT     Openssl111jStepT = 3
	Openssl111jStepTIP_OUT      Openssl111jStepT = 4
	Openssl111jStepTQDISC_OUT   Openssl111jStepT = 5
	Openssl111jStepTDEV_OUT     Openssl111jStepT = 6
	Openssl111jStepTNIC_OUT     Openssl111jStepT = 7
	Openssl111jStepTNIC_IN      Openssl111jStepT = 8
	Openssl111jStepTDEV_IN      Openssl111jStepT = 9
	Openssl111jStepTIP_IN       Openssl111jStepT = 10
	Openssl111jStepTTCP_IN      Openssl111jStepT = 11
	Openssl111jStepTUSER_COPY   Openssl111jStepT = 12
	Openssl111jStepTSYSCALL_IN  Openssl111jStepT = 13
	Openssl111jStepTSSL_IN      Openssl111jStepT = 14
	Openssl111jStepTEnd         Openssl111jStepT = 15
)

type Openssl111jTrafficDirectionT uint32

const (
	Openssl111jTrafficDirectionTKEgress  Openssl111jTrafficDirectionT = 0
	Openssl111jTrafficDirectionTKIngress Openssl111jTrafficDirectionT = 1
)

type Openssl111jTrafficProtocolT uint32

const (
	Openssl111jTrafficProtocolTKProtocolUnset   Openssl111jTrafficProtocolT = 0
	Openssl111jTrafficProtocolTKProtocolUnknown Openssl111jTrafficProtocolT = 1
	Openssl111jTrafficProtocolTKProtocolHTTP    Openssl111jTrafficProtocolT = 2
	Openssl111jTrafficProtocolTKProtocolHTTP2   Openssl111jTrafficProtocolT = 3
	Openssl111jTrafficProtocolTKProtocolMySQL   Openssl111jTrafficProtocolT = 4
	Openssl111jTrafficProtocolTKProtocolCQL     Openssl111jTrafficProtocolT = 5
	Openssl111jTrafficProtocolTKProtocolPGSQL   Openssl111jTrafficProtocolT = 6
	Openssl111jTrafficProtocolTKProtocolDNS     Openssl111jTrafficProtocolT = 7
	Openssl111jTrafficProtocolTKProtocolRedis   Openssl111jTrafficProtocolT = 8
	Openssl111jTrafficProtocolTKProtocolNATS    Openssl111jTrafficProtocolT = 9
	Openssl111jTrafficProtocolTKProtocolMongo   Openssl111jTrafficProtocolT = 10
	Openssl111jTrafficProtocolTKProtocolKafka   Openssl111jTrafficProtocolT = 11
	Openssl111jTrafficProtocolTKProtocolMux     Openssl111jTrafficProtocolT = 12
	Openssl111jTrafficProtocolTKProtocolAMQP    Openssl111jTrafficProtocolT = 13
	Openssl111jTrafficProtocolTKNumProtocols    Openssl111jTrafficProtocolT = 14
)

// LoadOpenssl111j returns the embedded CollectionSpec for Openssl111j.
func LoadOpenssl111j() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Openssl111jBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Openssl111j: %w", err)
	}

	return spec, err
}

// LoadOpenssl111jObjects loads Openssl111j and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*Openssl111jObjects
//	*Openssl111jPrograms
//	*Openssl111jMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadOpenssl111jObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadOpenssl111j()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// Openssl111jSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl111jSpecs struct {
	Openssl111jProgramSpecs
	Openssl111jMapSpecs
}

// Openssl111jSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl111jProgramSpecs struct {
	SSL_readEntryNestedSyscall    *ebpf.ProgramSpec `ebpf:"SSL_read_entry_nested_syscall"`
	SSL_readEntryOffset           *ebpf.ProgramSpec `ebpf:"SSL_read_entry_offset"`
	SSL_readExEntryNestedSyscall  *ebpf.ProgramSpec `ebpf:"SSL_read_ex_entry_nested_syscall"`
	SSL_readExRetNestedSyscall    *ebpf.ProgramSpec `ebpf:"SSL_read_ex_ret_nested_syscall"`
	SSL_readRetNestedSyscall      *ebpf.ProgramSpec `ebpf:"SSL_read_ret_nested_syscall"`
	SSL_readRetOffset             *ebpf.ProgramSpec `ebpf:"SSL_read_ret_offset"`
	SSL_writeEntryNestedSyscall   *ebpf.ProgramSpec `ebpf:"SSL_write_entry_nested_syscall"`
	SSL_writeEntryOffset          *ebpf.ProgramSpec `ebpf:"SSL_write_entry_offset"`
	SSL_writeExEntryNestedSyscall *ebpf.ProgramSpec `ebpf:"SSL_write_ex_entry_nested_syscall"`
	SSL_writeExRetNestedSyscall   *ebpf.ProgramSpec `ebpf:"SSL_write_ex_ret_nested_syscall"`
	SSL_writeRetNestedSyscall     *ebpf.ProgramSpec `ebpf:"SSL_write_ret_nested_syscall"`
	SSL_writeRetOffset            *ebpf.ProgramSpec `ebpf:"SSL_write_ret_offset"`
}

// Openssl111jMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl111jMapSpecs struct {
	ActiveSslReadArgsMap  *ebpf.MapSpec `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.MapSpec `ebpf:"active_ssl_write_args_map"`
	ConnEvtRb             *ebpf.MapSpec `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.MapSpec `ebpf:"conn_info_map"`
	Rb                    *ebpf.MapSpec `ebpf:"rb"`
	SslDataMap            *ebpf.MapSpec `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.MapSpec `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.MapSpec `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.MapSpec `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.MapSpec `ebpf:"syscall_rb"`
}

// Openssl111jObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl111jObjects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl111jObjects struct {
	Openssl111jPrograms
	Openssl111jMaps
}

func (o *Openssl111jObjects) Close() error {
	return _Openssl111jClose(
		&o.Openssl111jPrograms,
		&o.Openssl111jMaps,
	)
}

// Openssl111jMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl111jObjects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl111jMaps struct {
	ActiveSslReadArgsMap  *ebpf.Map `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.Map `ebpf:"active_ssl_write_args_map"`
	ConnEvtRb             *ebpf.Map `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.Map `ebpf:"conn_info_map"`
	Rb                    *ebpf.Map `ebpf:"rb"`
	SslDataMap            *ebpf.Map `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.Map `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.Map `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.Map `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.Map `ebpf:"syscall_rb"`
}

func (m *Openssl111jMaps) Close() error {
	return _Openssl111jClose(
		m.ActiveSslReadArgsMap,
		m.ActiveSslWriteArgsMap,
		m.ConnEvtRb,
		m.ConnInfoMap,
		m.Rb,
		m.SslDataMap,
		m.SslRb,
		m.SslUserSpaceCallMap,
		m.SyscallDataMap,
		m.SyscallRb,
	)
}

// Openssl111jPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl111jObjects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl111jPrograms struct {
	SSL_readEntryNestedSyscall    *ebpf.Program `ebpf:"SSL_read_entry_nested_syscall"`
	SSL_readEntryOffset           *ebpf.Program `ebpf:"SSL_read_entry_offset"`
	SSL_readExEntryNestedSyscall  *ebpf.Program `ebpf:"SSL_read_ex_entry_nested_syscall"`
	SSL_readExRetNestedSyscall    *ebpf.Program `ebpf:"SSL_read_ex_ret_nested_syscall"`
	SSL_readRetNestedSyscall      *ebpf.Program `ebpf:"SSL_read_ret_nested_syscall"`
	SSL_readRetOffset             *ebpf.Program `ebpf:"SSL_read_ret_offset"`
	SSL_writeEntryNestedSyscall   *ebpf.Program `ebpf:"SSL_write_entry_nested_syscall"`
	SSL_writeEntryOffset          *ebpf.Program `ebpf:"SSL_write_entry_offset"`
	SSL_writeExEntryNestedSyscall *ebpf.Program `ebpf:"SSL_write_ex_entry_nested_syscall"`
	SSL_writeExRetNestedSyscall   *ebpf.Program `ebpf:"SSL_write_ex_ret_nested_syscall"`
	SSL_writeRetNestedSyscall     *ebpf.Program `ebpf:"SSL_write_ret_nested_syscall"`
	SSL_writeRetOffset            *ebpf.Program `ebpf:"SSL_write_ret_offset"`
}

func (p *Openssl111jPrograms) Close() error {
	return _Openssl111jClose(
		p.SSL_readEntryNestedSyscall,
		p.SSL_readEntryOffset,
		p.SSL_readExEntryNestedSyscall,
		p.SSL_readExRetNestedSyscall,
		p.SSL_readRetNestedSyscall,
		p.SSL_readRetOffset,
		p.SSL_writeEntryNestedSyscall,
		p.SSL_writeEntryOffset,
		p.SSL_writeExEntryNestedSyscall,
		p.SSL_writeExRetNestedSyscall,
		p.SSL_writeRetNestedSyscall,
		p.SSL_writeRetOffset,
	)
}

func _Openssl111jClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed openssl111j_x86_bpfel.o
var _Openssl111jBytes []byte
