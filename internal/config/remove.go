/*
 * Flow CLI
 *
 * Copyright Flow Foundation
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

package config

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:              "remove <account|contract|deployment|network>",
	Short:            "Remove resource from configuration",
	Example:          "flow config remove account",
	Args:             cobra.ExactArgs(1),
	TraverseChildren: true,
}

func init() {
	removeAccountCommand.AddToParent(removeCmd)
	removeContractCommand.AddToParent(removeCmd)
	removeDeploymentCommand.AddToParent(removeCmd)
	removeNetworkCommand.AddToParent(removeCmd)
}
