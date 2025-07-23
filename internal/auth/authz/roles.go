package authz

var RolePermissions = map[string][]string{
	// Admin
	"admin": {"*"},

	// FIS roles
	"fis": {
		"GET:/v1/fis",
		"POST:/v1/fis",
		"PUT:/v1/fis",
		"DELETE:/v1/fis",
	},
	"fis_read": {
		"GET:/v1/fis",
	},

	// UTV roles
	"utv": {
		"GET:/v1/utv",
		"POST:/v1/utv",
		"PUT:/v1/utv",
		"DELETE:/v1/utv",
	},
	"utv_read": {
		"GET:/v1/utv",
	},

	// KAMK roles
	"kamk": {
		"GET:/v1/kamk",
		"POST:/v1/kamk",
		"PUT:/v1/kamk",
		"DELETE:/v1/kamk",
	},
	"kamk_read": {
		"GET:/v1/kamk",
	},

	// K-LAB roles
	"klab": {
		"GET:/v1/klab",
		"POST:/v1/klab",
		"PUT:/v1/klab",
		"DELETE:/v1/klab",
	},
	"klab_read": {
		"GET:/v1/klab",
	},

	// Tietoevry roles
	"tietoevry": {
		"GET:/v1/tietoevry",
		"POST:/v1/tietoevry",
		"PUT:/v1/tietoevry",
		"DELETE:/v1/tietoevry",
	},
	"tietoevry_read": {
		"GET:/v1/tietoevry",
	},

	// Coachtech roles
	"coachtech": {
		"GET:/v1/coachtech",
		"POST:/v1/coachtech",
		"PUT:/v1/coachtech",
		"DELETE:/v1/coachtech",
	},
	"coachtech_read": {
		"GET:/v1/coachtech",
	},

	// Archinisis roles
	"archinisis": {
		"GET:/v1/archinisis",
		"POST:/v1/archinisis",
		"PUT:/v1/archinisis",
		"DELETE:/v1/archinisis",
	},
	"archinisis_read": {
		"GET:/v1/archinisis",
	},
}
