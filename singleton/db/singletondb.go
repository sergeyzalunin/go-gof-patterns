package db

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

// thread safety approach can be made by using 2 options
// sync.Once (laziness approach) or init()
// init() — we could, but it's not lazy

var once sync.Once
var instance *SingletonDb

// Database is abstraction over db
type Database interface {
	GetPopulation(name string) int
}

// SingletonDb is an example of db
type SingletonDb struct {
	capitals map[string]int
}

// GetPopulation of the capital
func (db *SingletonDb) GetPopulation(name string) int {
	return db.capitals[name]
}

// GetTotalPopulation by presented cities
func GetTotalPopulation(cities []string) int {
	var result int
	for _, city := range cities {
		result += GetSingletonDb().GetPopulation(city)
	}
	return result
}

// GetTotalPopulationEx by presented cities from particular db
func GetTotalPopulationEx(db Database, cities []string) int {
	var result int
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

// GetSingletonDb ...
func GetSingletonDb() *SingletonDb {
	once.Do(func() {
		caps, e := readData("capitals.txt")
		db := SingletonDb{}
		if e == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}
