package main

var car = 4
var ped = 1
var bicycle = 2
var bus = 40
var train = 100

var dummyWL = WL{0, 0, 0}

// var lightGroups = [][]int{
// 	{4, 10, 23, 24, 33, 34},
// 	{2, 5, 6, 42},
// 	{2, 6, 27, 28, 37, 38, 42, 45},
// 	{31, 32, 33, 34, 35, 36, 37, 38},
// 	{21, 22, 23, 24, 25, 26, 27, 28},
// }

var nodes = []WL{
	dummyWL, // dummy node to ensure array indexing gets correct item
	WL{1, car, 0},
	WL{2, car, 0},
	WL{3, car, 0},
	WL{4, car, 0},
	WL{5, car, 0},
	WL{6, car, 0},
	WL{7, car, 0},
	WL{8, car, 0},
	WL{9, car, 0},
	WL{10, car, 0},
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	WL{21, ped, 0},
	WL{22, ped, 0},
	WL{23, ped, 0},
	WL{24, ped, 0},
	WL{25, ped, 0},
	WL{26, ped, 0},
	WL{27, ped, 0},
	WL{28, ped, 0},
	dummyWL,
	dummyWL,
	WL{31, bicycle, 0},
	WL{32, bicycle, 0},
	WL{33, bicycle, 0},
	WL{34, bicycle, 0},
	WL{35, bicycle, 0},
	WL{36, bicycle, 0},
	WL{37, bicycle, 0},
	WL{38, bicycle, 0},
	dummyWL,
	dummyWL,
	dummyWL,
	WL{42, bus, 0},
	dummyWL,
	dummyWL,
	WL{45, train, 0},
	WL{46, train, 0},
}

var newLightGroups = [][]WL{
	{
		nodes[4],
		nodes[10],
		nodes[23],
		nodes[24],
		nodes[33],
		nodes[34],
	},
	{
		nodes[2],
		nodes[5],
		nodes[6],
		nodes[42],
	},
	{
		nodes[2],
		nodes[6],
		nodes[27],
		nodes[28],
		nodes[37],
		nodes[38],
		nodes[42],
		nodes[45],
	},
	{
		nodes[31],
		nodes[32],
		nodes[33],
		nodes[34],
		nodes[35],
		nodes[36],
		nodes[37],
		nodes[38],
	},
	{
		nodes[21],
		nodes[22],
		nodes[23],
		nodes[24],
		nodes[25],
		nodes[26],
		nodes[27],
		nodes[28],
	},
}

// func initLights() {
// 	newLightGroups = make([][]WeightedLightNode,5)
// 	newLightGroups[0] = make([]WeightedLightNode,6)
// 	newLightGroups[0][0] = WeightedLightNode{4,car,0}
// 	newLightGroups[0][1] = WeightedLightNode{10,car,0}
// 	newLightGroups[0][2] = WeightedLightNode{23,ped,0}
// }
