package config

import (
	"errors"
	"fmt"
)

type Specification struct {
	Envirnoment Enviroment
	Port        string
}

type Enviroment string

const (
	Docker      Enviroment = "docker"
	Development Enviroment = "development"
)

func ValidateSpecification(s Specification) error {
	fmt.Println("enviroment is", s.Envirnoment)
	fmt.Println("port is", s.Port)
	if s.Envirnoment == "" {
		return errors.New("Enviroment is required")
	}

	if s.Port == "" {
		return errors.New("port is required")
	}

	return nil
}
