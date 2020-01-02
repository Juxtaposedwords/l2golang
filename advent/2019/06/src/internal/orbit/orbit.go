package orbit

import (
	"bufio"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"

	"strings"
)

func sanDistance(input io.Reader) (int, error) {
	orbitMap, err := orbitMapBuilder(input)
	if err != nil {
		return 0, err
	}
	youPath, ok := orbitPath("COM", "YOU", orbitMap)
	if !ok {
		return 0, status.Error(codes.FailedPrecondition, "no path found for you")
	}
	sanPath, ok := orbitPath("COM", "SAN", orbitMap)
	if !ok {
		return 0, status.Error(codes.FailedPrecondition, "no path found for you")
	}

	shortest, longest := youPath, sanPath
	if len(shortest) > len(longest) {
		longest, shortest = shortest, longest
	}
	shortest = append(shortest, "COM")
	reverse(shortest)
	longest = append(longest, "COM")
	reverse(longest)


	var index int
	for index = 0; index < len(shortest); index++ {
		if shortest[index] != longest[index] {
			break
		}
	}
	if index == len(shortest) -1 {
		return 0, status.Error(codes.NotFound, "unable to find a shared path")
	}
	return (len(shortest) -1 - index) + (len(longest) -1 - index) , nil
}

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
func reverse(input []string) {
	for i := 0; i < len(input)/2; i++ {
		input[len(input)-1-i], input[i] = input[i], input[len(input)-1-i]
	}
}
func orbitPath(body string, target string, orbitMap map[string][]string) ([]string, bool) {
	if body == target {
		return []string{}, true
	}
	for _, planet := range orbitMap[body] {
		path, isChild := orbitPath(planet, target, orbitMap)
		if isChild {
			path = append(path, planet)
			return path, true
		}
	}
	return nil, false
}

func orbitMapBuilder(input io.Reader) (map[string][]string, error) {
	orbitMap := map[string][]string{}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		bodies := strings.Split(strings.Replace(scanner.Text(), " ", "", -1), ")")
		if len(bodies) != 2 {
			return nil, status.Errorf(codes.Internal, "orbit pattern does not contain two bodies: %s", scanner.Text())
		}
		orbitee, orbiter, existingOrbiters := bodies[0], bodies[1], orbitMap[bodies[0]]
		if !contains(orbiter, existingOrbiters) {
			orbitMap[orbitee] = append(existingOrbiters, orbiter)
		}
	}
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
