// Copyright 2015 CoreOS, Inc.
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

// Package journal provides write bindings to the local systemd journal.
// It is implemented in pure Go and connects to the journal directly over its
// unix socket.
//
// To read from the journal, see the "sdjournal" package, which wraps the
// sd-journal a C API.
//
// http://www.freedesktop.org/software/systemd/man/systemd-journald.service.html
// +build js wasm

package journal

import (
	"fmt"
)

// Priority of a journal message
type Priority int

const (
	PriEmerg Priority = iota
	PriAlert
	PriCrit
	PriErr
	PriWarning
	PriNotice
	PriInfo
	PriDebug
)

// Enabled checks whether the local systemd journal is available for logging.
func Enabled() bool {
	return true
}

// Send a message to the local systemd journal. vars is a map of journald
// fields to values.  Fields must be composed of uppercase letters, numbers,
// and underscores, but must not start with an underscore. Within these
// restrictions, any arbitrary field name may be used.  Some names have special
// significance: see the journalctl documentation
// (http://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html)
// for more details.  vars may be nil.
func Send(message string, priority Priority, vars map[string]string) error {
	return nil
}

// Print prints a message to the local systemd journal using Send().
func Print(priority Priority, format string, a ...interface{}) error {
	return Send(fmt.Sprintf(format, a...), priority, nil)
}
