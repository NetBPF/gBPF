// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/khulnasoft/gbpf"
)

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*gbpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := gbpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See gbpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *gbpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed gbpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed gbpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	XdpProgFunc *gbpf.ProgramSpec `gbpf:"xdp_prog_func"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed gbpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	XdpStatsMap *gbpf.MapSpec `gbpf:"xdp_stats_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or gbpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or gbpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	XdpStatsMap *gbpf.Map `gbpf:"xdp_stats_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.XdpStatsMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or gbpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	XdpProgFunc *gbpf.Program `gbpf:"xdp_prog_func"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.XdpProgFunc,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfeb.o
var _BpfBytes []byte
