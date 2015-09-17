package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Server struct {
	ServerId   string
	ServerName string
}

func (s *Server) Validate() error {
	if (s.ServerId == "") == (s.ServerName == "") {
		return fmt.Errorf("Exactly one of the server-id and server-name must be set.")
	}
	return nil
}

func (s *Server) InferID(cn base.Connection) error {
	if s.ServerName == "" {
		return nil
	}

	ID, err := IDByName(cn, s.ServerName)
	if err != nil {
		return err
	}
	s.ServerId = ID
	return nil
}

func (s *Server) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "ServerName" {
		return nil, nil
	}

	return GetNames(cn, "all")
}
