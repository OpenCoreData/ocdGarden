package lookups

// LookUp is a stop gap lookup function for some data hacking
func LookUp(key string) string {

	m = make(map[string]string)

	m["key"] = "value"

	return m[key]

}
