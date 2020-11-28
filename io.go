package core

import (
	"encoding/json"
	"os"
)

func PrintJSON(obj interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(obj)
}

func ReadFromStdin(obj interface{}) error {
	dec := json.NewDecoder(os.Stdin)
	return dec.Decode(obj)
}
