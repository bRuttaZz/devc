package config

import (
	_ "embed"
)

//go:embed general.yml
var General []byte
