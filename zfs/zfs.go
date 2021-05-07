package zfs

import (
  "bytes"
  "os/exec"
)


func GetProp(dataset string, prop string, source string) ([]byte, error) {
  cmd := exec.Command("zfs", "get", "-H", "-p", "-o", "value", prop, "-s", source, dataset)

  out, err := cmd.Output()

  out = bytes.TrimRight(out, "\n")

  return out, err
}

func LoadKey(dataset string, noop bool, key []byte) ([]byte, error) {
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
