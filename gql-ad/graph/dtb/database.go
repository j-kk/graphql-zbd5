package dtb

import (
	"context"
	"fmt"
	"github.com/j-kk/go-graphql/graph/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
func (db *DB) addAdWords(words *[]string, adId int) (err error) {

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil

}

func (db *DB) AddAd(ad *model.Ad) (err error) {

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	if ad.Dimensions != nil {
		_, err = tx.Exec(context.Background(), "INSERT INTO ads(width, height, main_color) VALUES ($1, $2, $3);", ad.Dimensions.Width, ad.Dimensions.Height, ad.MainColor)
	} else {
		_, err = tx.Exec(context.Background(), "INSERT INTO ads(main_color) VALUES ($1);", ad.MainColor)
	}
	if err != nil {
		return err
	}
	ad.ID, err = retrieveId(&tx, "ads_id_seq")

	if len(ad.Texts) > 0 {
		for _, word := range ad.Texts {
			_, err = tx.Exec(context.Background(), "INSERT INTO adwords (word, ad_id) VALUES ($1, $2);", word, ad.ID)
			if err != nil {
				return err
			}
		}
	}

	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil

}
