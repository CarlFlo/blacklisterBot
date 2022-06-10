package database

import (
	"bytes"
	"encoding/json"

	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/corona10/goimagehash"
)

func SearchSHA1(h string) (bool, error) {

	var count int64
	if err := DB.Model(&Blacklist{}).Limit(1).Where("sha1 = ?", h).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func SearchAveragePerceptionDifference(aHash, dHash, pHash *goimagehash.ImageHash) (bool, error) {

	i := 0
	queryMax := 50

	for {

		var blacklist []Blacklist

		DB.Offset(i * queryMax).Limit(queryMax).Find(&blacklist)

		// Check each entry
		for _, e := range blacklist {

			var hash *goimagehash.ImageHash
			var distance int
			var err error
			var buf bytes.Buffer

			// Average
			createIOReader(&buf, e.getAverage(), int(goimagehash.AHash))

			hash, err = goimagehash.LoadImageHash(&buf)
			if err != nil {
				return false, err
			}
			distance, err = hash.Distance(aHash)
			if err != nil {
				return false, err
			}

			if distance <= config.CONFIG.Thresholds.Average {
				return true, nil
			}

			// Perception
			createIOReader(&buf, e.getDifference(), int(goimagehash.PHash))

			hash, err = goimagehash.LoadImageHash(&buf)
			if err != nil {
				return false, err
			}
			distance, err = hash.Distance(pHash)
			if err != nil {
				return false, err
			}

			if distance <= config.CONFIG.Thresholds.Perception {
				return true, nil
			}

			// Difference
			createIOReader(&buf, e.getPerception(), int(goimagehash.DHash))

			hash, err = goimagehash.LoadImageHash(&buf)
			if err != nil {
				return false, err
			}
			distance, err = hash.Distance(dHash)
			if err != nil {
				return false, err
			}

			if distance <= config.CONFIG.Thresholds.Difference {
				return true, nil
			}
		}

		// Got less than the limit, menaing it must be the last iteration
		if len(blacklist) < queryMax {
			break
		}

		i++
	}

	return false, nil
}

func createIOReader(buf *bytes.Buffer, hash uint64, kind int) error {

	type E struct {
		Hash uint64
		Kind int
	}

	return json.NewEncoder(buf).Encode(E{
		Hash: hash,
		Kind: kind,
	})
}
