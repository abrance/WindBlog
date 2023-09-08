package json_storage

import (
	"encoding/json"
	"github.com/WindBlog/util/errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/wonderivan/logger"
	"strings"
)

// 定义表时要像 Table 一样定一个结构体和方法

type Table struct {
	name string
}

func (t *Table) getDBEngine() *badger.DB {
	return db
}

func (t *Table) SetTableName(name string) *Table {
	t.name = name
	return t
}

func (t *Table) Get(id string) (TemplateStruct, error) {
	// 这里假设 obj 是 map[string]string 类型, 实际上可以为任意 go 结构体对象
	var objValue TemplateStruct
	err := t.getDBEngine().View(func(txn *badger.Txn) error {
		var err error
		item, err := txn.Get([]byte(t.name + id))
		if err != nil {
			logger.Error(errors.JsonDBError)
			return err
		}
		value, err := item.ValueCopy(nil)
		if err != nil {
			logger.Error(errors.JsonDBError)
			return err
		}
		err = json.Unmarshal(value, &objValue)
		if err != nil {
			logger.Error(errors.JsonDBError)
			return err
		}
		return nil
	})
	if err != nil {
		return objValue, err
	}
	return objValue, nil
}

type LikeOption struct {
	Like string
}

type FilterOption struct {
	NameFilterOption LikeOption
}

type TemplateStruct struct {
	TemplateField string
}

func (t *Table) List(filter *FilterOption) (map[string]TemplateStruct, error) {
	var mapObjValue map[string]TemplateStruct
	err := t.getDBEngine().View(func(txn *badger.Txn) error {
		var err error
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()
		var prefix []byte
		if filter != nil {
			// todo 判断有问题
			prefix = []byte(strings.Join([]string{t.name, filter.NameFilterOption.Like}, ":"))
		} else {
			prefix = []byte(t.name + ":")
		}

		for iter.Seek(prefix); iter.ValidForPrefix(prefix); iter.Next() {
			var objValue TemplateStruct
			var encoded []byte
			key := string(iter.Item().Key())
			_, err = iter.Item().ValueCopy(encoded)
			if err != nil {
				logger.Error(errors.JsonDBError, err)
				continue
			}
			err = json.Unmarshal(encoded, &objValue)
			if err != nil {
				logger.Error(errors.JsonDBError, err)
			}
			mapObjValue[key] = objValue
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	logger.Debug("map obj value: %v", mapObjValue)
	return mapObjValue, nil
}

func (t *Table) Insert(id string, fieldObj TemplateStruct) error {
	return t.getDBEngine().Update(func(txn *badger.Txn) error {
		encoded, err := json.Marshal(fieldObj)
		if err != nil {
			logger.Error(errors.JsonInsertError)
			return err
		}
		return txn.Set([]byte(strings.Join([]string{t.name, id}, ":")), encoded)
	})
}

func (t *Table) Update(id string, NewFieldObj TemplateStruct) error {
	return t.getDBEngine().Update(func(txn *badger.Txn) error {
		encoded, err := json.Marshal(NewFieldObj)
		if err != nil {
			logger.Error(errors.JsonInsertError)
			return err
		}
		return txn.Set([]byte(strings.Join([]string{t.name, id}, ":")), encoded)
	})
}

func (t *Table) Delete(id string) error {
	return t.getDBEngine().Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(strings.Join([]string{t.name, id}, ":")))
	})
}
