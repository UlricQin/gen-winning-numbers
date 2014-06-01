package utils

import (
	"fmt"
	"github.com/ulricqin/goutils/filetool"
	log "github.com/ulricqin/goutils/logtool"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

/**
 * each line of txt should be: \d{3}\r\n
 */
func NumExists(txt string, search []byte) (bool, error) {
	b, err := ioutil.ReadFile(txt)
	if err != nil {
		return false, err
	}

	s0 := search[0]
	s1 := search[1]
	s2 := search[2]

	sz := len(b)

	for i := 0; i < sz; i += 5 {
		if b[i] == s0 && b[i+1] == s1 && b[i+2] == s2 {
			return true, nil
		}
	}

	return false, err
}

func ReadDaDi(path string) []int {
	if !filetool.IsExist(path) {
		return []int{}
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return []int{}
	}

	sz := len(b)
	ret_sz := sz / 5

	ret := make([]int, ret_sz)
	ret_idx := 0
	for i := 0; i < sz; i += 5 {
		strNum := string([]byte{b[i], b[i+1], b[i+2]})
		tmp, _ := strconv.ParseInt(strNum, 10, 64)
		ret[ret_idx] = int(tmp)
		ret_idx++
	}

	return ret
}

func ReadRongCuoFile(rongCuoFile string) (min, max int, err error) {
	if !filetool.IsExist(rongCuoFile) {
		return
	}

	var content string
	content, err = filetool.ReadFileToStringNoLn(rongCuoFile)
	if err != nil {
		return
	}

	arr := strings.Split(content, "-")
	if len(arr) != 2 {
		err = fmt.Errorf("%s format error. right format like: 0-0 ", rongCuoFile)
		return
	}

	if v, e := strconv.ParseInt(arr[0], 10, 64); e != nil {
		err = fmt.Errorf("parse %s fail. error: %s", rongCuoFile, e)
		return
	} else {
		min = int(v)
	}

	if v, e := strconv.ParseInt(arr[1], 10, 64); e != nil {
		err = fmt.Errorf("parse %s fail. error: %s", rongCuoFile, e)
	} else {
		max = int(v)
	}
	return
}

func NextDuoYuanDirName(baseDir string) (string, error) {
	first := "duo_yuan_01"
	if !filetool.IsExist(baseDir) {
		err := filetool.InsureDir(baseDir)
		if err != nil {
			return "", err
		}
		return first, nil
	}

	dirs, e := ioutil.ReadDir(baseDir)

	if e != nil {
		return "", e
	}

	sz := len(dirs)
	if sz == 0 {
		return first, nil
	}

	var safeDirs []string = make([]string, 0)

	for i := 0; i < sz; i++ {
		dirName := dirs[i].Name()
		if dirName == "." || dirName == ".." {
			continue
		} else {
			safeDirs = append(safeDirs, dirName)
		}
	}

	safeDirsLen := len(safeDirs)
	if safeDirsLen == 0 {
		return first, nil
	}

	sort.Strings(safeDirs)

	lastDirName := safeDirs[safeDirsLen-1]

	arr := strings.Split(lastDirName, "_")
	dStr := arr[len(arr)-1]

	d, er := strconv.ParseInt(dStr, 10, 64)
	if er != nil {
		return "", er
	}

	d++
	return fmt.Sprintf("duo_yuan_%02d", d), nil

}

func GetNeedSlice(rongCuoMin, rongCuoMax, total int) []int {
	if rongCuoMin > rongCuoMax {
		log.Error("rongCuoMin > rongCuoMax")
		return []int{}
	}

	if rongCuoMax > total {
		log.Error("rongCuoMax > total")
		rongCuoMax = total
	}

	begin := total - rongCuoMax
	end := total - rongCuoMin
	size := end - begin + 1

	ret := make([]int, size)
	for i := 0; i < size; i++ {
		ret[i] = begin + i
	}

	return ret
}

func AntiCollection1000(raw []int) []int {
	raw_size := len(raw)
	if raw_size == 0 {
		ret := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			ret[i] = i
		}
		return ret
	}

	ret_size := 1000 - raw_size
	ret_list := make([]int, ret_size)
	ret_idx := 0

	for i := 0; i < raw[0]; i++ {
		ret_list[ret_idx] = i
		ret_idx++
	}

	if raw_size > 1 {
		for i := 1; i < raw_size; i++ {
			for j := raw[i-1] + 1; j < raw[i]; j++ {
				ret_list[ret_idx] = j
				ret_idx++
			}
		}
	}

	for i := raw[raw_size-1] + 1; i < 1000; i++ {
		ret_list[ret_idx] = i
		ret_idx++
	}

	return ret_list

}
