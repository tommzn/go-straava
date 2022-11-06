package strava

import (
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// BASE_URL for Strava API v3.
const BASE_URL = "https://www.strava.com/api/v3"

// TOKEN_ENDPOINT is used to convert authorization code to token or retrieve a new access token.
const TOKEN_ENDPOINT = "https://www.strava.com/oauth/token"

// StravaClient handles request to Strava APIs.
type Client struct {

	// BaseUrl for Strave API. If not set BASE_URL will be used.
	baseUrl string

	// Id of an athlete. If not set "/athlete" endpoint will be used to determine current authenticates athlete.
	athleteId *int64

	// TokenSource provides token for API authentication.
	tokenSource oauth2.TokenSource

	// Http client to call Strave API endpoints.
	httpClient *http.Client
}

// TimeFilter to restrict api requests with before and after values.
type TimeFilter struct {
	Before *time.Time
	After  *time.Time
}

// Pagination for api calls to define how much records should be returns in an api response
// or which page of a list of records should be returned.
type Pagination struct {
	Page    *int
	PerPage *int
}

// Strava API Models
// Details: https://developers.strava.com/docs/reference/#api-models-ActivityTotal

// A set of rolled-up statistics and totals for an athlete
type ActivityStats struct {

	// The longest distance ridden by the athlete.
	BiggestRideDistance float64 `json:"biggest_ride_distance"`

	// double	The highest climb ridden by the athlete.
	BiggestGlimbElevationGain float64 `json:"biggest_climb_elevation_gain"`

	// The recent (last 4 weeks) ride stats for the athlete.
	RecentRideTotals ActivityTotal `json:"recent_ride_totals"`

	// The recent (last 4 weeks) run stats for the athlete.
	RecentRunTotals ActivityTotal `json:"recent_run_totals"`

	// The recent (last 4 weeks) swim stats for the athlete.
	RecentSwimTotals ActivityTotal `json:"recent_swim_totals"`

	// The year to date ride stats for the athlete.
	YearToDateRideTotals ActivityTotal `json:"ytd_ride_totals"`

	// The year to date run stats for the athlete.
	YearToDateRunTotals ActivityTotal `json:"ytd_run_totals"`

	// The year to date swim stats for the athlete.
	YearToDateSwimTotals ActivityTotal `json:"ytd_swim_totals"`

	// The all time ride stats for the aTthlete.
	AllRideotals ActivityTotal `json:"all_ride_totals"`

	// The all time run stats for the athlete.
	AllRunTotals ActivityTotal `json:"all_run_totals"`

	// The all time swim stats for the athlete
	AllSwimTotals ActivityTotal `json:"all_swim_totals"`
}

// A roll-up of metrics pertaining to a set of activities. Values are in seconds and meters.
type ActivityTotal struct {

	// The number of activities considered in this total.
	Count int64 `json:"count"`

	// The total distance covered by the considered activities in meters.
	Distance float64 `json:"distance"`

	// The total moving time of the considered activities in seconds.
	MovingTime int64 `json:"moving_time"`

	// The total elapsed time of the considered activities in seconds.
	ElapsedTime int `json:"elapsed_time"`

	// The total elevation gain of the considered activities.
	ElevationGain float64 `json:"elevation_gain"`

	// The total number of achievements of the considered activities.
	AchievementCount int `json:"achievement_count"`
}

// Description of a single activity.
type SummaryActivity struct {

	// The unique identifier of the activity
	Id int `json:"id"`

	// The name of the activity
	Name string `json:"name"`

	// The activity's distance, in meters
	Distance float64 `json:"distance"`

	// The activity's moving time, in seconds
	MovingTime int64 `json:"moving_time"`

	// An enumeration of the sport types an activity may have. Distinct from ActivityType in that it has new types (e.g. MountainBikeRide)
	// May be one of the following values:
	//		AlpineSki, BackcountrySki, Canoeing, Crossfit, EBikeRide, Elliptical, EMountainBikeRide, Golf, GravelRide,
	//		Handcycle, Hike, IceSkate, InlineSkate, Kayaking, Kitesurf, MountainBikeRide, NordicSki, Ride, RockClimbing,
	//		RollerSki, Rowing, Run, Sail, Skateboard, Snowboard, Snowshoe, Soccer, StairStepper, StandUpPaddling, Surfing,
	//		Swim, TrailRun, Velomobile, VirtualRide, VirtualRun, Walk, WeightTraining, Wheelchair, Windsurf, Workout, Yoga
	SportType string `json:"sport_type"`

	// The time at which the activity was started.
	StartDateLocal time.Time `json:"start_date_local"`
}

// DetailedAthlete contains defail information of an athlete.
type DetailedAthlete struct {

	// The unique identifier of the athlete
	Id int64 `json:"id"`
}

// Error occurred in an api call.
type Error struct {

	// The code associated with this error.
	Code string `json:"code"`

	// The specific field or aspect of the resource associated with this error.
	Field string `json:"field"`

	// The type of resource associated with this error.
	Resource string `json:"resource"`
}

// Encapsulates the errors that may be returned from the API.
type Fault struct {

	// The set of specific errors associated with this fault, if any.
	Errors []Error `json:"errors"`

	// The message of the fault.
	Message string `json:"message"`
}
