package timeT

import (
	"database/sql/driver"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type timestamp time.Time

// MarshalJSON implements json.Marshaler.
func (t timestamp) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).UnixMilli())
	return []byte(stamp), nil
}

func (t *timestamp) UnmarshalJSON(data []byte) (err error) {
	var ts int64
	ts, err = strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	theTime := time.UnixMilli(ts)
	*t = timestamp(theTime)
	return nil
}

func (t timestamp) Value() (driver.Value, error) {
	if t.GetTime().UnixMilli() == 0 {
		return time.Time{}, nil
	}
	return time.Time(t), nil
}

func (t *timestamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = timestamp(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t timestamp) GetTime() time.Time {
	return time.Time(t)
}

func (t timestamp) GetBSON() (interface{}, error) {
	return time.Time(t).UnixMilli(), nil
}

func (t *timestamp) SetBSON(raw bson.Raw) error {
	var val int64
	if err := raw.Unmarshal(&val); err != nil {
		return err
	}
	*t = timestamp(time.UnixMilli(val))
	return nil
}

func (t timestamp) SpecifiedDate(days int, layout string) string {
	return t.GetTime().AddDate(0, 0, days).Format(layout)
}
