package om_ET

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type om_ET struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	timeSeparator      string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'om_ET' locale
func New() locales.Translator {
	return &om_ET{
		locale:             "om_ET",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "Br", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "Ama", "Gur", "Bitootessa", "Elb", "Cam", "Wax", "Ado", "Hag", "Ful", "Onk", "Sadaasa", "Mud"},
		monthsNarrow:       []string{"", "A", "G", "B", "E", "C", "W", "A", "H", "F", "O", "S", "M"},
		monthsWide:         []string{"", "Amajjii", "Guraandhala", "Bitootessa", "Eebila", "Caamsaa", "Waxabajjii", "Adoolessa", "Hagayya", "Fulbaana", "Onkoloolessa", "Sadaasa", "Mudde"},
		daysAbbreviated:    []string{"Dil", "Wix", "Kib", "Rob", "Kam", "Jim", "San"},
		daysNarrow:         []string{"D", "W", "K", "R", "K", "J", "S"},
		daysWide:           []string{"Dilbata", "Wiixata", "Kibxata", "Roobii", "Kamisa", "Jimaata", "Sanbata"},
		periodsAbbreviated: []string{"WD", "WB"},
		timezones:          map[string]string{"ACDT": "Sa’aatii Guyyaa Awustiraaliyaa Gidduugaleessaa", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Sa’aatii Guyyaa Dhiha Awustiraaliyaa Gidduugaleessaa", "ACWST": "Sa’aatii Istaandaardii Dhiha Awustiraaliyaa Gidduugaleessaa", "ADT": "Sa’aatii Guyyaa Atilaantiik", "ADT Arabia": "Sa’aatii Guyyaa Arabaa", "AEDT": "Sa’aatii Guyyaa Awustiraaliyaa Bahaa", "AEST": "Sa’aatii Istaandaardii Awustiraaliyaa Bahaa", "AFT": "Sa’aatii Afgaanistaan", "AKDT": "Sa’aatii Guyyaa Alaaskaa", "AKST": "Sa’aatii Istaandaardii Alaaskaa", "AMST": "Sa’aatii Bonaa Amazoon", "AMST Armenia": "Sa’aatii Bonaa Armaaniyaa", "AMT": "Sa’aatii Istaandaardii Amazoon", "AMT Armenia": "Sa’aatii Istaandaardii Armaaniyaa", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Sa’aatii Bonaa Arjentiinaa", "ART": "Sa’aatii Istaandaardii Arjentiinaa", "AST": "Sa’aatii Istaandaardii Atilaantiik", "AST Arabia": "Sa’aatii Istaandaardii Arabaa", "AWDT": "Sa’aatii Guuyyaa Awustiraaliyaa Dhihaa", "AWST": "Sa’aatii Sa’aatii Istaandaardii Awustiraaliyaa DhihaaDhiha Awustiraaliyaa", "AZST": "Sa’aatii Bonaa Azerbaajiyaan", "AZT": "Sa’aatii Istaandaardii Azerbaajiyaan", "BDT Bangladesh": "Sa’aatii Bonaa Baangilaadish", "BNT": "Sa’aatii Bruunee Darusalaam", "BOT": "Sa’aatii Boliiviyaa", "BRST": "Sa’aatii Bonaa Biraaziliyaa", "BRT": "Sa’aatii Istaandaardii Biraaziliyaa", "BST Bangladesh": "Sa’aatii Istaandaardii Baangilaadish", "BT": "Sa’aatii Bihutaan", "CAST": "CAST", "CAT": "Sa’aatii Afrikaa Gidduugaleessaa", "CCT": "Sa’aatii Odoloota Kokos", "CDT": "Sa’aatii Guyyaa Gidduugaleessaa", "CHADT": "Sa’aatii Guyyaa Chatham", "CHAST": "Sa’aatii Istaandaardii Chatham", "CHUT": "Sa’aatii Chuuk", "CKT": "Sa’aatii Istaandaardii Odoloota Kuuk", "CKT DST": "Sa’aatii Bona Walakkaa Odoloota Kuuk", "CLST": "Sa’aatii Bonaa Chiilii", "CLT": "Sa’aatii Istaandaardii Chiilii", "COST": "Sa’aatii Bonaa Kolombiyaa", "COT": "Sa’aatii Istaandaardii Kolombiyaa", "CST": "Sa’aatii Istaandaardii Gidduugaleessaa", "CST China": "Sa’aatii Istaandaardii Chaayinaa", "CST China DST": "Sa’aatii Guyyaa Chaayinaa", "CVST": "Sa’aatii Bonaa Keep Veerdee", "CVT": "Sa’aatii Istaandaardii Keep Veerdee", "CXT": "Sa’aatii Odola Kirismaas", "ChST": "Sa’aatii Istaandaardii Kamoroo", "ChST NMI": "Sa’aatii Odoloota Maariyaanaa Kaabaa", "CuDT": "Sa’aatii Guyyaa Kuubaa", "CuST": "Sa’aatii Istaandaardii Kuubaa", "DAVT": "Sa’aatii Daaviis", "DDUT": "Sa’aatii Dumont-d’Urville", "EASST": "Sa’aatii Bonaa Odola Bahaa", "EAST": "Sa’aatii Istaandaardii Odola Bahaa", "EAT": "Sa’aatii Baha Afrikaa", "ECT": "Sa’aatii Ikkuwaadoor", "EDT": "Sa’aatii Guyyaa Bahaa", "EGDT": "Sa’aatii Bonaa Giriinlaand Bahaa", "EGST": "Sa’aatii Istaandaardii Giriinlaand Bahaa", "EST": "Sa’aatii Istaandaardii Bahaa", "FEET": "Sa’aatii Awurooppaa Bahaa Dabalataa", "FJT": "Sa’aatii Istaandaardii Fiijii", "FJT Summer": "Sa’aatii Bonaa Fiijii", "FKST": "Sa’aatii Bonaa Odoloota Faalklaand", "FKT": "Sa’aatii Istaandaardii Odoloota Faalklaand", "FNST": "Sa’aatii Bonaa Fernando de Noronha", "FNT": "Sa’aatii Istaandaardii Fernando de Noronha", "GALT": "Sa’aatii Galaapagoos", "GAMT": "Sa’aatii Gaambiyeer", "GEST": "Sa’aatii Bonaa Joorjiyaa", "GET": "Sa’aatii Istaandaardii Joorjiyaa", "GFT": "Sa’aatii Fireench Guyinaa", "GIT": "Sa’aatii Odoloota Giilbert", "GMT": "Sa’aatii Giriinwiich Gidduugaleessaa", "GNSST": "GNSST", "GNST": "GNST", "GST": "Sa’aatii Istaandaardii Gaalfii", "GST Guam": "Sa’aatii Istaandaardii Guwaam", "GYT": "Sa’aatii Guyaanaa", "HADT": "Sa’aatii Istaandaardii Haawayi-Alewutiyan", "HAST": "Sa’aatii Istaandaardii Haawayi-Alewutiyan", "HKST": "Sa’aatii Bonaa Hoong Koong", "HKT": "Sa’aatii Istaandaardii Hoong Koong", "HOVST": "Sa’aatii Bonaa Hoovd", "HOVT": "Sa’aatii Istaandaardii Hoovd", "ICT": "Sa’aatii IndooChaayinaa", "IDT": "Sa’aatii Guyyaa Israa’eel", "IOT": "Sa’aatii Galaana Hindii", "IRKST": "Sa’aatii Bonaa Irkutsk", "IRKT": "Sa’aatii Istaandaardii Irkutsk", "IRST": "Sa’aatii Istaandaardii Iraan", "IRST DST": "Sa’aatii Guyyaa Iraan", "IST": "Sa’aatii Istaandaardii Hindii", "IST Israel": "Sa’aatii Istaandaardii Israa’eel", "JDT": "Sa’aatii Guyyaa Jaappaan", "JST": "Sa’aatii Istaandaardii Jaappaan", "KOST": "Sa’aatii Koosreyaa", "KRAST": "Sa’aatii Bonaa Krasnoyarsk", "KRAT": "Sa’aatii Istaandaardii Krasnoyarsk", "KST": "Sa’aatii Istaandaardii Kooriyaa", "KST DST": "Sa’aatii Guyyaa Kooriyaa", "LHDT": "Sa’aatii Guyyaa Lord Howe", "LHST": "Sa’aatii Istaandaardii Lord Howe", "LINT": "Sa’aatii Odoloota Line", "MAGST": "Sa’aatii Bonaa Magadan", "MAGT": "Sa’aatii Istaandaardii Magadan", "MART": "Sa’aatii Marquesas", "MAWT": "Sa’aatii Mawson", "MDT": "MDT", "MESZ": "Sa’aatii Bonaa Awurooppaa Gidduugaleessaa", "MEZ": "Sa’aatii Istaandaardii Awurooppaa Gidduugaleessaa", "MHT": "Sa’aatii Odoloota Maarshaal", "MMT": "Sa’aatii Maayinaamaar", "MSD": "Sa’aatii Bonaa Mooskoo", "MST": "MST", "MUST": "Sa’aatii Bonaa Moorishiyees", "MUT": "Sa’aatii Istaandaardii Moorishiyees", "MVT": "Sa’aatii Maaldiivs", "MYT": "Sa’aatii Maaleeshiyaa", "NCT": "Sa’aatii Istaandaardii Kaaledooniyaa Haaraa", "NDT": "Sa’aatii Guyyaa Newufaawondlaand", "NDT New Caledonia": "Sa’aatii Bonaa Kaaledooniyaa Haaraa", "NFDT": "Sa’aatii Guyyaa Norfolk Island", "NFT": "Sa’aatii Istaandaardii Norfolk Island", "NOVST": "Sa’aatii Bonaa Novosibirisk", "NOVT": "Sa’aatii Istaandaardii Novosibirisk", "NPT": "Sa’aatii Neeppaal", "NRT": "Sa’aatii Naawuruu", "NST": "Sa’aatii Istaandaardii Newufaawondlaand", "NUT": "Sa’aatii Niue", "NZDT": "Sa’aatii Guyyaa New Zealand", "NZST": "Sa’aatii Istaandaardii New Zealand", "OESZ": "Sa’aatii Bonaa Awurooppaa Bahaa", "OEZ": "Sa’aatii Istaandaardii Awurooppaa Bahaa", "OMSST": "Sa’aatii Bonaa Omsk", "OMST": "Sa’aatii Istaandaardii Omsk", "PDT": "Sa’aatii Guyyaa Paasfiik", "PDTM": "Sa’aatii Guyyaa Paasfiik Meksiikaan", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Sa’aatii Paapuwaa Giinii Haaraa", "PHOT": "Sa’aatii Odoloota Fooneeks", "PKT": "Sa’aatii Istaandaardii Paakistaan", "PKT DST": "Sa’aatii Bonaa Paakistaan", "PMDT": "Sa’aatii Guyyaa Ql. Piyeeree fi Mikuyelo", "PMST": "Sa’aatii Istaandaardii Ql. Piyeeree fi Mikuyelo", "PONT": "Sa’aatii Ponape", "PST": "Sa’aatii Istaandaardii Paasfiik", "PST Philippine": "Sa’aatii Istaandaardii Filippiins", "PST Philippine DST": "Sa’aatii Bonaa Filippiins", "PST Pitcairn": "Sa’aatii Pitcairn", "PSTM": "Sa’aatii Istaandaardii Paasfiik Meksiikaan", "PWT": "Sa’aatii Palawu", "PYST": "Sa’aatii Bonaa Paaraaguwaayi", "PYT": "Sa’aatii Istaandaardii Paaraaguwaayi", "PYT Korea": "Sa’aatii Piyoongyaang", "RET": "Sa’aatii Riiyuuniyeen", "ROTT": "Sa’aatii Rothera", "SAKST": "Sa’aatii Bonaa Sakhalin", "SAKT": "Sa’aatii Istaandaardii Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Sa’aatii Istaandaardii Afrikaa Kibbaa", "SBT": "Sa’aatii Odoloota Solomoon", "SCT": "Sa’aatii Siisheels", "SGT": "Sa’aatii Istaandaardii Singaapoor", "SLST": "Sa’aatii Laankaa", "SRT": "Sa’aatii Surinaame", "SST Samoa": "Sa’aatii Istaandaardii Saamowaa", "SST Samoa Apia": "Sa’aatii Istaandaardii Apia", "SST Samoa Apia DST": "Sa’aatii Guyyaa Apia", "SST Samoa DST": "Sa’aatii Guyyaa Saamowaa", "SYOT": "Sa’aatii Syowa", "TAAF": "Sa’aatii Firaans Kibbaa fi Antaarktikaa", "TAHT": "Sa’aatii Tahiti", "TJT": "Sa’aatii Tajikistaan", "TKT": "Sa’aatii Takelawu", "TLT": "Sa’aatii Tiimoor Bahaa", "TMST": "Sa’aatii Bonaa Turkemenistaan", "TMT": "Sa’aatii Istaandaardii Turkemenistaan", "TOST": "Sa’aatii Bonaa Tonga", "TOT": "Sa’aatii Istaandaardii Tonga", "TVT": "Sa’aatii Tuvalu", "TWT": "Sa’aatii Istaandaardii Tayipeyi", "TWT DST": "Sa’aatii Guyyaa Tayipeyi", "ULAST": "Sa’aatii Bonaa Ulaanbaatar", "ULAT": "Sa’aatii Istaandaardii Ulaanbaatar", "UYST": "Sa’aatii Bonaa Yuraagaayi", "UYT": "Sa’aatii Istaandaardii Yuraagaayi", "UZT": "Sa’aatii Istaandaardii Uzbeekistaan", "UZT DST": "Sa’aatii Bonaa Uzbeekistaan", "VET": "Sa’aatii Veenzuweelaa", "VLAST": "Sa’aatii Bonaa Vladivostok", "VLAT": "Sa’aatii Istaandaardii Vladivostok", "VOLST": "Sa’aatii Bonaa Volgograd", "VOLT": "Sa’aatii Istaandaardii Volgograd", "VOST": "Sa’aatii Vostok", "VUT": "Sa’aatii Istaandaardii Vanuwatu", "VUT DST": "Sa’aatii Bonaa Vanuwatu", "WAKT": "Sa’aatii Odola Wake", "WARST": "Sa’aatii Bonaa Arjentiinaa Dhihaa", "WART": "Sa’aatii Istaandaardii Arjentiinaa Dhihaa", "WAST": "Sa’aatii Afrikaa Dhihaa", "WAT": "Sa’aatii Afrikaa Dhihaa", "WESZ": "Sa’aatii Bonaa Awurooppaa Dhihaa", "WEZ": "Sa’aatii Istaandaardii Awurooppaa Dhihaa", "WFT": "Sa’aatii Wallis fi Futuna", "WGST": "Sa’aatii Bonaa Giriinlaand Dhihaa", "WGT": "Sa’aatii Istaandaardii Giriinlaand Dhihaa", "WIB": "Sa’aatii Indooneeshiyaa Dhihaa", "WIT": "Sa’aatii Indooneshiyaa Bahaa", "WITA": "Sa’aatii Indooneeshiyaa Gidduugaleessaa", "YAKST": "Sa’aatii Bonaa Yakutsk", "YAKT": "Sa’aatii Istaandardii Yakutsk", "YEKST": "Sa’aatii Bonaa Yekaterinburg", "YEKT": "Sa’aatii Istaandaardii Yekaterinburg", "YST": "Sa’aatii Yuukoon", "МСК": "Sa’aatii Istaandaardii Mooskoo", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Sa’aatii Kaazaakistaan Dhihaa", "شىعىش قازاق ەلى": "Sa’aatii Kaazaakistaan Bahaa", "قازاق ەلى": "Sa’aatii Kaazaakistaan", "قىرعىزستان": "Sa’aatii Kiyirigiyistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Sa’aatii Bonaa Azeeroos"},
	}
}

// Locale returns the current translators string locale
func (om *om_ET) Locale() string {
	return om.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'om_ET'
func (om *om_ET) PluralsCardinal() []locales.PluralRule {
	return om.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'om_ET'
func (om *om_ET) PluralsOrdinal() []locales.PluralRule {
	return om.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'om_ET'
func (om *om_ET) PluralsRange() []locales.PluralRule {
	return om.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'om_ET'
func (om *om_ET) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'om_ET'
func (om *om_ET) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'om_ET'
func (om *om_ET) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (om *om_ET) MonthAbbreviated(month time.Month) string {
	if len(om.monthsAbbreviated) == 0 {
		return ""
	}
	return om.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (om *om_ET) MonthsAbbreviated() []string {
	return om.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (om *om_ET) MonthNarrow(month time.Month) string {
	if len(om.monthsNarrow) == 0 {
		return ""
	}
	return om.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (om *om_ET) MonthsNarrow() []string {
	return om.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (om *om_ET) MonthWide(month time.Month) string {
	if len(om.monthsWide) == 0 {
		return ""
	}
	return om.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (om *om_ET) MonthsWide() []string {
	return om.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (om *om_ET) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(om.daysAbbreviated) == 0 {
		return ""
	}
	return om.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (om *om_ET) WeekdaysAbbreviated() []string {
	return om.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (om *om_ET) WeekdayNarrow(weekday time.Weekday) string {
	if len(om.daysNarrow) == 0 {
		return ""
	}
	return om.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (om *om_ET) WeekdaysNarrow() []string {
	return om.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (om *om_ET) WeekdayShort(weekday time.Weekday) string {
	if len(om.daysShort) == 0 {
		return ""
	}
	return om.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (om *om_ET) WeekdaysShort() []string {
	return om.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (om *om_ET) WeekdayWide(weekday time.Weekday) string {
	if len(om.daysWide) == 0 {
		return ""
	}
	return om.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (om *om_ET) WeekdaysWide() []string {
	return om.daysWide
}

// Decimal returns the decimal point of number
func (om *om_ET) Decimal() string {
	return om.decimal
}

// Group returns the group of number
func (om *om_ET) Group() string {
	return om.group
}

// Group returns the minus sign of number
func (om *om_ET) Minus() string {
	return om.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'om_ET' and handles both Whole and Real numbers based on 'v'
func (om *om_ET) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, om.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, om.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, om.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'om_ET' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (om *om_ET) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, om.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, om.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, om.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'om_ET'
func (om *om_ET) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := om.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, om.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, om.group[0])
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
		b = append(b, om.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, om.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'om_ET'
// in accounting notation.
func (om *om_ET) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := om.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, om.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, om.group[0])
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

		b = append(b, om.minus[0])

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
			b = append(b, om.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'om_ET'
func (om *om_ET) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'om_ET'
func (om *om_ET) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, om.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'om_ET'
func (om *om_ET) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, om.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'om_ET'
func (om *om_ET) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, om.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, om.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'om_ET'
func (om *om_ET) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, om.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, om.periodsAbbreviated[0]...)
	} else {
		b = append(b, om.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'om_ET'
func (om *om_ET) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, om.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, om.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, om.periodsAbbreviated[0]...)
	} else {
		b = append(b, om.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'om_ET'
func (om *om_ET) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, om.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, om.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, om.periodsAbbreviated[0]...)
	} else {
		b = append(b, om.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'om_ET'
func (om *om_ET) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, om.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, om.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, om.periodsAbbreviated[0]...)
	} else {
		b = append(b, om.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := om.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
