package validator

import (
	"errors"
	"fmt"
	"github.com/Pashgunt/converter/internal/enum"
	"github.com/Pashgunt/converter/internal/exception"
	"github.com/Pashgunt/converter/internal/resolver/alias/group/modify"
)

func ValidGroupAlias(rawData map[string]interface{}, alias string) error {
	if err := issetAliasConstraint(rawData, alias); err != nil {
		return err
	}

	data := modify.ConvertToNeedType(rawData, alias)

	if err := issetStructTagConstraint(data, alias); err != nil {
		return err
	}

	if err := issetPathTagConstraint(data, alias); err != nil {
		return err
	}

	return nil
}

func issetAliasConstraint(rawData map[string]interface{}, alias string) error {
	if _, isset := rawData[alias]; !isset {
		return errors.New(fmt.Sprintf(exception.InvalidGroupAliasException, alias))
	}

	return nil
}

func issetStructTagConstraint(data map[string]interface{}, alias string) error {
	if _, issetStructTag := data[enum.Struct]; !issetStructTag {
		return errors.New(fmt.Sprintf(exception.InvalidGroupAliasException, alias))
	}

	return nil
}

func issetPathTagConstraint(data map[string]interface{}, alias string) error {
	if _, issetPathTag := data[enum.Path]; !issetPathTag {
		return errors.New(fmt.Sprintf(exception.InvalidGroupAliasException, alias))
	}

	return nil
}
