package atsam

import (
	"errors"
	"io"
	"strings"

	"periph.io/x/periph"
	"periph.io/x/periph/host/distro"
	"periph.io/x/periph/host/fs"
)

// Present returns true if a TM AM335x processor is detected.
func Present() bool {
	if isArm {
		return strings.HasPrefix(distro.DTModel(), "Atmel SAMA5D27")
	}
	return false
}

// driver implements periph.Driver.
type driver struct {
}

func (d *driver) String() string {
	return "amsam"
}

func (d *driver) Prerequisites() []string {
	return nil
}

func (d *driver) After() []string {
	return nil
}

func (d *driver) Init() (bool, error) {
	if !Present() {
		return false, errors.New("atsam CPU not detected")
	}

	return true, nil
}

func init() {
	if isArm {
		periph.MustRegister(&drv)
	}
}

var drv driver

var ioctlOpen = ioctlOpenDefault

func ioctlOpenDefault(path string, flag int) (ioctlCloser, error) {
	f, err := fs.Open(path, flag)
	if err != nil {
		return nil, err
	}
	return f, nil
}

var fileIOOpen = fileIOOpenDefault

func fileIOOpenDefault(path string, flag int) (fileIO, error) {
	f, err := fs.Open(path, flag)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type ioctlCloser interface {
	io.Closer
	fs.Ioctler
}

type fileIO interface {
	Fd() uintptr
	fs.Ioctler
	io.Closer
	io.Reader
	io.Seeker
	io.Writer
}

// seekRead seeks to the beginning of a file and reads it.
func seekRead(f fileIO, b []byte) (int, error) {
	if _, err := f.Seek(0, 0); err != nil {
		return 0, err
	}
	return f.Read(b)
}

// seekWrite seeks to the beginning of a file and writes to it.
func seekWrite(f fileIO, b []byte) error {
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	_, err := f.Write(b)
	return err
}
