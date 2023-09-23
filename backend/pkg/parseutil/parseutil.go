package parseutil

import (
	"github.com/fahmifan/ulids"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/samber/lo"
)

func MustParseULID(s string) ulids.ULID {
	return ulids.ULID{ULID: ulid.MustParse(s)}
}

func WeakParseULID(s string) ulids.ULID {
	id, err := ulid.Parse(s)
	if err != nil {
		return ulids.ULID{}
	}
	return ulids.ULID{ULID: id}
}

func WeakParseUUID(s string) uuid.UUID {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func ParseULID(in string) (ulids.ULID, error) {
	id, err := ulid.Parse(in)
	if err != nil {
		return ulids.ULID{}, err
	}
	return ulids.ULID{ULID: id}, nil
}

func StringULIDs(ids []ulids.ULID) []string {
	return lo.Map(ids, func(id ulids.ULID, index int) string {
		return id.String()
	})
}
