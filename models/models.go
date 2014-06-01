package models

type DuoYuan struct {
	DirName    string
	RongCuoMin int
	RongCuoMax int
	Checked    int
	JiYiList   map[string]*JiYi
}

type JiYi struct {
	DirName    string
	RongCuoMin int
	RongCuoMax int
	Checked    int
	DaDiList   map[string]*DaDi
}

type DaDi struct {
	Name      string
	Checked   int
	SearchNum int
}

type GlobalConfig struct {
	BaseDir           string
	DuoYuanRongCuoMin int
	DuoYuanRongCuoMax int
	Last1             string
	Last2             string
	DuoYuanList       map[string]*DuoYuan
}

var (
	Global GlobalConfig
)

func SetBaseDir(baseDir string) {
	Global.BaseDir = baseDir
}
