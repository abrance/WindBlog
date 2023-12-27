package errors

const (
	SqliteDeleteError  = 2330
	SqliteSaveError    = 2320
	SqliteMigrateError = 2311
	SqliteInitError    = 2310
	SqliteError        = 2300

	FileDeleteError   = 2220
	FileNotExistError = 2210
	FileStorageError  = 2200

	JsonInsertError = 2121
	JsonSaveError   = 2120
	JsonDBInitError = 2110
	JsonDBError     = 2100

	StorageError = 2000
)
