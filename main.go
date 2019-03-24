package main

import "fmt"

func main() {
	start()
}

type element struct {
	atomicNumber           int
	symbol                 string
	element                string
	group                  int
	period                 int
	atomicWeight           float32
	density                float32
	meltingPoint           float32
	boilingPoint           float32
	heatCapacity           float32
	electronegativity      float32
	abundanceInEarthsCrust float32
}

func start() {
	var matter []element
	matter = append(matter, element{
		1,
		"H",
		"Hydrogen",
		1,
		1,
		1.008,
		0.00008988,
		14.01,
		20.28,
		14.304,
		2.20,
		1400.0,
	})

	matter = append(matter, element{
		2,
		"He",
		"Helium",
		18,
		1,
		4.002602,
		0.0001785,
		0,
		4.22,
		5.193,
		0,
		0.008,
	})

	matter = append(matter, element{
		3,         // int
		"Li",      // string
		"Lithium", // string
		1,         // int
		2,         // int
		9.0121831, // float32
		1.85,      // float32
		1560.0,    // float32
		2742.0,    // float32
		1.825,     // float32
		1.57,      // float32
		2.8,       // float32
	})

	matter = append(matter, element{
		4,           // int
		"Be",        // string
		"Beryllium", // string
		2,           // int
		2,           // int
		9.0121831,   // float32
		1.85,        // float32
		1560.0,      // float32
		2742.0,      // float32
		1.825,       // float32
		1.57,        // float32
		2.8,         // float32
	})

	matter = append(matter, element{
		5,       // int
		"B",     // string
		"Boron", // string
		13,      // int
		2,       // int
		10.81,   // float32
		2.34,    // float32
		2349.0,  // float32
		4200.0,  // float32
		1.026,   // float32
		2.04,    // float32
		10.0,    // float32
	})

	matter = append(matter, element{
		6,        // int
		"C",      // string
		"Carbon", // string
		14,       // int
		2,        // int
		12.011,   // float32
		2.267,    // float32
		3800,     // float32
		4300,     // float32
		0.709,    // float32
		2.55,     // float32
		200,      // float32
	})

	matter = append(matter, element{
		7,          // int
		"N",        // string
		"Nitrogen", // string
		15,         // int
		2,          // int
		14.007,     // float32
		0.0012506,  // float32
		63.15,      // float32
		77.36,      // float32
		1.04,       // float32
		3.04,       // float32
		19,         // float32
	})

	matter = append(matter, element{
		8,        // int
		"O",      // string
		"Oxygen", // string
		16,       // int
		2,        // int
		15.999,   // float32
		0.001429, // float32
		54.36,    // float32
		90.20,    // float32
		0.918,    // float32
		3.44,     // float32
		461000,   // float32
	})

	matter = append(matter, element{
		9,            // int
		"F",          // string
		"Fluorine",   // string
		17,           // int
		2,            // int
		18.998403163, // float32
		0.001696,     // float32
		53.53,        // float32
		85.03,        // float32
		0.824,        // float32
		3.98,         // float32
		585,          // float32
	})

	matter = append(matter, element{
		10,        // int
		"Ne",      // string
		"Neon",    // string
		18,        // int
		2,         // int
		20.1797,   // float32
		0.0008999, // float32
		24.56,     // float32
		27.07,     // float32
		1.03,      // float32
		0,         // float32
		0.005,     // float32
	})

	matter = append(matter, element{
		11,          // int
		"Na",        // string
		"Natrium",   // string
		1,           // int
		3,           // int
		22.98976928, // float32
		0.971,       // float32
		370.87,      // float32
		1156,        // float32
		1.228,       // float32
		0.93,        // float32
		23600,       // float32
	})

	matter = append(matter, element{
		12,          // int
		"Mg",        // string
		"Magnesium", // string
		2,           // int
		3,           // int
		24.305,      // float32
		1.738,       // float32
		923,         // float32
		1363,        // float32
		1.023,       // float32
		1.31,        // float32
		23300,       // float32
	})

	matter = append(matter, element{
		13,          // int
		"Al",        // string
		"Aluminium", // string
		13,          // int
		3,           // int
		26.9815384,  // float32
		2.698,       // float32
		933.47,      // float32
		2792,        // float32
		0.897,       // float32
		1.61,        // float32
		82300,       // float32
	})

	matter = append(matter, element{
		14,        // int
		"Si",      // string
		"Silicon", // string
		14,        // int
		3,         // int
		28.085,    // float32
		2.3296,    // float32
		1687,      // float32
		3538,      // float32
		0.705,     // float32
		1.9,       // float32
		282000,    // float32
	})

	matter = append(matter, element{
		15,           // int
		"P",          // string
		"Phosphorus", // string
		15,           // int
		3,            // int
		30.973761998, // float32
		1.82,         // float32
		317.30,       // float32
		550,          // float32
		0.769,        // float32
		2.19,         // float32
		1050,         // float32
	})

	matter = append(matter, element{
		16,        // int
		"S",       // string
		"Sulphur", // string
		16,        // int
		3,         // int
		32.06,     // float32
		2.067,     // float32
		388.36,    // float32
		717.87,    // float32
		0.71,      // float32
		2.58,      // float32
		350,       // float32
	})

	matter = append(matter, element{
		17,         // int
		"Cl",       // string
		"Chlorine", // string
		17,         // int
		3,          // int
		35.45,      // float32
		0.003214,   // float32
		171.6,      // float32
		239.11,     // float32
		0.479,      // float32
		3.16,       // float32
		145,        // float32
	})

	matter = append(matter, element{
		18,        // int
		"Ar",      // string
		"Argon",   // string
		18,        // int
		3,         // int
		39.948,    // float32
		0.0017837, // float32
		83.80,     // float32
		87.30,     // float32
		0.52,      // float32
		0,         // float32
		3.5,       // float32
	})

	matter = append(matter, element{
		19,       // int
		"K",      // string
		"Kalium", // string
		1,        // int
		4,        // int
		39.0983,  // float32
		0.862,    // float32
		336.53,   // float32
		1032,     // float32
		0.757,    // float32
		0.82,     // float32
		20900,    // float32
	})

	matter = append(matter, element{
		20,        // int
		"Ca",      // string
		"Calcium", // string
		2,         // int
		4,         // int
		40.078,    // float32
		1.54,      // float32
		1115,      // float32
		1757,      // float32
		0.647,     // float32
		1,         // float32
		41500,     // float32
	})

	matter = append(matter, element{
		21,         // int
		"Sc",       // string
		"Scandium", // string
		3,          // int
		4,          // int
		44.955908,  // float32
		2.989,      // float32
		1814,       // float32
		3109,       // float32
		0.568,      // float32
		1.36,       // float32
		22,         // float32
	})

	matter = append(matter, element{
		22,         // int
		"Ti",       // string
		"Titanium", // string
		4,          // int
		4,          // int
		47.867,     // float32
		4.54,       // float32
		1941,       // float32
		3560,       // float32
		0.523,      // float32
		1.54,       // float32
		5650,       // float32
	})

	matter = append(matter, element{
		23,         // int
		"V",        // string
		"Vanadium", // string
		5,          // int
		4,          // int
		50.9415,    // float32
		6.11,       // float32
		2183,       // float32
		3680,       // float32
		0.489,      // float32
		1.63,       // float32
		120,        // float32
	})

	matter = append(matter, element{
		24,         // int
		"Cr",       // string
		"Chromium", // string
		6,          // int
		4,          // int
		51.9961,    // float32
		7.15,       // float32
		2180,       // float32
		2944,       // float32
		0.449,      // float32
		1.66,       // float32
		102,        // float32
	})

	matter = append(matter, element{
		25,          // int
		"Mn",        // string
		"Manganese", // string
		7,           // int
		4,           // int
		54.938043,   // float32
		7.44,        // float32
		1519,        // float32
		2334,        // float32
		0.479,       // float32
		1.55,        // float32
		950,         // float32
	})

	matter = append(matter, element{
		26,       // int
		"Fe",     // string
		"Ferrum", // string
		8,        // int
		4,        // int
		55.845,   // float32
		7.874,    // float32
		1811,     // float32
		3134,     // float32
		0.449,    // float32
		1.83,     // float32
		56300,    // float32
	})

	matter = append(matter, element{
		27,        // int
		"Co",      // string
		"Cobalt",  // string
		9,         // int
		4,         // int
		58.933194, // float32
		8.86,      // float32
		1768,      // float32
		3200,      // float32
		0.421,     // float32
		1.88,      // float32
		25,        // float32
	})

	matter = append(matter, element{
		28,       // int
		"Ni",     // string
		"Nickel", // string
		10,       // int
		4,        // int
		58.6934,  // float32
		8.912,    // float32
		1728,     // float32
		3186,     // float32
		0.444,    // float32
		1.91,     // float32
		84,       // float32
	})

	matter = append(matter, element{
		29,       // int
		"Cu",     // string
		"Cuprum", // string
		11,       // int
		4,        // int
		63.546,   // float32
		8.96,     // float32
		1357.77,  // float32
		2835,     // float32
		0.385,    // float32
		1.9,      // float32
		60,       // float32
	})

	matter = append(matter, element{
		30,     // int
		"Zn",   // string
		"Zinc", // string
		12,     // int
		4,      // int
		65.38,  // float32
		7.134,  // float32
		692.88, // float32
		1180,   // float32
		0.388,  // float32
		1.65,   // float32
		70,     // float32
	})

	matter = append(matter, element{
		31,        // int
		"Ga",      // string
		"Gallium", // string
		13,        // int
		4,         // int
		69.723,    // float32
		5.907,     // float32
		302.9146,  // float32
		2673,      // float32
		0.371,     // float32
		1.81,      // float32
		19,        // float32
	})

	matter = append(matter, element{
		32,          // int
		"Ge",        // string
		"Germanium", // string
		14,          // int
		4,           // int
		69.723,      // float32
		5.907,       // float32
		302.9146,    // float32
		2673,        // float32
		0.371,       // float32
		1.81,        // float32
		19,          // float32
	})

	matter = append(matter, element{
		33,        // int
		"As",      // string
		"Arsenic", // string
		15,        // int
		4,         // int
		74.921595, // float32
		5.776,     // float32
		1090,      // float32
		887,       // float32
		0.329,     // float32
		2.18,      // float32
		1.8,       // float32
	})

	matter = append(matter, element{
		34,         // int
		"Se",       // string
		"Selenium", // string
		16,         // int
		4,          // int
		78.971,     // float32
		4.809,      // float32
		453,        // float32
		958,        // float32
		0.321,      // float32
		2.55,       // float32
		0.05,       // float32
	})

	matter = append(matter, element{
		35,        // int
		"Br",      // string
		"Bromine", // string
		17,        // int
		4,         // int
		79.904,    // float32
		3.122,     // float32
		265.8,     // float32
		332.0,     // float32
		0.474,     // float32
		2.96,      // float32
		2.4,       // float32
	})

	matter = append(matter, element{
		36,        // int
		"Kr",      // string
		"Krypton", // string
		18,        // int
		4,         // int
		83.798,    // float32
		0.003733,  // float32
		115.79,    // float32
		119.93,    // float32
		0.248,     // float32
		3,         // float32
		0.0001,    // float32
	})

	matter = append(matter, element{
		37,         // int
		"Rb",       // string
		"Rubidium", // string
		1,          // int
		5,          // int
		85.4678,    // float32
		1.532,      // float32
		312.46,     // float32
		961,        // float32
		0.363,      // float32
		0.82,       // float32
		90,         // float32
	})

	matter = append(matter, element{
		38,          // int
		"Sr",        // string
		"Strontium", // string
		2,           // int
		5,           // int
		87.62,       // float32
		2.64,        // float32
		1050,        // float32
		1655,        // float32
		0.301,       // float32
		0.95,        // float32
		370,         // float32
	})

	matter = append(matter, element{
		39,        // int
		"Y",       // string
		"Yttrium", // string
		3,         // int
		5,         // int
		88.90584,  // float32
		4.469,     // float32
		1799,      // float32
		3609,      // float32
		0.298,     // float32
		1.22,      // float32
		33,        // float32
	})

	matter = append(matter, element{
		40,          // int
		"Zr",        // string
		"Zirconium", // string
		4,           // int
		5,           // int
		91.224,      // float32
		6.506,       // float32
		2128,        // float32
		4682,        // float32
		0.278,       // float32
		1.33,        // float32
		165,         // float32
	})

	matter = append(matter, element{
		41,        // int
		"Nb",      // string
		"Niobium", // string
		5,         // int
		5,         // int
		92.90637,  // float32
		8.57,      // float32
		2750,      // float32
		5017,      // float32
		0.265,     // float32
		1.6,       // float32
		20,        // float32
	})

	matter = append(matter, element{
		42,           // int
		"Mo",         // string
		"Molybdenum", // string
		6,            // int
		5,            // int
		95.95,        // float32
		10.22,        // float32
		2896,         // float32
		4912,         // float32
		0.251,        // float32
		2.16,         // float32
		1.2,          // float32
	})

	matter = append(matter, element{
		43,           // int
		"Tc",         // string
		"Technetium", // string
		7,            // int
		5,            // int
		98,           // float32
		11.5,         // float32
		2430,         // float32
		4538,         // float32
		0,            // float32
		1.9,          // float32
		0.000000003,  // float32
	})

	matter = append(matter, element{
		44,          // int
		"Ru",        // string
		"Ruthenium", // string
		8,           // int
		5,           // int
		101.07,      // float32
		12.37,       // float32
		2607,        // float32
		4423,        // float32
		0.238,       // float32
		2.2,         // float32
		0.001,       // float32
	})

	matter = append(matter, element{
		45,        // int
		"Rh",      // string
		"Rhodium", // string
		9,         // int
		5,         // int
		102.90549, // float32
		12.41,     // float32
		2237,      // float32
		3968,      // float32
		0.243,     // float32
		2.28,      // float32
		0.001,     // float32
	})

	matter = append(matter, element{
		46,          // int
		"Pd",        // string
		"Palladium", // string
		10,          // int
		5,           // int
		106.42,      // float32
		12.02,       // float32
		1828.05,     // float32
		3236,        // float32
		0.244,       // float32
		2.2,         // float32
		0.015,       // float32
	})

	matter = append(matter, element{
		47,         // int
		"Ag",       // string
		"Argentum", // string
		11,         // int
		5,          // int
		107.8682,   // float32
		10.501,     // float32
		1234.93,    // float32
		2435,       // float32
		0.235,      // float32
		1.93,       // float32
		0.075,      // float32
	})

	matter = append(matter, element{
		48,        // int
		"Cd",      // string
		"Cadmium", // string
		12,        // int
		5,         // int
		112.414,   // float32
		8.69,      // float32
		594.22,    // float32
		1040,      // float32
		0.232,     // float32
		1.69,      // float32
		0.159,     // float32
	})

	matter = append(matter, element{
		49,       // int
		"In",     // string
		"Indium", // string
		13,       // int
		5,        // int
		114.818,  // float32
		7.31,     // float32
		429.75,   // float32
		2345,     // float32
		0.233,    // float32
		1.78,     // float32
		0.25,     // float32
	})

	matter = append(matter, element{
		50,        // int
		"Sn",      // string
		"Stannum", // string
		14,        // int
		5,         // int
		118.710,   // float32
		7.287,     // float32
		505.08,    // float32
		2875,      // float32
		0.228,     // float32
		1.96,      // float32
		2.3,       // float32
	})

	matter = append(matter, element{
		51,        // int
		"Sb",      // string
		"Stibium", // string
		15,        // int
		5,         // int
		121.760,   // float32
		6.685,     // float32
		903.78,    // float32
		1860,      // float32
		0.207,     // float32
		2.05,      // float32
		0.2,       // float32
	})

	matter = append(matter, element{
		52,          // int
		"Te",        // string
		"Tellurium", // string
		16,          // int
		5,           // int
		127.60,      // float32
		6.232,       // float32
		722.66,      // float32
		1261,        // float32
		0.202,       // float32
		2.1,         // float32
		0.001,       // float32
	})

	matter = append(matter, element{
		53,        // int
		"I",       // string
		"Iodine",  // string
		17,        // int
		5,         // int
		126.90447, // float32
		4.93,      // float32
		386.85,    // float32
		457.4,     // float32
		0.214,     // float32
		2.66,      // float32
		0.45,      // float32
	})

	matter = append(matter, element{
		54,       // int
		"Xe",     // string
		"Xenon",  // string
		18,       // int
		5,        // int
		131.293,  // float32
		0.005887, // float32
		161.4,    // float32
		165.03,   // float32
		0.158,    // float32
		2.6,      // float32
		0.00003,  // float32
	})

	matter = append(matter, element{
		55,           // int
		"Cs",         // string
		"Caesium",    // string
		1,            // int
		6,            // int
		132.90545196, // float32
		1.873,        // float32
		301.59,       // float32
		944,          // float32
		0.242,        // float32
		0.79,         // float32
		3,            // float32
	})

	matter = append(matter, element{
		56,       // int
		"Ba",     // string
		"Barium", // string
		2,        // int
		6,        // int
		137.327,  // float32
		3.594,    // float32
		1000,     // float32
		2170,     // float32
		0.204,    // float32
		0.89,     // float32
		425,      // float32
	})

	matter = append(matter, element{
		57,          // int
		"La",        // string
		"Lanthanum", // string
		3,           // int
		6,           // int
		138.90547,   // float32
		6.145,       // float32
		1193,        // float32
		3737,        // float32
		0.195,       // float32
		1.1,         // float32
		39,          // float32
	})

	matter = append(matter, element{
		58,       // int
		"Ce",     // string
		"Cerium", // string
		0,        // int
		6,        // int
		140.116,  // float32
		6.77,     // float32
		1068,     // float32
		3716,     // float32
		0.192,    // float32
		1.12,     // float32
		66.5,     // float32
	})

	matter = append(matter, element{
		59,             // int
		"Pr",           // string
		"Praseodymium", // string
		0,              // int
		6,              // int
		140.90766,      // float32
		6.773,          // float32
		1208,           // float32
		3793,           // float32
		0.193,          // float32
		1.13,           // float32
		9.2,            // float32
	})

	matter = append(matter, element{
		60,          // int
		"Nd",        // string
		"Neodymium", // string
		0,           // int
		6,           // int
		144.242,     // float32
		7.007,       // float32
		1297,        // float32
		3347,        // float32
		0.19,        // float32
		1.14,        // float32
		41.5,        // float32
	})

	matter = append(matter, element{
		61,                    // int
		"Pm",                  // string
		"Promethium",          // string
		0,                     // int
		6,                     // int
		145,                   // float32
		7.26,                  // float32
		1315,                  // float32
		3273,                  // float32
		0,                     // float32
		1.13,                  // float32
		0.0000000000000000002, // float32
	})

	matter = append(matter, element{
		62,         // int
		"Sm",       // string
		"Samarium", // string
		0,          // int
		6,          // int
		150.36,     // float32
		7.52,       // float32
		1345,       // float32
		2067,       // float32
		0.197,      // float32
		1.17,       // float32
		7.05,       // float32
	})

	matter = append(matter, element{
		63,         // int
		"Eu",       // string
		"Europium", // string
		0,          // int
		6,          // int
		151.964,    // float32
		5.243,      // float32
		1099,       // float32
		1802,       // float32
		0.182,      // float32
		1.2,        // float32
		2,          // float32
	})

	matter = append(matter, element{
		64,           // int
		"Gd",         // string
		"Gadolinium", // string
		0,            // int
		6,            // int
		157.25,       // float32
		7.895,        // float32
		1585,         // float32
		3546,         // float32
		0.236,        // float32
		1.2,          // float32
		6.2,          // float32
	})

	matter = append(matter, element{
		65,         // int
		"Tb",       // string
		"Terbium",  // string
		0,          // int
		6,          // int
		158.925354, // float32
		8.229,      // float32
		1629,       // float32
		3503,       // float32
		0.182,      // float32
		1.2,        // float32
		1.2,        // float32
	})

	matter = append(matter, element{
		66,           // int
		"Dy",         // string
		"Dysprosium", // string
		0,            // int
		6,            // int
		162.500,      // float32
		8.55,         // float32
		1680,         // float32
		2840,         // float32
		0.17,         // float32
		1.22,         // float32
		5.2,          // float32
	})

	matter = append(matter, element{
		67,         // int
		"Ho",       // string
		"Holmium",  // string
		0,          // int
		6,          // int
		164.930328, // float32
		8.795,      // float32
		1734,       // float32
		2993,       // float32
		0.165,      // float32
		1.23,       // float32
		1.3,        // float32
	})

	matter = append(matter, element{
		68,       // int
		"Er",     // string
		"Erbium", // string
		0,        // int
		6,        // int
		167.259,  // float32
		9.066,    // float32
		1802,     // float32
		3141,     // float32
		0.168,    // float32
		1.24,     // float32
		3.5,      // float32
	})

	matter = append(matter, element{
		69,         // int
		"Tm",       // string
		"Thulium",  // string
		0,          // int
		6,          // int
		168.934218, // float32
		9.321,      // float32
		1818,       // float32
		2223,       // float32
		0.16,       // float32
		1.25,       // float32
		0.52,       // float32
	})

	matter = append(matter, element{
		70,          // int
		"Yb",        // string
		"Ytterbium", // string
		0,           // int
		6,           // int
		173.045,     // float32
		6.965,       // float32
		1097,        // float32
		1469,        // float32
		0.155,       // float32
		1.1,         // float32
		3.2,         // float32
	})

	matter = append(matter, element{
		71,         // int
		"Lu",       // string
		"Lutetium", // string
		0,          // int
		6,          // int
		174.9668,   // float32
		9.84,       // float32
		1925,       // float32
		3675,       // float32
		0.154,      // float32
		1.27,       // float32
		0.8,        // float32
	})

	matter = append(matter, element{
		72,        // int
		"Hf",      // string
		"Hafnium", // string
		4,         // int
		6,         // int
		178.49,    // float32
		13.31,     // float32
		2506,      // float32
		4876,      // float32
		0.144,     // float32
		1.3,       // float32
		3,         // float32
	})

	matter = append(matter, element{
		73,         // int
		"Ta",       // string
		"Tantalum", // string
		5,          // int
		6,          // int
		180.94788,  // float32
		16.654,     // float32
		3290,       // float32
		5731,       // float32
		0.14,       // float32
		1.5,        // float32
		2,          // float32
	})

	matter = append(matter, element{
		74,        // int
		"W",       // string
		"Wolfram", // string
		6,         // int
		6,         // int
		183.84,    // float32
		19.25,     // float32
		3695,      // float32
		5828,      // float32
		0.132,     // float32
		2.36,      // float32
		1.3,       // float32
	})

	// Re-check the name, maybe latin is more correct?
	matter = append(matter, element{
		75,        // int
		"Re",      // string
		"Rhenium", // string
		7,         // int
		6,         // int
		186.207,   // float32
		21.02,     // float32
		3459,      // float32
		5869,      // float32
		0.137,     // float32
		1.9,       // float32
		0.0007,    // float32
	})

	matter = append(matter, element{
		76,       // int
		"Os",     // string
		"Osmium", // string
		8,        // int
		6,        // int
		190.23,   // float32
		22.61,    // float32
		3306,     // float32
		5285,     // float32
		0.13,     // float32
		2.2,      // float32
		0.002,    // float32
	})

	matter = append(matter, element{
		77,        // int
		"Ir",      // string
		"Iridium", // string
		9,         // int
		6,         // int
		192.217,   // float32
		22.56,     // float32
		2719,      // float32
		4701,      // float32
		0.131,     // float32
		2.2,       // float32
		0.001,     // float32
	})

	matter = append(matter, element{
		78,         // int
		"Pt",       // string
		"Platinum", // string
		10,         // int
		6,          // int
		195.084,    // float32
		21.46,      // float32
		2041.4,     // float32
		4098,       // float32
		0.133,      // float32
		2.28,       // float32
		0.005,      // float32
	})

	matter = append(matter, element{
		79,         // int
		"Au",       // string
		"Aurum",    // string
		11,         // int
		6,          // int
		196.966570, // float32
		19.282,     // float32
		1337.33,    // float32
		3129,       // float32
		0.129,      // float32
		2.54,       // float32
		0.004,      // float32
	})

	matter = append(matter, element{
		80,            // int
		"Hg",          // string
		"Hydrargyrum", // string
		12,            // int
		6,             // int
		200.592,       // float32
		13.5336,       // float32
		234.43,        // float32
		629.88,        // float32
		0.14,          // float32
		2,             // float32
		0.085,         // float32
	})

	matter = append(matter, element{
		81,         // int
		"Tl",       // string
		"Thallium", // string
		13,         // int
		6,          // int
		204.38,     // float32
		11.85,      // float32
		577,        // float32
		1746,       // float32
		0.129,      // float32
		1.62,       // float32
		0.85,       // float32
	})

	matter = append(matter, element{
		82,        // int
		"Pb",      // string
		"Plumbum", // string
		14,        // int
		6,         // int
		207.2,     // float32
		11.342,    // float32
		600.61,    // float32
		2022,      // float32
		0.129,     // float32
		1.87,      // float32
		14,        // float32
	})

	matter = append(matter, element{
		83,        // int
		"Bi",      // string
		"Bismuth", // string
		15,        // int
		6,         // int
		208.98040, // float32
		9.807,     // float32
		544.7,     // float32
		1837,      // float32
		0.122,     // float32
		2.02,      // float32
		0.009,     // float32
	})

	matter = append(matter, element{
		84,          // int
		"Po",        // string
		"Polonium",  // string
		16,          // int
		6,           // int
		209,         // float32
		9.32,        // float32
		527,         // float32
		1235,        // float32
		0,           // float32
		2.0,         // float32
		0.000000002, // float32
	})

	matter = append(matter, element{
		85,                     // int
		"At",                   // string
		"Astatine",             // string
		17,                     // int
		6,                      // int
		210,                    // float32
		7,                      // float32
		575,                    // float32
		610,                    // float32
		0,                      // float32
		2.2,                    // float32
		0.00000000000000000003, // float32
	})

	matter = append(matter, element{
		86,              // int
		"Rn",            // string
		"Radon",         // string
		18,              // int
		6,               // int
		222,             // float32
		0.00973,         // float32
		202,             // float32
		211.3,           // float32
		0.094,           // float32
		2.2,             // float32
		0.0000000000004, // float32
	})

	matter = append(matter, element{
		87,                   // int
		"Fr",                 // string
		"Francium",           // string
		1,                    // int
		7,                    // int
		223,                  // float32
		1.87,                 // float32
		300,                  // float32
		950,                  // float32
		0,                    // float32
		0.7,                  // float32
		0.000000000000000001, // float32
	})

	matter = append(matter, element{
		88,        // int
		"Ra",      // string
		"Radium",  // string
		2,         // int
		7,         // int
		226,       // float32
		5.5,       // float32
		973,       // float32
		2010,      // float32
		0.094,     // float32
		0.9,       // float32
		0.0000009, // float32
	})

	matter = append(matter, element{
		89,           // int
		"Ac",         // string
		"Actinium",   // string
		3,            // int
		7,            // int
		227,          // float32
		10.07,        // float32
		1323,         // float32
		3471,         // float32
		0.12,         // float32
		1.1,          // float32
		0.0000000055, // float32
	})

	matter = append(matter, element{
		90,        // int
		"Th",      // string
		"Thorium", // string
		0,         // int
		7,         // int
		232.0377,  // float32
		11.72,     // float32
		2115,      // float32
		5061,      // float32
		0.113,     // float32
		1.3,       // float32
		9.6,       // float32
	})

	matter = append(matter, element{
		91,             // int
		"Pa",           // string
		"Protactinium", // string
		0,              // int
		7,              // int
		231.03588,      // float32
		15.37,          // float32
		1841,           // float32
		4300,           // float32
		0,              // float32
		1.5,            // float32
		0.000014,       // float32
	})

	matter = append(matter, element{
		92,        // int
		"U",       // string
		"Uranium", // string
		0,         // int
		7,         // int
		238.02891, // float32
		18.95,     // float32
		1405.3,    // float32
		4404,      // float32
		0.116,     // float32
		1.38,      // float32
		2.7,       // float32
	})

	matter = append(matter, element{
		93,             // int
		"Np",           // string
		"Neptunium",    // string
		0,              // int
		7,              // int
		237,            // float32
		20.45,          // float32
		917,            // float32
		4273,           // float32
		0,              // float32
		1.36,           // float32
		0.000000000003, // float32
	})

	matter = append(matter, element{
		94,            // int
		"Pu",          // string
		"Plutonium",   // string
		0,             // int
		7,             // int
		244,           // float32
		19.84,         // float32
		912.5,         // float32
		3501,          // float32
		0,             // float32
		1.28,          // float32
		0.00000000003, // float32
	})

	matter = append(matter, element{
		95,          // int
		"Am",        // string
		"Americium", // string
		0,           // int
		7,           // int
		243,         // float32
		13.69,       // float32
		1449,        // float32
		2880,        // float32
		0,           // float32
		1.13,        // float32
		0,           // float32
	})

	matter = append(matter, element{
		96,       // int
		"Cm",     // string
		"Curium", // string
		0,        // int
		7,        // int
		247,      // float32
		13.51,    // float32
		1613,     // float32
		3383,     // float32
		0,        // float32
		1.28,     // float32
		0,        // float32
	})

	matter = append(matter, element{
		97,          // int
		"Bk",        // string
		"Berkelium", // string
		0,           // int
		7,           // int
		247,         // float32
		14.79,       // float32
		1259,        // float32
		2900,        // float32
		0,           // float32
		1.3,         // float32
		0,           // float32
	})

	matter = append(matter, element{
		98,            // int
		"Cf",          // string
		"Californium", // string
		0,             // int
		7,             // int
		251,           // float32
		15.1,          // float32
		1173,          // float32
		1743,          // float32
		0,             // float32
		1.3,           // float32
		0,             // float32
	})

	matter = append(matter, element{
		99,            // int
		"Es",          // string
		"Einsteinium", // string
		0,             // int
		7,             // int
		252,           // float32
		8.84,          // float32
		1133,          // float32
		1269,          // float32
		0,             // float32
		1.3,           // float32
		0,             // float32
	})

	matter = append(matter, element{
		100,       // int
		"Fm",      // string
		"Fermium", // string
		0,         // int
		7,         // int
		257,       // float32
		9.7,       // float32
		1125,      // float32
		0,         // float32
		0,         // float32
		1.3,       // float32
		0,         // float32
	})

	matter = append(matter, element{
		101,           // int
		"Md",          // string
		"Mendelevium", // string
		0,             // int
		7,             // int
		258,           // float32
		10.3,          // float32
		1100,          // float32
		0,             // float32
		0,             // float32
		1.3,           // float32
		0,             // float32
	})

	matter = append(matter, element{
		102,        // int
		"No",       // string
		"Nobelium", // string
		0,          // int
		7,          // int
		259,        // float32
		9.9,        // float32
		1100,       // float32
		0,          // float32
		0,          // float32
		1.3,        // float32
		0,          // float32
	})

	matter = append(matter, element{
		103,          // int
		"Lr",         // string
		"Lawrencium", // string
		0,            // int
		7,            // int
		266,          // float32
		15.6,         // float32
		1900,         // float32
		0,            // float32
		0,            // float32
		1.3,          // float32
		0,            // float32
	})

	matter = append(matter, element{
		104,             // int
		"Rf",            // string
		"Rutherfordium", // string
		4,               // int
		7,               // int
		267,             // float32
		23.2,            // float32
		2400,            // float32
		5800,            // float32
		0,               // float32
		0,               // float32
		0,               // float32
	})

	matter = append(matter, element{
		105,       // int
		"Db",      // string
		"Dubnium", // string
		5,         // int
		7,         // int
		268,       // float32
		29.3,      // float32
		0,         // float32
		0,         // float32
		0,         // float32
		0,         // float32
		0,         // float32
	})

	matter = append(matter, element{
		106,          // int
		"Sg",         // string
		"Seaborgium", // string
		6,            // int
		7,            // int
		269,          // float32
		35.0,         // float32
		0,            // float32
		0,            // float32
		0,            // float32
		0,            // float32
		0,            // float32
	})

	matter = append(matter, element{
		107,       // int
		"Bh",      // string
		"Bohrium", // string
		7,         // int
		7,         // int
		270,       // float32
		37.1,      // float32
		0,         // float32
		0,         // float32
		0,         // float32
		0,         // float32
		0,         // float32
	})

	matter = append(matter, element{
		108,       // int
		"Hs",      // string
		"Hassium", // string
		8,         // int
		7,         // int
		270,       // float32
		40.7,      // float32
		0,         // float32
		0,         // float32
		0,         // float32
		0,         // float32
		0,         // float32
	})

	matter = append(matter, element{
		109,          // int
		"Mt",         // string
		"Meitherium", // string
		9,            // int
		7,            // int
		278,          // float32
		37.4,         // float32
		0,            // float32
		0,            // float32
		0,            // float32
		0,            // float32
		0,            // float32
	})

	matter = append(matter, element{
		110,            // int
		"Ds",           // string
		"Darmstadtium", // string
		10,             // int
		7,              // int
		281,            // float32
		34.8,           // float32
		0,              // float32
		0,              // float32
		0,              // float32
		0,              // float32
		0,              // float32
	})

	matter = append(matter, element{
		111,           // int
		"Rg",          // string
		"Roentgenium", // string
		11,            // int
		7,             // int
		282,           // float32
		28.7,          // float32
		0,             // float32
		0,             // float32
		0,             // float32
		0,             // float32
		0,             // float32
	})

	matter = append(matter, element{
		112,           // int
		"Cn",          // string
		"Copernicium", // string
		12,            // int
		7,             // int
		285,           // float32
		23.7,          // float32
		0,             // float32
		357,           // float32
		0,             // float32
		0,             // float32
		0,             // float32
	})

	matter = append(matter, element{
		113,        // int
		"Nh",       // string
		"Nihonium", // string
		13,         // int
		7,          // int
		286,        // float32
		16,         // float32
		700,        // float32
		1400,       // float32
		0,          // float32
		0,          // float32
		0,          // float32
	})

	matter = append(matter, element{
		114,         // int
		"Fl",        // string
		"Flerovium", // string
		14,          // int
		7,           // int
		289,         // float32
		14,          // float32
		0,           // float32
		210,         // float32
		0,           // float32
		0,           // float32
		0,           // float32
	})

	matter = append(matter, element{
		115,         // int
		"Mc",        // string
		"Moscovium", // string
		15,          // int
		7,           // int
		290,         // float32
		13.5,        // float32
		700,         // float32
		1400,        // float32
		0,           // float32
		0,           // float32
		0,           // float32
	})

	matter = append(matter, element{
		116,           // int
		"Lv",          // string
		"Livermorium", // string
		16,            // int
		7,             // int
		293,           // float32
		12.9,          // float32
		709,           // float32
		1085,          // float32
		0,             // float32
		0,             // float32
		0,             // float32
	})

	matter = append(matter, element{
		117,          // int
		"Ts",         // string
		"Tennessine", // string
		17,           // int
		7,            // int
		294,          // float32
		7.2,          // float32
		723,          // float32
		883,          // float32
		0,            // float32
		0,            // float32
		0,            // float32
	})

	matter = append(matter, element{
		118,         // int
		"Og",        // string
		"Oganesson", // string
		18,          // int
		7,           // int
		294,         // float32
		5.0,         // float32
		0,           // float32
		350,         // float32
		0,           // float32
		0,           // float32
		0,           // float32
	})

	for i := 0; i < len(matter); i++ {
		fmt.Println(matter[i])
	}
}
