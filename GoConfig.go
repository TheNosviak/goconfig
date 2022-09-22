package goconfig


// GoConfig allows you to configure the current working env before actually running the execution routes
type GoConfig struct{
	configs    map[string]any
	inclusions map[string]Inclusion
}

// NewConfig will create a brand new GoConfig instance
func NewConfig() *GoConfig {
	return &GoConfig{
		inclusions: Defaults,
		configs:    nil,
	}
}

// Parse will completely scan all the directories passed inside the vardiac function args
func (GC *GoConfig) Parse(directories ...string) error {
	if GC.configs == nil {
		GC.configs = make(map[string]any)
	}

	// ranges through all the directories provided
	for _, dir := range directories {
		if err := GC.execDir(dir); err != nil {
			return err
		}
	}

	return nil
}