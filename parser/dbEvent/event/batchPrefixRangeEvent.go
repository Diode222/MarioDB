package event

type BatchPrefixRangeEvent struct {
	BasicInfo *BasicEventInfo
	Prefixes  [][]byte
}

func (e *BatchPrefixRangeEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *BatchPrefixRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
