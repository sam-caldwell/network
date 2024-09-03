package core

// DevlinkAttr - devlink_attr - represents the attributes of devlink as defined in linux/devlink.h
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkAttr uint8

/*
don't change the order or add anything between, this is ABI!
*/
const (
	// DevlinkAttrUnspec - DEVLINK_ATTR_UNSPEC - unspecified attribute
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrUnspec DevlinkAttr = 0

	// DevlinkAttrBusName - DEVLINK_ATTR_BUS_NAME (string) bus name + dev name together are a handle for devlink entity
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrBusName DevlinkAttr = 1

	// DevlinkAttrDevName - DEVLINK_ATTR_DEV_NAME (string) device name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDevName DevlinkAttr = 2

	// DevlinkAttrPortIndex - DEVLINK_ATTR_PORT_INDEX (u32) port index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortIndex DevlinkAttr = 3

	// DevlinkAttrPortType - DEVLINK_ATTR_PORT_TYPE (u16) port type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortType DevlinkAttr = 4

	// DevlinkAttrPortDesiredType - DEVLINK_ATTR_PORT_DESIRED_TYPE (u16) port desired type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortDesiredType DevlinkAttr = 5

	// DevlinkAttrPortNetdevIfindex - DEVLINK_ATTR_PORT_NETDEV_IFINDEX (u32) port netdev ifindex
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortNetdevIfindex DevlinkAttr = 6

	// DevlinkAttrPortNetdevName - DEVLINK_ATTR_PORT_NETDEV_NAME (string) port netdev name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortNetdevName DevlinkAttr = 7

	// DevlinkAttrPortIbdevName - DEVLINK_ATTR_PORT_IBDEV_NAME (string) port ibdev name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortIbdevName DevlinkAttr = 8

	// DevlinkAttrPortSplitCount - DEVLINK_ATTR_PORT_SPLIT_COUNT (u32) port split count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortSplitCount DevlinkAttr = 9

	// DevlinkAttrPortSplitGroup - DEVLINK_ATTR_PORT_SPLIT_GROUP (u32) port split group
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortSplitGroup DevlinkAttr = 10

	// DevlinkAttrSbIndex - DEVLINK_ATTR_SB_INDEX (u32) shared buffer index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbIndex DevlinkAttr = 11

	// DevlinkAttrSbSize - DEVLINK_ATTR_SB_SIZE (u32) shared buffer size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbSize DevlinkAttr = 12

	// DevlinkAttrSbIngressPoolCount - DEVLINK_ATTR_SB_INGRESS_POOL_COUNT (u16) shared buffer ingress pool count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbIngressPoolCount DevlinkAttr = 13

	// DevlinkAttrSbEgressPoolCount - DEVLINK_ATTR_SB_EGRESS_POOL_COUNT (u16) shared buffer egress pool count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbEgressPoolCount DevlinkAttr = 14

	// DevlinkAttrSbIngressTcCount - DEVLINK_ATTR_SB_INGRESS_TC_COUNT (u16) shared buffer ingress traffic class count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbIngressTcCount DevlinkAttr = 15

	// DevlinkAttrSbEgressTcCount - DEVLINK_ATTR_SB_EGRESS_TC_COUNT (u16) shared buffer egress traffic class count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbEgressTcCount DevlinkAttr = 16

	// DevlinkAttrSbPoolIndex - DEVLINK_ATTR_SB_POOL_INDEX (u16) shared buffer pool index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbPoolIndex DevlinkAttr = 17

	// DevlinkAttrSbPoolType - DEVLINK_ATTR_SB_POOL_TYPE (u8) shared buffer pool type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbPoolType DevlinkAttr = 18

	// DevlinkAttrSbPoolSize - DEVLINK_ATTR_SB_POOL_SIZE (u32) shared buffer pool size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbPoolSize DevlinkAttr = 19

	// DevlinkAttrSbPoolThresholdType - DEVLINK_ATTR_SB_POOL_THRESHOLD_TYPE (u8) shared buffer pool threshold type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbPoolThresholdType DevlinkAttr = 20

	// DevlinkAttrSbThreshold - DEVLINK_ATTR_SB_THRESHOLD (u32) shared buffer threshold
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbThreshold DevlinkAttr = 21

	// DevlinkAttrSbTcIndex - DEVLINK_ATTR_SB_TC_INDEX (u16) shared buffer traffic class index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbTcIndex DevlinkAttr = 22

	// DevlinkAttrSbOccCur - DEVLINK_ATTR_SB_OCC_CUR (u32) shared buffer occupancy current
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbOccCur DevlinkAttr = 23

	// DevlinkAttrSbOccMax - DEVLINK_ATTR_SB_OCC_MAX (u32) shared buffer occupancy maximum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbOccMax DevlinkAttr = 24

	// DevlinkAttrEswitchMode - DEVLINK_ATTR_ESWITCH_MODE (u16) eswitch mode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEswitchMode DevlinkAttr = 25

	// DevlinkAttrEswitchInlineMode - DEVLINK_ATTR_ESWITCH_INLINE_MODE (u8) eswitch inline mode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEswitchInlineMode DevlinkAttr = 26

	// DevlinkAttrDpipeTables - DEVLINK_ATTR_DPIPE_TABLES (nested) dpipe tables
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTables DevlinkAttr = 27

	// DevlinkAttrDpipeTable - DEVLINK_ATTR_DPIPE_TABLE (nested) dpipe table
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTable DevlinkAttr = 28

	// DevlinkAttrDpipeTableName - DEVLINK_ATTR_DPIPE_TABLE_NAME (string) dpipe table name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableName DevlinkAttr = 29

	// DevlinkAttrDpipeTableSize - DEVLINK_ATTR_DPIPE_TABLE_SIZE (u64) dpipe table size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableSize DevlinkAttr = 30

	// DevlinkAttrDpipeTableMatches - DEVLINK_ATTR_DPIPE_TABLE_MATCHES (nested) dpipe table matches
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableMatches DevlinkAttr = 31

	// DevlinkAttrDpipeTableActions - DEVLINK_ATTR_DPIPE_TABLE_ACTIONS (nested) dpipe table actions
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableActions DevlinkAttr = 32

	// DevlinkAttrDpipeTableCountersEnabled - DEVLINK_ATTR_DPIPE_TABLE_COUNTERS_ENABLED (u8) dpipe table counters enabled
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableCountersEnabled DevlinkAttr = 33

	// DevlinkAttrDpipeEntries - DEVLINK_ATTR_DPIPE_ENTRIES (nested) dpipe entries
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeEntries DevlinkAttr = 34

	// DevlinkAttrDpipeEntry - DEVLINK_ATTR_DPIPE_ENTRY (nested) dpipe entry
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeEntry DevlinkAttr = 35

	// DevlinkAttrDpipeEntryIndex - DEVLINK_ATTR_DPIPE_ENTRY_INDEX (u64) dpipe entry index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeEntryIndex DevlinkAttr = 36

	// DevlinkAttrDpipeEntryMatchValues - DEVLINK_ATTR_DPIPE_ENTRY_MATCH_VALUES (nested) dpipe entry match values
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeEntryMatchValues DevlinkAttr = 37

	// DevlinkAttrDpipeEntryActionValues - DEVLINK_ATTR_DPIPE_ENTRY_ACTION_VALUES (nested) dpipe entry action values
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeEntryActionValues DevlinkAttr = 38

	// DevlinkAttrDpipeEntryCounter - DEVLINK_ATTR_DPIPE_ENTRY_COUNTER (u64) dpipe entry counter
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeEntryCounter DevlinkAttr = 39

	// DevlinkAttrDpipeMatch - DEVLINK_ATTR_DPIPE_MATCH (nested) dpipe match
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeMatch DevlinkAttr = 40

	// DevlinkAttrDpipeMatchValue - DEVLINK_ATTR_DPIPE_MATCH_VALUE (nested) dpipe match value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeMatchValue DevlinkAttr = 41

	// DevlinkAttrDpipeMatchType - DEVLINK_ATTR_DPIPE_MATCH_TYPE (u32) dpipe match type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeMatchType DevlinkAttr = 42

	// DevlinkAttrDpipeAction - DEVLINK_ATTR_DPIPE_ACTION (nested) dpipe action
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeAction DevlinkAttr = 43

	// DevlinkAttrDpipeActionValue - DEVLINK_ATTR_DPIPE_ACTION_VALUE (nested) dpipe action value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeActionValue DevlinkAttr = 44

	// DevlinkAttrDpipeActionType - DEVLINK_ATTR_DPIPE_ACTION_TYPE (u32) dpipe action type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeActionType DevlinkAttr = 45

	// DevlinkAttrDpipeValue - DEVLINK_ATTR_DPIPE_VALUE (u32) dpipe value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeValue DevlinkAttr = 46

	// DevlinkAttrDpipeValueMask - DEVLINK_ATTR_DPIPE_VALUE_MASK (u32) dpipe value mask
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeValueMask DevlinkAttr = 47

	// DevlinkAttrDpipeValueMapping - DEVLINK_ATTR_DPIPE_VALUE_MAPPING (u32) dpipe value mapping
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeValueMapping DevlinkAttr = 48

	// DevlinkAttrDpipeHeaders - DEVLINK_ATTR_DPIPE_HEADERS (nested) dpipe headers
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeaders DevlinkAttr = 49

	// DevlinkAttrDpipeHeader - DEVLINK_ATTR_DPIPE_HEADER (nested) dpipe header
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeader DevlinkAttr = 50

	// DevlinkAttrDpipeHeaderName - DEVLINK_ATTR_DPIPE_HEADER_NAME (string) dpipe header name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeaderName DevlinkAttr = 51

	// DevlinkAttrDpipeHeaderId - DEVLINK_ATTR_DPIPE_HEADER_ID (u32) dpipe header ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeaderId DevlinkAttr = 52

	// DevlinkAttrDpipeHeaderFields - DEVLINK_ATTR_DPIPE_HEADER_FIELDS (nested) dpipe header fields
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeaderFields DevlinkAttr = 53

	// DevlinkAttrDpipeHeaderGlobal - DEVLINK_ATTR_DPIPE_HEADER_GLOBAL (u8) dpipe header global
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeaderGlobal DevlinkAttr = 54

	// DevlinkAttrDpipeHeaderIndex - DEVLINK_ATTR_DPIPE_HEADER_INDEX (u32) dpipe header index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeHeaderIndex DevlinkAttr = 55

	// DevlinkAttrDpipeField - DEVLINK_ATTR_DPIPE_FIELD (nested) dpipe field
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeField DevlinkAttr = 56

	// DevlinkAttrDpipeFieldName - DEVLINK_ATTR_DPIPE_FIELD_NAME (string) dpipe field name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeFieldName DevlinkAttr = 57

	// DevlinkAttrDpipeFieldId - DEVLINK_ATTR_DPIPE_FIELD_ID (u32) dpipe field ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeFieldId DevlinkAttr = 58

	// DevlinkAttrDpipeFieldBitwidth - DEVLINK_ATTR_DPIPE_FIELD_BITWIDTH (u32) dpipe field bitwidth
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeFieldBitwidth DevlinkAttr = 59

	// DevlinkAttrDpipeFieldMappingType - DEVLINK_ATTR_DPIPE_FIELD_MAPPING_TYPE (u32) dpipe field mapping type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeFieldMappingType DevlinkAttr = 60

	// DevlinkAttrPad - DEVLINK_ATTR_PAD (u32) padding attribute
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPad DevlinkAttr = 61

	// DevlinkAttrEswitchEncapMode - DEVLINK_ATTR_ESWITCH_ENCAP_MODE (u8) eswitch encapsulation mode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEswitchEncapMode DevlinkAttr = 62

	// DevlinkAttrResourceList - DEVLINK_ATTR_RESOURCE_LIST (nested) resource list
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceList DevlinkAttr = 63

	// DevlinkAttrResource - DEVLINK_ATTR_RESOURCE (nested) resource
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResource DevlinkAttr = 64

	// DevlinkAttrResourceName - DEVLINK_ATTR_RESOURCE_NAME (string) resource name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceName DevlinkAttr = 65

	// DevlinkAttrResourceId - DEVLINK_ATTR_RESOURCE_ID (u64) resource ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceId DevlinkAttr = 66

	// DevlinkAttrResourceSize - DEVLINK_ATTR_RESOURCE_SIZE (u64) resource size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceSize DevlinkAttr = 67

	// DevlinkAttrResourceSizeNew - DEVLINK_ATTR_RESOURCE_SIZE_NEW (u64) new resource size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceSizeNew DevlinkAttr = 68

	// DevlinkAttrResourceSizeValid - DEVLINK_ATTR_RESOURCE_SIZE_VALID (u8) resource size valid
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceSizeValid DevlinkAttr = 69

	// DevlinkAttrResourceSizeMin - DEVLINK_ATTR_RESOURCE_SIZE_MIN (u64) resource size minimum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceSizeMin DevlinkAttr = 70

	// DevlinkAttrResourceSizeMax - DEVLINK_ATTR_RESOURCE_SIZE_MAX (u64) resource size maximum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceSizeMax DevlinkAttr = 71

	// DevlinkAttrResourceSizeGran - DEVLINK_ATTR_RESOURCE_SIZE_GRAN (u64) resource size granularity
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceSizeGran DevlinkAttr = 72

	// DevlinkAttrResourceUnit - DEVLINK_ATTR_RESOURCE_UNIT (u8) resource unit
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceUnit DevlinkAttr = 73

	// DevlinkAttrResourceOcc - DEVLINK_ATTR_RESOURCE_OCC (u64) resource occupancy
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrResourceOcc DevlinkAttr = 74

	// DevlinkAttrDpipeTableResourceId - DEVLINK_ATTR_DPIPE_TABLE_RESOURCE_ID (u64) dpipe table resource ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableResourceId DevlinkAttr = 75

	// DevlinkAttrDpipeTableResourceUnits - DEVLINK_ATTR_DPIPE_TABLE_RESOURCE_UNITS (u64) dpipe table resource units
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDpipeTableResourceUnits DevlinkAttr = 76

	// DevlinkAttrPortFlavour - DEVLINK_ATTR_PORT_FLAVOUR (u64) port flavour
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortFlavour DevlinkAttr = 77

	// DevlinkAttrPortNumber - DEVLINK_ATTR_PORT_NUMBER (u32) port number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortNumber DevlinkAttr = 78

	// DevlinkAttrPortSplitSubportNumber - DEVLINK_ATTR_PORT_SPLIT_SUBPORT_NUMBER (u32) port split subport number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortSplitSubportNumber DevlinkAttr = 79

	// DevlinkAttrParam - DEVLINK_ATTR_PARAM (nested) parameter
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParam DevlinkAttr = 80

	// DevlinkAttrParamName - DEVLINK_ATTR_PARAM_NAME (string) parameter name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamName DevlinkAttr = 81

	// DevlinkAttrParamGeneric - DEVLINK_ATTR_PARAM_GENERIC (flag) parameter generic
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamGeneric DevlinkAttr = 82

	// DevlinkAttrParamType - DEVLINK_ATTR_PARAM_TYPE (u8) parameter type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamType DevlinkAttr = 83

	// DevlinkAttrParamValuesList - DEVLINK_ATTR_PARAM_VALUES_LIST (nested) parameter values list
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamValuesList DevlinkAttr = 84

	// DevlinkAttrParamValue - DEVLINK_ATTR_PARAM_VALUE (nested) parameter value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamValue DevlinkAttr = 85

	// DevlinkAttrParamValueData - DEVLINK_ATTR_PARAM_VALUE_DATA (dynamic) parameter value data
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamValueData DevlinkAttr = 86

	// DevlinkAttrParamValueCmode - DEVLINK_ATTR_PARAM_VALUE_CMODE (u8) parameter value cmode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrParamValueCmode DevlinkAttr = 87

	// DevlinkAttrRegionName - DEVLINK_ATTR_REGION_NAME (string) region name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionName DevlinkAttr = 88

	// DevlinkAttrRegionSize - DEVLINK_ATTR_REGION_SIZE (u64) region size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionSize DevlinkAttr = 89

	// DevlinkAttrRegionSnapshots - DEVLINK_ATTR_REGION_SNAPSHOTS (nested) region snapshots
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionSnapshots DevlinkAttr = 90

	// DevlinkAttrRegionSnapshot - DEVLINK_ATTR_REGION_SNAPSHOT (nested) region snapshot
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionSnapshot DevlinkAttr = 91

	// DevlinkAttrRegionSnapshotId - DEVLINK_ATTR_REGION_SNAPSHOT_ID (u32) region snapshot ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionSnapshotId DevlinkAttr = 92

	// DevlinkAttrRegionChunks - DEVLINK_ATTR_REGION_CHUNKS (nested) region chunks
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionChunks DevlinkAttr = 93

	// DevlinkAttrRegionChunk - DEVLINK_ATTR_REGION_CHUNK (nested) region chunk
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionChunk DevlinkAttr = 94

	// DevlinkAttrRegionChunkData - DEVLINK_ATTR_REGION_CHUNK_DATA (binary) region chunk data
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionChunkData DevlinkAttr = 95

	// DevlinkAttrRegionChunkAddr - DEVLINK_ATTR_REGION_CHUNK_ADDR (u64) region chunk address
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionChunkAddr DevlinkAttr = 96

	// DevlinkAttrRegionChunkLen - DEVLINK_ATTR_REGION_CHUNK_LEN (u64) region chunk length
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionChunkLen DevlinkAttr = 97

	// DevlinkAttrInfoDriverName - DEVLINK_ATTR_INFO_DRIVER_NAME (string) info driver name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoDriverName DevlinkAttr = 98

	// DevlinkAttrInfoSerialNumber - DEVLINK_ATTR_INFO_SERIAL_NUMBER (string) info serial number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoSerialNumber DevlinkAttr = 99

	// DevlinkAttrInfoVersionFixed - DEVLINK_ATTR_INFO_VERSION_FIXED (nested) info version fixed
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoVersionFixed DevlinkAttr = 100

	// DevlinkAttrInfoVersionRunning - DEVLINK_ATTR_INFO_VERSION_RUNNING (nested) info version running
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoVersionRunning DevlinkAttr = 101

	// DevlinkAttrInfoVersionStored - DEVLINK_ATTR_INFO_VERSION_STORED (nested) info version stored
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoVersionStored DevlinkAttr = 102

	// DevlinkAttrInfoVersionName - DEVLINK_ATTR_INFO_VERSION_NAME (string) info version name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoVersionName DevlinkAttr = 103

	// DevlinkAttrInfoVersionValue - DEVLINK_ATTR_INFO_VERSION_VALUE (string) info version value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoVersionValue DevlinkAttr = 104

	// DevlinkAttrSbPoolCellSize - DEVLINK_ATTR_SB_POOL_CELL_SIZE (u32) shared buffer pool cell size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSbPoolCellSize DevlinkAttr = 105

	// DevlinkAttrFmsg - DEVLINK_ATTR_FMSG (nested) fmsg
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsg DevlinkAttr = 106

	// DevlinkAttrFmsgObjNestStart - DEVLINK_ATTR_FMSG_OBJ_NEST_START (flag) fmsg object nest start
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgObjNestStart DevlinkAttr = 107

	// DevlinkAttrFmsgPairNestStart - DEVLINK_ATTR_FMSG_PAIR_NEST_START (flag) fmsg pair nest start
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgPairNestStart DevlinkAttr = 108

	// DevlinkAttrFmsgArrNestStart - DEVLINK_ATTR_FMSG_ARR_NEST_START (flag) fmsg array nest start
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgArrNestStart DevlinkAttr = 109

	// DevlinkAttrFmsgNestEnd - DEVLINK_ATTR_FMSG_NEST_END (flag) fmsg nest end
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgNestEnd DevlinkAttr = 110

	// DevlinkAttrFmsgObjName - DEVLINK_ATTR_FMSG_OBJ_NAME (string) fmsg object name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgObjName DevlinkAttr = 111

	// DevlinkAttrFmsgObjValueType - DEVLINK_ATTR_FMSG_OBJ_VALUE_TYPE (u8) fmsg object value type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgObjValueType DevlinkAttr = 112

	// DevlinkAttrFmsgObjValueData - DEVLINK_ATTR_FMSG_OBJ_VALUE_DATA (dynamic) fmsg object value data
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFmsgObjValueData DevlinkAttr = 113

	// DevlinkAttrHealthReporter - DEVLINK_ATTR_HEALTH_REPORTER (nested) health reporter
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporter DevlinkAttr = 114

	// DevlinkAttrHealthReporterName - DEVLINK_ATTR_HEALTH_REPORTER_NAME (string) health reporter name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterName DevlinkAttr = 115

	// DevlinkAttrHealthReporterState - DEVLINK_ATTR_HEALTH_REPORTER_STATE (u8) health reporter state
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterState DevlinkAttr = 116

	// DevlinkAttrHealthReporterErrCount - DEVLINK_ATTR_HEALTH_REPORTER_ERR_COUNT (u64) health reporter error count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterErrCount DevlinkAttr = 117

	// DevlinkAttrHealthReporterRecoverCount - DEVLINK_ATTR_HEALTH_REPORTER_RECOVER_COUNT (u64) health reporter recover count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterRecoverCount DevlinkAttr = 118

	// DevlinkAttrHealthReporterDumpTs - DEVLINK_ATTR_HEALTH_REPORTER_DUMP_TS (u64) health reporter dump timestamp
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterDumpTs DevlinkAttr = 119

	// DevlinkAttrHealthReporterGracefulPeriod - DEVLINK_ATTR_HEALTH_REPORTER_GRACEFUL_PERIOD (u64) health reporter graceful period
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterGracefulPeriod DevlinkAttr = 120

	// DevlinkAttrHealthReporterAutoRecover - DEVLINK_ATTR_HEALTH_REPORTER_AUTO_RECOVER (u8) health reporter auto recover
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterAutoRecover DevlinkAttr = 121

	// DevlinkAttrFlashUpdateFileName - DEVLINK_ATTR_FLASH_UPDATE_FILE_NAME (string) flash update file name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateFileName DevlinkAttr = 122

	// DevlinkAttrFlashUpdateComponent - DEVLINK_ATTR_FLASH_UPDATE_COMPONENT (string) flash update component
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateComponent DevlinkAttr = 123

	// DevlinkAttrFlashUpdateStatusMsg - DEVLINK_ATTR_FLASH_UPDATE_STATUS_MSG (u64) flash update status message
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateStatusMsg DevlinkAttr = 124

	// DevlinkAttrFlashUpdateStatusDone - DEVLINK_ATTR_FLASH_UPDATE_STATUS_DONE (u64) flash update status done
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateStatusDone DevlinkAttr = 125

	// DevlinkAttrFlashUpdateStatusTotal - DEVLINK_ATTR_FLASH_UPDATE_STATUS_TOTAL (u64) flash update status total
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateStatusTotal DevlinkAttr = 126

	// DevlinkAttrPortPciPfNumber - DEVLINK_ATTR_PORT_PCI_PF_NUMBER (u16) port PCI PF number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortPciPfNumber DevlinkAttr = 127

	// DevlinkAttrPortPciVfNumber - DEVLINK_ATTR_PORT_PCI_VF_NUMBER (u16) port PCI VF number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortPciVfNumber DevlinkAttr = 128

	// DevlinkAttrStats - DEVLINK_ATTR_STATS (nested) statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrStats DevlinkAttr = 129

	// DevlinkAttrTrapName - DEVLINK_ATTR_TRAP_NAME (string) trap name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapName DevlinkAttr = 130

	// DevlinkAttrTrapAction - DEVLINK_ATTR_TRAP_ACTION (u8) trap action (enum devlink_trap_action)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapAction DevlinkAttr = 131

	// DevlinkAttrTrapType - DEVLINK_ATTR_TRAP_TYPE (u8) trap type (enum devlink_trap_type)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapType DevlinkAttr = 132

	// DevlinkAttrTrapGeneric - DEVLINK_ATTR_TRAP_GENERIC (flag) generic trap
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapGeneric DevlinkAttr = 133

	// DevlinkAttrTrapMetadata - DEVLINK_ATTR_TRAP_METADATA (nested) trap metadata
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapMetadata DevlinkAttr = 134

	// DevlinkAttrTrapGroupName - DEVLINK_ATTR_TRAP_GROUP_NAME (string) trap group name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapGroupName DevlinkAttr = 135

	// DevlinkAttrReloadFailed - DEVLINK_ATTR_RELOAD_FAILED (u8) reload failed (0 or 1)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadFailed DevlinkAttr = 136

	// DevlinkAttrHealthReporterDumpTsNs - DEVLINK_ATTR_HEALTH_REPORTER_DUMP_TS_NS (u64) health reporter dump timestamp in nanoseconds
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterDumpTsNs DevlinkAttr = 137

	// DevlinkAttrNetnsFd - DEVLINK_ATTR_NETNS_FD (u32) network namespace file descriptor
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrNetnsFd DevlinkAttr = 138

	// DevlinkAttrNetnsPid - DEVLINK_ATTR_NETNS_PID (u32) network namespace process ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrNetnsPid DevlinkAttr = 139

	// DevlinkAttrNetnsId - DEVLINK_ATTR_NETNS_ID (u32) network namespace ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrNetnsId DevlinkAttr = 140

	// DevlinkAttrHealthReporterAutoDump - DEVLINK_ATTR_HEALTH_REPORTER_AUTO_DUMP (u8) health reporter auto dump
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrHealthReporterAutoDump DevlinkAttr = 141

	// DevlinkAttrTrapPolicerId - DEVLINK_ATTR_TRAP_POLICER_ID (u32) trap policer ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapPolicerId DevlinkAttr = 142

	// DevlinkAttrTrapPolicerRate - DEVLINK_ATTR_TRAP_POLICER_RATE (u64) trap policer rate
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapPolicerRate DevlinkAttr = 143

	// DevlinkAttrTrapPolicerBurst - DEVLINK_ATTR_TRAP_POLICER_BURST (u64) trap policer burst
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrTrapPolicerBurst DevlinkAttr = 144

	// DevlinkAttrPortFunction - DEVLINK_ATTR_PORT_FUNCTION (nested) port function
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortFunction DevlinkAttr = 145

	// DevlinkAttrInfoBoardSerialNumber - DEVLINK_ATTR_INFO_BOARD_SERIAL_NUMBER (string) info board serial number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrInfoBoardSerialNumber DevlinkAttr = 146

	// DevlinkAttrPortLanes - DEVLINK_ATTR_PORT_LANES (u32) port lanes
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortLanes DevlinkAttr = 147

	// DevlinkAttrPortSplittable - DEVLINK_ATTR_PORT_SPLITTABLE (u8) port splittable
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortSplittable DevlinkAttr = 148

	// DevlinkAttrPortExternal - DEVLINK_ATTR_PORT_EXTERNAL (u8) port external
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortExternal DevlinkAttr = 149

	// DevlinkAttrPortControllerNumber - DEVLINK_ATTR_PORT_CONTROLLER_NUMBER (u32) port controller number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortControllerNumber DevlinkAttr = 150

	// DevlinkAttrFlashUpdateStatusTimeout - DEVLINK_ATTR_FLASH_UPDATE_STATUS_TIMEOUT (u64) flash update status timeout
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateStatusTimeout DevlinkAttr = 151

	// DevlinkAttrFlashUpdateOverwriteMask - DEVLINK_ATTR_FLASH_UPDATE_OVERWRITE_MASK (bitfield32) flash update overwrite mask
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrFlashUpdateOverwriteMask DevlinkAttr = 152

	// DevlinkAttrReloadAction - DEVLINK_ATTR_RELOAD_ACTION (u8) reload action
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadAction DevlinkAttr = 153

	// DevlinkAttrReloadActionsPerformed - DEVLINK_ATTR_RELOAD_ACTIONS_PERFORMED (bitfield32) reload actions performed
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadActionsPerformed DevlinkAttr = 154

	// DevlinkAttrReloadLimits - DEVLINK_ATTR_RELOAD_LIMITS (bitfield32) reload limits
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadLimits DevlinkAttr = 155

	// DevlinkAttrDevStats - DEVLINK_ATTR_DEV_STATS (nested) device statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrDevStats DevlinkAttr = 156

	// DevlinkAttrReloadStats - DEVLINK_ATTR_RELOAD_STATS (nested) reload statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadStats DevlinkAttr = 157

	// DevlinkAttrReloadStatsEntry - DEVLINK_ATTR_RELOAD_STATS_ENTRY (nested) reload statistics entry
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadStatsEntry DevlinkAttr = 158

	// DevlinkAttrReloadStatsLimit - DEVLINK_ATTR_RELOAD_STATS_LIMIT (u8) reload statistics limit
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadStatsLimit DevlinkAttr = 159

	// DevlinkAttrReloadStatsValue - DEVLINK_ATTR_RELOAD_STATS_VALUE (u32) reload statistics value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadStatsValue DevlinkAttr = 160

	// DevlinkAttrRemoteReloadStats - DEVLINK_ATTR_REMOTE_RELOAD_STATS (nested) remote reload statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRemoteReloadStats DevlinkAttr = 161

	// DevlinkAttrReloadActionInfo - DEVLINK_ATTR_RELOAD_ACTION_INFO (nested) reload action information
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadActionInfo DevlinkAttr = 162

	// DevlinkAttrReloadActionStats - DEVLINK_ATTR_RELOAD_ACTION_STATS (nested) reload action statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrReloadActionStats DevlinkAttr = 163

	// DevlinkAttrPortPciSfNumber - DEVLINK_ATTR_PORT_PCI_SF_NUMBER (u32) port PCI SF number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrPortPciSfNumber DevlinkAttr = 164

	// DevlinkAttrRateType - DEVLINK_ATTR_RATE_TYPE (u16) rate type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateType DevlinkAttr = 165

	// DevlinkAttrRateTxShare - DEVLINK_ATTR_RATE_TX_SHARE (u64) rate TX share
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateTxShare DevlinkAttr = 166

	// DevlinkAttrRateTxMax - DEVLINK_ATTR_RATE_TX_MAX (u64) rate TX maximum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateTxMax DevlinkAttr = 167

	// DevlinkAttrRateNodeName - DEVLINK_ATTR_RATE_NODE_NAME (string) rate node name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateNodeName DevlinkAttr = 168

	// DevlinkAttrRateParentNodeName - DEVLINK_ATTR_RATE_PARENT_NODE_NAME (string) rate parent node name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateParentNodeName DevlinkAttr = 169

	// DevlinkAttrRegionMaxSnapshots - DEVLINK_ATTR_REGION_MAX_SNAPSHOTS (u32) region maximum snapshots
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionMaxSnapshots DevlinkAttr = 170

	// DevlinkAttrLinecardIndex - DEVLINK_ATTR_LINECARD_INDEX (u32) linecard index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrLinecardIndex DevlinkAttr = 171

	// DevlinkAttrLinecardState - DEVLINK_ATTR_LINECARD_STATE (u8) linecard state
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrLinecardState DevlinkAttr = 172

	// DevlinkAttrLinecardType - DEVLINK_ATTR_LINECARD_TYPE (string) linecard type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrLinecardType DevlinkAttr = 173

	// DevlinkAttrLinecardSupportedTypes - DEVLINK_ATTR_LINECARD_SUPPORTED_TYPES (nested) linecard supported types
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrLinecardSupportedTypes DevlinkAttr = 174

	// DevlinkAttrNestedDevlink - DEVLINK_ATTR_NESTED_DEVLINK (nested) nested devlink
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrNestedDevlink DevlinkAttr = 175

	// DevlinkAttrSelftests - DEVLINK_ATTR_SELFTESTS (nested) selftests
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrSelftests DevlinkAttr = 176

	// DevlinkAttrRateTxPriority - DEVLINK_ATTR_RATE_TX_PRIORITY (u32) rate TX priority
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateTxPriority DevlinkAttr = 177

	// DevlinkAttrRateTxWeight - DEVLINK_ATTR_RATE_TX_WEIGHT (u32) rate TX weight
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRateTxWeight DevlinkAttr = 178

	// DevlinkAttrRegionDirect - DEVLINK_ATTR_REGION_DIRECT (flag) region direct
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrRegionDirect DevlinkAttr = 179

	// DevlinkAttrMax - DEVLINK_ATTR_MAX - maximum value for devlink attributes
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrMax = iota - 1
)
