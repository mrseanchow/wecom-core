package callback

// 微盘容量不足事件
type WeDriveInsufficientCapacityEvent struct {
	BaseEvent
}

// 修改空间安全设置
type SpaceSecuritySettingsChangeEvent struct {
	BaseEvent
	SpaceId []string `xml:"SpaceId"`
}

// 修改空间成员
type SpaceMemberChangeEvent struct {
	BaseEvent
	SpaceId []string `xml:"SpaceId"`
}

// 解散空间
type DismissSpaceEvent struct {
	BaseEvent
	SpaceId []string `xml:"SpaceId"`
}

// 创建文件
type CreateFileEvent struct {
	BaseEvent
	FileId []string `xml:"FileId"`
}

// 重命名文件
type RenameFileEvent struct {
	BaseEvent
	FileId []string `xml:"FileId"`
}

// 更新文件内容
type UpdateFileEvent struct {
	BaseEvent
	FileId []string `xml:"FileId"`
}

// 删除文件
type DeleteFileEvent struct {
	BaseEvent
	FileId []string `xml:"FileId"`
}

// 移动文件
type MoveFileEvent struct {
	BaseEvent
	FileId []string `xml:"FileId"`
}

func ParseWeDriveInsufficientCapacity(data []byte) (*WeDriveInsufficientCapacityEvent, error) {
	return Parse[WeDriveInsufficientCapacityEvent](data)
}

func ParseSpaceSecuritySettingsChange(data []byte) (*SpaceSecuritySettingsChangeEvent, error) {
	return Parse[SpaceSecuritySettingsChangeEvent](data)
}

func ParseSpaceMemberChange(data []byte) (*SpaceMemberChangeEvent, error) {
	return Parse[SpaceMemberChangeEvent](data)
}

func ParseDismissSpace(data []byte) (*DismissSpaceEvent, error) {
	return Parse[DismissSpaceEvent](data)
}

func ParseCreateFile(data []byte) (*CreateFileEvent, error) {
	return Parse[CreateFileEvent](data)
}

func ParseRenameFile(data []byte) (*RenameFileEvent, error) {
	return Parse[RenameFileEvent](data)
}

func ParseUpdateFile(data []byte) (*UpdateFileEvent, error) {
	return Parse[UpdateFileEvent](data)
}

func ParseDeleteFile(data []byte) (*DeleteFileEvent, error) {
	return Parse[DeleteFileEvent](data)
}

func ParseMoveFile(data []byte) (*MoveFileEvent, error) {
	return Parse[MoveFileEvent](data)
}
