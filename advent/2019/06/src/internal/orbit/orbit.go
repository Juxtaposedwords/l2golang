package orbit

import (
	"bufio"
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"

	"strings"
)

func total(input io.Reader) (int, error) {
	orbitMap, err := orbitMapBuilder(input)
	if err != nil {
		return 0, err
	}
	return orbits("COM", orbitMap, 0), nil
}

func orbits(body string, orbitMap map[string][]string, distance int) int {
	if orbitMap[body] == nil {
		return distance
	}
	var output int
	for _, planet := range orbitMap[body] {
		output += orbits(planet, orbitMap, distance+1)
	}
	return output + distance
}

func orbitMapBuilder(input io.Reader) (map[string][]string, error) {

	defer logger.Init("LoggerExample", true, false, ioutil.Discard)

	orbitMap := map[string][]string{}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		bodies := strings.Split(strings.Replace(scanner.Text()," ","",-1), ")")
		if len(bodies) != 2 {
			return nil, status.Errorf(codes.Internal, "orbit pattern does not contain two bodies: %s", scanner.Text())
		}
		orbitee, orbiter, existingOrbiters := bodies[0], bodies[1], orbitMap[bodies[0]]
		if !contains(orbiter, existingOrbiters) {
			orbitMap[orbitee] = append(existingOrbiters, orbiter)
		}
	}
	logger.Infof("%#v\n", orbitMap)
	return orbitMap, nil
}

func contains(entry string, inputSlice []string) bool {
	for _, e := range inputSlice {
		if e == entry {
			return true
		}
	}
	return false
}
