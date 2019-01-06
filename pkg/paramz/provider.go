package paramz

type Provider interface {
	MustGetString(string, bool) string
}
