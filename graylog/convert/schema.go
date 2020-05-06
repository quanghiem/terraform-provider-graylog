package convert

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func GetSchema(data interface{}, sc *schema.Schema) (interface{}, error) {
	// ResourceData => Graylog API (Post/Update)
	switch sc.Type {
	case schema.TypeList:
		if sc.MinItems == 1 && sc.MaxItems == 1 {
			return data.([]interface{})[0], nil
		}
	case schema.TypeSet:
		return data.(*schema.Set).List(), nil
	case schema.TypeMap:
	}
	return data, nil
}

func SetSchema(data interface{}, sc *schema.Schema) (interface{}, error) {
	// Graylog API (Get) => ResourceData
	switch sc.Type {
	case schema.TypeList:
		return SetTypeList(data, sc)
	case schema.TypeSet:
	case schema.TypeMap:
	}
	return data, nil
}

func SetTypeList(data interface{}, sc *schema.Schema) (interface{}, error) {
	switch t := sc.Elem.(type) {
	case *schema.Resource:
		switch v := data.(type) {
		case []interface{}:
			ret := make([]interface{}, len(v))
			for i, a := range v {
				b, err := SetResource(a.(map[string]interface{}), t)
				if err != nil {
					return nil, err
				}
				ret[i] = b
			}
			return ret, nil
		}
	case *schema.Schema:
	}
	if sc.MinItems == 1 && sc.MaxItems == 1 {
		return []interface{}{data}, nil
	}
	return data, nil
}