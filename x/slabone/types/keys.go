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

const (
    SlabCreatedEventOriginator =	"EvCreatedOriginator"
	SlabCreatedEventOriginatorSocialId = "EvCreatedOriginatorSocialId"      
    SlabCreatedEventSlabId 			   ="EvCreatedSlabId"       
    SlabCreatedEventAssertion 		="EvCreatedAssertion"            
	SlabCreatedEventDirectedTowards ="EvCreatedDirectedTowards"
	SlabCreatedEventSlabState		= "EvCreatedSlabState"   
	SlabCreatedEventSlabCtxTime		= "EvCreatedSlabCtxTime"
	
	SlabVettedEventVetter 			= "EvVettedVetter"
	SlabVettedEventVetterSocialId   = "EvVettedVetterSocialId"
	SlabVettedEventSlabState 		= "EvVettedSlabState"
	SlabVettedEventSlabVettingNote	= "EvVettedVettingNote"
	SlabVettedEventSlabId 			= "EvVettedSlabId" 
	SlabVettedEventSlabCtxTime		= "EvVettedSlabCtxTime"

	SlabRevokedEventRevoker			= "EvRevokedRevoker"
	SlabRevokedEventRevokerSocialId = "EvRevokedRevokerSocialId"
	SlabRevokedEventSlabState		= "EvRevokedSlabState"
	SlabRevokedEventSlabRevokingNote= "EvRevokedRevokingNote"
	SlabRevokedEventSlabId			= "EvRevokedSlabId"
	SlabRevokedEventSlabCtxTime		= "EvRevokedSlabCtxTime"
)	
