package types

const (
	// ModuleName defines the module name
	ModuleName = "parrots"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_parrots"

	// ProfileKey
	ProfileKey = "Profile-value-"

	// ProfileCountKey
	ProfileCountKey = "Profile-count-"

	// BeakKey
	BeakKey = "Beak-value-"

	// BeakCountKey
	BeakCountKey = "Beak-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
