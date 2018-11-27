package fion

type config struct {
	tagLevel    int
	externalLog bool
	errorLog    string
}

func (c *config) setTagLevel(lv int) {
	c.tagLevel = lv
}

func (c *config) disableLog() {
	c.externalLog = false
}

func (c *config) enableLog() {
	c.externalLog = true
}

var fionConf = &config{
	tagLevel:    2,
	externalLog: true,
	errorLog:    "error.log",
}
