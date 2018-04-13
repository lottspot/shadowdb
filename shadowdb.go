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
  i := db.findRecord(r.Uname())
  if i >= 0 {
    db.records[i] = r
  } else {
    db.records = append(db.records, r)
  }
}

func (db *shadowDB) PurgeUser(uname string) {
  i := db.findRecord(uname)
  if i >= 0 {
    db.purgeRecordAt(uint(i))
  }
}

func (db *shadowDB) findRecord(uname string) int {
  for i, record := range db.records {
    if record.Uname() == uname {
      return i
    }
  }
  return -1
}

func (db *shadowDB) purgeRecordAt(i uint) {
  records := make([]DBRecord, len(db.records)-1)
  copy(records, db.records[:i])
  copy(records[i:], db.records[i+1:])
  db.records = records
}

func (db *shadowDB) insertRecordAt(u DBRecord, i uint) {
  db.records = append(db.records, &shadowUser{})
  copy(db.records[i+1:], db.records[i:])
  db.records[i] = u
}
