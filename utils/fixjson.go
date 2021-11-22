package utils

import (
	"encoding/json"
	"fmt"
	"magento2go/models"
	"regexp"
	"strings"
)

// FixJsonResponse
// one of the problems with the magento api is that it
// overloads the value:string fields with value:[]string
// and that breaks the json unmarshaller
// to fix this we use this function which intercepts the json string
// before it is json unmarshalled and alters the response
// flatten []string into a basic string
// EXPECTED:
//		"value":"4"
// ACTUAL:
// 		"value":["4","5","6"]
// CORRECTED:
// 		"value":"4,5,6"
func FlattenValueArrays(body string) string {
	reg := regexp.MustCompile(`"value":\[(".*?")\]`)
	strs := reg.FindAllString(body, -1)
	for _, match := range strs {
		reg = regexp.MustCompile(`"value":\[(".*?")\]`)
		toFix := reg.ReplaceAllString(match, "$1")
		leftMostMatch := fmt.Sprintf(`"value":"%s"`, strings.ReplaceAll(toFix, `"`, ``))
		body = reg.ReplaceAllString(string(body), leftMostMatch)
		// fmt.Println("FixResonseJson: ", leftMostMatch)
	}
	return body
}

// FixErrorParameters
// the error parameters come back from magento as objects
// rather than arrays of objects as their swagger docs specify
// EXPECTED:
// 		"parameters":[{"fieldName":"max","fieldValue":"300"}]
// ACTUAL:
// 		"parameters":{"max":300}
// CORRECTED:
// 		"parameters":[{"fieldName":"max","fieldValue":"300"}]
// however the
func FixErrorParameters(body string) string {

	var paramItems []*models.ErrorParametersItem

	reg := regexp.MustCompile(`"parameters":\{(.*?)\}`)
	strs := reg.FindAllString(body, -1)
	for _, match := range strs {
		reg = regexp.MustCompile(`"parameters":\{(.*?)\}`)
		toFix := reg.ReplaceAllString(match, "$1")
		toFix = strings.ReplaceAll(toFix, `"`, ``)
		items := strings.Split(toFix, ",")
		for _, item := range items {
			paramItem := &models.ErrorParametersItem{}
			part := strings.Split(item, ":")
			paramItem.FieldName = part[0]
			paramItem.FieldValue = part[1]
			paramItems = append(paramItems, paramItem)

			// replace any messages with the correct values
			if strings.Contains(body, "%"+paramItem.FieldName) {
				body = strings.Replace(body, "%"+paramItem.FieldName, paramItem.FieldValue, -1)
			}
		}
	}

	if len(paramItems) > 0 {
		b, err := json.Marshal(paramItems)
		if err != nil {
			return body
		}

		reg = regexp.MustCompile(`("parameters":)\{.*?\}`)
		fixed := reg.ReplaceAllString(body, "$1"+string(b))
		return fixed
	}
	return body
}
