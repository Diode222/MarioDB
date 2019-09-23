package event

type GetEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
}

func (e *GetEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *GetEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
