package shadowdb

import (
  "testing"
)

type stubDBRecord struct {
  record func() string
  uname func() string
}
func (s stubDBRecord) Record() string {if s.record == nil {return ""} else {return s.record()}}
func (s stubDBRecord) Uname() string {if s.uname == nil {return ""} else {return s.uname()}}

func TestInsertRecordAt(t *testing.T) {}
