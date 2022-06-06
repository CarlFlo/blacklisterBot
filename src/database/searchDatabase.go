package database

import "github.com/corona10/goimagehash"

func SearchSHA1(h string) bool {

	var b Blacklist

	DB.Where("sha1 = ?", h).First(b)

	// Check if anything returned

	return false
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
