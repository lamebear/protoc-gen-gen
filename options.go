package gserve

import (
	"errors"
	"log"
	"strings"
)

type Options struct {
	Package      string
	TemplatePath string
}

func ParseOptions(optionString *string) (Options, error) {
	o := Options{}

	if optionString == nil {
		return o, nil
	}

	var err error
	options := strings.Split(*optionString, ",")
	for _, option := range options {
		o, err = setOption(option, o)
		if err != nil {
			return o, err
		}
	}

	return o, nil
}

func setOption(optionString string, o Options) (Options, error) {
	parts := strings.Split(optionString, "=")
	if len(parts) == 1 || parts[1] == "" {
		return o, errors.New("Option value must be set")
	}

	switch parts[0] {
	case "package":
		o.Package = parts[1]
	case "template":
		o.TemplatePath = parts[1]
	default:
		log.Printf("Unknown option: %s=%s\n", parts[0], parts[1])
	}

	return o, nil
}
