package plugin

import (
	"fmt"
	"plugin"
	"reflect"
)

var (
	ErrNotAPointer      = fmt.Errorf("not a pointer")
	ErrNotAStruct       = fmt.Errorf("not a struct")
	ErrUnassignableType = fmt.Errorf("unable to assign plugin type to struct type")
	ErrLookupFailed     = fmt.Errorf("plugin property lookup failed")
)

func Load(plg interface{}, pathToPlugin string) error {
	// load the file
	p, err := plugin.Open(pathToPlugin)
	if err != nil {
		return err
	}

	return fill(plg, p)
}

// validateStruct validates that the struct being filled in
// is in fact a struct pointer pointing to a valid space in memory
func validateStruct(plg interface{}) (reflect.Type, error) {
	ptr := reflect.TypeOf(plg)
	if ptr.Kind() != reflect.Ptr {
		return nil, ErrNotAPointer
	}

	st := ptr.Elem()

	if st.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	return st, nil
}

func fill(plg interface{}, pluginSource *plugin.Plugin) error {
	st, err := validateStruct(plg)
	if err != nil {
		return err
	}
	p := reflect.ValueOf(plg).Elem()

	for i := 0; i < st.NumField(); i++ {
		fld := st.Field(i)
		plgFld := p.Field(i)

		// lookup the field name in the plugin
		name := fld.Tag.Get("lookup")
		v, err := pluginSource.Lookup(name)
		if err != nil {
			return fmt.Errorf("%v : %v", name, ErrLookupFailed)
		}

		// reflect on what was found in the plugin
		vVal := reflect.ValueOf(v)
		vType := reflect.TypeOf(v)

		// see if we can assign it to the struct
		if !vType.AssignableTo(fld.Type) {
			return fmt.Errorf("%v => %v : %v", vType, fld.Type, ErrUnassignableType)
		}

		// actually assign it to the struct
		plgFld.Set(vVal)
	}

	return nil
}
