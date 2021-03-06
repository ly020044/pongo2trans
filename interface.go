package pong2trans

type Exporter interface {
	Export(value string)
}

type Translator interface {
	Translate(in string) (out string)
}
