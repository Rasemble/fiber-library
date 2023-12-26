package database

import "github.com/Rasemble/Api-fiber-library/app/queries"

// Queries struct to collect all app queries
type Queries struct {
	*queries.BookQueries
}

// Func to open database connection
func OpenDBConnection() (*Queries, error) {
	// Define a new connection
	db, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}
	return &Queries{
		BookQueries: &queries.BookQueries{DB: db},
	}, nil
}
