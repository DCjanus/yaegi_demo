// Package symbol is a package that contains all symbols of the project.
// Packages that you want to be used in the rule file should be added with `yaegi extract` command.
//
//go:generate go install github.com/traefik/yaegi/cmd/yaegi@v0.15.1
//go:generate yaegi extract github.com/dcjanus/yaegi_demo/internal/helper
//go:generate yaegi extract github.com/sirupsen/logrus
package symbol

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}
