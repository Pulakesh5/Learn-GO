package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	// prefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	// re, err := regexp.Compile(`^(?:` + strings.Join(prefixes, "|") + `)`)
    re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    if(err != nil) {
        fmt.Println(err)
        return false
    } else {
		if re.MatchString(text) {
            return true
        } else {
        	return false    
        }
    }   
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[-=~*]*>`)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    splittedLines := re.Split(text, -1)
    return splittedLines
    
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)".*password.*"`)
    count := 0
    for _, line := range lines {
		if re.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line[0-9]*`)
    removedEndOfLineText := re.ReplaceAllString(text, "")
    return removedEndOfLineText
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string,len(lines))

    re, _ := regexp.Compile(`User[ ]+(?i)[a-zA-z0-9]+`)
    for index, line := range lines {
        usr := re.FindStringSubmatch(line)
        if(usr != nil) {
			usr := regexp.MustCompile(`[ ]+`).Split(usr[0], -1)
            line = "[USR] " + usr[1] + " " + line
            // fmt.Printf("line no %v has user %v: %v\n",index, usr, line)
        } else {
            // fmt.Printf("line no %v has no user: %v\n",index, line)
        }
        taggedLines[index] = line
    }
    return taggedLines
}
