package sanitizer

import (
	"errors"
	"loket-app/modules/location/model"
	"strings"
)

// ValidateLocation - function for validating create location request.
func ValidateLocation(payload *model.CreateLocationReq) error {

	if payload == nil {
		err := errors.New("invalid payload")
		return err
	}

	if strings.TrimSpace(payload.Name) == "" {
		err := errors.New("name is required")
		return err
	}

	if strings.TrimSpace(payload.Address) == "" {
		err := errors.New("address is required")
		return err
	}

	if strings.TrimSpace(payload.Province) == "" {
		err := errors.New("province name is required")
		return err
	}

	return nil
}
