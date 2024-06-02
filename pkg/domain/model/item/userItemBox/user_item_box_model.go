package userItemBox

// SetExistingItemMaps 過去に受け取ったことがあるアイテムをboolにする
func (s UserItemBoxes) SetExistingItemMaps() map[int64]bool {
	maps := make(map[int64]bool)
	for _, userItemBoxModel := range s {
		maps[userItemBoxModel.MasterItemId] = true
	}

	return maps
}
