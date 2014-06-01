package models

import (
	"fmt"
	"github.com/ulricqin/gen-winning-numbers/utils"
	"github.com/ulricqin/goutils/filetool"
	"math/rand"
	"path/filepath"
	// "sort"
	"strings"
	"time"
)

func GenerateOneDuoYuanDir(min, max int64) error {
	nextDir, err := utils.NextDuoYuanDirName(Global.BaseDir)
	if err != nil {
		return err
	}

	// duoyuan目录已经生成，接下来在下面生成20个记忆
	for i := 1; i <= 20; i++ {
		go GenerateOneJiYiDir(min, max, filepath.Join(Global.BaseDir, nextDir, fmt.Sprintf("ji_yi_%02d", i)), int64(i))
	}
	return nil
}

func GenerateOneJiYiDir(min, max int64, dir string, jiYiIndex int64) error {
	err := filetool.InsureDir(dir)
	if err != nil {
		return err
	}

	var t int64 = 100000

	for i := 0; i < 10000; i++ {
		filename := fmt.Sprintf("%04d.txt", i)
		if err = genOneFile(filepath.Join(dir, filename), jiYiIndex*t+int64(i)+time.Now().UnixNano(), min, max); err != nil {
			return err
		}
	}

	return nil
}

func genOneFile(filename string, seed, min, max int64) (err error) {
	r := rand.New(rand.NewSource(seed))

	diff := max - min
	numCnt := r.Intn(int(diff)) + int(min)

	// 生成numCnt个数字，写入filename，数字注意格式化
	m := make(map[int]bool)
	i := 0
	for {
		a := r.Intn(1000)
		if _, ok := m[a]; !ok {
			m[a] = true
			i++
			if i == numCnt {
				break
			}
		}
	}

	size := len(m)
	j := 0
	keys := make([]int, size)
	for k, _ := range m {
		keys[j] = k
		j++
	}

	// sort.Ints(keys)

	nums := make([]string, size)
	for k := 0; k < size; k++ {
		nums[k] = fmt.Sprintf("%03d\r\n", keys[k])
	}

	_, err = filetool.WriteStringToFile(filename, strings.Join(nums, ""))

	return
}
