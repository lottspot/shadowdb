package shadowdb

import (
  "io"
)

type DBRecord interface {
  Record() string
  Uname() string
}

type shadowDB struct {
  records []DBRecord
  recordParser interface{}
}

func NewDB() *shadowDB {
  db := new(shadowDB)
  db.records = make([]DBRecord, 0)
  db.recordParser = NewUserFromRecord
  return db
}

func (db *shadowDB) Load(r io.Reader) error {
  /*
   * parseFunc = reflect.ValueOf(db.recordParser)
   * recordVal = reflect.ValueOf(record)
   * parsedVal = parseFunc.Call([]reflect.Value{recordVal})
   * parsedRec = parseVal[0].Interface().(DBRecord)
   */
  return nil
}

func (db *shadowDB) Dump(w io.Writer) error {
  return nil
}

func (db *shadowDB) User(uname string) shadowUser {
  return shadowUser{}
}

func (db *shadowDB) ApplyRecord(r DBRecord) {
}

func (db *shadowDB) PurgeUser(uname string) {
}

func (db *shadowDB) findRecord(uname string) uint {
  return uint(0)
}

func (db *shadowDB) purgeRecordAt(i uint) {
}

func (db *shadowDB) insertRecordAt(u DBRecord, i uint) {
}
