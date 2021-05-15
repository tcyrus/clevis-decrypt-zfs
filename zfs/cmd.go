package zfs

import (
  "bytes"
  "os/exec"
)

type Dataset struct {
  Path   string
}

type Property struct {
  Value  string
  Source string
}

func DatasetOpen(dataset string) (Dataset, error) {
  return Dataset{Path: dataset}, nil
}

func (d *Dataset) Close() {
  return
}

func (d *Dataset) GetKeyStatus() (p Property, err error) {
  res, err := GetPropCmd(d.Path, "keystatus", "none")
  if (err != nil) {
    return
  }
  p = Property{Value: string(res), Source: "none"}
  return
}

func (d *Dataset) GetUserProperty(pName string) (p Property, err error) {
  res, err := GetPropCmd(d.Path, pName, "local")
  if (err != nil) {
    return
  }
  p = Property{Value: string(res), Source: "local"}
  return
}

func (d *Dataset) LoadKey(noop bool, key []byte) ([]byte, error) {
  return LoadKeyCmd(d.Path, noop, key)
}

func GetPropCmd(dataset string, prop string, source string) ([]byte, error) {
  cmd := exec.Command("zfs", "get", "-H", "-p", "-o", "value", prop, "-s", source, dataset)

  out, err := cmd.Output()

  out = bytes.TrimRight(out, "\n")

  return out, err
}

func LoadKeyCmd(dataset string, noop bool, key []byte) ([]byte, error) {
  var args []string

  if (noop) {
    args = []string{"load-key", "-n", "-L", "prompt", dataset}
  } else {
    args = []string{"load-key", "-L", "prompt", dataset}
  }

  cmd := exec.Command("zfs", args...)

  stdin, err := cmd.StdinPipe()
  if err != nil {
    return nil, err
  }

  go func() {
    defer stdin.Close()
    stdin.Write(key)
  }()

  out, err := cmd.CombinedOutput()

  out = bytes.TrimRight(out, "\n")

  return out, err
}
