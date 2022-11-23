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
	"path/filepath"
)

// SafeFileRotate safely rotates an existing file under path and replaces it with the tempfile
func SafeFileRotate(path, tempfile string) error {
	if e := os.Rename(tempfile, path); e != nil {
		return e
	}

	// best-effort fsync on parent directory. The fsync is required by some
	// filesystems, so to update the parents directory metadata to actually
	// contain the new file being rotated in.
	return SyncParent(path)
}

// SyncParent fsyncs parent directory
func SyncParent(path string) error {
	flags := os.O_RDONLY
	if runtime.GOOS == "aix" {
		// On AIX, fsync will fail if the file is opened in read-only mode.
		flags = os.O_RDWR
	}

	parent := filepath.Dir(path)
	f, err := os.Open(parent, flags, 0)
	if err != nil {
		return nil // ignore error
	}
	defer f.Close()

	return f.Sync()
}
