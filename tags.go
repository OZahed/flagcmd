package flagcmd

import (
	"flag"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func setupFlagsFromStructValue(fs *flag.FlagSet, sv reflect.Value) error {
	sv = reflect.Indirect(sv)
	st := sv.Type()

	for i := 0; i < sv.NumField(); i++ {
		f := sv.Field(i)
		fStruct := st.Field(i)

		// is Exported Field
		if fStruct.PkgPath != "" {
			continue
		}

		tag := fStruct.Tag.Get("cli")
		if tag == "" || tag == "-" {
			continue
		}

		// Parse the tag value
		tagVal := strings.Split(tag, ",")
		name := tagVal[0]
		var def string
		var usage string
		if len(tagVal) > 0 {
			def = tagVal[1]
		}
		if len(tagVal) > 1 {
			usage = tagVal[2]
		}

		switch f.Kind() {
		case reflect.String:
			fs.StringVar((*string)(unsafe.Pointer(f.Addr().Pointer())), name, f.String(), usage)
		case reflect.Bool:
			b, err := strconv.ParseBool(def)
			if def != "" && err != nil {
				return err
			}
			fs.BoolVar((*bool)(unsafe.Pointer(f.Addr().Pointer())), name, b, usage)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(def, 10, 64)
			if def != "" && err != nil {
				return err
			}
			fs.Int64Var((*int64)(unsafe.Pointer(f.Addr().Pointer())), name, n, usage)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			n, err := strconv.ParseUint(def, 10, 64)
			if def != "" && err != nil {
				return err
			}
			fs.Uint64Var((*uint64)(unsafe.Pointer(f.Addr().Pointer())), name, n, usage)
		case reflect.Float32, reflect.Float64:
			n, err := strconv.ParseFloat(def, 64)
			if def != "" && err != nil {
				return err
			}
			fs.Float64Var((*float64)(unsafe.Pointer(f.Addr().Pointer())), name, n, usage)
		case reflect.Struct:
			switch f.Type() {
			case reflect.TypeOf(time.Time{}):
				fs.StringVar((*string)(unsafe.Pointer(f.Addr().Pointer())), name, def, usage)
			case reflect.TypeOf(url.URL{}):
				fs.StringVar((*string)(unsafe.Pointer(f.Addr().Pointer())), name, def, usage)
			default:
				if err := setupFlagsFromStructValue(fs, f.Elem()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func ParseFlagSet(i interface{}) error {
	if i == nil {
		return ErrNilPointer
	}
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Pointer {
		return ErrNoPtr
	}

	fs := flag.NewFlagSet(v.Elem().Type().Name(), flag.ExitOnError)

	if err := setupFlagsFromStructValue(fs, v); err != nil {
		return err
	}

	return fs.Parse(os.Args)
}
