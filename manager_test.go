package manager

import (
	"errors"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/rlimit"
	"golang.org/x/sys/unix"
)

func TestVerifierError(t *testing.T) {
	err := rlimit.RemoveMemlock()
	if err != nil {
		t.Fatal(err)
	}
	m := &Manager{
		collectionSpec: &ebpf.CollectionSpec{
			Programs: map[string]*ebpf.ProgramSpec{"socket__filter": {
				Type: ebpf.SocketFilter,
				Instructions: asm.Instructions{
					asm.LoadImm(asm.R0, 0, asm.DWord),
					// Missing Return
				},
				License: "MIT",
			}},
		},
	}
	err = m.loadCollection()
	if err == nil {
		t.Fatal("expected error")
	}
	var ve *ebpf.VerifierError
	if !errors.As(err, &ve) {
		t.Fatal("expected to be able to unwrap to VerifierError")
	}
	if strings.Count(err.Error(), "\n") == 0 {
		t.Fatal("expected full verifier error")
	}
}

func TestExclude(t *testing.T) {
	err := rlimit.RemoveMemlock()
	if err != nil {
		t.Fatal(err)
	}

	m := &Manager{
		Probes: []*Probe{
			{ProbeIdentificationPair: ProbeIdentificationPair{EBPFFuncName: "access_map_one"}},
		},
		Maps: []*Map{
			{Name: "map_one"},
		},
	}
	opts := Options{
		RLimit: &unix.Rlimit{
			Cur: math.MaxUint64,
			Max: math.MaxUint64,
		},
		ExcludedMaps: []string{"map_two"},
	}

	f, err := os.Open("testdata/exclude.elf")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = f.Close() })
	err = m.InitWithOptions(f, opts)
	if err == nil || !strings.Contains(err.Error(), "missing map map_two") {
		t.Fatalf("expected error about missing map map_two, got `%s` instead", err)
	}

	opts.ExcludedFunctions = []string{"access_map_two"}
	err = m.InitWithOptions(f, opts)
	if err != nil {
		t.Fatal(err)
	}
}
