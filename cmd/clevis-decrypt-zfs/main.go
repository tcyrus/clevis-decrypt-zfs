package main

import (
  "log"
  "os"

  clevis "github.com/tcyrus/clevis-decrypt-zfs/clevis"
  zfs "github.com/tcyrus/clevis-decrypt-zfs/zfs"
)

func main() {
  if len(os.Args) < 2 {
    log.Fatalf("Usage: %s DATASET", os.Args[0])
  }

  var dataset string
  dataset = os.Args[1]

  decryptProp, err := zfs.GetProp(dataset, "latchset.clevis:decrypt", "local")
  if err != nil {
    log.Fatal(err)
  }

  if (string(decryptProp) != "yes") {
    log.Fatalf("%s dataset does not support clevis-like decryption: %q", dataset, decryptProp)
  }

  keystat, err := zfs.GetProp(dataset, "keystatus", "none")
  if err != nil {
    log.Fatal(err)
  }

  load_noop := (string(keystat) != "unavailable")

  jwe, err := zfs.GetProp(dataset, "latchset.clevis:jwe", "local")
  if err != nil {
    log.Fatal(err)
  }

  key, err := clevis.Decrypt(jwe)
  if err != nil {
    log.Fatal(err)
  }

  res, err := zfs.LoadKey(dataset, load_noop, key)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("%s\n", res)
}
