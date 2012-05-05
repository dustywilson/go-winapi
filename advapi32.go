// Copyright 2010 The go-winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unsafe"
)

const KEY_READ REGSAM = 0x20019

const (
	HKEY_CLASSES_ROOT     HKEY = 0x80000000
	HKEY_CURRENT_USER     HKEY = 0x80000001
	HKEY_LOCAL_MACHINE    HKEY = 0x80000002
	HKEY_USERS            HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA HKEY = 0x80000004
	HKEY_CURRENT_CONFIG   HKEY = 0x80000005
	HKEY_DYN_DATA         HKEY = 0x80000006
)

const (
	TYPE_REG_NONE                       REGTYPE = 0
	TYPE_REG_SZ                         REGTYPE = 1
	TYPE_REG_EXPAND_SZ                  REGTYPE = 2
	TYPE_REG_BINARY                     REGTYPE = 3
	TYPE_REG_DWORD                      REGTYPE = 4
	TYPE_REG_DWORD_LITTLE_ENDIAN        REGTYPE = 4
	TYPE_REG_DWORD_BIG_ENDIAN           REGTYPE = 5
	TYPE_REG_LINK                       REGTYPE = 6
	TYPE_REG_MULTI_SZ                   REGTYPE = 7
	TYPE_REG_RESOURCE_LIST              REGTYPE = 8
	TYPE_REG_FULL_RESOURCE_DESCRIPTOR   REGTYPE = 9
	TYPE_REG_RESOURCE_REQUIREMENTS_LIST REGTYPE = 10
	TYPE_REG_QWORD                      REGTYPE = 11
	TYPE_REG_QWORD_LITTLE_ENDIAN        REGTYPE = 11
)

type (
	ACCESS_MASK uint32
	HKEY        HANDLE
	REGTYPE     uint32
	REGSAM      ACCESS_MASK
)

var (
	// Library
	libadvapi32 uintptr

	// Functions
	regCloseKey     uintptr
	regOpenKeyEx    uintptr
	regQueryValueEx uintptr
	regSetValueEx   uintptr
)

func init() {
	// Library
	libadvapi32 = MustLoadLibrary("advapi32.dll")

	// Functions
	regCloseKey = MustGetProcAddress(libadvapi32, "RegCloseKey")
	regOpenKeyEx = MustGetProcAddress(libadvapi32, "RegOpenKeyExW")
	regQueryValueEx = MustGetProcAddress(libadvapi32, "RegQueryValueExW")
	regSetValueEx = MustGetProcAddress(libadvapi32, "RegSetValueExW")
}

func RegCloseKey(hKey HKEY) int32 {
	ret, _, _ := syscall.Syscall(regCloseKey, 1,
		uintptr(hKey),
		0,
		0)

	return int32(ret)
}

func RegOpenKeyEx(hKey HKEY, lpSubKey *uint16, ulOptions uint32, samDesired REGSAM, phkResult *HKEY) int32 {
	ret, _, _ := syscall.Syscall6(regOpenKeyEx, 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpSubKey)),
		uintptr(ulOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		0)

	return int32(ret)
}

func RegQueryValueEx(hKey HKEY, lpValueName *uint16, lpReserved, lpType *uint32, lpData *byte, lpcbData *uint32) int32 {
	ret, _, _ := syscall.Syscall6(regQueryValueEx, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)))

	return int32(ret)
}

func RegSetValueEx(hKey HKEY, lpValueName *uint16, lpReserved, dwType *REGTYPE, lpData *byte, cbData *uint32) int32 {
	ret, _, _ := syscall.Syscall6(regSetValueEx, 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(dwType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(cbData)))

	return int32(ret)
}
