package manager

import (
	"strconv"

	"github.com/boltdb/bolt"

	"github.com/jace-ys/taskar/pkg/task"
)

func (tm *TaskManager) Create(t *task.Task) error {
	err := tm.bolt.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(taskBucket)
		// Get subsequent ID for bucket and convert it from int64 to string
		id, _ := bkt.NextSequence()
		t.ID = strconv.Itoa(int(id))
		return bkt.Put([]byte(t.ID), encode(t))
	})
	if err != nil {
		return err
	}
	return nil
}
