package manager

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"

	"github.com/jace-ys/tusk/pkg/task"
)

var taskBucket = []byte("Tasks")

// TaskManager handles all CRUD operations with Bolt
type TaskManager struct {
	bolt *bolt.DB
}

func New(dbFilepath string) (*TaskManager, error) {
	// Open the `.db` file, otherwise create if it doesn't exist
	dbOptions := &bolt.Options{Timeout: 1 * time.Second}
	db, err := bolt.Open(dbFilepath, 0600, dbOptions)
	if err != nil {
		return nil, err
	}
	taskManager := &TaskManager{
		bolt: db,
	}
	// Create root bucket `task` to hold all tasks
	// db.Update returns an error
	createRootBkt := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
	return taskManager, createRootBkt
}

// Helper functions to encode/decode between struct and []byte
// Needed for storing structs in Bolt
func encode(t *task.Task) []byte {
	b, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	return b
}

func decode(b []byte) *task.Task {
	var t *task.Task
	err := json.Unmarshal(b, &t)
	if err != nil {
		return nil
	}
	return t
}

// Returns an 8-byte big endian representation of i
func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}
