package parser

type MethodType int

const (
	OPEN             MethodType = iota // Open db with db name, if no exists, create db with db name
	GET                                // Read with key
	BATCHGET                           // Batch read with multi keys
	PUT                                // Write with a kv pair
	BATCHPUT                           // Batch write with multi kv pairs
	DELETE                             // Delete with key
	BATCHDELETE                        // Batch delete with multi keys
	RANGE                              // Range db with start and limit parameters
	BATCHRANGE                         // Batch range with multi start and limit parameters pairs
	SEEKRANGE                          // Seek a data with key and start range from this key
	PREFIXRANGE                        // Range with prefix
	BATCHPREFIXRANGE                   // Batch prefix range with multi prefixes
)

type DBEventBasicInfo struct {
	Method MethodType
	DBName string
}

type OpenEventInfo struct {
	BasicInfo *DBEventBasicInfo
}

type GetEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Key       []byte
}

type BatchGetEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Keys      [][]byte
}

type PutEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Key       []byte
	Value     []byte
}

type BatchPutEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Keys      [][]byte
	Values    [][]byte
}

type DeleteEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Key       []byte
}

type BatchDeleteEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Keys      [][]byte
}

type RangeEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Start     []byte
	Limit     []byte // Range contains this key
}

type BatchRangeEventInfo struct {
	BasicInfo *DBEventBasicInfo
	// Need to promise len(Starts) == len(Limits)
	Starts [][]byte
	Limits [][]byte // Range contains this key
}

type SeekRangeEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Key       []byte
}

type PrefixRangeEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Prefix    []byte
}

type BatchPrefixRangeEventInfo struct {
	BasicInfo *DBEventBasicInfo
	Prefixes  [][]byte
}
