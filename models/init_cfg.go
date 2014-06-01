package models

import (
	"encoding/json"
	"github.com/ulricqin/gen-winning-numbers/g"
	"github.com/ulricqin/goutils/filetool"
	log "github.com/ulricqin/goutils/logtool"
	"os"
)

func InitBusinessCfg() {
	if filetool.IsExist(g.ConfigJsonFile) {
		b, err := filetool.ReadFileToBytes(g.ConfigJsonFile)
		if err != nil {
			log.Fetal("read %s fail. error: %s", g.ConfigJsonFile, err)
			os.Exit(1)
		}

		err = json.Unmarshal(b, &Global)
		if err != nil {
			log.Fetal("json.Unmarshal %s fail. error: %s", g.ConfigJsonFile, err)
			os.Exit(1)
		}
	}
}
