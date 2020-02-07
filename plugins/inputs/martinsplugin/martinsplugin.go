package martinsplugin

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

var MartinsConfig = `
  ## Set martins name
  name = Martin
`

type Martins struct {
	Cool string `toml:"name"`
}

func (s *Martins) SampleConfig() string {
	return MartinsConfig
}

func (s *Martins) Description() string {
	return "This is martin's name. You can rename martin."
}
func (s *Martins) Gather(acc telegraf.Accumulator) error {
	// We can do whatever we want on the OS here
	fields := make(map[string]interface{})
	fields["name"] = "Yaw"
	tags := make(map[string]string)
	acc.AddFields("martins", fields, tags)
	return nil
}
func init() {
	inputs.Add("martins", func() telegraf.Input {
		return &Martins{Cool: "martin"}
	})
}
