package shadowdb

import (
  "testing"
)

type stubShadowUser struct {
  asRecord func() string
  uname func() string
  pwhash func() string
  lastChange func() int
  minAge func() int
  maxAge func() int
  warnDays func() int
  graceDays func() int
  expires func () int
}
func (s stubShadowUser) AsRecord() string {if s.asRecord == nil {return ""} else {return s.asRecord()}}
func (s stubShadowUser) Uname() string {if s.uname == nil {return ""} else {return s.uname()}}
func (s stubShadowUser) Pwhash() string {if s.pwhash == nil {return ""} else {return s.pwhash()}}
func (s stubShadowUser) LastChange() int {if s.lastChange == nil {return -1} else {return s.lastChange()}}
func (s stubShadowUser) MinAge() int {if s.minAge == nil {return -1} else {return s.minAge()}}
func (s stubShadowUser) MaxAge() int {if s.maxAge == nil {return -1} else {return s.maxAge()}}
func (s stubShadowUser) WarnDays() int {if s.warnDays == nil {return -1} else {return s.warnDays()}}
func (s stubShadowUser) GraceDays() int {if s.graceDays == nil {return -1} else {return s.graceDays()}}
func (s stubShadowUser) Expires() int {if s.expires == nil {return -1} else {return s.expires()}}
func (s stubShadowUser) SetUname(uname string) {}
func (s stubShadowUser) SetPwhash(pwhash string) {}
func (s stubShadowUser) SetMinAge(minAge int) {}
func (s stubShadowUser) SetMaxAge(maxAge int) {}
func (s stubShadowUser) SetWarnDays(warnDays int) {}
func (s stubShadowUser) SetGraceDays(graceDays int) {}
func (s stubShadowUser) SetExpires(expires int) {}

func TestInsertRecordAt(t *testing.T) {}
