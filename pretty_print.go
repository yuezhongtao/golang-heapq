package golang_heapq

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func PrettyPrint[T Ordered](tree []T) {
	lines, _, _, _ := buildTreeString(tree, 0, false, "-")
	fmt.Printf("\n%s\n", strings.Join(lines, "\n"))
}

func buildTreeString[T Ordered](tree []T, currIndex int, includeIndex bool, delimiter string) ([]string, int, int, int) {
	if len(tree) == 0 {
		return []string{}, 0, 0, 0
	}
	if currIndex > len(tree)-1 {
		return []string{}, 0, 0, 0
	}
	var line1 []string
	var line2 []string
	root := tree[currIndex]
	nodeRepr := str(root)
	if includeIndex {
		nodeRepr = fmt.Sprintf("%s%s%s", strconv.FormatInt(int64(currIndex), 10), delimiter, str(root))
	}
	newRootWidth := len(nodeRepr)
	gapSize := len(nodeRepr)

	lBox, lBoxWidth, lRootStart, lRootEnd := buildTreeString(tree, 2*currIndex+1, includeIndex, delimiter)
	rBox, rBoxWidth, rRootStart, rRootEnd := buildTreeString(tree, 2*currIndex+2, includeIndex, delimiter)

	//# Draw the branch connecting the current root node to the left sub-box
	//# Pad the line with whitespaces where necessary
	newRootStart := 0
	if lBoxWidth > 0 {
		lRoot := (lRootStart+lRootEnd)/2 + 1
		line1 = append(line1, strRepeat(" ", (lRoot+1)))
		line1 = append(line1, strRepeat("_", (lBoxWidth-lRoot)))
		line2 = append(line2, strRepeat(" ", lRoot)+"/")
		line2 = append(line2, strRepeat(" ", (lBoxWidth-lRoot)))
		newRootStart = lBoxWidth + 1
		gapSize += 1
	}

	// # Draw the representation of the current root node
	line1 = append(line1, nodeRepr)
	line2 = append(line2, strRepeat(" ", newRootWidth))

	if rBoxWidth > 0 {
		rRoot := (rRootStart + rRootEnd) / 2
		line1 = append(line1, strRepeat("_", rRoot))
		line1 = append(line1, strRepeat(" ", (rBoxWidth-rRoot+1)))
		line2 = append(line2, strRepeat(" ", rRoot)+"\\")
		line2 = append(line2, strRepeat(" ", (rBoxWidth-rRoot)))
		gapSize += 1
	}
	newRootEnd := newRootStart + newRootWidth - 1

	gap := strRepeat(" ", gapSize)
	newBox := []string{strings.Join(line1, ""), strings.Join(line2, "")}
	for i := 0; i < max(len(lBox), len(rBox)); i++ {
		lLine := ""
		rLine := ""
		if i < len(lBox) {
			lLine = lBox[i]
		} else {
			lLine = strRepeat(" ", lBoxWidth)
		}
		if i < len(rBox) {
			rLine = rBox[i]
		} else {
			rLine = strRepeat(" ", rBoxWidth)
		}
		newBox = append(newBox, fmt.Sprintf("%s%s%s", lLine, gap, rLine))
	}
	return newBox, len(newBox[0]), newRootStart, newRootEnd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func strRepeat(v string, t int) string {
	s := ""
	for i := 0; i < t; i++ {
		s = s + v
	}
	return s
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func str[T Ordered](item T) string {
	switch typ := any(item).(type) {
	case int, int8, int16, int32, int64:
		v := *(*int64)(unsafe.Pointer(&item))
		return strconv.FormatInt(v, 10)
	case uint, uint8, uint16, uint32, uint64:
		v := *(*uint64)(unsafe.Pointer(&item))
		return strconv.FormatUint(v, 10)
	case float32, float64:
		v := *(*float64)(unsafe.Pointer(&item))
		return strconv.FormatFloat(v, 'e', -1, 64)
	case complex64, complex128:
		v := *(*complex128)(unsafe.Pointer(&item))
		return strconv.FormatComplex(v, 'g', -1, 100)
	default:
		panic(fmt.Sprintf("not support type: %v", typ))
	}
}
