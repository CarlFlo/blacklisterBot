package database

import (
	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/malm"
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

	// TODO: fix unexpected EOF

	for {

		var blacklist []Blacklist

		DB.Offset(i * queryMax).Limit(queryMax).Find(&blacklist)

		// Check each entry
		for _, e := range blacklist {

			var hash *goimagehash.ImageHash
			var distance int
			var err error

			// Average
			hash = goimagehash.NewImageHash(e.getAverage(), goimagehash.AHash)
			distance, err = hash.Distance(aHash)
			if err != nil {
				return false, err
			}

			if distance <= config.CONFIG.Thresholds.Average {
				malm.Debug("Average distance: %d <= %d", distance, config.CONFIG.Thresholds.Average)
				return true, nil
			}

			// Perception
			hash = goimagehash.NewImageHash(e.getDifference(), goimagehash.PHash)
			distance, err = hash.Distance(pHash)
			if err != nil {
				return false, err
			}

			if distance <= config.CONFIG.Thresholds.Perception {
				malm.Debug("Perception distance: %d <= %d", distance, config.CONFIG.Thresholds.Perception)
				return true, nil
			}

			// Difference
			hash = goimagehash.NewImageHash(e.getPerception(), goimagehash.DHash)
			distance, err = hash.Distance(dHash)
			if err != nil {
				return false, err
			}

			if distance <= config.CONFIG.Thresholds.Difference {
				malm.Debug("Difference distance: %d <= %d", distance, config.CONFIG.Thresholds.Difference)
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
