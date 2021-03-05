package pong2trans

type Translator interface {
	Translate(in string) (out string)
}
