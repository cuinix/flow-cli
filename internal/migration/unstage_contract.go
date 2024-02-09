/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package migration

import (
	"fmt"

	"github.com/onflow/flowkit"
	"github.com/onflow/flowkit/output"
	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/internal/transactions"
)

var unstageContractflags = transactions.Flags{}

var unstageContractCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "flow unstage-contract <NAME> --network <NETWORK> --signer <HOST_ACCOUNT>",
		Short:   "unstage a contract for migration",
		Example: `flow unstage-contract HelloWorld`,
		Args:    cobra.MinimumNArgs(1),
	},
	Flags: &unstageContractflags,
	RunS:  unstageContract,
}

func unstageContract(
	args []string,
	globalFlags command.GlobalFlags,
	_ output.Logger,
	flow flowkit.Services,
	state *flowkit.State,
) (command.Result, error) {
	code, err := RenderContractTemplate(UnstageContractTransactionFilepath, globalFlags.Network)
	if err != nil {
		return nil, fmt.Errorf("error loading staging contract file: %w", err)
	}

	contractName := args[0]

	res, err := transactions.SendTransaction(
		code,
		[]string{
			contractName,
		},
		"",
		flow,
		state,
		stageContractflags,
	)
	if err != nil {
		return nil, fmt.Errorf("error sending transaction: %w", err)
	}

	return res, nil
}
