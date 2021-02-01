package dtb

import (
	"context"
	"fmt"
	"github.com/j-kk/go-graphql/graph/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
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

func (db *DB) RegisterView(adView *model.View) error {

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO views(ads_id, user_id) VALUES ($1, $2);", adView.AdID, adView.UserID)
	if err != nil {
		return err
	}

	adView.ID, err = retrieveId(&tx, "views_id_seq")
	if err != nil {
		return err
	}

	rows, err := tx.Query(context.Background(), "SELECT t FROM views WHERE id = $1;", adView.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var timestamp time.Time

	for i := 0; rows.Next(); i++ {
		if i > 0 {
			return fmt.Errorf("too many results (currval)")
		}
		err = rows.Scan(&timestamp)
		if err != nil {
			return err
		}
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	adView.Timestamp = timestamp.String()

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetUser(userID int) (*model.User, error) {

	var newUser model.User
	newUser.ID = userID

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "SELECT gender, birth_year, income, geopos_id FROM users WHERE id = $1", newUser.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse user
	for i := 0; rows.Next(); i++ {
		if i > 0 {
			return nil, fmt.Errorf("too many results")
		}
		err = rows.Scan(&newUser.Gender, &newUser.BirthYear, &newUser.Income, &newUser.GeoPosID)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	interests, err := db.GetUserInterests(newUser.ID)
	if err != nil {
		return nil, err
	}
	newUser.Interests = *interests

	return &newUser, nil
}

func (db *DB) GetUserInterests(userID int) (*[]string, error) {

	var interests []string

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "SELECT i.name FROM usersinterests NATURAL JOIN interests i WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse user
	for i := 0; rows.Next(); i++ {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		interests = append(interests, name)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &interests, nil
}

func (db *DB) GetPosition(posID int) (*model.Position, error) {

	var pos model.Position
	pos.ID = posID

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "SELECT width, height FROM geopositions WHERE id = $1", posID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse user
	for i := 0; rows.Next(); i++ {
		if i > 0 {
			return nil, fmt.Errorf("too many results (currval)")
		}
		err = rows.Scan(&pos.Width, pos.Heigth)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &pos, nil
}

func (db *DB) GetAd(adID int) (*model.Ad, error) {

	var newAd model.Ad
	var adDim model.AdDimensions
	newAd.ID = adID

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "SELECT width, height, main_color FROM ads WHERE id = $1", newAd.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse user
	for i := 0; rows.Next(); i++ {
		if i > 0 {
			return nil, fmt.Errorf("too many results (currval)")
		}
		err = rows.Scan(&adDim.Width, &adDim.Height, &newAd.MainColor)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	newAd.Dimensions = &adDim

	words, err := db.GetAdWords(newAd.ID)
	if err != nil {
		return nil, err
	}
	newAd.Texts = *words

	return &newAd, nil
}

func (db *DB) GetAdWords(adID int) (*[]string, error) {

	var words []string

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "SELECT word FROM adwords WHERE ad_id = $1", adID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse user
	for i := 0; rows.Next(); i++ {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		words = append(words, name)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &words, nil
}
