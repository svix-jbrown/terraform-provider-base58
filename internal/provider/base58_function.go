// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/akamensky/base58"
)

var (
	_ function.Function = Base58Function{}
)

func NewBase58Function() function.Function {
	return Base58Function{}
}

type Base58Function struct{}

func (r Base58Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base58"
}

func (r Base58Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Encode a string using base58",
		MarkdownDescription: "Encode a string using base58",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "String to encode",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r Base58Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	var encoded string = base58.Encode([]byte(data))

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, encoded))
}
