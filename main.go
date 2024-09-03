package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCSVToMap(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, record := range records {
		if len(record) < 2 {
			continue // Skip rows that don't have at least two columns
		}
		key := record[0]
		value := record[1]
		result[key] = value
	}

	return result, nil
}

// as := []int
func main() {
	as := []int{4, 3}
	var a bool
	fmt.Println(a, as)
	// fmt.Println(sortArray(as))
}

// func sortArray(nums []int) []int {
// 	quickSortHelper(nums, 0, len(nums)-1)
// 	return nums
// }
// func quickSortHelper(arr []int, left, right int) {
// 	if left < right {
// 		pivot := partition(arr, left, right)
// 		quickSortHelper(arr, left, pivot-1)
// 		quickSortHelper(arr, pivot+1, right)
// 	}
// }

// func partition(arr []int, left, right int) int {
// 	pivot := arr[right]
// 	i := left

// 	for j := left; j < right; j++ {
// 		if arr[j] < pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}

// 	arr[i], arr[right] = arr[right], arr[i]
// 	return i
// }

func handler(nums []int, target int) int {
	return handleMiddleTask(nums, 0, len(nums)-1, target)
}

func handleMiddleTask(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2
	if arr[middle] == target {
		return middle
	}
	if arr[middle] > target {
		return handleMiddleTask(arr, left, middle-1, target)
	}
	return handleMiddleTask(arr, middle+1, right, target)
}

// CompressData
func CompressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	defer writer.Close()

	_, err := writer.Write(data) // try a new compoerser here
	if err != nil {
		return nil, err
	}

	// 确保所有数据都被写入
	err = writer.Close() // if not close, must close it
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// DecompressData 解压缩输入的二进制数据
func DecompressData(compressedData []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
