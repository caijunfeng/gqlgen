package exec

import (
	"fmt"
	"math"
	"reflect"

	"github.com/neelance/graphql-go/internal/schema"
)

type scalar struct {
	name        string
	reflectType reflect.Type
	coerceInput func(input interface{}) (interface{}, error)
}

func (*scalar) Kind() string       { return "SCALAR" }
func (t *scalar) TypeName() string { return t.name }

var builtinScalars = []*scalar{
	&scalar{
		name:        "Int",
		reflectType: reflect.TypeOf(int32(0)),
		coerceInput: func(input interface{}) (interface{}, error) {
			i := input.(int)
			if i < math.MinInt32 || i > math.MaxInt32 {
				return nil, fmt.Errorf("not a 32-bit integer: %d", i)
			}
			return int32(i), nil
		},
	},
	&scalar{
		name:        "Float",
		reflectType: reflect.TypeOf(float64(0)),
		coerceInput: func(input interface{}) (interface{}, error) {
			return input, nil // TODO
		},
	},
	&scalar{
		name:        "String",
		reflectType: reflect.TypeOf(""),
		coerceInput: func(input interface{}) (interface{}, error) {
			return input, nil // TODO
		},
	},
	&scalar{
		name:        "Boolean",
		reflectType: reflect.TypeOf(true),
		coerceInput: func(input interface{}) (interface{}, error) {
			return input, nil // TODO
		},
	},
	&scalar{
		name:        "ID",
		reflectType: reflect.TypeOf(""),
		coerceInput: func(input interface{}) (interface{}, error) {
			return input, nil // TODO
		},
	},
}

func AddBuiltinScalars(s *schema.Schema) {
	for _, scalar := range builtinScalars {
		s.Types[scalar.name] = scalar
	}
}