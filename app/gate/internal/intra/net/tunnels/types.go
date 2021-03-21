package tunnels

type TunnelType int32

const (
	PlayerTunnelType = TunnelType(iota)
	RoomTunnelType
	TeamTunnelType
	FightTunnelType
	ChatTunnelType
	MailTunnelType
)
