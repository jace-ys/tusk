package manager

import (
	"github.com/boltdb/bolt"

	"github.com/jace-ys/taskar/pkg/task"
)

func (tm *TaskManager) Update(t *task.Task) error {
	err := tm.bolt.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(taskBucket)
		return bkt.Put([]byte(t.ID), encode(t))
	})
	if err != nil {
		return err
	}
	return nil
}
