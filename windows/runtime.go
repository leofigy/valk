package windows

import (
	"errors"
	"syscall"
	"unsafe"

	win "golang.org/x/sys/windows"
)

const (
	DLL_NAME = "user32.dll"
)

var (
	methods = []string{
		"GetKeyState",
		"GetKeyboardLayout",
		"GetKeyboardState",
		"GetForegroundWindow",
		"ToUnicodeEx",
		"GetWindowTextW",
		"GetWindowTextLengthW",
	}
)

type (
	User32 struct {
		methods []string
		dll     *win.LazyDLL
		procs   map[string]*win.LazyProc
	}
)

func NewUser32() User32 {
	user := User32{
		methods,
		win.NewLazyDLL(DLL_NAME),
		make(map[string]*win.LazyProc, len(methods)),
	}
	user.LoadMethods()
	return user
}

func (u *User32) LoadMethods() error {

	if u.dll == nil {
		return errors.New("not working with a dll")
	}

	for _, method := range methods {
		u.procs[method] = u.dll.NewProc(method)
	}

	return nil
}

func (u *User32) GetWinTextLen(handler uintptr) (int, error) {
	ret, _, err := u.procs["GetWindowTextLengthW"].Call(
		handler,
	)
	return int(ret), err
}

func (u *User32) GetWinText(handler uintptr) (string, error) {
	lenWord, err := u.GetWinTextLen(handler)
	if err != nil {
		return "", err
	}
	lenWord += 1
	buf := make([]uint16, lenWord)
	u.procs["GetWindowTextW"].Call(
		handler,
		uintptr(
			unsafe.Pointer(&buf[0])),
		uintptr(lenWord),
	)

	return syscall.UTF16ToString(buf), nil
}
