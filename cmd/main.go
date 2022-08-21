package main

import (
	"hackattic_solutions/pkg/backup_restore"
	unpackbytes "hackattic_solutions/pkg/unpack_bytes"
)

func main() {

	backup_restore.BackupRestore()

	unpackbytes.Unpack()
}
