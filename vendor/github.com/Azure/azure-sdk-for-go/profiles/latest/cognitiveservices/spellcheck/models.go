// +build go1.9

// Copyright 2018 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package spellcheck

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/spellcheck"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}

type ActionType = original.ActionType

const (
	Edit ActionType = original.Edit
	Load ActionType = original.Load
)

func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}

type ErrorType = original.ErrorType

const (
	RepeatedToken ErrorType = original.RepeatedToken
	UnknownToken  ErrorType = original.UnknownToken
)

func PossibleErrorTypeValues() []ErrorType {
	return original.PossibleErrorTypeValues()
}

type Type = original.Type

const (
	TypeAnswer        Type = original.TypeAnswer
	TypeErrorResponse Type = original.TypeErrorResponse
	TypeIdentifiable  Type = original.TypeIdentifiable
	TypeResponse      Type = original.TypeResponse
	TypeResponseBase  Type = original.TypeResponseBase
	TypeSpellCheck    Type = original.TypeSpellCheck
)

func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}

type BasicAnswer = original.BasicAnswer
type Answer = original.Answer
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type BasicIdentifiable = original.BasicIdentifiable
type Identifiable = original.Identifiable
type BasicResponse = original.BasicResponse
type Response = original.Response
type BasicResponseBase = original.BasicResponseBase
type ResponseBase = original.ResponseBase
type SpellCheck = original.SpellCheck
type SpellingFlaggedToken = original.SpellingFlaggedToken
type SpellingTokenSuggestion = original.SpellingTokenSuggestion

func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
