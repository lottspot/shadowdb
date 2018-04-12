package shadowdb

import (
  "testing"
)

func TestAsRecord(t *testing.T) {
  u := shadowUser{}
  t.Log("successfully assigned ShadowUser:", u)
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
