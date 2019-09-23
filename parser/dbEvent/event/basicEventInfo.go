package event

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

type BasicEventInfo struct {
	Method MethodType
	DBName string
}
