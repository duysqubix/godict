package dict

import (
	"errors"
	"fmt"

	pp "github.com/k0kubun/pp/v3"
)

type dict map[interface{}]interface{}

type keyPair struct {
	key, value interface{}
}

type RawMap dict

func Dict(vals ...RawMap) *dict {
	t := make(dict)

	if len(vals) > 0 {
		for _, val := range vals {
			for k, v := range val {
				t[k] = v
			}
		}
	}
	return &t
}

func (d *dict) String() string {

	mypp := pp.New()
	mypp.SetColoringEnabled(false)
	return mypp.Sprint(d)

}

func (d *dict) Update(obj ...interface{}) error {
	if len(obj) == 0 {
		err_msg := "Must supply Update method with a value"
		return errors.New(err_msg)
	}

	main_obj := obj[0]
	switch main_obj.(type) {
	case *dict:
		dd, _ := main_obj.(*dict)
		for k, v := range *dd {
			(*d)[k] = v
		}

	case dict:
		dd, _ := main_obj.(dict)
		for k, v := range dd {
			(*d)[k] = v
		}

	default:
		// at this point the only other option valid is key,value
		if len(obj) != 2 {
			err_msg := "Must provide key value pair with Update method"
			panic(err_msg)
		}
		key, value := obj[0], obj[1]
		(*d)[key] = value
	}
	return nil
}

func (d *dict) Clear() {
	for k := range *d {
		delete(*d, k)
	}

}

func (d *dict) Copy() *dict {
	cpy_dict := Dict()
	for k, v := range *d {
		cpy_dict.Update(k, v)
	}
	return cpy_dict
}

func (d *dict) Get(key interface{}) (*interface{}, error) {
	err_msg := fmt.Sprintf("no such key: %s", key)

	value, exists := (*d)[key]
	if !exists {
		return nil, errors.New(err_msg)
	}
	return &value, nil
}

func (d *dict) Items() []keyPair {
	keyPairs := make([]keyPair, len(*d))

	i := 0
	for k, v := range *d {
		kp := keyPair{key: k, value: v}
		keyPairs[i] = kp
		i++
	}

	return keyPairs
}
