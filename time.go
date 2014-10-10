package pflag

import "time"

const modelTime string = "Mon Jan 2 15:04:05 MST 2006"

// -- time.Time Value
type timeValue time.Time

var timeFormat string = "2006-01-02" // YYYY MM DD
func SetTimeFormat(s string) error {
	_, err := time.Parse(modelTime, s)
	if err != nil {
		return err
	}

	timeFormat = s
	return nil
}

func newTimeValue(val time.Time, p *time.Time) *timeValue {
	*p = val
	return (*timeValue)(p)
}

func (d *timeValue) Set(s string) error {
	v, err := time.Parse(timeFormat, s)
	*d = timeValue(v)
	return err
}

func (d *timeValue) String() string {
	ret := (*time.Time)(d).String()
	return ret
}

// TimeVar defines a time.Time flag with specified name, default value, and usage string.
// The argument p points to a time.Time variable in which to store the value of the flag.
func (f *FlagSet) TimeVar(p *time.Time, name string, value time.Time, usage string) {
	f.VarP(newTimeValue(value, p), name, "", usage)
}

// Like TimeVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) TimeVarP(p *time.Time, name, shorthand string, value time.Time, usage string) {
	f.VarP(newTimeValue(value, p), name, shorthand, usage)
}

// TimeVar defines a time.Time flag with specified name, default value, and usage string.
// The argument p points to a time.Time variable in which to store the value of the flag.
func TimeVar(p *time.Time, name string, value time.Time, usage string) {
	CommandLine.VarP(newTimeValue(value, p), name, "", usage)
}

// Like TimeVar, but accepts a shorthand letter that can be used after a single dash.
func TimeVarP(p *time.Time, name, shorthand string, value time.Time, usage string) {
	CommandLine.VarP(newTimeValue(value, p), name, shorthand, usage)
}

// Time defines a time.Time flag with specified name, default value, and usage string.
// The return value is the address of a time.Time variable that stores the value of the flag.
func (f *FlagSet) Time(name string, value time.Time, usage string) *time.Time {
	p := new(time.Time)
	f.TimeVarP(p, name, "", value, usage)
	return p
}

// Like Time, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) TimeP(name, shorthand string, value time.Time, usage string) *time.Time {
	p := new(time.Time)
	f.TimeVarP(p, name, shorthand, value, usage)
	return p
}

// Time defines a time.Time flag with specified name, default value, and usage string.
// The return value is the address of a time.Time variable that stores the value of the flag.
func Time(name string, value time.Time, usage string) *time.Time {
	return CommandLine.TimeP(name, "", value, usage)
}

// Like Time, but accepts a shorthand letter that can be used after a single dash.
func TimeP(name, shorthand string, value time.Time, usage string) *time.Time {
	return CommandLine.TimeP(name, shorthand, value, usage)
}
