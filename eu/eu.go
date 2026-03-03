package eu

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type eu struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentPrefix          string
	timeSeparator          string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'eu' locale
func New() locales.Translator {
	return &eu{
		locale:                 "eu",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "₧", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentPrefix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "urt.", "ots.", "mar.", "api.", "mai.", "eka.", "uzt.", "abu.", "ira.", "urr.", "aza.", "abe."},
		monthsNarrow:           []string{"", "U", "O", "M", "A", "M", "E", "U", "A", "I", "U", "A", "A"},
		monthsWide:             []string{"", "urtarrila", "otsaila", "martxoa", "apirila", "maiatza", "ekaina", "uztaila", "abuztua", "iraila", "urria", "azaroa", "abendua"},
		daysAbbreviated:        []string{"ig.", "al.", "ar.", "az.", "og.", "or.", "lr."},
		daysNarrow:             []string{"I", "A", "A", "A", "O", "O", "L"},
		daysWide:               []string{"igandea", "astelehena", "asteartea", "asteazkena", "osteguna", "ostirala", "larunbata"},
		timezones:              map[string]string{"ACDT": "Australiako erdialdeko udako ordua", "ACST": "Australiako erdialdeko ordu estandarra", "ACT": "Acreko ordu estandarra", "ACWDT": "Australiako erdi-mendebaldeko udako ordua", "ACWST": "Australiako erdi-mendebaldeko ordu estandarra", "ADT": "Ipar Amerikako Atlantikoko udako ordua", "ADT Arabia": "Arabiako udako ordua", "AEDT": "Australiako ekialdeko udako ordua", "AEST": "Australiako ekialdeko ordu estandarra", "AFT": "Afganistango ordua", "AKDT": "Alaskako udako ordua", "AKST": "Alaskako ordu estandarra", "AMST": "Amazoniako udako ordua", "AMST Armenia": "Armeniako udako ordua", "AMT": "Amazoniako ordu estandarra", "AMT Armenia": "Armeniako ordu estandarra", "ANAST": "Anadyrreko udako ordua", "ANAT": "Anadyrreko ordu estandarra", "ARST": "Argentinako udako ordua", "ART": "Argentinako ordu estandarra", "AST": "Ipar Amerikako Atlantikoko ordu estandarra", "AST Arabia": "Arabiako ordu estandarra", "AWDT": "Australiako mendebaldeko udako ordua", "AWST": "Australiako mendebaldeko ordu estandarra", "AZST": "Azerbaijango udako ordua", "AZT": "Azerbaijango ordu estandarra", "BDT Bangladesh": "Bangladesheko udako ordua", "BNT": "Brunei Darussalamgo ordua", "BOT": "Boliviako ordua", "BRST": "Brasiliako udako ordua", "BRT": "Brasiliako ordu estandarra", "BST Bangladesh": "Bangladesheko ordu estandarra", "BT": "Bhutango ordua", "CAST": "Caseyko ordua", "CAT": "Afrikako erdialdeko ordua", "CCT": "Cocos uharteetako ordua", "CDT": "Ipar Amerikako erdialdeko udako ordua", "CHADT": "Chathamgo udako ordua", "CHAST": "Chathamgo ordu estandarra", "CHUT": "Chuukeko ordua", "CKT": "Cook uharteetako ordu estandarra", "CKT DST": "Cook uharteetako uda erdialdeko ordua", "CLST": "Txileko udako ordua", "CLT": "Txileko ordu estandarra", "COST": "Kolonbiako udako ordua", "COT": "Kolonbiako ordu estandarra", "CST": "Ipar Amerikako erdialdeko ordu estandarra", "CST China": "Txinako ordu estandarra", "CST China DST": "Txinako udako ordua", "CVST": "Cabo Verdeko udako ordua", "CVT": "Cabo Verdeko ordu estandarra", "CXT": "Christmas uharteko ordua", "ChST": "Chamorroko ordu estandarra", "ChST NMI": "Ipar Mariana uharteetako ordua", "CuDT": "Kubako udako ordua", "CuST": "Kubako ordu estandarra", "DAVT": "Daviseko ordua", "DDUT": "Dumont-d’Urvilleko ordua", "EASST": "Pazko uharteko udako ordua", "EAST": "Pazko uharteko ordu estandarra", "EAT": "Afrikako ekialdeko ordua", "ECT": "Ekuadorreko ordua", "EDT": "Ipar Amerikako ekialdeko udako ordua", "EGDT": "Groenlandiako ekialdeko udako ordua", "EGST": "Groenlandiako ekialdeko ordu estandarra", "EST": "Ipar Amerikako ekialdeko ordu estandarra", "FEET": "Europako ekialde urruneko ordua", "FJT": "Fijiko ordu estandarra", "FJT Summer": "Fijiko udako ordua", "FKST": "Falkland uharteetako udako ordua", "FKT": "Falkland uharteetako ordu estandarra", "FNST": "Fernando de Noronhako udako ordua", "FNT": "Fernando de Noronhako ordu estandarra", "GALT": "Galapagoetako ordua", "GAMT": "Gambierretako ordua", "GEST": "Georgiako udako ordua", "GET": "Georgiako ordu estandarra", "GFT": "Guyana Frantseseko ordua", "GIT": "Gilbert uharteetako ordua", "GMT": "Greenwichko meridianoaren ordua", "GNSST": "GNSST", "GNST": "GNST", "GST": "Hegoaldeko Georgietako ordua", "GST Guam": "Guameko ordu estandarra", "GYT": "Guyanako ordua", "HADT": "Hawaii-Aleutiar uharteetako udako ordua", "HAST": "Hawaii-Aleutiar uharteetako ordu estandarra", "HKST": "Hong Kongo udako ordua", "HKT": "Hong Kongo ordu estandarra", "HOVST": "Khovdeko udako ordua", "HOVT": "Khovdeko ordu estandarra", "ICT": "Indotxinako ordua", "IDT": "Israelgo udako ordua", "IOT": "Indiako ozeanoko ordua", "IRKST": "Irkutskeko udako ordua", "IRKT": "Irkutskeko ordu estandarra", "IRST": "Irango ordu estandarra", "IRST DST": "Irango udako ordua", "IST": "Indiako ordua", "IST Israel": "Israelgo ordu estandarra", "JDT": "Japoniako udako ordua", "JST": "Japoniako ordu estandarra", "KOST": "Kosraeko ordua", "KRAST": "Krasnoiarskeko udako ordua", "KRAT": "Krasnoiarskeko ordu estandarra", "KST": "Koreako ordu estandarra", "KST DST": "Koreako udako ordua", "LHDT": "Lord Howeko udako ordua", "LHST": "Lord Howeko ordu estandarra", "LINT": "Line uharteetako ordua", "MAGST": "Magadango udako ordua", "MAGT": "Magadango ordu estandarra", "MART": "Markesetako ordua", "MAWT": "Mawsoneko ordua", "MDT": "Macaoko udako ordua", "MESZ": "Europako erdialdeko udako ordua", "MEZ": "Europako erdialdeko ordu estandarra", "MHT": "Marshall Uharteetako ordua", "MMT": "Myanmarreko ordua", "MSD": "Moskuko udako ordua", "MST": "Macaoko ordu estandarra", "MUST": "Maurizioko udako ordua", "MUT": "Maurizioko ordu estandarra", "MVT": "Maldivetako ordua", "MYT": "Malaysiako ordua", "NCT": "Kaledonia Berriko ordu estandarra", "NDT": "Ternuako udako ordua", "NDT New Caledonia": "Kaledonia Berriko udako ordua", "NFDT": "Norfolk uharteetako udako ordua", "NFT": "Norfolk uharteetako ordu estandarra", "NOVST": "Novosibirskeko udako ordua", "NOVT": "Novosibirskeko ordu estandarra", "NPT": "Nepalgo ordua", "NRT": "Nauruko ordua", "NST": "Ternuako ordu estandarra", "NUT": "Niueko ordua", "NZDT": "Zeelanda Berriko udako ordua", "NZST": "Zeelanda Berriko ordu estandarra", "OESZ": "Europako ekialdeko udako ordua", "OEZ": "Europako ekialdeko ordu estandarra", "OMSST": "Omskeko udako ordua", "OMST": "Omskeko ordu estandarra", "PDT": "Ipar Amerikako Pazifikoko udako ordua", "PDTM": "Mexikoko Pazifikoko udako ordua", "PETDT": "Petropavlovsk-Kamchatskiko udako ordua", "PETST": "Petropavlovsk-Kamchatskiko ordu estandarra", "PGT": "Papua Ginea Berriko ordua", "PHOT": "Phoenix uharteetako ordua", "PKT": "Pakistango ordu estandarra", "PKT DST": "Pakistango udako ordua", "PMDT": "Saint-Pierre eta Mikeluneko udako ordua", "PMST": "Saint-Pierre eta Mikeluneko ordu estandarra", "PONT": "Ponapeko ordua", "PST": "Ipar Amerikako Pazifikoko ordu estandarra", "PST Philippine": "Filipinetako ordu estandarra", "PST Philippine DST": "Filipinetako udako ordua", "PST Pitcairn": "Pitcairneko ordua", "PSTM": "Mexikoko Pazifikoko ordu estandarra", "PWT": "Palauko ordua", "PYST": "Paraguaiko udako ordua", "PYT": "Paraguaiko ordu estandarra", "PYT Korea": "Piongiangeko ordua", "RET": "Reunioneko ordua", "ROTT": "Rotherako ordua", "SAKST": "Sakhalingo udako ordua", "SAKT": "Sakhalingo ordu estandarra", "SAMST": "Samarako udako ordua", "SAMT": "Samarako ordu estandarra", "SAST": "Afrikako hegoaldeko ordua", "SBT": "Salomon Uharteetako ordua", "SCT": "Seychelle uharteetako ordua", "SGT": "Singapurreko ordu estandarra", "SLST": "Lankako ordua", "SRT": "Surinamgo ordua", "SST Samoa": "Samoako ordu estandarra", "SST Samoa Apia": "Apiako ordu estandarra", "SST Samoa Apia DST": "Apiako udako ordua", "SST Samoa DST": "Samoako udako ordua", "SYOT": "Syowako ordua", "TAAF": "Frantziaren lurralde austral eta antartikoetako ordutegia", "TAHT": "Tahitiko ordua", "TJT": "Tadjikistango ordua", "TKT": "Tokelauko ordua", "TLT": "Ekialdeko Timorreko ordua", "TMST": "Turkmenistango udako ordua", "TMT": "Turkmenistango ordu estandarra", "TOST": "Tongako udako ordua", "TOT": "Tongako ordu estandarra", "TVT": "Tuvaluko ordua", "TWT": "Taipeiko ordu estandarra", "TWT DST": "Taipeiko udako ordua", "ULAST": "Ulan Batorreko udako ordua", "ULAT": "Ulan Batorreko ordu estandarra", "UYST": "Uruguaiko udako ordua", "UYT": "Uruguaiko ordu estandarra", "UZT": "Uzbekistango ordu estandarra", "UZT DST": "Uzbekistango udako ordua", "VET": "Venezuelako ordua", "VLAST": "Vladivostokeko udako ordua", "VLAT": "Vladivostokeko ordu estandarra", "VOLST": "Volgogradeko udako ordua", "VOLT": "Volgogradeko ordu estandarra", "VOST": "Vostokeko ordua", "VUT": "Vanuatuko ordu estandarra", "VUT DST": "Vanuatuko udako ordua", "WAKT": "Wake uharteko ordua", "WARST": "Argentina mendebaldeko udako ordua", "WART": "Argentina mendebaldeko ordu estandarra", "WAST": "Afrikako mendebaldeko ordua", "WAT": "Afrikako mendebaldeko ordua", "WESZ": "Europako mendebaldeko udako ordua", "WEZ": "Europako mendebaldeko ordu estandarra", "WFT": "Wallis eta Futunako ordutegia", "WGST": "Groenlandiako mendebaldeko udako ordua", "WGT": "Groenlandiako mendebaldeko ordu estandarra", "WIB": "Indonesiako mendebaldeko ordua", "WIT": "Indonesiako ekialdeko ordua", "WITA": "Indonesiako erdialdeko ordua", "YAKST": "Jakutskeko udako ordua", "YAKT": "Jakutskeko ordu estandarra", "YEKST": "Jekaterinburgeko udako ordua", "YEKT": "Jekaterinburgeko ordu estandarra", "YST": "Yukongo ordua", "МСК": "Moskuko ordu estandarra", "اقتاۋ": "Aktauko ordu estandarra", "اقتاۋ قالاسى": "Aktauko udako ordua", "اقتوبە": "Aktobeko ordu estandarra", "اقتوبە قالاسى": "Aktobeko udako ordua", "الماتى": "Almatyko ordu estandarra", "الماتى قالاسى": "Almatyko udako ordua", "باتىس قازاق ەلى": "Kazakhstango mendebaldeko ordua", "شىعىش قازاق ەلى": "Kazakhstango ekialdeko ordua", "قازاق ەلى": "Kazakhstango ordua", "قىرعىزستان": "Kirgizistango ordua", "قىزىلوردا": "Kyzylordako ordu estandarra", "قىزىلوردا قالاسى": "Kyzylordako udako ordua", "∅∅∅": "Azoreetako udako ordua"},
	}
}

// Locale returns the current translators string locale
func (eu *eu) Locale() string {
	return eu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'eu'
func (eu *eu) PluralsCardinal() []locales.PluralRule {
	return eu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'eu'
func (eu *eu) PluralsOrdinal() []locales.PluralRule {
	return eu.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'eu'
func (eu *eu) PluralsRange() []locales.PluralRule {
	return eu.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'eu'
func (eu *eu) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'eu'
func (eu *eu) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'eu'
func (eu *eu) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (eu *eu) MonthAbbreviated(month time.Month) string {
	if len(eu.monthsAbbreviated) == 0 {
		return ""
	}
	return eu.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (eu *eu) MonthsAbbreviated() []string {
	return eu.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (eu *eu) MonthNarrow(month time.Month) string {
	if len(eu.monthsNarrow) == 0 {
		return ""
	}
	return eu.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (eu *eu) MonthsNarrow() []string {
	return eu.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (eu *eu) MonthWide(month time.Month) string {
	if len(eu.monthsWide) == 0 {
		return ""
	}
	return eu.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (eu *eu) MonthsWide() []string {
	return eu.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (eu *eu) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(eu.daysAbbreviated) == 0 {
		return ""
	}
	return eu.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (eu *eu) WeekdaysAbbreviated() []string {
	return eu.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (eu *eu) WeekdayNarrow(weekday time.Weekday) string {
	if len(eu.daysNarrow) == 0 {
		return ""
	}
	return eu.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (eu *eu) WeekdaysNarrow() []string {
	return eu.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (eu *eu) WeekdayShort(weekday time.Weekday) string {
	if len(eu.daysShort) == 0 {
		return ""
	}
	return eu.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (eu *eu) WeekdaysShort() []string {
	return eu.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (eu *eu) WeekdayWide(weekday time.Weekday) string {
	if len(eu.daysWide) == 0 {
		return ""
	}
	return eu.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (eu *eu) WeekdaysWide() []string {
	return eu.daysWide
}

// Decimal returns the decimal point of number
func (eu *eu) Decimal() string {
	return eu.decimal
}

// Group returns the group of number
func (eu *eu) Group() string {
	return eu.group
}

// Group returns the minus sign of number
func (eu *eu) Minus() string {
	return eu.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'eu' and handles both Whole and Real numbers based on 'v'
func (eu *eu) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(eu.minus) - 1; j >= 0; j-- {
			b = append(b, eu.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'eu' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (eu *eu) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(eu.minus) - 1; j >= 0; j-- {
			b = append(b, eu.minus[j])
		}
	}

	for j := len(eu.percentPrefix) - 1; j >= 0; j-- {
		b = append(b, eu.percentPrefix[j])
	}

	b = append(b, eu.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'eu'
func (eu *eu) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := eu.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(eu.minus) - 1; j >= 0; j-- {
			b = append(b, eu.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, eu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, eu.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'eu'
// in accounting notation.
func (eu *eu) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := eu.currencies[currency]
	l := len(s) + len(symbol) + 8 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, eu.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, eu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, eu.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, eu.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'eu'
func (eu *eu) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'eu'
func (eu *eu) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x28, 0x65}...)
	b = append(b, []byte{0x29, 0x6b, 0x6f}...)
	b = append(b, []byte{0x20}...)
	b = append(b, eu.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x28, 0x61}...)
	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'eu'
func (eu *eu) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x28, 0x65}...)
	b = append(b, []byte{0x29, 0x6b, 0x6f}...)
	b = append(b, []byte{0x20}...)
	b = append(b, eu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x72, 0x65, 0x6e}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x28, 0x61}...)
	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'eu'
func (eu *eu) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x28, 0x65}...)
	b = append(b, []byte{0x29, 0x6b, 0x6f}...)
	b = append(b, []byte{0x20}...)
	b = append(b, eu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x72, 0x65, 0x6e}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x28, 0x61}...)
	b = append(b, []byte{0x29, 0x2c, 0x20}...)
	b = append(b, eu.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'eu'
func (eu *eu) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'eu'
func (eu *eu) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'eu'
func (eu *eu) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'eu'
func (eu *eu) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := eu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
