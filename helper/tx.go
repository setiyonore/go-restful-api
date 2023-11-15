package helper

import "database/sql"

func PanicOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			PanicIfError(errorCommit)
		}
	}
}
