package semver

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

func (r Version) Value() (driver.Value, error) {
	return r.String(), nil
}

func (r *Version) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		vv := loadVersion(string(v))
		*r = *vv
	case string:
		vv := loadVersion(string(v))
		*r = *vv
	default:
		return fmt.Errorf("not supported %s", reflect.TypeOf(src).Kind())
	}
	return nil
}
