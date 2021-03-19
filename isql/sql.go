package isql

import "strings"

// SQLHandler Sql Handler
type SQLHandler struct {
	segment []string
	param   []interface{}
}

// NewSQLHandler New SQLHandler
func NewSQLHandler() *SQLHandler {
	return &SQLHandler{
		segment: make([]string, 0),
		param:   make([]interface{}, 0),
	}
}

// Append append
func (s *SQLHandler) Append(sql string, args ...interface{}) {
	s.segment = append(s.segment, sql)
	s.param = append(s.param, args...)
}

func (s *SQLHandler) buildSQL(sep string) string {
	if sep == "" {
		sep = " "
	}
	return strings.Join(s.segment, " ")
}

// Build build
func (s *SQLHandler) Build(sep string) (sql string, args []interface{}) {
	return s.buildSQL(sep), s.param
}
