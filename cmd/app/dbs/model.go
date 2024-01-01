package dbs

import (
	"encoding/json"
	"fmt"
)

type Sticker struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Stock uint   `json:"stock"`
	Price string `json:"price"`
}

func FetchAll() ([]*Sticker, error) {
	var collection, collectionErr = postgres.Query(`SELECT * FROM stickers;`)
	if collectionErr != nil {
		return nil, collectionErr
	}
	defer collection.Close()
	var stickers = []*Sticker{}
	for collection.Next() {
		var sticker = &Sticker{}
		var err = collection.Scan(&sticker.ID, &sticker.Name, &sticker.Stock, &sticker.Price)
		if err != nil {
			continue
		}
		stickers = append(stickers, sticker)
	}
	return stickers, nil
}

func FetchByID(id string) (Sticker, error) {
	var collection, collectionErr = postgres.Query(`SELECT * FROM stickers WHERE id=$1`, id)
	if collectionErr != nil {
		return Sticker{}, collectionErr
	}
	defer collection.Close()
	var sticker = Sticker{}
	for collection.Next() {
		var err = collection.Scan(&sticker.ID, &sticker.Name, &sticker.Stock, &sticker.Price)
		if err != nil {
			return Sticker{}, err
		}
	}
	return sticker, nil
}

func InsertOne(b []byte) (Sticker, error) {
	// done
	var sticker = Sticker{}
	json.Unmarshal(b, &sticker)
	var query = fmt.Sprintf(`INSERT INTO stickers(name,stock,price) VALUES('%v',%v,%v);`, sticker.Name, sticker.Stock, sticker.Price)
	var collection, collectionErr = postgres.Query(query)
	if collectionErr != nil {
		return Sticker{}, collectionErr
	}
	defer collection.Close()
	for collection.Next() {
		var err = collection.Scan(&sticker)
		if err != nil {
			return Sticker{}, err
		}
	}
	return sticker, nil
}

func Update(id string, b []byte) (Sticker, error) {
	// done
	var sticker = Sticker{}
	json.Unmarshal(b, &sticker)
	var query = fmt.Sprintf(`UPDATE stickers SET name='%v', stock=%v, price=%v WHERE id=%v`, sticker.Name, sticker.Stock, sticker.Price, id)
	var collection, collectionErr = postgres.Query(query)
	if collectionErr != nil {
		return Sticker{}, collectionErr
	}
	defer collection.Close()
	for collection.Next() {
		var err = collection.Scan(&sticker)
		if err != nil {
			return Sticker{}, err
		}
	}
	return sticker, nil
}

func DeleteOne(id string) (Sticker, error) {
	// done
	var collection, collectionErr = postgres.Query(`DELETE FROM stickers WHERE id=$1`, id)
	if collectionErr != nil {
		return Sticker{}, collectionErr
	}
	defer collection.Close()
	var sticker = Sticker{}
	for collection.Next() {
		var err = collection.Scan(&sticker)
		if err != nil {
			return Sticker{}, err
		}
	}
	return sticker, nil
}
