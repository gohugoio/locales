package ps

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ps struct {
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

// New returns a new instance of translator for the 'ps' locale
func New() locales.Translator {
	return &ps{
		locale:                 "ps",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		monthsNarrow:           []string{"", "ج", "ف", "م", "ا", "م", "ج", "ج", "ا", "س", "ا", "ن", "د"},
		monthsWide:             []string{"", "جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysWide:               []string{"يونۍ", "دونۍ", "درېنۍ", "څلرنۍ", "پينځنۍ", "جمعه", "اونۍ"},
		timezones:              map[string]string{"ACDT": "آسترالوي مرکزي د ورځې روښانه وخت", "ACST": "ACST", "ACT": "ACT", "ACWDT": "آسترالوي مرکزي لوېديځ د ورځې روښانه وخت", "ACWST": "آسترالوي مرکزي لوېديځ معياري وخت", "ADT": "اتلانتیک د رڼا ورځې وخت", "ADT Arabia": "عربي د ورځې روښانه وخت", "AEDT": "آسترالوي ختيځ د ورځې روښانه وخت", "AEST": "آسترالوي ختيځ معياري وخت", "AFT": "افغانستان وخت", "AKDT": "الاسکا د ورځې روښانه وخت", "AKST": "الاسکا معياري وخت", "AMST": "ایمیزون اوړي وخت", "AMST Armenia": "ارمنستان اوړي وخت", "AMT": "ایمیزون معیاری وخت", "AMT Armenia": "ارمنستان معياري وخت", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ارجنټاین اوړي وخت", "ART": "ارجنټاین معیاری وخت", "AST": "اتلانتیک معياري وخت", "AST Arabia": "عربي معیاري وخت", "AWDT": "آسترالوي لوېديځ د ورځې روښانه وخت", "AWST": "آسترالوي لوېديځ معياري وخت", "AZST": "اذرباییجان اوړي وخت", "AZT": "آذربايجان معياري وخت", "BDT Bangladesh": "بنګله ديش اوړي وخت", "BNT": "برونايي دارالسلام وخت", "BOT": "بولیویا وخت", "BRST": "برسلیا اوړي وخت", "BRT": "برسلیا معیاری وخت", "BST Bangladesh": "بنګلادیش معیاري وخت", "BT": "بهوټان وخت", "CAST": "CAST", "CAT": "منځنی افريقا وخت", "CCT": "کوکوز ټاپوګانو وخت", "CDT": "مرکزي د ورځې روښانه وخت", "CHADT": "چاتام د ورځې روښانه وخت", "CHAST": "چاتام معياري وخت", "CHUT": "چوک وخت", "CKT": "کوک ټاپوګانو معياري وخت", "CKT DST": "کوک ټاپوګانو نيم اوړي وخت", "CLST": "چلی اوړي وخت", "CLT": "چلی معیاری وخت", "COST": "کولمبیا اوړي وخت", "COT": "کولمبیا معیاری وخت", "CST": "مرکزي معياري وخت", "CST China": "چین معیاري وخت", "CST China DST": "د چين د رڼا ورځې وخت", "CVST": "کیپ وردډ سمر وخت", "CVT": "کیپ وردډ معياري وخت", "CXT": "کريسمس ټاپو وخت", "ChST": "چمارو معياري وخت", "ChST NMI": "ChST NMI", "CuDT": "کیوبا د رڼا ورځې وخت", "CuST": "کیوبا معياري وخت", "DAVT": "ډيوس وخت", "DDUT": "ډومونټ ډي ارول", "EASST": "ايستر ټاپو اوړي وخت", "EAST": "ايستر ټاپو معياري وخت", "EAT": "ختيځ افريقا وخت", "ECT": "د اکوادور وخت", "EDT": "ختيځ د رڼا ورځې وخت", "EGDT": "د ختیځ ګرینلینډ اوړي وخت", "EGST": "د ختیځ ګرینلینډ معياري وخت", "EST": "ختيځ معياري وخت", "FEET": "لرې ختيځ اروپايي وخت", "FJT": "فجی معياري وخت", "FJT Summer": "فجي د اوړي وخت", "FKST": "د فوکلنډ ټاپو اوړي وخت", "FKT": "د فوکلنډ ټاپو معیاری وخت", "FNST": "فرنانڈو دي نورونھا اوړي وخت", "FNT": "فرنانڈو دي نورونها معیاری وخت", "GALT": "ګالپګوس وخت", "GAMT": "ګيمبير وخت", "GEST": "د جورجيا د اوړي وخت", "GET": "جورجیا معیاري وخت", "GFT": "د فرانسوي ګانا وخت", "GIT": "جلبرټ ټاپوګانو وخت", "GMT": "ګرينويچ معياري وخت", "GNSST": "GNSST", "GNST": "GNST", "GST": "د سویل جورجیا وخت", "GST Guam": "GST Guam", "GYT": "د ګوانانا وخت", "HADT": "هوایی الیوتین رڼا ورځې وخت", "HAST": "هوایی الیوتین معیاری وخت", "HKST": "هانګ کانګ اوړي وخت", "HKT": "هانګ کانګ معياري وخت", "HOVST": "هاوډ اوړي وخت", "HOVT": "هاوډ معیاری وخت", "ICT": "انډوچاینه وخت", "IDT": "اسراييل د ورځې روښانه وخت", "IOT": "د هند سمندر وخت", "IRKST": "ارکوټسک اوړي وخت", "IRKT": "ارکوټسک معياري وخت", "IRST": "ایران معياري وخت", "IRST DST": "ايران د ورځې روښانه وخت", "IST": "هند معیاري وخت", "IST Israel": "اسراییل معياري وخت", "JDT": "جاپان د ورځې روښانه وخت", "JST": "جاپان معياري وخت", "KOST": "کوسراي وخت", "KRAST": "کريسنويارسک اوړي وخت", "KRAT": "کريسنويارسک معياري وخت", "KST": "کوريايي معياري وخت", "KST DST": "کوريايي د ورځې روښانه وخت", "LHDT": "لارډ هوي د ورځې روښانه وخت", "LHST": "لارډ هوي معياري وخت", "LINT": "لاين ټاپوګانو وخت", "MAGST": "ميګډان اوړي وخت", "MAGT": "ميګډان معياري وخت", "MART": "مارکسس وخت", "MAWT": "ماوسن وخت", "MDT": "MDT", "MESZ": "مرکزي اروپايياوړي وخت", "MEZ": "د مرکزي اروپا معیاري وخت", "MHT": "مارشل ټاپوګانو وخت", "MMT": "میانمار وخت", "MSD": "ماسکو سمر وخت", "MST": "MST", "MUST": "ماريشيس د اوړي وخت", "MUT": "ماریشیس معياري وخت", "MVT": "مالديپ وخت", "MYT": "ملائیشیا وخت", "NCT": "نيو کالیډونیا معياري وخت", "NDT": "نيو فاونډلېنډ د ورځې روښانه وخت", "NDT New Caledonia": "نيو کايډونيا اوړي وخت", "NFDT": "د نورفکاس ټاپو اوړي وخت", "NFT": "د نورفکاس ټاپو معياري وخت", "NOVST": "نووسيبرسک اوړي وخت", "NOVT": "نووسيبرسک معياري وخت", "NPT": "نیپال وخت", "NRT": "ناورو وخت", "NST": "د نوي فیلډلینډ معیاری وخت", "NUT": "نییو وخت", "NZDT": "نيوزي لېنډ د ورځې روښانه وخت", "NZST": "نيوزي لېنډ معياري وخت", "OESZ": "ختيځ اروپايي اوړي وخت", "OEZ": "ختيځ اروپايي معياري وخت", "OMSST": "اومسک اوړي وخت", "OMST": "اومسک معياري وخت", "PDT": "پیسفک د رڼا ورځې وخت", "PDTM": "مکسیکن پیسفک رڼا ورځې وخت", "PETDT": "PETDT", "PETST": "PETST", "PGT": "پاپوا نیو ګنی وخت", "PHOT": "د فینکس ټاپو وخت", "PKT": "پاکستان معیاري وخت", "PKT DST": "پاکستان اوړي وخت", "PMDT": "سینټ پییرا و ميکلين رڼا ورځې وخت", "PMST": "سینټ پییرا و ميکلين معیاری وخت", "PONT": "پونيپ وخت", "PST": "د پیسفک معياري وخت", "PST Philippine": "فلپاين معياري وخت", "PST Philippine DST": "فلپاين اوړي وخت", "PST Pitcairn": "پیټ کارین وخت", "PSTM": "مکسیکن پیسفک معیاری وخت", "PWT": "پالاو وخت", "PYST": "پاراګوای اوړي وخت", "PYT": "پیراګوای معياري وخت", "PYT Korea": "پيانګ يانګ وخت", "RET": "ري يونين وخت", "ROTT": "رودرا وخت", "SAKST": "سخلين اوړي وخت", "SAKT": "سخلین معياري وخت", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "جنوبي افريقا معياري وخت", "SBT": "سلیمان ټاپوګانو وخت", "SCT": "سیچیلس وخت", "SGT": "سنګاپور معیاري وخت", "SLST": "SLST", "SRT": "سورینام وخت", "SST Samoa": "سموا معياري وخت", "SST Samoa Apia": "اپیا معياري وخت", "SST Samoa Apia DST": "اپيا د ورځې روښانه وخت", "SST Samoa DST": "سموا د ورځې روښانه وخت", "SYOT": "سیوا وخت", "TAAF": "د فرانسې سویل او انټارټيک وخت", "TAHT": "ټهيټي وخت", "TJT": "تاجکستان وخت", "TKT": "توکیلاو وخت", "TLT": "ختيځ تيمور وخت", "TMST": "ترکمنستان اوړي وخت", "TMT": "ترکمنستان معياري وخت", "TOST": "ټونګا اوړي وخت", "TOT": "د ټونګ معياري وخت", "TVT": "تووالو وخت", "TWT": "تايپي معياري وخت", "TWT DST": "تايپي د ورځې روښانه وخت", "ULAST": "اولان باټر اوړي وخت", "ULAT": "اولان باټر معیاري وخت", "UYST": "یوروګوای اوړي وخت", "UYT": "یوروګوای معياري وخت", "UZT": "ازبکستان معياري وخت", "UZT DST": "ازبکستان اوړي وخت", "VET": "وینزویلا وخت", "VLAST": "ولاديوستاک اوړي وخت", "VLAT": "ولاديوستاک معياري وخت", "VOLST": "والګوګراد اوړي وخت", "VOLT": "والګوګراد معياري وخت", "VOST": "واستوک وخت", "VUT": "ونواتو معياري وخت", "VUT DST": "ونواتو اوړي وخت", "WAKT": "ويک تاپو وخت", "WARST": "لوېديځ ارجنټاين اوړي وخت", "WART": "لوېديځ ارجنټاين معياري وخت", "WAST": "لوېديځ افريقا وخت", "WAT": "لوېديځ افريقا وخت", "WESZ": "لوېديځ اروپايي اوړي وخت", "WEZ": "لوېديځ اروپايي معياري وخت", "WFT": "والس او فوتونا وخت", "WGST": "لویدیځ ګرینلینډ اوړي وخت", "WGT": "لویدیځ ګرینلینډ معياري وخت", "WIB": "لویدیځ اندونیزیا وخت", "WIT": "اندونیزیا وخت", "WITA": "مرکزي ادونيزيا وخت", "YAKST": "ياکوټسک د اوړي وخت", "YAKT": "ياکوټسک معياري وخت", "YEKST": "د ياکټرنبرګ د اوړي وخت", "YEKT": "د ياکيټرنبرګ معياري وخت", "YST": "د یوکون وخت", "МСК": "ماسکو معياري وخت", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "لویدیځ قزاقستان وخت", "شىعىش قازاق ەلى": "ختيځ قازقستان وخت", "قازاق ەلى": "قزاقستان وخت", "قىرعىزستان": "کرغیزستان وخت", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ايزورس اوړي وخت"},
	}
}

// Locale returns the current translators string locale
func (ps *ps) Locale() string {
	return ps.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ps'
func (ps *ps) PluralsCardinal() []locales.PluralRule {
	return ps.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ps'
func (ps *ps) PluralsOrdinal() []locales.PluralRule {
	return ps.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ps'
func (ps *ps) PluralsRange() []locales.PluralRule {
	return ps.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ps'
func (ps *ps) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ps'
func (ps *ps) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ps'
func (ps *ps) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ps.CardinalPluralRule(num1, v1)
	end := ps.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ps *ps) MonthAbbreviated(month time.Month) string {
	if len(ps.monthsAbbreviated) == 0 {
		return ""
	}
	return ps.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ps *ps) MonthsAbbreviated() []string {
	return ps.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ps *ps) MonthNarrow(month time.Month) string {
	if len(ps.monthsNarrow) == 0 {
		return ""
	}
	return ps.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ps *ps) MonthsNarrow() []string {
	return ps.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ps *ps) MonthWide(month time.Month) string {
	if len(ps.monthsWide) == 0 {
		return ""
	}
	return ps.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ps *ps) MonthsWide() []string {
	return ps.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ps *ps) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ps.daysAbbreviated) == 0 {
		return ""
	}
	return ps.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ps *ps) WeekdaysAbbreviated() []string {
	return ps.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ps *ps) WeekdayNarrow(weekday time.Weekday) string {
	if len(ps.daysNarrow) == 0 {
		return ""
	}
	return ps.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ps *ps) WeekdaysNarrow() []string {
	return ps.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ps *ps) WeekdayShort(weekday time.Weekday) string {
	if len(ps.daysShort) == 0 {
		return ""
	}
	return ps.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ps *ps) WeekdaysShort() []string {
	return ps.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ps *ps) WeekdayWide(weekday time.Weekday) string {
	if len(ps.daysWide) == 0 {
		return ""
	}
	return ps.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ps *ps) WeekdaysWide() []string {
	return ps.daysWide
}

// Decimal returns the decimal point of number
func (ps *ps) Decimal() string {
	return ps.decimal
}

// Group returns the group of number
func (ps *ps) Group() string {
	return ps.group
}

// Group returns the minus sign of number
func (ps *ps) Minus() string {
	return ps.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ps' and handles both Whole and Real numbers based on 'v'
func (ps *ps) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ps.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ps.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ps' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ps *ps) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ps.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ps.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ps'
func (ps *ps) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ps.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ps.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(ps.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ps.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, ps.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ps.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ps'
// in accounting notation.
func (ps *ps) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ps.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ps.group[0])
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

		for j := len(ps.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ps.currencyNegativePrefix[j])
		}

	} else {

		for j := len(ps.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ps.currencyPositivePrefix[j])
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
			b = append(b, ps.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ps.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ps'
func (ps *ps) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ps'
func (ps *ps) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ps.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ps'
func (ps *ps) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ps.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ps'
func (ps *ps) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ps.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0xd8, 0xaf, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd8, 0xaf, 0x20}...)
	b = append(b, ps.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ps'
func (ps *ps) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ps'
func (ps *ps) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ps'
func (ps *ps) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ps.timeSeparator...)

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

// FmtTimeFull returns the full time representation of 't' for 'ps'
func (ps *ps) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := ps.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
