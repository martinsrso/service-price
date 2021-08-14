// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate" //migrate
	_ "golang.org/x/tools/cmd/goimports"                 // updates imports and formats code
)
