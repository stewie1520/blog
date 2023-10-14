package types

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}

func NewNullString(nilString *string) NullString {
	if nilString == nil {
		return NullString{
			sql.NullString{
				String: "",
				Valid:  false,
			},
		}
	}

	return NullString{
		sql.NullString{
			String: *nilString,
			Valid:  true,
		},
	}
}
