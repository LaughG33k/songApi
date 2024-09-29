package repository

import "github.com/jackc/pgx"

func IsErrCode(err error, code string) bool {
	pgerr, ok := err.(pgx.PgError)
	return ok && pgerr.Code == code
}
