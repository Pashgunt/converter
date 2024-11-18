package group

import (
	"fmt"
	"github.com/Pashgunt/converter/config"
	"github.com/Pashgunt/converter/internal/enum"
	"github.com/Pashgunt/converter/internal/helper"
	"github.com/Pashgunt/converter/internal/resolver/alias/group/modify"
	"github.com/Pashgunt/converter/internal/resolver/alias/group/validator"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

const delimiter = ","

func existGroup(alias string) bool {
	if _, err := os.Stat(getGroupFilePath(alias)); err != nil {
		return false
	}

	return true
}

func getGroupFilePath(alias string) string {
	return fmt.Sprintf("%s/%s.yaml", config.PathGroup, alias)
}

func GetGroups(groupAlias []string) ([]string, error) {
	var groups []string

	for _, alias := range groupAlias {
		if !existGroup(alias) {
			continue
		}

		data, err := os.ReadFile(getGroupFilePath(alias))

		if err != nil {
			return nil, err
		}

		rawData := map[string]interface{}{}

		if err = yaml.Unmarshal(data, &rawData); err != nil {
			return nil, err
		}

		if err = validator.ValidGroupAlias(rawData, alias); err != nil {
			return nil, err
		}

		convertedData := modify.ConvertToNeedType(rawData, alias)

		var groupPrefix string = getGroupPrefix(convertedData[enum.Struct].(string))

		switch pathDataByType := convertedData[enum.Path].(type) {
		case []interface{}:
			for index := 0; index < len(pathDataByType); index++ {
				groups = append(groups, getPaths(pathDataByType[index].(string), []string{}, groupPrefix, 0)...)
			}
		}
	}

	return groups, nil
}

func getGroupPrefix(structName string) string {
	match := regexp.MustCompile(enum.GroupPrefix).FindStringSubmatch(structName)

	if len(match) > 1 {
		return helper.CamelToSnake(fmt.Sprintf("%s__", match[1]))
	}

	return ""
}

func getPaths(pathDataByType string, groups []string, groupPrefix string, offset int) []string {
	matches := regexp.MustCompile(getRegexByOffset(offset)).FindAllString(pathDataByType[offset:], -1)

	if matches != nil {
		offset += len(matches[0])

		for _, match := range strings.Split(matches[0][1:len(matches[0])-1], delimiter) {
			groups = append(groups, groupPrefix+match)
		}
	}

	if matches = regexp.MustCompile(enum.GroupPrefixList).FindAllString(pathDataByType[offset:], -1); matches != nil {
		offset += len(matches[0])

		for _, match := range strings.Split(matches[0], delimiter) {
			groups = getPaths(pathDataByType, groups, getGroupPrefix(match), offset)
		}
	}

	return groups
}

func getRegexByOffset(offset int) string {
	if offset == 0 {
		return enum.GroupNameDetailInit
	}

	return enum.GroupNameDetail
}
