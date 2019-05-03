package manager

import (
	"github.com/boltdb/bolt"

	"github.com/jace-ys/taskar/pkg/task"
)

func (tm *TaskManager) FindOne(key string) (*task.Task, error) {
	var t *task.Task
	err := tm.bolt.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(taskBucket)
		v := bkt.Get([]byte(key))
		t = decode(v)
		return nil
	})
	if err != nil {
		return t, err
	}
	return t, nil
}

func (tm *TaskManager) FindAll() (task.TaskSlice, error) {
	var t task.TaskSlice
	err := tm.bolt.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(taskBucket)
		c := bkt.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t = append(t, decode(v))
		}
		return nil
	})
	if err != nil {
		return t, err
	}
	return t, nil
}
