package core

import (
    "bytes"
)

/** Generic parser class.
* All parsers inherit this class.
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ParserCore struct {
    //public static final boolean debug = Debug.parserDebug;
    nesting_level int //protected static int

    lexer *LexerCore
}

func NewParserCore(buffer string) *ParserCore {
    this := &ParserCore{}

    this.lexer = NewLexerCore("CharLexer", buffer)

    return this
}

func (this *ParserCore) GetNameValue(separator byte) *NameValue {
    if Debug.parserDebug {
        this.Dbg_enter("nameValue")
        defer this.Dbg_leave("nameValue")
    }

    this.lexer.Match(LexerCore_ID)
    name := this.lexer.GetNextToken()
    // eat white space.
    this.lexer.SPorHT()
    //try {
    quoted := false
    la, err := this.lexer.LookAheadK(0)
    if la == separator && err == nil {
        this.lexer.ConsumeK(1)
        this.lexer.SPorHT()
        var str string
        if la, err = this.lexer.LookAheadK(0); la == '"' && err == nil {
            str, _ = this.lexer.QuotedString()
            quoted = true
        } else {
            this.lexer.Match(LexerCore_ID)
            value := this.lexer.GetNextToken()
            str = value.tokenValue
        }
        nv := NewNameValue(name.tokenValue, str)
        if quoted {
            nv.SetQuotedValue()
        }
        return nv
    }   //else {
    return NewNameValue(name.tokenValue, "")
    //}
    //} catch (ParseException ex) {
    //	return new NameValue(name.tokenValue,null);
    //}
}

func (this *ParserCore) Dbg_enter(rule string) {
    var stringBuffer bytes.Buffer //= new StringBuffer();
    for i := 0; i < this.nesting_level; i++ {
        stringBuffer.WriteString(">")
    }
    if Debug.parserDebug {
        println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
    }
    this.nesting_level++
}

func (this *ParserCore) Dbg_leave(rule string) {
    var stringBuffer bytes.Buffer //= new StringBuffer();
    for i := 0; i < this.nesting_level; i++ {
        stringBuffer.WriteString("<")
    }
    if Debug.parserDebug {
        println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
    }
    this.nesting_level--
}

/*func (this *ParserCore)  NameValue nameValue() * {
	return nameValue('=');
}*/

func (this *ParserCore) PeekLine(rule string) {
    if Debug.parserDebug {
        Debug.println(rule + " " + this.lexer.PeekLine())
    }
}
