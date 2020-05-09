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
	return acmeEntry{
		Location: filepath.Join(w.Location, w.Dictionary),
		Word:     w.Word,
		Slang:    w.Slang,
	}
}

type acmeEntry struct {
	Location string `json:"dictionary_location"`
	Word     string `json:"dictionary_word"`
	Slang    bool   `json:"political_correctness"`
}

func (v acmeEntry) toEntry() Entry {
	time.Now().UTC().UnixNano()
	dir, dic := filepath.Split(v.Location)
	return Entry{
		Dictionary: dic,
		Location:   dir,
		Word:       v.Word,
		Slang:      v.Slang,
		origin:     "ACME",
	}
}

// Note: Compile check
var _ json.Marshaler = (*Entry)(nil)

// MarshalJSON converts entry into raw json
func (w *Entry) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.toACMEEntry())
}

// UnmarshalJSON hydrates an entry from raw json
func (w *Entry) UnmarshalJSON(bb []byte) error {
	var ae acmeEntry
	if err := json.Unmarshal(bb, &ae); err != nil {
		return err
	}
	*w = ae.toEntry()

	return nil
}
