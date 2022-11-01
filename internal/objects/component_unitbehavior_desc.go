package objects

type UnitBehaviorDesc struct {
	*Component
}

func NewUnitBehaviorDesc(gctype string) *UnitBehaviorDesc {
	return &UnitBehaviorDesc{
		Component: NewComponent(gctype, "UnitBehaviorDesc"),
	}
}
