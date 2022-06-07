package database

import (
	"github.com/corona10/goimagehash"
)

func SearchSHA1(h string) (bool, error) {

	var count int64
	if err := DB.Model(&Blacklist{}).Limit(1).Where("sha1 = ?", h).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func SearchAverage(h *goimagehash.ImageHash) bool {

	return false
}

func SearchDifference(h *goimagehash.ImageHash) bool {

	return false
}

func SearchPerception(h *goimagehash.ImageHash) bool {

	return false
}
