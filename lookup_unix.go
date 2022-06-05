//go:build darwin || freebsd || linux || netbsd || openbsd

package goaccountcredential

import (
	"math"
	"os/user"
	"strconv"
	"syscall"
)

func castAccountID(s string) (accountID uint32, err error) {
	var v uint64
	if v, err = strconv.ParseUint(s, 10, 32); nil != err {
		return math.MaxUint32, err
	}
	accountID = uint32(v)
	return accountID, nil
}
func currentAccountID() (uid, gid uint32, err error) {
	u, err := user.Current()
	if nil != err {
		return math.MaxUint32, math.MaxUint32, err
	}
	if uid, err = castAccountID(u.Uid); nil != err {
		return math.MaxUint32, math.MaxUint32, err
	}
	if gid, err = castAccountID(u.Gid); nil != err {
		return math.MaxUint32, math.MaxUint32, err
	}
	return uid, gid, nil
}

func Lookup(userName, groupName string) (c *syscall.Credential, err error) {
	if (userName == "") && (groupName == "") {
		return
	}
	uid, gid, err := currentAccountID()
	if nil != err {
		return nil, err
	}
	if userName != "" {
		if u, err := user.Lookup(userName); nil != err {
			return nil, err
		} else if uid, err = castAccountID(u.Uid); nil != err {
			return nil, err
		} else if gid, err = castAccountID(u.Gid); nil != err {
			return nil, err
		}
	}
	if groupName != "" {
		if g, err := user.LookupGroup(groupName); nil != err {
			return nil, err
		} else if gid, err = castAccountID(g.Gid); nil != err {
			return nil, err
		}
	}
	c = &syscall.Credential{
		Uid:         uid,
		Gid:         gid,
		NoSetGroups: true,
	}
	return
}
