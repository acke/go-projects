package parsinglogfiles

import (
    "regexp"
)


func IsValidLine(text string) bool {
	prepend_string := []string {"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}

    valid_line := false

    if len(text) < 5 {
        return valid_line
    }

    for _, str := range prepend_string {
        if str == text[:5]  {
            valid_line = true
        }
	}

	return valid_line
}

func SplitLogLine(text string) []string {

    re, err := regexp.Compile(`[<][-=\*~>]*[>]`)

    if err != nil { return nil }

    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`\".*(?i)password.*\"`)
	counter := 0

     if err != nil { return 0 }

    for _, line := range lines {
        if re.MatchString(line) {
            counter++

        }
    }

    return counter
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line[0-9]*`)

    if err != nil { return "" }

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	re, err := regexp.Compile(`User\s+[a-zA-Z0-9]+`)
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

