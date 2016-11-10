// Generated from ONCRPCv2.g4 by ANTLR 4.5.3.

package parser // ONCRPCv2

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pboyer/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 42, 294, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18,
	4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 67, 10,
	3, 12, 3, 14, 3, 70, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 7, 4, 83, 10, 4, 12, 4, 14, 4, 86, 11, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 95, 10, 5, 3, 6, 3, 6, 5, 6, 99,
	10, 6, 3, 7, 3, 7, 7, 7, 103, 10, 7, 12, 7, 14, 7, 106, 11, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 121, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 135, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 142,
	10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 150, 10, 8, 3, 9, 3, 9,
	5, 9, 154, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 160, 10, 11, 3, 11,
	5, 11, 163, 10, 11, 3, 11, 5, 11, 166, 10, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 179, 10, 11,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 7, 13, 193, 10, 13, 12, 13, 14, 13, 196, 11, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7,
	15, 210, 10, 15, 12, 15, 14, 15, 213, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 227, 10,
	17, 12, 17, 14, 17, 230, 11, 17, 3, 17, 5, 17, 233, 10, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 245, 10,
	18, 12, 18, 14, 18, 248, 11, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 283, 10, 21, 3, 22,
	3, 22, 5, 22, 287, 10, 22, 3, 23, 6, 23, 290, 10, 23, 13, 23, 14, 23, 291,
	3, 23, 2, 2, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 2, 3, 3, 2, 37, 40, 312, 2, 46, 3, 2, 2, 2,
	4, 61, 3, 2, 2, 2, 6, 76, 3, 2, 2, 2, 8, 94, 3, 2, 2, 2, 10, 98, 3, 2,
	2, 2, 12, 104, 3, 2, 2, 2, 14, 149, 3, 2, 2, 2, 16, 153, 3, 2, 2, 2, 18,
	155, 3, 2, 2, 2, 20, 178, 3, 2, 2, 2, 22, 180, 3, 2, 2, 2, 24, 183, 3,
	2, 2, 2, 26, 199, 3, 2, 2, 2, 28, 202, 3, 2, 2, 2, 30, 216, 3, 2, 2, 2,
	32, 219, 3, 2, 2, 2, 34, 236, 3, 2, 2, 2, 36, 252, 3, 2, 2, 2, 38, 257,
	3, 2, 2, 2, 40, 282, 3, 2, 2, 2, 42, 286, 3, 2, 2, 2, 44, 289, 3, 2, 2,
	2, 46, 47, 7, 3, 2, 2, 47, 48, 7, 41, 2, 2, 48, 49, 7, 4, 2, 2, 49, 53,
	5, 4, 3, 2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2,
	53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3,
	2, 2, 2, 56, 57, 7, 5, 2, 2, 57, 58, 7, 6, 2, 2, 58, 59, 5, 18, 10, 2,
	59, 60, 7, 7, 2, 2, 60, 3, 3, 2, 2, 2, 61, 62, 7, 8, 2, 2, 62, 63, 7, 41,
	2, 2, 63, 64, 7, 4, 2, 2, 64, 68, 5, 6, 4, 2, 65, 67, 5, 6, 4, 2, 66, 65,
	3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 5, 2, 2, 72, 73, 7,
	6, 2, 2, 73, 74, 5, 18, 10, 2, 74, 75, 7, 7, 2, 2, 75, 5, 3, 2, 2, 2, 76,
	77, 5, 8, 5, 2, 77, 78, 7, 41, 2, 2, 78, 79, 7, 9, 2, 2, 79, 84, 5, 10,
	6, 2, 80, 81, 7, 10, 2, 2, 81, 83, 5, 20, 11, 2, 82, 80, 3, 2, 2, 2, 83,
	86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 87, 3, 2, 2,
	2, 86, 84, 3, 2, 2, 2, 87, 88, 7, 11, 2, 2, 88, 89, 7, 6, 2, 2, 89, 90,
	5, 18, 10, 2, 90, 91, 7, 7, 2, 2, 91, 7, 3, 2, 2, 2, 92, 95, 7, 12, 2,
	2, 93, 95, 5, 20, 11, 2, 94, 92, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 9,
	3, 2, 2, 2, 96, 99, 7, 12, 2, 2, 97, 99, 5, 20, 11, 2, 98, 96, 3, 2, 2,
	2, 98, 97, 3, 2, 2, 2, 99, 11, 3, 2, 2, 2, 100, 103, 5, 44, 23, 2, 101,
	103, 5, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106,
	3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 13, 3, 2,
	2, 2, 106, 104, 3, 2, 2, 2, 107, 108, 5, 20, 11, 2, 108, 109, 7, 41, 2,
	2, 109, 150, 3, 2, 2, 2, 110, 111, 5, 20, 11, 2, 111, 112, 7, 41, 2, 2,
	112, 113, 7, 13, 2, 2, 113, 114, 5, 16, 9, 2, 114, 115, 7, 14, 2, 2, 115,
	150, 3, 2, 2, 2, 116, 117, 5, 20, 11, 2, 117, 118, 7, 41, 2, 2, 118, 120,
	7, 15, 2, 2, 119, 121, 5, 16, 9, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3,
	2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 7, 16, 2, 2, 123, 150, 3, 2, 2,
	2, 124, 125, 7, 17, 2, 2, 125, 126, 7, 41, 2, 2, 126, 127, 7, 13, 2, 2,
	127, 128, 5, 16, 9, 2, 128, 129, 7, 14, 2, 2, 129, 150, 3, 2, 2, 2, 130,
	131, 7, 17, 2, 2, 131, 132, 7, 41, 2, 2, 132, 134, 7, 15, 2, 2, 133, 135,
	5, 16, 9, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2,
	2, 2, 136, 150, 7, 16, 2, 2, 137, 138, 7, 18, 2, 2, 138, 139, 7, 41, 2,
	2, 139, 141, 7, 15, 2, 2, 140, 142, 5, 16, 9, 2, 141, 140, 3, 2, 2, 2,
	141, 142, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 150, 7, 16, 2, 2, 144,
	145, 5, 20, 11, 2, 145, 146, 7, 19, 2, 2, 146, 147, 7, 41, 2, 2, 147, 150,
	3, 2, 2, 2, 148, 150, 7, 12, 2, 2, 149, 107, 3, 2, 2, 2, 149, 110, 3, 2,
	2, 2, 149, 116, 3, 2, 2, 2, 149, 124, 3, 2, 2, 2, 149, 130, 3, 2, 2, 2,
	149, 137, 3, 2, 2, 2, 149, 144, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150,
	15, 3, 2, 2, 2, 151, 154, 5, 18, 10, 2, 152, 154, 7, 41, 2, 2, 153, 151,
	3, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154, 17, 3, 2, 2, 2, 155, 156, 9, 2,
	2, 2, 156, 19, 3, 2, 2, 2, 157, 163, 7, 20, 2, 2, 158, 160, 7, 20, 2, 2,
	159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161,
	163, 7, 21, 2, 2, 162, 157, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 163, 179,
	3, 2, 2, 2, 164, 166, 7, 20, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 167, 3, 2, 2, 2, 167, 179, 7, 22, 2, 2, 168, 179, 7, 23, 2,
	2, 169, 179, 7, 24, 2, 2, 170, 179, 7, 25, 2, 2, 171, 179, 7, 26, 2, 2,
	172, 173, 7, 27, 2, 2, 173, 179, 7, 41, 2, 2, 174, 179, 5, 22, 12, 2, 175,
	179, 5, 26, 14, 2, 176, 179, 5, 30, 16, 2, 177, 179, 7, 41, 2, 2, 178,
	162, 3, 2, 2, 2, 178, 165, 3, 2, 2, 2, 178, 168, 3, 2, 2, 2, 178, 169,
	3, 2, 2, 2, 178, 170, 3, 2, 2, 2, 178, 171, 3, 2, 2, 2, 178, 172, 3, 2,
	2, 2, 178, 174, 3, 2, 2, 2, 178, 175, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2,
	178, 177, 3, 2, 2, 2, 179, 21, 3, 2, 2, 2, 180, 181, 7, 28, 2, 2, 181,
	182, 5, 24, 13, 2, 182, 23, 3, 2, 2, 2, 183, 184, 7, 4, 2, 2, 184, 185,
	7, 41, 2, 2, 185, 186, 7, 6, 2, 2, 186, 187, 5, 16, 9, 2, 187, 194, 3,
	2, 2, 2, 188, 189, 7, 10, 2, 2, 189, 190, 7, 41, 2, 2, 190, 191, 7, 6,
	2, 2, 191, 193, 5, 16, 9, 2, 192, 188, 3, 2, 2, 2, 193, 196, 3, 2, 2, 2,
	194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2, 196,
	194, 3, 2, 2, 2, 197, 198, 7, 5, 2, 2, 198, 25, 3, 2, 2, 2, 199, 200, 7,
	27, 2, 2, 200, 201, 5, 28, 15, 2, 201, 27, 3, 2, 2, 2, 202, 203, 7, 4,
	2, 2, 203, 204, 5, 14, 8, 2, 204, 205, 7, 7, 2, 2, 205, 211, 3, 2, 2, 2,
	206, 207, 5, 14, 8, 2, 207, 208, 7, 7, 2, 2, 208, 210, 3, 2, 2, 2, 209,
	206, 3, 2, 2, 2, 210, 213, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212,
	3, 2, 2, 2, 212, 214, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 214, 215, 7, 5,
	2, 2, 215, 29, 3, 2, 2, 2, 216, 217, 7, 29, 2, 2, 217, 218, 5, 32, 17,
	2, 218, 31, 3, 2, 2, 2, 219, 220, 7, 30, 2, 2, 220, 221, 7, 9, 2, 2, 221,
	222, 5, 14, 8, 2, 222, 223, 7, 11, 2, 2, 223, 224, 7, 4, 2, 2, 224, 228,
	5, 34, 18, 2, 225, 227, 5, 34, 18, 2, 226, 225, 3, 2, 2, 2, 227, 230, 3,
	2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 232, 3, 2, 2,
	2, 230, 228, 3, 2, 2, 2, 231, 233, 5, 36, 19, 2, 232, 231, 3, 2, 2, 2,
	232, 233, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 7, 5, 2, 2, 235,
	33, 3, 2, 2, 2, 236, 237, 7, 31, 2, 2, 237, 238, 5, 16, 9, 2, 238, 239,
	7, 32, 2, 2, 239, 246, 3, 2, 2, 2, 240, 241, 7, 31, 2, 2, 241, 242, 5,
	16, 9, 2, 242, 243, 7, 32, 2, 2, 243, 245, 3, 2, 2, 2, 244, 240, 3, 2,
	2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2,
	247, 249, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 250, 5, 14, 8, 2, 250,
	251, 7, 7, 2, 2, 251, 35, 3, 2, 2, 2, 252, 253, 7, 33, 2, 2, 253, 254,
	7, 32, 2, 2, 254, 255, 5, 14, 8, 2, 255, 256, 7, 7, 2, 2, 256, 37, 3, 2,
	2, 2, 257, 258, 7, 34, 2, 2, 258, 259, 7, 41, 2, 2, 259, 260, 7, 6, 2,
	2, 260, 261, 5, 18, 10, 2, 261, 262, 7, 7, 2, 2, 262, 39, 3, 2, 2, 2, 263,
	264, 7, 35, 2, 2, 264, 265, 5, 14, 8, 2, 265, 266, 7, 7, 2, 2, 266, 283,
	3, 2, 2, 2, 267, 268, 7, 28, 2, 2, 268, 269, 7, 41, 2, 2, 269, 270, 5,
	24, 13, 2, 270, 271, 7, 7, 2, 2, 271, 283, 3, 2, 2, 2, 272, 273, 7, 27,
	2, 2, 273, 274, 7, 41, 2, 2, 274, 275, 5, 28, 15, 2, 275, 276, 7, 7, 2,
	2, 276, 283, 3, 2, 2, 2, 277, 278, 7, 29, 2, 2, 278, 279, 7, 41, 2, 2,
	279, 280, 5, 32, 17, 2, 280, 281, 7, 7, 2, 2, 281, 283, 3, 2, 2, 2, 282,
	263, 3, 2, 2, 2, 282, 267, 3, 2, 2, 2, 282, 272, 3, 2, 2, 2, 282, 277,
	3, 2, 2, 2, 283, 41, 3, 2, 2, 2, 284, 287, 5, 40, 21, 2, 285, 287, 5, 38,
	20, 2, 286, 284, 3, 2, 2, 2, 286, 285, 3, 2, 2, 2, 287, 43, 3, 2, 2, 2,
	288, 290, 5, 42, 22, 2, 289, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291,
	289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 45, 3, 2, 2, 2, 26, 53, 68,
	84, 94, 98, 102, 104, 120, 134, 141, 149, 153, 159, 162, 165, 178, 194,
	211, 228, 232, 246, 282, 286, 291,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'program'", "'{'", "'}'", "'='", "';'", "'version'", "'('", "','",
	"')'", "'void'", "'['", "']'", "'<'", "'>'", "'opaque'", "'string'", "'*'",
	"'unsigned'", "'int'", "'hyper'", "'float'", "'double'", "'quadruple'",
	"'bool'", "'struct'", "'enum'", "'union'", "'switch'", "'case'", "':'",
	"'default'", "'const'", "'typedef'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMENT",
	"OCTAL", "DECIMAL", "HEXADECIMAL", "BOOLEAN", "IDENTIFIER", "WS",
}

var ruleNames = []string{
	"programDef", "versionDef", "procedureDef", "procReturn", "procFirstArg",
	"oncrpcv2Specification", "declaration", "value", "constant", "typeSpecifier",
	"enumTypeSpec", "enumBody", "structTypeSpec", "structBody", "unionTypeSpec",
	"unionBody", "caseSpec", "defaultCaseSpec", "constantDef", "typeDef", "definition",
	"xdrSpecification",
}

type ONCRPCv2Parser struct {
	*antlr.BaseParser
}

func NewONCRPCv2Parser(input antlr.TokenStream) *ONCRPCv2Parser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(ONCRPCv2Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ONCRPCv2.g4"

	return this
}

// ONCRPCv2Parser tokens.
const (
	ONCRPCv2ParserEOF         = antlr.TokenEOF
	ONCRPCv2ParserT__0        = 1
	ONCRPCv2ParserT__1        = 2
	ONCRPCv2ParserT__2        = 3
	ONCRPCv2ParserT__3        = 4
	ONCRPCv2ParserT__4        = 5
	ONCRPCv2ParserT__5        = 6
	ONCRPCv2ParserT__6        = 7
	ONCRPCv2ParserT__7        = 8
	ONCRPCv2ParserT__8        = 9
	ONCRPCv2ParserT__9        = 10
	ONCRPCv2ParserT__10       = 11
	ONCRPCv2ParserT__11       = 12
	ONCRPCv2ParserT__12       = 13
	ONCRPCv2ParserT__13       = 14
	ONCRPCv2ParserT__14       = 15
	ONCRPCv2ParserT__15       = 16
	ONCRPCv2ParserT__16       = 17
	ONCRPCv2ParserT__17       = 18
	ONCRPCv2ParserT__18       = 19
	ONCRPCv2ParserT__19       = 20
	ONCRPCv2ParserT__20       = 21
	ONCRPCv2ParserT__21       = 22
	ONCRPCv2ParserT__22       = 23
	ONCRPCv2ParserT__23       = 24
	ONCRPCv2ParserT__24       = 25
	ONCRPCv2ParserT__25       = 26
	ONCRPCv2ParserT__26       = 27
	ONCRPCv2ParserT__27       = 28
	ONCRPCv2ParserT__28       = 29
	ONCRPCv2ParserT__29       = 30
	ONCRPCv2ParserT__30       = 31
	ONCRPCv2ParserT__31       = 32
	ONCRPCv2ParserT__32       = 33
	ONCRPCv2ParserCOMMENT     = 34
	ONCRPCv2ParserOCTAL       = 35
	ONCRPCv2ParserDECIMAL     = 36
	ONCRPCv2ParserHEXADECIMAL = 37
	ONCRPCv2ParserBOOLEAN     = 38
	ONCRPCv2ParserIDENTIFIER  = 39
	ONCRPCv2ParserWS          = 40
)

// ONCRPCv2Parser rules.
const (
	ONCRPCv2ParserRULE_programDef            = 0
	ONCRPCv2ParserRULE_versionDef            = 1
	ONCRPCv2ParserRULE_procedureDef          = 2
	ONCRPCv2ParserRULE_procReturn            = 3
	ONCRPCv2ParserRULE_procFirstArg          = 4
	ONCRPCv2ParserRULE_oncrpcv2Specification = 5
	ONCRPCv2ParserRULE_declaration           = 6
	ONCRPCv2ParserRULE_value                 = 7
	ONCRPCv2ParserRULE_constant              = 8
	ONCRPCv2ParserRULE_typeSpecifier         = 9
	ONCRPCv2ParserRULE_enumTypeSpec          = 10
	ONCRPCv2ParserRULE_enumBody              = 11
	ONCRPCv2ParserRULE_structTypeSpec        = 12
	ONCRPCv2ParserRULE_structBody            = 13
	ONCRPCv2ParserRULE_unionTypeSpec         = 14
	ONCRPCv2ParserRULE_unionBody             = 15
	ONCRPCv2ParserRULE_caseSpec              = 16
	ONCRPCv2ParserRULE_defaultCaseSpec       = 17
	ONCRPCv2ParserRULE_constantDef           = 18
	ONCRPCv2ParserRULE_typeDef               = 19
	ONCRPCv2ParserRULE_definition            = 20
	ONCRPCv2ParserRULE_xdrSpecification      = 21
)

// IProgramDefContext is an interface to support dynamic dispatch.
type IProgramDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramDefContext differentiates from other interfaces.
	IsProgramDefContext()
}

type ProgramDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramDefContext() *ProgramDefContext {
	var p = new(ProgramDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_programDef
	return p
}

func (*ProgramDefContext) IsProgramDefContext() {}

func NewProgramDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramDefContext {
	var p = new(ProgramDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_programDef

	return p
}

func (s *ProgramDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *ProgramDefContext) AllVersionDef() []IVersionDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVersionDefContext)(nil)).Elem())
	var tst = make([]IVersionDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVersionDefContext)
		}
	}

	return tst
}

func (s *ProgramDefContext) VersionDef(i int) IVersionDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVersionDefContext)
}

func (s *ProgramDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProgramDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterProgramDef(s)
	}
}

func (s *ProgramDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitProgramDef(s)
	}
}

func (p *ONCRPCv2Parser) ProgramDef() (localctx IProgramDefContext) {
	localctx = NewProgramDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ONCRPCv2ParserRULE_programDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(44)
	p.Match(ONCRPCv2ParserT__0)

	p.SetState(45)
	p.Match(ONCRPCv2ParserIDENTIFIER)

	p.SetState(46)
	p.Match(ONCRPCv2ParserT__1)

	p.SetState(47)
	p.VersionDef()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ONCRPCv2ParserT__5 {
		p.SetState(48)
		p.VersionDef()

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(54)
	p.Match(ONCRPCv2ParserT__2)

	p.SetState(55)
	p.Match(ONCRPCv2ParserT__3)

	p.SetState(56)
	p.Constant()

	p.SetState(57)
	p.Match(ONCRPCv2ParserT__4)

	return localctx
}

// IVersionDefContext is an interface to support dynamic dispatch.
type IVersionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionDefContext differentiates from other interfaces.
	IsVersionDefContext()
}

type VersionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionDefContext() *VersionDefContext {
	var p = new(VersionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_versionDef
	return p
}

func (*VersionDefContext) IsVersionDefContext() {}

func NewVersionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionDefContext {
	var p = new(VersionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_versionDef

	return p
}

func (s *VersionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *VersionDefContext) AllProcedureDef() []IProcedureDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcedureDefContext)(nil)).Elem())
	var tst = make([]IProcedureDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcedureDefContext)
		}
	}

	return tst
}

func (s *VersionDefContext) ProcedureDef(i int) IProcedureDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedureDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcedureDefContext)
}

func (s *VersionDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *VersionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterVersionDef(s)
	}
}

func (s *VersionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitVersionDef(s)
	}
}

func (p *ONCRPCv2Parser) VersionDef() (localctx IVersionDefContext) {
	localctx = NewVersionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ONCRPCv2ParserRULE_versionDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.Match(ONCRPCv2ParserT__5)

	p.SetState(60)
	p.Match(ONCRPCv2ParserIDENTIFIER)

	p.SetState(61)
	p.Match(ONCRPCv2ParserT__1)

	p.SetState(62)
	p.ProcedureDef()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(ONCRPCv2ParserT__9-10))|(1<<(ONCRPCv2ParserT__17-10))|(1<<(ONCRPCv2ParserT__18-10))|(1<<(ONCRPCv2ParserT__19-10))|(1<<(ONCRPCv2ParserT__20-10))|(1<<(ONCRPCv2ParserT__21-10))|(1<<(ONCRPCv2ParserT__22-10))|(1<<(ONCRPCv2ParserT__23-10))|(1<<(ONCRPCv2ParserT__24-10))|(1<<(ONCRPCv2ParserT__25-10))|(1<<(ONCRPCv2ParserT__26-10))|(1<<(ONCRPCv2ParserIDENTIFIER-10)))) != 0 {
		p.SetState(63)
		p.ProcedureDef()

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(69)
	p.Match(ONCRPCv2ParserT__2)

	p.SetState(70)
	p.Match(ONCRPCv2ParserT__3)

	p.SetState(71)
	p.Constant()

	p.SetState(72)
	p.Match(ONCRPCv2ParserT__4)

	return localctx
}

// IProcedureDefContext is an interface to support dynamic dispatch.
type IProcedureDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcedureDefContext differentiates from other interfaces.
	IsProcedureDefContext()
}

type ProcedureDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureDefContext() *ProcedureDefContext {
	var p = new(ProcedureDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_procedureDef
	return p
}

func (*ProcedureDefContext) IsProcedureDefContext() {}

func NewProcedureDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureDefContext {
	var p = new(ProcedureDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_procedureDef

	return p
}

func (s *ProcedureDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureDefContext) ProcReturn() IProcReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcReturnContext)
}

func (s *ProcedureDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *ProcedureDefContext) ProcFirstArg() IProcFirstArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcFirstArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcFirstArgContext)
}

func (s *ProcedureDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProcedureDefContext) AllTypeSpecifier() []ITypeSpecifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem())
	var tst = make([]ITypeSpecifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecifierContext)
		}
	}

	return tst
}

func (s *ProcedureDefContext) TypeSpecifier(i int) ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcedureDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterProcedureDef(s)
	}
}

func (s *ProcedureDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitProcedureDef(s)
	}
}

func (p *ONCRPCv2Parser) ProcedureDef() (localctx IProcedureDefContext) {
	localctx = NewProcedureDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ONCRPCv2ParserRULE_procedureDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.ProcReturn()

	p.SetState(75)
	p.Match(ONCRPCv2ParserIDENTIFIER)

	p.SetState(76)
	p.Match(ONCRPCv2ParserT__6)

	p.SetState(77)
	p.ProcFirstArg()

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ONCRPCv2ParserT__7 {
		p.SetState(78)
		p.Match(ONCRPCv2ParserT__7)

		p.SetState(79)
		p.TypeSpecifier()

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(85)
	p.Match(ONCRPCv2ParserT__8)

	p.SetState(86)
	p.Match(ONCRPCv2ParserT__3)

	p.SetState(87)
	p.Constant()

	p.SetState(88)
	p.Match(ONCRPCv2ParserT__4)

	return localctx
}

// IProcReturnContext is an interface to support dynamic dispatch.
type IProcReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcReturnContext differentiates from other interfaces.
	IsProcReturnContext()
}

type ProcReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcReturnContext() *ProcReturnContext {
	var p = new(ProcReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_procReturn
	return p
}

func (*ProcReturnContext) IsProcReturnContext() {}

func NewProcReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcReturnContext {
	var p = new(ProcReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_procReturn

	return p
}

func (s *ProcReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcReturnContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterProcReturn(s)
	}
}

func (s *ProcReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitProcReturn(s)
	}
}

func (p *ONCRPCv2Parser) ProcReturn() (localctx IProcReturnContext) {
	localctx = NewProcReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ONCRPCv2ParserRULE_procReturn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(92)

	switch p.GetTokenStream().LA(1) {
	case ONCRPCv2ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(90)
		p.Match(ONCRPCv2ParserT__9)

	case ONCRPCv2ParserT__17, ONCRPCv2ParserT__18, ONCRPCv2ParserT__19, ONCRPCv2ParserT__20, ONCRPCv2ParserT__21, ONCRPCv2ParserT__22, ONCRPCv2ParserT__23, ONCRPCv2ParserT__24, ONCRPCv2ParserT__25, ONCRPCv2ParserT__26, ONCRPCv2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(91)
		p.TypeSpecifier()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProcFirstArgContext is an interface to support dynamic dispatch.
type IProcFirstArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcFirstArgContext differentiates from other interfaces.
	IsProcFirstArgContext()
}

type ProcFirstArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcFirstArgContext() *ProcFirstArgContext {
	var p = new(ProcFirstArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_procFirstArg
	return p
}

func (*ProcFirstArgContext) IsProcFirstArgContext() {}

func NewProcFirstArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcFirstArgContext {
	var p = new(ProcFirstArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_procFirstArg

	return p
}

func (s *ProcFirstArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcFirstArgContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcFirstArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcFirstArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcFirstArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterProcFirstArg(s)
	}
}

func (s *ProcFirstArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitProcFirstArg(s)
	}
}

func (p *ONCRPCv2Parser) ProcFirstArg() (localctx IProcFirstArgContext) {
	localctx = NewProcFirstArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ONCRPCv2ParserRULE_procFirstArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(96)

	switch p.GetTokenStream().LA(1) {
	case ONCRPCv2ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(94)
		p.Match(ONCRPCv2ParserT__9)

	case ONCRPCv2ParserT__17, ONCRPCv2ParserT__18, ONCRPCv2ParserT__19, ONCRPCv2ParserT__20, ONCRPCv2ParserT__21, ONCRPCv2ParserT__22, ONCRPCv2ParserT__23, ONCRPCv2ParserT__24, ONCRPCv2ParserT__25, ONCRPCv2ParserT__26, ONCRPCv2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(95)
		p.TypeSpecifier()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOncrpcv2SpecificationContext is an interface to support dynamic dispatch.
type IOncrpcv2SpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOncrpcv2SpecificationContext differentiates from other interfaces.
	IsOncrpcv2SpecificationContext()
}

type Oncrpcv2SpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOncrpcv2SpecificationContext() *Oncrpcv2SpecificationContext {
	var p = new(Oncrpcv2SpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_oncrpcv2Specification
	return p
}

func (*Oncrpcv2SpecificationContext) IsOncrpcv2SpecificationContext() {}

func NewOncrpcv2SpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Oncrpcv2SpecificationContext {
	var p = new(Oncrpcv2SpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_oncrpcv2Specification

	return p
}

func (s *Oncrpcv2SpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Oncrpcv2SpecificationContext) AllXdrSpecification() []IXdrSpecificationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXdrSpecificationContext)(nil)).Elem())
	var tst = make([]IXdrSpecificationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXdrSpecificationContext)
		}
	}

	return tst
}

func (s *Oncrpcv2SpecificationContext) XdrSpecification(i int) IXdrSpecificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXdrSpecificationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXdrSpecificationContext)
}

func (s *Oncrpcv2SpecificationContext) AllProgramDef() []IProgramDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProgramDefContext)(nil)).Elem())
	var tst = make([]IProgramDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProgramDefContext)
		}
	}

	return tst
}

func (s *Oncrpcv2SpecificationContext) ProgramDef(i int) IProgramDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProgramDefContext)
}

func (s *Oncrpcv2SpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Oncrpcv2SpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Oncrpcv2SpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterOncrpcv2Specification(s)
	}
}

func (s *Oncrpcv2SpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitOncrpcv2Specification(s)
	}
}

func (p *ONCRPCv2Parser) Oncrpcv2Specification() (localctx IOncrpcv2SpecificationContext) {
	localctx = NewOncrpcv2SpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ONCRPCv2ParserRULE_oncrpcv2Specification)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ONCRPCv2ParserT__0)|(1<<ONCRPCv2ParserT__24)|(1<<ONCRPCv2ParserT__25)|(1<<ONCRPCv2ParserT__26))) != 0) || _la == ONCRPCv2ParserT__31 || _la == ONCRPCv2ParserT__32 {
		p.SetState(100)

		switch p.GetTokenStream().LA(1) {
		case ONCRPCv2ParserT__24, ONCRPCv2ParserT__25, ONCRPCv2ParserT__26, ONCRPCv2ParserT__31, ONCRPCv2ParserT__32:
			p.SetState(98)
			p.XdrSpecification()

		case ONCRPCv2ParserT__0:
			p.SetState(99)
			p.ProgramDef()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CopyFrom(ctx *DeclarationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclOpaqueFixedContext struct {
	*DeclarationContext
}

func NewDeclOpaqueFixedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclOpaqueFixedContext {
	var p = new(DeclOpaqueFixedContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclOpaqueFixedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclOpaqueFixedContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclOpaqueFixedContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclOpaqueFixedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclOpaqueFixed(s)
	}
}

func (s *DeclOpaqueFixedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclOpaqueFixed(s)
	}
}

type DeclVarContext struct {
	*DeclarationContext
}

func NewDeclVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclVarContext {
	var p = new(DeclVarContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclVarContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclVarContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclVarContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclVar(s)
	}
}

func (s *DeclVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclVar(s)
	}
}

type DeclOptionalContext struct {
	*DeclarationContext
}

func NewDeclOptionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclOptionalContext {
	var p = new(DeclOptionalContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclOptionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclOptionalContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclOptionalContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclOptionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclOptional(s)
	}
}

func (s *DeclOptionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclOptional(s)
	}
}

type DeclStringContext struct {
	*DeclarationContext
}

func NewDeclStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclStringContext {
	var p = new(DeclStringContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclStringContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclStringContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclString(s)
	}
}

func (s *DeclStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclString(s)
	}
}

type DeclFixedContext struct {
	*DeclarationContext
}

func NewDeclFixedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclFixedContext {
	var p = new(DeclFixedContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclFixedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclFixedContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclFixedContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclFixedContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclFixedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclFixed(s)
	}
}

func (s *DeclFixedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclFixed(s)
	}
}

type DeclVoidContext struct {
	*DeclarationContext
}

func NewDeclVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclVoidContext {
	var p = new(DeclVoidContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclVoid(s)
	}
}

func (s *DeclVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclVoid(s)
	}
}

type DeclSimpleContext struct {
	*DeclarationContext
}

func NewDeclSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclSimpleContext {
	var p = new(DeclSimpleContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclSimpleContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclSimpleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclSimple(s)
	}
}

func (s *DeclSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclSimple(s)
	}
}

type DeclOpaqueVarContext struct {
	*DeclarationContext
}

func NewDeclOpaqueVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclOpaqueVarContext {
	var p = new(DeclOpaqueVarContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *DeclOpaqueVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclOpaqueVarContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *DeclOpaqueVarContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclOpaqueVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDeclOpaqueVar(s)
	}
}

func (s *DeclOpaqueVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDeclOpaqueVar(s)
	}
}

func (p *ONCRPCv2Parser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ONCRPCv2ParserRULE_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	la_ := p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	switch la_ {
	case 1:
		localctx = NewDeclSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(105)
		p.TypeSpecifier()

		p.SetState(106)
		p.Match(ONCRPCv2ParserIDENTIFIER)

	case 2:
		localctx = NewDeclFixedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(108)
		p.TypeSpecifier()

		p.SetState(109)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(110)
		p.Match(ONCRPCv2ParserT__10)

		p.SetState(111)
		p.Value()

		p.SetState(112)
		p.Match(ONCRPCv2ParserT__11)

	case 3:
		localctx = NewDeclVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(114)
		p.TypeSpecifier()

		p.SetState(115)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(116)
		p.Match(ONCRPCv2ParserT__12)

		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ONCRPCv2ParserOCTAL-35))|(1<<(ONCRPCv2ParserDECIMAL-35))|(1<<(ONCRPCv2ParserHEXADECIMAL-35))|(1<<(ONCRPCv2ParserBOOLEAN-35))|(1<<(ONCRPCv2ParserIDENTIFIER-35)))) != 0 {
			p.SetState(117)
			p.Value()

		}
		p.SetState(120)
		p.Match(ONCRPCv2ParserT__13)

	case 4:
		localctx = NewDeclOpaqueFixedContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(122)
		p.Match(ONCRPCv2ParserT__14)

		p.SetState(123)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(124)
		p.Match(ONCRPCv2ParserT__10)

		p.SetState(125)
		p.Value()

		p.SetState(126)
		p.Match(ONCRPCv2ParserT__11)

	case 5:
		localctx = NewDeclOpaqueVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(128)
		p.Match(ONCRPCv2ParserT__14)

		p.SetState(129)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(130)
		p.Match(ONCRPCv2ParserT__12)

		p.SetState(132)
		_la = p.GetTokenStream().LA(1)

		if ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ONCRPCv2ParserOCTAL-35))|(1<<(ONCRPCv2ParserDECIMAL-35))|(1<<(ONCRPCv2ParserHEXADECIMAL-35))|(1<<(ONCRPCv2ParserBOOLEAN-35))|(1<<(ONCRPCv2ParserIDENTIFIER-35)))) != 0 {
			p.SetState(131)
			p.Value()

		}
		p.SetState(134)
		p.Match(ONCRPCv2ParserT__13)

	case 6:
		localctx = NewDeclStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(135)
		p.Match(ONCRPCv2ParserT__15)

		p.SetState(136)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(137)
		p.Match(ONCRPCv2ParserT__12)

		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ONCRPCv2ParserOCTAL-35))|(1<<(ONCRPCv2ParserDECIMAL-35))|(1<<(ONCRPCv2ParserHEXADECIMAL-35))|(1<<(ONCRPCv2ParserBOOLEAN-35))|(1<<(ONCRPCv2ParserIDENTIFIER-35)))) != 0 {
			p.SetState(138)
			p.Value()

		}
		p.SetState(141)
		p.Match(ONCRPCv2ParserT__13)

	case 7:
		localctx = NewDeclOptionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		p.SetState(142)
		p.TypeSpecifier()

		p.SetState(143)
		p.Match(ONCRPCv2ParserT__16)

		p.SetState(144)
		p.Match(ONCRPCv2ParserIDENTIFIER)

	case 8:
		localctx = NewDeclVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(146)
		p.Match(ONCRPCv2ParserT__9)

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ONCRPCv2Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ONCRPCv2ParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(151)

	switch p.GetTokenStream().LA(1) {
	case ONCRPCv2ParserOCTAL, ONCRPCv2ParserDECIMAL, ONCRPCv2ParserHEXADECIMAL, ONCRPCv2ParserBOOLEAN:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(149)
		p.Constant()

	case ONCRPCv2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(150)
		p.Match(ONCRPCv2ParserIDENTIFIER)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserDECIMAL, 0)
}

func (s *ConstantContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserHEXADECIMAL, 0)
}

func (s *ConstantContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserOCTAL, 0)
}

func (s *ConstantContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserBOOLEAN, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *ONCRPCv2Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ONCRPCv2ParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ONCRPCv2ParserOCTAL-35))|(1<<(ONCRPCv2ParserDECIMAL-35))|(1<<(ONCRPCv2ParserHEXADECIMAL-35))|(1<<(ONCRPCv2ParserBOOLEAN-35)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.Consume()
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) CopyFrom(ctx *TypeSpecifierContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeSpecStructContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecStructContext {
	var p = new(TypeSpecStructContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecStructContext) StructTypeSpec() IStructTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructTypeSpecContext)
}

func (s *TypeSpecStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecStruct(s)
	}
}

func (s *TypeSpecStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecStruct(s)
	}
}

type TypeSpecHyperContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecHyperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecHyperContext {
	var p = new(TypeSpecHyperContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecHyperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecHyperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecHyper(s)
	}
}

func (s *TypeSpecHyperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecHyper(s)
	}
}

type TypeSpecUnionContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecUnionContext {
	var p = new(TypeSpecUnionContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecUnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecUnionContext) UnionTypeSpec() IUnionTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeSpecContext)
}

func (s *TypeSpecUnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecUnion(s)
	}
}

func (s *TypeSpecUnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecUnion(s)
	}
}

type TypeSpecFloatContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecFloatContext {
	var p = new(TypeSpecFloatContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecFloat(s)
	}
}

func (s *TypeSpecFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecFloat(s)
	}
}

type TypeSpecEnumContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecEnumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecEnumContext {
	var p = new(TypeSpecEnumContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecEnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecEnumContext) EnumTypeSpec() IEnumTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeSpecContext)
}

func (s *TypeSpecEnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecEnum(s)
	}
}

func (s *TypeSpecEnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecEnum(s)
	}
}

type TypeSpecIntContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecIntContext {
	var p = new(TypeSpecIntContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecInt(s)
	}
}

func (s *TypeSpecIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecInt(s)
	}
}

type TypeSpecDoubleContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecDoubleContext {
	var p = new(TypeSpecDoubleContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecDoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecDoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecDouble(s)
	}
}

func (s *TypeSpecDoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecDouble(s)
	}
}

type TypeSpecQuadContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecQuadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecQuadContext {
	var p = new(TypeSpecQuadContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecQuadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecQuadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecQuad(s)
	}
}

func (s *TypeSpecQuadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecQuad(s)
	}
}

type TypeSpecStructIdentifierContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecStructIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecStructIdentifierContext {
	var p = new(TypeSpecStructIdentifierContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecStructIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecStructIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *TypeSpecStructIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecStructIdentifier(s)
	}
}

func (s *TypeSpecStructIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecStructIdentifier(s)
	}
}

type TypeSpecIdentifierContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecIdentifierContext {
	var p = new(TypeSpecIdentifierContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *TypeSpecIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecIdentifier(s)
	}
}

func (s *TypeSpecIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecIdentifier(s)
	}
}

type TypeSpecBoolContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecBoolContext {
	var p = new(TypeSpecBoolContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeSpecBool(s)
	}
}

func (s *TypeSpecBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeSpecBool(s)
	}
}

func (p *ONCRPCv2Parser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ONCRPCv2ParserRULE_typeSpecifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	la_ := p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	switch la_ {
	case 1:
		localctx = NewTypeSpecIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(160)
		p.GetErrorHandler().Sync(p)

		la_ := p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

		switch la_ {
		case 1:
			p.SetState(155)
			p.Match(ONCRPCv2ParserT__17)

		case 2:
			p.SetState(157)
			_la = p.GetTokenStream().LA(1)

			if _la == ONCRPCv2ParserT__17 {
				p.SetState(156)
				p.Match(ONCRPCv2ParserT__17)

			}
			p.SetState(159)
			p.Match(ONCRPCv2ParserT__18)

		}

	case 2:
		localctx = NewTypeSpecHyperContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if _la == ONCRPCv2ParserT__17 {
			p.SetState(162)
			p.Match(ONCRPCv2ParserT__17)

		}
		p.SetState(165)
		p.Match(ONCRPCv2ParserT__19)

	case 3:
		localctx = NewTypeSpecFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(166)
		p.Match(ONCRPCv2ParserT__20)

	case 4:
		localctx = NewTypeSpecDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(167)
		p.Match(ONCRPCv2ParserT__21)

	case 5:
		localctx = NewTypeSpecQuadContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(168)
		p.Match(ONCRPCv2ParserT__22)

	case 6:
		localctx = NewTypeSpecBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(169)
		p.Match(ONCRPCv2ParserT__23)

	case 7:
		localctx = NewTypeSpecStructIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		p.SetState(170)
		p.Match(ONCRPCv2ParserT__24)

		p.SetState(171)
		p.Match(ONCRPCv2ParserIDENTIFIER)

	case 8:
		localctx = NewTypeSpecEnumContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(172)
		p.EnumTypeSpec()

	case 9:
		localctx = NewTypeSpecStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		p.SetState(173)
		p.StructTypeSpec()

	case 10:
		localctx = NewTypeSpecUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		p.SetState(174)
		p.UnionTypeSpec()

	case 11:
		localctx = NewTypeSpecIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		p.SetState(175)
		p.Match(ONCRPCv2ParserIDENTIFIER)

	}

	return localctx
}

// IEnumTypeSpecContext is an interface to support dynamic dispatch.
type IEnumTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeSpecContext differentiates from other interfaces.
	IsEnumTypeSpecContext()
}

type EnumTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeSpecContext() *EnumTypeSpecContext {
	var p = new(EnumTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_enumTypeSpec
	return p
}

func (*EnumTypeSpecContext) IsEnumTypeSpecContext() {}

func NewEnumTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeSpecContext {
	var p = new(EnumTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_enumTypeSpec

	return p
}

func (s *EnumTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeSpecContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterEnumTypeSpec(s)
	}
}

func (s *EnumTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitEnumTypeSpec(s)
	}
}

func (p *ONCRPCv2Parser) EnumTypeSpec() (localctx IEnumTypeSpecContext) {
	localctx = NewEnumTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ONCRPCv2ParserRULE_enumTypeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(178)
	p.Match(ONCRPCv2ParserT__25)

	p.SetState(179)
	p.EnumBody()

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ONCRPCv2ParserIDENTIFIER)
}

func (s *EnumBodyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, i)
}

func (s *EnumBodyContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *ONCRPCv2Parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ONCRPCv2ParserRULE_enumBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.Match(ONCRPCv2ParserT__1)

	p.SetState(182)
	p.Match(ONCRPCv2ParserIDENTIFIER)

	p.SetState(183)
	p.Match(ONCRPCv2ParserT__3)

	p.SetState(184)
	p.Value()

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ONCRPCv2ParserT__7 {
		p.SetState(186)
		p.Match(ONCRPCv2ParserT__7)

		p.SetState(187)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(188)
		p.Match(ONCRPCv2ParserT__3)

		p.SetState(189)
		p.Value()

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(195)
	p.Match(ONCRPCv2ParserT__2)

	return localctx
}

// IStructTypeSpecContext is an interface to support dynamic dispatch.
type IStructTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructTypeSpecContext differentiates from other interfaces.
	IsStructTypeSpecContext()
}

type StructTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeSpecContext() *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_structTypeSpec
	return p
}

func (*StructTypeSpecContext) IsStructTypeSpecContext() {}

func NewStructTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_structTypeSpec

	return p
}

func (s *StructTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeSpecContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitStructTypeSpec(s)
	}
}

func (p *ONCRPCv2Parser) StructTypeSpec() (localctx IStructTypeSpecContext) {
	localctx = NewStructTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ONCRPCv2ParserRULE_structTypeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(197)
	p.Match(ONCRPCv2ParserT__24)

	p.SetState(198)
	p.StructBody()

	return localctx
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_structBody
	return p
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *StructBodyContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (p *ONCRPCv2Parser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ONCRPCv2ParserRULE_structBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(200)
	p.Match(ONCRPCv2ParserT__1)

	p.SetState(201)
	p.Declaration()

	p.SetState(202)
	p.Match(ONCRPCv2ParserT__4)

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(ONCRPCv2ParserT__9-10))|(1<<(ONCRPCv2ParserT__14-10))|(1<<(ONCRPCv2ParserT__15-10))|(1<<(ONCRPCv2ParserT__17-10))|(1<<(ONCRPCv2ParserT__18-10))|(1<<(ONCRPCv2ParserT__19-10))|(1<<(ONCRPCv2ParserT__20-10))|(1<<(ONCRPCv2ParserT__21-10))|(1<<(ONCRPCv2ParserT__22-10))|(1<<(ONCRPCv2ParserT__23-10))|(1<<(ONCRPCv2ParserT__24-10))|(1<<(ONCRPCv2ParserT__25-10))|(1<<(ONCRPCv2ParserT__26-10))|(1<<(ONCRPCv2ParserIDENTIFIER-10)))) != 0 {
		p.SetState(204)
		p.Declaration()

		p.SetState(205)
		p.Match(ONCRPCv2ParserT__4)

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(212)
	p.Match(ONCRPCv2ParserT__2)

	return localctx
}

// IUnionTypeSpecContext is an interface to support dynamic dispatch.
type IUnionTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeSpecContext differentiates from other interfaces.
	IsUnionTypeSpecContext()
}

type UnionTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeSpecContext() *UnionTypeSpecContext {
	var p = new(UnionTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_unionTypeSpec
	return p
}

func (*UnionTypeSpecContext) IsUnionTypeSpecContext() {}

func NewUnionTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeSpecContext {
	var p = new(UnionTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_unionTypeSpec

	return p
}

func (s *UnionTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeSpecContext) UnionBody() IUnionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *UnionTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterUnionTypeSpec(s)
	}
}

func (s *UnionTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitUnionTypeSpec(s)
	}
}

func (p *ONCRPCv2Parser) UnionTypeSpec() (localctx IUnionTypeSpecContext) {
	localctx = NewUnionTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ONCRPCv2ParserRULE_unionTypeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(214)
	p.Match(ONCRPCv2ParserT__26)

	p.SetState(215)
	p.UnionBody()

	return localctx
}

// IUnionBodyContext is an interface to support dynamic dispatch.
type IUnionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionBodyContext differentiates from other interfaces.
	IsUnionBodyContext()
}

type UnionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionBodyContext() *UnionBodyContext {
	var p = new(UnionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_unionBody
	return p
}

func (*UnionBodyContext) IsUnionBodyContext() {}

func NewUnionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionBodyContext {
	var p = new(UnionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_unionBody

	return p
}

func (s *UnionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionBodyContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *UnionBodyContext) AllCaseSpec() []ICaseSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseSpecContext)(nil)).Elem())
	var tst = make([]ICaseSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseSpecContext)
		}
	}

	return tst
}

func (s *UnionBodyContext) CaseSpec(i int) ICaseSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseSpecContext)
}

func (s *UnionBodyContext) DefaultCaseSpec() IDefaultCaseSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultCaseSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultCaseSpecContext)
}

func (s *UnionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterUnionBody(s)
	}
}

func (s *UnionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitUnionBody(s)
	}
}

func (p *ONCRPCv2Parser) UnionBody() (localctx IUnionBodyContext) {
	localctx = NewUnionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ONCRPCv2ParserRULE_unionBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(217)
	p.Match(ONCRPCv2ParserT__27)

	p.SetState(218)
	p.Match(ONCRPCv2ParserT__6)

	p.SetState(219)
	p.Declaration()

	p.SetState(220)
	p.Match(ONCRPCv2ParserT__8)

	p.SetState(221)
	p.Match(ONCRPCv2ParserT__1)

	p.SetState(222)
	p.CaseSpec()

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ONCRPCv2ParserT__28 {
		p.SetState(223)
		p.CaseSpec()

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(230)
	_la = p.GetTokenStream().LA(1)

	if _la == ONCRPCv2ParserT__30 {
		p.SetState(229)
		p.DefaultCaseSpec()

	}
	p.SetState(232)
	p.Match(ONCRPCv2ParserT__2)

	return localctx
}

// ICaseSpecContext is an interface to support dynamic dispatch.
type ICaseSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseSpecContext differentiates from other interfaces.
	IsCaseSpecContext()
}

type CaseSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseSpecContext() *CaseSpecContext {
	var p = new(CaseSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_caseSpec
	return p
}

func (*CaseSpecContext) IsCaseSpecContext() {}

func NewCaseSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseSpecContext {
	var p = new(CaseSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_caseSpec

	return p
}

func (s *CaseSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseSpecContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *CaseSpecContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *CaseSpecContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CaseSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterCaseSpec(s)
	}
}

func (s *CaseSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitCaseSpec(s)
	}
}

func (p *ONCRPCv2Parser) CaseSpec() (localctx ICaseSpecContext) {
	localctx = NewCaseSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ONCRPCv2ParserRULE_caseSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(234)
	p.Match(ONCRPCv2ParserT__28)

	p.SetState(235)
	p.Value()

	p.SetState(236)
	p.Match(ONCRPCv2ParserT__29)

	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ONCRPCv2ParserT__28 {
		p.SetState(238)
		p.Match(ONCRPCv2ParserT__28)

		p.SetState(239)
		p.Value()

		p.SetState(240)
		p.Match(ONCRPCv2ParserT__29)

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(247)
	p.Declaration()

	p.SetState(248)
	p.Match(ONCRPCv2ParserT__4)

	return localctx
}

// IDefaultCaseSpecContext is an interface to support dynamic dispatch.
type IDefaultCaseSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultCaseSpecContext differentiates from other interfaces.
	IsDefaultCaseSpecContext()
}

type DefaultCaseSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultCaseSpecContext() *DefaultCaseSpecContext {
	var p = new(DefaultCaseSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_defaultCaseSpec
	return p
}

func (*DefaultCaseSpecContext) IsDefaultCaseSpecContext() {}

func NewDefaultCaseSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseSpecContext {
	var p = new(DefaultCaseSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_defaultCaseSpec

	return p
}

func (s *DefaultCaseSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultCaseSpecContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DefaultCaseSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultCaseSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDefaultCaseSpec(s)
	}
}

func (s *DefaultCaseSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDefaultCaseSpec(s)
	}
}

func (p *ONCRPCv2Parser) DefaultCaseSpec() (localctx IDefaultCaseSpecContext) {
	localctx = NewDefaultCaseSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ONCRPCv2ParserRULE_defaultCaseSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	p.Match(ONCRPCv2ParserT__30)

	p.SetState(251)
	p.Match(ONCRPCv2ParserT__29)

	p.SetState(252)
	p.Declaration()

	p.SetState(253)
	p.Match(ONCRPCv2ParserT__4)

	return localctx
}

// IConstantDefContext is an interface to support dynamic dispatch.
type IConstantDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantDefContext differentiates from other interfaces.
	IsConstantDefContext()
}

type ConstantDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDefContext() *ConstantDefContext {
	var p = new(ConstantDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_constantDef
	return p
}

func (*ConstantDefContext) IsConstantDefContext() {}

func NewConstantDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefContext {
	var p = new(ConstantDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_constantDef

	return p
}

func (s *ConstantDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *ConstantDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterConstantDef(s)
	}
}

func (s *ConstantDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitConstantDef(s)
	}
}

func (p *ONCRPCv2Parser) ConstantDef() (localctx IConstantDefContext) {
	localctx = NewConstantDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ONCRPCv2ParserRULE_constantDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.Match(ONCRPCv2ParserT__31)

	p.SetState(256)
	p.Match(ONCRPCv2ParserIDENTIFIER)

	p.SetState(257)
	p.Match(ONCRPCv2ParserT__3)

	p.SetState(258)
	p.Constant()

	p.SetState(259)
	p.Match(ONCRPCv2ParserT__4)

	return localctx
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) CopyFrom(ctx *TypeDefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeDefEnumContext struct {
	*TypeDefContext
}

func NewTypeDefEnumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefEnumContext {
	var p = new(TypeDefEnumContext)

	p.TypeDefContext = NewEmptyTypeDefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeDefContext))

	return p
}

func (s *TypeDefEnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefEnumContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *TypeDefEnumContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *TypeDefEnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeDefEnum(s)
	}
}

func (s *TypeDefEnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeDefEnum(s)
	}
}

type TypeDefStructContext struct {
	*TypeDefContext
}

func NewTypeDefStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefStructContext {
	var p = new(TypeDefStructContext)

	p.TypeDefContext = NewEmptyTypeDefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeDefContext))

	return p
}

func (s *TypeDefStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefStructContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *TypeDefStructContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *TypeDefStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeDefStruct(s)
	}
}

func (s *TypeDefStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeDefStruct(s)
	}
}

type TypeDefTypedefContext struct {
	*TypeDefContext
}

func NewTypeDefTypedefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefTypedefContext {
	var p = new(TypeDefTypedefContext)

	p.TypeDefContext = NewEmptyTypeDefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeDefContext))

	return p
}

func (s *TypeDefTypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefTypedefContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TypeDefTypedefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeDefTypedef(s)
	}
}

func (s *TypeDefTypedefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeDefTypedef(s)
	}
}

type TypeDefUnionContext struct {
	*TypeDefContext
}

func NewTypeDefUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefUnionContext {
	var p = new(TypeDefUnionContext)

	p.TypeDefContext = NewEmptyTypeDefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeDefContext))

	return p
}

func (s *TypeDefUnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefUnionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ONCRPCv2ParserIDENTIFIER, 0)
}

func (s *TypeDefUnionContext) UnionBody() IUnionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *TypeDefUnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterTypeDefUnion(s)
	}
}

func (s *TypeDefUnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitTypeDefUnion(s)
	}
}

func (p *ONCRPCv2Parser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ONCRPCv2ParserRULE_typeDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(280)

	switch p.GetTokenStream().LA(1) {
	case ONCRPCv2ParserT__32:
		localctx = NewTypeDefTypedefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(261)
		p.Match(ONCRPCv2ParserT__32)

		p.SetState(262)
		p.Declaration()

		p.SetState(263)
		p.Match(ONCRPCv2ParserT__4)

	case ONCRPCv2ParserT__25:
		localctx = NewTypeDefEnumContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(265)
		p.Match(ONCRPCv2ParserT__25)

		p.SetState(266)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(267)
		p.EnumBody()

		p.SetState(268)
		p.Match(ONCRPCv2ParserT__4)

	case ONCRPCv2ParserT__24:
		localctx = NewTypeDefStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(270)
		p.Match(ONCRPCv2ParserT__24)

		p.SetState(271)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(272)
		p.StructBody()

		p.SetState(273)
		p.Match(ONCRPCv2ParserT__4)

	case ONCRPCv2ParserT__26:
		localctx = NewTypeDefUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(275)
		p.Match(ONCRPCv2ParserT__26)

		p.SetState(276)
		p.Match(ONCRPCv2ParserIDENTIFIER)

		p.SetState(277)
		p.UnionBody()

		p.SetState(278)
		p.Match(ONCRPCv2ParserT__4)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) TypeDef() ITypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *DefinitionContext) ConstantDef() IConstantDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantDefContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *ONCRPCv2Parser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ONCRPCv2ParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(284)

	switch p.GetTokenStream().LA(1) {
	case ONCRPCv2ParserT__24, ONCRPCv2ParserT__25, ONCRPCv2ParserT__26, ONCRPCv2ParserT__32:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(282)
		p.TypeDef()

	case ONCRPCv2ParserT__31:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(283)
		p.ConstantDef()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IXdrSpecificationContext is an interface to support dynamic dispatch.
type IXdrSpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXdrSpecificationContext differentiates from other interfaces.
	IsXdrSpecificationContext()
}

type XdrSpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXdrSpecificationContext() *XdrSpecificationContext {
	var p = new(XdrSpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ONCRPCv2ParserRULE_xdrSpecification
	return p
}

func (*XdrSpecificationContext) IsXdrSpecificationContext() {}

func NewXdrSpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XdrSpecificationContext {
	var p = new(XdrSpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ONCRPCv2ParserRULE_xdrSpecification

	return p
}

func (s *XdrSpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *XdrSpecificationContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *XdrSpecificationContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *XdrSpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XdrSpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XdrSpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.EnterXdrSpecification(s)
	}
}

func (s *XdrSpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ONCRPCv2Listener); ok {
		listenerT.ExitXdrSpecification(s)
	}
}

func (p *ONCRPCv2Parser) XdrSpecification() (localctx IXdrSpecificationContext) {
	localctx = NewXdrSpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ONCRPCv2ParserRULE_xdrSpecification)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_alt := 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(286)
			p.Definition()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}
