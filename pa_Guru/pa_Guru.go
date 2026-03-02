package pa_Guru

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type pa_Guru struct {
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
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'pa_Guru' locale
func New() locales.Translator {
	return &pa_Guru{
		locale:                 "pa_Guru",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ਜਨ", "ਫ਼ਰ", "ਮਾਰਚ", "ਅਪ੍ਰੈ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾ", "ਅਗ", "ਸਤੰ", "ਅਕਤੂ", "ਨਵੰ", "ਦਸੰ"},
		monthsNarrow:           []string{"", "ਜ", "ਫ਼", "ਮਾ", "ਅ", "ਮ", "ਜੂ", "ਜੁ", "ਅ", "ਸ", "ਅ", "ਨ", "ਦ"},
		monthsWide:             []string{"", "ਜਨਵਰੀ", "ਫ਼ਰਵਰੀ", "ਮਾਰਚ", "ਅਪ੍ਰੈਲ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾਈ", "ਅਗਸਤ", "ਸਤੰਬਰ", "ਅਕਤੂਬਰ", "ਨਵੰਬਰ", "ਦਸੰਬਰ"},
		daysAbbreviated:        []string{"ਐਤ", "ਸੋਮ", "ਮੰਗਲ", "ਬੁੱਧ", "ਵੀਰ", "ਸ਼ੁੱਕਰ", "ਸ਼ਨੀ"},
		daysNarrow:             []string{"ਐ", "ਸੋ", "ਮੰ", "ਬੁੱ", "ਵੀ", "ਸ਼ੁੱ", "ਸ਼"},
		daysShort:              []string{"ਐਤ", "ਸੋਮ", "ਮੰਗ", "ਬੁੱਧ", "ਵੀਰ", "ਸ਼ੁੱਕ", "ਸ਼ਨੀ"},
		daysWide:               []string{"ਐਤਵਾਰ", "ਸੋਮਵਾਰ", "ਮੰਗਲਵਾਰ", "ਬੁੱਧਵਾਰ", "ਵੀਰਵਾਰ", "ਸ਼ੁੱਕਰਵਾਰ", "ਸ਼ਨੀਵਾਰ"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪੱਛਮੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "ACWST": "ਆਸਟ੍ਰੇਲੀਆਈ ਕੇਂਦਰੀ ਪੱਛਮੀ ਮਿਆਰੀ ਵੇਲਾ", "ADT": "ਅਟਲਾਂਟਿਕ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "ADT Arabia": "ਅਰਬੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "AEDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੂਰਬੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "AEST": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੂਰਬੀ ਮਿਆਰੀ ਵੇਲਾ", "AFT": "ਅਫ਼ਗਾਨਿਸਤਾਨ ਵੇਲਾ", "AKDT": "ਅਲਾਸਕਾ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "AKST": "ਅਲਾਸਕਾ ਮਿਆਰੀ ਵੇਲਾ", "AMST": "ਅਮੇਜ਼ਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "AMST Armenia": "ਅਰਮੀਨੀਆ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "AMT": "ਅਮੇਜ਼ਨ ਮਿਆਰੀ ਵੇਲਾ", "AMT Armenia": "ਅਰਮੀਨੀਆ ਮਿਆਰੀ ਵੇਲਾ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ਅਰਜਨਟੀਨਾ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "ART": "ਅਰਜਨਟੀਨਾ ਮਿਆਰੀ ਵੇਲਾ", "AST": "ਅਟਲਾਂਟਿਕ ਮਿਆਰੀ ਵੇਲਾ", "AST Arabia": "ਅਰਬੀ ਮਿਆਰੀ ਵੇਲਾ", "AWDT": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੱਛਮੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "AWST": "ਆਸਟ੍ਰੇਲੀਆਈ ਪੱਛਮੀ ਮਿਆਰੀ ਵੇਲਾ", "AZST": "ਅਜ਼ਰਬਾਈਜਾਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "AZT": "ਅਜ਼ਰਬਾਈਜਾਨ ਮਿਆਰੀ ਵੇਲਾ", "BDT Bangladesh": "ਬੰਗਲਾਦੇਸ਼ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "BNT": "ਬਰੂਨੇਈ ਦਾਰੂਸਲਾਮ ਵੇਲਾ", "BOT": "ਬੋਲੀਵੀਆ ਵੇਲਾ", "BRST": "ਬ੍ਰਾਜ਼ੀਲੀਆ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "BRT": "ਬ੍ਰਾਜ਼ੀਲੀਆ ਮਿਆਰੀ ਵੇਲਾ", "BST Bangladesh": "ਬੰਗਲਾਦੇਸ਼ ਮਿਆਰੀ ਵੇਲਾ", "BT": "ਭੂਟਾਨ ਵੇਲਾ", "CAST": "ਕੇਸੀ ਸਮਾਂ", "CAT": "ਕੇਂਦਰੀ ਅਫ਼ਰੀਕਾ ਵੇਲਾ", "CCT": "ਕੋਕਸ ਆਈਲੈਂਡ ਵੇਲਾ", "CDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਕੇਂਦਰੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "CHADT": "ਚੈਥਮ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "CHAST": "ਚੈਥਮ ਮਿਆਰੀ ਵੇਲਾ", "CHUT": "ਚੂਕ ਵੇਲਾ", "CKT": "ਕੁੱਕ ਆਈਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "CKT DST": "ਕੁੱਕ ਆਈਲੈਂਡ ਅੱਧ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "CLST": "ਚਿਲੀ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "CLT": "ਚਿਲੀ ਮਿਆਰੀ ਵੇਲਾ", "COST": "ਕੋਲੰਬੀਆ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "COT": "ਕੋਲੰਬੀਆ ਮਿਆਰੀ ਵੇਲਾ", "CST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਕੇਂਦਰੀ ਮਿਆਰੀ ਵੇਲਾ", "CST China": "ਚੀਨ ਮਿਆਰੀ ਵੇਲਾ", "CST China DST": "ਚੀਨ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "CVST": "ਕੇਪ ਵਰਡ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "CVT": "ਕੇਪ ਵਰਡ ਮਿਆਰੀ ਵੇਲਾ", "CXT": "ਕ੍ਰਿਸਮਸ ਆਈਲੈਂਡ ਵੇਲਾ", "ChST": "ਚਾਮੋਰੋ ਮਿਆਰੀ ਵੇਲਾ", "ChST NMI": "ਉੱਤਰੀ ਮਰਿਆਨਾ ਆਈਲੈਂਡ ਸਮਾਂ", "CuDT": "ਕਿਊਬਾ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "CuST": "ਕਿਊਬਾ ਮਿਆਰੀ ਵੇਲਾ", "DAVT": "ਡੇਵਿਸ ਵੇਲਾ", "DDUT": "ਡਿਉਮੋਂਟ ਡਿਉਰਵਿਲੇ ਵੇਲਾ", "EASST": "ਈਸਟਰ ਆਈਲੈਂਡ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "EAST": "ਈਸਟਰ ਆਈਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "EAT": "ਪੂਰਬੀ ਅਫ਼ਰੀਕਾ ਵੇਲਾ", "ECT": "ਇਕਵੇਡੋਰ ਵੇਲਾ", "EDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੂਰਬੀ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "EGDT": "ਪੂਰਬੀ ਗ੍ਰੀਨਲੈਂਡ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "EGST": "ਪੂਰਬੀ ਗ੍ਰੀਨਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "EST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੂਰਬੀ ਮਿਆਰੀ ਵੇਲਾ", "FEET": "ਹੋਰ-ਪੂਰਬੀ ਯੂਰਪੀ ਵੇਲਾ", "FJT": "ਫ਼ਿਜ਼ੀ ਮਿਆਰੀ ਵੇਲਾ", "FJT Summer": "ਫ਼ਿਜ਼ੀ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "FKST": "ਫ਼ਾਕਲੈਂਡ ਆਈਲੈਂਡਸ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "FKT": "ਫ਼ਾਕਲੈਂਡ ਆਈਲੈਂਡਸ ਮਿਆਰੀ ਵੇਲਾ", "FNST": "ਫਰਨਾਂਡੋ ਡੇ ਨੋਰੋਨਹਾ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "FNT": "ਫਰਨਾਂਡੋ ਡੇ ਨੋਰੋਨਹਾ ਮਿਆਰੀ ਵੇਲਾ", "GALT": "ਗਲਾਪਾਗੋਸ ਵੇਲਾ", "GAMT": "ਗੈਂਬੀਅਰ ਵੇਲਾ", "GEST": "ਜਾਰਜੀਆ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "GET": "ਜਾਰਜੀਆ ਮਿਆਰੀ ਵੇਲਾ", "GFT": "ਫ੍ਰੈਂਚ ਗੁਏਨਾ ਵੇਲਾ", "GIT": "ਗਿਲਬਰਟ ਆਈਲੈਂਡ ਵੇਲਾ", "GMT": "ਗ੍ਰੀਨਵਿਚ ਮੀਨ ਵੇਲਾ", "GNSST": "GNSST", "GNST": "GNST", "GST": "ਖਾੜੀ ਮਿਆਰੀ ਵੇਲਾ", "GST Guam": "ਗੁਆਮ ਸਮਾਂ", "GYT": "ਗੁਯਾਨਾ ਵੇਲਾ", "HADT": "ਹਵਾਈ-ਅਲੇਯੂਸ਼ਿਅਨ ਮਿਆਰੀ ਵੇਲਾ", "HAST": "ਹਵਾਈ-ਅਲੇਯੂਸ਼ਿਅਨ ਮਿਆਰੀ ਵੇਲਾ", "HKST": "ਹਾਂਗ ਕਾਂਗ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "HKT": "ਹਾਂਗ ਕਾਂਗ ਮਿਆਰੀ ਵੇਲਾ", "HOVST": "ਹੋਵਡ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "HOVT": "ਹੋਵਡ ਮਿਆਰੀ ਵੇਲਾ", "ICT": "ਇੰਡੋਚਾਈਨਾ ਵੇਲਾ", "IDT": "ਇਜ਼ਰਾਈਲ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "IOT": "ਹਿੰਦ ਮਹਾਂਸਾਗਰ ਵੇਲਾ", "IRKST": "ਇਰਕੁਤਸਕ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "IRKT": "ਇਰਕੁਤਸਕ ਮਿਆਰੀ ਵੇਲਾ", "IRST": "ਈਰਾਨ ਮਿਆਰੀ ਵੇਲਾ", "IRST DST": "ਈਰਾਨ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "IST": "ਭਾਰਤੀ ਮਿਆਰੀ ਵੇਲਾ", "IST Israel": "ਇਜ਼ਰਾਈਲ ਮਿਆਰੀ ਵੇਲਾ", "JDT": "ਜਪਾਨ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "JST": "ਜਪਾਨ ਮਿਆਰੀ ਵੇਲਾ", "KOST": "ਕੋਸਰੇ ਵੇਲਾ", "KRAST": "ਕ੍ਰਾਸਨੋਯਾਰਸਕ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "KRAT": "ਕ੍ਰਾਸਨੋਯਾਰਸਕ ਮਿਆਰੀ ਵੇਲਾ", "KST": "ਕੋਰੀਆਈ ਮਿਆਰੀ ਵੇਲਾ", "KST DST": "ਕੋਰੀਆਈ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "LHDT": "ਲੌਰਡ ਹੋਵੇ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "LHST": "ਲੌਰਡ ਹੋਵੇ ਮਿਆਰੀ ਵੇਲਾ", "LINT": "ਲਾਈਨ ਆਈਲੈਂਡ ਵੇਲਾ", "MAGST": "ਮੈਗੇਡਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "MAGT": "ਮੈਗੇਡਨ ਮਿਆਰੀ ਵੇਲਾ", "MART": "ਮਾਰਕਿਸਾਸ ਵੇਲਾ", "MAWT": "ਮੌਸਨ ਵੇਲਾ", "MDT": "ਮਕਾਉ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "MESZ": "ਮੱਧ ਯੂਰਪੀ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "MEZ": "ਮੱਧ ਯੂਰਪੀ ਮਿਆਰੀ ਵੇਲਾ", "MHT": "ਮਾਰਸ਼ਲ ਆਈਲੈਂਡ ਵੇਲਾ", "MMT": "ਮਿਆਂਮਾਰ ਵੇਲਾ", "MSD": "ਮਾਸਕੋ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "MST": "ਮਕਾਉ ਮਿਆਰੀ ਸਮਾਂ", "MUST": "ਮੌਰਿਸ਼ਸ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "MUT": "ਮੌਰਿਸ਼ਸ ਮਿਆਰੀ ਵੇਲਾ", "MVT": "ਮਾਲਦੀਵ ਵੇਲਾ", "MYT": "ਮਲੇਸ਼ੀਆ ਵੇਲਾ", "NCT": "ਨਿਊ ਕੈਲੇਡੋਨੀਆ ਮਿਆਰੀ ਵੇਲਾ", "NDT": "ਨਿਊਫਾਉਂਡਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "NDT New Caledonia": "ਨਿਊ ਕੈਲੇਡੋਨੀਆ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "NFDT": "ਨੋਰਫੌਕ ਆਈਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "NFT": "ਨੋਰਫੌਕ ਆਈਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "NOVST": "ਨੌਵੋਸਿਬੀਰਸਕ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "NOVT": "ਨੌਵੋਸਿਬੀਰਸਕ ਮਿਆਰੀ ਵੇਲਾ", "NPT": "ਨੇਪਾਲ ਵੇਲਾ", "NRT": "ਨਾਉਰੂ ਵੇਲਾ", "NST": "ਨਿਊਫਾਉਂਡਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "NUT": "ਨੀਊ ਵੇਲਾ", "NZDT": "ਨਿਊਜ਼ੀਲੈਂਡ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "NZST": "ਨਿਊਜ਼ੀਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "OESZ": "ਪੂਰਬੀ ਯੂਰਪੀ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "OEZ": "ਪੂਰਬੀ ਯੂਰਪੀ ਮਿਆਰੀ ਵੇਲਾ", "OMSST": "ਓਮਸਕ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "OMST": "ਓਮਸਕ ਮਿਆਰੀ ਵੇਲਾ", "PDT": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੈਸਿਫਿਕ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "PDTM": "ਮੈਕਸੀਕਨ ਪੈਸਿਫਿਕ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ਪਾਪੂਆ ਨਿਊ ਗਿਨੀ ਵੇਲਾ", "PHOT": "ਫਿਨਿਕਸ ਆਈਲੈਂਡ ਵੇਲਾ", "PKT": "ਪਾਕਿਸਤਾਨ ਮਿਆਰੀ ਵੇਲਾ", "PKT DST": "ਪਾਕਿਸਤਾਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "PMDT": "ਸੈਂਟ ਪੀਅਰੇ ਅਤੇ ਮਿਕੇਲਨ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "PMST": "ਸੈਂਟ ਪੀਅਰੇ ਅਤੇ ਮਿਕੇਲਨ ਮਿਆਰੀ ਵੇਲਾ", "PONT": "ਪੋਨਾਪੇ ਵੇਲਾ", "PST": "ਉੱਤਰੀ ਅਮਰੀਕੀ ਪੈਸਿਫਿਕ ਮਿਆਰੀ ਵੇਲਾ", "PST Philippine": "ਫਿਲਿਪੀਨੀ ਮਿਆਰੀ ਵੇਲਾ", "PST Philippine DST": "ਫਿਲਿਪੀਨੀ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "PST Pitcairn": "ਪਿਟਕੈਰਨ ਵੇਲਾ", "PSTM": "ਮੈਕਸੀਕਨ ਪੈਸਿਫਿਕ ਮਿਆਰੀ ਵੇਲਾ", "PWT": "ਪਲਾਉ ਵੇਲਾ", "PYST": "ਪੈਰਾਗਵੇ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "PYT": "ਪੈਰਾਗਵੇ ਮਿਆਰੀ ਵੇਲਾ", "PYT Korea": "ਪਯੋਂਗਯਾਂਗ ਵੇਲਾ", "RET": "ਰਿਯੂਨੀਅਨ ਵੇਲਾ", "ROTT": "ਰੋਥੇਰਾ ਵੇਲਾ", "SAKST": "ਸਖਲੀਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "SAKT": "ਸਖਲੀਨ ਮਿਆਰੀ ਵੇਲਾ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ਦੱਖਣੀ ਅਫ਼ਰੀਕਾ ਮਿਆਰੀ ਵੇਲਾ", "SBT": "ਸੋਲੋਮਨ ਆਈਲੈਂਡਸ ਵੇਲਾ", "SCT": "ਸੇਸ਼ਲਸ ਵੇਲਾ", "SGT": "ਸਿੰਗਾਪੁਰ ਮਿਆਰੀ ਵੇਲਾ", "SLST": "ਲੰਕਾ ਸਮਾਂ", "SRT": "ਸੂਰੀਨਾਮ ਵੇਲਾ", "SST Samoa": "ਸਾਮੋਆ ਮਿਆਰੀ ਵੇਲਾ", "SST Samoa Apia": "ਐਪੀਆ ਮਿਆਰੀ ਵੇਲਾ", "SST Samoa Apia DST": "ਐਪੀਆ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "SST Samoa DST": "ਸਾਮੋਆ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "SYOT": "ਸਿਓਵਾ ਵੇਲਾ", "TAAF": "ਫ੍ਰੈਂਚ ਦੱਖਣੀ ਅਤੇ ਐਂਟਾਰਟਿਕ ਵੇਲਾ", "TAHT": "ਤਾਹੀਤੀ ਵੇਲਾ", "TJT": "ਤਾਜਿਕਿਸਤਾਨ ਵੇਲਾ", "TKT": "ਟੋਕੇਲਾਉ ਵੇਲਾ", "TLT": "ਪੂਰਬੀ ਤਿਮੂਰ ਵੇਲਾ", "TMST": "ਤੁਰਕਮੇਨਿਸਤਾਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "TMT": "ਤੁਰਕਮੇਨਿਸਤਾਨ ਮਿਆਰੀ ਵੇਲਾ", "TOST": "ਟੋਂਗਾ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "TOT": "ਟੋਂਗਾ ਮਿਆਰੀ ਵੇਲਾ", "TVT": "ਟੁਵਾਲੂ ਵੇਲਾ", "TWT": "ਤੈਪਈ ਮਿਆਰੀ ਵੇਲਾ", "TWT DST": "ਤੈਪਈ ਪ੍ਰਕਾਸ਼ ਵੇਲਾ", "ULAST": "ਉਲਨ ਬਟੋਰ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "ULAT": "ਉਲਨ ਬਟੋਰ ਮਿਆਰੀ ਵੇਲਾ", "UYST": "ਉਰੂਗਵੇ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "UYT": "ਉਰੂਗਵੇ ਮਿਆਰੀ ਵੇਲਾ", "UZT": "ਉਜ਼ਬੇਕਿਸਤਾਨ ਮਿਆਰੀ ਵੇਲਾ", "UZT DST": "ਉਜ਼ਬੇਕਿਸਤਾਨ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "VET": "ਵੈਨੇਜ਼ੂਏਲਾ ਵੇਲਾ", "VLAST": "ਵਲਾਦੀਵੋਸਤਕ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "VLAT": "ਵਲਾਦੀਵੋਸਤਕ ਮਿਆਰੀ ਵੇਲਾ", "VOLST": "ਵੋਲਗੋਗ੍ਰੇਡ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "VOLT": "ਵੋਲਗੋਗ੍ਰੇਡ ਮਿਆਰੀ ਵੇਲਾ", "VOST": "ਵੋਸਟੋਕ ਵੇਲਾ", "VUT": "ਵਾਨੂਆਟੂ ਮਿਆਰੀ ਵੇਲਾ", "VUT DST": "ਵਾਨੂਆਟੂ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "WAKT": "ਵੇਕ ਆਈਲੈਂਡ ਵੇਲਾ", "WARST": "ਪੱਛਮੀ ਅਰਜਨਟੀਨਾ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "WART": "ਪੱਛਮੀ ਅਰਜਨਟੀਨਾ ਮਿਆਰੀ ਵੇਲਾ", "WAST": "ਪੱਛਮੀ ਅਫਰੀਕਾ ਵੇਲਾ", "WAT": "ਪੱਛਮੀ ਅਫਰੀਕਾ ਵੇਲਾ", "WESZ": "ਪੱਛਮੀ ਯੂਰਪੀ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "WEZ": "ਪੱਛਮੀ ਯੂਰਪੀ ਮਿਆਰੀ ਵੇਲਾ", "WFT": "ਵਾਲਿਸ ਅਤੇ ਫੁਟੂਨਾ ਵੇਲਾ", "WGST": "ਪੱਛਮੀ ਗ੍ਰੀਨਲੈਂਡ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "WGT": "ਪੱਛਮੀ ਗ੍ਰੀਨਲੈਂਡ ਮਿਆਰੀ ਵੇਲਾ", "WIB": "ਪੱਛਮੀ ਇੰਡੋਨੇਸ਼ੀਆ ਵੇਲਾ", "WIT": "ਪੂਰਬੀ ਇੰਡੋਨੇਸ਼ੀਆ ਵੇਲਾ", "WITA": "ਮੱਧ ਇੰਡੋਨੇਸ਼ੀਆਈ ਵੇਲਾ", "YAKST": "ਯਕੁਤਸਕ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "YAKT": "ਯਕੁਤਸਕ ਮਿਆਰੀ ਵੇਲਾ", "YEKST": "ਯਕੇਤਰਿਨਬਰਗ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ", "YEKT": "ਯਕੇਤਰਿਨਬਰਗ ਮਿਆਰੀ ਵੇਲਾ", "YST": "ਯੂਕੋਨ ਸਮਾਂ", "МСК": "ਮਾਸਕੋ ਮਿਆਰੀ ਵੇਲਾ", "اقتاۋ": "ਅਕਤਾਉ ਮਿਆਰੀ ਸਮਾਂ", "اقتاۋ قالاسى": "ਅਕਤਾਉ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "اقتوبە": "ਅਕਤੋਬ ਮਿਆਰੀ ਸਮਾਂ", "اقتوبە قالاسى": "ਅਕਤੋਬ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "الماتى": "ਅਲਮਾਟੀ ਮਿਆਰੀ ਸਮਾਂ", "الماتى قالاسى": "ਅਲਮਾਟੀ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "باتىس قازاق ەلى": "ਪੱਛਮੀ ਕਜ਼ਾਖ਼ਸਤਾਨ ਵੇਲਾ", "شىعىش قازاق ەلى": "ਪੂਰਬੀ ਕਜ਼ਾਖ਼ਸਤਾਨ ਵੇਲਾ", "قازاق ەلى": "ਕਜ਼ਾਖ਼ਸਤਾਨ ਵੇਲਾ", "قىرعىزستان": "ਕਿਰਗਿਸਤਾਨ ਵੇਲਾ", "قىزىلوردا": "ਕਿਜ਼ਲੋਰਡਾ ਮਿਆਰੀ ਸਮਾਂ", "قىزىلوردا قالاسى": "ਕਿਜ਼ਲੋਰਡਾ ਗਰਮੀ-ਰੁੱਤ ਸਮਾਂ", "∅∅∅": "ਪੇਰੂ ਗਰਮੀਆਂ ਦਾ ਵੇਲਾ"},
	}
}

// Locale returns the current translators string locale
func (pa *pa_Guru) Locale() string {
	return pa.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pa_Guru'
func (pa *pa_Guru) PluralsCardinal() []locales.PluralRule {
	return pa.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pa_Guru'
func (pa *pa_Guru) PluralsOrdinal() []locales.PluralRule {
	return pa.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pa_Guru'
func (pa *pa_Guru) PluralsRange() []locales.PluralRule {
	return pa.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pa_Guru'
func (pa *pa_Guru) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pa_Guru'
func (pa *pa_Guru) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pa_Guru'
func (pa *pa_Guru) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := pa.CardinalPluralRule(num1, v1)
	end := pa.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pa *pa_Guru) MonthAbbreviated(month time.Month) string {
	return pa.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pa *pa_Guru) MonthsAbbreviated() []string {
	return pa.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pa *pa_Guru) MonthNarrow(month time.Month) string {
	return pa.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pa *pa_Guru) MonthsNarrow() []string {
	return pa.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pa *pa_Guru) MonthWide(month time.Month) string {
	return pa.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pa *pa_Guru) MonthsWide() []string {
	return pa.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pa *pa_Guru) WeekdayAbbreviated(weekday time.Weekday) string {
	return pa.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pa *pa_Guru) WeekdaysAbbreviated() []string {
	return pa.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pa *pa_Guru) WeekdayNarrow(weekday time.Weekday) string {
	return pa.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pa *pa_Guru) WeekdaysNarrow() []string {
	return pa.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pa *pa_Guru) WeekdayShort(weekday time.Weekday) string {
	return pa.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pa *pa_Guru) WeekdaysShort() []string {
	return pa.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pa *pa_Guru) WeekdayWide(weekday time.Weekday) string {
	return pa.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pa *pa_Guru) WeekdaysWide() []string {
	return pa.daysWide
}

// Decimal returns the decimal point of number
func (pa *pa_Guru) Decimal() string {
	return pa.decimal
}

// Group returns the group of number
func (pa *pa_Guru) Group() string {
	return pa.group
}

// Group returns the minus sign of number
func (pa *pa_Guru) Minus() string {
	return pa.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pa_Guru' and handles both Whole and Real numbers based on 'v'
func (pa *pa_Guru) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, pa.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pa_Guru' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pa *pa_Guru) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pa.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pa_Guru'
func (pa *pa_Guru) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pa.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, pa.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(pa.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, pa.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, pa.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pa_Guru'
// in accounting notation.
func (pa *pa_Guru) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pa.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pa.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, pa.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(pa.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, pa.currencyNegativePrefix[j])
		}

		b = append(b, pa.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(pa.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, pa.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pa.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, pa.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pa.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pa_Guru'
func (pa *pa_Guru) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, pa.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pa.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, pa.periodsAbbreviated[0]...)
	} else {
		b = append(b, pa.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pa.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
