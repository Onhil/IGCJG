package main

// MetaInfo struct for api uptime and meta info
type MetaInfo struct {
	uptime  string
	info    string
	version string
}

// Uptime returns app uptime in the ISO 8601 standard
func Uptime() MetaInfo {
	return MetaInfo{"P", "info", "v"}
}
