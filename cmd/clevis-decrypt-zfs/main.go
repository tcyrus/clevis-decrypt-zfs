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

  var datasetPath string
  datasetPath = os.Args[1]

  dataset, err := zfs.DatasetOpen(datasetPath)
  if err != nil {
    log.Fatal(err)
  }
  defer dataset.Close()

  prop, err := dataset.GetUserProperty("latchset.clevis:decrypt")
  if err != nil {
    log.Fatal(err)
  }

  if (prop.Value != "yes") {
    log.Fatalf("%s dataset does not support clevis-like decryption: %s", datasetPath, prop.Value)
  }

  prop, err = dataset.GetKeyStatus()
  if err != nil {
    log.Fatal(err)
  }

  load_noop := (prop.Value != "unavailable")

  prop, err = dataset.GetUserProperty("latchset.clevis:jwe")
  if err != nil {
    log.Fatal(err)
  }

  key, err := clevis.Decrypt([]byte(prop.Value))
  if err != nil {
    log.Fatal(err)
  }

  res, err := dataset.LoadKey(load_noop, key)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("%s\n", res)
}
