// +build libzfs

package zfs

import (
//  "os"
//  "syscall"

  zfs "github.com/bicomsystems/go-libzfs"
)

type Dataset = zfs.Dataset

func  (d *Dataset) GetKeyStatus() (string, error) {
  var p zfs.Property
  if p, err = d.GetProperty(zfs.DatasetPropKeyStatus); err != nil {
    return nil, err
  }

  return p.Value, nil
}

/*
// There's no easy way to pass raw byte keys to libzfs
// because of the way their crypto api is structured.
// That being said, I have some scratch code here that
// might work, but probably isn't worth the hassle or
// risk (it would need to be added to the zfs wrapper
// library that we're using)
//
// This bad idea was based on:
// https://play.golang.org/p/Xg2iajdiuNN
// https://github.com/cbreak-black/ZetaWatch/blob/r47/ZFSWrapper/ZFSUtils.cpp#L426-L503
// https://github.com/bicomsystems/go-libzfs/blob/v0.3.5/zfs.go#L388-L412
func (d *Dataset) LoadKey(k []byte, noop bool) (err error) {
  Global.Mtx.Lock()
  defer Global.Mtx.Unlock()

  if d.list == nil {
    err = errors.New(msgDatasetIsNil)
    return
  }

  r, w, err := os.Pipe()
  if err != nil {
    return
  }

  restore := func(r, fd) {
    syscall.Dup2(fd, syscall.Stdin)
    syscall.Close(fd)
    r.Close()
  }

  oldStdin, err := syscall.Dup(syscall.Stdin)
  defer restore(r, oldStdin)
  if err != nil {
    return
  }

  err = syscall.Dup2(int(r.Fd()), syscall.Stdin)
  if err != nil {
    return
  }

  future := make(chan error)

  go func() {
    defer close(future)
    _, ferr := w.Write(k)
    future <- ferr
  }()

  prompt := C.CString("prompt")
  defer C.free(unsafe.Pointer(prompt))

  errcode := C.zfs_crypto_load_key(d.list.zh, booleanT(noop), prompt)

  err = <- future
  if err != nil {
    return
  }

  if errcode != 0 {
    err = LastError()
    return
  }
}
*/
