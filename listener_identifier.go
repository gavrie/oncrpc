package main

type identifier string

// TODO: Convert to method
func unexportedIdentifier(s string) identifier {
	return identifier(camelCase(s, false))
}

// TODO: Convert to method
func exportedIdentifier(s string) identifier {
	return identifier(camelCase(s, true))
}
