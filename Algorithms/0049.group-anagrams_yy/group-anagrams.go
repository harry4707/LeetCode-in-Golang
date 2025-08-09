package Problem0049

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	record := make(map[string][]string)

	for _, str := range strs {
		temp := sortString(str)
		record[temp] = append(record[temp], str)
	}
	for _, v := range record {
		sort.Strings(v)
		res = append(res, v)
	}

	return res
}

func sortString(s string) string {
	bytes := []byte(s)

	temp := make([]int, len(bytes))
	for i, b := range bytes {
		temp[i] = int(b)
	}

	sort.Ints(temp)

	for i, v := range temp {
		bytes[i] = byte(v)
	}

	return string(bytes)
}

func groupAnagramsyy(strs []string) [][]string {
	var record = make(map[string][]string)
	for _, str := range strs {
		charSlice := []byte(str)
		sort.Slice(charSlice, func(i, j int) bool {
			return charSlice[i] < charSlice[j]
		})

		record[string(charSlice)] = append(record[string(charSlice)], str)
	}
	var res [][]string
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

func groupAnagramsCountNum(strs []string) [][]string {
	var record = make(map[string][]string)
	for _, str := range strs {
		countSlice := [26]int{}
		// 构造每个char 的树木
		for _, char := range str {
			countSlice[char-'a']++
		}
		fmt.Println(fmt.Sprint(countSlice))
		fmt.Println(strings.Fields(fmt.Sprint(countSlice)))
		key := strings.Join(strings.Fields(fmt.Sprint(countSlice)), "#")
		fmt.Println(key)
		record[key] = append(record[key], str)
	}
	var res [][]string
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
