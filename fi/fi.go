package fi

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type fi struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'fi' locale
func New() locales.Translator {
	return &fi{
		locale:                 "fi",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ".",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "mk", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "tammi", "helmi", "maalis", "huhti", "touko", "kesä", "heinä", "elo", "syys", "loka", "marras", "joulu"},
		monthsNarrow:           []string{"", "T", "H", "M", "H", "T", "K", "H", "E", "S", "L", "M", "J"},
		monthsWide:             []string{"", "tammikuuta", "helmikuuta", "maaliskuuta", "huhtikuuta", "toukokuuta", "kesäkuuta", "heinäkuuta", "elokuuta", "syyskuuta", "lokakuuta", "marraskuuta", "joulukuuta"},
		daysAbbreviated:        []string{"su", "ma", "ti", "ke", "to", "pe", "la"},
		daysNarrow:             []string{"S", "M", "T", "K", "T", "P", "L"},
		daysWide:               []string{"sunnuntaina", "maanantaina", "tiistaina", "keskiviikkona", "torstaina", "perjantaina", "lauantaina"},
		periodsAbbreviated:     []string{"ap.", "ip."},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"eKr.", "jKr."},
		erasNarrow:             []string{"eKr", "jKr"},
		erasWide:               []string{"ennen Kristuksen syntymää", "jälkeen Kristuksen syntymän"},
		timezones:              map[string]string{"ACDT": "Keski-Australian kesäaika", "ACST": "Keski-Australian normaaliaika", "ACT": "Acren normaaliaika", "ACWDT": "Läntisen Keski-Australian kesäaika", "ACWST": "Läntisen Keski-Australian normaaliaika", "ADT": "Kanadan Atlantin kesäaika", "ADT Arabia": "Saudi-Arabian kesäaika", "AEDT": "Itä-Australian kesäaika", "AEST": "Itä-Australian normaaliaika", "AFT": "Afganistanin aika", "AKDT": "Alaskan kesäaika", "AKST": "Alaskan normaaliaika", "AMST": "Amazonin kesäaika", "AMST Armenia": "Armenian kesäaika", "AMT": "Amazonin normaaliaika", "AMT Armenia": "Armenian normaaliaika", "ANAST": "Anadyrin kesäaika", "ANAT": "Anadyrin normaaliaika", "ARST": "Argentiinan kesäaika", "ART": "Argentiinan normaaliaika", "AST": "Kanadan Atlantin normaaliaika", "AST Arabia": "Saudi-Arabian normaaliaika", "AWDT": "Länsi-Australian kesäaika", "AWST": "Länsi-Australian normaaliaika", "AZST": "Azerbaidžanin kesäaika", "AZT": "Azerbaidžanin normaaliaika", "BDT Bangladesh": "Bangladeshin kesäaika", "BNT": "Brunein aika", "BOT": "Bolivian aika", "BRST": "Brasilian kesäaika", "BRT": "Brasilian normaaliaika", "BST Bangladesh": "Bangladeshin normaaliaika", "BT": "Bhutanin aika", "CAST": "Caseyn aika", "CAT": "Keski-Afrikan aika", "CCT": "Kookossaarten aika", "CDT": "Yhdysvaltain keskinen kesäaika", "CHADT": "Chathamin kesäaika", "CHAST": "Chathamin normaaliaika", "CHUT": "Chuukin aika", "CKT": "Cookinsaarten normaaliaika", "CKT DST": "Cookinsaarten kesäaika", "CLST": "Chilen kesäaika", "CLT": "Chilen normaaliaika", "COST": "Kolumbian kesäaika", "COT": "Kolumbian normaaliaika", "CST": "Yhdysvaltain keskinen normaaliaika", "CST China": "Kiinan normaaliaika", "CST China DST": "Kiinan kesäaika", "CVST": "Kap Verden kesäaika", "CVT": "Kap Verden normaaliaika", "CXT": "Joulusaaren aika", "ChST": "Tšamorron aika", "ChST NMI": "Pohjois-Mariaanien aika", "CuDT": "Kuuban kesäaika", "CuST": "Kuuban normaaliaika", "DAVT": "Davisin aika", "DDUT": "Dumont d’Urvillen aika", "EASST": "Pääsiäissaaren kesäaika", "EAST": "Pääsiäissaaren normaaliaika", "EAT": "Itä-Afrikan aika", "ECT": "Ecuadorin aika", "EDT": "Yhdysvaltain itäinen kesäaika", "EGDT": "Itä-Grönlannin kesäaika", "EGST": "Itä-Grönlannin normaaliaika", "EST": "Yhdysvaltain itäinen normaaliaika", "FEET": "Itäisemmän Euroopan aika", "FJT": "Fidžin normaaliaika", "FJT Summer": "Fidžin kesäaika", "FKST": "Falklandinsaarten kesäaika", "FKT": "Falklandinsaarten normaaliaika", "FNST": "Fernando de Noronhan kesäaika", "FNT": "Fernando de Noronhan normaaliaika", "GALT": "Galápagossaarten aika", "GAMT": "Gambiersaarten aika", "GEST": "Georgian kesäaika", "GET": "Georgian normaaliaika", "GFT": "Ranskan Guayanan aika", "GIT": "Gilbertsaarten aika", "GMT": "Greenwichin normaaliaika", "GNSST": "GNSST", "GNST": "GNST", "GST": "Arabiemiirikuntien normaaliaika", "GST Guam": "Guamin aika", "GYT": "Guyanan aika", "HADT": "Havaijin-Aleuttien normaaliaika", "HAST": "Havaijin-Aleuttien normaaliaika", "HKST": "Hongkongin kesäaika", "HKT": "Hongkongin normaaliaika", "HOVST": "Hovdin kesäaika", "HOVT": "Hovdin normaaliaika", "ICT": "Indokiinan aika", "IDT": "Israelin kesäaika", "IOT": "Intian valtameren aika", "IRKST": "Irkutskin kesäaika", "IRKT": "Irkutskin normaaliaika", "IRST": "Iranin normaaliaika", "IRST DST": "Iranin kesäaika", "IST": "Intian aika", "IST Israel": "Israelin normaaliaika", "JDT": "Japanin kesäaika", "JST": "Japanin normaaliaika", "KOST": "Kosraen aika", "KRAST": "Krasnojarskin kesäaika", "KRAT": "Krasnojarskin normaaliaika", "KST": "Korean normaaliaika", "KST DST": "Korean kesäaika", "LHDT": "Lord Howen kesäaika", "LHST": "Lord Howen normaaliaika", "LINT": "Linesaarten aika", "MAGST": "Magadanin kesäaika", "MAGT": "Magadanin normaaliaika", "MART": "Marquesassaarten aika", "MAWT": "Mawsonin aika", "MDT": "Macaon kesäaika", "MESZ": "Keski-Euroopan kesäaika", "MEZ": "Keski-Euroopan normaaliaika", "MHT": "Marshallinsaarten aika", "MMT": "Myanmarin aika", "MSD": "Moskovan kesäaika", "MST": "Macaon normaaliaika", "MUST": "Mauritiuksen kesäaika", "MUT": "Mauritiuksen normaaliaika", "MVT": "Malediivien aika", "MYT": "Malesian aika", "NCT": "Uuden-Kaledonian normaaliaika", "NDT": "Newfoundlandin kesäaika", "NDT New Caledonia": "Uuden-Kaledonian kesäaika", "NFDT": "Norfolkinsaaren kesäaika", "NFT": "Norfolkinsaaren normaaliaika", "NOVST": "Novosibirskin kesäaika", "NOVT": "Novosibirskin normaaliaika", "NPT": "Nepalin aika", "NRT": "Naurun aika", "NST": "Newfoundlandin normaaliaika", "NUT": "Niuen aika", "NZDT": "Uuden-Seelannin kesäaika", "NZST": "Uuden-Seelannin normaaliaika", "OESZ": "Itä-Euroopan kesäaika", "OEZ": "Itä-Euroopan normaaliaika", "OMSST": "Omskin kesäaika", "OMST": "Omskin normaaliaika", "PDT": "Yhdysvaltain Tyynenmeren kesäaika", "PDTM": "Meksikon Tyynenmeren kesäaika", "PETDT": "Kamtšatkan kesäaika", "PETST": "Kamtšatkan normaaliaika", "PGT": "Papua-Uuden-Guinean aika", "PHOT": "Phoenixsaarten aika", "PKT": "Pakistanin normaaliaika", "PKT DST": "Pakistanin kesäaika", "PMDT": "Saint-Pierren ja Miquelonin kesäaika", "PMST": "Saint-Pierren ja Miquelonin normaaliaika", "PONT": "Pohnpein aika", "PST": "Yhdysvaltain Tyynenmeren normaaliaika", "PST Philippine": "Filippiinien normaaliaika", "PST Philippine DST": "Filippiinien kesäaika", "PST Pitcairn": "Pitcairnin aika", "PSTM": "Meksikon Tyynenmeren normaaliaika", "PWT": "Palaun aika", "PYST": "Paraguayn kesäaika", "PYT": "Paraguayn normaaliaika", "PYT Korea": "Pjongjangin aika", "RET": "Réunionin aika", "ROTT": "Rotheran aika", "SAKST": "Sahalinin kesäaika", "SAKT": "Sahalinin normaaliaika", "SAMST": "Samaran kesäaika", "SAMT": "Samaran normaaliaika", "SAST": "Etelä-Afrikan aika", "SBT": "Salomonsaarten aika", "SCT": "Seychellien aika", "SGT": "Singaporen aika", "SLST": "Sri Lankan aika", "SRT": "Surinamen aika", "SST Samoa": "Samoan normaaliaika", "SST Samoa Apia": "Apian normaaliaika", "SST Samoa Apia DST": "Apian kesäaika", "SST Samoa DST": "Samoan kesäaika", "SYOT": "Syowan aika", "TAAF": "Ranskan eteläisten ja antarktisten alueiden aika", "TAHT": "Tahitin aika", "TJT": "Tadžikistanin aika", "TKT": "Tokelaun aika", "TLT": "Itä-Timorin aika", "TMST": "Turkmenistanin kesäaika", "TMT": "Turkmenistanin normaaliaika", "TOST": "Tongan kesäaika", "TOT": "Tongan normaaliaika", "TVT": "Tuvalun aika", "TWT": "Taipein normaaliaika", "TWT DST": "Taipein kesäaika", "ULAST": "Ulan Batorin kesäaika", "ULAT": "Ulan Batorin normaaliaika", "UYST": "Uruguayn kesäaika", "UYT": "Uruguayn normaaliaika", "UZT": "Uzbekistanin normaaliaika", "UZT DST": "Uzbekistanin kesäaika", "VET": "Venezuelan aika", "VLAST": "Vladivostokin kesäaika", "VLAT": "Vladivostokin normaaliaika", "VOLST": "Volgogradin kesäaika", "VOLT": "Volgogradin normaaliaika", "VOST": "Vostokin aika", "VUT": "Vanuatun normaaliaika", "VUT DST": "Vanuatun kesäaika", "WAKT": "Waken aika", "WARST": "Länsi-Argentiinan kesäaika", "WART": "Länsi-Argentiinan normaaliaika", "WAST": "Länsi-Afrikan aika", "WAT": "Länsi-Afrikan aika", "WESZ": "Länsi-Euroopan kesäaika", "WEZ": "Länsi-Euroopan normaaliaika", "WFT": "Wallisin ja Futunan aika", "WGST": "Länsi-Grönlannin kesäaika", "WGT": "Länsi-Grönlannin normaaliaika", "WIB": "Länsi-Indonesian aika", "WIT": "Itä-Indonesian aika", "WITA": "Keski-Indonesian aika", "YAKST": "Jakutskin kesäaika", "YAKT": "Jakutskin normaaliaika", "YEKST": "Jekaterinburgin kesäaika", "YEKT": "Jekaterinburgin normaaliaika", "YST": "Yukonin aika", "МСК": "Moskovan normaaliaika", "اقتاۋ": "Aqtaw’n normaaliaika", "اقتاۋ قالاسى": "Aqtaw’n kesäaika", "اقتوبە": "Aqtöben normaaliaika", "اقتوبە قالاسى": "Aqtöben kesäaika", "الماتى": "Almatyn normaaliaika", "الماتى قالاسى": "Almatyn kesäaika", "باتىس قازاق ەلى": "Länsi-Kazakstanin aika", "شىعىش قازاق ەلى": "Itä-Kazakstanin aika", "قازاق ەلى": "Kazakstanin aika", "قىرعىزستان": "Kirgisian aika", "قىزىلوردا": "Qızılordan normaaliaika", "قىزىلوردا قالاسى": "Qızılordan kesäaika", "∅∅∅": "Perun kesäaika"},
	}
}

// Locale returns the current translators string locale
func (fi *fi) Locale() string {
	return fi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fi'
func (fi *fi) PluralsCardinal() []locales.PluralRule {
	return fi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fi'
func (fi *fi) PluralsOrdinal() []locales.PluralRule {
	return fi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fi'
func (fi *fi) PluralsRange() []locales.PluralRule {
	return fi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fi'
func (fi *fi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fi'
func (fi *fi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fi'
func (fi *fi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fi *fi) MonthAbbreviated(month time.Month) string {
	return fi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fi *fi) MonthsAbbreviated() []string {
	return fi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fi *fi) MonthNarrow(month time.Month) string {
	return fi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fi *fi) MonthsNarrow() []string {
	return fi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fi *fi) MonthWide(month time.Month) string {
	return fi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fi *fi) MonthsWide() []string {
	return fi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fi *fi) WeekdayAbbreviated(weekday time.Weekday) string {
	return fi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fi *fi) WeekdaysAbbreviated() []string {
	return fi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fi *fi) WeekdayNarrow(weekday time.Weekday) string {
	return fi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fi *fi) WeekdaysNarrow() []string {
	return fi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fi *fi) WeekdayShort(weekday time.Weekday) string {
	return fi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fi *fi) WeekdaysShort() []string {
	return fi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fi *fi) WeekdayWide(weekday time.Weekday) string {
	return fi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fi *fi) WeekdaysWide() []string {
	return fi.daysWide
}

// Decimal returns the decimal point of number
func (fi *fi) Decimal() string {
	return fi.decimal
}

// Group returns the group of number
func (fi *fi) Group() string {
	return fi.group
}

// Group returns the minus sign of number
func (fi *fi) Minus() string {
	return fi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fi' and handles both Whole and Real numbers based on 'v'
func (fi *fi) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fi.group) - 1; j >= 0; j-- {
					b = append(b, fi.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fi' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fi *fi) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fi.percentSuffix...)

	b = append(b, fi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fi'
func (fi *fi) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fi.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fi.group) - 1; j >= 0; j-- {
					b = append(b, fi.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, fi.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fi'
// in accounting notation.
func (fi *fi) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fi.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fi.group) - 1; j >= 0; j-- {
					b = append(b, fi.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fi.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, fi.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fi'
func (fi *fi) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fi'
func (fi *fi) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fi'
func (fi *fi) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, fi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fi'
func (fi *fi) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0x63, 0x63, 0x63, 0x63, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, fi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fi'
func (fi *fi) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fi'
func (fi *fi) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fi'
func (fi *fi) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fi'
func (fi *fi) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
