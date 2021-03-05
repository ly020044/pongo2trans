package pong2trans

import (
	"github.com/flosch/pongo2/v4"
)

var t Translator

type tagTransNode struct {
	name       string
	translator Translator
	value      pongo2.IEvaluator
}

func newTagTransNode(value pongo2.IEvaluator) *tagTransNode {
	return &tagTransNode{
		translator: t,
		value:      value,
	}
}

func (node *tagTransNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	value, err := node.value.Evaluate(ctx)
	if err != nil {
		return err
	}

	if node.translator == nil {
		writer.WriteString(value.String())
		return nil
	}

	writer.WriteString(node.translator.Translate(value.String()))

	return nil
}

func tagTransParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	value, err := arguments.ParseExpression()
	if err != nil {
		return nil, err
	}

	node := newTagTransNode(value)
	if arguments.Remaining() > 0 {
		return nil, arguments.Error("error trans tag", nil)
	}

	return node, nil
}

func RegisterTransTag(translator Translator) error {
	t = translator
	return pongo2.RegisterTag("trans", tagTransParser)
}
