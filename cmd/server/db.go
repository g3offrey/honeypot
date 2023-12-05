package main

import (
	"github.com/g3offrey/honeypot/internal/config"
	"github.com/g3offrey/honeypot/internal/persistence"
	"github.com/g3offrey/honeypot/pkg/catch"
)

type repositories struct {
	catch catch.CatchRepository
}

func getRepositories(cfg *config.Config) (repositories, error) {
	db, err := persistence.Init(cfg.DBPath)
	if err != nil {
		return repositories{}, err
	}

	return repositories{
		catch: persistence.NewGormCatchRepository(db),
	}, nil
}
