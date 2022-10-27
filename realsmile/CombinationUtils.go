package realsmile

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"strconv"
)

type CombinationUtils struct {
	m, n int
}

var (
	Combination CombinationUtils
)

func (this *CombinationUtils) GetCombination(m, n int) (*arraylist.List, error) {
	this.n = n
	this.m = m
	if n > m {
		return nil, this
	}
	flagArray := make([]int, m)

	for i := 0; i < m; i++ {
		if i < n {
			flagArray[i] = 1
		} else {
			flagArray[i] = 0
		}
	}
	list := arraylist.New()
	list.Add(this.checkEveryCombination(flagArray, m))
	var bFind = true
	for bFind {
		bFind = false
		for i := 0; i < m-1; i++ {
			if flagArray[i] == 1 && flagArray[i+1] == 0 {
				flagArray[i] = 0
				flagArray[i+1] = 1
				bFind = true
				if flagArray[0] == 0 {
					j := 0
					for k := 0; k < i; k++ {
						if flagArray[k] == 1 {
							flagArray[k] = 0
							flagArray[j] = 1
							j++
						}
					}
				}
				break
			}
		}
		if bFind {
			list.Add(this.checkEveryCombination(flagArray, m))
		}
	}
	return list, nil
}
func (this *CombinationUtils) Error() string {
	return fmt.Sprintf("传入参数错误，M选N的组合，N ≤ M。%v ≤ %v", this.n, this.m)
}

func (this *CombinationUtils) checkEveryCombination(flagArray []int, m int) string {
	var result string
	for index := 0; index < m; index++ {
		if flagArray[index] == 1 {
			if len(result) > 0 {
				result += ","
			}
			result += strconv.Itoa(index + 1)
		}
	}
	return result
}
func (this *CombinationUtils) CombinationSize(m, n int) int {
	if m < n {
		panic("m ≥ n")
	}
	mn := m - n
	nFact := this.mFactorial(n)
	mmn := this.mnFactorial(m, mn)
	return mmn / nFact
}
func (this *CombinationUtils) mFactorial(m int) int {
	if 0 == m {
		return 1
	}
	return m * this.mFactorial(m-1)
}
func (this *CombinationUtils) mnFactorial(m, n int) int {
	if m == n {
		return 1
	}
	return m * this.mnFactorial(m-1, n)
}
