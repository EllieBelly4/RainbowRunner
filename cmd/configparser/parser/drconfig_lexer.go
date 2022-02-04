// Code generated from C:/Users/Sophie/go/src/RainbowRunner/antlr\DRConfig.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 169,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 54, 10, 2, 12, 2, 14,
	2, 57, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 65, 10, 3, 12,
	3, 14, 3, 68, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 18, 6, 18, 128, 10, 18, 13, 18, 14, 18, 129, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 7, 20, 138, 10, 20, 12, 20, 14, 20,
	141, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 147, 10, 21, 12, 21, 14,
	21, 150, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 6, 22, 156, 10, 22, 13, 22,
	14, 22, 157, 3, 23, 5, 23, 161, 10, 23, 3, 23, 6, 23, 164, 10, 23, 13,
	23, 14, 23, 165, 3, 24, 3, 24, 5, 66, 139, 148, 2, 25, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 3, 2, 9, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34,
	3, 2, 41, 41, 3, 2, 36, 36, 8, 2, 41, 41, 47, 47, 50, 59, 67, 92, 97, 97,
	99, 124, 3, 2, 47, 47, 4, 2, 48, 48, 50, 59, 2, 177, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5,
	60, 3, 2, 2, 2, 7, 74, 3, 2, 2, 2, 9, 80, 3, 2, 2, 2, 11, 88, 3, 2, 2,
	2, 13, 95, 3, 2, 2, 2, 15, 105, 3, 2, 2, 2, 17, 107, 3, 2, 2, 2, 19, 109,
	3, 2, 2, 2, 21, 111, 3, 2, 2, 2, 23, 113, 3, 2, 2, 2, 25, 115, 3, 2, 2,
	2, 27, 117, 3, 2, 2, 2, 29, 119, 3, 2, 2, 2, 31, 121, 3, 2, 2, 2, 33, 123,
	3, 2, 2, 2, 35, 127, 3, 2, 2, 2, 37, 133, 3, 2, 2, 2, 39, 135, 3, 2, 2,
	2, 41, 144, 3, 2, 2, 2, 43, 155, 3, 2, 2, 2, 45, 160, 3, 2, 2, 2, 47, 167,
	3, 2, 2, 2, 49, 50, 7, 49, 2, 2, 50, 51, 7, 49, 2, 2, 51, 55, 3, 2, 2,
	2, 52, 54, 10, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53,
	3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2,
	58, 59, 8, 2, 2, 2, 59, 4, 3, 2, 2, 2, 60, 61, 7, 49, 2, 2, 61, 62, 7,
	44, 2, 2, 62, 66, 3, 2, 2, 2, 63, 65, 11, 2, 2, 2, 64, 63, 3, 2, 2, 2,
	65, 68, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 69, 3,
	2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7, 44, 2, 2, 70, 71, 7, 49, 2, 2,
	71, 72, 3, 2, 2, 2, 72, 73, 8, 3, 2, 2, 73, 6, 3, 2, 2, 2, 74, 75, 5, 45,
	23, 2, 75, 76, 5, 27, 14, 2, 76, 77, 5, 45, 23, 2, 77, 78, 5, 27, 14, 2,
	78, 79, 5, 45, 23, 2, 79, 8, 3, 2, 2, 2, 80, 81, 7, 103, 2, 2, 81, 82,
	7, 122, 2, 2, 82, 83, 7, 118, 2, 2, 83, 84, 7, 103, 2, 2, 84, 85, 7, 112,
	2, 2, 85, 86, 7, 102, 2, 2, 86, 87, 7, 117, 2, 2, 87, 10, 3, 2, 2, 2, 88,
	89, 7, 117, 2, 2, 89, 90, 7, 118, 2, 2, 90, 91, 7, 99, 2, 2, 91, 92, 7,
	118, 2, 2, 92, 93, 7, 107, 2, 2, 93, 94, 7, 101, 2, 2, 94, 12, 3, 2, 2,
	2, 95, 96, 7, 118, 2, 2, 96, 97, 7, 116, 2, 2, 97, 98, 7, 99, 2, 2, 98,
	99, 7, 112, 2, 2, 99, 100, 7, 117, 2, 2, 100, 101, 7, 107, 2, 2, 101, 102,
	7, 103, 2, 2, 102, 103, 7, 112, 2, 2, 103, 104, 7, 118, 2, 2, 104, 14,
	3, 2, 2, 2, 105, 106, 7, 63, 2, 2, 106, 16, 3, 2, 2, 2, 107, 108, 7, 125,
	2, 2, 108, 18, 3, 2, 2, 2, 109, 110, 7, 127, 2, 2, 110, 20, 3, 2, 2, 2,
	111, 112, 7, 35, 2, 2, 112, 22, 3, 2, 2, 2, 113, 114, 7, 61, 2, 2, 114,
	24, 3, 2, 2, 2, 115, 116, 7, 48, 2, 2, 116, 26, 3, 2, 2, 2, 117, 118, 7,
	46, 2, 2, 118, 28, 3, 2, 2, 2, 119, 120, 7, 60, 2, 2, 120, 30, 3, 2, 2,
	2, 121, 122, 7, 44, 2, 2, 122, 32, 3, 2, 2, 2, 123, 124, 7, 49, 2, 2, 124,
	125, 7, 49, 2, 2, 125, 34, 3, 2, 2, 2, 126, 128, 9, 3, 2, 2, 127, 126,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2,
	2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 8, 18, 2, 2, 132, 36, 3, 2, 2, 2,
	133, 134, 7, 12, 2, 2, 134, 38, 3, 2, 2, 2, 135, 139, 9, 4, 2, 2, 136,
	138, 11, 2, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2, 139, 140,
	3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141, 139, 3, 2,
	2, 2, 142, 143, 9, 4, 2, 2, 143, 40, 3, 2, 2, 2, 144, 148, 9, 5, 2, 2,
	145, 147, 11, 2, 2, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148,
	149, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148,
	3, 2, 2, 2, 151, 152, 9, 5, 2, 2, 152, 42, 3, 2, 2, 2, 153, 156, 9, 6,
	2, 2, 154, 156, 5, 25, 13, 2, 155, 153, 3, 2, 2, 2, 155, 154, 3, 2, 2,
	2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158,
	44, 3, 2, 2, 2, 159, 161, 9, 7, 2, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3,
	2, 2, 2, 161, 163, 3, 2, 2, 2, 162, 164, 9, 8, 2, 2, 163, 162, 3, 2, 2,
	2, 164, 165, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	46, 3, 2, 2, 2, 167, 168, 11, 2, 2, 2, 168, 48, 3, 2, 2, 2, 12, 2, 55,
	66, 129, 139, 148, 155, 157, 160, 165, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'extends'", "'static'", "'transient'", "'='", "'{'", "'}'",
	"'!'", "';'", "'.'", "','", "':'", "'*'", "'//'", "", "'\n'",
}

var lexerSymbolicNames = []string{
	"", "COMMENT", "MLCOMMENT", "VECTOR3", "EXTENDS", "STATIC", "TRANSIENT",
	"ASSIGN", "PARENL", "PARENR", "EXCL", "SEMI", "DOT", "COMMA", "COLON",
	"ASTERISK", "SLASHSLASH", "WS", "EOL", "SINGLESTR", "DOUBLESTR", "IDENTIFIER",
	"NUMBER", "ANY",
}

var lexerRuleNames = []string{
	"COMMENT", "MLCOMMENT", "VECTOR3", "EXTENDS", "STATIC", "TRANSIENT", "ASSIGN",
	"PARENL", "PARENR", "EXCL", "SEMI", "DOT", "COMMA", "COLON", "ASTERISK",
	"SLASHSLASH", "WS", "EOL", "SINGLESTR", "DOUBLESTR", "IDENTIFIER", "NUMBER",
	"ANY",
}

type DRConfigLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDRConfigLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DRConfigLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDRConfigLexer(input antlr.CharStream) *DRConfigLexer {
	l := new(DRConfigLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "DRConfig.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DRConfigLexer tokens.
const (
	DRConfigLexerCOMMENT    = 1
	DRConfigLexerMLCOMMENT  = 2
	DRConfigLexerVECTOR3    = 3
	DRConfigLexerEXTENDS    = 4
	DRConfigLexerSTATIC     = 5
	DRConfigLexerTRANSIENT  = 6
	DRConfigLexerASSIGN     = 7
	DRConfigLexerPARENL     = 8
	DRConfigLexerPARENR     = 9
	DRConfigLexerEXCL       = 10
	DRConfigLexerSEMI       = 11
	DRConfigLexerDOT        = 12
	DRConfigLexerCOMMA      = 13
	DRConfigLexerCOLON      = 14
	DRConfigLexerASTERISK   = 15
	DRConfigLexerSLASHSLASH = 16
	DRConfigLexerWS         = 17
	DRConfigLexerEOL        = 18
	DRConfigLexerSINGLESTR  = 19
	DRConfigLexerDOUBLESTR  = 20
	DRConfigLexerIDENTIFIER = 21
	DRConfigLexerNUMBER     = 22
	DRConfigLexerANY        = 23
)
