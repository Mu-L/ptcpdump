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

// loadBpf_legacy returns the embedded CollectionSpec for bpf_legacy.
func loadBpf_legacy() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_legacyBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_legacy: %w", err)
	}

	return spec, err
}

// loadBpf_legacyObjects loads bpf_legacy and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_legacyObjects
//	*bpf_legacyPrograms
//	*bpf_legacyMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_legacyObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_legacy()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_legacySpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_legacySpecs struct {
	bpf_legacyProgramSpecs
	bpf_legacyMapSpecs
}

// bpf_legacySpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_legacyProgramSpecs struct {
	KprobeNfNatManipPkt              *ebpf.ProgramSpec `ebpf:"kprobe__nf_nat_manip_pkt"`
	KprobeNfNatPacket                *ebpf.ProgramSpec `ebpf:"kprobe__nf_nat_packet"`
	KprobeSecuritySkClassifyFlow     *ebpf.ProgramSpec `ebpf:"kprobe__security_sk_classify_flow"`
	KprobeTcpSendmsg                 *ebpf.ProgramSpec `ebpf:"kprobe__tcp_sendmsg"`
	KprobeUdpSendSkb                 *ebpf.ProgramSpec `ebpf:"kprobe__udp_send_skb"`
	KprobeUdpSendmsg                 *ebpf.ProgramSpec `ebpf:"kprobe__udp_sendmsg"`
	RawTracepointSchedProcessExec    *ebpf.ProgramSpec `ebpf:"raw_tracepoint__sched_process_exec"`
	RawTracepointSchedProcessExit    *ebpf.ProgramSpec `ebpf:"raw_tracepoint__sched_process_exit"`
	RawTracepointSchedProcessFork    *ebpf.ProgramSpec `ebpf:"raw_tracepoint__sched_process_fork"`
	TcEgress                         *ebpf.ProgramSpec `ebpf:"tc_egress"`
	TcIngress                        *ebpf.ProgramSpec `ebpf:"tc_ingress"`
	UprobeGoBuiltinTlsWriteKeyLog    *ebpf.ProgramSpec `ebpf:"uprobe__go_builtin__tls__write_key_log"`
	UprobeGoBuiltinTlsWriteKeyLogRet *ebpf.ProgramSpec `ebpf:"uprobe__go_builtin__tls__write_key_log__ret"`
}

// bpf_legacyMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_legacyMapSpecs struct {
	ConfigMap           *ebpf.MapSpec `ebpf:"config_map"`
	ExecEventStack      *ebpf.MapSpec `ebpf:"exec_event_stack"`
	ExecEvents          *ebpf.MapSpec `ebpf:"exec_events"`
	ExitEvents          *ebpf.MapSpec `ebpf:"exit_events"`
	FilterByKernelCount *ebpf.MapSpec `ebpf:"filter_by_kernel_count"`
	FilterMntnsMap      *ebpf.MapSpec `ebpf:"filter_mntns_map"`
	FilterNetnsMap      *ebpf.MapSpec `ebpf:"filter_netns_map"`
	FilterPidMap        *ebpf.MapSpec `ebpf:"filter_pid_map"`
	FilterPidnsMap      *ebpf.MapSpec `ebpf:"filter_pidns_map"`
	FlowPidMap          *ebpf.MapSpec `ebpf:"flow_pid_map"`
	GoKeylogBufStorage  *ebpf.MapSpec `ebpf:"go_keylog_buf_storage"`
	GoKeylogEvents      *ebpf.MapSpec `ebpf:"go_keylog_events"`
	NatFlowMap          *ebpf.MapSpec `ebpf:"nat_flow_map"`
	PacketEventStack    *ebpf.MapSpec `ebpf:"packet_event_stack"`
	PacketEvents        *ebpf.MapSpec `ebpf:"packet_events"`
	SockCookiePidMap    *ebpf.MapSpec `ebpf:"sock_cookie_pid_map"`
}

// bpf_legacyObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_legacyObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_legacyObjects struct {
	bpf_legacyPrograms
	bpf_legacyMaps
}

func (o *bpf_legacyObjects) Close() error {
	return _Bpf_legacyClose(
		&o.bpf_legacyPrograms,
		&o.bpf_legacyMaps,
	)
}

// bpf_legacyMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_legacyObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_legacyMaps struct {
	ConfigMap           *ebpf.Map `ebpf:"config_map"`
	ExecEventStack      *ebpf.Map `ebpf:"exec_event_stack"`
	ExecEvents          *ebpf.Map `ebpf:"exec_events"`
	ExitEvents          *ebpf.Map `ebpf:"exit_events"`
	FilterByKernelCount *ebpf.Map `ebpf:"filter_by_kernel_count"`
	FilterMntnsMap      *ebpf.Map `ebpf:"filter_mntns_map"`
	FilterNetnsMap      *ebpf.Map `ebpf:"filter_netns_map"`
	FilterPidMap        *ebpf.Map `ebpf:"filter_pid_map"`
	FilterPidnsMap      *ebpf.Map `ebpf:"filter_pidns_map"`
	FlowPidMap          *ebpf.Map `ebpf:"flow_pid_map"`
	GoKeylogBufStorage  *ebpf.Map `ebpf:"go_keylog_buf_storage"`
	GoKeylogEvents      *ebpf.Map `ebpf:"go_keylog_events"`
	NatFlowMap          *ebpf.Map `ebpf:"nat_flow_map"`
	PacketEventStack    *ebpf.Map `ebpf:"packet_event_stack"`
	PacketEvents        *ebpf.Map `ebpf:"packet_events"`
	SockCookiePidMap    *ebpf.Map `ebpf:"sock_cookie_pid_map"`
}

func (m *bpf_legacyMaps) Close() error {
	return _Bpf_legacyClose(
		m.ConfigMap,
		m.ExecEventStack,
		m.ExecEvents,
		m.ExitEvents,
		m.FilterByKernelCount,
		m.FilterMntnsMap,
		m.FilterNetnsMap,
		m.FilterPidMap,
		m.FilterPidnsMap,
		m.FlowPidMap,
		m.GoKeylogBufStorage,
		m.GoKeylogEvents,
		m.NatFlowMap,
		m.PacketEventStack,
		m.PacketEvents,
		m.SockCookiePidMap,
	)
}

// bpf_legacyPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_legacyObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_legacyPrograms struct {
	KprobeNfNatManipPkt              *ebpf.Program `ebpf:"kprobe__nf_nat_manip_pkt"`
	KprobeNfNatPacket                *ebpf.Program `ebpf:"kprobe__nf_nat_packet"`
	KprobeSecuritySkClassifyFlow     *ebpf.Program `ebpf:"kprobe__security_sk_classify_flow"`
	KprobeTcpSendmsg                 *ebpf.Program `ebpf:"kprobe__tcp_sendmsg"`
	KprobeUdpSendSkb                 *ebpf.Program `ebpf:"kprobe__udp_send_skb"`
	KprobeUdpSendmsg                 *ebpf.Program `ebpf:"kprobe__udp_sendmsg"`
	RawTracepointSchedProcessExec    *ebpf.Program `ebpf:"raw_tracepoint__sched_process_exec"`
	RawTracepointSchedProcessExit    *ebpf.Program `ebpf:"raw_tracepoint__sched_process_exit"`
	RawTracepointSchedProcessFork    *ebpf.Program `ebpf:"raw_tracepoint__sched_process_fork"`
	TcEgress                         *ebpf.Program `ebpf:"tc_egress"`
	TcIngress                        *ebpf.Program `ebpf:"tc_ingress"`
	UprobeGoBuiltinTlsWriteKeyLog    *ebpf.Program `ebpf:"uprobe__go_builtin__tls__write_key_log"`
	UprobeGoBuiltinTlsWriteKeyLogRet *ebpf.Program `ebpf:"uprobe__go_builtin__tls__write_key_log__ret"`
}

func (p *bpf_legacyPrograms) Close() error {
	return _Bpf_legacyClose(
		p.KprobeNfNatManipPkt,
		p.KprobeNfNatPacket,
		p.KprobeSecuritySkClassifyFlow,
		p.KprobeTcpSendmsg,
		p.KprobeUdpSendSkb,
		p.KprobeUdpSendmsg,
		p.RawTracepointSchedProcessExec,
		p.RawTracepointSchedProcessExit,
		p.RawTracepointSchedProcessFork,
		p.TcEgress,
		p.TcIngress,
		p.UprobeGoBuiltinTlsWriteKeyLog,
		p.UprobeGoBuiltinTlsWriteKeyLogRet,
	)
}

func _Bpf_legacyClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_legacy_x86_bpfel.o
var _Bpf_legacyBytes []byte
