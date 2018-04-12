package shadowdb

import (
  "testing"
)

func TestNewUserFromRecord(t *testing.T) {
  record := "foo:bar:1::120:7:7::"
  u, e := NewUserFromRecord(record)
  if e != nil {
    t.Fatal("Error initializing user:", e.Error())
  }
  if u.uname != "foo" {
    t.Error("uname: expected \"foo\", got", u.uname)
  }
  if u.pwhash != "bar" {
    t.Error("pwhash: expected \"bar\", got", u.pwhash)
  }
  if u.lastChange != "1" {
    t.Error("lastChange: expected \"1\", got", u.lastChange)
  }
  if u.minAge != "" {
    t.Error("minAge: expected \"\", got", u.minAge)
  }
  if u.maxAge != "120" {
    t.Error("maxAge: expected \"120\", got", u.maxAge)
  }
  if u.warnDays != "7" {
    t.Error("warnDays: expected \"7\", got", u.warnDays)
  }
  if u.graceDays != "7" {
    t.Error("graceDays: expected \"7\", got", u.graceDays)
  }
  if u.expires != "" {
    t.Error("expires: expected \"\", got", u.expires)
  }
}

func TestAsRecord(t *testing.T) {
  u := shadowUser{
    uname: "foo",
    pwhash: "bar",
    lastChange: "1",
    minAge: "",
    maxAge: "120",
    warnDays: "7",
    graceDays: "7",
    expires: "",
  }
  record := u.AsRecord()
  expect := "foo:bar:1::120:7:7::"
  if record != expect {
    t.Error("\nexpected:", expect, "\nreceived:", record)
  }
}

func TestUname(t *testing.T) {
  u := shadowUser{}
  u.SetUname("foo")
  result := u.Uname()
  if result != "foo" {
    t.Error("Unexpected value for uname:", result)
  }
}

func TestPwhash(t *testing.T) {
  u := shadowUser{}
  u.SetPwhash("foo")
  result := u.Pwhash()
  if result != "foo" {
    t.Error("Unexpected value for pwhash:", result)
  }
}

func TestLastChange(t *testing.T) {
  u := shadowUser{}
  u.SetLastChange(1)
  result := u.LastChange()
  internal := u.lastChange
  if result != 1 {
    t.Error("Unexpected value for lastChange:", result)
  }
  if internal != "1" {
    t.Error("Unexpected internal storage value for lastChange:", internal)
  }
}

func TestMinAge(t *testing.T) {
  u := shadowUser{}
  u.SetMinAge(1)
  result := u.MinAge()
  internal := u.minAge
  if result != 1 {
    t.Error("Unexpected value for minAge:", result)
  }
  if internal != "1" {
    t.Error("Unexpected internal storage value for minAge:", internal)
  }
}

func TestMaxAge(t *testing.T) {
  u := shadowUser{}
  u.SetMaxAge(1)
  result := u.MaxAge()
  internal := u.maxAge
  if result != 1 {
    t.Error("Unexpected value for maxAge:", result)
  }
  if internal != "1" {
    t.Error("Unexpected internal storage value for maxAge:", internal)
  }
}

func TestWarnDays(t *testing.T) {
  u := shadowUser{}
  u.SetWarnDays(1)
  result := u.WarnDays()
  internal := u.warnDays
  if result != 1 {
    t.Error("Unexpected value for warnDays:", result)
  }
  if internal != "1" {
    t.Error("Unexpected internal storage value for warnDays:", internal)
  }
}

func TestGraceDays(t *testing.T) {
  u := shadowUser{}
  u.SetGraceDays(1)
  result := u.GraceDays()
  internal := u.graceDays
  if result != 1 {
    t.Error("Unexpected value for graceDays:", result)
  }
  if internal != "1" {
    t.Error("Unexpected internal storage value for graceDays:", internal)
  }
}

func TestExpires(t *testing.T) {
  u := shadowUser{}
  u.SetExpires(1)
  result := u.Expires()
  internal := u.expires
  if result != 1 {
    t.Error("Unexpected value for expires:", result)
  }
  if internal != "1" {
    t.Error("Unexpected internal storage value for expires:", internal)
  }
}
