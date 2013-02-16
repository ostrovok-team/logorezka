package outputs

import "strings"

func makeRuneFilenameSafe(c rune) rune {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '-' || c == '.' {
		return c
	}
	return '_'
}

func MakeFilenameSafe(name string) string {
	return strings.Map(makeRuneFilenameSafe, name)
}

func makeRunePathSafe(c rune) rune {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '-' || c == '.' || c == '/' {
		return c
	}
	return '_'
}

func MakePathSafe(name string) string {
	return strings.Map(makeRunePathSafe, name)
}
