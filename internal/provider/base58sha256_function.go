// Copyright Svix, Inc. 2025, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"crypto/sha256"

	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/akamensky/base58"
)

var (
	_ function.Function = Base58Sha256Function{}
)

func NewBase58Sha256Function() function.Function {
	return Base58Sha256Function{}
}

type Base58Sha256Function struct{}

func (r Base58Sha256Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base58sha256"
}

func (r Base58Sha256Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Encode a string using base58(sha256())",
		MarkdownDescription: "Encode a string using base58(sha256())",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "String to encode",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r Base58Sha256Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	var hashed = sha256.Sum256([]byte(data))
	var encoded = base58.Encode(hashed[:])

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, encoded))
}
