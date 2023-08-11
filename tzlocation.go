package dayone2md

import (
	"log/slog"
	"time"
)

func makeTimeZoneCache() func(string) *time.Location {
	tzLocations := make(map[string]*time.Location)
	return func(tz string) *time.Location {
		var err error
		tzLocation, ok := tzLocations[tz]
		if !ok {
			tzLocation, err = time.LoadLocation(tz)
			if err != nil {
				slog.Warn("timezone not found", "tz", tz, "err", err)
				tzLocation = time.Local // use local
			}
			tzLocations[tz] = tzLocation
		}
		return tzLocation
	}
}
