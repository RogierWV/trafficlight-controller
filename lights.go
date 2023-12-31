package main

// weights
var car = 10
var ped = 1
var bicycle = 2
var bus = 40
var train = 1000

var dummyWL = WL{0, 0, 0}

var nodes = []WL{
	dummyWL, // dummy node to ensure array indexing gets correct item
	{1, car, 0},
	{2, car, 0},
	{3, car, 0},
	{4, car, 0},
	{5, car, 0},
	{6, car, 0},
	{7, car, 0},
	{8, car, 0},
	{9, car, 0},
	{10, car, 0},
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
	{21, ped, 0},
	{22, ped, 0},
	{23, ped, 0},
	{24, ped, 0},
	{25, ped, 0},
	{26, ped, 0},
	{27, ped, 0},
	{28, ped, 0},
	dummyWL,
	dummyWL,
	{31, bicycle, 0},
	{32, bicycle, 0},
	{33, bicycle, 0},
	{34, bicycle, 0},
	{35, bicycle, 0},
	{36, bicycle, 0},
	{37, bicycle, 0},
	{38, bicycle, 0},
	dummyWL,
	dummyWL,
	dummyWL,
	{42, bus, 0},
	dummyWL,
	dummyWL,
	{45, train, 0},
	{46, train, 0},
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
	dummyWL,
}

var newLightGroups = [][]WL{
	{
		nodes[4],
		nodes[23],
		nodes[24],
		nodes[33],
		nodes[34],
	},
	{
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
	{
		nodes[8],
		nodes[9],
		nodes[21],
		nodes[22],
		nodes[31],
		nodes[32],
	},
	{
		nodes[2],
		nodes[6],
		nodes[27],
		nodes[28],
		nodes[37],
		nodes[38],
		nodes[42],
		nodes[46],
	},
	{
		nodes[1],
		nodes[6],
		nodes[46],
	},
	{
		nodes[7],
		nodes[3],
		nodes[23],
		nodes[24],
		nodes[33],
		nodes[34],
	},
}
