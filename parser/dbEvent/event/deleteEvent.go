package event

type DeleteEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
}

func (e *DeleteEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *DeleteEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
