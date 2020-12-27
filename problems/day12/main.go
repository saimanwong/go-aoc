package day12

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type ins struct {
	act rune
	n   int
}

func parseInput(input []string) []ins {
	ret := []ins{}
	re := regexp.MustCompile("([A-Z]{1})([0-9]+)")

	for _, line := range input {
		s := re.FindStringSubmatch(line)
		act := s[1]
		n, _ := strconv.Atoi(s[2])

		ret = append(ret, ins{
			act: []rune(act)[0],
			n:   n,
		})
	}

	return ret
}

func updateRot(curr ins, currDir rune) rune {
	MAP := map[rune]map[rune]rune{
		'R': {
			'N': 'E',
			'E': 'S',
			'S': 'W',
			'W': 'N',
		},

		'L': {
			'N': 'W',
			'W': 'S',
			'S': 'E',
			'E': 'N',
		},
	}
	if curr.act == 'R' || curr.act == 'L' {
		for i := 0; i < curr.n/90; i++ {
			currDir = MAP[curr.act][currDir]
		}
	}
	return currDir
}

func updateLat(curr ins, currDir rune, currLat int) int {
	DIR := map[rune]int{
		'N': 1,
		'S': -1,
	}

	if curr.act == 'N' || curr.act == 'S' {
		currLat += DIR[curr.act] * curr.n
	}

	return currLat
}

func updateLon(curr ins, currDir rune, currLon int) int {
	DIR := map[rune]int{
		'E': 1,
		'W': -1,
	}

	if curr.act == 'E' || curr.act == 'W' {
		currLon += DIR[curr.act] * curr.n
	}

	return currLon
}

func updateFor(curr ins, currDir rune, currLon int, currLat int) (int, int) {
	DIR := map[rune]int{
		'N': 1,
		'S': -1,
		'E': 1,
		'W': -1,
	}

	if curr.act == 'F' {
		if currDir == 'E' || currDir == 'W' {
			currLon += DIR[currDir] * curr.n
		}
	}

	if curr.act == 'F' {
		if currDir == 'N' || currDir == 'S' {
			currLat += DIR[currDir] * curr.n
		}
	}

	return currLon, currLat
}

func problemOne(in []ins) string {
	currDir := 'E'
	currLon := 0
	currLat := 0
	for _, i := range in {
		if i.act == 'F' {
			currLon, currLat = updateFor(i, currDir, currLon, currLat)
		}

		if i.act == 'E' || i.act == 'W' {
			currLon = updateLon(i, currDir, currLon)
		}

		if i.act == 'N' || i.act == 'S' {
			currLat = updateLat(i, currDir, currLat)
		}

		if i.act == 'R' || i.act == 'L' {
			currDir = updateRot(i, currDir)
		}
	}
	return fmt.Sprintf("%d", int(math.Abs(float64(currLat)))+int(math.Abs(float64(currLon))))
}

func getDir(t string, n int) rune {
	if t == "lon" {
		ret := 'E'
		if n < 0 {
			ret = 'W'
		}
		return ret

	}
	if t == "lat" {
		ret := 'N'
		if n < 0 {
			ret = 'S'
		}
		return ret
	}

	return '0'
}

func status(note string, c map[string]int) {
	dirLon := getDir("lon", c["lon"])
	dirLat := getDir("lat", c["lat"])
	if c["lat"] < 0 {
		dirLat = 'S'
	}
	fmt.Printf("[%s] Lon: %d (%s), Lat: %d (%s)\n", note, c["lon"], string(dirLon), c["lat"], string(dirLat))
}

func problemTwo(in []ins) string {
	wayp := map[string]int{
		"lon": 10,
		"lat": 1,
	}
	ship := map[string]int{
		"lon": 0,
		"lat": 0,
	}
	for _, i := range in {
		// fmt.Printf("%s%d\n", string(i.act), i.n)
		// status("BEFORE WAYP", wayp)
		// status("BEFORE SHIP", ship)
		// F move ship
		if i.act == 'F' {
			ship["lon"] += wayp["lon"] * i.n
			ship["lat"] += wayp["lat"] * i.n
		}

		// NESW move wayp
		if i.act == 'N' || i.act == 'E' || i.act == 'S' || i.act == 'W' {
			dirLon := getDir("lon", wayp["lon"])
			dirLat := getDir("lat", wayp["lat"])
			wayp["lon"] = updateLon(i, dirLon, wayp["lon"])
			wayp["lat"] = updateLat(i, dirLat, wayp["lat"])
		}

		// RL rotate wayp
		// func updateRot(curr ins, currDir rune) rune {
		if i.act == 'R' || i.act == 'L' {
			oldDirLon := getDir("lon", wayp["lon"])
			oldDirLat := getDir("lat", wayp["lat"])
			newDirLon := updateRot(i, oldDirLon)
			newDirLat := updateRot(i, oldDirLat)
			// fmt.Println("OLD", string(oldDirLon), string(oldDirLat))
			// fmt.Println("NEW", string(newDirLon), string(newDirLat))

			if (oldDirLon == 'E' || oldDirLon == 'W') && (newDirLon == 'E' || newDirLon == 'W') {
				wayp["lon"] = wayp["lon"] * -1
				// fmt.Println("LON -> LON")
			}

			if (oldDirLat == 'N' || oldDirLat == 'S') && (newDirLat == 'N' || newDirLat == 'S') {
				wayp["lat"] = wayp["lat"] * -1
				// fmt.Println("LAT -> LAT")
			}

			oldLat := wayp["lat"]
			oldLon := wayp["lon"]
			if (oldDirLon == 'E' || oldDirLon == 'W') && (newDirLon == 'N' || newDirLon == 'S') {
				wayp["lat"] = oldLon
				if newDirLon == 'N' {
					wayp["lat"] = int(math.Abs(float64(wayp["lat"])))
				}
				if newDirLon == 'S' {
					wayp["lat"] = int(math.Abs(float64(wayp["lat"]))) * -1
				}
				// fmt.Println("LON -> LAT")
			}

			if (oldDirLat == 'N' || oldDirLat == 'S') && (newDirLat == 'E' || newDirLat == 'W') {
				wayp["lon"] = oldLat
				if newDirLat == 'E' {
					wayp["lon"] = int(math.Abs(float64(wayp["lon"])))
				}
				if newDirLat == 'W' {
					wayp["lon"] = int(math.Abs(float64(wayp["lon"]))) * -1
				}
				// fmt.Println("LAT -> LON")
			}
		}

		// status("AFTER WAYP", wayp)
		// status("AFTER SHIP", ship)
		// fmt.Println("")
	}

	// return ""
	return fmt.Sprintf("%d", int(math.Abs(float64(ship["lon"])))+int(math.Abs(float64(ship["lat"]))))
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}
