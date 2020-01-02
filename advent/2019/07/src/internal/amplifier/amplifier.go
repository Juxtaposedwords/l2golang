package amplifier

import (
//	"github.com/google/logger"
	//	"google.golang.org/grpc/codes"
	//	"google.golang.org/grpc/status"
	"internal/intcode"
)

// LargestPhaseCombinationResult takes the given software and tries all possible orderings of individual phases.
func LargestPhaseCombinationResult(software []int, phases []int, intialInput int) (int, []int, error) {
	phasePermtuations := generateIntPermutations(phases)
	//logger.Infof("%#v\n", phasePermtuations)
	var highest int
	var highestCombo []int
	for _, combination := range phasePermtuations {
	//	logger.Infof("%d\n", combination)
		result, err := thrusterValue(software, combination, intialInput)
		if err != nil {
			return 0, nil, err
		}
		if result > highest {
			highest, highestCombo = result, combination
		}
	}
	return highest, highestCombo, nil
}

func thrusterValue(software []int, combination []int, intialInput int) (int, error) {
	input := intialInput
	//	logger.Infof("*** %#v", combination)

	for _, phase := range combination {
		softwareCopy := make([]int, len(software))
		copy(softwareCopy, software)
		var err error
		input, err = intcode.DiagnosticCode(softwareCopy, []int{phase, input})
		if err != nil {
			return 0, err
		}
	}
	return input, nil
}

// shamelessly stolen

func generateIntPermutations(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
