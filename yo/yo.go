package yo

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type yo struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
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
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'yo' locale
func New() locales.Translator {
	return &yo{
		locale:                 "yo",
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
		monthsNarrow:           []string{"", "S", "È", "Ẹ", "Ì", "Ẹ̀", "Ò", "A", "Ò", "O", "Ọ̀", "B", "Ọ̀"},
		monthsWide:             []string{"", "Oṣù Ṣẹ́rẹ́", "Oṣù Èrèlè", "Oṣù Ẹrẹ̀nà", "Oṣù Ìgbé", "Oṣù Ẹ̀bibi", "Oṣù Òkúdu", "Oṣù Agẹmọ", "Oṣù Ògún", "Oṣù Owewe", "Oṣù Ọ̀wàrà", "Oṣù Bélú", "Oṣù Ọ̀pẹ̀"},
		daysAbbreviated:        []string{"Àìkú", "Ajé", "Ìsẹ́gun", "Ọjọ́rú", "Ọjọ́bọ", "Ẹtì", "Àbámẹ́ta"},
		daysNarrow:             []string{"À", "A", "Ì", "Ọ", "Ọ", "Ẹ", "À"},
		daysWide:               []string{"Ọjọ́ Àìkú", "Ọjọ́ Ajé", "Ọjọ́ Ìsẹ́gun", "Ọjọ́rú", "Ọjọ́bọ", "Ọjọ́ Ẹtì", "Ọjọ́ Àbámẹ́ta"},
		periodsAbbreviated:     []string{"Àárọ̀", "Ọ̀sán"},
		erasAbbreviated:        []string{"AD", "CE"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "Àkókò Ojúmọmọ Ààrin Gùngùn Australia", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Àkókò Ojúmọmọ Ààrin Gùngùn Ìwọ̀-Oòrùn Australia", "ACWST": "Àkókò Àfẹnukò Ààrin Gùngùn Ìwọ̀-Oòrùn Australia", "ADT": "Àkókò Ìyálẹta Àtìláńtíìkì", "ADT Arabia": "Àkókò Ojúmọmọ Arabia", "AEDT": "Àkókò Ojúmọmọ Ìlà-Oòrùn Australia", "AEST": "Àkókò Àfẹnukò Ìlà-Oòrùn Australia", "AFT": "Àkókò Afghanistan", "AKDT": "Àkókò Ojúmọ́ Alásíkà", "AKST": "Àkókò Àfẹnukò Alásíkà", "AMST": "Àkókò Oru Amásọ́nì", "AMST Armenia": "Àkókò Sọmà Arabia", "AMT": "Àkókò Afẹnukò Amásọ́nì", "AMT Armenia": "Àkókò Àfẹnukò Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Aago Soma Argentina", "ART": "Aago àsìkò Argentina", "AST": "Àkókò àsikò Àtìláńtíìkì", "AST Arabia": "Àkókò Àfẹnukò Arabia", "AWDT": "Àkókò Ojúmọmọ Ìwọ̀-Oòrùn Australia", "AWST": "Àkókò Àfẹnukò Ìwọ̀-Oòrùn Australia", "AZST": "Àkókò Sọmà Azerbaijan", "AZT": "Àkókò Àfẹnukò Azerbaijan", "BDT Bangladesh": "Àkókò Sọmà Bangladesh", "BNT": "Brunei Darussalam Time", "BOT": "Aago Bolivia", "BRST": "Aago Soma Brasilia", "BRT": "Aago àsìkò Bùràsílíà", "BST Bangladesh": "Àkókò Àfẹnukò Bangladesh", "BT": "Àkókò Bhutan", "CAST": "CAST", "CAT": "Àkókò Àárín Afírikà", "CCT": "Àkókò Àwọn Erékùsù Cocos", "CDT": "Akókò àárín gbùngbùn ojúmọmọ", "CHADT": "Àkókò Ojúmọmọ Chatam", "CHAST": "Àkókò Àfẹnukò Chatam", "CHUT": "Àkókò Chuuk", "CKT": "Àkókò Àfẹnukò Àwọn Erekusu Kuuku", "CKT DST": "Àkókò Ilaji Sọma Àwọn Erekusu Kuuku", "CLST": "Àkókò Oru Ṣílè", "CLT": "Àkókò Àfẹnukò Ṣílè", "COST": "Aago Soma Colombia", "COT": "Aago àsìkò Kolombia", "CST": "àkókò asiko àárín gbùngbùn", "CST China": "Àkókò Ìfẹnukòsí Ṣáínà", "CST China DST": "Àkókò Ojúmọmọ Ṣáínà", "CVST": "Àkókò Ẹ̀rún Képú Fáàdì", "CVT": "Àkókò Àfẹnukò Képú Fáàdì", "CXT": "Àkókò Erékùsù Christmas", "ChST": "Àkókò Àfẹnukò Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Àkókò Ojúmọmọ Kúbà", "CuST": "Àkókò Àfẹnukò Kúbà", "DAVT": "Àkókò Davis", "DDUT": "Àkókò Dumont-d’Urville", "EASST": "Aago Soma Easter Island", "EAST": "Aago àsìkò Easter Island", "EAT": "Àkókò Ìlà-Oòrùn Afírikà", "ECT": "Aago Ecuador", "EDT": "Àkókò ojúmọmọ Ìhà Ìlà Oòrun", "EGDT": "Àkókò ìgbà Ooru Greenland", "EGST": "Àkókò Ìwọ̀ Ìfẹnukò oorùn Greenland", "EST": "Akókò Àsikò Ìha Ìla Oòrùn", "FEET": "Àkókò Iwájú Ìlà Oòrùn Yúróòpù", "FJT": "Àkókò Àfẹnukò Fiji", "FJT Summer": "Àkókò Sọma Fiji", "FKST": "Àkókò Ooru Etíkun Fókílándì", "FKT": "Àkókò Àfẹnukò Etíkun Fókílándì", "FNST": "Aago Soma Fernando de Noronha", "FNT": "Aago àsìkò Fenando de Norona", "GALT": "Aago Galapago", "GAMT": "Àkókò Gambia", "GEST": "Àkókò Sọmà Georgia", "GET": "Àkókò Àfẹnukò Georgia", "GFT": "Àkókò Gúyánà Fáránsè", "GIT": "Àkókò Àwọn Erekusu Gilibati", "GMT": "Greenwich Mean Time", "GNSST": "GNSST", "GNST": "GNST", "GST": "Àkókò Àfẹnukò Gulf", "GST Guam": "GST Guam", "GYT": "Àkókò Gúyànà", "HADT": "Àkókò Àfẹnukò Hawaii-Aleutian", "HAST": "Àkókò Àfẹnukò Hawaii-Aleutian", "HKST": "Àkókò Sọmà Hong Kong", "HKT": "Àkókò Ìfẹnukòsí Hong Kong", "HOVST": "Àkókò Sọmà Hofidi", "HOVT": "Àkókò Ìfẹnukòsí Hofidi", "ICT": "Àkókò Indochina", "IDT": "Àkókò Ojúmọmọ Israel", "IOT": "Àkókò Etíkun Índíà", "IRKST": "Àkókò Sọmà Íkúsíkì", "IRKT": "Àkókò Àfẹnukò Íkósíkì", "IRST": "Àkókò Àfẹnukò Irani", "IRST DST": "Àkókò Ojúmọmọ Irani", "IST": "Àkókò Àfẹnukò India", "IST Israel": "Àkókò Àfẹnukò Israel", "JDT": "Àkókò Sọmà Japan", "JST": "Àkókò Ìfẹnukòsí Japan", "KOST": "Àkókò Kosirai", "KRAST": "Àkókò Sọmà Krasinoyasiki", "KRAT": "Àkókò Àfẹnukò Krasinoyasiki", "KST": "Àkókò Ìfẹnukòsí Koria", "KST DST": "Àkókò Ojúmọmọ Koria", "LHDT": "Àkókò Ojúmọmọ Lord Howe", "LHST": "Àkókò Àfẹnukò Lord Howe", "LINT": "Àkókò Àwọn Erekusu Laini", "MAGST": "Àkókò Sọmà Magadani", "MAGT": "Àkókò Àfẹnukò Magadani", "MART": "Àkókò Makuesasi", "MAWT": "Àkókò Mawson", "MDT": "MDT", "MESZ": "Àkókò Àárin Sọmà Europe", "MEZ": "Àkókò Àárin àsikò Europe", "MHT": "Àkókò Àwọn Erekusu Masaali", "MMT": "Àkókò Ìlà Myanmar", "MSD": "Àkókò Sọmà Mosiko", "MST": "MST", "MUST": "Àkókò Ooru Máríṣúṣì", "MUT": "Àkókò Àfẹnukò Máríṣúṣì", "MVT": "Àkókò Maldives", "MYT": "Àkókò Malaysia", "NCT": "Àkókò Àfẹnukò Kalidonia Tuntun", "NDT": "Àkókò Ojúmọmọ Newfoundland", "NDT New Caledonia": "Àkókò Sọma Kalidonia Tuntun", "NFDT": "Àkókò Ojúmọmọ Erékùsù Norfolk", "NFT": "Àkókò Àfẹnukò Erékùsù Norfolk", "NOVST": "Àkókò Sọmà Noforibisiki", "NOVT": "Àkókò Àfẹnukò Nofosibiriki", "NPT": "Àkókò Nepali", "NRT": "Àkókò Nauru", "NST": "Àkókò Àfẹnukò Newfoundland", "NUT": "Àkókò Niue", "NZDT": "Àkókò Ojúmọmọ New Zealand", "NZST": "Àkókò Àfẹnukò New zealand", "OESZ": "Àkókò Sọmà Ìha Ìlà Oòrùn Europe", "OEZ": "Àkókò àsikò Ìhà Ìlà Oòrùn Europe", "OMSST": "Àkókò Sọmà Omisiki", "OMST": "Àkókò Àfẹnukò Omisiki", "PDT": "Àkókò Ìyálẹta Pàsífíìkì", "PDTM": "Àkókò Ojúmọmọ Pásífíìkì Mẹ́síkò", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Àkókò Papua New Guinea", "PHOT": "Àkókò Àwọn Erékùsù Phoenix", "PKT": "Àkókò Àfẹnukò Pakistani", "PKT DST": "Àkókò Sọmà Pakistani", "PMDT": "Àkókò Ojúmọmọ Pierre & Miquelon", "PMST": "Àkókò Àfẹnukò Pierre & Miquelon", "PONT": "Àkókò Ponape", "PST": "Àkókò àsikò Pàsífíìkì", "PST Philippine": "Àkókò Àfẹnukò Filipininni", "PST Philippine DST": "Àkókò Sọmà Filipininni", "PST Pitcairn": "Àkókò Pitcairn", "PSTM": "Àkókò Àfẹnukò Pásífíìkì Mẹ́síkò", "PWT": "Àkókò Palau", "PYST": "Àkókò Ooru Párágúwè", "PYT": "Àkókò Àfẹnukò Párágúwè", "PYT Korea": "Àkókò Pyongyangi", "RET": "Àkókò Rẹ́yúníọ́nì", "ROTT": "Àkókò Rothera", "SAKST": "Àkókò Sọmà Sakhalin", "SAKT": "Àkókò Àfẹnukò Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "South Africa Standard Time", "SBT": "Àkókò Àwọn Erekusu Solomon", "SCT": "Àkókò Sèṣẹ́ẹ̀lì", "SGT": "Àkókò Àfẹnukò Singapore", "SLST": "SLST", "SRT": "Àkókò Súrínámù", "SST Samoa": "Àkókò Àfẹnukò Samoa", "SST Samoa Apia": "Àkókò Àfẹnukò Apia", "SST Samoa Apia DST": "Àkókò Ojúmọmọ Apia", "SST Samoa DST": "Àkókò Ojúmọmọ Samoa", "SYOT": "Àkókò Syowa", "TAAF": "Àkókò Gúsù Fáransé àti Àntátíìkì", "TAHT": "Àkókò Tahiti", "TJT": "Àkókò Tajikisitaani", "TKT": "Àkókò Tokelau", "TLT": "Àkókò Ìlà oorùn Timor", "TMST": "Àkókò Sọmà Turkmenistani", "TMT": "Àkókò Àfẹnukò Turkimenistani", "TOST": "Àkókò Sọmà Tonga", "TOT": "Àkókò Àfẹnukò Tonga", "TVT": "Àkókò Tufalu", "TWT": "Àkókò Ìfẹnukòsí Taipei", "TWT DST": "Àkókò Ojúmọmọ Taipei", "ULAST": "Àkókò Sọmà Ulaanbaatar", "ULAT": "Àkókò Ìfẹnukòsí Ulaanbaatar", "UYST": "Aago Soma Uruguay", "UYT": "Àkókò Àfẹnukò Úrúgúwè", "UZT": "Àkókò Àfẹnukò Usibekistani", "UZT DST": "Àkókò Sọmà Usibekistani", "VET": "Aago Venezuela", "VLAST": "Àkókò Sọmà Filadifositoki", "VLAT": "Àkókò Àfẹnukò Filadifositoki", "VOLST": "Àkókò Sọmà Foligogiradi", "VOLT": "Àkókò Àfẹnukò Foligogiradi", "VOST": "Àkókò Vostok", "VUT": "Àkókò Àfẹnukò Fanuatu", "VUT DST": "Àkókò Sọmà Fanuatu", "WAKT": "Àkókò Erékùsù Wake", "WARST": "Àkókò Oru Iwọ́-oòrùn Ajẹ́ntínà", "WART": "Àkókò Iwọ́-oòrùn Àfẹnukò Ajẹ́ntínà", "WAST": "Àkókò Ìwọ̀-Oòrùn Afírikà", "WAT": "Àkókò Ìwọ̀-Oòrùn Afírikà", "WESZ": "Àkókò Sọmà Ìhà Ìwọ Oòrùn Europe", "WEZ": "Àkókò àsikò Ìwọ Oòrùn Europe", "WFT": "Àkókò Wallis & Futuina", "WGST": "Àkókò Àfẹnukò Ìgba Oòru Greenland", "WGT": "Àkókò Àfẹnukò Ìwọ̀ Oòrùn Greenland", "WIB": "Àkókò Ìwọ̀ oorùn Indonesia", "WIT": "Àkókò Ìlà oorùn Indonesia", "WITA": "Àkókò Ààrin Gbùngbùn Indonesia", "YAKST": "Àkókò Sọmà Yatutsk", "YAKT": "Àkókò Àfẹnukò Yatutsk", "YEKST": "Àkókò Sọmà Yekaterinburg", "YEKT": "Àkókò Àfẹnukò Yekaterinburg", "YST": "Àkókò Yúkọ́nì", "МСК": "Àkókò Àfẹnukò Mosiko", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Àkókò Ìwọ̀-Oòrùn Kasasitáànì", "شىعىش قازاق ەلى": "Àkókò Ìlà-Oòrùn Kasasitáànì", "قازاق ەلى": "Aago Kasasitáànì", "قىرعىزستان": "Àkókò Kirigisitaani", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Àkókò Ooru Pérù"},
	}
}

// Locale returns the current translators string locale
func (yo *yo) Locale() string {
	return yo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yo'
func (yo *yo) PluralsCardinal() []locales.PluralRule {
	return yo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yo'
func (yo *yo) PluralsOrdinal() []locales.PluralRule {
	return yo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yo'
func (yo *yo) PluralsRange() []locales.PluralRule {
	return yo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yo'
func (yo *yo) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yo'
func (yo *yo) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yo'
func (yo *yo) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yo *yo) MonthAbbreviated(month time.Month) string {
	return yo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yo *yo) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yo *yo) MonthNarrow(month time.Month) string {
	return yo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yo *yo) MonthsNarrow() []string {
	return yo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yo *yo) MonthWide(month time.Month) string {
	return yo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yo *yo) MonthsWide() []string {
	return yo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yo *yo) WeekdayAbbreviated(weekday time.Weekday) string {
	return yo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yo *yo) WeekdaysAbbreviated() []string {
	return yo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yo *yo) WeekdayNarrow(weekday time.Weekday) string {
	return yo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yo *yo) WeekdaysNarrow() []string {
	return yo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yo *yo) WeekdayShort(weekday time.Weekday) string {
	return yo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yo *yo) WeekdaysShort() []string {
	return yo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yo *yo) WeekdayWide(weekday time.Weekday) string {
	return yo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yo *yo) WeekdaysWide() []string {
	return yo.daysWide
}

// Decimal returns the decimal point of number
func (yo *yo) Decimal() string {
	return yo.decimal
}

// Group returns the group of number
func (yo *yo) Group() string {
	return yo.group
}

// Group returns the minus sign of number
func (yo *yo) Minus() string {
	return yo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yo' and handles both Whole and Real numbers based on 'v'
func (yo *yo) FmtNumber(num float64, v uint64) string {
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

// FmtPercent returns 'num' with digits/precision of 'v' for 'yo' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yo *yo) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yo'
func (yo *yo) FmtCurrency(num float64, v uint64, currency currency.Type) string {
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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yo'
// in accounting notation.
func (yo *yo) FmtAccounting(num float64, v uint64, currency currency.Type) string {
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

// FmtDateShort returns the short date representation of 't' for 'yo'
func (yo *yo) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'yo'
func (yo *yo) FmtDateMedium(t time.Time) string {
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

// FmtDateLong returns the long date representation of 't' for 'yo'
func (yo *yo) FmtDateLong(t time.Time) string {
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

// FmtDateFull returns the full date representation of 't' for 'yo'
func (yo *yo) FmtDateFull(t time.Time) string {
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

// FmtTimeShort returns the short time representation of 't' for 'yo'
func (yo *yo) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'yo'
func (yo *yo) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'yo'
func (yo *yo) FmtTimeLong(t time.Time) string {
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

// FmtTimeFull returns the full time representation of 't' for 'yo'
func (yo *yo) FmtTimeFull(t time.Time) string {
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
