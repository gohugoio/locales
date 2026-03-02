package ja

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ja struct {
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

// New returns a new instance of translator for the 'ja' locale
func New() locales.Translator {
	return &ja{
		locale:                 "ja",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "元", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "￥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "レイ", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "Cg", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsWide:             []string{"", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		daysAbbreviated:        []string{"日", "月", "火", "水", "木", "金", "土"},
		daysNarrow:             []string{"日", "月", "火", "水", "木", "金", "土"},
		daysWide:               []string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},
		timezones:              map[string]string{"ACDT": "オーストラリア中部夏時間", "ACST": "オーストラリア中部標準時", "ACT": "アクレ標準時", "ACWDT": "オーストラリア中西部夏時間", "ACWST": "オーストラリア中西部標準時", "ADT": "大西洋夏時間", "ADT Arabia": "アラビア夏時間", "AEDT": "オーストラリア東部夏時間", "AEST": "オーストラリア東部標準時", "AFT": "アフガニスタン時間", "AKDT": "アラスカ夏時間", "AKST": "アラスカ標準時", "AMST": "アマゾン夏時間", "AMST Armenia": "アルメニア夏時間", "AMT": "アマゾン標準時", "AMT Armenia": "アルメニア標準時", "ANAST": "アナディリ夏時間", "ANAT": "アナディリ標準時", "ARST": "アルゼンチン夏時間", "ART": "アルゼンチン標準時", "AST": "大西洋標準時", "AST Arabia": "アラビア標準時", "AWDT": "オーストラリア西部夏時間", "AWST": "オーストラリア西部標準時", "AZST": "アゼルバイジャン夏時間", "AZT": "アゼルバイジャン標準時", "BDT Bangladesh": "バングラデシュ夏時間", "BNT": "ブルネイ時間", "BOT": "ボリビア時間", "BRST": "ブラジリア夏時間", "BRT": "ブラジリア標準時", "BST Bangladesh": "バングラデシュ標準時", "BT": "ブータン時間", "CAST": "ケイシー基地時間", "CAT": "中央アフリカ時間", "CCT": "ココス諸島時間", "CDT": "米国中部夏時間", "CHADT": "チャタム夏時間", "CHAST": "チャタム標準時", "CHUT": "チューク時間", "CKT": "クック諸島標準時", "CKT DST": "クック諸島夏時間", "CLST": "チリ夏時間", "CLT": "チリ標準時", "COST": "コロンビア夏時間", "COT": "コロンビア標準時", "CST": "米国中部標準時", "CST China": "中国標準時", "CST China DST": "中国夏時間", "CVST": "カーボベルデ夏時間", "CVT": "カーボベルデ標準時", "CXT": "クリスマス島時間", "ChST": "チャモロ時間", "ChST NMI": "北マリアナ諸島時間", "CuDT": "キューバ夏時間", "CuST": "キューバ標準時", "DAVT": "デービス基地時間", "DDUT": "デュモン・デュルヴィル基地時間", "EASST": "イースター島夏時間", "EAST": "イースター島標準時", "EAT": "東アフリカ時間", "ECT": "エクアドル時間", "EDT": "米国東部夏時間", "EGDT": "グリーンランド東部夏時間", "EGST": "グリーンランド東部標準時", "EST": "米国東部標準時", "FEET": "極東ヨーロッパ時間", "FJT": "フィジー標準時", "FJT Summer": "フィジー夏時間", "FKST": "フォークランド諸島夏時間", "FKT": "フォークランド諸島標準時", "FNST": "フェルナンド・デ・ノローニャ夏時間", "FNT": "フェルナンド・デ・ノローニャ標準時", "GALT": "ガラパゴス時間", "GAMT": "ガンビエ諸島時間", "GEST": "ジョージア夏時間", "GET": "ジョージア標準時", "GFT": "仏領ギアナ時間", "GIT": "ギルバート諸島時間", "GMT": "グリニッジ標準時", "GNSST": "GNSST", "GNST": "GNST", "GST": "湾岸標準時", "GST Guam": "グアム時間", "GYT": "ガイアナ時間", "HADT": "ハワイ・アリューシャン夏時間", "HAST": "ハワイ・アリューシャン標準時", "HKST": "香港夏時間", "HKT": "香港標準時", "HOVST": "ホブド夏時間", "HOVT": "ホブド標準時", "ICT": "インドシナ時間", "IDT": "イスラエル夏時間", "IOT": "インド洋時間", "IRKST": "イルクーツク夏時間", "IRKT": "イルクーツク標準時", "IRST": "イラン標準時", "IRST DST": "イラン夏時間", "IST": "インド標準時", "IST Israel": "イスラエル標準時", "JDT": "日本夏時間", "JST": "日本標準時", "KOST": "コスラエ時間", "KRAST": "クラスノヤルスク夏時間", "KRAT": "クラスノヤルスク標準時", "KST": "韓国標準時", "KST DST": "韓国夏時間", "LHDT": "ロードハウ夏時間", "LHST": "ロードハウ標準時", "LINT": "ライン諸島時間", "MAGST": "マガダン夏時間", "MAGT": "マガダン標準時", "MART": "マルキーズ時間", "MAWT": "モーソン基地時間", "MDT": "米国山岳夏時間", "MESZ": "中央ヨーロッパ夏時間", "MEZ": "中央ヨーロッパ標準時", "MHT": "マーシャル諸島時間", "MMT": "ミャンマー時間", "MSD": "モスクワ夏時間", "MST": "米国山岳標準時", "MUST": "モーリシャス夏時間", "MUT": "モーリシャス標準時", "MVT": "モルディブ時間", "MYT": "マレーシア時間", "NCT": "ニューカレドニア標準時", "NDT": "ニューファンドランド夏時間", "NDT New Caledonia": "ニューカレドニア夏時間", "NFDT": "ノーフォーク島夏時間", "NFT": "ノーフォーク島標準時", "NOVST": "ノヴォシビルスク夏時間", "NOVT": "ノヴォシビルスク標準時", "NPT": "ネパール時間", "NRT": "ナウル時間", "NST": "ニューファンドランド標準時", "NUT": "ニウエ時間", "NZDT": "ニュージーランド夏時間", "NZST": "ニュージーランド標準時", "OESZ": "東ヨーロッパ夏時間", "OEZ": "東ヨーロッパ標準時", "OMSST": "オムスク夏時間", "OMST": "オムスク標準時", "PDT": "米国太平洋夏時間", "PDTM": "メキシコ太平洋夏時間", "PETDT": "ペトロパブロフスク・カムチャツキー夏時間", "PETST": "ペトロパブロフスク・カムチャツキー標準時", "PGT": "パプアニューギニア時間", "PHOT": "フェニックス諸島時間", "PKT": "パキスタン標準時", "PKT DST": "パキスタン夏時間", "PMDT": "サンピエール島・ミクロン島夏時間", "PMST": "サンピエール島・ミクロン島標準時", "PONT": "ポンペイ時間", "PST": "米国太平洋標準時", "PST Philippine": "フィリピン標準時", "PST Philippine DST": "フィリピン夏時間", "PST Pitcairn": "ピトケアン時間", "PSTM": "メキシコ太平洋標準時", "PWT": "パラオ時間", "PYST": "パラグアイ夏時間", "PYT": "パラグアイ標準時", "PYT Korea": "北朝鮮時間", "RET": "レユニオン時間", "ROTT": "ロゼラ基地時間", "SAKST": "サハリン夏時間", "SAKT": "サハリン標準時", "SAMST": "サマラ夏時間", "SAMT": "サマラ標準時", "SAST": "南アフリカ標準時", "SBT": "ソロモン諸島時間", "SCT": "セーシェル時間", "SGT": "シンガポール標準時", "SLST": "ランカ時間", "SRT": "スリナム時間", "SST Samoa": "米領サモア標準時", "SST Samoa Apia": "サモア標準時", "SST Samoa Apia DST": "サモア夏時間", "SST Samoa DST": "米領サモア夏時間", "SYOT": "昭和基地時間", "TAAF": "仏領南方南極時間", "TAHT": "タヒチ時間", "TJT": "タジキスタン時間", "TKT": "トケラウ時間", "TLT": "東ティモール時間", "TMST": "トルクメニスタン夏時間", "TMT": "トルクメニスタン標準時", "TOST": "トンガ夏時間", "TOT": "トンガ標準時", "TVT": "ツバル時間", "TWT": "台湾標準時", "TWT DST": "台湾夏時間", "ULAST": "ウランバートル夏時間", "ULAT": "ウランバートル標準時", "UYST": "ウルグアイ夏時間", "UYT": "ウルグアイ標準時", "UZT": "ウズベキスタン標準時", "UZT DST": "ウズベキスタン夏時間", "VET": "ベネズエラ時間", "VLAST": "ウラジオストク夏時間", "VLAT": "ウラジオストク標準時", "VOLST": "ボルゴグラード夏時間", "VOLT": "ボルゴグラード標準時", "VOST": "ボストーク基地時間", "VUT": "バヌアツ標準時", "VUT DST": "バヌアツ夏時間", "WAKT": "ウェーク島時間", "WARST": "西部アルゼンチン夏時間", "WART": "西部アルゼンチン標準時", "WAST": "西アフリカ時間", "WAT": "西アフリカ時間", "WESZ": "西ヨーロッパ夏時間", "WEZ": "西ヨーロッパ標準時", "WFT": "ウォリス・フツナ時間", "WGST": "グリーンランド西部夏時間", "WGT": "グリーンランド西部標準時", "WIB": "インドネシア西部時間", "WIT": "インドネシア東部時間", "WITA": "インドネシア中部時間", "YAKST": "ヤクーツク夏時間", "YAKT": "ヤクーツク標準時", "YEKST": "エカテリンブルグ夏時間", "YEKT": "エカテリンブルグ標準時", "YST": "ユーコン時間", "МСК": "モスクワ標準時", "اقتاۋ": "アクタウ標準時", "اقتاۋ قالاسى": "アクタウ夏時間", "اقتوبە": "アクトベ標準時", "اقتوبە قالاسى": "アクトベ夏時間", "الماتى": "アルトマイ標準時", "الماتى قالاسى": "アルマトイ夏時間", "باتىس قازاق ەلى": "西カザフスタン時間", "شىعىش قازاق ەلى": "東カザフスタン時間", "قازاق ەلى": "カザフスタン時間", "قىرعىزستان": "キルギス時間", "قىزىلوردا": "クズロルダ標準時", "قىزىلوردا قالاسى": "クズロルダ夏時間", "∅∅∅": "ペルー夏時間"},
	}
}

// Locale returns the current translators string locale
func (ja *ja) Locale() string {
	return ja.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ja'
func (ja *ja) PluralsCardinal() []locales.PluralRule {
	return ja.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ja'
func (ja *ja) PluralsOrdinal() []locales.PluralRule {
	return ja.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ja'
func (ja *ja) PluralsRange() []locales.PluralRule {
	return ja.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ja'
func (ja *ja) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ja'
func (ja *ja) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ja'
func (ja *ja) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ja *ja) MonthAbbreviated(month time.Month) string {
	if len(ja.monthsAbbreviated) == 0 {
		return ""
	}
	return ja.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ja *ja) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ja *ja) MonthNarrow(month time.Month) string {
	if len(ja.monthsNarrow) == 0 {
		return ""
	}
	return ja.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ja *ja) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (ja *ja) MonthWide(month time.Month) string {
	if len(ja.monthsWide) == 0 {
		return ""
	}
	return ja.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ja *ja) MonthsWide() []string {
	return ja.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ja *ja) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ja.daysAbbreviated) == 0 {
		return ""
	}
	return ja.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ja *ja) WeekdaysAbbreviated() []string {
	return ja.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ja *ja) WeekdayNarrow(weekday time.Weekday) string {
	if len(ja.daysNarrow) == 0 {
		return ""
	}
	return ja.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ja *ja) WeekdaysNarrow() []string {
	return ja.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ja *ja) WeekdayShort(weekday time.Weekday) string {
	if len(ja.daysShort) == 0 {
		return ""
	}
	return ja.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ja *ja) WeekdaysShort() []string {
	return ja.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ja *ja) WeekdayWide(weekday time.Weekday) string {
	if len(ja.daysWide) == 0 {
		return ""
	}
	return ja.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ja *ja) WeekdaysWide() []string {
	return ja.daysWide
}

// Decimal returns the decimal point of number
func (ja *ja) Decimal() string {
	return ja.decimal
}

// Group returns the group of number
func (ja *ja) Group() string {
	return ja.group
}

// Group returns the minus sign of number
func (ja *ja) Minus() string {
	return ja.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ja' and handles both Whole and Real numbers based on 'v'
func (ja *ja) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ja.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ja.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ja' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ja *ja) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ja.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ja.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ja'
func (ja *ja) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ja.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ja.group[0])
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
		b = append(b, ja.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ja.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ja'
// in accounting notation.
func (ja *ja) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ja.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ja.group[0])
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

		b = append(b, ja.currencyNegativePrefix[0])

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
			b = append(b, ja.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ja.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ja'
func (ja *ja) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ja'
func (ja *ja) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ja'
func (ja *ja) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ja'
func (ja *ja) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)
	b = append(b, ja.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ja'
func (ja *ja) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ja'
func (ja *ja) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ja'
func (ja *ja) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ja'
func (ja *ja) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0xe6, 0x99, 0x82}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe5, 0x88, 0x86}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe7, 0xa7, 0x92, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ja.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
