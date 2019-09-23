package event

type BatchPutEvent struct {
	BasicInfo *BasicEventInfo
	Keys      [][]byte
	Values    [][]byte
}

func (e *BatchPutEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *BatchPutEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
