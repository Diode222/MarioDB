package event

type PrefixRangeEvent struct {
	BasicInfo *BasicEventInfo
	Prefix    []byte
}

func (e *PrefixRangeEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *PrefixRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
