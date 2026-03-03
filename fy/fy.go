package fy

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type fy struct {
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
	currencyPositivePrefix string
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

// New returns a new instance of translator for the 'fy' locale
func New() locales.Translator {
	return &fy{
		locale:                 "fy",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "C$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJ$", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SI$", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mrt", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Jannewaris", "Febrewaris", "Maart", "April", "Maaie", "Juny", "July", "Augustus", "Septimber", "Oktober", "Novimber", "Desimber"},
		daysAbbreviated:        []string{"si", "mo", "ti", "wo", "to", "fr", "so"},
		daysWide:               []string{"snein", "moandei", "tiisdei", "woansdei", "tongersdei", "freed", "sneon"},
		timezones:              map[string]string{"ACDT": "Midden-Australyske simmertiid", "ACST": "Midden-Australyske standerttiid", "ACT": "Acre-standerttiid", "ACWDT": "Midden-Australyske westelijke simmertiid", "ACWST": "Midden-Australyske westelijke standerttiid", "ADT": "Atlantic-simmertiid", "ADT Arabia": "Arabyske simmertiid", "AEDT": "East-Australyske simmertiid", "AEST": "East-Australyske standerttiid", "AFT": "Afghaanske tiid", "AKDT": "Alaska-simmertiid", "AKST": "Alaska-standerttiid", "AMST": "Amazone-simmertiid", "AMST Armenia": "Armeense simmertiid", "AMT": "Amazone-standerttiid", "AMT Armenia": "Armeense standerttiid", "ANAST": "Anadyr-simmertiid", "ANAT": "Anadyr-standerttiid", "ARST": "Argentynske simmertiid", "ART": "Argentynske standerttiid", "AST": "Atlantic-standerttiid", "AST Arabia": "Arabyske standerttiid", "AWDT": "West-Australyske simmertiid", "AWST": "West-Australyske standerttiid", "AZST": "Azerbeidzjaanske simmertiid", "AZT": "Azerbeidzjaanske standerttiid", "BDT Bangladesh": "Bengalese simmertiid", "BNT": "Bruneise tiid", "BOT": "Boliviaanske tiid", "BRST": "Brazyljaanske simmertiid", "BRT": "Brazyljaanske standerttiid", "BST Bangladesh": "Bengalese standerttiid", "BT": "Bhutaanske tiid", "CAST": "CAST", "CAT": "Sintraal-Afrikaanske tiid", "CCT": "Kokoseilânske tiid", "CDT": "Central-simmertiid", "CHADT": "Chatham simmertiid", "CHAST": "Chatham standerttiid", "CHUT": "Chuukse tiid", "CKT": "Cookeilânse standerttiid", "CKT DST": "Cookeilânse halve simmertiid", "CLST": "Sileenske simmertiid", "CLT": "Sileenske standerttiid", "COST": "Kolombiaanske simmertiid", "COT": "Kolombiaanske standerttiid", "CST": "Central-standerttiid", "CST China": "Sineeske standerttiid", "CST China DST": "Sineeske simmertiid", "CVST": "Kaapverdyske simmertiid", "CVT": "Kaapverdyske standerttiid", "CXT": "Krysteilânske tiid", "ChST": "Chamorro-tiid", "ChST NMI": "Noardlike Mariaanske tiid", "CuDT": "Kubaanske simmertiid", "CuST": "Kubaanske standerttiid", "DAVT": "Davis tiid", "DDUT": "Dumont-d’Urville tiid", "EASST": "Peaskeeilânske simmertiid", "EAST": "Peaskeeilânske standerttiid", "EAT": "East-Afrikaanske tiid", "ECT": "Ecuadoraanske tiid", "EDT": "Eastern-simmertiid", "EGDT": "East-Groenlânske simmertiid", "EGST": "East-Groenlânske standerttiid", "EST": "Eastern-standerttiid", "FEET": "Fierder-eastlik-Europeeske tiid", "FJT": "Fijyske standerttiid", "FJT Summer": "Fijyske simmertiid", "FKST": "Falklâneilânske simmertiid", "FKT": "Falklâneilânske standerttiid", "FNST": "Fernando de Noronha-simmertiid", "FNT": "Fernando de Noronha-standerttiid", "GALT": "Galapagoseilânske tiid", "GAMT": "Gambiereilânske tiid", "GEST": "Georgyske simmertiid", "GET": "Georgyske standerttiid", "GFT": "Frâns-Guyaanske tiid", "GIT": "Gilberteilânske tiid", "GMT": "Greenwich Mean Time", "GNSST": "GNSST", "GNST": "GNST", "GST": "Sûd-Georgyske tiid", "GST Guam": "Guamese standerttiid", "GYT": "Guyaanske tiid", "HADT": "Hawaii-Aleoetyske simmertiid", "HAST": "Hawaii-Aleoetyske standerttiid", "HKST": "Hongkongse simmertiid", "HKT": "Hongkongse standerttiid", "HOVST": "Hovd simmertiid", "HOVT": "Hovd standerttiid", "ICT": "Yndochinese tiid", "IDT": "Israëlyske simmertiid", "IOT": "Yndyske Oceaan-tiid", "IRKST": "Irkoetsk-simmertiid", "IRKT": "Irkoetsk-standerttiid", "IRST": "Iraanske standerttiid", "IRST DST": "Iraanske simmertiid", "IST": "Yndiaaske tiid", "IST Israel": "Israëlyske standerttiid", "JDT": "Japanske simmertiid", "JST": "Japanske standerttiid", "KOST": "Kosraese tiid", "KRAST": "Krasnojarsk-simmertiid", "KRAT": "Krasnojarsk-standerttiid", "KST": "Koreaanske standerttiid", "KST DST": "Koreaanske simmertiid", "LHDT": "Lord Howe-eilânske simmertiid", "LHST": "Lord Howe-eilânske standerttiid", "LINT": "Line-eilânske tiid", "MAGST": "Magadan-simmertiid", "MAGT": "Magadan-standerttiid", "MART": "Marquesaseilânske tiid", "MAWT": "Mawson tiid", "MDT": "Mountain-simmertiid", "MESZ": "Midden-Europeeske simmertiid", "MEZ": "Midden-Europeeske standerttiid", "MHT": "Marshalleilânske tiid", "MMT": "Myanmarese tiid", "MSD": "Moskou-simmertiid", "MST": "Mountain-standerttiid", "MUST": "Mauritiaanske simmertiid", "MUT": "Mauritiaanske standerttiid", "MVT": "Maldivyske tiid", "MYT": "Maleisyske tiid", "NCT": "Nij-Kaledonyske standerttiid", "NDT": "Newfoundlânske-simmertiid", "NDT New Caledonia": "Nij-Kaledonyske simmertiid", "NFDT": "Norfolkeilânske simmertiid", "NFT": "Norfolkeilânske standerttiid", "NOVST": "Novosibirsk-simmertiid", "NOVT": "Novosibirsk-standerttiid", "NPT": "Nepalese tiid", "NRT": "Nauruaanske tiid", "NST": "Newfoundlânske-standerttiid", "NUT": "Niuese tiid", "NZDT": "Nij-Seelânske simmertiid", "NZST": "Nij-Seelânske standerttiid", "OESZ": "East-Europeeske simmertiid", "OEZ": "East-Europeeske standerttiid", "OMSST": "Omsk-simmertiid", "OMST": "Omsk-standerttiid", "PDT": "Pasifik-simmertiid", "PDTM": "Meksikaanske Pacific-simmerttiid", "PETDT": "Petropavlovsk-Kamtsjatski-simmertiid", "PETST": "Petropavlovsk-Kamtsjatski-standerttiid", "PGT": "Papoea-Nij-Guineeske tiid", "PHOT": "Phoenixeilânske tiid", "PKT": "Pakistaanske standerttiid", "PKT DST": "Pakistaanske simmertiid", "PMDT": "Saint Pierre en Miquelon-simmertiid", "PMST": "Saint Pierre en Miquelon-standerttiid", "PONT": "Pohnpei tiid", "PST": "Pasifik-standerttiid", "PST Philippine": "Filipijnse standerttiid", "PST Philippine DST": "Filipijnse simmertiid", "PST Pitcairn": "Pitcairneillânske tiid", "PSTM": "Meksikaanske Pacific-standerttiid", "PWT": "Belause tiid", "PYST": "Paraguayaanske simmertiid", "PYT": "Paraguayaanske standerttiid", "PYT Korea": "Pyongyang-tiid", "RET": "Réunionse tiid", "ROTT": "Rothera tiid", "SAKST": "Sachalin-simmertiid", "SAKT": "Sachalin-standerttiid", "SAMST": "Samara-simmertiid", "SAMT": "Samara-standerttiid", "SAST": "Sûd-Afrikaanske tiid", "SBT": "Salomonseilânske tiid", "SCT": "Seychelse tiid", "SGT": "Singaporese standerttiid", "SLST": "Lanka-tiid", "SRT": "Surinaamske tiid", "SST Samoa": "Samoaanske standerttiid", "SST Samoa Apia": "Apia-standerttiid", "SST Samoa Apia DST": "Apia-simmertiid", "SST Samoa DST": "Samoaanske simmertiid", "SYOT": "Syowa tiid", "TAAF": "Frânske Súdlike en Antarctyske tiid", "TAHT": "Tahitiaanske tiid", "TJT": "Tadzjiekse tiid", "TKT": "Tokelau-eilânske tiid", "TLT": "East-Timorese tiid", "TMST": "Turkmeense simmertiid", "TMT": "Turkmeense standerttiid", "TOST": "Tongaanske simmertiid", "TOT": "Tongaanske standerttiid", "TVT": "Tuvaluaanske tiid", "TWT": "Taipei standerttiid", "TWT DST": "Taipei simmertiid", "ULAST": "Ulaanbaatar simmertiid", "ULAT": "Ulaanbaatar standerttiid", "UYST": "Uruguayaanske simmertiid", "UYT": "Uruguayaanske standerttiid", "UZT": "Oezbeekse standerttiid", "UZT DST": "Oezbeekse simmertiid", "VET": "Fenezolaanske tiid", "VLAST": "Vladivostok-simmertiid", "VLAT": "Vladivostok-standerttiid", "VOLST": "Wolgograd-simmertiid", "VOLT": "Wolgograd-standerttiid", "VOST": "Vostok tiid", "VUT": "Vanuatuaanske standerttiid", "VUT DST": "Vanuatuaanske simmertiid", "WAKT": "Wake-eilânske tiid", "WARST": "West-Argentynske simmertiid", "WART": "West-Argentynske standerttiid", "WAST": "West-Afrikaanske tiid", "WAT": "West-Afrikaanske tiid", "WESZ": "West-Europeeske simmertiid", "WEZ": "West-Europeeske standerttiid", "WFT": "Wallis en Futunase tiid", "WGST": "West-Groenlânske simmertiid", "WGT": "West-Groenlânske standerttiid", "WIB": "West-Yndonezyske tiid", "WIT": "East-Yndonezyske tiid", "WITA": "Sintraal-Yndonezyske tiid", "YAKST": "Jakoetsk-simmertiid", "YAKT": "Jakoetsk-standerttiid", "YEKST": "Jekaterinenburg-simmertiid", "YEKT": "Jekaterinenburg-standerttiid", "YST": "Yukon-tiid", "МСК": "Moskou-standerttiid", "اقتاۋ": "Aqtau-standerttiid", "اقتاۋ قالاسى": "Aqtau-simmertiid", "اقتوبە": "Aqtöbe-standerttiid", "اقتوبە قالاسى": "Aqtöbe-simmertiid", "الماتى": "Alma-Ata-standerttiid", "الماتى قالاسى": "Alma-Ata-simmertiid", "باتىس قازاق ەلى": "West-Kazachse tiid", "شىعىش قازاق ەلى": "East-Kazachse tiid", "قازاق ەلى": "Kazakhstan-tiid", "قىرعىزستان": "Kirgizyske tiid", "قىزىلوردا": "Qyzylorda-standerttiid", "قىزىلوردا قالاسى": "Qyzylorda-simmertiid", "∅∅∅": "Peruaanske simmertiid"},
	}
}

// Locale returns the current translators string locale
func (fy *fy) Locale() string {
	return fy.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fy'
func (fy *fy) PluralsCardinal() []locales.PluralRule {
	return fy.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fy'
func (fy *fy) PluralsOrdinal() []locales.PluralRule {
	return fy.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fy'
func (fy *fy) PluralsRange() []locales.PluralRule {
	return fy.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fy'
func (fy *fy) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fy'
func (fy *fy) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fy'
func (fy *fy) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fy *fy) MonthAbbreviated(month time.Month) string {
	if len(fy.monthsAbbreviated) == 0 {
		return ""
	}
	return fy.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fy *fy) MonthsAbbreviated() []string {
	return fy.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fy *fy) MonthNarrow(month time.Month) string {
	if len(fy.monthsNarrow) == 0 {
		return ""
	}
	return fy.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fy *fy) MonthsNarrow() []string {
	return fy.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fy *fy) MonthWide(month time.Month) string {
	if len(fy.monthsWide) == 0 {
		return ""
	}
	return fy.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fy *fy) MonthsWide() []string {
	return fy.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fy *fy) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(fy.daysAbbreviated) == 0 {
		return ""
	}
	return fy.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fy *fy) WeekdaysAbbreviated() []string {
	return fy.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fy *fy) WeekdayNarrow(weekday time.Weekday) string {
	if len(fy.daysNarrow) == 0 {
		return ""
	}
	return fy.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fy *fy) WeekdaysNarrow() []string {
	return fy.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fy *fy) WeekdayShort(weekday time.Weekday) string {
	if len(fy.daysShort) == 0 {
		return ""
	}
	return fy.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fy *fy) WeekdaysShort() []string {
	return fy.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fy *fy) WeekdayWide(weekday time.Weekday) string {
	if len(fy.daysWide) == 0 {
		return ""
	}
	return fy.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fy *fy) WeekdaysWide() []string {
	return fy.daysWide
}

// Decimal returns the decimal point of number
func (fy *fy) Decimal() string {
	return fy.decimal
}

// Group returns the group of number
func (fy *fy) Group() string {
	return fy.group
}

// Group returns the minus sign of number
func (fy *fy) Minus() string {
	return fy.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fy' and handles both Whole and Real numbers based on 'v'
func (fy *fy) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fy.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fy' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fy *fy) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fy.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, fy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fy.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fy'
func (fy *fy) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fy.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fy.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(fy.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, fy.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, fy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fy'
// in accounting notation.
func (fy *fy) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fy.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, fy.group[0])
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

		for j := len(fy.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, fy.currencyNegativePrefix[j])
		}

	} else {

		for j := len(fy.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, fy.currencyPositivePrefix[j])
		}

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
			b = append(b, fy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fy.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fy'
func (fy *fy) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fy'
func (fy *fy) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fy.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fy'
func (fy *fy) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fy'
func (fy *fy) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, fy.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fy'
func (fy *fy) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fy'
func (fy *fy) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fy'
func (fy *fy) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fy'
func (fy *fy) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fy.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
