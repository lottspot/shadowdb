package shadowdb

import (
  "io"
)

type shadowDB struct {
  records []shadowUser
}

func NewDB() *shadowDB {
  db := new(shadowDB)
  db.records = make([]shadowUser, 0)
  return db
}

func (db *shadowDB) Load(src io.Reader) error {
  return nil
}

func (db *shadowDB) Dump(dst io.Writer) error {
  return nil
}

func (db *shadowDB) User(uname string) shadowUser {
  return shadowUser{}
}

func (db *shadowDB) ApplyUser(u *shadowUser) {
}

func (db *shadowDB) PurgeUser(uname string) {
}

func (db *shadowDB) findRecord(uname string) uint {
  return uint(0)
}

func (db *shadowDB) purgeRecordAt(i uint) {
}

func (db *shadowDB) insertRecordAt(u *shadowUser, i uint) {
}
