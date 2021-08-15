package config

type Rollbar struct {
	Enabled bool
	Token   string
}

func (r Rollbar) IsEnabled() bool {
	return r.Enabled && r.Token != ""
}
