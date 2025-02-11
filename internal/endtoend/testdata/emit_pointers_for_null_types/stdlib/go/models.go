// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package datatype

import (
	"database/sql"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type DtCharacter struct {
	A sql.NullString
	B sql.NullString
	C sql.NullString
	D sql.NullString
	E sql.NullString
}

type DtCharacterNotNull struct {
	A string
	B string
	C string
	D string
	E string
}

type DtDatetime struct {
	A sql.NullTime
	B sql.NullTime
	C sql.NullTime
	D sql.NullTime
	E sql.NullTime
	F sql.NullTime
	G sql.NullTime
	H sql.NullTime
}

type DtDatetimeNotNull struct {
	A time.Time
	B time.Time
	C time.Time
	D time.Time
	E time.Time
	F time.Time
	G time.Time
	H time.Time
}

type DtNetType struct {
	A pqtype.Inet
	B pqtype.CIDR
	C pqtype.Macaddr
}

type DtNetTypesNotNull struct {
	A pqtype.Inet
	B pqtype.CIDR
	C pqtype.Macaddr
}

type DtNumeric struct {
	A sql.NullInt16
	B sql.NullInt32
	C sql.NullInt64
	D sql.NullString
	E sql.NullString
	F sql.NullFloat64
	G sql.NullFloat64
	H sql.NullInt16
	I sql.NullInt32
	J sql.NullInt64
	K sql.NullInt16
	L sql.NullInt32
	M sql.NullInt64
}

type DtNumericNotNull struct {
	A int16
	B int32
	C int64
	D string
	E string
	F float32
	G float64
	H int16
	I int32
	J int64
	K int16
	L int32
	M int64
}

type DtRange struct {
	A interface{}
	B interface{}
	C interface{}
	D interface{}
	E interface{}
	F interface{}
}

type DtRangeNotNull struct {
	A interface{}
	B interface{}
	C interface{}
	D interface{}
	E interface{}
	F interface{}
}
