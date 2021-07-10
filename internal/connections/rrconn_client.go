package connections

import (
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type RRConnClient struct {
	ID        int
	Conn      *RRConn
	Zone      string
	IsSpawned bool

	TicksSinceLastUpdate int
	MoveQueue            *datatypes.MovementUpdateQueue
	LastPosition         pkg.Vector3
}

func (p *RRConnClient) GetID() int {
	return p.ID
}

func (p *RRConnClient) Send(b *byter.Byter) error {
	_, err := p.Conn.NetConn.Write(b.Data())

	return err
}

func NewRRConnClient(ID int, conn *RRConn) (p *RRConnClient) {
	p = &RRConnClient{
		ID:        ID,
		Conn:      conn,
		MoveQueue: datatypes.NewMovementUpdateQueue(1024),
	}

	return
}

//CrashLog: ClientEntityManager::processComponentUpdate ERROR: Invalid ComponentID(5) from server. UpdateID(100)

func (p *RRConnClient) Tick() {
}
