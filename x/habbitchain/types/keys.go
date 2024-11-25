package types

const (
	// ModuleName defines the module name
	ModuleName = "habbitchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_habbitchain"
)

var (
	ParamsKey = []byte("p_habbitchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
