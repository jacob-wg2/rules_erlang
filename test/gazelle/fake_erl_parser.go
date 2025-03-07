package erlang_test

import (
	erlang "github.com/rabbitmq/rules_erlang/gazelle"
)

type fakeErlParserCall struct {
	erlFile   string
	erlangApp *erlang.ErlangApp
	macros    erlang.ErlParserMacros
}

type erlParserFake struct {
	Responses map[string]*erlang.ErlAttrs
	Calls     []fakeErlParserCall
}

func fakeErlParser(responses map[string]*erlang.ErlAttrs) *erlParserFake {
	return &erlParserFake{Responses: responses}
}

func (p *erlParserFake) DeepParseErl(erlFile string, erlangApp *erlang.ErlangApp, macros erlang.ErlParserMacros) (*erlang.ErlAttrs, error) {
	p.Calls = append(p.Calls, fakeErlParserCall{
		erlFile:   erlFile,
		erlangApp: erlangApp,
		macros:    macros,
	})
	if r, ok := p.Responses[erlFile]; ok {
		return r, nil
	}
	return &erlang.ErlAttrs{}, nil
}
