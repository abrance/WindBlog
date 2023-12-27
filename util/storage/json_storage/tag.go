package json_storage

import (
	"encoding/json"
	"fmt"
	"github.com/WindBlog/util/errors"
	"github.com/coreos/etcd/pkg/idutil"
	"github.com/dgraph-io/badger/v3"
	"github.com/wonderivan/logger"
	"strings"
	"time"
)

type TagTable struct {
	name string
	gen  *idutil.Generator
}

type Tag struct {
	Id         string // unique key, 数字整型
	Name       string // 书名
	Url        string // 地址, file://  表示本地
	IsArchive  bool   // 是否已归档
	ArchiveId  string //归档id
	CreateTime int64
	UpdateTime int64
}

func (f *TagTable) Init() {
	f.SetTableName("tag")
	f.gen = idutil.NewGenerator(0, time.Now())
}

func (f *TagTable) getNextId() string {
	return fmt.Sprintf("%d", f.gen.Next())
}

func (f *TagTable) getDBEngine() *badger.DB {
	return db
}

func (f *TagTable) SetTableName(name string) *TagTable {
	f.name = name
	return f
}

func (f *TagTable) Get(id string) (*Tag, error) {
	// 这里假设 obj 是 map[string]string 类型, 实际上可以为任意 go 结构体对象
	var objValue Tag

	logger.Info(id)
	err := f.getDBEngine().View(func(txn *badger.Txn) error {
		var err error
		item, err := txn.Get([]byte(strings.Join([]string{f.name, id}, ":")))
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
		return &objValue, err
	}
	return &objValue, nil
}

func (f *TagTable) List(filter *FilterOption) (map[string]*Tag, error) {
	mapObjValue := make(map[string]*Tag)
	err := f.getDBEngine().View(func(txn *badger.Txn) error {
		//var err error
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()
		prefix := make([]byte, 0)
		if filter != nil {
			// todo 判断有问题
			prefix = []byte(strings.Join([]string{f.name, filter.NameFilterOption.Like}, ":"))
		} else {
			prefix = []byte(f.name + ":")
		}

		for iter.Seek(prefix); iter.ValidForPrefix(prefix); iter.Next() {
			var objValue Tag
			key := string(iter.Item().Key())
			encoded, err := iter.Item().ValueCopy(nil)
			if err != nil {
				logger.Error(errors.JsonDBError, err)
				continue
			}
			err = json.Unmarshal(encoded, &objValue)
			if err != nil {
				logger.Error(errors.JsonDBError, err)
				continue
			}
			mapObjValue[key] = &objValue
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	logger.Debug("map obj value: %v", mapObjValue)
	return mapObjValue, nil
}

func (f *TagTable) Insert(fieldObj *Tag) error {
	id := f.getNextId()
	fieldObj.Id = id
	return f.getDBEngine().Update(func(txn *badger.Txn) error {
		encoded, err := json.Marshal(*fieldObj)
		if err != nil {
			logger.Error(errors.JsonInsertError)
			return err
		}
		return txn.Set([]byte(strings.Join([]string{f.name, id}, ":")), encoded)
	})
}

func (f *TagTable) Update(id string, NewFieldObj *Tag) error {
	return f.getDBEngine().Update(func(txn *badger.Txn) error {
		encoded, err := json.Marshal(NewFieldObj)
		if err != nil {
			logger.Error(errors.JsonInsertError)
			return err
		}
		return txn.Set([]byte(strings.Join([]string{f.name, id}, ":")), encoded)
	})
}

func (f *TagTable) Delete(id string) error {
	return f.getDBEngine().Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(strings.Join([]string{f.name, id}, ":")))
	})
}
