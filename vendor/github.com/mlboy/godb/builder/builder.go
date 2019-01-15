package builder

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var tagKey = "db"
var identKey = "`"

//SQLSegments ...
type SQLSegments struct {
	table   []Table
	fields  []string
	flags   []string
	join    []map[string]string
	where   Clause
	groupBy []string
	having  Clause
	orderBy []string
	limit   struct {
		limit  int
		offset int
	}
	union     []func(*SQLSegments)
	forUpdate bool
	returning bool
	// params    []interface{}
	params []map[string]interface{}
	render struct {
		args []interface{}
	}
	checks map[int]bool
}

//Table ..
type Table struct {
	Name  string
	Alias string
}

//Add ..
type Add int

//Sub ..
type Sub int

//New ..
func New() *SQLSegments {
	return &SQLSegments{checks: make(map[int]bool, 0)}
}

//NewSQLSegment ..
func NewSQLSegment() *SQLSegments {
	return &SQLSegments{}
}

//Table SQLSegments
func (s *SQLSegments) Table(name interface{}) *SQLSegments {
	switch v := name.(type) {
	case Table:
		s.table = append(s.table, v)
	case []Table:
		s.table = append(s.table, v...)
	case string:
		var t = &Table{v, ""}
		s.table = append(s.table, *t)
	}
	return s
}

//Field SQLSegments
func (s *SQLSegments) Field(fields ...string) *SQLSegments {
	if len(fields) > 0 {
		s.fields = append(s.fields, fields...)
	}
	return s
}

//Flag SQLSegments
func (s *SQLSegments) Flag(flags ...string) *SQLSegments {
	if len(flags) > 0 {
		s.flags = append(s.flags, flags...)
	}
	return s
}

//Join SQLSegments
func (s *SQLSegments) Join(table string, conditionA, logic, conditionB string) *SQLSegments {
	s.addJoin("", table, conditionA, logic, conditionB)
	return s
}

//LeftJoin SQLSegments
func (s *SQLSegments) LeftJoin(table string, conditionA, logic, conditionB string) *SQLSegments {
	s.addJoin("LEFT", table, conditionA, logic, conditionB)
	return s
}

//RightJoin SQLSegments
func (s *SQLSegments) RightJoin(table string, conditionA, logic, conditionB string) *SQLSegments {
	s.addJoin("RIGHT", table, conditionA, logic, conditionB)
	return s
}

//InnerJoin SQLSegments
func (s *SQLSegments) InnerJoin(table string, conditionA, logic, conditionB string) *SQLSegments {
	s.addJoin("INNER", table, conditionA, logic, conditionB)
	return s
}

//CorssJoin SQLSegments
func (s *SQLSegments) CorssJoin(table string, conditionA, logic, conditionB string) *SQLSegments {
	s.addJoin("CROSS", table, conditionA, logic, conditionB)
	return s
}

//addJoin SQLSegments
func (s *SQLSegments) addJoin(typ string, table string, conditionA, logic, conditionB string) *SQLSegments {
	var t = make(map[string]string)
	t["type"] = typ
	t["table"] = table
	t["logic"] = logic
	t["conditionA"] = conditionA
	t["conditionB"] = conditionB
	s.join = append(s.join, t)
	return s
}

//OrderBy SQLSegments
func (s *SQLSegments) OrderBy(fields ...string) *SQLSegments {
	if len(fields) > 0 {
		s.orderBy = append(s.orderBy, fields...)
	}
	return s
}

//GroupBy SQLSegments
func (s *SQLSegments) GroupBy(fields ...string) *SQLSegments {
	if len(fields) > 0 {
		s.groupBy = append(s.groupBy, fields...)
	}
	return s
}

//Offset SQLSegments
func (s *SQLSegments) Offset(n int) *SQLSegments {
	s.limit.offset = n
	return s
}

//Limit SQLSegments
func (s *SQLSegments) Limit(n int) *SQLSegments {
	s.limit.limit = n
	return s
}

//ForUpdate SQLSegments
func (s *SQLSegments) ForUpdate() *SQLSegments {
	s.forUpdate = true
	return s
}

//Clause ...
type Clause struct {
	key    interface{}
	val    interface{}
	logic  string
	clause []*Clause
}

func (p *Clause) addClause(logic string, key interface{}, vals ...interface{}) *Clause {
	var c = &Clause{}
	c.logic = logic
	switch k := key.(type) {
	case func(*Clause):
		k(c)
		// p.clause = append(p.clause, c)
	default:
		c.key = key
		if len(vals) > 0 {
			c.val = vals[0]
		}
	}
	// fmt.Println(p.clause)
	p.clause = append(p.clause, c)
	return p
}

//Where ..
func (p *Clause) Where(key interface{}, vals ...interface{}) *Clause {
	p.addClause("AND", key, vals...)
	return p
}

//OrWhere ..
func (p *Clause) OrWhere(key interface{}, vals ...interface{}) *Clause {
	p.addClause("OR", key, vals...)
	return p
}

//Build ...
func (p *Clause) Build(i int) (string, []interface{}) {
	var sql = ""
	var args []interface{}
	if p.logic != "" && i > 0 {
		sql += " " + p.logic
	}
	switch k := p.key.(type) {
	case string:
		r, _ := regexp.Compile(`\[(\>\=|\<\=|\>|\<|\<\>|\!\=|\=|\~|\!\~|like|!like|in|!in|is|!is|exists|!exists|#)\]?([a-zA-Z0-9_.\-\=\s\?\(\)]*)`)
		match := r.FindStringSubmatch(k)
		var context string
		if len(match) > 0 {
			// fmt.Println(len(match), match[1])
			switch match[1] {
			case "~", "like":
				context = buildIdent(match[2]) + " LIKE ?"
				args = append(args, p.val)
			case "!~", "!like":
				context = buildIdent(match[2]) + "` NOT LIKE ?"
				args = append(args, p.val)
			case ">":
				context = buildIdent(match[2]) + " > ?"
				args = append(args, p.val)
			case ">=":
				context = buildIdent(match[2]) + "` >= ?"
				args = append(args, p.val)
			case "<":
				context = buildIdent(match[2]) + " < ?"
				args = append(args, p.val)
			case "<=":
				context = buildIdent(match[2]) + " <= ?"
				args = append(args, p.val)
			case "<>", "!=":
				context = buildIdent(match[2]) + " != ?"
				args = append(args, p.val)
			case "=":
				context = buildIdent(match[2]) + " = ?"
				args = append(args, p.val)
			case "in":
				context = buildIdent(match[2]) + " IN ("
				var holder string
				if reflect.TypeOf(p.val).Kind() == reflect.Slice {
					v := reflect.ValueOf(p.val)
					holder = buildPlaceholder(v.Len(), "?", " ,")
					for n := 0; n < v.Len(); n++ {
						args = append(args, v.Index(n).Interface())
					}
				} else {
					holder = "?"
					args = append(args, p.val)
				}
				context += holder + ")"
			case "!in":
				context = buildIdent(match[2]) + " NOT IN ("
				var holder string
				if reflect.TypeOf(p.val).Kind() == reflect.Slice {
					v := reflect.ValueOf(p.val)
					holder = buildPlaceholder(v.Len(), "?", " ,")
					for n := 0; n < v.Len(); n++ {
						args = append(args, v.Index(n).Interface())
					}
				} else {
					holder = "?"
					args = append(args, p.val)
				}
				context += holder + ")"
			case "exists":
				switch p.val.(type) {
				case string:
					context = "EXISTS (" + p.val.(string) + ")"
				case func(s *SQLSegments):
					s := NewSQLSegment()
					p.val.(func(s *SQLSegments))(s)
					context = "EXISTS (" + s.BuildSelect() + ")"
					args = append(args, s.render.args...)
				}
			case "!exists":
				switch p.val.(type) {
				case string:
					context = "NOT EXISTS (" + p.val.(string) + ")"
				case func(s *SQLSegments):
					s := NewSQLSegment()
					p.val.(func(s *SQLSegments))(s)
					context = "NOT EXISTS (" + s.BuildSelect() + ")"
					args = append(args, s.render.args...)
				}
			case "is":
				if p.val == nil {
					context = buildIdent(match[2]) + " IS NULL"
				} else {
					context = buildIdent(match[2]) + " IS ?"
					args = append(args, p.val)
				}
			case "!is":
				if p.val == nil {
					context = buildIdent(match[2]) + " IS NOT NULL"
				} else {
					context = buildIdent(match[2]) + " IS NOT ?"
					args = append(args, p.val)
				}
			case "#":
				context = match[2]
				if reflect.TypeOf(p.val).Kind() == reflect.Slice {
					v := reflect.ValueOf(p.val)
					for n := 0; n < v.Len(); n++ {
						args = append(args, v.Index(n).Interface())
					}
				} else {
					args = append(args, p.val)
				}
			}
			sql += " " + context
		} else {
			if p.val != nil {
				sql += " " + buildIdent(k) + " = ?"
				args = append(args, p.val)
			} else {
				sql += " " + k
			}
		}
	case nil:
		sql += " ("
		for j, c := range p.clause {
			part, arg := c.Build(j)
			sql += part
			args = append(args, arg...)
		}
		sql += ")"
	}
	return sql, args
}

//Where ..
func (s *SQLSegments) Where(key interface{}, vals ...interface{}) *SQLSegments {
	s.where.Where(key, vals...)
	return s
}

//OrWhere ..
func (s *SQLSegments) OrWhere(key interface{}, vals ...interface{}) *SQLSegments {
	s.where.OrWhere(key, vals...)
	return s
}

//BuildWhereClause ...
func (s *SQLSegments) buildWhereClause() string {
	var sql string
	if len(s.where.clause) > 0 {
		sql = " WHERE"
		for i, c := range s.where.clause {
			part, args := c.Build(i)
			sql += part
			s.render.args = append(s.render.args, args...)
		}
	}
	return sql
}

//Having ...
func (s *SQLSegments) Having(key interface{}, vals ...interface{}) *SQLSegments {
	s.having.Where(key, vals...)
	return s
}

//buildHavingClause ...
func (s *SQLSegments) buildHavingClause() string {
	var sql string
	if len(s.having.clause) > 0 {
		sql = " HAVING"
		for i, c := range s.having.clause {
			part, args := c.Build(i)
			sql += part
			s.render.args = append(s.render.args, args...)
		}
	}
	return sql
}
func (s *SQLSegments) buildFlags() string {
	var sql string
	for _, v := range s.flags {
		sql += " " + v
	}
	return sql

}
func (s *SQLSegments) buildField() string {
	var sql string
	if len(s.fields) == 0 {
		sql += " *"
	} else {
		for i, v := range s.fields {
			if i > 0 {
				sql += ","
			}
			if v == "*" {
				sql += " " + v
			} else {
				sql += " " + buildIdent(v)
			}
		}
	}
	return sql
}
func (s *SQLSegments) buildTable() string {
	var sql string
	for i, v := range s.table {
		if i > 0 {
			sql += ","
		}
		sql += " " + buildIdent(v.Name)
		if v.Alias != "" {
			sql += " AS " + buildIdent(v.Alias)
		}
	}
	return sql
}
func (s *SQLSegments) buildJoin() string {
	var sql string
	for _, t := range s.join {
		sql += " " + t["type"] + "JOIN " + buildIdent(t["table"]) + " ON " + buildIdent(t["conditionA"]) + " " + t["logic"] + " " + buildIdent(t["conditionB"])
	}
	return sql
}
func (s *SQLSegments) buildGroupBy() string {
	var sql string
	if len(s.groupBy) > 0 {
		sql += " GROUP BY"
	}
	for i, v := range s.groupBy {
		if i > 0 {
			sql += ","
		}
		sql += " " + buildIdent(v)
	}
	return sql
}
func (s *SQLSegments) buildOrderBy() string {
	var sql string
	if len(s.orderBy) > 0 {
		sql += " ORDER BY"
	}
	for i, v := range s.orderBy {
		if i > 0 {
			sql += ","
		}
		sql += " " + buildIdent(v)
	}
	return sql
}
func (s *SQLSegments) buildLimit() string {
	var sql string
	if s.limit.limit != 0 {
		sql += fmt.Sprintf(" LIMIT %d", s.limit.limit)
	}
	if s.limit.offset != 0 {
		sql += fmt.Sprintf(" OFFSET %d", s.limit.offset)
	}
	return sql
}

//Union ...
func (s *SQLSegments) Union(f func(*SQLSegments)) *SQLSegments {
	s.union = append(s.union, f)
	return s
}
func (s *SQLSegments) buildUnion() string {
	var sql string
	if len(s.union) > 0 {
		sql += " UNION ("
	}
	for _, f := range s.union {
		var ss = &SQLSegments{}
		f(ss)
		sql += ss.BuildSelect()
	}
	if len(s.union) > 0 {
		sql += ")"
	}
	return sql
}
func (s *SQLSegments) buildForUpdate() string {
	if s.forUpdate == true {
		return " FOR UPDATE"
	}
	return ""
}

//BuildSelect ...
func (s *SQLSegments) BuildSelect() string {
	var sql = fmt.Sprintf("SELECT%s%s FROM%s%s%s%s%s%s%s%s%s",
		s.buildFlags(),
		s.buildField(),
		s.buildTable(),
		s.buildJoin(),
		s.buildWhereClause(),
		s.buildGroupBy(),
		s.buildHavingClause(),
		s.buildOrderBy(),
		s.buildLimit(),
		s.buildUnion(),
		s.buildForUpdate(),
	)
	// fmt.Println(s.render.args)
	return sql
}

//Insert ...
func (s *SQLSegments) Insert(vals ...map[string]interface{}) *SQLSegments {
	s.params = append(s.params, vals...)
	return s
}

//BuildInsert ...
func (s *SQLSegments) BuildInsert() string {
	var sql = fmt.Sprintf("INSERT%s INTO%s%s%s",
		s.buildFlags(),
		s.buildTable(),
		s.buildValuesForInsert(),
		s.buildReturning(),
	)
	return sql
}

//BuildReplace ...
func (s *SQLSegments) BuildReplace() string {
	var sql = fmt.Sprintf("REPLACE%s INTO%s%s%s",
		s.buildFlags(),
		s.buildTable(),
		s.buildValuesForInsert(),
		s.buildReturning(),
	)
	return sql
}

//BuildInsert ...
func (s *SQLSegments) buildValuesForInsert() string {
	var fields string
	var values string
	var fieldSlice []string
	for i, vals := range s.params {
		if i == 0 {
			for arg := range vals {
				fieldSlice = append(fieldSlice, arg)
			}
		}
	}
	fieldLen := len(fieldSlice)
	fields += buildString(fieldSlice, ",", " (", ")", true)
	for i, vals := range s.params {
		if i == 0 {
			values += " ("
		} else {
			values += ",("
		}
		for _, arg := range fieldSlice {
			s.render.args = append(s.render.args, vals[arg])
		}
		values += buildPlaceholder(fieldLen, "?", ",")
		values += ")"
	}
	var sql = fields + " VALUES" + values
	return sql
}

// func (s *SQLSegments) buildValuesForInsert() string {
// 	var fields string
// 	var values string
// 	for i, param := range s.params {
// 		v := reflect.ValueOf(param).Elem()
// 		t := reflect.TypeOf(param).Elem()
// 		if i == 0 {
// 			values += " ("
// 			fields += " ("
// 		} else {
// 			values += ",("
// 		}
// 		for j := 0; j < v.NumField(); j++ {
// 			if v.Interface() == nil {
// 				continue
// 			}
// 			var arg string
// 			if t.Field(j).Tag.Get(tagKey) == "" {
// 				arg = t.Field(j).Name
// 			} else {
// 				arg = t.Field(j).Tag.Get(tagKey)
// 			}
// 			s.render.args = append(s.render.args, v.Field(j).Interface())
// 			// if v.Field(j).Kind() == reflect.String {
// 			// 	// fmt.Printf(3"t:%v      v:%+v", arg, v.Field(j).Interface().(string))
// 			// }
// 			if j > 0 {
// 				values += ","
// 			}
// 			values += "?"
// 			if i == 0 {
// 				if j > 0 {
// 					fields += ","
// 				}
// 				fields += arg
// 			}
// 		}
// 		if i == 0 {
// 			fields += ")"
// 		}
// 		values += ")"
// 	}
// 	var sql = fields + " VALUES" + values
// 	return sql
// }
func (s *SQLSegments) UpdateField(key string, val interface{}) *SQLSegments {
	if len(s.params) == 0 {
		s.params = append(s.params, make(map[string]interface{}, 0))
	}
	s.params[0][key] = val
	return s
}

//Update ..
// list := make(map[string]interface{}, 1)
// list["[+]Expires"] = 1
func (s *SQLSegments) Update(vals map[string]interface{}) *SQLSegments {
	//panic("Update method only one parameter is supported")
	if len(vals) < 1 {
		panic("Must be have values")
	}
	s.params = append(s.params, vals)
	return s
}

//UnsafeUpdate 可以没有where条件更新 ,Update 更新必须指定where条件才能更新否则panic
// func (s *SQLSegments) UnsafeUpdate(vals ...map[string]interface{}) *SQLSegments {
// 	if len(vals) > 1 {
// 		panic("Update method only one parameter is supported")
// 	}
// 	s.params = append(s.params, vals...)
// 	return s
// }

//buildReturning ...
func (s *SQLSegments) buildReturning() string {
	if s.returning == true {
		return " RETURNING"
	}
	return ""
}

//BuildUpdate ...
func (s *SQLSegments) BuildUpdate() string {
	var sql = fmt.Sprintf("UPDATE%s%s%s%s%s%s%s",
		s.buildFlags(),
		s.buildTable(),
		s.buildValuesForUpdate(),
		s.buildWhereClause(),
		s.buildOrderBy(),
		s.buildLimit(),
		s.buildReturning(),
	)
	// fmt.Println(s.render.args)
	return sql
}

//buildValuesForUpdate ...
func (s *SQLSegments) buildValuesForUpdate() string {
	var buffer bytes.Buffer
	buffer.WriteString(" SET ")
	// var fieldSlice []string
	if len(s.params) == 0 {
		panic(fmt.Sprintf("Must be have values after 'UPDATE %s SET'", s.buildTable()))
	}
	r, _ := regexp.Compile(`\[(\+|\-)\]?([a-zA-Z0-9_.\-\=\s\?\(\)]*)`)
	for i, vals := range s.params {
		if i == 0 {
			if len(vals) == 0 {
				panic(fmt.Sprintf("Must be have values after 'UPDATE %s SET'", s.buildTable()))
			}
			j := 0
			for arg, val := range vals {
				// fieldSlice = append(fieldSlice, arg)
				// s.render.args = append(s.render.args, val)
				if j > 0 {
					buffer.WriteString(", ")
				}

				match := r.FindStringSubmatch(arg)
				if len(match) > 1 {
					buffer.WriteString(buildIdent(match[2]))
					buffer.WriteString(" = ")
					buffer.WriteString(buildIdent(match[2]))
					buffer.WriteString(" ")
					buffer.WriteString(match[1])
					buffer.WriteString(" ?")
					s.render.args = append(s.render.args, val)
				} else {
					buffer.WriteString(buildIdent(arg))
					buffer.WriteString(" = ?")
					s.render.args = append(s.render.args, val)
				}
				j++
			}
		} else {
			if len(vals) == 0 {
				panic("just support one of vals")
			}
			//just support one of vals
			break
		}
	}
	// for i, s := range fieldSlice {
	// 	if i > 0 {
	// 		buffer.WriteString(", ")
	// 	}
	// 	buffer.WriteString(buildIdent(s))
	// 	buffer.WriteString(" = ?")
	// }
	return buffer.String()
}

// func (s *SQLSegments) buildValuesForUpdate() string {
// 	var sql = " SET"
// 	for i, param := range s.params {
// 		v := reflect.ValueOf(param).Elem()
// 		t := reflect.TypeOf(param).Elem()
// 		if i == 0 {
// 			for j := 0; j < v.NumField(); j++ {
// 				if v.Interface() == nil {
// 					continue
// 				}
// 				var arg string
// 				if t.Field(j).Tag.Get(tagKey) == "" {
// 					arg = t.Field(j).Name
// 				} else {
// 					arg = t.Field(j).Tag.Get(tagKey)
// 				}
// 				s.render.args = append(s.render.args, v.Field(j).Interface())
// 				if j > 0 {
// 					sql += ","
// 				} else {
// 					sql += " "
// 				}
// 				sql += arg + " = ?"
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	return sql
// }

//Delete ...
func (s *SQLSegments) Delete() *SQLSegments {
	return s
}

//BuildDelete ...
func (s *SQLSegments) BuildDelete() string {
	var sql = fmt.Sprintf("DELETE%s FROM%s%s%s%s%s",
		s.buildFlags(),
		s.buildTable(),
		s.buildWhereClause(),
		s.buildOrderBy(),
		s.buildLimit(),
		s.buildReturning(),
	)
	// fmt.Println(s.render.args)
	return sql
}

//buildIdent
func buildIdent(name string) string {
	return identKey + strings.Replace(name, ".", identKey+"."+identKey, -1) + identKey
}

func buildString(vals []string, sep string, header, footer string, ident bool) string {
	var buffer bytes.Buffer
	buffer.WriteString(header)
	for i, s := range vals {
		if i > 0 {
			buffer.WriteString(sep)
		}
		if ident {
			buffer.WriteString(buildIdent(s))
		} else {
			buffer.WriteString(s)
		}
	}
	buffer.WriteString(footer)
	return buffer.String()
}

//buildPlaceholder
func buildPlaceholder(l int, holder, sep string) string {
	var buffer bytes.Buffer
	for i := 0; i < l; i++ {
		if i > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(holder)
	}
	return buffer.String()
}

//Args ..
func (s *SQLSegments) Args() []interface{} {
	return s.render.args
}

/*
//Connect ...
type Connect struct {
	db *sql.DB
}

//NewConnect ...
func NewConnect() *Connect {
	return &Connect{}
}

//TT ...
type TT struct {
	ID string `db:"id"`
}

type mapingStruct struct {
	rt      reflect.Type
	mapping map[string]*reflect.StructField
}

func (ms *mapingStruct) FieldByMaping(name string) *reflect.StructField {
	if ptr, ok := ms.mapping[name]; ok {
		return ptr
	}
	return nil
}

func newMapingStruct(src interface{}) mapingStruct {
	rv := reflect.ValueOf(src)
	t := rv.Type().Elem()
	var ms mapingStruct
	ms.rt = t
	ms.mapping = make(map[string]*reflect.StructField, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !strings.Contains(string(field.Tag), tagKey+":") {
			ms.mapping[field.Name] = &field
			continue
		}
		tag := t.Field(i).Tag.Get(tagKey)
		if tag == "-" {
			continue
		}
		//todo db:"id,pk"
		ms.mapping[tag] = &field
	}
	return ms
}

func (c *Connect) scanAll(rows *sql.Rows, dest interface{}) {
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	rv := reflect.ValueOf(dest)
	fmt.Println("type:", rv.Kind())
	if rv.Kind() == reflect.Ptr {
		rv = reflect.Indirect(rv)
	} else if rv.Kind() == reflect.Slice {
		// slice := rv.Type()
		// sliceTpye := rv.Type().Elem()
		// ms := newMapingStruct(slice.Elem())
		// fmt.Println(sliceTpye)
		// fields := m.TraversalsByName(slice.Elem(), columns)
		for rows.Next() {
			// elem := reflect.New(sliceTpye).Elem()
			// vp := reflect.New(slice.Elem())
			// v := reflect.Indirect(vp)
			// fmt.Println("===", v.Kind())
			// fmt.Println(v.FieldByName("ID").Interface())
			// values[0] = elem.FieldByName("ID").Interface()
			// values[0] = elem.Field(0).Interface()
			// if v.Kind() == reflect.Struct {
			// for i := range columns {
			// 	ptr := ms.FieldByMaping("id")
			// 	if ptr != nil {
			// 		values[i] = v.FieldByIndex(ptr.Index).Interface()
			// 	} else {
			// 		fmt.Println(v.Field(i).Interface())
			// 		values[i] = v.Field(i).Interface()
			// 	}
			// }
			// values[i] = v.Field(i).Addr().Interface()
			rows.Scan(values...)
		}
		// for i := range columns {
		//
		// 	// values[i] = v.Field(i).Addr().Interface()
		// 	// rt := reflect.TypeOf(vp.Interface())
		// 	// values[i] = vp.Elem().Field(0).Addr().Interface()
		// }
	} else {

	}
	fmt.Println(dest)
}

//Connect ...
func (c *Connect) Connect() {
	var err error
	c.db, err = sql.Open("mysql", "root:123321@tcp(192.168.33.10:3306)/auth")

	if err != nil {
		log.Println(err)
	}
	defer c.db.Close()
	err = c.db.Ping()
	if err != nil {
		log.Println(err)
	}
	rows, err := c.db.Query("select * from accounts")
	if err != nil {
		fmt.Println(err)
	}
	// columns, _ := rows.Columns()
	// values := make([]interface{}, len(columns))
	// for rows.Next() {
	// 	for i := range columns {
	// 		values[i] = new(sql.RawBytes)
	// 	}
	// 	rows.Scan(values...)
	// }
	fmt.Println(rows)
	var dest []TT
	c.scanAll(rows, dest)

	// for i := range columns {
	// 	values[i] = rv.Elem().Field(i).Addr().Interface()
	// }
	// fmt.Println("type:", rv.Kind())
	// if rv.Kind() == reflect.Ptr {
	// 	values[0] = rv.Elem().Field(0).Addr().Interface()
	// } else if rv.Kind() == reflect.Slice {

	// } else {
	// 	rn := reflect.New(reflect.PtrTo(rv.Type()))

	// 	// fmt.Println(rn.Elem().Type())
	// 	// fmt.Println("----", rv.Kind() == reflect.Struct)
	// 	// rn.Elem().Set(rv)

	// 	// fmt.Println(rv.Addr())
	// 	// rv.Set(rn.Elem())
	// 	values[0] = rn.Elem().Addr().Interface()

	// 	// values[0] = rv.Field(0).Addr().Interface()
	// }
	// fmt.Printf("%+v %+v\r\n", values, dest)
	// fmt.Println("---", aa)
	//在这里进行一些数据库操作
}
*/
