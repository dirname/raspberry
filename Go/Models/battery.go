package Models

import (
	"github.com/pkg/errors"
	"regexp"
)

type BatteryReport struct {
	State uint
	Info  string
	ID    int
}

type UPS struct {
	Info          string
	Version       string
	Capacity      string
	OutputVoltage string
	InVoltage     string
}

func (u *UPS) Parse() error {
	re := regexp.MustCompile("\\$ (.*?) \\$")
	res := re.FindAllString(u.Info, -1)
	if len(res) > 0 {
		tmp := res[0]
		re = regexp.MustCompile("SmartUPS (.*?),")
		u.Version = re.ReplaceAllString(re.FindString(tmp), "${1}")
		re = regexp.MustCompile(",Vin (.*?),")
		if vin := re.ReplaceAllString(re.FindString(tmp), "${1}"); vin == "NG" {
			u.InVoltage = "Not charging"
		} else {
			u.InVoltage = vin
		}
		re = regexp.MustCompile("BATCAP (.*?),")
		u.Capacity = re.ReplaceAllString(re.FindString(tmp), "${1}")
		re = regexp.MustCompile(",Vout (.*) \\$")
		u.OutputVoltage = re.ReplaceAllString(re.FindString(tmp), "${1}")
		return nil
	} else {
		return errors.New("Not a normal signal")
	}
}
