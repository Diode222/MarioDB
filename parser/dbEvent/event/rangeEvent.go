package event

type RangeEvent struct {
	BasicInfo *BasicEventInfo
	Start     []byte
	Limit     []byte // Range contains this key
}

func (e *RangeEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *RangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
