package objects

type IActivatable interface {
	Activate(player *RRPlayer, u *UnitBehavior, id byte)
}
