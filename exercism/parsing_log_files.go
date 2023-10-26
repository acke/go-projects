package parsinglogfiles

import (
    "regexp"
)


func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WARN|ERR|FTL)\]`)

    if err != nil { return false }

	return re.MatchString(text)
}

//In regexp * is 0 or more, + is 1 or more.
func SplitLogLine(text string) []string {

    re, err := regexp.Compile(`[<][-=\*~>]*[>]`)

    if err != nil { return nil }

    return re.Split(text, -1)
}

//In regexp (?i) makes expression case-insensitive
func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)".*password.*"`)
	counter := 0

    if err != nil { return 0 }

    for _, line := range lines {
        if re.MatchString(line) {
            counter++

        }
    }

    return counter
}

//In regexp \d replaces 0-9, any number
func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d*`)

    if err != nil { return "" }

    return re.ReplaceAllString(text, "")
}

//In regexp \w replaces [a-zA-Z0-9]
func TagWithUserName(lines []string) []string {

    re, err := regexp.Compile(`User\s+[\w]+`)
    if err != nil { return nil }

    regexp_replace, err := regexp.Compile(`User\s+`)
	if err != nil { return nil }

	for i, line := range lines {
        submatch := re.FindStringSubmatch(line)

        if len(submatch) > 0 {
            newline := regexp_replace.ReplaceAllString(submatch[0], "[USR] ")
            lines[i] = newline + " " + line
        }
    }

    return lines
}

