package event

type SeekRangeEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
}

func (e *SeekRangeEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *SeekRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
