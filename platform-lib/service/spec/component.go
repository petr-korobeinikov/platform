package spec

import (
	"fmt"
	"strings"
)

type Component struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

func (c *Component) ID() string {
	return fmt.Sprintf("component-%s-%s", c.Type, c.Name)
}

func (c *Component) FormatEnvVarName(v string) string {
	up := strings.ToUpper(c.ID())
	uv := strings.ToUpper(v)

	return strings.ReplaceAll(
		fmt.Sprintf("%s_%s", up, uv),
		"-",
		"_",
	)
}

func (c *Component) FormatEnvVarNameEscaped(v string) string {
	return fmt.Sprintf("${%s}", c.FormatEnvVarName(v))
}
