// Generated from ONCRPCv2.g4 by ANTLR 4.5.3.

package parser

import (
	"fmt"
	"unicode"

	"github.com/pboyer/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 42, 306, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9,
	23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28,
	4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4,
	34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39,
	9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 248, 10,
	35, 12, 35, 14, 35, 251, 11, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 7, 36, 261, 10, 36, 12, 36, 14, 36, 264, 11, 36, 3, 37,
	5, 37, 267, 10, 37, 3, 37, 6, 37, 270, 10, 37, 13, 37, 14, 37, 271, 3,
	38, 3, 38, 3, 38, 3, 38, 6, 38, 278, 10, 38, 13, 38, 14, 38, 279, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 291, 10,
	39, 3, 40, 3, 40, 7, 40, 295, 10, 40, 12, 40, 14, 40, 298, 11, 40, 3, 41,
	6, 41, 301, 10, 41, 13, 41, 14, 41, 302, 3, 41, 3, 41, 3, 249, 2, 42, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40,
	79, 41, 81, 42, 3, 2, 9, 3, 2, 51, 57, 3, 2, 50, 57, 3, 2, 50, 59, 5, 2,
	50, 59, 67, 72, 99, 104, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 313, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2,
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2,
	2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2,
	2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3,
	2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81,
	3, 2, 2, 2, 3, 83, 3, 2, 2, 2, 5, 91, 3, 2, 2, 2, 7, 93, 3, 2, 2, 2, 9,
	95, 3, 2, 2, 2, 11, 97, 3, 2, 2, 2, 13, 99, 3, 2, 2, 2, 15, 107, 3, 2,
	2, 2, 17, 109, 3, 2, 2, 2, 19, 111, 3, 2, 2, 2, 21, 113, 3, 2, 2, 2, 23,
	118, 3, 2, 2, 2, 25, 120, 3, 2, 2, 2, 27, 122, 3, 2, 2, 2, 29, 124, 3,
	2, 2, 2, 31, 126, 3, 2, 2, 2, 33, 133, 3, 2, 2, 2, 35, 140, 3, 2, 2, 2,
	37, 142, 3, 2, 2, 2, 39, 151, 3, 2, 2, 2, 41, 155, 3, 2, 2, 2, 43, 161,
	3, 2, 2, 2, 45, 167, 3, 2, 2, 2, 47, 174, 3, 2, 2, 2, 49, 184, 3, 2, 2,
	2, 51, 189, 3, 2, 2, 2, 53, 196, 3, 2, 2, 2, 55, 201, 3, 2, 2, 2, 57, 207,
	3, 2, 2, 2, 59, 214, 3, 2, 2, 2, 61, 219, 3, 2, 2, 2, 63, 221, 3, 2, 2,
	2, 65, 229, 3, 2, 2, 2, 67, 235, 3, 2, 2, 2, 69, 243, 3, 2, 2, 2, 71, 257,
	3, 2, 2, 2, 73, 266, 3, 2, 2, 2, 75, 273, 3, 2, 2, 2, 77, 290, 3, 2, 2,
	2, 79, 292, 3, 2, 2, 2, 81, 300, 3, 2, 2, 2, 83, 84, 7, 114, 2, 2, 84,
	85, 7, 116, 2, 2, 85, 86, 7, 113, 2, 2, 86, 87, 7, 105, 2, 2, 87, 88, 7,
	116, 2, 2, 88, 89, 7, 99, 2, 2, 89, 90, 7, 111, 2, 2, 90, 4, 3, 2, 2, 2,
	91, 92, 7, 125, 2, 2, 92, 6, 3, 2, 2, 2, 93, 94, 7, 127, 2, 2, 94, 8, 3,
	2, 2, 2, 95, 96, 7, 63, 2, 2, 96, 10, 3, 2, 2, 2, 97, 98, 7, 61, 2, 2,
	98, 12, 3, 2, 2, 2, 99, 100, 7, 120, 2, 2, 100, 101, 7, 103, 2, 2, 101,
	102, 7, 116, 2, 2, 102, 103, 7, 117, 2, 2, 103, 104, 7, 107, 2, 2, 104,
	105, 7, 113, 2, 2, 105, 106, 7, 112, 2, 2, 106, 14, 3, 2, 2, 2, 107, 108,
	7, 42, 2, 2, 108, 16, 3, 2, 2, 2, 109, 110, 7, 46, 2, 2, 110, 18, 3, 2,
	2, 2, 111, 112, 7, 43, 2, 2, 112, 20, 3, 2, 2, 2, 113, 114, 7, 120, 2,
	2, 114, 115, 7, 113, 2, 2, 115, 116, 7, 107, 2, 2, 116, 117, 7, 102, 2,
	2, 117, 22, 3, 2, 2, 2, 118, 119, 7, 93, 2, 2, 119, 24, 3, 2, 2, 2, 120,
	121, 7, 95, 2, 2, 121, 26, 3, 2, 2, 2, 122, 123, 7, 62, 2, 2, 123, 28,
	3, 2, 2, 2, 124, 125, 7, 64, 2, 2, 125, 30, 3, 2, 2, 2, 126, 127, 7, 113,
	2, 2, 127, 128, 7, 114, 2, 2, 128, 129, 7, 99, 2, 2, 129, 130, 7, 115,
	2, 2, 130, 131, 7, 119, 2, 2, 131, 132, 7, 103, 2, 2, 132, 32, 3, 2, 2,
	2, 133, 134, 7, 117, 2, 2, 134, 135, 7, 118, 2, 2, 135, 136, 7, 116, 2,
	2, 136, 137, 7, 107, 2, 2, 137, 138, 7, 112, 2, 2, 138, 139, 7, 105, 2,
	2, 139, 34, 3, 2, 2, 2, 140, 141, 7, 44, 2, 2, 141, 36, 3, 2, 2, 2, 142,
	143, 7, 119, 2, 2, 143, 144, 7, 112, 2, 2, 144, 145, 7, 117, 2, 2, 145,
	146, 7, 107, 2, 2, 146, 147, 7, 105, 2, 2, 147, 148, 7, 112, 2, 2, 148,
	149, 7, 103, 2, 2, 149, 150, 7, 102, 2, 2, 150, 38, 3, 2, 2, 2, 151, 152,
	7, 107, 2, 2, 152, 153, 7, 112, 2, 2, 153, 154, 7, 118, 2, 2, 154, 40,
	3, 2, 2, 2, 155, 156, 7, 106, 2, 2, 156, 157, 7, 123, 2, 2, 157, 158, 7,
	114, 2, 2, 158, 159, 7, 103, 2, 2, 159, 160, 7, 116, 2, 2, 160, 42, 3,
	2, 2, 2, 161, 162, 7, 104, 2, 2, 162, 163, 7, 110, 2, 2, 163, 164, 7, 113,
	2, 2, 164, 165, 7, 99, 2, 2, 165, 166, 7, 118, 2, 2, 166, 44, 3, 2, 2,
	2, 167, 168, 7, 102, 2, 2, 168, 169, 7, 113, 2, 2, 169, 170, 7, 119, 2,
	2, 170, 171, 7, 100, 2, 2, 171, 172, 7, 110, 2, 2, 172, 173, 7, 103, 2,
	2, 173, 46, 3, 2, 2, 2, 174, 175, 7, 115, 2, 2, 175, 176, 7, 119, 2, 2,
	176, 177, 7, 99, 2, 2, 177, 178, 7, 102, 2, 2, 178, 179, 7, 116, 2, 2,
	179, 180, 7, 119, 2, 2, 180, 181, 7, 114, 2, 2, 181, 182, 7, 110, 2, 2,
	182, 183, 7, 103, 2, 2, 183, 48, 3, 2, 2, 2, 184, 185, 7, 100, 2, 2, 185,
	186, 7, 113, 2, 2, 186, 187, 7, 113, 2, 2, 187, 188, 7, 110, 2, 2, 188,
	50, 3, 2, 2, 2, 189, 190, 7, 117, 2, 2, 190, 191, 7, 118, 2, 2, 191, 192,
	7, 116, 2, 2, 192, 193, 7, 119, 2, 2, 193, 194, 7, 101, 2, 2, 194, 195,
	7, 118, 2, 2, 195, 52, 3, 2, 2, 2, 196, 197, 7, 103, 2, 2, 197, 198, 7,
	112, 2, 2, 198, 199, 7, 119, 2, 2, 199, 200, 7, 111, 2, 2, 200, 54, 3,
	2, 2, 2, 201, 202, 7, 119, 2, 2, 202, 203, 7, 112, 2, 2, 203, 204, 7, 107,
	2, 2, 204, 205, 7, 113, 2, 2, 205, 206, 7, 112, 2, 2, 206, 56, 3, 2, 2,
	2, 207, 208, 7, 117, 2, 2, 208, 209, 7, 121, 2, 2, 209, 210, 7, 107, 2,
	2, 210, 211, 7, 118, 2, 2, 211, 212, 7, 101, 2, 2, 212, 213, 7, 106, 2,
	2, 213, 58, 3, 2, 2, 2, 214, 215, 7, 101, 2, 2, 215, 216, 7, 99, 2, 2,
	216, 217, 7, 117, 2, 2, 217, 218, 7, 103, 2, 2, 218, 60, 3, 2, 2, 2, 219,
	220, 7, 60, 2, 2, 220, 62, 3, 2, 2, 2, 221, 222, 7, 102, 2, 2, 222, 223,
	7, 103, 2, 2, 223, 224, 7, 104, 2, 2, 224, 225, 7, 99, 2, 2, 225, 226,
	7, 119, 2, 2, 226, 227, 7, 110, 2, 2, 227, 228, 7, 118, 2, 2, 228, 64,
	3, 2, 2, 2, 229, 230, 7, 101, 2, 2, 230, 231, 7, 113, 2, 2, 231, 232, 7,
	112, 2, 2, 232, 233, 7, 117, 2, 2, 233, 234, 7, 118, 2, 2, 234, 66, 3,
	2, 2, 2, 235, 236, 7, 118, 2, 2, 236, 237, 7, 123, 2, 2, 237, 238, 7, 114,
	2, 2, 238, 239, 7, 103, 2, 2, 239, 240, 7, 102, 2, 2, 240, 241, 7, 103,
	2, 2, 241, 242, 7, 104, 2, 2, 242, 68, 3, 2, 2, 2, 243, 244, 7, 49, 2,
	2, 244, 245, 7, 44, 2, 2, 245, 249, 3, 2, 2, 2, 246, 248, 11, 2, 2, 2,
	247, 246, 3, 2, 2, 2, 248, 251, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 249,
	247, 3, 2, 2, 2, 250, 252, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 252, 253,
	7, 44, 2, 2, 253, 254, 7, 49, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 8,
	35, 2, 2, 256, 70, 3, 2, 2, 2, 257, 258, 7, 50, 2, 2, 258, 262, 9, 2, 2,
	2, 259, 261, 9, 3, 2, 2, 260, 259, 3, 2, 2, 2, 261, 264, 3, 2, 2, 2, 262,
	260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 72, 3, 2, 2, 2, 264, 262, 3,
	2, 2, 2, 265, 267, 7, 47, 2, 2, 266, 265, 3, 2, 2, 2, 266, 267, 3, 2, 2,
	2, 267, 269, 3, 2, 2, 2, 268, 270, 9, 4, 2, 2, 269, 268, 3, 2, 2, 2, 270,
	271, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 74, 3,
	2, 2, 2, 273, 274, 7, 50, 2, 2, 274, 275, 7, 122, 2, 2, 275, 277, 3, 2,
	2, 2, 276, 278, 9, 5, 2, 2, 277, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2,
	279, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 76, 3, 2, 2, 2, 281, 282,
	7, 86, 2, 2, 282, 283, 7, 84, 2, 2, 283, 284, 7, 87, 2, 2, 284, 291, 7,
	71, 2, 2, 285, 286, 7, 72, 2, 2, 286, 287, 7, 67, 2, 2, 287, 288, 7, 78,
	2, 2, 288, 289, 7, 85, 2, 2, 289, 291, 7, 71, 2, 2, 290, 281, 3, 2, 2,
	2, 290, 285, 3, 2, 2, 2, 291, 78, 3, 2, 2, 2, 292, 296, 9, 6, 2, 2, 293,
	295, 9, 7, 2, 2, 294, 293, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294,
	3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 80, 3, 2, 2, 2, 298, 296, 3, 2,
	2, 2, 299, 301, 9, 8, 2, 2, 300, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2,
	302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304,
	305, 8, 41, 2, 2, 305, 82, 3, 2, 2, 2, 11, 2, 249, 262, 266, 271, 279,
	290, 296, 302, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'program'", "'{'", "'}'", "'='", "';'", "'version'", "'('", "','",
	"')'", "'void'", "'['", "']'", "'<'", "'>'", "'opaque'", "'string'", "'*'",
	"'unsigned'", "'int'", "'hyper'", "'float'", "'double'", "'quadruple'",
	"'bool'", "'struct'", "'enum'", "'union'", "'switch'", "'case'", "':'",
	"'default'", "'const'", "'typedef'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMENT",
	"OCTAL", "DECIMAL", "HEXADECIMAL", "BOOLEAN", "IDENTIFIER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"COMMENT", "OCTAL", "DECIMAL", "HEXADECIMAL", "BOOLEAN", "IDENTIFIER",
	"WS",
}

type ONCRPCv2Lexer struct {
	*antlr.BaseLexer
	modeNames []string
	// TODO: EOF string
}

func NewONCRPCv2Lexer(input antlr.CharStream) *ONCRPCv2Lexer {
	var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	l := new(ONCRPCv2Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ONCRPCv2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ONCRPCv2Lexer tokens.
const (
	ONCRPCv2LexerT__0        = 1
	ONCRPCv2LexerT__1        = 2
	ONCRPCv2LexerT__2        = 3
	ONCRPCv2LexerT__3        = 4
	ONCRPCv2LexerT__4        = 5
	ONCRPCv2LexerT__5        = 6
	ONCRPCv2LexerT__6        = 7
	ONCRPCv2LexerT__7        = 8
	ONCRPCv2LexerT__8        = 9
	ONCRPCv2LexerT__9        = 10
	ONCRPCv2LexerT__10       = 11
	ONCRPCv2LexerT__11       = 12
	ONCRPCv2LexerT__12       = 13
	ONCRPCv2LexerT__13       = 14
	ONCRPCv2LexerT__14       = 15
	ONCRPCv2LexerT__15       = 16
	ONCRPCv2LexerT__16       = 17
	ONCRPCv2LexerT__17       = 18
	ONCRPCv2LexerT__18       = 19
	ONCRPCv2LexerT__19       = 20
	ONCRPCv2LexerT__20       = 21
	ONCRPCv2LexerT__21       = 22
	ONCRPCv2LexerT__22       = 23
	ONCRPCv2LexerT__23       = 24
	ONCRPCv2LexerT__24       = 25
	ONCRPCv2LexerT__25       = 26
	ONCRPCv2LexerT__26       = 27
	ONCRPCv2LexerT__27       = 28
	ONCRPCv2LexerT__28       = 29
	ONCRPCv2LexerT__29       = 30
	ONCRPCv2LexerT__30       = 31
	ONCRPCv2LexerT__31       = 32
	ONCRPCv2LexerT__32       = 33
	ONCRPCv2LexerCOMMENT     = 34
	ONCRPCv2LexerOCTAL       = 35
	ONCRPCv2LexerDECIMAL     = 36
	ONCRPCv2LexerHEXADECIMAL = 37
	ONCRPCv2LexerBOOLEAN     = 38
	ONCRPCv2LexerIDENTIFIER  = 39
	ONCRPCv2LexerWS          = 40
)