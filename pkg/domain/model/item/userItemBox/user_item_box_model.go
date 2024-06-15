package userItemBox

// SetUserItemBoxMaps Mapにする（key=masterItemId, value=UserItemBox）
func (s UserItemBoxes) SetUserItemBoxMaps() map[int64]*UserItemBox {
	maps := make(map[int64]*UserItemBox)
	for _, userItemBoxModel := range s {
		maps[userItemBoxModel.MasterItemId] = userItemBoxModel
	}

	return maps
}

// SetUserItemBoxExistingMaps 過去に受け取ったことがあるアイテムをboolにする（key=masterItemId, value=過去に受け取ったことがある）
func (s UserItemBoxes) SetUserItemBoxExistingMaps() map[int64]bool {
	maps := make(map[int64]bool)
	for _, userItemBoxModel := range s {
		maps[userItemBoxModel.MasterItemId] = true
	}

	return maps
}
