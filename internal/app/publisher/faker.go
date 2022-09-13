package publisher

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v4"
)

func customFakerGenerator() error {
	// Custom timestamp generator in correct format
	if err := faker.AddProvider("customFakerTimestamp", func(v reflect.Value) (interface{}, error) {
		randTimestamp := time.Unix(rand.Int63n(time.Now().Unix()), 0)
		return randTimestamp.Format("2006-01-02T15:04:05Z"), nil
	}); err != nil {
		return err
	}

	if err := faker.SetRandomMapAndSliceSize(5); err != nil {
		return err
	}
	return nil
}
