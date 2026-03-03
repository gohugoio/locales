package pcm

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type pcm struct {
	locale            string
	pluralsCardinal   []locales.PluralRule
	pluralsOrdinal    []locales.PluralRule
	pluralsRange      []locales.PluralRule
	decimal           string
	group             string
	minus             string
	percent           string
	timeSeparator     string
	currencies        []string // idx = enum of currency code
	monthsAbbreviated []string
	monthsNarrow      []string
	monthsWide        []string
	daysAbbreviated   []string
	daysNarrow        []string
	daysShort         []string
	daysWide          []string
	timezones         map[string]string
}

// New returns a new instance of translator for the 'pcm' locale
func New() locales.Translator {
	return &pcm{
		locale:            "pcm",
		pluralsCardinal:   []locales.PluralRule{2, 6},
		pluralsOrdinal:    nil,
		pluralsRange:      []locales.PluralRule{6},
		decimal:           ".",
		group:             ",",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "p.", "BYR", "BZD", "KA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "₦", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "Jén", "Fẹ́b", "Mach", "Épr", "Mee", "Jun", "Jul", "Ọgọ", "Sẹp", "Ọkt", "Nọv", "Dis"},
		monthsNarrow:      []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:        []string{"", "Jénúári", "Fẹ́búári", "Mach", "Éprel", "Mee", "Jun", "Julai", "Ọgọst", "Sẹptẹ́mba", "Ọktóba", "Nọvẹ́mba", "Disẹ́mba"},
		daysAbbreviated:   []string{"Sọ́n", "Mọ́n", "Tiú", "Wẹ́n", "Tọ́z", "Fraí", "Sát"},
		daysWide:          []string{"Sọ́ndè", "Mọ́ndè", "Tiúzdè", "Wẹ́nẹ́zdè", "Tọ́zdè", "Fraídè", "Sátọdè"},
		timezones:         map[string]string{"ACDT": "Ọstrélia Mídúl Délaít Taim", "ACST": "Ọstrélia Mídúl Fíksd Taim", "ACT": "ACT", "ACWDT": "Ọstrélia Mídúl Wẹ́stán Délaít Taim", "ACWST": "Ọstrélia Mídúl Wẹ́stán Fíksd Taim", "ADT": "Atlántík Délaít Taim", "ADT Arabia": "Arébiá Délaít Taim", "AEDT": "Ọstrélia Ístán Délaít Taim", "AEST": "Ọstrélia Ístán Fíksd Taim", "AFT": "Afgánístan Taim", "AKDT": "Aláská Délaít Taim", "AKST": "Aláská Fíksd Taim", "AMST": "Ámázọn Họ́t Sízín Taim", "AMST Armenia": "Armẹ́nia Họ́t Sízin Taim", "AMT": "Ámázọn Fíksd Taim", "AMT Armenia": "Armẹ́nia Fíksd Taim", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Ajẹntína Họ́t Sízín Taim", "ART": "Ajẹntína Fíksd Taim", "AST": "Atlántík Fíksd Taim", "AST Arabia": "Arébiá Fíksd Taim", "AWDT": "Ọstrélia Wẹ́stán Délaít Taim", "AWST": "Ọstrélia Wẹ́stán Fíksd Taim", "AZST": "Azẹrbaijan Họ́t Sízin Taim", "AZT": "Azẹrbaijan Fíksd Taim", "BDT Bangladesh": "Bangladẹsh Délaít Taim", "BNT": "Brunẹi Darúsalam Taim", "BOT": "Bolívia Fíksd Taim", "BRST": "Brasília Họ́t Sízín Taim", "BRT": "Brasília Fíksd Taim", "BST Bangladesh": "Bangladẹsh Fíksd Taim", "BT": "Butan Taim", "CAST": "CAST", "CAT": "Mídúl Áfríká Taim", "CCT": "Kókós Aílands Taim", "CDT": "Nọ́t Amẹ́ríká Mídúl Ériá Délaít Taim", "CHADT": "Chátam Délaít Taim", "CHAST": "Chátam Fíksd Taim", "CHUT": "Chuk Taim", "CKT": "Kúk Aílands Fíksd Taim", "CKT DST": "Kúk Aílands Haf Họ́t Sízin Taim", "CLST": "Chílẹ Họ́t Sízín Taim", "CLT": "Chílẹ Fíksd Taim", "COST": "Kolómbia Họ́t Sízín Taim", "COT": "Kolómbia Fíksd Taim", "CST": "Nọ́t Amẹ́ríká Mídúl Ériá Fíksd Taim", "CST China": "Chaína Fíksd Taim", "CST China DST": "Chaína Délaít Taim", "CVST": "Kep Vẹ́d Họ́t Sízin Taim", "CVT": "Kep Vẹ́d Fíksd Taim", "CXT": "Krísmás Aíland Taim", "ChST": "Chamóro Fíksd Taim", "ChST NMI": "ChST NMI", "CuDT": "Kúba Délaít Taim", "CuST": "Kúba Fíksd Taim", "DAVT": "Dévis Taim", "DDUT": "Diúmọ́n-d’Uvil Taim", "EASST": "Ísta Họ́t Sízín Taim", "EAST": "Ísta Fíksd Taim", "EAT": "Íst Áfríká Taim", "ECT": "Ẹ́kwuádọ Taim", "EDT": "Nọ́t Amẹ́ríká Ístán Ériá Délaít Taim", "EGDT": "Íist Grínlánd Họ́t Sízin Taim", "EGST": "Íist Grínlánd Fíksd Taim", "EST": "Nọ́t Amẹ́ríká Ístán Ériá Fíksd Taim", "FEET": "Faá-Ístán Yúrop Taim", "FJT": "Fíji Fíksd Taim", "FJT Summer": "Fíji Họ́t Sízín Taim", "FKST": "Fọ́lkland Họ́t Sízín Taim", "FKT": "Fọ́lkland Fíksd Taim", "FNST": "Fẹrnándó di Nọrónia Họ́t Sízín Taim", "FNT": "Fẹrnándó di Nọrónia Fíksd Taim", "GALT": "Galápágọs Taim", "GAMT": "Gámbiẹr Taim", "GEST": "Jọ́jia Họ́t Sízin Taim", "GET": "Jọ́jia Fíksd Taim", "GFT": "Frẹ́nch Giána Taim", "GIT": "Gílbat Aílands Taim", "GMT": "Grínwích Mín Taim", "GNSST": "GNSST", "GNST": "GNST", "GST": "Saút Jọ́jia Taim", "GST Guam": "GST Guam", "GYT": "Gayána Taim", "HADT": "Hawaií-Elúshián Délaít Taim", "HAST": "Hawaií-Elúshián Fíksd Taim", "HKST": "Họng Kọng Họ́t Sízin Taim", "HKT": "Họng Kọng Fíksd Taim", "HOVST": "Hovd Họ́t Sízin Taim", "HOVT": "Hovd Fíksd Taim", "ICT": "Indochaína Taim", "IDT": "Ízrẹl Délaít Taim", "IOT": "Índián Óshẹ́n Taim", "IRKST": "Irkútsk Họ́t Sízin Taim", "IRKT": "Irkútsk Fíksd Taim", "IRST": "Iran Fíksd Taim", "IRST DST": "Iran Délaít Taim", "IST": "Índia Fíksd Taim", "IST Israel": "Ízrẹl Fíksd Taim", "JDT": "Japan Délaít Taim", "JST": "Japan Fíksd Taim", "KOST": "Kọ́sraẹ Taim", "KRAST": "Krasnoyask Họ́t Sízin Taim", "KRAT": "Krasnoyask Fíksd Taim", "KST": "Koria Fíksd Taim", "KST DST": "Koria Délaít Taim", "LHDT": "Lọd Haú Délaít Taim", "LHST": "Lọd Haú Fíksd Taim", "LINT": "Laín Aílands Taim", "MAGST": "Mágádan Họ́t Sízin Taim", "MAGT": "Mágádan Fíksd Taim", "MART": "Makwẹ́sas Taim", "MAWT": "Mọ́sọn Taim", "MDT": "Nọ́t Amẹ́ríká Maúntin Ériá Délaít Taim", "MESZ": "Mídúl Yúrop Họ́t Sízin Taim", "MEZ": "Mídúl Yúrop Fíksd Taim", "MHT": "Máshal Aílands Taim", "MMT": "Miánma Taim", "MSD": "Mọ́sko Họ́t Sízin Taim", "MST": "Nọ́t Amẹ́ríká Maúntin Ériá Fíksd Taim", "MUST": "Mọríshọs Họ́t Sízin Taim", "MUT": "Mọríshọs Fíksd Taim", "MVT": "Mọ́divs Taim", "MYT": "Maléshia Taim", "NCT": "Niú Kalẹdónia Fíksd Taim", "NDT": "Niúfaúndlánd Délaít Taim", "NDT New Caledonia": "Niú Kalẹdónia Họ́t Sízin Taim", "NFDT": "Nọ́rfọ́lk Aíland Délaít Taim", "NFT": "Nọ́rfọ́lk Aíland Fíksd Taim", "NOVST": "Novosibisk Họ́t Sízin Taim", "NOVT": "Novosibisk Fíksd Taim", "NPT": "Nẹpọl Taim", "NRT": "Naúru Taim", "NST": "Niúfaúndlánd Fíksd Taim", "NUT": "Niúẹ Taim", "NZDT": "Niú Ziland Délaít Taim", "NZST": "Niú Ziland Fíksd Taim", "OESZ": "Ístán Yúrop Họ́t Sízin Taim", "OEZ": "Ístán Yúrop Fíksd Taim", "OMSST": "Ọmsk Họ́t Sízin Taim", "OMST": "Ọmsk Fíksd Taim", "PDT": "Nọ́t Amẹ́ríká Pasífík Ériá Délaít Taim", "PDTM": "Mẹ́ksíkó Pasífík Délaít Taim", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Pápuá Niú Gíni Taim", "PHOT": "Fíniks Aílands Taim", "PKT": "Pákístan Fíksd Taim", "PKT DST": "Pákístan Họ́t Sízin Taim", "PMDT": "Sent Piẹr an Míkẹlọn Délaít Taim", "PMST": "Sent Piẹr an Míkẹlọn Fíksd Taim", "PONT": "Pónápẹ Taim", "PST": "Nọ́t Amẹ́ríká Pasífík Ériá Fíksd Taim", "PST Philippine": "Fílípin Fíksd Taim", "PST Philippine DST": "Fílípin Họt Sízin Taim", "PST Pitcairn": "Pítkan Taim", "PSTM": "Mẹ́ksíkó Pasífík Fíksd Taim", "PWT": "Paláu Taim", "PYST": "Párágwue Họ́t Sízín Taim", "PYT": "Párágwue Fíksd Taim", "PYT Korea": "Piọngyang Taim", "RET": "Riyúniọn Taim", "ROTT": "Rotẹ́ra Taim", "SAKST": "Sákhalin Họ́t Sízin Taim", "SAKT": "Sákhalin Fíksd Taim", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Saút Áfríká Fíksd Taim", "SBT": "Sólómọ́n Aílands Taim", "SCT": "Sẹ́chẹls Taim", "SGT": "Singapọ Taim", "SLST": "SLST", "SRT": "Súrínam Taim", "SST Samoa": "Sámoá Fíksd Taim", "SST Samoa Apia": "Ápia Fíksd Taim", "SST Samoa Apia DST": "Ápia Délaít Taim", "SST Samoa DST": "Sámoá Délaít Taim", "SYOT": "Siówa Taim", "TAAF": "Frẹ́nch Saútan an Antátík Taim", "TAHT": "Tahíti Taim", "TJT": "Tajíkistan Taim", "TKT": "Tokẹláu Taim", "TLT": "Íst Tímọ Taim", "TMST": "Tọkmẹnistan Họ́t Sízin Taim", "TMT": "Tọkmẹnistan Fíksd Taim", "TOST": "Tọ́nga Họ́t Sízin Taim", "TOT": "Tọ́nga Fíksd Taim", "TVT": "Tuválu Taim", "TWT": "Taipẹi Fíksd Taim", "TWT DST": "Taipẹi Délaít Taim", "ULAST": "Mọngólia Họ́t Sízin Taim", "ULAT": "Mọngólia Fíksd Taim", "UYST": "Yúrugwue Họ́t Sízín Taim", "UYT": "Yúrugwue Fíksd Taim", "UZT": "Uzbẹkistan Fíksd Taim", "UZT DST": "Uzbẹkistan Họ́t Sízin Taim", "VET": "Vẹnẹzuẹ́la Taim", "VLAST": "Vladivostok Họ́t Sízin Taim", "VLAT": "Vladivọstọk Fíksd Taim", "VOLST": "Volvógrad Họ́t Sízin Taim", "VOLT": "Volvógrad Fíksd Taim", "VOST": "Vọ́stọk Taim", "VUT": "Vanuátu Fíksd Taim", "VUT DST": "Vanuátu Sízin Taim", "WAKT": "Wék Aíland Taim", "WARST": "Wẹ́stán Ajẹntína Họ́t Sízín Taim", "WART": "Wẹ́stán Ajẹntína Fíksd Taim", "WAST": "Wẹ́st Áfríká Taim", "WAT": "Wẹ́st Áfríká Taim", "WESZ": "Wẹ́stán Yúrop Họ́t Sízin Taim", "WEZ": "Wẹ́stán Yúrop Fíksd Taim", "WFT": "Wális an Fútúna Taim", "WGST": "Wẹ́st Grínlánd Họ́t Sízin Taim", "WGT": "Wẹ́st Grínlánd Fíksd Taim", "WIB": "Wẹ́stán Indonẹ́shia Taim", "WIT": "Ístán Indonẹ́shia Taim", "WITA": "Mídúl Indonẹ́shia Taim", "YAKST": "Yékútsk Họ́t Sízin Taim", "YAKT": "Yékútsk Fíksd Taim", "YEKST": "Yẹketẹrínbug Họ́t Sízin Taim", "YEKT": "Yẹketẹrínbug Fíksd Taim", "YST": "Yukón Taim", "МСК": "Mọ́sko Fíksd Taim", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Wẹ́st Kazékstan Taim", "شىعىش قازاق ەلى": "Íst Kazékstan Taim", "قازاق ەلى": "Kazékstan Taim", "قىرعىزستان": "Kẹgistan Taim", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Pẹru Họ́t Sízín Taim"},
	}
}

// Locale returns the current translators string locale
func (pcm *pcm) Locale() string {
	return pcm.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pcm'
func (pcm *pcm) PluralsCardinal() []locales.PluralRule {
	return pcm.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pcm'
func (pcm *pcm) PluralsOrdinal() []locales.PluralRule {
	return pcm.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pcm'
func (pcm *pcm) PluralsRange() []locales.PluralRule {
	return pcm.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pcm'
func (pcm *pcm) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pcm'
func (pcm *pcm) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pcm'
func (pcm *pcm) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pcm *pcm) MonthAbbreviated(month time.Month) string {
	if len(pcm.monthsAbbreviated) == 0 {
		return ""
	}
	return pcm.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pcm *pcm) MonthsAbbreviated() []string {
	return pcm.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pcm *pcm) MonthNarrow(month time.Month) string {
	if len(pcm.monthsNarrow) == 0 {
		return ""
	}
	return pcm.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pcm *pcm) MonthsNarrow() []string {
	return pcm.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pcm *pcm) MonthWide(month time.Month) string {
	if len(pcm.monthsWide) == 0 {
		return ""
	}
	return pcm.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pcm *pcm) MonthsWide() []string {
	return pcm.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pcm *pcm) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(pcm.daysAbbreviated) == 0 {
		return ""
	}
	return pcm.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pcm *pcm) WeekdaysAbbreviated() []string {
	return pcm.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pcm *pcm) WeekdayNarrow(weekday time.Weekday) string {
	if len(pcm.daysNarrow) == 0 {
		return ""
	}
	return pcm.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pcm *pcm) WeekdaysNarrow() []string {
	return pcm.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pcm *pcm) WeekdayShort(weekday time.Weekday) string {
	if len(pcm.daysShort) == 0 {
		return ""
	}
	return pcm.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pcm *pcm) WeekdaysShort() []string {
	return pcm.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pcm *pcm) WeekdayWide(weekday time.Weekday) string {
	if len(pcm.daysWide) == 0 {
		return ""
	}
	return pcm.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pcm *pcm) WeekdaysWide() []string {
	return pcm.daysWide
}

// Decimal returns the decimal point of number
func (pcm *pcm) Decimal() string {
	return pcm.decimal
}

// Group returns the group of number
func (pcm *pcm) Group() string {
	return pcm.group
}

// Group returns the minus sign of number
func (pcm *pcm) Minus() string {
	return pcm.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pcm' and handles both Whole and Real numbers based on 'v'
func (pcm *pcm) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pcm.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pcm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pcm' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pcm *pcm) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pcm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pcm.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pcm'
func (pcm *pcm) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pcm.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pcm.group[0])
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
		b = append(b, pcm.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pcm.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pcm'
// in accounting notation.
func (pcm *pcm) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pcm.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pcm.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, pcm.group[0])
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

		b = append(b, pcm.minus[0])

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
			b = append(b, pcm.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pcm'
func (pcm *pcm) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'pcm'
func (pcm *pcm) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pcm.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pcm'
func (pcm *pcm) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pcm.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pcm'
func (pcm *pcm) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pcm.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pcm.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pcm'
func (pcm *pcm) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pcm'
func (pcm *pcm) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pcm'
func (pcm *pcm) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pcm'
func (pcm *pcm) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pcm.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pcm.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
