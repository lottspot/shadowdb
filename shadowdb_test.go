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

func TestInsertRecordAt(t *testing.T) {
  records := []DBRecord{
    stubDBRecord{},
    stubDBRecord{},
    stubDBRecord{},
  }
  db := shadowDB{
    records: records,
  }
  newRecord := stubDBRecord{
    uname: func() (string) {return "user"},
  }
  db.insertRecordAt(newRecord, 1)
  if db.records[1].Uname() != "user" {
    t.Error("user did not appear at index 1")
  }
}

func TestPurgeRecordAt(t *testing.T) {
  toPurge := stubDBRecord{
    uname: func() (string) {return "purgeme"},
  }
  records := []DBRecord{
    stubDBRecord{},
    toPurge,
    stubDBRecord{},
  }
  db := shadowDB{
    records: records,
  }
  db.purgeRecordAt(1)
  if db.records[1].Uname() == "purgeme" {
    t.Error("user was not purged")
  }
}

func TestFindRecord(t *testing.T) {
  toFind := stubDBRecord{
    uname: func() (string) {return "findme"},
  }
  records := []DBRecord{
    stubDBRecord{},
    stubDBRecord{},
    toFind,
  }
  db := shadowDB{
    records: records,
  }
  i := db.findRecord("findme")
  if db.records[i].Uname() != "findme" {
    t.Error("incorrect user found")
  }
}