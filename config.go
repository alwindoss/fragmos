package fragmos

import "time"

type Config struct {
	Home             string        `env:"HOME"`
	MiningDifficulty int           `env:"FRAGMOS_MINING_DIFFICULTY"`
	Port             int           `env:"PORT" envDefault:"3030"`
	Password         string        `env:"PASSWORD,unset"`
	IsProduction     bool          `env:"PRODUCTION"`
	Hosts            []string      `env:"HOSTS" envSeparator:":"`
	Duration         time.Duration `env:"DURATION"`
	TempFolder       string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}
