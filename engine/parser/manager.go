package parser

import "olive/engine/log"

var SharedManager = &Manager{}

type Manager struct {
	savers map[string]Parser
}

func (p *Manager) Register(parsers ...Parser) {
	if p.savers == nil {
		p.savers = map[string]Parser{}
	}
	for _, parser := range parsers {
		_, ok := p.savers[parser.Type()]
		if ok {
			log.Logger.Warnf("parser %s has registered", parser.Type())
		}
		p.savers[parser.Type()] = parser
	}
}

func (p *Manager) Parser(typ string) (Parser, bool) {
	v, ok := p.savers[typ]
	return v, ok
}

type Parser interface {
	New() Parser
	Type() string
	Parse(streamURL string, out string) error
	Stop()
}
