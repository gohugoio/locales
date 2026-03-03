package yo_BJ

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type yo_BJ struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	timeSeparator          string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'yo_BJ' locale
func New() locales.Translator {
	return &yo_BJ{
		locale:                 "yo_BJ",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "₦", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsNarrow:           []string{"", "S", "È", "Ɛ", "Ì", "Ɛ̀", "Ò", "A", "Ò", "O", "Ɔ̀", "B", "Ɔ̀"},
		monthsWide:             []string{"", "Oshù Shɛ́rɛ́", "Oshù Èrèlè", "Oshù Ɛrɛ̀nà", "Oshù Ìgbé", "Oshù Ɛ̀bibi", "Oshù Òkúdu", "Oshù Agɛmɔ", "Oshù Ògún", "Oshù Owewe", "Oshù Ɔ̀wàrà", "Oshù Bélú", "Oshù Ɔ̀pɛ̀"},
		daysAbbreviated:        []string{"Àìkú", "Ajé", "Ìsɛ́gun", "Ɔjɔ́rú", "Ɔjɔ́bɔ", "Ɛtì", "Àbámɛ́ta"},
		daysNarrow:             []string{"À", "A", "Ì", "Ɔ", "Ɔ", "Ɛ", "À"},
		daysWide:               []string{"Ɔjɔ́ Àìkú", "Ɔjɔ́ Ajé", "Ɔjɔ́ Ìsɛ́gun", "Ɔjɔ́rú", "Ɔjɔ́bɔ", "Ɔjɔ́ Ɛtì", "Ɔjɔ́ Àbámɛ́ta"},
		timezones:              map[string]string{"ACDT": "Àkókò Ojúmɔmɔ Ààrin Gùngùn Australia", "ACST": "Àkókò Àfɛnukò Ààrin Gùngùn Australia", "ACT": "ACT", "ACWDT": "Àkókò Ojúmɔmɔ Ààrin Gùngùn Ìwɔ̀-Oòrùn Australia", "ACWST": "Àkókò Àfɛnukò Ààrin Gùngùn Ìwɔ̀-Oòrùn Australia", "ADT": "Àkókò Ìyálɛta Àtìláńtíìkì", "ADT Arabia": "Àkókò Ojúmɔmɔ Arabia", "AEDT": "Àkókò Ojúmɔmɔ Ìlà-Oòrùn Australia", "AEST": "Àkókò Àfɛnukò Ìlà-Oòrùn Australia", "AFT": "Àkókò Afghanistan", "AKDT": "Àkókò Ojúmɔ́ Alásíkà", "AKST": "Àkókò Àfɛnukò Alásíkà", "AMST": "Àkókò Oru Amásɔ́nì", "AMST Armenia": "Àkókò Sɔmà Arabia", "AMT": "Àkókò Afɛnukò Amásɔ́nì", "AMT Armenia": "Àkókò Àfɛnukò Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Aago Soma Argentina", "ART": "Aago àsìkò Argentina", "AST": "Àkókò àsikò Àtìláńtíìkì", "AST Arabia": "Àkókò Àfɛnukò Arabia", "AWDT": "Àkókò Ojúmɔmɔ Ìwɔ̀-Oòrùn Australia", "AWST": "Àkókò Àfɛnukò Ìwɔ̀-Oòrùn Australia", "AZST": "Àkókò Sɔmà Azerbaijan", "AZT": "Àkókò Àfɛnukò Azerbaijan", "BDT Bangladesh": "Àkókò Sɔmà Bangladesh", "BNT": "Brunei Darussalam Time", "BOT": "Aago Bolivia", "BRST": "Aago Soma Brasilia", "BRT": "Aago àsìkò Bùràsílíà", "BST Bangladesh": "Àkókò Àfɛnukò Bangladesh", "BT": "Àkókò Bhutan", "CAST": "CAST", "CAT": "Àkókò Àárín Afírikà", "CCT": "Àkókò Àwɔn Erékùsù Cocos", "CDT": "Akókò àárín gbùngbùn ojúmɔmɔ", "CHADT": "Àkókò Ojúmɔmɔ Chatam", "CHAST": "Àkókò Àfɛnukò Chatam", "CHUT": "Àkókò Chuuk", "CKT": "Àkókò Àfɛnukò Àwɔn Erekusu Kuuku", "CKT DST": "Àkókò Ilaji Sɔma Àwɔn Erekusu Kuuku", "CLST": "Àkókò Oru Shílè", "CLT": "Àkókò Àfɛnukò Shílè", "COST": "Aago Soma Colombia", "COT": "Aago àsìkò Kolombia", "CST": "àkókò asiko àárín gbùngbùn", "CST China": "Àkókò Ìfɛnukòsí Sháínà", "CST China DST": "Àkókò Ojúmɔmɔ Sháínà", "CVST": "Àkókò Ɛ̀rún Képú Fáàdì", "CVT": "Àkókò Àfɛnukò Képú Fáàdì", "CXT": "Àkókò Erékùsù Christmas", "ChST": "Àkókò Àfɛnukò Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Àkókò Ojúmɔmɔ Kúbà", "CuST": "Àkókò Àfɛnukò Kúbà", "DAVT": "Àkókò Davis", "DDUT": "Àkókò Dumont-d’Urville", "EASST": "Aago Soma Easter Island", "EAST": "Aago àsìkò Easter Island", "EAT": "Àkókò Ìlà-Oòrùn Afírikà", "ECT": "Aago Ecuador", "EDT": "Àkókò ojúmɔmɔ Ìhà Ìlà Oòrun", "EGDT": "Àkókò ìgbà Ooru Greenland", "EGST": "Àkókò Ìwɔ̀ Ìfɛnukò oorùn Greenland", "EST": "Akókò Àsikò Ìha Ìla Oòrùn", "FEET": "Àkókò Iwájú Ìlà Oòrùn Yúróòpù", "FJT": "Àkókò Àfɛnukò Fiji", "FJT Summer": "Àkókò Sɔma Fiji", "FKST": "Àkókò Ooru Etíkun Fókílándì", "FKT": "Àkókò Àfɛnukò Etíkun Fókílándì", "FNST": "Aago Soma Fernando de Noronha", "FNT": "Aago àsìkò Fenando de Norona", "GALT": "Aago Galapago", "GAMT": "Àkókò Gambia", "GEST": "Àkókò Sɔmà Georgia", "GET": "Àkókò Àfɛnukò Georgia", "GFT": "Àkókò Gúyánà Fáránsè", "GIT": "Àkókò Àwɔn Erekusu Gilibati", "GMT": "Greenwich Mean Time", "GNSST": "GNSST", "GNST": "GNST", "GST": "Àkókò Àfɛnukò Gulf", "GST Guam": "GST Guam", "GYT": "Àkókò Gúyànà", "HADT": "Àkókò Àfɛnukò Hawaii-Aleutian", "HAST": "Àkókò Àfɛnukò Hawaii-Aleutian", "HKST": "Àkókò Sɔmà Hong Kong", "HKT": "Àkókò Ìfɛnukòsí Hong Kong", "HOVST": "Àkókò Sɔmà Hofidi", "HOVT": "Àkókò Ìfɛnukòsí Hofidi", "ICT": "Àkókò Indochina", "IDT": "Àkókò Ojúmɔmɔ Israel", "IOT": "Àkókò Etíkun Índíà", "IRKST": "Àkókò Sɔmà Íkúsíkì", "IRKT": "Àkókò Àfɛnukò Íkósíkì", "IRST": "Àkókò Àfɛnukò Irani", "IRST DST": "Àkókò Ojúmɔmɔ Irani", "IST": "Àkókò Àfɛnukò India", "IST Israel": "Àkókò Àfɛnukò Israel", "JDT": "Àkókò Sɔmà Japan", "JST": "Àkókò Ìfɛnukòsí Japan", "KOST": "Àkókò Kosirai", "KRAST": "Àkókò Sɔmà Krasinoyasiki", "KRAT": "Àkókò Àfɛnukò Krasinoyasiki", "KST": "Àkókò Ìfɛnukòsí Koria", "KST DST": "Àkókò Ojúmɔmɔ Koria", "LHDT": "Àkókò Ojúmɔmɔ Lord Howe", "LHST": "Àkókò Àfɛnukò Lord Howe", "LINT": "Àkókò Àwɔn Erekusu Laini", "MAGST": "Àkókò Sɔmà Magadani", "MAGT": "Àkókò Àfɛnukò Magadani", "MART": "Àkókò Makuesasi", "MAWT": "Àkókò Mawson", "MDT": "Àkókò ojúmɔmɔ Ori-òkè", "MESZ": "Àkókò Àárin Sɔmà Europe", "MEZ": "Àkókò Àárin àsikò Europe", "MHT": "Àkókò Àwɔn Erekusu Masaali", "MMT": "Àkókò Ìlà Myanmar", "MSD": "Àkókò Sɔmà Mosiko", "MST": "Àkókò asiko òkè", "MUST": "Àkókò Ooru Máríshúshì", "MUT": "Àkókò Àfɛnukò Máríshúshì", "MVT": "Àkókò Maldives", "MYT": "Àkókò Malaysia", "NCT": "Àkókò Àfɛnukò Kalidonia Tuntun", "NDT": "Àkókò Ojúmɔmɔ Newfoundland", "NDT New Caledonia": "Àkókò Sɔma Kalidonia Tuntun", "NFDT": "Àkókò Ojúmɔmɔ Erékùsù Norfolk", "NFT": "Àkókò Àfɛnukò Erékùsù Norfolk", "NOVST": "Àkókò Sɔmà Noforibisiki", "NOVT": "Àkókò Àfɛnukò Nofosibiriki", "NPT": "Àkókò Nepali", "NRT": "Àkókò Nauru", "NST": "Àkókò Àfɛnukò Newfoundland", "NUT": "Àkókò Niue", "NZDT": "Àkókò Ojúmɔmɔ New Zealand", "NZST": "Àkókò Àfɛnukò New zealand", "OESZ": "Àkókò Sɔmà Ìha Ìlà Oòrùn Europe", "OEZ": "Àkókò àsikò Ìhà Ìlà Oòrùn Europe", "OMSST": "Àkókò Sɔmà Omisiki", "OMST": "Àkókò Àfɛnukò Omisiki", "PDT": "Àkókò Ìyálɛta Pàsífíìkì", "PDTM": "Àkókò Ojúmɔmɔ Pásífíìkì Mɛ́síkò", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Àkókò Papua New Guinea", "PHOT": "Àkókò Àwɔn Erékùsù Phoenix", "PKT": "Àkókò Àfɛnukò Pakistani", "PKT DST": "Àkókò Sɔmà Pakistani", "PMDT": "Àkókò Ojúmɔmɔ Pierre & Miquelon", "PMST": "Àkókò Àfɛnukò Pierre & Miquelon", "PONT": "Àkókò Ponape", "PST": "Àkókò àsikò Pàsífíìkì", "PST Philippine": "Àkókò Àfɛnukò Filipininni", "PST Philippine DST": "Àkókò Sɔmà Filipininni", "PST Pitcairn": "Àkókò Pitcairn", "PSTM": "Àkókò Àfɛnukò Pásífíìkì Mɛ́síkò", "PWT": "Àkókò Palau", "PYST": "Àkókò Ooru Párágúwè", "PYT": "Àkókò Àfɛnukò Párágúwè", "PYT Korea": "Àkókò Pyongyangi", "RET": "Àkókò Rɛ́yúníɔ́nì", "ROTT": "Àkókò Rothera", "SAKST": "Àkókò Sɔmà Sakhalin", "SAKT": "Àkókò Àfɛnukò Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "South Africa Standard Time", "SBT": "Àkókò Àwɔn Erekusu Solomon", "SCT": "Àkókò Sèshɛ́ɛ̀lì", "SGT": "Àkókò Àfɛnukò Singapore", "SLST": "SLST", "SRT": "Àkókò Súrínámù", "SST Samoa": "Àkókò Àfɛnukò Samoa", "SST Samoa Apia": "Àkókò Àfɛnukò Apia", "SST Samoa Apia DST": "Àkókò Ojúmɔmɔ Apia", "SST Samoa DST": "Àkókò Ojúmɔmɔ Samoa", "SYOT": "Àkókò Syowa", "TAAF": "Àkókò Gúsù Fáransé àti Àntátíìkì", "TAHT": "Àkókò Tahiti", "TJT": "Àkókò Tajikisitaani", "TKT": "Àkókò Tokelau", "TLT": "Àkókò Ìlà oorùn Timor", "TMST": "Àkókò Sɔmà Turkmenistani", "TMT": "Àkókò Àfɛnukò Turkimenistani", "TOST": "Àkókò Sɔmà Tonga", "TOT": "Àkókò Àfɛnukò Tonga", "TVT": "Àkókò Tufalu", "TWT": "Àkókò Ìfɛnukòsí Taipei", "TWT DST": "Àkókò Ojúmɔmɔ Taipei", "ULAST": "Àkókò Sɔmà Ulaanbaatar", "ULAT": "Àkókò Ìfɛnukòsí Ulaanbaatar", "UYST": "Aago Soma Uruguay", "UYT": "Àkókò Àfɛnukò Úrúgúwè", "UZT": "Àkókò Àfɛnukò Usibekistani", "UZT DST": "Àkókò Sɔmà Usibekistani", "VET": "Aago Venezuela", "VLAST": "Àkókò Sɔmà Filadifositoki", "VLAT": "Àkókò Àfɛnukò Filadifositoki", "VOLST": "Àkókò Sɔmà Foligogiradi", "VOLT": "Àkókò Àfɛnukò Foligogiradi", "VOST": "Àkókò Vostok", "VUT": "Àkókò Àfɛnukò Fanuatu", "VUT DST": "Àkókò Sɔmà Fanuatu", "WAKT": "Àkókò Erékùsù Wake", "WARST": "Àkókò Oru Iwɔ́-oòrùn Ajɛ́ntínà", "WART": "Àkókò Iwɔ́-oòrùn Àfɛnukò Ajɛ́ntínà", "WAST": "Àkókò Ìwɔ̀-Oòrùn Afírikà", "WAT": "Àkókò Ìwɔ̀-Oòrùn Afírikà", "WESZ": "Àkókò Sɔmà Ìhà Ìwɔ Oòrùn Europe", "WEZ": "Àkókò àsikò Ìwɔ Oòrùn Europe", "WFT": "Àkókò Wallis & Futuina", "WGST": "Àkókò Àfɛnukò Ìgba Oòru Greenland", "WGT": "Àkókò Àfɛnukò Ìwɔ̀ Oòrùn Greenland", "WIB": "Àkókò Ìwɔ̀ oorùn Indonesia", "WIT": "Àkókò Ìlà oorùn Indonesia", "WITA": "Àkókò Ààrin Gbùngbùn Indonesia", "YAKST": "Àkókò Sɔmà Yatutsk", "YAKT": "Àkókò Àfɛnukò Yatutsk", "YEKST": "Àkókò Sɔmà Yekaterinburg", "YEKT": "Àkókò Àfɛnukò Yekaterinburg", "YST": "Àkókò Yúkɔ́nì", "МСК": "Àkókò Àfɛnukò Mosiko", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Àkókò Ìwɔ̀-Oòrùn Kasasitáànì", "شىعىش قازاق ەلى": "Àkókò Ìlà-Oòrùn Kasasitáànì", "قازاق ەلى": "Aago Kasasitáànì", "قىرعىزستان": "Àkókò Kirigisitaani", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Àkókò Ooru Ásɔ́sì"},
	}
}

// Locale returns the current translators string locale
func (yo *yo_BJ) Locale() string {
	return yo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsCardinal() []locales.PluralRule {
	return yo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsOrdinal() []locales.PluralRule {
	return yo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsRange() []locales.PluralRule {
	return yo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yo_BJ'
func (yo *yo_BJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yo *yo_BJ) MonthAbbreviated(month time.Month) string {
	if len(yo.monthsAbbreviated) == 0 {
		return ""
	}
	return yo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yo *yo_BJ) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yo *yo_BJ) MonthNarrow(month time.Month) string {
	if len(yo.monthsNarrow) == 0 {
		return ""
	}
	return yo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yo *yo_BJ) MonthsNarrow() []string {
	return yo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yo *yo_BJ) MonthWide(month time.Month) string {
	if len(yo.monthsWide) == 0 {
		return ""
	}
	return yo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yo *yo_BJ) MonthsWide() []string {
	return yo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(yo.daysAbbreviated) == 0 {
		return ""
	}
	return yo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yo *yo_BJ) WeekdaysAbbreviated() []string {
	return yo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayNarrow(weekday time.Weekday) string {
	if len(yo.daysNarrow) == 0 {
		return ""
	}
	return yo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yo *yo_BJ) WeekdaysNarrow() []string {
	return yo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayShort(weekday time.Weekday) string {
	if len(yo.daysShort) == 0 {
		return ""
	}
	return yo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yo *yo_BJ) WeekdaysShort() []string {
	return yo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayWide(weekday time.Weekday) string {
	if len(yo.daysWide) == 0 {
		return ""
	}
	return yo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yo *yo_BJ) WeekdaysWide() []string {
	return yo.daysWide
}

// Decimal returns the decimal point of number
func (yo *yo_BJ) Decimal() string {
	return yo.decimal
}

// Group returns the group of number
func (yo *yo_BJ) Group() string {
	return yo.group
}

// Group returns the minus sign of number
func (yo *yo_BJ) Minus() string {
	return yo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yo_BJ' and handles both Whole and Real numbers based on 'v'
func (yo *yo_BJ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yo_BJ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yo *yo_BJ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yo.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yo.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, yo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yo_BJ'
// in accounting notation.
func (yo *yo_BJ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yo.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, yo.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, yo.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, yo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := yo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
