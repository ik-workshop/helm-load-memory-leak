package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func cleanup(files []string) {
	for _, file := range files {
		_ = os.Remove(file)
	}
}

func TestV3RunnerWith_Fixture_Chart_SpecialCharacters(t *testing.T) {
	files, err := copyFile("testdata/chart-benchmark/tests", "main_test.yaml", 2)
	assert.NoError(t, err)
	t.Cleanup(func() { cleanup(files) })
	inALoop(len(files))
}

func BenchmarkNewTestForCPUAndMemoryHelm(b *testing.B) {
	files, err := copyFile("testdata/chart-benchmark/templates", "configmap.yaml", 400)
	assert.NoError(b, err)
	b.Cleanup(func() { cleanup(files) })
	for b.Loop() {
		inALoop(len(files))
	}
}

func BenchmarkTestV3RunnerWith_Fixture_Chart_SpecialCharactersHelm(b *testing.B) {
	files, err := copyFile("testdata/chart-benchmark/templates", "configmap.yaml", 10)
	assert.NoError(b, err)

	b.Cleanup(func() {
		for _, file := range files {
			err := os.Remove(file)
			assert.NoError(b, err)
		}
	})

	for b.Loop() {
		inALoop(len(files))
	}
}

func BenchmarkNewTestForCPUAndMemory(b *testing.B) {
	files, err := copyFile("testdata/chart-benchmark/tests", "main_test.yaml", 400)
	assert.NoError(b, err)

	b.Cleanup(func() {
		for _, file := range files {
			err := os.Remove(file)
			assert.NoError(b, err)
		}
	})

	for b.Loop() {
		inALoop(len(files))
	}
}

func copyFile(dir, src string, times int) ([]string, error) {
	// Open the source file for reading
	in, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, src))
	if err != nil {
		return nil, err
	}
	var result []string

	for i := 0; i < times; i++ {
		fileName := fmt.Sprintf("%s/%d-%s", dir, i, src)
		err := os.WriteFile(fileName, in, 0644)
		if err != nil {
			return nil, err
		}
		result = append(result, fileName)
	}

	return result, nil
}
