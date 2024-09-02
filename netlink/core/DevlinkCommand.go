package core

// DevlinkCommand - devlink_command
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkCommand uint8

/*
warning: don't change the order or add anything between this is ABI!
*/
const (

	//DevlinkCmdUnspec - DEVLINK_CMD_UNSPEC
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdUnspec DevlinkCommand = 0

	// DevlinkCmdGet - DEVLINK_CMD_GET - can dump
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdGet DevlinkCommand = 1

	// DevlinkCmdSet - DEVLINK_CMD_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSet DevlinkCommand = 2

	// DevlinkCmdNew - DEVLINK_CMD_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdNew DevlinkCommand = 3

	// DevlinkCmdDel - DEVLINK_CMD_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDel DevlinkCommand = 3

	// DevlinkCmdPortGet - DEVLINK_CMD_PORT_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortGet DevlinkCommand = 4

	// DevlinkCmdPortSet - DEVLINK_CMD_PORT_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortSet DevlinkCommand = 5

	// DevlinkCmdPortNew - DEVLINK_CMD_PORT_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortNew DevlinkCommand = 6

	// DevlinkCmdPortDel - DEVLINK_CMD_PORT_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortDel DevlinkCommand = 7

	// DevlinkCmdPortSplit - DEVLINK_CMD_PORT_SPLIT
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortSplit DevlinkCommand = 8

	// DevlinkCmdPortUnsplit - DEVLINK_CMD_PORT_UNSPLIT
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortUnsplit DevlinkCommand = 9

	// DevlinkCmdSbGet - DEVLINK_CMD_SB_GET - /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbGet DevlinkCommand = 10

	// DevlinkCmdSbSet - DEVLINK_CMD_SB_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbSet DevlinkCommand = 11

	// DevlinkCmdSbNew - DEVLINK_CMD_SB_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbNew DevlinkCommand = 12

	// DevlinkCmdSbDel - DEVLINK_CMD_SB_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbDel DevlinkCommand = 13

	// DevlinkCmdSbPoolGet - DEVLINK_CMD_SB_POOL_GET - /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolGet DevlinkCommand = 14

	// DevlinkCmdSbPoolSet - DEVLINK_CMD_SB_POOL_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolSet DevlinkCommand = 15

	// DevlinkCmdSbPoolNew - DEVLINK_CMD_SB_POOL_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolNew DevlinkCommand = 16

	// DevlinkCmdSbPoolDel - DEVLINK_CMD_SB_POOL_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPoolDel DevlinkCommand = 17

	// DevlinkCmdSbPortPoolGet - DEVLINK_CMD_SB_PORT_POOL_GET - /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolGet DevlinkCommand = 18

	// DevlinkCmdSbPortPoolSet - DEVLINK_CMD_SB_PORT_POOL_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolSet DevlinkCommand = 19

	// DevlinkCmdSbPortPoolNew - DEVLINK_CMD_SB_PORT_POOL_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolNew DevlinkCommand = 20

	// DevlinkCmdSbPortPoolDel - DEVLINK_CMD_SB_PORT_POOL_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbPortPoolDel DevlinkCommand = 21

	// DevlinkCmdSbTcPoolBindGet - DEVLINK_CMD_SB_TC_POOL_BIND_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindGet DevlinkCommand = 22

	// DevlinkCmdSbTcPoolBindSet - DEVLINK_CMD_SB_TC_POOL_BIND_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindSet DevlinkCommand = 23

	// DevlinkCmdSbTcPoolBindNew - DEVLINK_CMD_SB_TC_POOL_BIND_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindNew DevlinkCommand = 24

	// DevlinkCmdSbTcPoolBindDel - DEVLINK_CMD_SB_TC_POOL_BIND_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbTcPoolBindDel DevlinkCommand = 25

	// DevlinkCmdSbOccSnapshot - DEVLINK_CMD_SB_OCC_SNAPSHOT - Shared buffer occupancy monitoring command
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbOccSnapshot DevlinkCommand = 26

	// DevlinkCmdSbOccMaxClear - DEVLINK_CMD_SB_OCC_MAX_CLEAR - Shared buffer occupancy monitoring command
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSbOccMaxClear DevlinkCommand = 27

	// DevlinkCmdEswitchGet - DEVLINK_CMD_ESWITCH_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchGet DevlinkCommand = 28

	// DevlinkCmdEswitchModeGet - DEVLINK_CMD_ESWITCH_MODE_GET /* obsolete never use this! */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchModeGet = DevlinkCmdEswitchGet

	// DevlinkCmdEswitchSet - DEVLINK_CMD_ESWITCH_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchSet DevlinkCommand = 29

	// DevlinkCmdEswitchModeSet - DEVLINK_CMD_ESWITCH_MODE_SET /* obsolete never use this! */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdEswitchModeSet = DevlinkCmdEswitchSet

	// DevlinkCmdDpipeTableGet - DEVLINK_CMD_DPIPE_TABLE_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeTableGet DevlinkCommand = 30

	// DevlinkCmdDpipeEntriesGet - DEVLINK_CMD_DPIPE_ENTRIES_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeEntriesGet DevlinkCommand = 31

	// DevlinkCmdDpipeHeadersGet - DEVLINK_CMD_DPIPE_HEADERS_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeHeadersGet DevlinkCommand = 32

	// DevlinkCmdDpipeTableCountersSet - DEVLINK_CMD_DPIPE_TABLE_COUNTERS_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdDpipeTableCountersSet DevlinkCommand = 33

	// DevlinkCmdResourceSet - DEVLINK_CMD_RESOURCE_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdResourceSet DevlinkCommand = 34

	// DevlinkCmdResourceDump - DEVLINK_CMD_RESOURCE_DUMP
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdResourceDump DevlinkCommand = 35

	// DevlinkCmdReload - DEVLINK_CMD_RELOAD - Hot driver reload makes configuration changes take place. The
	// devlink instance is not released during the process.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdReload DevlinkCommand = 36

	// DevlinkCmdParamGet - DEVLINK_CMD_PARAM_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamGet DevlinkCommand = 37

	// DevlinkCmdParamSet - DEVLINK_CMD_PARAM_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamSet DevlinkCommand = 38

	// DevlinkCmdParamNew - DEVLINK_CMD_PARAM_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamNew DevlinkCommand = 39

	// DevlinkCmdParamDel - DEVLINK_CMD_PARAM_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdParamDel DevlinkCommand = 40

	// DevlinkCmdRegionGet - DEVLINK_CMD_REGION_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionGet DevlinkCommand = 41

	// DevlinkCmdRegionSet - DEVLINK_CMD_REGION_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionSet DevlinkCommand = 42

	// DevlinkCmdRegionNew - DEVLINK_CMD_REGION_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionNew DevlinkCommand = 43

	// DevlinkCmdRegionDel - DEVLINK_CMD_REGION_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionDel DevlinkCommand = 44

	// DevlinkCmdRegionRead - DEVLINK_CMD_REGION_READ
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRegionRead DevlinkCommand = 45

	// DevlinkCmdPortParamGet - DEVLINK_CMD_PORT_PARAM_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamGet DevlinkCommand = 46

	// DevlinkCmdPortParamSet - DEVLINK_CMD_PORT_PARAM_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamSet DevlinkCommand = 47

	// DevlinkCmdPortParamNew - DEVLINK_CMD_PORT_PARAM_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamNew DevlinkCommand = 48

	// DevlinkCmdPortParamDel - DEVLINK_CMD_PORT_PARAM_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdPortParamDel DevlinkCommand = 49

	// DevlinkCmdInfoGet - DEVLINK_CMD_INFO_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdInfoGet DevlinkCommand = 50

	// DevlinkCmdHealthReporterGet - DEVLINK_CMD_HEALTH_REPORTER_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterGet DevlinkCommand = 51

	// DevlinkCmdHealthReporterSet - DEVLINK_CMD_HEALTH_REPORTER_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterSet DevlinkCommand = 52

	// DevlinkCmdHealthReporterRecover - DEVLINK_CMD_HEALTH_REPORTER_RECOVER
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterRecover DevlinkCommand = 53

	// DevlinkCmdHealthReporterDiagnose - DEVLINK_CMD_HEALTH_REPORTER_DIAGNOSE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterDiagnose DevlinkCommand = 54

	// DevlinkCmdHealthReporterDumpGet - DEVLINK_CMD_HEALTH_REPORTER_DUMP_GET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterDumpGet DevlinkCommand = 55

	// DevlinkCmdHealthReporterDumpClear - DEVLINK_CMD_HEALTH_REPORTER_DUMP_CLEAR
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterDumpClear DevlinkCommand = 56

	// DevlinkCmdFlashUpdate - DEVLINK_CMD_FLASH_UPDATE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdFlashUpdate DevlinkCommand = 57

	// DevlinkCmdFlashUpdateEnd - DEVLINK_CMD_FLASH_UPDATE_END /* notification only */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdFlashUpdateEnd DevlinkCommand = 58

	// DevlinkCmdFlashUpdateStatus - DEVLINK_CMD_FLASH_UPDATE_STATUS /* notification only */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdFlashUpdateStatus DevlinkCommand = 59

	// DevlinkCmdTrapGet - DEVLINK_CMD_TRAP_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGet DevlinkCommand = 60

	// DevlinkCmdTrapSet - DEVLINK_CMD_TRAP_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapSet DevlinkCommand = 61

	// DevlinkCmdTrapNew - DEVLINK_CMD_TRAP_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapNew DevlinkCommand = 62

	// DevlinkCmdTrapDel - DEVLINK_CMD_TRAP_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapDel DevlinkCommand = 63

	// DevlinkCmdTrapGroupGet - DEVLINK_CMD_TRAP_GROUP_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupGet DevlinkCommand = 64

	// DevlinkCmdTrapGroupSet - DEVLINK_CMD_TRAP_GROUP_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupSet DevlinkCommand = 65

	// DevlinkCmdTrapGroupNew - DEVLINK_CMD_TRAP_GROUP_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupNew DevlinkCommand = 66

	// DevlinkCmdTrapGroupDel - DEVLINK_CMD_TRAP_GROUP_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapGroupDel DevlinkCommand = 67

	// DevlinkCmdTrapPolicerGet - DEVLINK_CMD_TRAP_POLICER_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerGet DevlinkCommand = 68

	// DevlinkCmdTrapPolicerSet - DEVLINK_CMD_TRAP_POLICER_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerSet DevlinkCommand = 69

	// DevlinkCmdTrapPolicerNew - DEVLINK_CMD_TRAP_POLICER_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerNew DevlinkCommand = 70

	// DevlinkCmdTrapPolicerDel - DEVLINK_CMD_TRAP_POLICER_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdTrapPolicerDel DevlinkCommand = 71

	// DevlinkCmdHealthReporterTest - DEVLINK_CMD_HEALTH_REPORTER_TEST
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdHealthReporterTest DevlinkCommand = 72

	// DevlinkCmdRateGet - DEVLINK_CMD_RATE_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateGet DevlinkCommand = 73

	// DevlinkCmdRateSet - DEVLINK_CMD_RATE_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateSet DevlinkCommand = 74

	// DevlinkCmdRateNew - DEVLINK_CMD_RATE_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateNew DevlinkCommand = 75

	// DevlinkCmdRateDel - DEVLINK_CMD_RATE_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdRateDel DevlinkCommand = 76

	// DevlinkCmdLinecardGet - DEVLINK_CMD_LINECARD_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardGet DevlinkCommand = 77

	// DevlinkCmdLinecardSet - DEVLINK_CMD_LINECARD_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardSet DevlinkCommand = 78

	// DevlinkCmdLinecardNew - DEVLINK_CMD_LINECARD_NEW
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardNew DevlinkCommand = 79

	// DevlinkCmdLinecardDel - DEVLINK_CMD_LINECARD_DEL
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdLinecardDel DevlinkCommand = 80

	// DevlinkCmdSelftestsGet - DEVLINK_CMD_SELFTESTS_GET /* can dump */
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSelftestsGet DevlinkCommand = 81

	// DevlinkCmdSelftestsRun - DEVLINK_CMD_SELFTESTS_RUN
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdSelftestsRun DevlinkCommand = 82

	// DevlinkCmdNotifyFilterSet - DEVLINK_CMD_NOTIFY_FILTER_SET
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkCmdNotifyFilterSet DevlinkCommand = 83

	/*


		add new commands above here


	*/

	// __devlinkCmdMax - __DEVLINK_CMD_MAX
	__devlinkCmdMax DevlinkCommand = 84

	// DevlinkCmdMax - DEVLINK_CMD_MAX
	DevlinkCmdMax = __devlinkCmdMax - 1
)
