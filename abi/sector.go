package abi

import (
	"fmt"
	"math"
	"strconv"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/big"
)

// SectorNumber is a numeric identifier for a sector. It is usually relative to a miner.
type SectorNumber uint64

func (s SectorNumber) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// The maximum assignable sector number.
// Raising this would require modifying our AMT implementation.
const MaxSectorNumber = math.MaxInt64

// SectorSize indicates one of a set of possible sizes in the network.
// Ideally, SectorSize would be an enum
// type SectorSize enum {
//   1KiB = 1024
//   1MiB = 1048576
//   1GiB = 1073741824
//   1TiB = 1099511627776
//   1PiB = 1125899906842624
//   1EiB = 1152921504606846976
//   max  = 18446744073709551615
// }
type SectorSize uint64

// Formats the size as a decimal string.
func (s SectorSize) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// Abbreviates the size as a human-scale number.
// This approximates (truncates) the size unless it is a power of 1024.
func (s SectorSize) ShortString() string {
	var biUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	unit := 0
	for s >= 1024 && unit < len(biUnits)-1 {
		s /= 1024
		unit++
	}
	return fmt.Sprintf("%d%s", s, biUnits[unit])
}

type SectorID struct {
	Miner  ActorID
	Number SectorNumber
}

// The unit of storage power (measured in bytes)
type StoragePower = big.Int

type SectorQuality = big.Int

func NewStoragePower(n int64) StoragePower {
	return big.NewInt(n)
}

// These enumerations must match the proofs library and never change.
type RegisteredSealProof int64

const (
	RegisteredSealProof_StackedDrg2KiBV1   = RegisteredSealProof(0)
	RegisteredSealProof_StackedDrg8MiBV1   = RegisteredSealProof(1)
	RegisteredSealProof_StackedDrg512MiBV1 = RegisteredSealProof(2)
	RegisteredSealProof_StackedDrg32GiBV1  = RegisteredSealProof(3)
	RegisteredSealProof_StackedDrg64GiBV1  = RegisteredSealProof(4)

	RegisteredSealProof_StackedDrg2KiBV1_1   = RegisteredSealProof(5)
	RegisteredSealProof_StackedDrg8MiBV1_1   = RegisteredSealProof(6)
	RegisteredSealProof_StackedDrg512MiBV1_1 = RegisteredSealProof(7)
	RegisteredSealProof_StackedDrg32GiBV1_1  = RegisteredSealProof(8)
	RegisteredSealProof_StackedDrg64GiBV1_1  = RegisteredSealProof(9)
)

type RegisteredPoStProof int64

const (
	RegisteredPoStProof_StackedDrgWinning2KiBV1   = RegisteredPoStProof(0)
	RegisteredPoStProof_StackedDrgWinning8MiBV1   = RegisteredPoStProof(1)
	RegisteredPoStProof_StackedDrgWinning512MiBV1 = RegisteredPoStProof(2)
	RegisteredPoStProof_StackedDrgWinning32GiBV1  = RegisteredPoStProof(3)
	RegisteredPoStProof_StackedDrgWinning64GiBV1  = RegisteredPoStProof(4)
	RegisteredPoStProof_StackedDrgWindow2KiBV1    = RegisteredPoStProof(5)
	RegisteredPoStProof_StackedDrgWindow8MiBV1    = RegisteredPoStProof(6)
	RegisteredPoStProof_StackedDrgWindow512MiBV1  = RegisteredPoStProof(7)
	RegisteredPoStProof_StackedDrgWindow32GiBV1   = RegisteredPoStProof(8)
	RegisteredPoStProof_StackedDrgWindow64GiBV1   = RegisteredPoStProof(9)
)

// Metadata about a seal proof type.
type SealProofInfo struct {
	SectorSize       SectorSize
	WinningPoStProof RegisteredPoStProof
	WindowPoStProof  RegisteredPoStProof
}

const (
	ss2KiB   = 2 << 10
	ss8MiB   = 8 << 20
	ss512MiB = 512 << 20
	ss32GiB  = 32 << 30
	ss64GiB  = 64 << 30
)

var SealProofInfos = map[RegisteredSealProof]*SealProofInfo{
	RegisteredSealProof_StackedDrg2KiBV1: {
		SectorSize:       ss2KiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning2KiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow2KiBV1,
	},
	RegisteredSealProof_StackedDrg8MiBV1: {
		SectorSize:       ss8MiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning8MiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow8MiBV1,
	},
	RegisteredSealProof_StackedDrg512MiBV1: {
		SectorSize:       ss512MiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning512MiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow512MiBV1,
	},
	RegisteredSealProof_StackedDrg32GiBV1: {
		SectorSize:       ss32GiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning32GiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow32GiBV1,
	},
	RegisteredSealProof_StackedDrg64GiBV1: {
		SectorSize:       ss64GiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning64GiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow64GiBV1,
	},

	RegisteredSealProof_StackedDrg2KiBV1_1: {
		SectorSize:       ss2KiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning2KiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow2KiBV1,
	},
	RegisteredSealProof_StackedDrg8MiBV1_1: {
		SectorSize:       ss8MiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning8MiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow8MiBV1,
	},
	RegisteredSealProof_StackedDrg512MiBV1_1: {
		SectorSize:       ss512MiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning512MiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow512MiBV1,
	},
	RegisteredSealProof_StackedDrg32GiBV1_1: {
		SectorSize:       ss32GiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning32GiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow32GiBV1,
	},
	RegisteredSealProof_StackedDrg64GiBV1_1: {
		SectorSize:       ss64GiB,
		WinningPoStProof: RegisteredPoStProof_StackedDrgWinning64GiBV1,
		WindowPoStProof:  RegisteredPoStProof_StackedDrgWindow64GiBV1,
	},
}

func (p RegisteredSealProof) SectorSize() (SectorSize, error) {
	info, ok := SealProofInfos[p]
	if !ok {
		return 0, xerrors.Errorf("unsupported proof type: %v", p)
	}
	return info.SectorSize, nil
}

// RegisteredWinningPoStProof produces the PoSt-specific RegisteredProof corresponding
// to the receiving RegisteredProof.
func (p RegisteredSealProof) RegisteredWinningPoStProof() (RegisteredPoStProof, error) {
	info, ok := SealProofInfos[p]
	if !ok {
		return 0, xerrors.Errorf("unsupported proof type: %v", p)
	}
	return info.WinningPoStProof, nil
}

// RegisteredWindowPoStProof produces the PoSt-specific RegisteredProof corresponding
// to the receiving RegisteredProof.
func (p RegisteredSealProof) RegisteredWindowPoStProof() (RegisteredPoStProof, error) {
	info, ok := SealProofInfos[p]
	if !ok {
		return 0, xerrors.Errorf("unsupported proof type: %v", p)
	}
	return info.WindowPoStProof, nil
}

// Metadata about a PoSt proof type.
type PoStProofInfo struct {
	SectorSize SectorSize
}

var PoStProofInfos = map[RegisteredPoStProof]*PoStProofInfo{
	RegisteredPoStProof_StackedDrgWinning2KiBV1: {
		SectorSize: ss2KiB,
	},
	RegisteredPoStProof_StackedDrgWinning8MiBV1: {
		SectorSize: ss8MiB,
	},
	RegisteredPoStProof_StackedDrgWinning512MiBV1: {
		SectorSize: ss512MiB,
	},
	RegisteredPoStProof_StackedDrgWinning32GiBV1: {
		SectorSize: ss32GiB,
	},
	RegisteredPoStProof_StackedDrgWinning64GiBV1: {
		SectorSize: ss64GiB,
	},
	RegisteredPoStProof_StackedDrgWindow2KiBV1: {
		SectorSize: ss2KiB,
	},
	RegisteredPoStProof_StackedDrgWindow8MiBV1: {
		SectorSize: ss8MiB,
	},
	RegisteredPoStProof_StackedDrgWindow512MiBV1: {
		SectorSize: ss512MiB,
	},
	RegisteredPoStProof_StackedDrgWindow32GiBV1: {
		SectorSize: ss32GiB,
	},
	RegisteredPoStProof_StackedDrgWindow64GiBV1: {
		SectorSize: ss64GiB,
	},
}

func (p RegisteredPoStProof) SectorSize() (SectorSize, error) {
	info, ok := PoStProofInfos[p]
	if !ok {
		return 0, xerrors.Errorf("unsupported proof type: %v", p)
	}
	return info.SectorSize, nil
}

type SealRandomness Randomness
type InteractiveSealRandomness Randomness
type PoStRandomness Randomness
