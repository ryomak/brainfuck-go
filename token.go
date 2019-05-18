package bf

import (
	"regexp"
	"strings"
)

//commands is list of bf command
type commands struct {
	NEXT  string
	PREV  string
	INC   string
	DEC   string
	READ  string
	WRITE string
	OPEN  string
	CLOSE string
}

// Ops is getting opelation for regexp
func (c *commands) Ops() string {
	//for regexp
	return strings.Join(c.list(), "|")
}

// list is getting match list
func (c *commands) list() []string {
	return []string{
		regexp.QuoteMeta(c.NEXT),
		regexp.QuoteMeta(c.PREV),
		regexp.QuoteMeta(c.INC),
		regexp.QuoteMeta(c.DEC),
		regexp.QuoteMeta(c.READ),
		regexp.QuoteMeta(c.WRITE),
		regexp.QuoteMeta(c.OPEN),
		regexp.QuoteMeta(c.CLOSE),
	}
}
