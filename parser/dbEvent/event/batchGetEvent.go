package event

type BatchGetEvent struct {
	BasicInfo *BasicEventInfo
	Keys      [][]byte
}

func (e *BatchGetEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *BatchGetEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
