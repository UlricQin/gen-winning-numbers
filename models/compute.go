package models

import (
	"fmt"
	"github.com/ulricqin/gen-winning-numbers/utils"
	log "github.com/ulricqin/goutils/logtool"
	"path/filepath"
	"strconv"
	"sync"
)

// 直接返回一批奖号
func ComputeModel2() ([]string, error) {
	// 本期中奖号在上期对应的文本中进行搜索。判断第几个文本是正确的
	// 概率上讲：上期正确的顺序本期应该不会出现了

	var total int

	var duoYuanResult map[string][]int = make(map[string][]int)
	var wg sync.WaitGroup
	for _, dy := range Global.DuoYuanList {
		if dy.Checked > 0 {
			total++
			duoYuanResult[dy.DirName] = []int{}
			wg.Add(1)
			go OneDuoYuan(&wg, dy, duoYuanResult)
		}
	}

	wg.Wait()

	log.Info(">>>>> has %d duo yuan for compute", total)

	needSlice := utils.GetNeedSlice(Global.DuoYuanRongCuoMin, Global.DuoYuanRongCuoMax, total)
	log.Info(">>>>> duo yuan need silce: %v", needSlice)

	retInts := utils.IntsIntersection(duoYuanResult, needSlice)

	sz := len(retInts)
	ret := make([]string, sz)
	for i := 0; i < sz; i++ {
		ret[i] = fmt.Sprintf("%03d", retInts[i])
	}

	log.Info(">>>>> ret list size: %d", sz)

	return ret, nil
}

func OneDuoYuan(wg *sync.WaitGroup, dy *DuoYuan, duoYuanResult map[string][]int) {
	defer wg.Done()

	var total int

	var jiYiResult map[string][]int = make(map[string][]int)
	var jiYiWG sync.WaitGroup
	for _, jiYi := range dy.JiYiList {
		if jiYi.Checked > 0 {
			total++
			jiYiResult[jiYi.DirName] = []int{}
			jiYiWG.Add(1)
			go OneJiYi(&jiYiWG, dy, jiYi, jiYiResult)
		}
	}

	log.Info(">>>>> there are %d ji yi in %s", total, dy.DirName)

	jiYiWG.Wait()

	needSlice := utils.GetNeedSlice(dy.RongCuoMin, dy.RongCuoMax, total)

	ints := utils.IntsIntersection(jiYiResult, needSlice)
	log.Info(">>>>> ret list size: %d of %s", len(ints), dy.DirName)
	duoYuanResult[dy.DirName] = ints
}

func OneJiYi(wg *sync.WaitGroup, dy *DuoYuan, jiYi *JiYi, jiYiResult map[string][]int) {
	defer wg.Done()

	// 拿着这期奖号去上期奖号对应的文本中搜索，看看有哪几个是正确的，比如0,1,2,3,4,59,66是正确的
	// 要拿到所有这些顺序在本期对应的文本
	orders := currPeriodOrders(dy, jiYi)
	log.Info(">>>>>> %s/%s orders: %v", dy.DirName, jiYi.DirName, orders)
	sz := len(orders)
	if sz == 0 {
		log.Warn("no orders in %s/%s", dy.DirName, jiYi.DirName)
		return
	}

	txtBegin64, err := strconv.ParseInt(Global.Last1[1:], 10, 64)
	if err != nil {
		log.Error("parse Global.last1[1:] fail. error: %s", err)
		return
	}
	txtBegin := int(txtBegin64)

	txtList := make([]string, sz)
	for i := 0; i < sz; i++ {
		tmp := txtBegin + orders[i]
		if tmp > 9999 {
			tmp -= 10000
		}
		txtList[i] = fmt.Sprintf("%04d.txt", tmp)
	}

	pathPrefix := filepath.Join(Global.BaseDir, dy.DirName, jiYi.DirName)

	// 处理各个txt文件，读取出数字，然后求交集
	txtInts := make(map[string][]int)
	var readWG sync.WaitGroup

	for i := 0; i < sz; i++ {
		txtInts[txtList[i]] = []int{}
		readWG.Add(1)
		go ReadOneTxt(&readWG, filepath.Join(pathPrefix, txtList[i]), txtList[i], txtInts)
	}

	readWG.Wait()

	tmpInts := utils.IntsIntersection(txtInts, []int{sz})
	log.Info("[%s/%s] after intersection len: %d", dy.DirName, jiYi.DirName, len(tmpInts))
	jiYiResult[jiYi.DirName] = utils.AntiCollection1000(tmpInts)
}

func ReadOneTxt(wg *sync.WaitGroup, txtPath, txtFilename string, txtInts map[string][]int) {
	defer wg.Done()
	txtInts[txtFilename] = utils.ReadDaDi(txtPath)
}

func currPeriodOrders(dy *DuoYuan, jiYi *JiYi) []int {
	daDi := jiYi.DaDiList["da_di_01"]
	searchCnt := daDi.SearchNum
	searchCntLess1 := searchCnt - 1
	// e.g. []int{0,3,9,11,20}
	var retIntSlice []int = make([]int, searchCnt)
	retIdx := 0

	searchBytes := []byte(Global.Last1[2:])

	beginTxtIndex64, err := strconv.ParseInt(Global.Last2[1:], 10, 64)
	if err != nil {
		log.Error("parse last2 winning number fail. error: %s", err)
		return []int{}
	}
	beginTxtIndex := int(beginTxtIndex64)

	var done bool
	var exists bool

	prefixPath := filepath.Join(Global.BaseDir, dy.DirName, jiYi.DirName)

	order := 0
	for i := beginTxtIndex; i <= 9999; i++ {
		txt := filepath.Join(prefixPath, fmt.Sprintf("%04d.txt", i))
		exists, err = utils.NumExists(txt, searchBytes)
		if err != nil {
			log.Error("check number: %s exists in %s fail. error: %s", string(searchBytes), txt, err)
			break
		}

		if exists {
			retIntSlice[retIdx] = order
			if retIdx == searchCntLess1 {
				done = true
				break
			}
			retIdx++
		}
		order++
	}

	if !done {
		for i := 0; i < beginTxtIndex; i++ {
			txt := filepath.Join(prefixPath, fmt.Sprintf("%04d.txt", i))
			exists, err = utils.NumExists(txt, searchBytes)
			if err != nil {
				log.Error("check number: %s exists in %s fail. error: %s", string(searchBytes), txt, err)
				break
			}

			if exists {
				retIntSlice[retIdx] = order
				if retIdx == searchCntLess1 {
					break
				}
				retIdx++
			}
			order++
		}
	}

	return retIntSlice
}
