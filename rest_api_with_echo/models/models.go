package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/adwinugroho/go-collection/config"
	"github.com/arangodb/go-driver"
)

type (
	// create struct modelling data in DB
	MyURL struct {
		ID     string `json:"_key,omitempty"`
		Data   Data   `json:"data"`
		Logger Logger `json:"logger"`
	}

	// add struct Data
	Data struct {
		HashID   string `json:"hashId"`
		ShortUrl string `json:"shortUrl"`
		LongUrl  string `json:"longUrl"`
	}
	// for log
	Logger struct {
		CurrNo    int    `json:"curr_no"`
		Inputtime string `json:"inputtime"`
		LogReason string `json:"logReason"`
	}

	BindVars struct {
		KeyID    string
		Date     string
		RoomID   string
		IsActive bool    `json:"isActive"`
		Owner    string  `json:"owner,omitempty"`
		Offset   int     `json:"offset"`
		Limit    int     `json:"limit"`
		Page     int     `json:"page"`
		Search   Search  `json:"search"`
		Filters  Filters `json:"filter,omitempty"`
	}

	Filters struct {
		Date     []string `json:"date,omitempty"`
		Type     string   `json:"type,omitempty"`
		Category string   `json:"category,omitempty"`
	}

	// initial variable DB we use (db for database live and dbLog for database log)
	DB struct {
		dbLive driver.Database
		dbLog  driver.Database
	}
)

// get connection from config
func NewConnection(conn *config.Connection) *DB {
	return &DB{
		dbLive: conn.DBLive,
		dbLog:  conn.DBLog,
	}
}

// func add data to DB (insert)
func (db *DB) AddData(data MyURL) (*MyURL, error) {
	// variable parent context
	ctx := context.Background()
	// declare model we used
	myURL := MyURL{}
	// open connection to DB collection
	col, err := db.dbLive.Collection(ctx, "shorten_url")
	if err != nil {
		log.Printf("[models.go:AddData, db.Collection] Error open connection to collection, cause: %+v\n", err)
		return nil, err
	}
	//WithReturnNew is used to configure a context to make create, update & replace document
	driverCtx := driver.WithReturnNew(ctx, &myURL)
	// call func CreateDocument and passing driverCtx and parameter
	meta, err := col.CreateDocument(driverCtx, data)
	if err != nil {
		log.Printf("[models.go:AddData, col.CreateDoc] Error while creating document, cause: %+v\n", err)
		return nil, err
	}
	// passing _key in db in our model (myURL)
	myURL.ID = meta.Key
	// print _key
	fmt.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)
	// return deference model (myURL) and error == nil
	return &myURL, nil
}

func (db *DB) DeleteByKey(id string) (*[]MyURL, error) {
	ctx := context.Background()
	var result []MyURL
	var bindVars = map[string]interface{}{
		"keyid": id,
	}
	var query = "FOR x IN shorten_url FILTER x._key == @keyid REMOVE x IN shorten_url LET removed = OLD RETURN removed"
	cursor, err := db.dbLive.Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	for {
		var data MyURL
		_, err := cursor.ReadDocument(ctx, &data)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, errors.New("error read document")
		}
		result = append(result, data)
	}
	return &result, nil
}

func (db *DB) GetDataByKey(id string) (*MyURL, error) {
	ctx := context.Background()
	myURL := MyURL{}
	col, err := db.dbLive.Collection(ctx, "shorten_url")
	if err != nil {
		log.Printf("[models.go:GetDataByKey, db.Collection] Error open connection to collection, cause: %+v\n", err)
		return nil, err
	}
	_, err = col.ReadDocument(ctx, id, &myURL)
	if err != nil {
		log.Printf("[models.go:GetDataByKey, col.ReadDoc] Error reading document, cause: %+v\n", err)
		return nil, err
	}
	return &myURL, nil
}

func (db *DB) GetListAllData(vars BindVars) ([]MyURL, int64, error) {
	var query, queryCount, search string
	var filter bool
	var ctx = context.Background()
	var datas []MyURL
	var bindVars = map[string]interface{}{}
	//Search by
	if vars.Search.Text != "" {
		vars.Search.Text = "%" + vars.Search.Text + "%"
		search = fmt.Sprintf(`LIKE(x.data.longUrl, "%s", true)`, vars.Search.Text)
		filter = true
	}

	if filter {
		query = fmt.Sprintf(`FOR x IN shorten_url %s %s LIMIT %v, %v RETURN x`, "FILTER", search, vars.Offset, vars.Limit)
		queryCount = fmt.Sprintf(`FOR x IN shorten_url %s %s LIMIT %v, %v RETURN x`, "FILTER", search, vars.Offset, vars.Limit)
		//fmt.Printf("query: %+v\n", query)
	} else {
		query = fmt.Sprintf(`FOR x IN shorten_url LIMIT %v, %v RETURN x`, vars.Offset, vars.Limit)
		queryCount = `FOR x IN shorten_url RETURN x`
	}

	cursor, err := db.dbLive.Query(ctx, query, bindVars)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close()
	for {
		data := new(MyURL)
		meta, err := cursor.ReadDocument(ctx, &data)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Printf("Error reading document: %+v \n", err)
		}
		data.ID = meta.Key
		datas = append(datas, *data)
	}
	if len(datas) == 0 {
		return nil, 0, nil
	}

	ctxCount := driver.WithQueryCount(context.Background())
	cursorCount, err := db.dbLive.Query(ctxCount, queryCount, bindVars)
	if err != nil {
		fmt.Printf("[model:ListAllData]: Error execute query count [%s], cause: %+v \n", queryCount, err)
		return nil, 0, err
	}
	defer cursorCount.Close()
	fmt.Printf("query count func:%+v\n", cursorCount.Count())
	return datas, cursorCount.Count(), nil
}

//Add log to DB
func (db *DB) SaveLog(model *MyURL) (*string, error) {
	var ctx = context.Background()
	model.ID = fmt.Sprintf("%s-%s", model.ID, strconv.Itoa(model.Logger.CurrNo))
	col, err := db.dbLog.Collection(ctx, "shorten_url_log")
	if err != nil {
		return nil, err
	}

	meta, err := col.CreateDocument(ctx, &model)
	if err != nil {
		return nil, err
	}

	return &meta.Key, nil
}

//Update data
func (db *DB) UpdateData(model *MyURL) (*MyURL, error) {
	ctx := context.Background()
	data := MyURL{}
	col, err := db.dbLive.Collection(ctx, "shorten_url")
	if err != nil {
		return nil, err
	}

	driverCtx := driver.WithReturnNew(ctx, &data)
	meta, err := col.ReplaceDocument(driverCtx, model.ID, model)
	fmt.Printf("Doc Revision: %s \n", meta.Rev)
	if err != nil {
		return nil, err
	}

	data.ID = meta.Key
	return &data, nil
}
