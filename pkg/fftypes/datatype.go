// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fftypes

import (
	"context"

	"github.com/kaleido-io/firefly/internal/i18n"
)

type ValidatorType string

const (
	// ValidatorTypeJSON is the validator type for JSON Schema validation
	ValidatorTypeJSON ValidatorType = "json"
	// ValidatorTypeBLOB is the validator type for binary blob data, that is passed through without any parsing or validation
	ValidatorTypeBLOB ValidatorType = "blob"
	// ValidatorTypeDatatype is the validator type for data definitions, which are a built in type that defines other types
	ValidatorTypeDatatype ValidatorType = "datadef"
)

// Datatype is the structure defining a data definition, such as a JSON schema
type Datatype struct {
	ID        *UUID         `json:"id,omitempty"`
	Validator ValidatorType `json:"validator"`
	Namespace string        `json:"namespace,omitempty"`
	Name      string        `json:"name,omitempty"`
	Version   string        `json:"version,omitempty"`
	Hash      *Bytes32      `json:"hash,omitempty"`
	Created   *FFTime       `json:"created,omitempty"`
	Value     Byteable      `json:"value,omitempty"`
}

func (dt *Datatype) Validate(ctx context.Context, existing bool) (err error) {
	if dt.Validator != ValidatorTypeJSON {
		return i18n.NewError(ctx, i18n.MsgUnknownFieldValue, "validator", dt.Validator)
	}
	if err = ValidateFFNameField(ctx, dt.Namespace, "namespace"); err != nil {
		return err
	}
	if err = ValidateFFNameField(ctx, dt.Name, "name"); err != nil {
		return err
	}
	if err = ValidateFFNameField(ctx, dt.Version, "version"); err != nil {
		return err
	}
	if len(dt.Value) == 0 {
		return i18n.NewError(ctx, i18n.MsgMissingRequiredField, "value")
	}
	if existing {
		if dt.ID == nil {
			return i18n.NewError(ctx, i18n.MsgNilID)
		}
		hash := dt.Value.Hash()
		if dt.Hash == nil || *dt.Hash != *hash {
			return i18n.NewError(ctx, i18n.MsgDataInvalidHash, hash, dt.Hash)
		}
	}
	return nil
}
