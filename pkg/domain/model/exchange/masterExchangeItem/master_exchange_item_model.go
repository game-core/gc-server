package masterExchangeItem

// ExtractMasterExchangeId masterExchangeIdを抽出する
func (s MasterExchangeItems) ExtractMasterExchangeId() int64 {
	for _, masterExchangeItem := range s {
		if masterExchangeItem != nil {
			return masterExchangeItem.MasterItemId
		}
	}

	return 0
}
