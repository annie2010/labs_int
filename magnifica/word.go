// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package magnifica

import (
	"time"

	"encoding/json"
	"path/filepath"
)

// Entry represents a dictionary entry
type Entry struct {
	Dictionary string
	Location   string
	Word       string
	Slang      bool
	origin     string
}

func (w Entry) toACMEEntry() acmeEntry {
	<<!!YOUR_CODE!!>> -- implement an ACME word converter
}

type acmeEntry struct {
	<<!!YOUR_CODE!!>> -- implement an ACME entry
}

func (v acmeEntry) toEntry() Entry {
	<<!!YOUR_CODE!!>> -- implement an ACME to Magnifica Entry
}

// MarshalJSON converts entry into raw json
func (w *Entry) MarshalJSON() ([]byte, error) {
	<<!!YOUR_CODE!!>>
}

// UnmarshalJSON hydrates an entry from raw json
func (w *Entry) UnmarshalJSON(bb []byte) error {
	<<!!YOUR_CODE!!>>
}
