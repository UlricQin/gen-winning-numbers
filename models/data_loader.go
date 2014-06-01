package models

import (
	"fmt"
	"github.com/ulricqin/gen-winning-numbers/g"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadDuoYuanDirs() error {
	dirs, err := ioutil.ReadDir(Global.BaseDir)
	if err != nil {
		return err
	}

	sz := len(dirs)
	if sz == 0 {
		return nil
	}

	if Global.DuoYuanList == nil {
		Global.DuoYuanList = make(map[string]*DuoYuan)
	}

	for i := 0; i < sz; i++ {
		dirName := dirs[i].Name()
		if dirName == "." || dirName == ".." {
			continue
		}

		if _, ok := Global.DuoYuanList[dirName]; ok {
			continue
		}

		dy := &DuoYuan{DirName: dirName}
		if err = dy.FillDetail(); err != nil {
			return err
		}

		Global.DuoYuanList[dirName] = dy
	}

	return nil
}

func (this *DuoYuan) FillDetail() (err error) {
	duoYuanDir := filepath.Join(Global.BaseDir, this.DirName)

	var jiYiDirs []os.FileInfo
	jiYiDirs, err = ioutil.ReadDir(duoYuanDir)
	if err != nil {
		return
	}

	sz := len(jiYiDirs)
	if sz == 0 {
		return
	}

	this.JiYiList = make(map[string]*JiYi)
	for i := 0; i < sz; i++ {
		jiYiDirName := jiYiDirs[i].Name()
		if jiYiDirName == "." || jiYiDirName == ".." {
			continue
		}

		// da di list
		ddl := make(map[string]*DaDi)
		for i := 1; i <= g.DaDiNum; i++ {
			name := fmt.Sprintf("da_di_%02d", i)
			ddl[name] = &DaDi{Name: name}
		}

		jiyi := &JiYi{DirName: jiYiDirName, DaDiList: ddl}
		this.JiYiList[jiYiDirName] = jiyi
	}
	return
}
