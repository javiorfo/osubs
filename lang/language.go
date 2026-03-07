package lang

import "fmt"

type Language interface {
	fmt.Stringer
	Code() string
	isLanguage()
}

type Lang uint

const (
	Abkhazian Lang = iota
	Afrikaans
	Albanian
	Amharic
	Arabic
	Aragonese
	Armenian
	Assamese
	Asturian
	Azerbaijani
	Basque
	Belarusian
	Bengali
	Bosnian
	Breton
	Bulgarian
	Burmese
	Catalan
	ChineseCantonese
	ChineseSimplified
	ChineseTraditional
	ChineseBilingual
	Croatian
	Czech
	Danish
	Dari
	Dutch
	English
	Esperanto
	Estonian
	Extremaduran
	Finnish
	French
	Gaelic
	Galician
	Georgian
	German
	Greek
	Hebrew
	Hindi
	Hungarian
	Icelandic
	Igbo
	Indonesian
	Interlingua
	Irish
	Italian
	Japanese
	Kannada
	Kazakh
	Khmer
	Korean
	Kurdish
	Kyrgyz
	Latvian
	Lithuanian
	Luxembourgish
	Macedonian
	Malay
	Malayalam
	Manipuri
	Marathi
	Mongolian
	Montenegrin
	Navajo
	Nepali
	NorthernSami
	Norwegian
	Occitan
	Odia
	Persian
	Polish
	Portuguese
	PortugueseBr
	PortugueseMz
	Pushto
	Romanian
	Russian
	Santali
	Serbian
	Sindhi
	Sinhalese
	Slovak
	Slovenian
	Somali
	SorbianLanguages
	SouthAzerbaijani
	Spanish
	SpanishEU
	SpanishLA
	Swahili
	Swedish
	Syriac
	Tagalog
	Tamil
	Tatar
	Telugu
	Tetum
	Thai
	TokiPona
	Turkish
	Turkmen
	Ukrainian
	Urdu
	Uzbek
	Vietnamese
	Welsch
)

func (l Lang) isLanguage() {}

func (l Lang) Code() string {
	switch l {
	case Abkhazian:
		return "abk"
	case Afrikaans:
		return "afr"
	case Albanian:
		return "alb"
	case Amharic:
		return "Amh"
	case Arabic:
		return "ara"
	case Aragonese:
		return "arg"
	case Armenian:
		return "arm"
	case Assamese:
		return "asm"
	case Asturian:
		return "ast"
	case Azerbaijani:
		return "aze"
	case Basque:
		return "baq"
	case Belarusian:
		return "bel"
	case Bengali:
		return "ben"
	case Bosnian:
		return "bos"
	case Breton:
		return "bre"
	case Bulgarian:
		return "bul"
	case Burmese:
		return "bur"
	case Catalan:
		return "cat"
	case ChineseCantonese:
		return "zhc"
	case ChineseSimplified:
		return "chi"
	case ChineseTraditional:
		return "zht"
	case ChineseBilingual:
		return "zhe"
	case Croatian:
		return "hrv"
	case Czech:
		return "cze"
	case Danish:
		return "dan"
	case Dari:
		return "prs"
	case Dutch:
		return "dut"
	case English:
		return "eng"
	case Esperanto:
		return "epo"
	case Estonian:
		return "est"
	case Extremaduran:
		return "ext"
	case Finnish:
		return "fin"
	case French:
		return "fre"
	case Gaelic:
		return "gla"
	case Galician:
		return "glb"
	case Georgian:
		return "geo"
	case German:
		return "ger"
	case Greek:
		return "ell"
	case Hebrew:
		return "heb"
	case Hindi:
		return "hin"
	case Hungarian:
		return "hun"
	case Icelandic:
		return "ice"
	case Igbo:
		return "ibo"
	case Indonesian:
		return "ind"
	case Interlingua:
		return "ina"
	case Irish:
		return "gle"
	case Italian:
		return "ita"
	case Japanese:
		return "jpn"
	case Kannada:
		return "kan"
	case Kazakh:
		return "kaz"
	case Khmer:
		return "khm"
	case Korean:
		return "kor"
	case Kurdish:
		return "kur"
	case Kyrgyz:
		return "kir"
	case Latvian:
		return "lav"
	case Lithuanian:
		return "lit"
	case Luxembourgish:
		return "ltz"
	case Macedonian:
		return "mac"
	case Malay:
		return "may"
	case Malayalam:
		return "mal"
	case Manipuri:
		return "mni"
	case Marathi:
		return "mar"
	case Mongolian:
		return "mon"
	case Montenegrin:
		return "mne"
	case Navajo:
		return "nav"
	case Nepali:
		return "nep"
	case NorthernSami:
		return "sme"
	case Norwegian:
		return "nor"
	case Occitan:
		return "oci"
	case Odia:
		return "ori"
	case Persian:
		return "per"
	case Polish:
		return "pol"
	case Portuguese:
		return "por"
	case PortugueseBr:
		return "pob"
	case PortugueseMz:
		return "pom"
	case Pushto:
		return "pus"
	case Romanian:
		return "rum"
	case Russian:
		return "rus"
	case Santali:
		return "sat"
	case Serbian:
		return "scc"
	case Sindhi:
		return "snd"
	case Sinhalese:
		return "sin"
	case Slovak:
		return "slo"
	case Slovenian:
		return "slv"
	case Somali:
		return "som"
	case SorbianLanguages:
		return "wen"
	case SouthAzerbaijani:
		return "azb"
	case Spanish:
		return "spa"
	case SpanishEU:
		return "spn"
	case SpanishLA:
		return "spl"
	case Swahili:
		return "swa"
	case Swedish:
		return "swe"
	case Syriac:
		return "syr"
	case Tagalog:
		return "tgl"
	case Tamil:
		return "tam"
	case Tatar:
		return "tat"
	case Telugu:
		return "tel"
	case Tetum:
		return "tet"
	case Thai:
		return "tha"
	case TokiPona:
		return "tok"
	case Turkish:
		return "tur"
	case Turkmen:
		return "tuk"
	case Ukrainian:
		return "ukr"
	case Urdu:
		return "urd"
	case Uzbek:
		return "uzb"
	case Vietnamese:
		return "vie"
	case Welsch:
		return "wel"
	default:
		return "eng"
	}
}

func (l Lang) String() string {
	switch l {
	case Abkhazian:
		return "abkhazian"
	case Afrikaans:
		return "afrikaans"
	case Albanian:
		return "albanian"
	case Amharic:
		return "amharic"
	case Arabic:
		return "arabic"
	case Aragonese:
		return "aragonese"
	case Armenian:
		return "armenian"
	case Assamese:
		return "assamese"
	case Asturian:
		return "asturian"
	case Azerbaijani:
		return "azerbaijani"
	case Basque:
		return "basque"
	case Belarusian:
		return "belarusian"
	case Bengali:
		return "bengali"
	case Bosnian:
		return "bosnian"
	case Breton:
		return "breton"
	case Bulgarian:
		return "bulgarian"
	case Burmese:
		return "burmese"
	case Catalan:
		return "catalan"
	case ChineseCantonese:
		return "chinese cantonese"
	case ChineseSimplified:
		return "chinese simplified"
	case ChineseTraditional:
		return "chinese traditional"
	case ChineseBilingual:
		return "chinese bilingual"
	case Croatian:
		return "croatian"
	case Czech:
		return "czech"
	case Danish:
		return "danish"
	case Dari:
		return "dari"
	case Dutch:
		return "dutch"
	case English:
		return "english"
	case Esperanto:
		return "esperanto"
	case Estonian:
		return "estonian"
	case Extremaduran:
		return "extremaduran"
	case Finnish:
		return "finnish"
	case French:
		return "french"
	case Gaelic:
		return "gaelic"
	case Galician:
		return "galician"
	case Georgian:
		return "georgian"
	case German:
		return "german"
	case Greek:
		return "greek"
	case Hebrew:
		return "hebrew"
	case Hindi:
		return "hindi"
	case Hungarian:
		return "hungarian"
	case Icelandic:
		return "icelandic"
	case Igbo:
		return "igbo"
	case Indonesian:
		return "indonesian"
	case Interlingua:
		return "interlingua"
	case Irish:
		return "irish"
	case Italian:
		return "italian"
	case Japanese:
		return "japanese"
	case Kannada:
		return "kannada"
	case Kazakh:
		return "kazakh"
	case Khmer:
		return "khmer"
	case Korean:
		return "korean"
	case Kurdish:
		return "kurdish"
	case Kyrgyz:
		return "kyrgyz"
	case Latvian:
		return "latvian"
	case Lithuanian:
		return "lithuanian"
	case Luxembourgish:
		return "luxembourgish"
	case Macedonian:
		return "macedonian"
	case Malay:
		return "malay"
	case Malayalam:
		return "malayalam"
	case Manipuri:
		return "manipuri"
	case Marathi:
		return "marathi"
	case Mongolian:
		return "mongolian"
	case Montenegrin:
		return "montenegrin"
	case Navajo:
		return "navajo"
	case Nepali:
		return "nepali"
	case NorthernSami:
		return "northern sami"
	case Norwegian:
		return "norwegian"
	case Occitan:
		return "occitan"
	case Odia:
		return "odia"
	case Persian:
		return "persian"
	case Polish:
		return "polish"
	case Portuguese:
		return "portuguese"
	case PortugueseBr:
		return "portuguese br"
	case PortugueseMz:
		return "portuguese mz"
	case Pushto:
		return "pushto"
	case Romanian:
		return "romanian"
	case Russian:
		return "russian"
	case Santali:
		return "santali"
	case Serbian:
		return "serbian"
	case Sindhi:
		return "sindhi"
	case Sinhalese:
		return "sinhalese"
	case Slovak:
		return "slovak"
	case Slovenian:
		return "slovenian"
	case Somali:
		return "somali"
	case SorbianLanguages:
		return "sorbian languages"
	case SouthAzerbaijani:
		return "south azerbaijani"
	case Spanish:
		return "spanish"
	case SpanishEU:
		return "spanish eu"
	case SpanishLA:
		return "spanish la"
	case Swahili:
		return "swahili"
	case Swedish:
		return "swedish"
	case Syriac:
		return "syriac"
	case Tagalog:
		return "tagalog"
	case Tamil:
		return "tamil"
	case Tatar:
		return "tatar"
	case Telugu:
		return "telugu"
	case Tetum:
		return "tetum"
	case Thai:
		return "thai"
	case TokiPona:
		return "toki pona"
	case Turkish:
		return "turkish"
	case Turkmen:
		return "turkmen"
	case Ukrainian:
		return "ukrainian"
	case Urdu:
		return "urdu"
	case Uzbek:
		return "uzbek"
	case Vietnamese:
		return "vietnamese"
	case Welsch:
		return "welsch"
	default:
		return "unknown"
	}
}
