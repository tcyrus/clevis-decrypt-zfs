package clevis

import (
  "fmt"

  "github.com/lestrrat-go/jwx/jwe"
)

// Decrypt decrypts a clevis bound message. The message format can be either compact or JSON.
func Decrypt(data []byte) ([]byte, error) {
  msg, err := jwe.Parse(data)
  if err != nil {
    return nil, err
  }

  clevis, err := GetClevisNode(msg)
  if err != nil {
    return nil, err
  }

  pin, err := GetClevisPin(clevis)
  if err != nil {
    return nil, err
  }

/*
  pin, ok := clevis["pin"].(string)
  if !ok {
    return nil, fmt.Errorf("clevis.go: provided message does not contain 'clevis.pin' node")
  }
*/

  switch pin {
    case "tang":
      return DecryptTang(msg, clevis)
    case "sss":
      return DecryptSss(msg, clevis)
    default:
      return nil, fmt.Errorf("clevis.go: unknown pin '%v'", pin)
  }
}

func GetClevisNode(msg *jwe.Message) (map[string]interface{}, error) {
  node, ok := msg.Recipients()[0].Headers().PrivateParams()["clevis"].(map[string]interface{})
  if !ok {
    return nil, fmt.Errorf("clevis.go: provided message does not contain 'clevis' node")
  }
  return node, nil
}

func GetClevisPin(node map[string]interface{}) (string, error) {
  pin, ok := node["pin"].(string)
  if !ok {
    return "", fmt.Errorf("clevis.go: provided message does not contain 'clevis.pin' node")
  }
  return pin, nil
}
