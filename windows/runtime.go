package windows

import (
	"errors"

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

type User32 struct {
	methods []string
	dll     *win.LazyDLL
	procs   map[string]*win.LazyProc
}

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
