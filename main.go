package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {

	args := os.Args[1:]
	initialPath := args[0]

	result := traverse(initialPath)
	for _, v := range result {
		fmt.Printf("%s/%s\n", initialPath, v)
	}
}

func traverse(dir string) []string {
	directoryEntries, _ := os.ReadDir(dir)
	result := flattenNested(mapToNewArray(func(entry fs.DirEntry) []string {
		isDir := entry.IsDir()
		name := entry.Name()
		if isDir {
			nestedPath := fmt.Sprintf("%s/%s", dir, name)
			nestedResults := traverse(nestedPath)
			return mapToNewArray[string, string](func(s string) string {
				return fmt.Sprintf("%s/%s", name, s)
			}, nestedResults)
		} else {
			return []string{name}
		}
	}, directoryEntries))
	return result
}

func mapToNewArray[T any, K any](mapper func(t T) K, tArr []T) []K {
	result := make([]K, 0)
	for _, v := range tArr {
		result = append(result, mapper(v))
	}
	return result
}

func flattenNested[T any](tMatrix [][]T) []T {
	result := make([]T, 0)
	for _, strArr := range tMatrix {
		result = append(result, strArr...)
	}
	return result
}
