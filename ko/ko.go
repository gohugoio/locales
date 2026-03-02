package ko

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ko struct {
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

// New returns a new instance of translator for the 'ko' locale
func New() locales.Translator {
	return &ko{
		locale:                 "ko",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "‏-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsNarrow:           []string{"", "1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
		monthsWide:             []string{"", "1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
		daysAbbreviated:        []string{"일", "월", "화", "수", "목", "금", "토"},
		daysNarrow:             []string{"일", "월", "화", "수", "목", "금", "토"},
		daysWide:               []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"},
		periodsAbbreviated:     []string{"", ""},
		periodsWide:            []string{"오전", "오후"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"기원전", "서기"},
		timezones:              map[string]string{"ACDT": "호주 중부 하계 표준시", "ACST": "아크레 하계 표준시", "ACT": "아크레 표준시", "ACWDT": "호주 중서부 하계 표준시", "ACWST": "호주 중서부 표준시", "ADT": "대서양 하계 표준시", "ADT Arabia": "아라비아 하계 표준시", "AEDT": "호주 동부 하계 표준시", "AEST": "호주 동부 표준시", "AFT": "아프가니스탄 시간", "AKDT": "알래스카 하계 표준시", "AKST": "알래스카 표준시", "AMST": "아마존 하계 표준시", "AMST Armenia": "아르메니아 하계 표준시", "AMT": "아마존 표준시", "AMT Armenia": "아르메니아 표준시", "ANAST": "아나디리 하계 표준시", "ANAT": "아나디리 표준시", "ARST": "아르헨티나 하계 표준시", "ART": "아르헨티나 표준시", "AST": "대서양 표준시", "AST Arabia": "아라비아 표준시", "AWDT": "호주 서부 하계 표준시", "AWST": "호주 서부 표준시", "AZST": "아제르바이잔 하계 표준시", "AZT": "아제르바이잔 표준시", "BDT Bangladesh": "방글라데시 하계 표준시", "BNT": "브루나이 시간", "BOT": "볼리비아 시간", "BRST": "브라질리아 하계 표준시", "BRT": "브라질리아 표준시", "BST Bangladesh": "방글라데시 표준시", "BT": "부탄 시간", "CAST": "케이시 시간", "CAT": "중앙아프리카 시간", "CCT": "코코스 제도 시간", "CDT": "미 중부 하계 표준시", "CHADT": "채텀 하계 표준시", "CHAST": "채텀 표준시", "CHUT": "추크 시간", "CKT": "쿡 제도 표준시", "CKT DST": "쿡 제도 절반 하계 표준시", "CLST": "칠레 하계 표준시", "CLT": "칠레 표준시", "COST": "콜롬비아 하계 표준시", "COT": "콜롬비아 표준시", "CST": "미 중부 표준시", "CST China": "중국 표준시", "CST China DST": "중국 하계 표준시", "CVST": "카보 베르데 하계 표준시", "CVT": "카보 베르데 표준시", "CXT": "크리스마스섬 시간", "ChST": "차모로 시간", "ChST NMI": "북마리아나 제도 표준 시간", "CuDT": "쿠바 하계 표준시", "CuST": "쿠바 표준시", "DAVT": "데이비스 시간", "DDUT": "뒤몽뒤르빌 시간", "EASST": "이스터섬 하계 표준시", "EAST": "이스터섬 표준시", "EAT": "동아프리카 시간", "ECT": "에콰도르 시간", "EDT": "미 동부 하계 표준시", "EGDT": "그린란드 동부 하계 표준시", "EGST": "그린란드 동부 표준시", "EST": "미 동부 표준시", "FEET": "극동유럽 표준시", "FJT": "피지 표준시", "FJT Summer": "피지 하계 표준시", "FKST": "포클랜드 제도 하계 표준시", "FKT": "포클랜드 제도 표준시", "FNST": "페르난도 데 노로냐 하계 표준시", "FNT": "페르난도 데 노로냐 표준시", "GALT": "갈라파고스 시간", "GAMT": "감비에 시간", "GEST": "조지아 하계 표준시", "GET": "조지아 표준시", "GFT": "프랑스령 가이아나 시간", "GIT": "길버트 제도 시간", "GMT": "그리니치 표준시", "GNSST": "GNSST", "GNST": "GNST", "GST": "사우스 조지아 시간", "GST Guam": "괌 표준 시간", "GYT": "가이아나 시간", "HADT": "하와이 알류샨 하계 표준시", "HAST": "하와이 알류샨 표준시", "HKST": "홍콩 하계 표준시", "HKT": "홍콩 표준시", "HOVST": "호브드 하계 표준시", "HOVT": "호브드 표준시", "ICT": "인도차이나 시간", "IDT": "이스라엘 하계 표준시", "IOT": "인도양 시간", "IRKST": "이르쿠츠크 하계 표준시", "IRKT": "이르쿠츠크 표준시", "IRST": "이란 표준시", "IRST DST": "이란 하계 표준시", "IST": "인도 표준시", "IST Israel": "이스라엘 표준시", "JDT": "일본 하계 표준시", "JST": "일본 표준시", "KOST": "코스라에섬 시간", "KRAST": "크라스노야르스크 하계 표준시", "KRAT": "크라스노야르스크 표준시", "KST": "한국 표준시", "KST DST": "한국 하계 표준시", "LHDT": "로드 하우 하계 표준시", "LHST": "로드 하우 표준시", "LINT": "라인 제도 시간", "MAGST": "마가단 하계 표준시", "MAGT": "마가단 표준시", "MART": "마르키즈 제도 시간", "MAWT": "모슨 시간", "MDT": "미 산지 하계 표준시", "MESZ": "중부유럽 하계 표준시", "MEZ": "중부유럽 표준시", "MHT": "마셜 제도 시간", "MMT": "미얀마 시간", "MSD": "모스크바 하계 표준시", "MST": "미 산악 표준시", "MUST": "모리셔스 하계 표준시", "MUT": "모리셔스 표준시", "MVT": "몰디브 시간", "MYT": "말레이시아 시간", "NCT": "뉴칼레도니아 표준시", "NDT": "뉴펀들랜드 하계 표준시", "NDT New Caledonia": "뉴칼레도니아 하계 표준시", "NFDT": "노퍽섬 하계 표준시", "NFT": "노퍽섬 표준시", "NOVST": "노보시비르스크 하계 표준시", "NOVT": "노보시비르스크 표준시", "NPT": "네팔 시간", "NRT": "나우루 시간", "NST": "뉴펀들랜드 표준시", "NUT": "니우에 시간", "NZDT": "뉴질랜드 하계 표준시", "NZST": "뉴질랜드 표준시", "OESZ": "동유럽 하계 표준시", "OEZ": "동유럽 표준시", "OMSST": "옴스크 하계 표준시", "OMST": "옴스크 표준시", "PDT": "미 태평양 하계 표준시", "PDTM": "멕시코 태평양 하계 표준시", "PETDT": "페트로파블롭스크-캄차츠키 하계 표준시", "PETST": "페트로파블롭스크-캄차츠키 표준시", "PGT": "파푸아뉴기니 시간", "PHOT": "피닉스 제도 시간", "PKT": "파키스탄 표준시", "PKT DST": "파키스탄 하계 표준시", "PMDT": "세인트피에르 미클롱 하계 표준시", "PMST": "세인트피에르 미클롱 표준시", "PONT": "포나페 시간", "PST": "미 태평양 표준시", "PST Philippine": "필리핀 표준시", "PST Philippine DST": "필리핀 하계 표준시", "PST Pitcairn": "핏케언 시간", "PSTM": "멕시코 태평양 표준시", "PWT": "팔라우 시간", "PYST": "파라과이 하계 표준시", "PYT": "파라과이 표준시", "PYT Korea": "평양 시간", "RET": "레위니옹 시간", "ROTT": "로데라 시간", "SAKST": "사할린 하계 표준시", "SAKT": "사할린 표준시", "SAMST": "사마라 하계 표준시", "SAMT": "사마라 표준시", "SAST": "남아프리카 시간", "SBT": "솔로몬 제도 시간", "SCT": "세이셸 시간", "SGT": "싱가포르 표준시", "SLST": "랑카 표준 시간", "SRT": "수리남 시간", "SST Samoa": "사모아 표준시", "SST Samoa Apia": "아피아 표준시", "SST Samoa Apia DST": "아피아 하계 표준시", "SST Samoa DST": "사모아 하계 표준시", "SYOT": "쇼와 시간", "TAAF": "프랑스령 남부 식민지 및 남극 시간", "TAHT": "타히티 시간", "TJT": "타지키스탄 시간", "TKT": "토켈라우 시간", "TLT": "동티모르 시간", "TMST": "투르크메니스탄 하계 표준시", "TMT": "투르크메니스탄 표준시", "TOST": "통가 하계 표준시", "TOT": "통가 표준시", "TVT": "투발루 시간", "TWT": "대만 표준시", "TWT DST": "대만 하계 표준시", "ULAST": "울란바토르 하계 표준시", "ULAT": "울란바토르 표준시", "UYST": "우루과이 하계 표준시", "UYT": "우루과이 표준시", "UZT": "우즈베키스탄 표준시", "UZT DST": "우즈베키스탄 하계 표준시", "VET": "베네수엘라 시간", "VLAST": "블라디보스토크 하계 표준시", "VLAT": "블라디보스토크 표준시", "VOLST": "볼고그라드 하계 표준시", "VOLT": "볼고그라드 표준시", "VOST": "보스톡 시간", "VUT": "바누아투 표준시", "VUT DST": "바누아투 하계 표준시", "WAKT": "웨이크섬 시간", "WARST": "아르헨티나 서부 하계 표준시", "WART": "아르헨티나 서부 표준시", "WAST": "서아프리카 시간", "WAT": "서아프리카 시간", "WESZ": "서유럽 하계 표준시", "WEZ": "서유럽 표준시", "WFT": "월리스푸투나 제도 시간", "WGST": "그린란드 서부 하계 표준시", "WGT": "그린란드 서부 표준시", "WIB": "서부 인도네시아 시간", "WIT": "동부 인도네시아 시간", "WITA": "중부 인도네시아 시간", "YAKST": "야쿠츠크 하계 표준시", "YAKT": "야쿠츠크 표준시", "YEKST": "예카테린부르크 하계 표준시", "YEKT": "예카테린부르크 표준시", "YST": "유콘 시간", "МСК": "모스크바 표준시", "اقتاۋ": "악타우 표준 표준시", "اقتاۋ قالاسى": "악타우 하계 표준시", "اقتوبە": "악퇴베 표준 표준시", "اقتوبە قالاسى": "악퇴베 하계 표준시", "الماتى": "알마티 표준 표준시", "الماتى قالاسى": "알마티 하계 표준시", "باتىس قازاق ەلى": "서부 카자흐스탄 시간", "شىعىش قازاق ەلى": "동부 카자흐스탄 시간", "قازاق ەلى": "카자흐스탄 시간", "قىرعىزستان": "키르기스스탄 시간", "قىزىلوردا": "키질로르다 표준 시간", "قىزىلوردا قالاسى": "키질로르다 하계 표준시", "∅∅∅": "페루 하계 표준시"},
	}
}

// Locale returns the current translators string locale
func (ko *ko) Locale() string {
	return ko.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ko'
func (ko *ko) PluralsCardinal() []locales.PluralRule {
	return ko.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ko'
func (ko *ko) PluralsOrdinal() []locales.PluralRule {
	return ko.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ko'
func (ko *ko) PluralsRange() []locales.PluralRule {
	return ko.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ko'
func (ko *ko) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ko'
func (ko *ko) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ko'
func (ko *ko) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ko *ko) MonthAbbreviated(month time.Month) string {
	return ko.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ko *ko) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ko *ko) MonthNarrow(month time.Month) string {
	return ko.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ko *ko) MonthsNarrow() []string {
	return ko.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ko *ko) MonthWide(month time.Month) string {
	return ko.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ko *ko) MonthsWide() []string {
	return ko.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ko *ko) WeekdayAbbreviated(weekday time.Weekday) string {
	return ko.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ko *ko) WeekdaysAbbreviated() []string {
	return ko.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ko *ko) WeekdayNarrow(weekday time.Weekday) string {
	return ko.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ko *ko) WeekdaysNarrow() []string {
	return ko.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ko *ko) WeekdayShort(weekday time.Weekday) string {
	return ko.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ko *ko) WeekdaysShort() []string {
	return ko.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ko *ko) WeekdayWide(weekday time.Weekday) string {
	return ko.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ko *ko) WeekdaysWide() []string {
	return ko.daysWide
}

// Decimal returns the decimal point of number
func (ko *ko) Decimal() string {
	return ko.decimal
}

// Group returns the group of number
func (ko *ko) Group() string {
	return ko.group
}

// Group returns the minus sign of number
func (ko *ko) Minus() string {
	return ko.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ko' and handles both Whole and Real numbers based on 'v'
func (ko *ko) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ko.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ko.minus) - 1; j >= 0; j-- {
			b = append(b, ko.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ko' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ko *ko) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ko.minus) - 1; j >= 0; j-- {
			b = append(b, ko.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ko.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ko'
func (ko *ko) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ko.currencies[currency]
	l := len(s) + len(symbol) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ko.group[0])
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
		for j := len(ko.minus) - 1; j >= 0; j-- {
			b = append(b, ko.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ko.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ko'
// in accounting notation.
func (ko *ko) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ko.currencies[currency]
	l := len(s) + len(symbol) + 7 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ko.group[0])
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

		b = append(b, ko.currencyNegativePrefix[0])

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
			b = append(b, ko.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ko.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ko'
func (ko *ko) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ko'
func (ko *ko) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ko'
func (ko *ko) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xeb, 0x85, 0x84, 0x20}...)
	b = append(b, ko.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xec, 0x9d, 0xbc}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ko'
func (ko *ko) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xeb, 0x85, 0x84, 0x20}...)
	b = append(b, ko.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xec, 0x9d, 0xbc, 0x20}...)
	b = append(b, ko.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ko'
func (ko *ko) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ko.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ko'
func (ko *ko) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ko.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ko.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ko'
func (ko *ko) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0xec, 0x8b, 0x9c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xeb, 0xb6, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xec, 0xb4, 0x88, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ko'
func (ko *ko) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0xec, 0x8b, 0x9c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xeb, 0xb6, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xec, 0xb4, 0x88, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ko.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
