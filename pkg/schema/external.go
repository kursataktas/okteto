//  Copyright 2024 The Okteto Authors
// Copyright 2023|2024 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
// http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package schema

import "github.com/invopop/jsonschema"

type external struct{}

func (external) JSONSchema() *jsonschema.Schema {
	endpointProps := jsonschema.NewProperties()
	endpointProps.Set("name", &jsonschema.Schema{
		Type: "string",
	})
	endpointProps.Set("url", &jsonschema.Schema{
		Type:   "string",
		Format: "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)",
	})

	externalProps := jsonschema.NewProperties()
	externalProps.Set("notes", &jsonschema.Schema{
		Type: "string",
	})
	externalProps.Set("icon", &jsonschema.Schema{
		Type: "string",
	})
	externalProps.Set("endpoints", &jsonschema.Schema{
		Type: "array",
		Items: &jsonschema.Schema{
			Type:       "object",
			Properties: endpointProps,
			Required:   []string{"url"},
		},
	})

	return &jsonschema.Schema{
		Type:                 "object",
		AdditionalProperties: jsonschema.FalseSchema,
		PatternProperties: map[string]*jsonschema.Schema{
			".*": {
				Type:                 "object",
				AdditionalProperties: jsonschema.FalseSchema,
				Properties:           externalProps,
			},
		},
	}
}
