//go:build linux

package core

const (
	_ = iota
	IFLA_VF_INFO
	IFLA_VF_INFO_MAX = IFLA_VF_INFO
)

const (
	_           = iota
	IFLA_VF_MAC /* Hardware queue specific attributes */
	IFLA_VF_VLAN
	IFLA_VF_TX_RATE      /* Max TX Bandwidth Allocation */
	IFLA_VF_SPOOFCHK     /* Spoof Checking on/off switch */
	IFLA_VF_LINK_STATE   /* link state enable/disable/auto switch */
	IFLA_VF_RATE         /* Min and Max TX Bandwidth Allocation */
	IFLA_VF_RSS_QUERY_EN /* RSS Redirection Table and Hash Key query on/off switch*/
	IFLA_VF_STATS        /* network device statistics */
	IFLA_VF_TRUST        /* Trust state of VF */
	IFLA_VF_IB_NODE_GUID /* VF Infiniband node GUID */
	IFLA_VF_IB_PORT_GUID /* VF Infiniband port GUID */
	IFLA_VF_VLAN_LIST    /* nested list of vlans, option for QinQ */

	IFLA_VF_MAX = IFLA_VF_IB_PORT_GUID
)

const (
	_                 = iota
	IFLA_VF_VLAN_INFO /* VLAN ID, QoS and VLAN protocol */
	__IFLA_VF_VLAN_INFO_MAX
)

const (
	IFLA_VF_LINK_STATE_AUTO    = iota /* link state of the uplink */
	IFLA_VF_LINK_STATE_ENABLE         /* link always up */
	IFLA_VF_LINK_STATE_DISABLE        /* link always down */
	IFLA_VF_LINK_STATE_MAX     = IFLA_VF_LINK_STATE_DISABLE
)

const (
	IFLA_VF_STATS_RX_PACKETS = iota
	IFLA_VF_STATS_TX_PACKETS
	IFLA_VF_STATS_RX_BYTES
	IFLA_VF_STATS_TX_BYTES
	IFLA_VF_STATS_BROADCAST
	IFLA_VF_STATS_MULTICAST
	IFLA_VF_STATS_RX_DROPPED
	IFLA_VF_STATS_TX_DROPPED
	IFLA_VF_STATS_MAX = IFLA_VF_STATS_TX_DROPPED
)
