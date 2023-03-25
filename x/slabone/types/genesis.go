package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SlabList: []Slab{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in slab
	slabIdMap := make(map[uint64]bool)
	slabCount := gs.GetSlabCount()
	for _, elem := range gs.SlabList {
		if _, ok := slabIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for slab")
		}
		if elem.Id >= slabCount {
			return fmt.Errorf("slab id should be lower or equal than the last id")
		}
		slabIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
