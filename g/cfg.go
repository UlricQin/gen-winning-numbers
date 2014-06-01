package g

var (
	RootName string
	RootPass string
	DaDiNum  int
)

func initCfg() {
	RootName = Cfg.String("root_name")
	RootPass = Cfg.String("root_pass")
	DaDiNum, _ = Cfg.Int("da_di_num")
}
