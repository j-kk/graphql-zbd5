package dtb

import (
	"context"
	"fmt"
	"github.com/j-kk/go-graphql/graph/model"
	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func InitDB(dtbUrl string) (db *DB, err error) {
	pool, err := pgxpool.Connect(context.Background(), dtbUrl)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	db = &DB{pool}

	return db, nil
}

func (db *DB) CloseDB() {
	db.pool.Close()
}

func retrieveId(tx *pgx.Tx, tableName string) (id int, err error) {

	rows, err := (*tx).Query(context.Background(), "SELECT * FROM currval($1)", tableName)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if i > 0 {
			return 0, fmt.Errorf("too many results (currval)")
		}
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	if rows.Err() != nil {
		return 0, rows.Err()
	}
	return
}

func (db *DB) addInterests(interests []string) (err error) { // TODO bulk insert
	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	for _, interest := range interests {
		_, err = tx.Exec(context.Background(), "INSERT INTO interests(name) VALUES ($1) ON CONFLICT DO NOTHING;", interest)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) matchInterests(interests []string, userId int) (err error) {
	var sqlCommand = fmt.Sprintf("INSERT INTO usersinterests (user_id, interest_id) SELECT usr.uid, interests.id FROM (VALUES (%v)) as usr (uid) CROSS JOIN (SELECT id FROM interests WHERE interests.name = ($1)) as interests;", userId)
	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	for _, interest := range interests {
		_, err = tx.Exec(context.Background(), sqlCommand, interest)
		if err != nil {
			return err
		}

	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) AddUser(user *model.User) (err error) {

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}

	// Add user
	defer tx.Rollback(context.Background())
	if user.GeoPos != nil {
		_, err = tx.Exec(context.Background(), "INSERT INTO GeoPositions (width, height) VALUES ($1,$2)", user.GeoPos.Width, user.GeoPos.Heigth)
		if err != nil {
			return err
		}
		_, err = tx.Exec(context.Background(), "INSERT INTO users (gender, birth_year, income, geopos_id) VALUES($1, $2, $3, currval('geopositions_id_seq'))", user.Gender, user.BirthYear, user.Income)
		if err != nil {
			return err
		}
		_, err = tx.Exec(context.Background(), "UPDATE geopositions SET user_id = lastval() WHERE id = currval('geopositions_id_seq')")
	} else {
		_, err = tx.Exec(context.Background(), "INSERT INTO users (gender, birth_year, income) VALUES($1, $2, $3)", user.Gender, user.BirthYear, user.Income)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	userId, err := retrieveId(&tx, "users_id_seq")
	if err != nil {
		return err
	}

	if user.Interests != nil {
		err = db.addInterests(user.Interests)
		if err != nil {
			return err
		}
		err = db.matchInterests(user.Interests, userId)
		if err != nil {
			return err
		}

	}

	return nil
}
