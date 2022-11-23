// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// This build tag is equivalent of "unix" in Go 1.19.
//go:build aix || android || darwin || dragonfly || freebsd || hurd || illumos || ios || linux || netbsd || openbsd || solaris
// +build aix android darwin dragonfly freebsd hurd illumos ios linux netbsd openbsd solaris

package file

import (
	"os"

	"golang.org/x/sys/unix"
)

// RedirectStandardError causes all standard error output to be directed to the
// given file.
func RedirectStandardError(toFile *os.File) error {
	return unix.Dup2(int(toFile.Fd()), 2)
}
