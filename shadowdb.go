package shadowdb

import (
  "io"
)

type ShadowUser interface {
  AsRecord() string
  Uname() string
  Pwhash() string
  LastChange() int
  MinAge() int
  MaxAge() int
  WarnDays() int
  GraceDays() int
  Expires() int
  SetUname(uname string)
  SetPwhash(pwhash string)
  SetMinAge(minAge int)
  SetMaxAge(maxAge int)
  SetWarnDays(warnDays int)
  SetGraceDays(graceDays int)
  SetExpires(expires int)
}

type shadowDB struct {
  records []ShadowUser
}

func NewDB() *shadowDB {
  db := new(shadowDB)
  db.records = make([]ShadowUser, 0)
  return db
}

func (db *shadowDB) Load(r io.Reader) error {
  return nil
}

func (db *shadowDB) Dump(w io.Writer) error {
  return nil
}

func (db *shadowDB) User(uname string) shadowUser {
  return shadowUser{}
}

func (db *shadowDB) ApplyUser(u ShadowUser) {
}

func (db *shadowDB) PurgeUser(uname string) {
}

func (db *shadowDB) findRecord(uname string) uint {
  return uint(0)
}

func (db *shadowDB) purgeRecordAt(i uint) {
}

func (db *shadowDB) insertRecordAt(u ShadowUser, i uint) {
}
