package database

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Client struct {
	db     *sql.DB
	logger zerolog.Logger

	Address string `key:"address" description:"address and port of the databse" default:"database"`

	Username string `key:"username" description:"username of the databse" default:"pgx"`
	Password string `key:"password" description:"password of the databse" default:"pgx2021"`

	Database string `key:"database" description:"databse" default:"wallet"`
}

func (c *Client) Init() (err error) {
	c.db, err = sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", c.Username, c.Password, c.Address, c.Database))
	if err != nil {
		log.Error().Err(err).Msgf("Unable to connect to database: %v", err)
		return err
	}

	c.logger = log.With().Str("module", "database").Logger()
	c.logger.Debug().Msgf("Connected to %s", c.Address)

	return nil
}
