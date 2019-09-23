package event

type BatchDeleteEvent struct {
	BasicInfo *BasicEventInfo
	Keys      [][]byte
}

func (e *BatchDeleteEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *BatchDeleteEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
