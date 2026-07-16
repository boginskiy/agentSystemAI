package config

type Conf struct {
}

func NewConf() *Conf {
	return &Conf{}
}

func (c *Conf) GetStorePath() string { // STORE_PATH
	return "store"
}
