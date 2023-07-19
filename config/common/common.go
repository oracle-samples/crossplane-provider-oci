/*
 * Copyright (c) 2022, 2023 Oracle and/or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"fmt"
	"strings"
)

const (
	SelfPackagePath = "github.com/oracle/provider-oci/config/common"

	VersionAlpha1 = "v1alpha1"

	ErrFmtNoAttribute = "the attribute %s is not found"

	ErrFmtUnexpectedType = "the attribute %s is not in the expected format"
)

/*
GetNameFromFullyQualifiedID
Given the Terraform ID in the pattern like <resource1>/<resource2>/.../<resourceN>
This function extracts the last part of the ID and returns <resourceN>
*/
func GetNameFromFullyQualifiedID(tfState map[string]any) (string, error) {
	id, ok := tfState["id"]
	if !ok {
		return "", fmt.Errorf(ErrFmtNoAttribute, "id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", fmt.Errorf(ErrFmtUnexpectedType, "id")
	}
	words := strings.Split(idStr, "/")
	return words[len(words)-1], nil
}

func GetIDFromParamsFunc(parameters map[string]any, attrsMap map[string]string) (string, error) {
	var result string
	for key, value := range attrsMap {
		paramValue, ok := parameters[value]
		if !ok {
			return "", fmt.Errorf(ErrFmtNoAttribute, value)
		}
		paramValueStr, ok := paramValue.(string)
		if !ok {
			return "", fmt.Errorf(ErrFmtUnexpectedType, value)
		}
		result = fmt.Sprintf("%s/%s", key, paramValueStr)
	}
	return result, nil
}
