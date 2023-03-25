package types

const (
	// ModuleName defines the module name
	ModuleName = "slabone"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_slabone"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SlabKey      = "Slab/value/"
	SlabCountKey = "Slab/count/"
)
