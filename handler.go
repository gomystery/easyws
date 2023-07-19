package easyws

type IEasyWs interface {
	OnStart() error

	OnConnect() error

	OnUpgraded() error

	OnReceive(msg []byte) ([]byte, error)

	OnShutdown() error

	OnClose(err error) error

	// todo to add more
}
