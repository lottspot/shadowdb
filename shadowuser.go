package shadowdb

import (
  "strings"
  "strconv"
)

type shadowUser struct {
  uname string
  pwhash string
  lastChange string
  minAge string
  maxAge string
  warnDays string
  graceDays string
  expires string
}

func NewUser() *shadowUser {
  return new(shadowUser)
}

func UserFromRecord(record string) *shadowUser {
  recordSlice := strings.Split(record, ":")
  return &shadowUser{
    uname: recordSlice[0],
    pwhash: recordSlice[1],
    lastChange: recordSlice[2],
    minAge: recordSlice[3],
    maxAge: recordSlice[4],
    warnDays: recordSlice[5],
    graceDays: recordSlice[6],
    expires: recordSlice[7],
  }
}

func (u *shadowUser) AsRecord() string {
  recordSlice := []string{
    u.uname,
    u.pwhash,
    u.lastChange,
    u.minAge,
    u.maxAge,
    u.warnDays,
    u.graceDays,
    u.expires,
    "", // ensure record has terminating ":" char
  }
  return strings.Join(recordSlice, ":")
}

func (u *shadowUser) Uname() string {
  return u.uname
}

func (u *shadowUser) Pwhash() string {
  return u.pwhash
}

func (u *shadowUser) LastChange() int {
  lastChange, e := strconv.ParseInt(u.lastChange, 10, 0)
  if e != nil { // This covers the case of an empty string
    return -1
  } else {
    return int(lastChange)
  }
}

func (u *shadowUser) MinAge() int {
  minAge, e := strconv.ParseInt(u.minAge, 10, 0)
  if e != nil { // This covers the case of an empty string
    return -1
  } else {
    return int(minAge)
  }
}

func (u *shadowUser) MaxAge() int {
  maxAge, e := strconv.ParseInt(u.maxAge, 10, 0)
  if e != nil { // This covers the case of an empty string
    return -1
  } else {
    return int(maxAge)
  }
}

func (u *shadowUser) WarnDays() int {
  warnDays, e := strconv.ParseInt(u.warnDays, 10, 0)
  if e != nil { // This covers the case of an empty string
    return -1
  } else {
    return int(warnDays)
  }
}

func (u *shadowUser) GraceDays() int {
  graceDays, e := strconv.ParseInt(u.graceDays, 10, 0)
  if e != nil { // This covers the case of an empty string
    return -1
  } else {
    return int(graceDays)
  }
}

func (u *shadowUser) Expires() int {
  expires, e := strconv.ParseInt(u.expires, 10, 0)
  if e != nil { // This covers the case of an empty string
    return -1
  } else {
    return int(expires)
  }
}

func (u *shadowUser) SetUname(uname string) {
  u.uname = uname
}

func (u *shadowUser) SetPwhash(pwhash string) {
  u.pwhash = pwhash
}

func (u *shadowUser) SetLastChange(lastChange int) {
  if lastChange < 0 {
    u.lastChange = ""
  } else {
    u.lastChange = strconv.FormatInt(int64(lastChange), 10)
  }
}

func (u *shadowUser) SetMinAge(minAge int) {
  if minAge < 0 {
    u.minAge = ""
  } else {
    u.minAge = strconv.FormatInt(int64(minAge), 10)
  }
}

func (u *shadowUser) SetMaxAge(maxAge int) {
  if maxAge < 0 {
    u.maxAge = ""
  } else {
    u.maxAge = strconv.FormatInt(int64(maxAge), 10)
  }
}

func (u *shadowUser) SetWarnDays(warnDays int) {
  if warnDays < 0 {
    u.warnDays = ""
  } else {
    u.warnDays = strconv.FormatInt(int64(warnDays), 10)
  }
}

func (u *shadowUser) SetGraceDays(graceDays int) {
  if graceDays < 0 {
    u.graceDays = ""
  } else {
    u.graceDays = strconv.FormatInt(int64(graceDays), 10)
  }
}

func (u *shadowUser) SetExpires(expires int) {
  if expires < 0 {
    u.expires = ""
  } else {
    u.expires = strconv.FormatInt(int64(expires), 10)
  }
}
