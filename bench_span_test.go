package safe

import (
	"math"
)

func benchSpanAdd() ([]int8, []int8) {
	span := int8Full()
	return span, span
}

func benchSpanAddU() ([]uint8, []uint8) {
	span := uint8Full()
	return span, span
}

func benchSpanAdd3() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func benchSpanAdd3U() ([]uint8, []uint8, []uint8) {
	span := uint8Full()
	return span, span, span
}

// Six result arguments are used for the following reasons:
//
// 1. To use in the benchmark that part of the function code that sorts
// the input arguments and increase the impact of this sorting on the benchmark results
//
// 2. To reduce the chances of using in the benchmark a number of nested loops
// different from the intended one
//
//nolint:gocritic // The reasons are stated above.
func benchSpanAddM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := int8Short()
	return span, span, span, span, span, span
}

//nolint:gocritic // For reasons, see the description of the [benchSpanAddM] function.
func benchSpanAddMU() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := uint8Short()
	return span, span, span, span, span, span
}

func benchSpanSub() ([]int8, []int8) {
	span := int8Full()
	return span, span
}

func benchSpanSubU() ([]uint8, []uint8) {
	span := uint8Full()
	return span, span
}

func benchSpanSub3() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func benchSpanSub3U() ([]uint8, []uint8, []uint8) {
	span := uint8Full()
	return span, span, span
}

//nolint:gocritic // For reasons, see the description of the [benchSpanAddM] function.
func benchSpanSubM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := int8Short()
	return span, span, span, span, span, span
}

//nolint:gocritic // For reasons, see the description of the [benchSpanAddM] function.
func benchSpanSubMU() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := uint8Short()
	return span, span, span, span, span, span
}

func benchSpanMul() ([]int8, []int8) {
	span := int8Full()
	return span, span
}

func benchSpanMulU() ([]uint8, []uint8) {
	span := uint8Full()
	return span, span
}

func benchSpanMul3() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func benchSpanMul3U() ([]uint8, []uint8, []uint8) {
	span := uint8Full()
	return span, span, span
}

//nolint:gocritic // For reasons, see the description of the [benchSpanAddM] function.
func benchSpanMulM() ([]int8, []int8, []int8, []int8, []int8, []int8) {
	span := int8Short()
	return span, span, span, span, span, span
}

//nolint:gocritic // For reasons, see the description of the [benchSpanAddM] function.
func benchSpanMulMU() ([]uint8, []uint8, []uint8, []uint8, []uint8, []uint8) {
	span := uint8Short()
	return span, span, span, span, span, span
}

func benchSpanDiv() ([]int8, []int8) {
	span := int8Full()
	return span, span
}

func benchSpanDivM() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func benchSpanNegate() ([]int8, []uint8) {
	signed := int8Full()
	unsigned := uint8Full()

	return signed, unsigned
}

func benchSpanIToI() ([]int8, []uint8, []uint16) {
	s8 := int8Full()
	u8 := uint8Full()
	u16 := uint16Short()

	return s8, u8, u16
}

func benchSpanIToF() []int64 {
	return []int64{1, 2, 9007199254740993, 9007199254740995}
}

func benchSpanFToI() []float64 {
	span := []float64{
		-18446744073709551617,
		-18446744073709551616,
		1,
		2,
		3,
		4,
		18446744073709551616,
		18446744073709551617,
		math.NaN(),
		math.NaN(),
		math.NaN(),
		math.NaN(),
	}

	return span
}

func benchSpanAddSub() ([]uint8, []uint8, []uint8) {
	span := uint8Full()
	return span, span, span
}

func benchSpanAddDiv() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func benchSpanAddDivU() ([]uint8, []uint8, []uint8) {
	span := uint8Full()
	return span, span, span
}

func benchSpanSubDiv() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func benchSpanSubDivU() ([]uint8, []uint8, []uint8) {
	span := uint8Full()
	return span, span, span
}

func benchSpanShift() ([]int8, []int8) {
	span := int8Full()
	return span, span
}

func benchSpanAddSubDiv() ([]int8, []int8, []int8, []int8) {
	span := int8Short()
	return span, span, span, span
}

func benchSpanAddOneSubDiv() ([]int8, []int8, []int8) {
	span := int8Full()
	return span, span, span
}

func int8Full() []int8 {
	span := []int8{
		-128, -127, -126, -125, -124, -123, -122, -121, -120,
		-119, -118, -117, -116, -115, -114, -113, -112, -111, -110,
		-109, -108, -107, -106, -105, -104, -103, -102, -101, -100,
		-99, -98, -97, -96, -95, -94, -93, -92, -91, -90,
		-89, -88, -87, -86, -85, -84, -83, -82, -81, -80,
		-79, -78, -77, -76, -75, -74, -73, -72, -71, -70,
		-69, -68, -67, -66, -65, -64, -63, -62, -61, -60,
		-59, -58, -57, -56, -55, -54, -53, -52, -51, -50,
		-49, -48, -47, -46, -45, -44, -43, -42, -41, -40,
		-39, -38, -37, -36, -35, -34, -33, -32, -31, -30,
		-29, -28, -27, -26, -25, -24, -23, -22, -21, -20,
		-19, -18, -17, -16, -15, -14, -13, -12, -11, -10,
		-9, -8, -7, -6, -5, -4, -3, -2, -1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
		110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
		120, 121, 122, 123, 124, 125, 126, 127,
	}

	return span
}

func int8Short() []int8 {
	span := []int8{
		-128, -127, -126, -125, -124,
		-5, -4, -3, -2, -1,
		0,
		1, 2, 3, 4, 5,
		123, 124, 125, 126, 127,
	}

	return span
}

func uint8Full() []uint8 {
	span := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
		110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
		120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
		130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
		140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
		150, 151, 152, 153, 154, 155, 156, 157, 158, 159,
		160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
		170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
		180, 181, 182, 183, 184, 185, 186, 187, 188, 189,
		190, 191, 192, 193, 194, 195, 196, 197, 198, 199,
		200, 201, 202, 203, 204, 205, 206, 207, 208, 209,
		210, 211, 212, 213, 214, 215, 216, 217, 218, 219,
		220, 221, 222, 223, 224, 225, 226, 227, 228, 229,
		230, 231, 232, 233, 234, 235, 236, 237, 238, 239,
		240, 241, 242, 243, 244, 245, 246, 247, 248, 249,
		250, 251, 252, 253, 254, 255,
	}

	return span
}

func uint8Short() []uint8 {
	span := []uint8{
		128, 129, 130, 131, 132,
		251, 252, 253, 254, 255,
		0,
		1, 2, 3, 4, 5,
		123, 124, 125, 126, 127,
	}

	return span
}

func uint16Short() []uint16 {
	span := []uint16{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
		110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
		120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
		130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
		140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
		150, 151, 152, 153, 154, 155, 156, 157, 158, 159,
		160, 161, 162, 163, 164, 165, 166, 167, 168, 169,
		170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
		180, 181, 182, 183, 184, 185, 186, 187, 188, 189,
		190, 191, 192, 193, 194, 195, 196, 197, 198, 199,
		200, 201, 202, 203, 204, 205, 206, 207, 208, 209,
		210, 211, 212, 213, 214, 215, 216, 217, 218, 219,
		220, 221, 222, 223, 224, 225, 226, 227, 228, 229,
		230, 231, 232, 233, 234, 235, 236, 237, 238, 239,
		240, 241, 242, 243, 244, 245, 246, 247, 248, 249,
		250, 251, 252, 253, 254, 255, 256, 257, 258, 259,
		260, 261, 262, 263, 264, 265, 266, 267, 268, 269,
		270, 271, 272, 273, 274, 275, 276, 277, 278, 279,
		280, 281, 282, 283, 284, 285, 286, 287, 288, 289,
		290, 291, 292, 293, 294, 295, 296, 297, 298, 299,
		300, 301, 302, 303, 304, 305, 306, 307, 308, 309,
		310, 311, 312, 313, 314, 315, 316, 317, 318, 319,
		320, 321, 322, 323, 324, 325, 326, 327, 328, 329,
		330, 331, 332, 333, 334, 335, 336, 337, 338, 339,
		340, 341, 342, 343, 344, 345, 346, 347, 348, 349,
		350, 351, 352, 353, 354, 355, 356, 357, 358, 359,
		360, 361, 362, 363, 364, 365, 366, 367, 368, 369,
		370, 371, 372, 373, 374, 375, 376, 377, 378, 379,
		380, 381, 382, 383, 384, 385, 386, 387, 388, 389,
		390, 391, 392, 393, 394, 395, 396, 397, 398, 399,
		400, 401, 402, 403, 404, 405, 406, 407, 408, 409,
		410, 411, 412, 413, 414, 415, 416, 417, 418, 419,
		420, 421, 422, 423, 424, 425, 426, 427, 428, 429,
		430, 431, 432, 433, 434, 435, 436, 437, 438, 439,
		440, 441, 442, 443, 444, 445, 446, 447, 448, 449,
		450, 451, 452, 453, 454, 455, 456, 457, 458, 459,
		460, 461, 462, 463, 464, 465, 466, 467, 468, 469,
		470, 471, 472, 473, 474, 475, 476, 477, 478, 479,
		480, 481, 482, 483, 484, 485, 486, 487, 488, 489,
		490, 491, 492, 493, 494, 495, 496, 497, 498, 499,
		500, 501, 502, 503, 504, 505, 506, 507, 508, 509,
		510, 511,
		65024, 65025, 65026, 65027, 65028, 65029, 65030, 65031, 65032, 65033,
		65034, 65035, 65036, 65037, 65038, 65039, 65040, 65041, 65042, 65043,
		65044, 65045, 65046, 65047, 65048, 65049, 65050, 65051, 65052, 65053,
		65054, 65055, 65056, 65057, 65058, 65059, 65060, 65061, 65062, 65063,
		65064, 65065, 65066, 65067, 65068, 65069, 65070, 65071, 65072, 65073,
		65074, 65075, 65076, 65077, 65078, 65079, 65080, 65081, 65082, 65083,
		65084, 65085, 65086, 65087, 65088, 65089, 65090, 65091, 65092, 65093,
		65094, 65095, 65096, 65097, 65098, 65099, 65100, 65101, 65102, 65103,
		65104, 65105, 65106, 65107, 65108, 65109, 65110, 65111, 65112, 65113,
		65114, 65115, 65116, 65117, 65118, 65119, 65120, 65121, 65122, 65123,
		65124, 65125, 65126, 65127, 65128, 65129, 65130, 65131, 65132, 65133,
		65134, 65135, 65136, 65137, 65138, 65139, 65140, 65141, 65142, 65143,
		65144, 65145, 65146, 65147, 65148, 65149, 65150, 65151, 65152, 65153,
		65154, 65155, 65156, 65157, 65158, 65159, 65160, 65161, 65162, 65163,
		65164, 65165, 65166, 65167, 65168, 65169, 65170, 65171, 65172, 65173,
		65174, 65175, 65176, 65177, 65178, 65179, 65180, 65181, 65182, 65183,
		65184, 65185, 65186, 65187, 65188, 65189, 65190, 65191, 65192, 65193,
		65194, 65195, 65196, 65197, 65198, 65199, 65200, 65201, 65202, 65203,
		65204, 65205, 65206, 65207, 65208, 65209, 65210, 65211, 65212, 65213,
		65214, 65215, 65216, 65217, 65218, 65219, 65220, 65221, 65222, 65223,
		65224, 65225, 65226, 65227, 65228, 65229, 65230, 65231, 65232, 65233,
		65234, 65235, 65236, 65237, 65238, 65239, 65240, 65241, 65242, 65243,
		65244, 65245, 65246, 65247, 65248, 65249, 65250, 65251, 65252, 65253,
		65254, 65255, 65256, 65257, 65258, 65259, 65260, 65261, 65262, 65263,
		65264, 65265, 65266, 65267, 65268, 65269, 65270, 65271, 65272, 65273,
		65274, 65275, 65276, 65277, 65278, 65279, 65280, 65281, 65282, 65283,
		65284, 65285, 65286, 65287, 65288, 65289, 65290, 65291, 65292, 65293,
		65294, 65295, 65296, 65297, 65298, 65299, 65300, 65301, 65302, 65303,
		65304, 65305, 65306, 65307, 65308, 65309, 65310, 65311, 65312, 65313,
		65314, 65315, 65316, 65317, 65318, 65319, 65320, 65321, 65322, 65323,
		65324, 65325, 65326, 65327, 65328, 65329, 65330, 65331, 65332, 65333,
		65334, 65335, 65336, 65337, 65338, 65339, 65340, 65341, 65342, 65343,
		65344, 65345, 65346, 65347, 65348, 65349, 65350, 65351, 65352, 65353,
		65354, 65355, 65356, 65357, 65358, 65359, 65360, 65361, 65362, 65363,
		65364, 65365, 65366, 65367, 65368, 65369, 65370, 65371, 65372, 65373,
		65374, 65375, 65376, 65377, 65378, 65379, 65380, 65381, 65382, 65383,
		65384, 65385, 65386, 65387, 65388, 65389, 65390, 65391, 65392, 65393,
		65394, 65395, 65396, 65397, 65398, 65399, 65400, 65401, 65402, 65403,
		65404, 65405, 65406, 65407, 65408, 65409, 65410, 65411, 65412, 65413,
		65414, 65415, 65416, 65417, 65418, 65419, 65420, 65421, 65422, 65423,
		65424, 65425, 65426, 65427, 65428, 65429, 65430, 65431, 65432, 65433,
		65434, 65435, 65436, 65437, 65438, 65439, 65440, 65441, 65442, 65443,
		65444, 65445, 65446, 65447, 65448, 65449, 65450, 65451, 65452, 65453,
		65454, 65455, 65456, 65457, 65458, 65459, 65460, 65461, 65462, 65463,
		65464, 65465, 65466, 65467, 65468, 65469, 65470, 65471, 65472, 65473,
		65474, 65475, 65476, 65477, 65478, 65479, 65480, 65481, 65482, 65483,
		65484, 65485, 65486, 65487, 65488, 65489, 65490, 65491, 65492, 65493,
		65494, 65495, 65496, 65497, 65498, 65499, 65500, 65501, 65502, 65503,
		65504, 65505, 65506, 65507, 65508, 65509, 65510, 65511, 65512, 65513,
		65514, 65515, 65516, 65517, 65518, 65519, 65520, 65521, 65522, 65523,
		65524, 65525, 65526, 65527, 65528, 65529, 65530, 65531, 65532, 65533,
		65534, 65535,
	}

	return span
}
