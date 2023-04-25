package constants

type EntryType string

const (
	Added       EntryType = "Added"
	ModifiedNew EntryType = "Modified_New"
	ModifiedOld EntryType = "Modified_Old"
	Deleted     EntryType = "Deleted"
)

type (
	MapPathHash      map[string]string
	MapEntryPathHash map[EntryType]MapPathHash
)

const (
	FileHashesName     = "file_hashes.csv"
	ComparedHashesName = "compared_hashes.csv"
)
