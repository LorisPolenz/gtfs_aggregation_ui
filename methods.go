package main

import (
	"database/sql"
	"fmt"
)

func fetch_by_routeShortName(route_short_name string) []Entity {
	rows, err := get_db().Query("SELECT * FROM p1 WHERE route_short_name = $1", route_short_name)
	check(err)
	defer rows.Close()

	entities := []Entity{}

	for rows.Next() {
		var entity Entity
		var arrivalDelay, departureDelay sql.NullInt64
		var platform sql.NullString

		err = rows.Scan(
			&entity.StopName,
			&entity.StopSequence,
			&entity.StartDatetime,
			&platform,
			&arrivalDelay,
			&departureDelay,
			&entity.RouteShortName,
			&entity.TripHeadsign)

		// Set to 0 if nil
		entity.ArrivalDelay = 0
		if arrivalDelay.Valid {
			entity.ArrivalDelay = int(arrivalDelay.Int64)
		}

		entity.DepartureDelay = 0
		if departureDelay.Valid {
			entity.DepartureDelay = int(departureDelay.Int64)
		}

		entity.Platform = ""
		if platform.Valid {
			entity.Platform = platform.String
		}

		check(err)
		entities = append(entities, entity)
	}

	return entities
}

func fetch_by_query(query Query) []Entity {
	conditions := []string{}
	args := []interface{}{}
	argIndex := 1
	baseQuery := "SELECT * FROM p1"

	if query.StopName != "" {
		conditions = append(conditions, fmt.Sprintf("stop_name = $%d", argIndex))
		args = append(args, query.StopName)
		argIndex++
	}

	if query.RouteShortname != "" {
		conditions = append(conditions, fmt.Sprintf("route_short_name = $%d", argIndex))
		args = append(args, query.RouteShortname)
		argIndex++
	}

	if query.TripHeadsign != "" {
		conditions = append(conditions, fmt.Sprintf("trip_headsign = $%d", argIndex))
		args = append(args, query.TripHeadsign)
		argIndex++
	}

	if query.DelayThreshold > 0 {
		conditions = append(conditions, fmt.Sprintf("arrival_delay > $%d", argIndex))
		args = append(args, query.DelayThreshold)
		argIndex++
	}

	// Append conditions to the base query if any
	if len(conditions) > 0 {
		baseQuery += " WHERE " + conditions[0]
		for _, cond := range conditions[1:] {
			baseQuery += " AND " + cond
		}
	}

	rows, err := get_db().Query(baseQuery, args...)

	check(err)
	defer rows.Close()

	entities := []Entity{}

	for rows.Next() {
		var entity Entity
		var arrivalDelay, departureDelay sql.NullInt64
		var platform sql.NullString

		err = rows.Scan(
			&entity.StopName,
			&entity.StopSequence,
			&entity.StartDatetime,
			&platform,
			&arrivalDelay,
			&departureDelay,
			&entity.RouteShortName,
			&entity.TripHeadsign)

		// Set to 0 if nil
		entity.ArrivalDelay = 0
		if arrivalDelay.Valid {
			entity.ArrivalDelay = int(arrivalDelay.Int64)
		}

		entity.DepartureDelay = 0
		if departureDelay.Valid {
			entity.DepartureDelay = int(departureDelay.Int64)
		}

		entity.Platform = ""
		if platform.Valid {
			entity.Platform = platform.String
		}

		check(err)
		entities = append(entities, entity)
	}

	return entities
}
