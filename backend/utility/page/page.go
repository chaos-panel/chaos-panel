package page

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/chaos-plus/chaos-plus/utility/utils"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var SQL_OPERAORS = map[string]string{
	"eq":         "=",
	"neq":        "!=",
	"gt":         ">",
	"lt":         "<",
	"gte":        ">=",
	"lte":        "<=",
	"like":       "LIKE",
	"notLike":    "NOT LIKE",
	"in":         "IN",
	"notIn":      "NOT IN",
	"between":    "BETWEEN",
	"notBetween": "NOT BETWEEN",
	"null":       "IS NULL",
	"notNull":    "IS NOT NULL",
	// 可根据需要扩展更多操作符
}

type PageReq[F interface{}, O interface{}] struct {
	Offset int `json:"offset"  dc:"query offset"`
	Limit  int `json:"limit"   dc:"query limit"`
	Filter F   `json:"filter"  dc:"query filter map"`
	Order  O   `json:"order"   dc:"query order map"`
}
type PageRes[T interface{}] struct {
	Offset int `json:"offset"  dc:"query offset"`
	Limit  int `json:"limit"   dc:"query limit"`
	Count  int `json:"count"   dc:"list count"`
	Total  int `json:"total"   dc:"total count"`
	List   T   `json:"list"    dc:"data array"`
}

// FilterCondition 定义了一个过滤条件
type FilterCondition struct {
	FieldName   string
	FieldType   reflect.Type
	RawOperator string
	RawValue    interface{}
	SqlField    string
	SqlOperator string
	SqlValue    interface{}
}

// OrderCondition 定义了一个排序条件
type OrderCondition struct {
	FieldName string
	Direction string
}

func NonDeleted(m *gdb.Model) *gdb.Model {
	return m.Where("deleted_at <= ?", "1970-01-01 00:00:00")
}
func NonLocked(m *gdb.Model) *gdb.Model {
	return m.Where("locked_at <= ?", "1970-01-01 00:00:00")
}

func ParseQueriesFromCtx(c context.Context, m *gdb.Model, pageReq any) *gdb.Model {
	return ParseQueries(ghttp.RequestFromCtx(c), m, pageReq)
}

func ParseQueries(r *ghttp.Request, m *gdb.Model, pageReq any) *gdb.Model {
	if r == nil || m == nil || pageReq == nil {
		return m
	}
	val := utils.ReflectStruct(pageReq)
	if val.Kind() != reflect.Struct {
		return m
	}

	var filterMap map[string]FilterCondition
	var orderMap []OrderCondition

	var offset *gvar.Var = gvar.New(nil, false)
	var limit *gvar.Var = gvar.New(nil, false)

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldValue := val.Field(i).Interface()
		// fmt.Printf("%s: %v\n", fieldName, fieldValue)
		if strings.EqualFold(fieldName, "pageReq") {
			m = ParseQueries(r, m, fieldValue)
		}
		if strings.EqualFold(fieldName, "offset") {
			offset = gvar.New(fieldValue, true)
		}
		if strings.EqualFold(fieldName, "limit") {
			limit = gvar.New(fieldValue, true)
		}
		if strings.EqualFold(fieldName, "filter") {
			filterMap = ParseFilter(r, &fieldValue)
		}
		if strings.EqualFold(fieldName, "order") {
			orderMap = ParseOrer(r, &fieldValue)
		}
	}
	// fmt.Println("===> ",offset, limit, filterMap, orderMap)

	for f, c := range filterMap {
		_ = f
		// fmt.Println("filter ===> ", f, c)
		m = m.Where(c.SqlField+" "+c.SqlOperator+" ? ", c.SqlValue)
	}
	for i := range orderMap {
		c := orderMap[i]
		// fmt.Println("order ===> ", i, c)
		m = m.Order(c.FieldName + " " + c.Direction)
	}

	if offset != nil && !offset.IsNil() && limit != nil && !limit.IsNil() {
		m = m.Limit(offset.Int(), limit.Int())
	} else if limit != nil && !limit.IsNil() {
		m = m.Limit(limit.Int())
	}

	return m
}

func ParseFilter(r *ghttp.Request, f *interface{}) map[string]FilterCondition {
	filterMap := make(map[string]FilterCondition)

	filterParams := r.GetMap()["filter"]
	if filterParams == nil {
		return filterMap
	}
	filterMapper := (filterParams).(map[string]interface{})
	//
	resType := utils.ReflectStruct(f)
	for i := 0; i < resType.NumField(); i++ {
		field := resType.Type().Field(i)
		paramName := field.Tag.Get("p")
		if paramName == "" {
			paramName = field.Name
		}
		fieldName := field.Tag.Get("field")
		if fieldName == "" {
			fieldName = field.Name
		}
		filterValue := filterMapper[paramName]
		// fmt.Println("=======> ", fieldName, filterValue)
		if filterValue == nil {
			continue
		}
		tableName := field.Tag.Get("table")
		if tableName != "" {
			tableName = tableName + "."
		}
		operator, value := parseFiterValue(gconv.String(filterValue))
		condition := &FilterCondition{
			FieldName:   paramName,
			FieldType:   field.Type,
			RawOperator: operator,
			RawValue:    value,
			SqlField:    tableName + utils.CamelToSnake(fieldName),
			SqlOperator: SQL_OPERAORS[operator],
			SqlValue:    parseSqlValue(field.Type, value, operator),
		}

		filterMap[fieldName] = *condition
	}
	return filterMap
}

func ParseOrer(r *ghttp.Request, o *interface{}) []OrderCondition {
	orderMap := []OrderCondition{}

	orderParams := r.GetMap()["order"]
	if orderParams == nil {
		return orderMap
	}
	orderMapper := (orderParams).(map[string]interface{})

	resType := utils.ReflectStruct(o)
	for i := 0; i < resType.NumField(); i++ {
		field := resType.Type().Field(i)
		paramName := field.Tag.Get("p")
		if paramName == "" {
			paramName = field.Name
		}
		fieldName := field.Tag.Get("field")
		if fieldName == "" {
			fieldName = field.Name
		}
		filterValue := orderMapper[paramName]
		if filterValue == nil {
			continue
		}
		tableName := field.Tag.Get("table")
		if tableName != "" {
			tableName = tableName + "."
		}
		direction := gconv.String(filterValue)

		allowedValues := []string{"ASC", "DESC"}
		isAllowed := false
		for _, v := range allowedValues {
			if direction == v {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			continue
		}

		condition := OrderCondition{
			FieldName: tableName + fieldName,
			Direction: direction,
		}

		orderMap = append(orderMap, condition)
	}
	return orderMap
}

// AddFilter adds a filter to the PageRequestMap.
func parseSqlValue(valueType reflect.Type, value string, operator string) interface{} {
	switch strings.ToLower(operator) {
	case "like", "not like", "notLike":
		return "'%" + escapeSQL(value) + "%'"
	case "between", "not between", "notBetween":
		return strings.Join(getEscapesArray(valueType, value), " and ")
	case "in", "not in", "notIn":
		escapes := strings.Join(getEscapesArray(valueType, value), ",")
		if len(escapes) > 0 {
			return "(" + escapes + ")"
		} else {
			return "( null )"
		}
	default:
		return getTypedValue(valueType, value)
	}

}

func parseFiterValue(val string) (string, string) {
	if !strings.Contains(val, ":") {
		return "eq", val
	}
	split := strings.Split(val, ":")
	if len(split) <= 0 {
		return "eq", val
	} else if len(split) == 1 {
		return "eq", split[0]
	} else {
		if split[0] == "" {
			return "eq", split[1]
		} else {
			return split[0], split[1]
		}
	}
}

// getEscapesArray processes and escapes a JSON array string into a slice of strings.
func getEscapesArray(valueType reflect.Type, value string) []string {
	if !strings.HasPrefix(value, "[") {
		value = "[" + value
	}
	if !strings.HasSuffix(value, "]") {
		value = value + "]"
	}
	var arr []interface{}
	if err := json.Unmarshal([]byte(value), &arr); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil
	}
	var escapedValues []string
	for _, v := range arr {
		escapedValues = append(escapedValues, escapeSQL(getTypedValue(valueType, v).(string)))
	}
	return escapedValues
}

// getTypedValue converts the value into the appropriate type based on the provided reflect.Type.
func getTypedValue(valueType reflect.Type, value interface{}) interface{} {
	if value == nil {
		return nil
	}
	switch valueType.Kind() {
	case reflect.String:
		return "'" + escapeSQL(value.(string)) + "'"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value
	case reflect.Float32, reflect.Float64:
		return value
	case reflect.Bool:
		return value
	default:
		// Handle enums and other types as needed
		if valueType.Kind() == reflect.Struct {
			if _, ok := value.(time.Time); ok {
				return "'" + value.(time.Time).Format("2006-01-02 15:04:05") + "'"
			}
		}
	}
	return value
}

// escapeSQL escapes special characters in a string for use in SQL queries.
func escapeSQL(str string) string {
	replacer := strings.NewReplacer(`'`, `''`, `\`, `\\`)
	return replacer.Replace(str)
}
