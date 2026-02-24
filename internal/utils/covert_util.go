package utils

import "github.com/google/uuid"

func StringPtrToUUIDPtr(str *string) *uuid.UUID {

	if str == nil || *str == "" {
		return nil
	}
	u := uuid.MustParse(*str)

	return &u
}

func StringToUUID(str string) uuid.UUID {
	return uuid.MustParse(str)
}
