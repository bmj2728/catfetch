package catdb

import (
	"fmt"
	"hash/fnv"
	"log"
	"strings"
	"time"

	"github.com/bmj2728/catfetch/pkg/shared/metadata"
	"go.etcd.io/bbolt"
)

const (
	bucketCats         = "cats"
	bucketMetadata     = "metadata"
	bucketData         = "data"
	dbKeyImgData       = "img_data"
	dbKeyMetaId        = "cat_id"
	dbKeyMetaTags      = "tags"
	dbKeyMetaCreatedAt = "created_at"
	dbKeyMetaURL       = "url"
	dbKeyMetaMIMEType  = "mime_type"
)

type CatDB struct {
	db *bbolt.DB
}

func OpenDB(path string) (*CatDB, error) {
	db, err := bbolt.Open(path, 0644, nil)
	if err != nil {
		return nil, err
	}
	cdb := &CatDB{db: db}
	err = cdb.InitDB()
	if err != nil {
		return nil, err
	}

	return &CatDB{db: db}, nil
}

func (c *CatDB) Close() error {
	fmt.Println("Closing CatDB")
	fmt.Println(c.DB().Stats())

	return c.db.Close()
}

func (c *CatDB) DB() *bbolt.DB {
	return c.db
}

func (c *CatDB) InitDB() error {
	return c.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketCats))
		if err != nil {
			return err
		}
		return nil
	})
}

func (c *CatDB) AddCatVersion(metadata *metadata.CatMetadata, catData []byte) (string, string, error) {
	catId := metadata.ID
	versionId, err := hashURL(metadata.URL)
	if err != nil {
		return "", "", err
	}

	err = c.db.Update(func(tx *bbolt.Tx) error {

		b := tx.Bucket([]byte(bucketCats))
		cat, err := b.CreateBucketIfNotExists([]byte(catId))
		if err != nil {
			return err
		}
		version, err := cat.CreateBucketIfNotExists([]byte(versionId))
		if err != nil {
			return err
		}
		md, err := version.CreateBucketIfNotExists([]byte(bucketMetadata))
		if err != nil {
			return err
		}

		err = md.Put([]byte(dbKeyMetaId), []byte(metadata.ID))
		if err != nil {
			return err
		}

		err = md.Put([]byte(dbKeyMetaTags), []byte(strings.Join(metadata.Tags, ", ")))
		if err != nil {
			return err
		}

		err = md.Put([]byte(dbKeyMetaCreatedAt), []byte(metadata.CreatedAt.Format(time.RFC3339)))
		if err != nil {
			return err
		}

		err = md.Put([]byte(dbKeyMetaURL), []byte(metadata.URL))
		if err != nil {
			return err
		}

		err = md.Put([]byte(dbKeyMetaMIMEType), []byte(metadata.MIMEType))
		if err != nil {
			return err
		}

		data, err := version.CreateBucketIfNotExists([]byte(bucketData))
		if err != nil {
			return err
		}
		err = data.Put([]byte(dbKeyImgData), catData)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", "", err
	}
	return catId, versionId, nil
}

func hashURL(url string) (string, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(url))
	if err != nil {
		log.Default().Println(err)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum64()), nil
}
