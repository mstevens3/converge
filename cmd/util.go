// Copyright © 2016 Asteris, LLC
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

package cmd

import (
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// bind a set of PFlags to Viper, failing and exiting on error
func viperBindPFlags(flags *pflag.FlagSet) {
	if err := viper.BindPFlags(flags); err != nil {
		logrus.WithError(err).Fatal("could not bind flags")
	}
}

// UseColor tells us whether or not to print colors using ANSI escape sequences
// based on the following: 1. If we're in a color terminal 2. If the user has
// specified the `nocolor` option (deduced via Viper) 3. If we're on Windows.
func UseColor() bool {
	isColorTerminal := logrus.IsTerminal() && (runtime.GOOS != "windows")
	return !viper.GetBool("nocolor") && isColorTerminal
}