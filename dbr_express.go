// Package vasiliyaltunin/dbr_express is gocraft/dbr wrapper that allow you do simple sql query - select, insert, update and delete just in one string
package dbre

import (
	"reflect"

	"github.com/vasiliyaltunin/dbr"
)

type dbrExpressStmt struct {
	Session *dbr.Session
}

var sess dbr.Session

func DbrExpress(session *dbr.Session) *dbrExpressStmt {
	return &dbrExpressStmt{Session: session}
}

//Select - selects data from database
func (s *dbrExpressStmt) Select(table, fields, where string, data interface{}) interface{} {

	var ptr reflect.Value

	ptr = reflect.New(reflect.TypeOf(data)) // create new pointer

	obj := ptr.Interface()

	sess := s.Session

	query := sess.Select(fields).
		From(table)

	if len(where) > 0 {
		query.Where(where)
	}

	_, err := query.Load(obj)

	if err != nil {
		panic(err)
	}

	return obj

}

//Insert - inserts data into database
func (s *dbrExpressStmt) Insert(table string, columns []string, data interface{}) interface{} {

	obj := reflect.ValueOf(data).Interface()

	sess := s.Session

	sess.InsertInto(table).
		Columns(columns...).
		Record(obj).
		Exec()

	// id is set automatically
	return obj

}

//Update - updates data into database
func (s *dbrExpressStmt) Update(table string, columns []string, data interface{}, whereStr, whereValue string) {

	obj := reflect.ValueOf(data).Interface()

	sess := s.Session

	// m := structs.Map(data)

	sess.Update(table).
		Columns(columns...).
		Record(obj).
		Where(whereStr, whereValue).
		Exec()

}

//Delete - updates data into database
func (s *dbrExpressStmt) Delete(table string, whereStr, whereValue string) {

	sess := s.Session

	sess.DeleteFrom(table).
		Where(whereStr, whereValue).
		Exec()

}