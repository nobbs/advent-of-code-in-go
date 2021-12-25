package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type XYZ struct {
	x, y, z int
}

type Triangle struct {
	p, q, r    XYZ
	circ, area int
}

type Translation struct {
	original, foreign, rotationID int
}

type MatchingPair struct {
	p          XYZ
	rotationID int
}

type Scanner struct {
	beacons      []XYZ
	fingerprints []Triangle
}

type Pair struct {
	a, b int
}

func (p XYZ) Add(q XYZ) XYZ {
	p.x += q.x
	p.y += q.y
	p.z += q.z

	return p
}

func (p XYZ) Sub(q XYZ) XYZ {
	p.x -= q.x
	p.y -= q.y
	p.z -= q.z

	return p
}

func (p XYZ) Negate() XYZ {
	p.x = -p.x
	p.y = -p.y
	p.z = -p.z

	return p
}

func (p XYZ) Equals(q XYZ) bool {
	return p.x == q.x && p.y == q.y && p.z == q.z
}

func rotateX90(p XYZ) XYZ {
	p.y, p.z = -p.z, p.y
	return p
}

func rotateY90(p XYZ) XYZ {
	p.x, p.y, p.z = -p.z, p.y, p.x
	return p
}

func rotateZ90(p XYZ) XYZ {
	p.x, p.y = p.y, -p.x
	return p
}

func rotate(p XYZ, rid int) XYZ {
	switch rid {
	case 0:
		// (0, 0, 0) - {1 2 3}
		p = p
	case 1:
		// (0, 0, 1) - {2 -1 3}
		p = rotateZ90(p)
	case 2:
		// (0, 0, 2) - {-1 -2 3}
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 3:
		// (0, 0, 3) - {-2 1 3}
		p = rotateZ90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 4:
		// (0, 1, 0) - {-3 2 1}
		p = rotateY90(p)
	case 5:
		// (0, 1, 1) - {2 3 1}
		p = rotateY90(p)
		p = rotateZ90(p)
	case 6:
		// (0, 1, 2) - {3 -2 1}
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 7:
		// (0, 1, 3) - {-2 -3 1}
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 8:
		// (0, 2, 0) - {-1 2 -3}
		p = rotateY90(p)
		p = rotateY90(p)
	case 9:
		// (0, 2, 1) - {2 1 -3}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
	case 10:
		// (0, 2, 2) - {1 -2 -3}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 11:
		// (0, 2, 3) - {-2 -1 -3}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 12:
		// (0, 3, 0) - {3 2 -1}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateY90(p)
	case 13:
		// (0, 3, 1) - {2 -3 -1}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
	case 14:
		// (0, 3, 2) - {-3 -2 -1}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 15:
		// (0, 3, 3) - {-2 3 -1}
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 16:
		// (1, 0, 0) - {1 -3 2}
		p = rotateX90(p)
	case 17:
		// (1, 0, 1) - {-3 -1 2}
		p = rotateX90(p)
		p = rotateZ90(p)
	case 18:
		// (1, 0, 2) - {-1 3 2}
		p = rotateX90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 19:
		// (1, 0, 3) - {3 1 2}
		p = rotateX90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 20:
		// (1, 2, 0) - {-1 -3 -2}
		p = rotateX90(p)
		p = rotateY90(p)
		p = rotateY90(p)
	case 21:
		// (1, 2, 1) - {-3 1 -2}
		p = rotateX90(p)
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
	case 22:
		// (1, 2, 2) - {1 3 -2}
		p = rotateX90(p)
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	case 23:
		// (1, 2, 3) - {3 -1 -2}
		p = rotateX90(p)
		p = rotateY90(p)
		p = rotateY90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
		p = rotateZ90(p)
	}

	return p
}

func l1norm(p XYZ) int {
	return util.AbsInt(p.x) + util.AbsInt(p.y) + util.AbsInt(p.z)
}

func (p XYZ) l1distance(q XYZ) int {
	return util.AbsInt(p.x-q.x) + util.AbsInt(p.y-q.y) + util.AbsInt(p.z-q.z)
}

// calculate the determinant of the vectors a, b, c
func det(a, b, c XYZ) int {
	return a.x*(b.y*c.z-b.z*c.y) - a.y*(b.x*c.z-b.z*c.x) + a.z*(b.x*c.y-b.y*c.x)
}

// calculate the "circumference" of the triangle (a, b, c) using the l1 distance
func l1circumference(a, b, c XYZ) int {
	return a.l1distance(b) + b.l1distance(c) + c.l1distance(a)
}

// calculate the "area" of the triangle (a, b, c) using the l1 norm of the cross product (b-a) x (c-a)
func l1area(a, b, c XYZ) int {
	return l1norm(crossprod(b.Sub(a), c.Sub(a)))
}

// calculate the cross product a x b
func crossprod(a, b XYZ) XYZ {
	c := XYZ{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - a.x*b.z,
		z: a.x*b.y - a.y*b.x,
	}

	return c
}

func (p XYZ) findNearestTwo(list []XYZ) (q XYZ, r XYZ) {
	distances := []struct {
		val   int
		index int
	}{}

	for i, u := range list {
		distances = append(distances, struct {
			val   int
			index int
		}{val: p.l1distance(u), index: i})
	}

	sort.Slice(distances, func(i, j int) bool { return distances[i].val < distances[j].val })
	q = list[distances[1].index]
	r = list[distances[2].index]

	return q, r
}

func parseInput(lines []string) []Scanner {
	scanners := []Scanner{}
	var current Scanner

	// iterate over the input lines
	for _, line := range lines {
		if strings.Contains(line, "scanner") {
			// if the line contains "scanner", start a new scanner object
			current = Scanner{beacons: make([]XYZ, 0)}
		} else if strings.TrimSpace(line) != "" {
			// otherwise, parse the coordinates and add it to the current scanner
			coords := strings.Split(line, ",")
			current.beacons = append(current.beacons, XYZ{
				x: util.ParseInt(coords[0]),
				y: util.ParseInt(coords[1]),
				z: util.ParseInt(coords[2]),
			})
		} else {
			// if the line is empty, the current scanner is done and can be added to the list of read scanners. but
			// first, also compute the fingerprints of the beacons
			current.fingerprints = computeFingerprints(current.beacons)
			scanners = append(scanners, current)
		}
	}
	// also don't forget to add the last one
	current.fingerprints = computeFingerprints(current.beacons)
	scanners = append(scanners, current)

	return scanners
}

func computeFingerprints(beacons []XYZ) []Triangle {
	fingerprints := make([]Triangle, 0, len(beacons))
	duplicateCheck := map[[2]int]struct{}{}

	for _, p := range beacons {
		q, r := p.findNearestTwo(beacons)
		circ := l1circumference(p, q, r)
		area := l1area(p, q, r)

		if _, ok := duplicateCheck[[2]int{area, circ}]; !ok {
			fingerprints = append(fingerprints, Triangle{p: p, q: q, r: r, circ: circ, area: area})
		}
		duplicateCheck[[2]int{area, circ}] = struct{}{}
	}

	return fingerprints
}

func matchScanners(scanners []Scanner) (map[Pair]MatchingPair, []int) {

	// some auxiliary vars
	// store a
	matched := map[Pair]MatchingPair{}

	// store every already matched scanner
	newlyMatchedScanners := []int{0}

	// store every remaining unmatched scanner
	unmatched := map[int]bool{}
	for i := 1; i < len(scanners); i++ {
		unmatched[i] = true
	}

	// prevMatch is used to store the scanner, that the scanner with the respective index was matched against - this
	// allows the reconstruction of all required rotations starting from scanner 0
	prevMatch := make([]int, len(scanners))

	// try to match a scanner that is yet unmatched
	for len(unmatched) > 0 {
		// only check against newly matched scanners
		nextNewlyMatchedScanners := []int{}
		for _, matchedScannerID := range newlyMatchedScanners {
			// also only check not yet matched scanners
			for unmatchedScannerID := range unmatched {
				alreadyMatched := false
				_, ok := matched[Pair{matchedScannerID, unmatchedScannerID}]
				if !ok {
					// compare all fingerprints of the matched scanner and the unmatched scanner
					for _, finger := range scanners[matchedScannerID].fingerprints {
						for _, fingerCheck := range scanners[unmatchedScannerID].fingerprints {
							// check fingerprint equality
							if finger.area == fingerCheck.area && finger.circ == fingerCheck.circ {
								// we found a matching fingerprint in both scanners
								// now lets try to figure out the rotation and translation

								for rotationID := 0; rotationID < 24 && !alreadyMatched; rotationID++ {
									translations := map[XYZ][]Translation{}
									for originalIndex, original := range []XYZ{finger.p, finger.q, finger.r} {
										for foreignIndex, foreign := range []XYZ{fingerCheck.p, fingerCheck.q, fingerCheck.r} {
											// compute translatio of points of the triangle in the matched scanner to
											// all possible combinations of the triangle in the not yet matched scanners
											// position system after rotation
											diff := rotate(foreign, rotationID).Sub(original)
											v, ok := translations[diff]
											if ok {
												v = append(v, Translation{originalIndex, foreignIndex, rotationID})
												translations[diff] = v
											} else {
												translations[diff] = []Translation{{originalIndex, foreignIndex, rotationID}}
											}
										}

									}
									// now, we've to check if we found a common translation, i.e. if there's a
									// translation vector that appeared three times, this must be the common translation
									for k, v := range translations {
										if len(v) >= 3 {
											// found a match, stop prco
											prevMatch[unmatchedScannerID] = matchedScannerID
											matched[Pair{matchedScannerID, unmatchedScannerID}] = MatchingPair{k, v[0].rotationID}
											nextNewlyMatchedScanners = append(nextNewlyMatchedScanners, unmatchedScannerID)
											delete(unmatched, unmatchedScannerID)
											alreadyMatched = true
											break
										}
									}
								}
							}
						}
					}
				}
			}
		}
		newlyMatchedScanners = nextNewlyMatchedScanners
	}

	return matched, prevMatch
}

func computeRotationSequence(scanners []Scanner, prevMatch []int) [][]int {
	requiredRotations := [][]int{}
	for i := 0; i < len(scanners); i++ {
		requiredRotations = append(requiredRotations, make([]int, 0))
	}
	for i := 1; i < len(scanners); i++ {
		for x := i; x != 0; x = prevMatch[x] {
			requiredRotations[i] = append(requiredRotations[i], x)
		}
		requiredRotations[i] = append(requiredRotations[i], 0)
	}
	return requiredRotations
}

func computeTranslationsFromOrigin(matched map[Pair]MatchingPair, rotationSequences [][]int) map[int]XYZ {
	translationsFromOrigin := map[int]XYZ{}
	translationsFromOrigin[0] = XYZ{0, 0, 0}
	for len(translationsFromOrigin) != len(matched)+1 {
		for k, v := range matched {
			if _, ok := translationsFromOrigin[k.b]; !ok {
				if q, ok := translationsFromOrigin[k.a]; ok {
					p := v.p
					if rots := rotationSequences[k.b]; len(rots) > 1 {
						for prv := 2; prv < len(rots); prv++ {
							rid := matched[Pair{rots[prv], rots[prv-1]}].rotationID
							p = rotate(p, rid)
						}
						translationsFromOrigin[k.b] = q.Sub(p)
					} else {
						// don't ask me why this has to be switched...
						translationsFromOrigin[k.b] = p.Sub(q)
					}
				}
			}
		}
	}
	return translationsFromOrigin
}

func partOne(lines []string) int {
	// read the input and compute fingerprints
	scanners := parseInput(lines)

	matched, prevMatch := matchScanners(scanners)

	requiredRotations := computeRotationSequence(scanners, prevMatch)

	// got all matches, now it's finally time to compute all the translations
	translationsFromOrigin := computeTranslationsFromOrigin(matched, requiredRotations)

	allBeacons := map[XYZ]int{}
	for scannerID := 0; scannerID < len(scanners); scannerID++ {
		for _, beacon := range scanners[scannerID].beacons {
			translated := XYZ{}
			if scannerID != 0 {
				p := beacon
				if rots := requiredRotations[scannerID]; len(rots) > 1 {
					for prv := 1; prv < len(rots); prv++ {
						rid := matched[Pair{rots[prv], rots[prv-1]}].rotationID
						p = rotate(p, rid)
					}
					translated = translationsFromOrigin[scannerID].Add(p)
				}
			} else {
				translated = beacon
			}
			allBeacons[translated]++
		}
	}

	return len(allBeacons)
}

func partTwo(lines []string) int {
	scanners := parseInput(lines)

	matched, prevMatch := matchScanners(scanners)

	requiredRotations := computeRotationSequence(scanners, prevMatch)

	// got all matches, now it's finally time to compute all the translations
	translationsFromOrigin := computeTranslationsFromOrigin(matched, requiredRotations)
	maxDistance := 0
	for i := range scanners {
		for j := range scanners {
			if i != j {
				val := translationsFromOrigin[i].l1distance(translationsFromOrigin[j])
				if val > maxDistance {
					maxDistance = val
				}
			}
		}
	}

	return maxDistance
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
