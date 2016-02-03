package schema

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"math"
	"strconv"

	google_protobuf "github.com/opsee/protobuf/proto/google/protobuf"
)

func coerceString(value interface{}) interface{} {
	switch value.(type) {
	case string:
		return value
	case []byte:
		return string(value.([]byte))
	}
	return fmt.Sprintf("%v", value)
}

// String is the GraphQL string type definition
var String *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "String",
	Serialize:  coerceString,
	ParseValue: coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coerceInt(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1
		}
		return 0
	case int:
		return value
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case int64:
		if value < int64(math.MinInt32) || value > int64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case uint:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint32:
		if value > uint32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case uint64:
		if value > uint64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case float32:
		if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case float64:
		if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceInt(val)
	case *google_protobuf.Timestamp:
		millis := value.Seconds * 1000
		millis = millis + int64(value.Nanos/1000000)
		return millis
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

// Int is the GraphQL Integer type definition.
var Int *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "Int",
	Serialize:  coerceInt,
	ParseValue: coerceInt,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})
