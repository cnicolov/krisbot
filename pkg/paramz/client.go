package paramz

type ParameterProvider int

const (
	SSMParameterProvider ParameterProvider = 0
)

func New(config *Config) Provider {
	switch config.Provider {
	case SSMParameterProvider:
		return &SSMProvider{config: config, client: NewSSMClient()}
	default:
		panic("No such provider")
	}
}
