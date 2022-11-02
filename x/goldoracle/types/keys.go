package types

const (
	// ModuleName defines the module name
	ModuleName = "goldoracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_goldoracle"

	// Version defines the current version the IBC module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to
	PortID = "goldoracle"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("goldoracle-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
