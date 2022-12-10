package objects

type IActivatable interface {
	Activate(player *RRPlayer, id byte)
}
