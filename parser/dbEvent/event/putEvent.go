package event

type PutEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
	Value     []byte
}

func (e *PutEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *PutEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
