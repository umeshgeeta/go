// author: Umesh Patil
// March 2020

// The program builds a binary tree based on insertions done. The level
// where root resides is designated as '1', next at '2' and so forth.
// If height of the binary tree is 'h'; there will be h levels in the tree.
//
// It prints the binary tree on console as a pyramid. It computes
// spaces so that the pyramid is depicted as closely as possible.
package bt

import (
	"fmt"
	"math"
	"strconv"
)

// *************************************
// Struct & Constant definitions
// *************************************
type Bt struct {
	Nodes []int
	Height int
	MaxVal int			// maximum value encountered so that we know how many digits to provision per node
}

const TileSize = 4
const PadChar = " "
const TileChar = " "

// *************************************
// Constructor
// *************************************
func New(h int) *Bt {
	if h < 1 {
		fmt.Printf("Invalid Binary Tree Hieght, minimum value is 1.")
		return nil
	}
	bt := new(Bt)
	bt.Height = h
	// make sure that height is set before invoking nodeCount
	sz := bt.nodeCount()
	bt.Nodes = make([]int, sz)
	bt.MaxVal = math.MinInt64
	return bt
}

// *************************************
// Methods - Exposed
// *************************************

// Simply inserts the given value at the specified index.
func (bt* Bt) Insert(index int, value int){
	if index < 0 || index >= len(bt.Nodes) {
		fmt.Printf("Index %d is invalid for binary tree with height %d", index, bt.Height)
	}
	bt.Nodes[index] = value
	if bt.MaxVal < value {
		bt.MaxVal = value
	}
}

// Returns the pyramid string which can be printed.
// Gaps are called as tiles. Numbers are padded to maximum value so as alignment works.
func (bt Bt) Pyramid() string {
	result := ""
	tpr := tilesPerRow(bt.Height)
	for level := 1; level <= bt.Height; level++ {
		td := tileDist(tpr, level, bt.Height)
		result = result + bt.fillRow(level, td) + "\n"
	}
	return result
}

// *************************************
// Methods - Not Exposed
// *************************************

// Fills the row for the given level
func (bt Bt) fillRow(level int, td []int) string {
	result := ""
	nodesInLevel := numPerLevel(level)
	charsToAdjust := (bt.maxNodesInARow() - nodesInLevel) * bt.maxNodeLength()
	start, _ := indices(level)
	charsPerNode := charsToAdjust / nodesInLevel
	tdl := len(td)
	for k := 0; k < tdl-1; k++ {
		result += tileStr(td[k])
		result += bt.padNode(start+k, charsPerNode)
	}
	result += tileStr(td[tdl-1])
	return result
}

// Pad the value at a node so as it aligns to the value with maximum digits.
func (bt Bt) padNode(index int, additional int) string {
	return pad(bt.Nodes[index], bt.maxNodeLength(), additional)
}

func (bt Bt) maxNodeLength() int {
	return len(strconv.Itoa(bt.MaxVal))
}

func (bt Bt) nodeCount() int {
	return int(math.Pow(2, float64(bt.Height))) - 1
}

// Indicates how many node - values - can be maximum in a row,
// that row is the last row or the last level.
func (bt Bt) maxNodesInARow() int {
	return numPerLevel(bt.Height)
}

// *************************************
// Functions
// *************************************

// returns number of nodes for a given level
func numPerLevel(level int) int {
	return int(math.Pow(2, float64(level-1)))
}

// start and end indices of nodes in a given level
// end index is excluded, upto that value
func indices(l int) (int, int){
	start := int(math.Pow(2, float64(l-1))) - 1
	end := start + int(math.Pow(2, float64(l-1)))
	return start, end
}

// We pad the value to max. digit length, plus additionally
// more padding because the number of values at that level will
// be less than the maximum. For a binary tree with height 4,
// the last row / level will have 8 numbers which second last
// will have only 4 numbers. So the padding for those 4 additional
// numbers on the last row needs to be distributed among 4 numbers
// when we deal with padding of numbers on level 3.
func pad(num int, ps int, additional int) string {
	ns := strconv.Itoa(num)
	d := len(ns)
	result := ""
	totalChars := additional + ps - d
	half := totalChars / 2
	for j := 0; j < half; j++ {
		result += PadChar
	}
	result += ns
	for j := half; j < totalChars; j++ {
		result += PadChar
	}
	return result
}

func tileStr(tiles int) string {
	result := ""
	if tiles > 0 {
		for k := 0; k < tiles; k++ {
			result += singleTile()
		}
	}
	return result
}

func singleTile() string {
	result := ""
	for i := 0; i < TileSize; i++ {
		result += TileChar
	}
	return result
}

func tilesPerRow(height int) int {
	return (int(math.Pow(2, float64(height)))) - 2
}

// Method returns tile distribution for a row of the specified level.
// Key thing to understand is 'second row' in a binary tree of depth /
// height 4 will have different count of tiles compared to the same
// 'second row' in a binary tree of depth 8. That is because the last
// row will have more numbers for a taller binary tree and we want to
// start with 2 tile separation between 2 numbers at the bottom row.
func tileDist(tileNum int, level int, btHeight int) []int {
	tileSeg := int(math.Pow(2, float64(level-1))) + 1
	result := make([]int, tileSeg)
	boundary := int(math.Pow(2, float64(btHeight-level))) - 1
	result[0] = boundary
	result[tileSeg-1] = boundary
	tilesLeft := tileNum - (boundary * 2)
	if tilesLeft > 0 {
		perSeg := tilesLeft / (tileSeg - 2)
		for k := 1; k < tileSeg-1; k++ {
			result[k] = perSeg
		}
		leftOver := tilesLeft % (tileSeg - 2)
		if leftOver > 0 {
			middle := tileSeg / 2
			// for now add all to the middle segment
			// TODO - distribute around
			result[middle] += leftOver
		}
	}
	return result
}