package pointers

func Int32ToPointer(i int32) *int32 {
	return &i
}

func Int64ToPointer(i int64) *int64 {
	return &i
}

func StringToPointer(s string) *string {
	return &s
}
