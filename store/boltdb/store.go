package boltdb

import (
	"time"

	"github.com/asdine/brazier"
	"github.com/asdine/brazier/store"
	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

// NewStore returns a BoltDB store
func NewStore(path string) (*Store, error) {
	var err error

	db, err := storm.Open(
		path,
		storm.AutoIncrement(),
		storm.Codec(new(rawCodec)),
		storm.BoltOptions(0644, &bolt.Options{
			Timeout: time.Duration(50) * time.Millisecond,
		}),
	)

	if err != nil {
		return nil, errors.Wrap(err, "Can't open database")
	}

	return &Store{
		DB: db,
	}, nil
}

// Store is a BoltDB store
type Store struct {
	DB *storm.DB
}

// Create a bucket
func (s *Store) Create(name string) error {
	_, err := s.Bucket(name)
	if err == nil {
		return store.ErrAlreadyExists
	}
	if err != store.ErrNotFound {
		return err
	}

	return s.DB.Set("buckets", name, nil)
}

// Bucket returns the bucket associated with the given id
func (s *Store) Bucket(name string) (brazier.Bucket, error) {
	var str []byte
	err := s.DB.Get("buckets", name, &str)
	if err != nil {
		if err == storm.ErrNotFound {
			return nil, store.ErrNotFound
		}
		return nil, err
	}

	return NewBucket(s, name, s.DB.From(name)), nil
}

// List returns the list of all buckets
func (s *Store) List() ([]string, error) {
	var buckets []string
	err := s.DB.Select().Bucket("buckets").RawEach(func(k, v []byte) error {
		buckets = append(buckets, string(k))
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "boltdb.store.List failed to fetch buckets")
	}

	return buckets, nil
}

// Close BoltDB connection
func (s *Store) Close() error {
	return s.DB.Close()
}
