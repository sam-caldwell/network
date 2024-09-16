package core

// DevlinkAttr - devlink_attr - represents the attributes of devlink as defined in linux/devlink.h
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkAttr uint8

/*
don't change the order or add anything between, this is ABI!
*/
const (
	// DevlinkAttrEnumUnspec - DEVLINK_ATTR_UNSPEC - unspecified attribute
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumUnspec DevlinkAttr = 0

	// DevlinkAttrEnumBusName - DEVLINK_ATTR_BUS_NAME (string) bus name + dev name together are a handle for
	// devlink entity
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumBusName DevlinkAttr = 1

	// DevlinkAttrEnumDevName - DEVLINK_ATTR_DEV_NAME (string) device name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDevName DevlinkAttr = 2

	// DevlinkAttrEnumPortIndex - DEVLINK_ATTR_PORT_INDEX (u32) port index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortIndex DevlinkAttr = 3

	// DevlinkAttrEnumPortType - DEVLINK_ATTR_PORT_TYPE (u16) port type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortType DevlinkAttr = 4

	// DevlinkAttrEnumPortDesiredType - DEVLINK_ATTR_PORT_DESIRED_TYPE (u16) port desired type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortDesiredType DevlinkAttr = 5

	// DevlinkAttrEnumPortNetdevIfindex - DEVLINK_ATTR_PORT_NETDEV_IFINDEX (u32) port netdev ifindex
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortNetdevIfindex DevlinkAttr = 6

	// DevlinkAttrEnumPortNetdevName - DEVLINK_ATTR_PORT_NETDEV_NAME (string) port netdev name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortNetdevName DevlinkAttr = 7

	// DevlinkAttrEnumPortIbdevName - DEVLINK_ATTR_PORT_IBDEV_NAME (string) port ibdev name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortIbdevName DevlinkAttr = 8

	// DevlinkAttrEnumPortSplitCount - DEVLINK_ATTR_PORT_SPLIT_COUNT (u32) port split count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortSplitCount DevlinkAttr = 9

	// DevlinkAttrEnumPortSplitGroup - DEVLINK_ATTR_PORT_SPLIT_GROUP (u32) port split group
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortSplitGroup DevlinkAttr = 10

	// DevlinkAttrEnumSbIndex - DEVLINK_ATTR_SB_INDEX (u32) shared buffer index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbIndex DevlinkAttr = 11

	// DevlinkAttrEnumSbSize - DEVLINK_ATTR_SB_SIZE (u32) shared buffer size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbSize DevlinkAttr = 12

	// DevlinkAttrEnumSbIngressPoolCount - DEVLINK_ATTR_SB_INGRESS_POOL_COUNT (u16) shared buffer ingress pool count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbIngressPoolCount DevlinkAttr = 13

	// DevlinkAttrEnumSbEgressPoolCount - DEVLINK_ATTR_SB_EGRESS_POOL_COUNT (u16) shared buffer egress pool count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbEgressPoolCount DevlinkAttr = 14

	// DevlinkAttrEnumSbIngressTcCount - DEVLINK_ATTR_SB_INGRESS_TC_COUNT (u16) shared buffer ingress traffic
	// class count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbIngressTcCount DevlinkAttr = 15

	// DevlinkAttrEnumSbEgressTcCount - DEVLINK_ATTR_SB_EGRESS_TC_COUNT (u16) shared buffer egress traffic
	// class count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbEgressTcCount DevlinkAttr = 16

	// DevlinkAttrEnumSbPoolIndex - DEVLINK_ATTR_SB_POOL_INDEX (u16) shared buffer pool index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbPoolIndex DevlinkAttr = 17

	// DevlinkAttrEnumSbPoolType - DEVLINK_ATTR_SB_POOL_TYPE (u8) shared buffer pool type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbPoolType DevlinkAttr = 18

	// DevlinkAttrEnumSbPoolSize - DEVLINK_ATTR_SB_POOL_SIZE (u32) shared buffer pool size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbPoolSize DevlinkAttr = 19

	// DevlinkAttrEnumSbPoolThresholdType - DEVLINK_ATTR_SB_POOL_THRESHOLD_TYPE (u8) shared buffer pool
	// threshold type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbPoolThresholdType DevlinkAttr = 20

	// DevlinkAttrEnumSbThreshold - DEVLINK_ATTR_SB_THRESHOLD (u32) shared buffer threshold
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbThreshold DevlinkAttr = 21

	// DevlinkAttrEnumSbTcIndex - DEVLINK_ATTR_SB_TC_INDEX (u16) shared buffer traffic class index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbTcIndex DevlinkAttr = 22

	// DevlinkAttrEnumSbOccCur - DEVLINK_ATTR_SB_OCC_CUR (u32) shared buffer occupancy current
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbOccCur DevlinkAttr = 23

	// DevlinkAttrEnumSbOccMax - DEVLINK_ATTR_SB_OCC_MAX (u32) shared buffer occupancy maximum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbOccMax DevlinkAttr = 24

	// DevlinkAttrEnumEswitchMode - DEVLINK_ATTR_ESWITCH_MODE (u16) eswitch mode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumEswitchMode DevlinkAttr = 25

	// DevlinkAttrEnumEswitchInlineMode - DEVLINK_ATTR_ESWITCH_INLINE_MODE (u8) eswitch inline mode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumEswitchInlineMode DevlinkAttr = 26

	// DevlinkAttrEnumDpipeTables - DEVLINK_ATTR_DPIPE_TABLES (nested) dpipe tables
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTables DevlinkAttr = 27

	// DevlinkAttrEnumDpipeTable - DEVLINK_ATTR_DPIPE_TABLE (nested) dpipe table
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTable DevlinkAttr = 28

	// DevlinkAttrEnumDpipeTableName - DEVLINK_ATTR_DPIPE_TABLE_NAME (string) dpipe table name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableName DevlinkAttr = 29

	// DevlinkAttrEnumDpipeTableSize - DEVLINK_ATTR_DPIPE_TABLE_SIZE (u64) dpipe table size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableSize DevlinkAttr = 30

	// DevlinkAttrEnumDpipeTableMatches - DEVLINK_ATTR_DPIPE_TABLE_MATCHES (nested) dpipe table matches
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableMatches DevlinkAttr = 31

	// DevlinkAttrEnumDpipeTableActions - DEVLINK_ATTR_DPIPE_TABLE_ACTIONS (nested) dpipe table actions
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableActions DevlinkAttr = 32

	// DevlinkAttrEnumDpipeTableCountersEnabled - DEVLINK_ATTR_DPIPE_TABLE_COUNTERS_ENABLED (u8) dpipe table
	// counters enabled
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableCountersEnabled DevlinkAttr = 33

	// DevlinkAttrEnumDpipeEntries - DEVLINK_ATTR_DPIPE_ENTRIES (nested) dpipe entries
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeEntries DevlinkAttr = 34

	// DevlinkAttrEnumDpipeEntry - DEVLINK_ATTR_DPIPE_ENTRY (nested) dpipe entry
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeEntry DevlinkAttr = 35

	// DevlinkAttrEnumDpipeEntryIndex - DEVLINK_ATTR_DPIPE_ENTRY_INDEX (u64) dpipe entry index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeEntryIndex DevlinkAttr = 36

	// DevlinkAttrEnumDpipeEntryMatchValues - DEVLINK_ATTR_DPIPE_ENTRY_MATCH_VALUES (nested) dpipe entry match values
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeEntryMatchValues DevlinkAttr = 37

	// DevlinkAttrEnumDpipeEntryActionValues - DEVLINK_ATTR_DPIPE_ENTRY_ACTION_VALUES (nested) dpipe entry action
	// values
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeEntryActionValues DevlinkAttr = 38

	// DevlinkAttrEnumDpipeEntryCounter - DEVLINK_ATTR_DPIPE_ENTRY_COUNTER (u64) dpipe entry counter
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeEntryCounter DevlinkAttr = 39

	// DevlinkAttrEnumDpipeMatch - DEVLINK_ATTR_DPIPE_MATCH (nested) dpipe match
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeMatch DevlinkAttr = 40

	// DevlinkAttrEnumDpipeMatchValue - DEVLINK_ATTR_DPIPE_MATCH_VALUE (nested) dpipe match value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeMatchValue DevlinkAttr = 41

	// DevlinkAttrEnumDpipeMatchType - DEVLINK_ATTR_DPIPE_MATCH_TYPE (u32) dpipe match type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeMatchType DevlinkAttr = 42

	// DevlinkAttrEnumDpipeAction - DEVLINK_ATTR_DPIPE_ACTION (nested) dpipe action
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeAction DevlinkAttr = 43

	// DevlinkAttrEnumDpipeActionValue - DEVLINK_ATTR_DPIPE_ACTION_VALUE (nested) dpipe action value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeActionValue DevlinkAttr = 44

	// DevlinkAttrEnumDpipeActionType - DEVLINK_ATTR_DPIPE_ACTION_TYPE (u32) dpipe action type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeActionType DevlinkAttr = 45

	// DevlinkAttrEnumDpipeValue - DEVLINK_ATTR_DPIPE_VALUE (u32) dpipe value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeValue DevlinkAttr = 46

	// DevlinkAttrEnumDpipeValueMask - DEVLINK_ATTR_DPIPE_VALUE_MASK (u32) dpipe value mask
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeValueMask DevlinkAttr = 47

	// DevlinkAttrEnumDpipeValueMapping - DEVLINK_ATTR_DPIPE_VALUE_MAPPING (u32) dpipe value mapping
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeValueMapping DevlinkAttr = 48

	// DevlinkAttrEnumDpipeHeaders - DEVLINK_ATTR_DPIPE_HEADERS (nested) dpipe headers
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeaders DevlinkAttr = 49

	// DevlinkAttrEnumDpipeHeader - DEVLINK_ATTR_DPIPE_HEADER (nested) dpipe header
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeader DevlinkAttr = 50

	// DevlinkAttrEnumDpipeHeaderName - DEVLINK_ATTR_DPIPE_HEADER_NAME (string) dpipe header name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeaderName DevlinkAttr = 51

	// DevlinkAttrEnumDpipeHeaderId - DEVLINK_ATTR_DPIPE_HEADER_ID (u32) dpipe header ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeaderId DevlinkAttr = 52

	// DevlinkAttrEnumDpipeHeaderFields - DEVLINK_ATTR_DPIPE_HEADER_FIELDS (nested) dpipe header fields
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeaderFields DevlinkAttr = 53

	// DevlinkAttrEnumDpipeHeaderGlobal - DEVLINK_ATTR_DPIPE_HEADER_GLOBAL (u8) dpipe header global
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeaderGlobal DevlinkAttr = 54

	// DevlinkAttrEnumDpipeHeaderIndex - DEVLINK_ATTR_DPIPE_HEADER_INDEX (u32) dpipe header index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeHeaderIndex DevlinkAttr = 55

	// DevlinkAttrEnumDpipeField - DEVLINK_ATTR_DPIPE_FIELD (nested) dpipe field
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeField DevlinkAttr = 56

	// DevlinkAttrEnumDpipeFieldName - DEVLINK_ATTR_DPIPE_FIELD_NAME (string) dpipe field name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeFieldName DevlinkAttr = 57

	// DevlinkAttrEnumDpipeFieldId - DEVLINK_ATTR_DPIPE_FIELD_ID (u32) dpipe field ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeFieldId DevlinkAttr = 58

	// DevlinkAttrEnumDpipeFieldBitwidth - DEVLINK_ATTR_DPIPE_FIELD_BITWIDTH (u32) dpipe field bitwidth
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeFieldBitwidth DevlinkAttr = 59

	// DevlinkAttrEnumDpipeFieldMappingType - DEVLINK_ATTR_DPIPE_FIELD_MAPPING_TYPE (u32) dpipe field mapping type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeFieldMappingType DevlinkAttr = 60

	// DevlinkAttrEnumPad - DEVLINK_ATTR_PAD (u32) padding attribute
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPad DevlinkAttr = 61

	// DevlinkAttrEnumEswitchEncapMode - DEVLINK_ATTR_ESWITCH_ENCAP_MODE (u8) eswitch encapsulation mode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumEswitchEncapMode DevlinkAttr = 62

	// DevlinkAttrEnumResourceList - DEVLINK_ATTR_RESOURCE_LIST (nested) resource list
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceList DevlinkAttr = 63

	// DevlinkAttrEnumResource - DEVLINK_ATTR_RESOURCE (nested) resource
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResource DevlinkAttr = 64

	// DevlinkAttrEnumResourceName - DEVLINK_ATTR_RESOURCE_NAME (string) resource name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceName DevlinkAttr = 65

	// DevlinkAttrEnumResourceId - DEVLINK_ATTR_RESOURCE_ID (u64) resource ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceId DevlinkAttr = 66

	// DevlinkAttrEnumResourceSize - DEVLINK_ATTR_RESOURCE_SIZE (u64) resource size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceSize DevlinkAttr = 67

	// DevlinkAttrEnumResourceSizeNew - DEVLINK_ATTR_RESOURCE_SIZE_NEW (u64) new resource size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceSizeNew DevlinkAttr = 68

	// DevlinkAttrEnumResourceSizeValid - DEVLINK_ATTR_RESOURCE_SIZE_VALID (u8) resource size valid
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceSizeValid DevlinkAttr = 69

	// DevlinkAttrEnumResourceSizeMin - DEVLINK_ATTR_RESOURCE_SIZE_MIN (u64) resource size minimum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceSizeMin DevlinkAttr = 70

	// DevlinkAttrEnumResourceSizeMax - DEVLINK_ATTR_RESOURCE_SIZE_MAX (u64) resource size maximum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceSizeMax DevlinkAttr = 71

	// DevlinkAttrEnumResourceSizeGran - DEVLINK_ATTR_RESOURCE_SIZE_GRAN (u64) resource size granularity
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceSizeGran DevlinkAttr = 72

	// DevlinkAttrEnumResourceUnit - DEVLINK_ATTR_RESOURCE_UNIT (u8) resource unit
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceUnit DevlinkAttr = 73

	// DevlinkAttrEnumResourceOcc - DEVLINK_ATTR_RESOURCE_OCC (u64) resource occupancy
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumResourceOcc DevlinkAttr = 74

	// DevlinkAttrEnumDpipeTableResourceId - DEVLINK_ATTR_DPIPE_TABLE_RESOURCE_ID (u64) dpipe table resource ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableResourceId DevlinkAttr = 75

	// DevlinkAttrEnumDpipeTableResourceUnits - DEVLINK_ATTR_DPIPE_TABLE_RESOURCE_UNITS (u64) dpipe table
	// resource units
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDpipeTableResourceUnits DevlinkAttr = 76

	// DevlinkAttrEnumPortFlavour - DEVLINK_ATTR_PORT_FLAVOUR (u64) port flavour
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortFlavour DevlinkAttr = 77

	// DevlinkAttrEnumPortNumber - DEVLINK_ATTR_PORT_NUMBER (u32) port number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortNumber DevlinkAttr = 78

	// DevlinkAttrEnumPortSplitSubportNumber - DEVLINK_ATTR_PORT_SPLIT_SUBPORT_NUMBER (u32) port split subport number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortSplitSubportNumber DevlinkAttr = 79

	// DevlinkAttrEnumParam - DEVLINK_ATTR_PARAM (nested) parameter
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParam DevlinkAttr = 80

	// DevlinkAttrEnumParamName - DEVLINK_ATTR_PARAM_NAME (string) parameter name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamName DevlinkAttr = 81

	// DevlinkAttrEnumParamGeneric - DEVLINK_ATTR_PARAM_GENERIC (flag) parameter generic
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamGeneric DevlinkAttr = 82

	// DevlinkAttrEnumParamType - DEVLINK_ATTR_PARAM_TYPE (u8) parameter type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamType DevlinkAttr = 83

	// DevlinkAttrEnumParamValuesList - DEVLINK_ATTR_PARAM_VALUES_LIST (nested) parameter values list
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamValuesList DevlinkAttr = 84

	// DevlinkAttrEnumParamValue - DEVLINK_ATTR_PARAM_VALUE (nested) parameter value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamValue DevlinkAttr = 85

	// DevlinkAttrEnumParamValueData - DEVLINK_ATTR_PARAM_VALUE_DATA (dynamic) parameter value data
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamValueData DevlinkAttr = 86

	// DevlinkAttrEnumParamValueCmode - DEVLINK_ATTR_PARAM_VALUE_CMODE (u8) parameter value cmode
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumParamValueCmode DevlinkAttr = 87

	// DevlinkAttrEnumRegionName - DEVLINK_ATTR_REGION_NAME (string) region name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionName DevlinkAttr = 88

	// DevlinkAttrEnumRegionSize - DEVLINK_ATTR_REGION_SIZE (u64) region size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionSize DevlinkAttr = 89

	// DevlinkAttrEnumRegionSnapshots - DEVLINK_ATTR_REGION_SNAPSHOTS (nested) region snapshots
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionSnapshots DevlinkAttr = 90

	// DevlinkAttrEnumRegionSnapshot - DEVLINK_ATTR_REGION_SNAPSHOT (nested) region snapshot
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionSnapshot DevlinkAttr = 91

	// DevlinkAttrEnumRegionSnapshotId - DEVLINK_ATTR_REGION_SNAPSHOT_ID (u32) region snapshot ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionSnapshotId DevlinkAttr = 92

	// DevlinkAttrEnumRegionChunks - DEVLINK_ATTR_REGION_CHUNKS (nested) region chunks
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionChunks DevlinkAttr = 93

	// DevlinkAttrEnumRegionChunk - DEVLINK_ATTR_REGION_CHUNK (nested) region chunk
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionChunk DevlinkAttr = 94

	// DevlinkAttrEnumRegionChunkData - DEVLINK_ATTR_REGION_CHUNK_DATA (binary) region chunk data
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionChunkData DevlinkAttr = 95

	// DevlinkAttrEnumRegionChunkAddr - DEVLINK_ATTR_REGION_CHUNK_ADDR (u64) region chunk address
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionChunkAddr DevlinkAttr = 96

	// DevlinkAttrEnumRegionChunkLen - DEVLINK_ATTR_REGION_CHUNK_LEN (u64) region chunk length
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionChunkLen DevlinkAttr = 97

	// DevlinkAttrEnumInfoDriverName - DEVLINK_ATTR_INFO_DRIVER_NAME (string) info driver name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoDriverName DevlinkAttr = 98

	// DevlinkAttrEnumInfoSerialNumber - DEVLINK_ATTR_INFO_SERIAL_NUMBER (string) info serial number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoSerialNumber DevlinkAttr = 99

	// DevlinkAttrEnumInfoVersionFixed - DEVLINK_ATTR_INFO_VERSION_FIXED (nested) info version fixed
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoVersionFixed DevlinkAttr = 100

	// DevlinkAttrEnumInfoVersionRunning - DEVLINK_ATTR_INFO_VERSION_RUNNING (nested) info version running
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoVersionRunning DevlinkAttr = 101

	// DevlinkAttrEnumInfoVersionStored - DEVLINK_ATTR_INFO_VERSION_STORED (nested) info version stored
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoVersionStored DevlinkAttr = 102

	// DevlinkAttrEnumInfoVersionName - DEVLINK_ATTR_INFO_VERSION_NAME (string) info version name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoVersionName DevlinkAttr = 103

	// DevlinkAttrEnumInfoVersionValue - DEVLINK_ATTR_INFO_VERSION_VALUE (string) info version value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoVersionValue DevlinkAttr = 104

	// DevlinkAttrEnumSbPoolCellSize - DEVLINK_ATTR_SB_POOL_CELL_SIZE (u32) shared buffer pool cell size
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSbPoolCellSize DevlinkAttr = 105

	// DevlinkAttrEnumFmsg - DEVLINK_ATTR_FMSG (nested) fmsg
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsg DevlinkAttr = 106

	// DevlinkAttrEnumFmsgObjNestStart - DEVLINK_ATTR_FMSG_OBJ_NEST_START (flag) fmsg object nest start
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgObjNestStart DevlinkAttr = 107

	// DevlinkAttrEnumFmsgPairNestStart - DEVLINK_ATTR_FMSG_PAIR_NEST_START (flag) fmsg pair nest start
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgPairNestStart DevlinkAttr = 108

	// DevlinkAttrEnumFmsgArrNestStart - DEVLINK_ATTR_FMSG_ARR_NEST_START (flag) fmsg array nest start
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgArrNestStart DevlinkAttr = 109

	// DevlinkAttrEnumFmsgNestEnd - DEVLINK_ATTR_FMSG_NEST_END (flag) fmsg nest end
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgNestEnd DevlinkAttr = 110

	// DevlinkAttrEnumFmsgObjName - DEVLINK_ATTR_FMSG_OBJ_NAME (string) fmsg object name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgObjName DevlinkAttr = 111

	// DevlinkAttrEnumFmsgObjValueType - DEVLINK_ATTR_FMSG_OBJ_VALUE_TYPE (u8) fmsg object value type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgObjValueType DevlinkAttr = 112

	// DevlinkAttrEnumFmsgObjValueData - DEVLINK_ATTR_FMSG_OBJ_VALUE_DATA (dynamic) fmsg object value data
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFmsgObjValueData DevlinkAttr = 113

	// DevlinkAttrEnumHealthReporter - DEVLINK_ATTR_HEALTH_REPORTER (nested) health reporter
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporter DevlinkAttr = 114

	// DevlinkAttrEnumHealthReporterName - DEVLINK_ATTR_HEALTH_REPORTER_NAME (string) health reporter name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterName DevlinkAttr = 115

	// DevlinkAttrEnumHealthReporterState - DEVLINK_ATTR_HEALTH_REPORTER_STATE (u8) health reporter state
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterState DevlinkAttr = 116

	// DevlinkAttrEnumHealthReporterErrCount - DEVLINK_ATTR_HEALTH_REPORTER_ERR_COUNT (u64) health reporter error
	// count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterErrCount DevlinkAttr = 117

	// DevlinkAttrEnumHealthReporterRecoverCount - DEVLINK_ATTR_HEALTH_REPORTER_RECOVER_COUNT (u64) health reporter
	// recover count
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterRecoverCount DevlinkAttr = 118

	// DevlinkAttrEnumHealthReporterDumpTs - DEVLINK_ATTR_HEALTH_REPORTER_DUMP_TS (u64) health reporter dump
	// timestamp
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterDumpTs DevlinkAttr = 119

	// DevlinkAttrEnumHealthReporterGracefulPeriod - DEVLINK_ATTR_HEALTH_REPORTER_GRACEFUL_PERIOD (u64) health
	// reporter graceful period
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterGracefulPeriod DevlinkAttr = 120

	// DevlinkAttrEnumHealthReporterAutoRecover - DEVLINK_ATTR_HEALTH_REPORTER_AUTO_RECOVER (u8) health reporter
	// auto recover
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterAutoRecover DevlinkAttr = 121

	// DevlinkAttrEnumFlashUpdateFileName - DEVLINK_ATTR_FLASH_UPDATE_FILE_NAME (string) flash update file name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateFileName DevlinkAttr = 122

	// DevlinkAttrEnumFlashUpdateComponent - DEVLINK_ATTR_FLASH_UPDATE_COMPONENT (string) flash update component
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateComponent DevlinkAttr = 123

	// DevlinkAttrEnumFlashUpdateStatusMsg - DEVLINK_ATTR_FLASH_UPDATE_STATUS_MSG (u64) flash update status message
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateStatusMsg DevlinkAttr = 124

	// DevlinkAttrEnumFlashUpdateStatusDone - DEVLINK_ATTR_FLASH_UPDATE_STATUS_DONE (u64) flash update status done
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateStatusDone DevlinkAttr = 125

	// DevlinkAttrEnumFlashUpdateStatusTotal - DEVLINK_ATTR_FLASH_UPDATE_STATUS_TOTAL (u64) flash update status total
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateStatusTotal DevlinkAttr = 126

	// DevlinkAttrEnumPortPciPfNumber - DEVLINK_ATTR_PORT_PCI_PF_NUMBER (u16) port PCI PF number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortPciPfNumber DevlinkAttr = 127

	// DevlinkAttrEnumPortPciVfNumber - DEVLINK_ATTR_PORT_PCI_VF_NUMBER (u16) port PCI VF number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortPciVfNumber DevlinkAttr = 128

	// DevlinkAttrEnumStats - DEVLINK_ATTR_STATS (nested) statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumStats DevlinkAttr = 129

	// DevlinkAttrEnumTrapName - DEVLINK_ATTR_TRAP_NAME (string) trap name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapName DevlinkAttr = 130

	// DevlinkAttrEnumTrapAction - DEVLINK_ATTR_TRAP_ACTION (u8) trap action (enum devlink_trap_action)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapAction DevlinkAttr = 131

	// DevlinkAttrEnumTrapType - DEVLINK_ATTR_TRAP_TYPE (u8) trap type (enum devlink_trap_type)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapType DevlinkAttr = 132

	// DevlinkAttrEnumTrapGeneric - DEVLINK_ATTR_TRAP_GENERIC (flag) generic trap
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapGeneric DevlinkAttr = 133

	// DevlinkAttrEnumTrapMetadata - DEVLINK_ATTR_TRAP_METADATA (nested) trap metadata
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapMetadata DevlinkAttr = 134

	// DevlinkAttrEnumTrapGroupName - DEVLINK_ATTR_TRAP_GROUP_NAME (string) trap group name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapGroupName DevlinkAttr = 135

	// DevlinkAttrEnumReloadFailed - DEVLINK_ATTR_RELOAD_FAILED (u8) reload failed (0 or 1)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadFailed DevlinkAttr = 136

	// DevlinkAttrEnumHealthReporterDumpTsNs - DEVLINK_ATTR_HEALTH_REPORTER_DUMP_TS_NS (u64) health reporter dump
	// timestamp in nanoseconds
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterDumpTsNs DevlinkAttr = 137

	// DevlinkAttrEnumNetnsFd - DEVLINK_ATTR_NETNS_FD (u32) network namespace file descriptor
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumNetnsFd DevlinkAttr = 138

	// DevlinkAttrEnumNetnsPid - DEVLINK_ATTR_NETNS_PID (u32) network namespace process ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumNetnsPid DevlinkAttr = 139

	// DevlinkAttrEnumNetnsId - DEVLINK_ATTR_NETNS_ID (u32) network namespace ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumNetnsId DevlinkAttr = 140

	// DevlinkAttrEnumHealthReporterAutoDump - DEVLINK_ATTR_HEALTH_REPORTER_AUTO_DUMP (u8) health reporter auto dump
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumHealthReporterAutoDump DevlinkAttr = 141

	// DevlinkAttrEnumTrapPolicerId - DEVLINK_ATTR_TRAP_POLICER_ID (u32) trap policer ID
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapPolicerId DevlinkAttr = 142

	// DevlinkAttrEnumTrapPolicerRate - DEVLINK_ATTR_TRAP_POLICER_RATE (u64) trap policer rate
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapPolicerRate DevlinkAttr = 143

	// DevlinkAttrEnumTrapPolicerBurst - DEVLINK_ATTR_TRAP_POLICER_BURST (u64) trap policer burst
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumTrapPolicerBurst DevlinkAttr = 144

	// DevlinkAttrEnumPortFunction - DEVLINK_ATTR_PORT_FUNCTION (nested) port function
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortFunction DevlinkAttr = 145

	// DevlinkAttrEnumInfoBoardSerialNumber - DEVLINK_ATTR_INFO_BOARD_SERIAL_NUMBER (string) info board serial number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumInfoBoardSerialNumber DevlinkAttr = 146

	// DevlinkAttrEnumPortLanes - DEVLINK_ATTR_PORT_LANES (u32) port lanes
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortLanes DevlinkAttr = 147

	// DevlinkAttrEnumPortSplittable - DEVLINK_ATTR_PORT_SPLITTABLE (u8) port splittable
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortSplittable DevlinkAttr = 148

	// DevlinkAttrEnumPortExternal - DEVLINK_ATTR_PORT_EXTERNAL (u8) port external
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortExternal DevlinkAttr = 149

	// DevlinkAttrEnumPortControllerNumber - DEVLINK_ATTR_PORT_CONTROLLER_NUMBER (u32) port controller number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortControllerNumber DevlinkAttr = 150

	// DevlinkAttrEnumFlashUpdateStatusTimeout - DEVLINK_ATTR_FLASH_UPDATE_STATUS_TIMEOUT (u64) flash update status
	// timeout
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateStatusTimeout DevlinkAttr = 151

	// DevlinkAttrEnumFlashUpdateOverwriteMask - DEVLINK_ATTR_FLASH_UPDATE_OVERWRITE_MASK (bitfield32) flash update
	// overwrite mask
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumFlashUpdateOverwriteMask DevlinkAttr = 152

	// DevlinkAttrEnumReloadAction - DEVLINK_ATTR_RELOAD_ACTION (u8) reload action
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadAction DevlinkAttr = 153

	// DevlinkAttrEnumReloadActionsPerformed - DEVLINK_ATTR_RELOAD_ACTIONS_PERFORMED (bitfield32) reload actions
	// performed
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadActionsPerformed DevlinkAttr = 154

	// DevlinkAttrEnumReloadLimits - DEVLINK_ATTR_RELOAD_LIMITS (bitfield32) reload limits
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadLimits DevlinkAttr = 155

	// DevlinkAttrEnumDevStats - DEVLINK_ATTR_DEV_STATS (nested) device statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumDevStats DevlinkAttr = 156

	// DevlinkAttrEnumReloadStats - DEVLINK_ATTR_RELOAD_STATS (nested) reload statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadStats DevlinkAttr = 157

	// DevlinkAttrEnumReloadStatsEntry - DEVLINK_ATTR_RELOAD_STATS_ENTRY (nested) reload statistics entry
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadStatsEntry DevlinkAttr = 158

	// DevlinkAttrEnumReloadStatsLimit - DEVLINK_ATTR_RELOAD_STATS_LIMIT (u8) reload statistics limit
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadStatsLimit DevlinkAttr = 159

	// DevlinkAttrEnumReloadStatsValue - DEVLINK_ATTR_RELOAD_STATS_VALUE (u32) reload statistics value
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadStatsValue DevlinkAttr = 160

	// DevlinkAttrEnumRemoteReloadStats - DEVLINK_ATTR_REMOTE_RELOAD_STATS (nested) remote reload statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRemoteReloadStats DevlinkAttr = 161

	// DevlinkAttrEnumReloadActionInfo - DEVLINK_ATTR_RELOAD_ACTION_INFO (nested) reload action information
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadActionInfo DevlinkAttr = 162

	// DevlinkAttrEnumReloadActionStats - DEVLINK_ATTR_RELOAD_ACTION_STATS (nested) reload action statistics
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumReloadActionStats DevlinkAttr = 163

	// DevlinkAttrEnumPortPciSfNumber - DEVLINK_ATTR_PORT_PCI_SF_NUMBER (u32) port PCI SF number
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumPortPciSfNumber DevlinkAttr = 164

	// DevlinkAttrEnumRateType - DEVLINK_ATTR_RATE_TYPE (u16) rate type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateType DevlinkAttr = 165

	// DevlinkAttrEnumRateTxShare - DEVLINK_ATTR_RATE_TX_SHARE (u64) rate TX share
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateTxShare DevlinkAttr = 166

	// DevlinkAttrEnumRateTxMax - DEVLINK_ATTR_RATE_TX_MAX (u64) rate TX maximum
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateTxMax DevlinkAttr = 167

	// DevlinkAttrEnumRateNodeName - DEVLINK_ATTR_RATE_NODE_NAME (string) rate node name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateNodeName DevlinkAttr = 168

	// DevlinkAttrEnumRateParentNodeName - DEVLINK_ATTR_RATE_PARENT_NODE_NAME (string) rate parent node name
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateParentNodeName DevlinkAttr = 169

	// DevlinkAttrEnumRegionMaxSnapshots - DEVLINK_ATTR_REGION_MAX_SNAPSHOTS (u32) region maximum snapshots
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionMaxSnapshots DevlinkAttr = 170

	// DevlinkAttrEnumLinecardIndex - DEVLINK_ATTR_LINECARD_INDEX (u32) linecard index
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumLinecardIndex DevlinkAttr = 171

	// DevlinkAttrEnumLinecardState - DEVLINK_ATTR_LINECARD_STATE (u8) linecard state
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumLinecardState DevlinkAttr = 172

	// DevlinkAttrEnumLinecardType - DEVLINK_ATTR_LINECARD_TYPE (string) linecard type
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumLinecardType DevlinkAttr = 173

	// DevlinkAttrEnumLinecardSupportedTypes - DEVLINK_ATTR_LINECARD_SUPPORTED_TYPES (nested) linecard supported types
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumLinecardSupportedTypes DevlinkAttr = 174

	// DevlinkAttrEnumNestedDevlink - DEVLINK_ATTR_NESTED_DEVLINK (nested) nested devlink
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumNestedDevlink DevlinkAttr = 175

	// DevlinkAttrEnumSelftests - DEVLINK_ATTR_SELFTESTS (nested) selftests
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumSelftests DevlinkAttr = 176

	// DevlinkAttrEnumRateTxPriority - DEVLINK_ATTR_RATE_TX_PRIORITY (u32) rate TX priority
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateTxPriority DevlinkAttr = 177

	// DevlinkAttrEnumRateTxWeight - DEVLINK_ATTR_RATE_TX_WEIGHT (u32) rate TX weight
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRateTxWeight DevlinkAttr = 178

	// DevlinkAttrEnumRegionDirect - DEVLINK_ATTR_REGION_DIRECT (flag) region direct
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumRegionDirect DevlinkAttr = 179

	// DevlinkAttrEnumMax - DEVLINK_ATTR_MAX - maximum value for devlink attributes
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkAttrEnumMax = iota - 1
)
