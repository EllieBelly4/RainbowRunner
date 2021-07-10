package managers

var Players = NewPlayerManager()

type PlayerManager struct {
	//Players map[int]*game.RRConn
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{}
}
