package tcp

type Config struct {
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	ServerName  string `json:"server_name"`  // server name of the tcp server
	Network     string `json:"network"`      // tcp、tcp4 or tcp6
	ReadBuffer  int    `json:"read_buffer"`  // tcp read buffer
	WriteBuffer int    `json:"write_buffer"` // tcp write buffer
	Log         struct {
		FileName   string `json:"file_name"`   // log output file name
		MaxSize    int    `json:"max_size"`    // Maximum size of log file, in megabytes
		MaxBackups int    `json:"max_backups"` // Maximum number of old files to keep
		MaxAges    int    `json:"max_ages"`    // Maximum days of old files to keep
		Compress   bool   `json:"compress"`    // Compress old files
	} `json:"log"`
}

func (c *Config) SetUp() error {
	return nil
}
