package shadowdb

import (
  "testing"
  "bytes"
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
  for _, record := range db.records {
    if record.Uname() == "purgeme" {
      t.Error("user was not purged")
    }
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

func TestPurgeUser(t *testing.T){
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
  db.PurgeUser("purgeme")
  for _, record := range db.records {
    if record.Uname() == "purgeme" {
      t.Error("user was not purged")
    }
  }
}

func TestApplyRecord(t *testing.T){
  user1Old := stubDBRecord{
    uname: func() (string) {return "user1"},
    record: func() (string) {return "old"},
  }
  user1New := stubDBRecord{
    uname: func() (string) {return "user1"},
    record: func() (string) {return "new"},
  }
  user2 := stubDBRecord{
    uname: func() (string) {return "user2"},
  }
  records := []DBRecord{
    user1Old,
  }
  db := shadowDB{
    records: records,
  }
  db.ApplyRecord(user1New)
  db.ApplyRecord(user2)
  if db.records[0].Record() != "new" {
    t.Error("user1 record was not updated")
  }
  if db.records[len(db.records)-1].Uname() != "user2" {
    t.Error("user2 was not added to db")
  }
}

func TestUserGetter(t *testing.T){
  toGet := shadowUser{
    uname: "getme",
  }
  records := []DBRecord{
    &shadowUser{},
    &shadowUser{},
    &toGet,
  }
  db := shadowDB{
    records: records,
  }
  got := db.User("getme")
  if got.uname != toGet.uname {
    t.Error("expected to get user", toGet.uname, ", got", got.uname)
  }
}

func TestDump(t *testing.T){
  records := []DBRecord{
    stubDBRecord{record: func()(string){return"user1"}},
    stubDBRecord{record: func()(string){return"user2"}},
    stubDBRecord{record: func()(string){return"user3"}},
  }
  file := bytes.NewBuffer([]byte{})
  db := shadowDB{
    records: records,
  }
  db.Dump(file)
  result := file.String()
  if result != "user1\nuser2\nuser3\n" {
    t.Error("\nexpected:\nuser1\nuser2\nuser3\n\ngot:\n", result)
  }
}

func TestLoad(t *testing.T){
  fileContents :=
`user1
user2
user3
`
  file := bytes.NewBuffer([]byte(fileContents))
  db := shadowDB{
    recordParser: func (record string) stubDBRecord {return stubDBRecord{}},
  }
  db.Load(file)
  nLoaded := len(db.records)
  if nLoaded != 3 {
    t.Error("expected 3 records loaded, got", nLoaded)
  }
}