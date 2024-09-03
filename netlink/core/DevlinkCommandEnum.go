package core

// DevlinkCommandEnum - devlink_command
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkCommandEnum uint8

/*
warning: don't change the order or add anything between this is ABI!
*/
const (

	//DevlinkCmdUnspec - DEVLINK_CMD_UNSPEC
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdUnspec DevlinkCommandEnum = 0

	// DevlinkCmdGet - DEVLINK_CMD_GET - can dump
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdGet DevlinkCommandEnum = 1

	// DevlinkCmdSet - DEVLINK_CMD_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSet DevlinkCommandEnum = 2

	// DevlinkCmdNew - DEVLINK_CMD_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdNew DevlinkCommandEnum = 3

	// DevlinkCmdDel - DEVLINK_CMD_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDel DevlinkCommandEnum = 3

	// DevlinkCmdPortGet - DEVLINK_CMD_PORT_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortGet DevlinkCommandEnum = 4

	// DevlinkCmdPortSet - DEVLINK_CMD_PORT_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortSet DevlinkCommandEnum = 5

	// DevlinkCmdPortNew - DEVLINK_CMD_PORT_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortNew DevlinkCommandEnum = 6

	// DevlinkCmdPortDel - DEVLINK_CMD_PORT_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortDel DevlinkCommandEnum = 7

	// DevlinkCmdPortSplit - DEVLINK_CMD_PORT_SPLIT
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortSplit DevlinkCommandEnum = 8

	// DevlinkCmdPortUnsplit - DEVLINK_CMD_PORT_UNSPLIT
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortUnsplit DevlinkCommandEnum = 9

	// DevlinkCmdSbGet - DEVLINK_CMD_SB_GET - /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbGet DevlinkCommandEnum = 10

	// DevlinkCmdSbSet - DEVLINK_CMD_SB_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbSet DevlinkCommandEnum = 11

	// DevlinkCmdSbNew - DEVLINK_CMD_SB_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbNew DevlinkCommandEnum = 12

	// DevlinkCmdSbDel - DEVLINK_CMD_SB_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbDel DevlinkCommandEnum = 13

	// DevlinkCmdSbPoolGet - DEVLINK_CMD_SB_POOL_GET - /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolGet DevlinkCommandEnum = 14

	// DevlinkCmdSbPoolSet - DEVLINK_CMD_SB_POOL_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolSet DevlinkCommandEnum = 15

	// DevlinkCmdSbPoolNew - DEVLINK_CMD_SB_POOL_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolNew DevlinkCommandEnum = 16

	// DevlinkCmdSbPoolDel - DEVLINK_CMD_SB_POOL_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolDel DevlinkCommandEnum = 17

	// DevlinkCmdSbPortPoolGet - DEVLINK_CMD_SB_PORT_POOL_GET - /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolGet DevlinkCommandEnum = 18

	// DevlinkCmdSbPortPoolSet - DEVLINK_CMD_SB_PORT_POOL_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolSet DevlinkCommandEnum = 19

	// DevlinkCmdSbPortPoolNew - DEVLINK_CMD_SB_PORT_POOL_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolNew DevlinkCommandEnum = 20

	// DevlinkCmdSbPortPoolDel - DEVLINK_CMD_SB_PORT_POOL_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolDel DevlinkCommandEnum = 21

	// DevlinkCmdSbTcPoolBindGet - DEVLINK_CMD_SB_TC_POOL_BIND_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindGet DevlinkCommandEnum = 22

	// DevlinkCmdSbTcPoolBindSet - DEVLINK_CMD_SB_TC_POOL_BIND_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindSet DevlinkCommandEnum = 23

	// DevlinkCmdSbTcPoolBindNew - DEVLINK_CMD_SB_TC_POOL_BIND_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindNew DevlinkCommandEnum = 24

	// DevlinkCmdSbTcPoolBindDel - DEVLINK_CMD_SB_TC_POOL_BIND_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindDel DevlinkCommandEnum = 25

	// DevlinkCmdSbOccSnapshot - DEVLINK_CMD_SB_OCC_SNAPSHOT - Shared buffer occupancy monitoring command
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbOccSnapshot DevlinkCommandEnum = 26

	// DevlinkCmdSbOccMaxClear - DEVLINK_CMD_SB_OCC_MAX_CLEAR - Shared buffer occupancy monitoring command
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbOccMaxClear DevlinkCommandEnum = 27

	// DevlinkCmdEswitchGet - DEVLINK_CMD_ESWITCH_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchGet DevlinkCommandEnum = 28

	// DevlinkCmdEswitchModeGet - DEVLINK_CMD_ESWITCH_MODE_GET /* obsolete never use this! */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchModeGet = DevlinkCmdEswitchGet

	// DevlinkCmdEswitchSet - DEVLINK_CMD_ESWITCH_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchSet DevlinkCommandEnum = 29

	// DevlinkCmdEswitchModeSet - DEVLINK_CMD_ESWITCH_MODE_SET /* obsolete never use this! */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchModeSet = DevlinkCmdEswitchSet

	// DevlinkCmdDpipeTableGet - DEVLINK_CMD_DPIPE_TABLE_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeTableGet DevlinkCommandEnum = 30

	// DevlinkCmdDpipeEntriesGet - DEVLINK_CMD_DPIPE_ENTRIES_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeEntriesGet DevlinkCommandEnum = 31

	// DevlinkCmdDpipeHeadersGet - DEVLINK_CMD_DPIPE_HEADERS_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeHeadersGet DevlinkCommandEnum = 32

	// DevlinkCmdDpipeTableCountersSet - DEVLINK_CMD_DPIPE_TABLE_COUNTERS_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeTableCountersSet DevlinkCommandEnum = 33

	// DevlinkCmdResourceSet - DEVLINK_CMD_RESOURCE_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdResourceSet DevlinkCommandEnum = 34

	// DevlinkCmdResourceDump - DEVLINK_CMD_RESOURCE_DUMP
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdResourceDump DevlinkCommandEnum = 35

	// DevlinkCmdReload - DEVLINK_CMD_RELOAD - Hot driver reload makes configuration changes take place. The
	// devlink instance is not released during the process.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdReload DevlinkCommandEnum = 36

	// DevlinkCmdParamGet - DEVLINK_CMD_PARAM_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamGet DevlinkCommandEnum = 37

	// DevlinkCmdParamSet - DEVLINK_CMD_PARAM_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamSet DevlinkCommandEnum = 38

	// DevlinkCmdParamNew - DEVLINK_CMD_PARAM_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamNew DevlinkCommandEnum = 39

	// DevlinkCmdParamDel - DEVLINK_CMD_PARAM_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamDel DevlinkCommandEnum = 40

	// DevlinkCmdRegionGet - DEVLINK_CMD_REGION_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionGet DevlinkCommandEnum = 41

	// DevlinkCmdRegionSet - DEVLINK_CMD_REGION_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionSet DevlinkCommandEnum = 42

	// DevlinkCmdRegionNew - DEVLINK_CMD_REGION_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionNew DevlinkCommandEnum = 43

	// DevlinkCmdRegionDel - DEVLINK_CMD_REGION_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionDel DevlinkCommandEnum = 44

	// DevlinkCmdRegionRead - DEVLINK_CMD_REGION_READ
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionRead DevlinkCommandEnum = 45

	// DevlinkCmdPortParamGet - DEVLINK_CMD_PORT_PARAM_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamGet DevlinkCommandEnum = 46

	// DevlinkCmdPortParamSet - DEVLINK_CMD_PORT_PARAM_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamSet DevlinkCommandEnum = 47

	// DevlinkCmdPortParamNew - DEVLINK_CMD_PORT_PARAM_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamNew DevlinkCommandEnum = 48

	// DevlinkCmdPortParamDel - DEVLINK_CMD_PORT_PARAM_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamDel DevlinkCommandEnum = 49

	// DevlinkCmdInfoGet - DEVLINK_CMD_INFO_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdInfoGet DevlinkCommandEnum = 50

	// DevlinkCmdHealthReporterGet - DEVLINK_CMD_HEALTH_REPORTER_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterGet DevlinkCommandEnum = 51

	// DevlinkCmdHealthReporterSet - DEVLINK_CMD_HEALTH_REPORTER_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterSet DevlinkCommandEnum = 52

	// DevlinkCmdHealthReporterRecover - DEVLINK_CMD_HEALTH_REPORTER_RECOVER
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterRecover DevlinkCommandEnum = 53

	// DevlinkCmdHealthReporterDiagnose - DEVLINK_CMD_HEALTH_REPORTER_DIAGNOSE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterDiagnose DevlinkCommandEnum = 54

	// DevlinkCmdHealthReporterDumpGet - DEVLINK_CMD_HEALTH_REPORTER_DUMP_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterDumpGet DevlinkCommandEnum = 55

	// DevlinkCmdHealthReporterDumpClear - DEVLINK_CMD_HEALTH_REPORTER_DUMP_CLEAR
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterDumpClear DevlinkCommandEnum = 56

	// DevlinkCmdFlashUpdate - DEVLINK_CMD_FLASH_UPDATE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdFlashUpdate DevlinkCommandEnum = 57

	// DevlinkCmdFlashUpdateEnd - DEVLINK_CMD_FLASH_UPDATE_END /* notification only */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdFlashUpdateEnd DevlinkCommandEnum = 58

	// DevlinkCmdFlashUpdateStatus - DEVLINK_CMD_FLASH_UPDATE_STATUS /* notification only */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdFlashUpdateStatus DevlinkCommandEnum = 59

	// DevlinkCmdTrapGet - DEVLINK_CMD_TRAP_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGet DevlinkCommandEnum = 60

	// DevlinkCmdTrapSet - DEVLINK_CMD_TRAP_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapSet DevlinkCommandEnum = 61

	// DevlinkCmdTrapNew - DEVLINK_CMD_TRAP_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapNew DevlinkCommandEnum = 62

	// DevlinkCmdTrapDel - DEVLINK_CMD_TRAP_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapDel DevlinkCommandEnum = 63

	// DevlinkCmdTrapGroupGet - DEVLINK_CMD_TRAP_GROUP_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupGet DevlinkCommandEnum = 64

	// DevlinkCmdTrapGroupSet - DEVLINK_CMD_TRAP_GROUP_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupSet DevlinkCommandEnum = 65

	// DevlinkCmdTrapGroupNew - DEVLINK_CMD_TRAP_GROUP_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupNew DevlinkCommandEnum = 66

	// DevlinkCmdTrapGroupDel - DEVLINK_CMD_TRAP_GROUP_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupDel DevlinkCommandEnum = 67

	// DevlinkCmdTrapPolicerGet - DEVLINK_CMD_TRAP_POLICER_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerGet DevlinkCommandEnum = 68

	// DevlinkCmdTrapPolicerSet - DEVLINK_CMD_TRAP_POLICER_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerSet DevlinkCommandEnum = 69

	// DevlinkCmdTrapPolicerNew - DEVLINK_CMD_TRAP_POLICER_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerNew DevlinkCommandEnum = 70

	// DevlinkCmdTrapPolicerDel - DEVLINK_CMD_TRAP_POLICER_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerDel DevlinkCommandEnum = 71

	// DevlinkCmdHealthReporterTest - DEVLINK_CMD_HEALTH_REPORTER_TEST
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterTest DevlinkCommandEnum = 72

	// DevlinkCmdRateGet - DEVLINK_CMD_RATE_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateGet DevlinkCommandEnum = 73

	// DevlinkCmdRateSet - DEVLINK_CMD_RATE_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateSet DevlinkCommandEnum = 74

	// DevlinkCmdRateNew - DEVLINK_CMD_RATE_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateNew DevlinkCommandEnum = 75

	// DevlinkCmdRateDel - DEVLINK_CMD_RATE_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateDel DevlinkCommandEnum = 76

	// DevlinkCmdLinecardGet - DEVLINK_CMD_LINECARD_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardGet DevlinkCommandEnum = 77

	// DevlinkCmdLinecardSet - DEVLINK_CMD_LINECARD_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardSet DevlinkCommandEnum = 78

	// DevlinkCmdLinecardNew - DEVLINK_CMD_LINECARD_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardNew DevlinkCommandEnum = 79

	// DevlinkCmdLinecardDel - DEVLINK_CMD_LINECARD_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardDel DevlinkCommandEnum = 80

	// DevlinkCmdSelftestsGet - DEVLINK_CMD_SELFTESTS_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSelftestsGet DevlinkCommandEnum = 81

	// DevlinkCmdSelftestsRun - DEVLINK_CMD_SELFTESTS_RUN
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSelftestsRun DevlinkCommandEnum = 82

	// DevlinkCmdNotifyFilterSet - DEVLINK_CMD_NOTIFY_FILTER_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdNotifyFilterSet DevlinkCommandEnum = 83

	/*


		add new commands above here


	*/
	// DevlinkCmdMax - DEVLINK_CMD_MAX
	DevlinkCmdMax = DevlinkCmdNotifyFilterSet
)
