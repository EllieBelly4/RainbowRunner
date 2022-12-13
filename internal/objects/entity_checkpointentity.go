package objects

//go:generate go run ../../scripts/generatelua -type=CheckpointEntity -extends=WorldEntity
type CheckpointEntity struct {
	*WorldEntity
}

func NewCheckpointEntity(gctype string) *CheckpointEntity {
	worldEntity := NewWorldEntity(gctype)

	worldEntity.WorldEntityFlags = 0x07

	return &CheckpointEntity{
		WorldEntity: worldEntity,
	}
}
