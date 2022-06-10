package database

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/jpeg"
	"strconv"

	"github.com/CarlFlo/malm"
	"github.com/corona10/goimagehash"
)

type Blacklist struct {
	Model
	Sha1           string `gorm:"unique;not null;index"`
	AverageHash    string `gorm:"not null;index"`
	DifferenceHash string `gorm:"not null;index"`
	PerceptionHash string `gorm:"not null;index"`
	URL            string `gorm:"unique;not null"`
}

func (Blacklist) TableName() string {
	return "Blacklist"
}

// Saves the data to the database
func (b *Blacklist) Save() {
	tx := DB.Save(&b)
	if tx.Error != nil {
		malm.Error("%s", tx.Error)
	}
}

func (b *Blacklist) DeleteEntry() {
	DB.Delete(&b)
}

func (b *Blacklist) getAverage() uint64 {

	// ParseInt uint64
	i, err := strconv.ParseUint(b.AverageHash, 10, 64)
	if err != nil {
		malm.Error("%s", err)
	}
	return i
}

func (b *Blacklist) getDifference() uint64 {

	// ParseInt uint64
	i, err := strconv.ParseUint(b.DifferenceHash, 10, 64)
	if err != nil {
		malm.Error("%s", err)
	}
	return i
}

func (b *Blacklist) getPerception() uint64 {

	// ParseInt uint64
	i, err := strconv.ParseUint(b.PerceptionHash, 10, 64)
	if err != nil {
		malm.Error("%s", err)
	}
	return i
}

func (b *Blacklist) New(img image.Image, link string) error {

	var hash *goimagehash.ImageHash
	var err error

	b.URL = link

	// SHA1
	buf := new(bytes.Buffer)
	if err = jpeg.Encode(buf, img, nil); err != nil {
		return err
	}

	hasher := sha1.New()
	hasher.Write(buf.Bytes())
	b.Sha1 = fmt.Sprintf("%x", hasher.Sum(nil))

	// Average
	hash, err = goimagehash.AverageHash(img)
	if err != nil {
		return err
	}
	b.AverageHash = strconv.FormatUint(hash.GetHash(), 10)

	// Difference
	hash, err = goimagehash.DifferenceHash(img)
	if err != nil {
		return err
	}
	b.DifferenceHash = strconv.FormatUint(hash.GetHash(), 10)

	// Perception
	hash, err = goimagehash.PerceptionHash(img)
	if err != nil {
		return err
	}
	b.PerceptionHash = strconv.FormatUint(hash.GetHash(), 10)

	return nil
}
