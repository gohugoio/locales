package km

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type km struct {
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
	periodsAbbreviated     []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'km' locale
func New() locales.Translator {
	return &km{
		locale:                 "km",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "៛", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "ឡូទី", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsNarrow:           []string{"", "ម", "ក", "ម", "ម", "ឧ", "ម", "ក", "ស", "ក", "ត", "វ", "ធ"},
		monthsWide:             []string{"", "មករា", "កុម្ភៈ", "មីនា", "មេសា", "ឧសភា", "មិថុនា", "កក្កដា", "សីហា", "កញ្ញា", "តុលា", "វិច្ឆិកា", "ធ្នូ"},
		daysAbbreviated:        []string{"អាទិត្យ", "ចន្ទ", "អង្គារ", "ពុធ", "ព្រហ", "សុក្រ", "សៅរ៍"},
		daysNarrow:             []string{"អ", "ច", "អ", "ព", "ព", "ស", "ស"},
		daysShort:              []string{"អា", "ច", "អ", "ពុ", "ព្រ", "សុ", "ស"},
		daysWide:               []string{"អាទិត្យ", "ច័ន្ទ", "អង្គារ", "ពុធ", "ព្រហស្បតិ៍", "សុក្រ", "សៅរ៍"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "ម៉ោង\u200bពេលថ្ងៃ\u200b\u200b\u200b\u200bនៅ\u200bអូស្ត្រាលី\u200bកណ្ដាល", "ACST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអូស្ត្រាលី\u200bកណ្ដាល", "ACT": "ACT", "ACWDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200b\u200bភាគ\u200bខាង\u200bលិច\u200bនៃ\u200bអូស្ត្រាលី\u200bកណ្ដាល", "ACWST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bភាគ\u200bខាង\u200bលិច\u200bនៃ\u200bអូស្ត្រាលី\u200bកណ្ដាល", "ADT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអាត្លង់ទិក", "ADT Arabia": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអារ៉ាប់", "AEDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអូស្ត្រាលី\u200bខាង\u200bកើត", "AEST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអូស្ត្រាលី\u200bខាង\u200bកើត", "AFT": "ម៉ោង\u200bនៅ\u200bអាហ្វហ្គានីស្ថាន", "AKDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200b\u200bអាឡាស្កា", "AKST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអាឡាស្កា", "AMST": "ម៉ោង\u200bនៅ\u200bអាម៉ាហ្សូននារដូវក្តៅ", "AMST Armenia": "ម៉ោង\u200bនៅ\u200bអាមេនីនារដូវ\u200bក្ដៅ\u200b", "AMT": "ម៉ោងស្តង់ដារ\u200bនៅ\u200bអាម៉ាហ្សូន", "AMT Armenia": "ម៉ោង\u200bស្ដង់ដារ\u200bនៅ\u200bអាមេនី", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ម៉ោង\u200bនៅ\u200bអាហ្សង់ទីននារដូវក្តៅ", "ART": "ម៉ោងស្តង់ដារ\u200bនៅ\u200bអាហ្សង់ទីន", "AST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអាត្លង់ទិក", "AST Arabia": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអារ៉ាប់", "AWDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអូស្ត្រាលី\u200bខាង\u200bលិច", "AWST": "ម៉ោង\u200b\u200bស្តង់ដារ\u200bនៅ\u200bអូស្ត្រាលី\u200bខាង\u200bលិច", "AZST": "ម៉ោង\u200b\u200bនៅ\u200bអាស៊ែបៃហ្សង់នារដូវ\u200bក្ដៅ", "AZT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអាស៊ែបៃហ្សង់", "BDT Bangladesh": "ម៉ោង\u200b\u200bនៅ\u200bបង់ក្លាដែសនារដូវ\u200bក្ដៅ", "BNT": "ម៉ោងនៅព្រុយណេដារូសាឡឹម", "BOT": "ម៉ោង\u200bនៅ\u200bបូលីវី", "BRST": "ម៉ោង\u200bនៅ\u200bប្រាស៊ីលីយ៉ានា\u200b\u200bរដូវ\u200bក្ដៅ", "BRT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bប្រាស៊ីលីយ៉ា", "BST Bangladesh": "ម៉ោង\u200bស្ដង់ដារ\u200bនៅ\u200bបង់ក្លាដែស", "BT": "ម៉ោងនៅប៊ូតង់", "CAST": "CAST", "CAT": "ម៉ោង\u200bនៅ\u200bអាហ្វ្រិក\u200bកណ្ដាល", "CCT": "ម៉ោង\u200bនៅ\u200bប្រជុំកោះ\u200bកូកូស", "CDT": "ម៉ោង\u200b\u200bពេលថ្ងៃនៅ\u200bទ្វីបអាមេរិក\u200bខាង\u200bជើងភាគកណ្តាល", "CHADT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bចាថាំ", "CHAST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bចាថាំ", "CHUT": "ម៉ោង\u200bនៅ\u200bចូអុក", "CKT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bប្រជុំកោះ\u200bខូក", "CKT DST": "ម៉ោង\u200bនៅប្រជុំ\u200bកោះ\u200bខូកនាពាក់កណ្ដាល\u200bរដូវ\u200b\u200b\u200bក្ដៅ", "CLST": "ម៉ោងនៅស៊ីលីនារដូវក្តៅ", "CLT": "ម៉ោងស្តង់ដារនៅស៊ីលី", "COST": "ម៉ោង\u200bនៅ\u200bកូឡុំប៊ីនា\u200bរដូវ\u200bក្ដៅ", "COT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bកូឡុំប៊ី", "CST": "ម៉ោង\u200b\u200bស្តង់ដារនៅ\u200bទ្វីបអាមេរិក\u200bខាង\u200bជើងភាគកណ្តាល", "CST China": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bចិន", "CST China DST": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bចិន", "CVST": "ម៉ោង\u200b\u200bនៅ\u200bកាប់វែរនារដូវ\u200bក្ដៅ", "CVT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bកាប់វែរ", "CXT": "ម៉ោង\u200bនៅ\u200bកោះ\u200bគ្រីស្មាស", "ChST": "ម៉ោង\u200bស្តង់ដារនៅ\u200bចាំម៉ូរ៉ូ", "ChST NMI": "ChST NMI", "CuDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bគុយបា", "CuST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bគុយបា", "DAVT": "ម៉ោង\u200bនៅ\u200bដាវីស", "DDUT": "ម៉ោង\u200bនៅ\u200bឌុយម៉ុងដឺអ៊ុយវីល", "EASST": "ម៉ោងនៅកោះអ៊ីស្ទ័រនារដូវក្តៅ", "EAST": "ម៉ោងស្តង់ដារនៅកោះអ៊ីស្ទ័រ", "EAT": "ម៉ោង\u200bនៅ\u200bអាហ្វ្រិក\u200bខាង\u200bកើត", "ECT": "ម៉ោង\u200bនៅ\u200bអេក្វាទ័រ", "EDT": "ម៉ោងពេលថ្ងៃនៅទ្វីបអាមេរិកខាងជើងភាគខាងកើត", "EGDT": "ម៉ោង\u200bនៅ\u200bហ្គ្រីនលែនខាង\u200bកើតនា\u200bរដូវ\u200bក្ដៅ", "EGST": "ម៉ោង\u200b\u200b\u200bស្តង់ដារ\u200bនៅ\u200b\u200bហ្គ្រីនលែន\u200bខាង\u200bកើត", "EST": "ម៉ោងស្តង់ដារនៅទ្វីបអាមេរិកខាងជើងភាគខាងកើត", "FEET": "ម៉ោង\u200bនៅ\u200bចុងបូព៌ានៃទ្វីប\u200bអឺរ៉ុប\u200b", "FJT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bហ្វីជី", "FJT Summer": "ម៉ោង\u200bនៅ\u200b\u200bហ្វីជីនា\u200b\u200bរដូវ\u200bក្ដៅ", "FKST": "ម៉ោង\u200b\u200bនៅប្រជុំ\u200bកោះ\u200bហ្វក់ឡែននារដូវ\u200bក្ដៅ", "FKT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅប្រជុំ\u200bកោះ\u200bហ្វក់ឡែន", "FNST": "ម៉ោង\u200bនៅហ្វ៊ែណាន់ដូ\u200bដឺណូរ៉ូញ៉ានារដូវក្តៅ", "FNT": "ម៉ោង\u200bស្តង់ដារនៅហ្វ៊ែណាន់ដូ\u200bដឺណូរ៉ូញ៉ា", "GALT": "ម៉ោង\u200bនៅ\u200bកាឡាប៉ាកូស", "GAMT": "ម៉ោង\u200bនៅ\u200bកាំបៀ", "GEST": "ម៉ោង\u200bនៅ\u200bហ្សកហ្ស៊ីនា\u200b\u200bរដូវ\u200bក្ដៅ", "GET": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bហ្សកហ្ស៊ី", "GFT": "ម៉ោង\u200bនៅ\u200bហ្គីយ៉ាន\u200bបារាំង", "GIT": "ម៉ោង\u200bនៅ\u200bកោះ\u200bកីប៊ឺត", "GMT": "ម៉ោងនៅគ្រីនវិច", "GNSST": "GNSST", "GNST": "GNST", "GST": "ម៉ោង\u200bស្តង់ដា\u200bនៅ\u200bកាល់", "GST Guam": "GST Guam", "GYT": "ម៉ោង\u200bនៅ\u200bហ្គីយ៉ាន", "HADT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bហាវៃ-អាល់ដ្យូសិន", "HAST": "ម៉ោង\u200bស្តង់ដារ\u200b\u200bនៅ\u200bហាវៃ-អាល់ដ្យូសិន", "HKST": "ម៉ោងនៅ\u200bហុងកុងនា\u200bរដូវ\u200bក្ដៅ\u200b", "HKT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bហុងកុង", "HOVST": "ម៉ោងនៅ\u200bហូវនា\u200bរដូវ\u200bក្ដៅ\u200b", "HOVT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅហូវ", "ICT": "ម៉ោង\u200bនៅ\u200bឥណ្ឌូចិន", "IDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអ៊ីស្រាអែល", "IOT": "ម៉ោង\u200bនៅ\u200bមហាសមុទ្រ\u200bឥណ្ឌា", "IRKST": "ម៉ោងនៅអៀរគុតស្កិ៍នារដូវក្តៅ", "IRKT": "ម៉ោងស្តង់ដារនៅអៀរគុតស្កិ៍", "IRST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអ៊ីរ៉ង់", "IRST DST": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអ៊ីរ៉ង់", "IST": "ម៉ោង\u200bស្តង់ដារនៅ\u200bឥណ្ឌា", "IST Israel": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអ៊ីស្រាអែល", "JDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅជប៉ុន", "JST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bជប៉ុន", "KOST": "ម៉ោង\u200bនៅ\u200bកូស្រៃ", "KRAST": "ម៉ោង\u200bនៅ\u200bក្រាណូយ៉ាសនា\u200bរដូវ\u200bក្ដៅ", "KRAT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bក្រាណូយ៉ាស", "KST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bកូរ៉េ", "KST DST": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bកូរ៉េ", "LHDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bឡតហៅ", "LHST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bឡត\u200bហៅ", "LINT": "ម៉ោង\u200bនៅ\u200bកោះ\u200bឡាញ", "MAGST": "ម៉ោង\u200bនៅ\u200bម៉ាហ្កាដាន\u200bនារដូវ\u200bក្ដៅ", "MAGT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bម៉ាហ្កាដាន", "MART": "ម៉ោង\u200bនៅ\u200bកោះ\u200bម៉ាគឺសាស់", "MAWT": "ម៉ោង\u200bនៅ\u200bម៉ៅ\u200bសាន់", "MDT": "MDT", "MESZ": "ម៉ោង\u200bនៅ\u200bអឺរ៉ុប\u200bកណ្ដាលនា\u200bរដូវ\u200bក្ដៅ", "MEZ": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអឺរ៉ុប\u200bកណ្ដាល", "MHT": "ម៉ោង\u200bនៅ\u200bម៉ាសាល", "MMT": "ម៉ោង\u200bនៅ\u200bមីយ៉ាន់ម៉ា", "MSD": "ម៉ោង\u200bនៅ\u200bមូស្គូ\u200bនារដូវ\u200bក្ដៅ", "MST": "MST", "MUST": "ម៉ោង\u200b\u200bរដូវ\u200bក្ដៅនៅ\u200bម៉ូរីស", "MUT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bម៉ូរីស", "MVT": "ម៉ោង\u200bនៅ\u200bម៉ាល់ឌីវ", "MYT": "ម៉ោង\u200bនៅ\u200bម៉ាឡេស៊ី", "NCT": "ម៉ោងស្តង់ដារ\u200bនៅណូវ៉ែលកាឡេដូនៀ", "NDT": "ម៉ោង\u200bពេលថ្ងៃ\u200bនៅ\u200bញូវហ្វោនឡែន", "NDT New Caledonia": "ម៉ោង\u200bនៅណូវ៉ែលកាឡេដូនៀនារដូវក្តៅ", "NFDT": "ម៉ោងនៅ\u200bណ័រហ្វក់នា\u200bរដូវ\u200bក្ដៅ", "NFT": "ម៉ោង\u200bស្ដង់ដារ\u200bនៅ\u200bណ័រហ្វក់", "NOVST": "ម៉ោង\u200bនៅ\u200bណូវ៉ូស៊ីប៊ីកនា\u200bរដូវ\u200bក្ដៅ", "NOVT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bណូវ៉ូស៊ីប៊ីក", "NPT": "ម៉ោងនៅនេប៉ាល់", "NRT": "ម៉ោង\u200bនៅ\u200bណូរូ", "NST": "ម៉ោង\u200b\u200bស្តង់ដារ\u200b\u200bនៅ\u200bញូវហ្វោនឡែន", "NUT": "ម៉ោងនៅ\u200bនីវ៉េ", "NZDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bនូវែលសេឡង់", "NZST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bនូវែលសេឡង់", "OESZ": "ម៉ោង\u200bនៅ\u200bអឺរ៉ុប\u200b\u200bខាង\u200bកើត\u200bនា\u200bរដូវ\u200bក្ដៅ", "OEZ": "ម៉ោង\u200bស្តង់ដារ\u200b\u200bនៅ\u200bអឺរ៉ុប\u200b\u200bខាង\u200bកើត\u200b", "OMSST": "ម៉ោង\u200bនៅ\u200bអូមនា\u200bរដូវ\u200bក្ដៅ", "OMST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអូម", "PDT": "ម៉ោងពេលថ្ងៃនៅប៉ាស៊ីហ្វិកអាមេរិក", "PDTM": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bប៉ាស៊ីហ្វិក\u200bម៉ិកស៊ិក", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ម៉ោង\u200bនៅប៉ាពូអាស៊ី នូវែលហ្គីណេ", "PHOT": "ម៉ោង\u200bនៅ\u200bកោះ\u200bផូនីក", "PKT": "ម៉ោង\u200bស្ដង់ដារ\u200bនៅ\u200bប៉ាគីស្ថាន", "PKT DST": "ម៉ោងនៅ\u200bប៉ាគីស្ថាននា\u200bរដូវ\u200bក្ដៅ\u200b", "PMDT": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅសង់\u200bព្យែរ និង\u200bមីគុយឡុង", "PMST": "ម៉ោង\u200bស្តង់ដារ\u200bនៅសង់\u200bព្យែរ និង\u200bមីគុយឡុង", "PONT": "ម៉ោង\u200bនៅ\u200bប៉ូណាប់", "PST": "ម៉ោងស្ដង់ដារនៅប៉ាស៊ីហ្វិកអាមេរិក", "PST Philippine": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bហ្វីលីពីន", "PST Philippine DST": "ម៉ោង\u200b\u200bនៅ\u200bហ្វីលីពីននា\u200bរដូវ\u200bក្ដៅ", "PST Pitcairn": "ម៉ោង\u200bនៅ\u200bភីឃឺន", "PSTM": "ម៉ោង\u200bស្តង់ដា\u200bនៅ\u200bប៉ាស៊ីហ្វិក\u200bម៉ិកស៊ិក", "PWT": "ម៉ោង\u200bនៅ\u200bផាឡៅ", "PYST": "ម៉ោង\u200bនៅប៉ារ៉ាហ្គាយនា\u200bរដូវ\u200bក្ដៅ", "PYT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bប៉ារ៉ាហ្គាយ", "PYT Korea": "ម៉ោងនៅព្យុងយ៉ាង", "RET": "ម៉ោងនៅរេអ៊ុយ៉ុង", "ROTT": "ម៉ោង\u200bនៅ\u200bរ៉ូធឺរ៉ា", "SAKST": "ម៉ោង\u200bនៅ\u200bសាក់ខាលីននា\u200bរដូវ\u200bក្ដៅ", "SAKT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bសាក់ខាលីន", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ម៉ោង\u200bនៅ\u200bអាហ្វ្រិក\u200bខាង\u200bត្បូង", "SBT": "ម៉ោង\u200bនៅ\u200bកោះ\u200bសូឡូម៉ុន", "SCT": "ម៉ោង\u200bនៅ\u200bសីស្ហែល", "SGT": "ម៉ោង\u200bនៅ\u200bសិង្ហបូរី", "SLST": "SLST", "SRT": "ម៉ោង\u200bនៅ\u200bសូរីណាម", "SST Samoa": "ម៉ោង\u200bស្តង់ដារនៅ\u200bសាម័រ", "SST Samoa Apia": "ម៉ោង\u200bស្តង់ដា\u200bនៅ\u200bអាប្យា", "SST Samoa Apia DST": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bអាប្យា", "SST Samoa DST": "ម៉ោង\u200bនៅ\u200bសាម័រនារដូវក្តៅ", "SYOT": "ម៉ោង\u200bនៅ\u200bស៊ីអូវ៉ា", "TAAF": "ម៉ោងនៅបារាំងខាងត្បូង និងនៅអង់តាំងទិក", "TAHT": "ម៉ោង\u200bនៅ\u200bតាហិទី", "TJT": "ម៉ោងនៅតាជីគីស្ថាន", "TKT": "ម៉ោង\u200bនៅ\u200bតូខេឡៅ", "TLT": "ម៉ោង\u200bនៅ\u200b\u200bទីម័រ\u200bខាង\u200bកើត", "TMST": "ម៉ោង\u200bរដូវ\u200bក្ដៅ\u200bនៅ\u200bតួកម៉េនីស្ថាន\u200b", "TMT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅតួកម៉េនីស្ថាន", "TOST": "ម៉ោង\u200b\u200bនៅ\u200bតុងហ្គានារដូវ\u200bក្ដៅ", "TOT": "ម៉ោង\u200bស្តង់ដារ\u200b\u200bនៅ\u200bតុងហ្គា", "TVT": "ម៉ោង\u200bនៅ\u200bទុយវ៉ាលូ", "TWT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bតៃប៉ិ", "TWT DST": "ម៉ោង\u200bពេល\u200bថ្ងៃ\u200bនៅ\u200bតៃប៉ិ", "ULAST": "ម៉ោងនៅ\u200bអ៊ូឡាន\u200bបាទូនា\u200bរដូវ\u200bក្ដៅ\u200b", "ULAT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអ៊ូឡាន\u200bបាទូ", "UYST": "ម៉ោង\u200bនៅ\u200bអ៊ុយរូហ្គាយនា\u200b\u200bរដូវ\u200bក្ដៅ", "UYT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអ៊ុយរូហ្គាយ", "UZT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអ៊ូសបេគីស្ថាន", "UZT DST": "ម៉ោង\u200bនៅ\u200bអ៊ូសបេគីស្ថាននារដូវ\u200bក្ដៅ\u200b", "VET": "ម៉ោង\u200bនៅ\u200bវ៉េណេស៊ុយអេឡា", "VLAST": "ម៉ោង\u200bនៅ\u200bវ៉្លាឌីវ៉ូស្តុកនា\u200bរដូវ\u200bក្ដៅ", "VLAT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bវ៉្លាឌីវ៉ូស្តុក", "VOLST": "ម៉ោង\u200bនៅ\u200bវ៉ូហ្កោក្រាដនា\u200bរដូវ\u200bក្ដៅ", "VOLT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bវ៉ូហ្កោក្រាដ", "VOST": "ម៉ោង\u200bនៅ\u200bវ័រស្តុក", "VUT": "ម៉ោង\u200b\u200bស្តង់ដារ\u200bនៅ\u200bវ៉ានូទូ", "VUT DST": "ម៉ោង\u200bនៅ\u200bវ៉ានូទូនារដូវ\u200bក្ដៅ\u200b", "WAKT": "ម៉ោង\u200bនៅ\u200bកោះវេក", "WARST": "ម៉ោង\u200bនៅ\u200bអាហ្សង់ទីនភាគខាងលិចនារដូវក្តៅ", "WART": "ម៉ោងស្តង់ដារ\u200bនៅ\u200bអាហ្សង់ទីនភាគខាងលិច", "WAST": "ម៉ោង\u200bនៅ\u200bអាហ្វ្រិក\u200bខាង\u200bលិច", "WAT": "ម៉ោង\u200bនៅ\u200bអាហ្វ្រិក\u200bខាង\u200bលិច", "WESZ": "ម៉ោង\u200bនៅ\u200bអឺរ៉ុប\u200bខាង\u200bលិចនារដូវ\u200bក្ដៅ\u200b", "WEZ": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអឺរ៉ុប\u200bខាង\u200bលិច", "WFT": "ម៉ោង\u200bនៅ\u200bវ៉ាលីស និងហ្វ៊ុទូណា", "WGST": "ម៉ោងនៅហ្គ្រីនលែនខាងលិចនារដូវក្តៅ", "WGT": "ម៉ោងស្តង់ដារនៅហ្គ្រីនលែនខាងលិច", "WIB": "ម៉ោង\u200bនៅ\u200bឥណ្ឌូណេស៊ី\u200b\u200bខាង\u200bលិច", "WIT": "ម៉ោង\u200bនៅ\u200bឥណ្ឌូណេស៊ី\u200b\u200bខាង\u200bកើត", "WITA": "ម៉ោង\u200bនៅ\u200bឥណ្ឌូណេស៊ី\u200b\u200b\u200bកណ្ដាល", "YAKST": "ម៉ោង\u200bនៅ\u200bយ៉ាគុតស្កិ៍នា\u200bរដូវ\u200bក្ដៅ", "YAKT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bយ៉ាគុតស្កិ៍", "YEKST": "ម៉ោង\u200bនៅ\u200bអ៊ិខាធឺរីនប៊័កនា\u200bរដូវ\u200b\u200bក្ដៅ", "YEKT": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bអ៊ិខាធឺរីនប៊័ក", "YST": "ម៉ោងនៅយូខន់", "МСК": "ម៉ោង\u200bស្តង់ដារ\u200bនៅ\u200bមូស្គូ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ម៉ោង\u200bនៅ\u200bកាហ្សាក់ស្ថាន\u200bខាង\u200b\u200b\u200bលិច", "شىعىش قازاق ەلى": "ម៉ោង\u200bកាហ្សាក់ស្ថាន\u200b\u200bខាង\u200bកើត", "قازاق ەلى": "ពេលវេលានៅកាហ្សាក់ស្ថាន", "قىرعىزستان": "ម៉ោងនៅកៀហ្ស៊ីស៊ីស្ថាន", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ម៉ោង\u200b\u200bនៅ\u200bអេហ្សសនារដូវ\u200bក្ដៅ"},
	}
}

// Locale returns the current translators string locale
func (km *km) Locale() string {
	return km.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'km'
func (km *km) PluralsCardinal() []locales.PluralRule {
	return km.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'km'
func (km *km) PluralsOrdinal() []locales.PluralRule {
	return km.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'km'
func (km *km) PluralsRange() []locales.PluralRule {
	return km.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'km'
func (km *km) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'km'
func (km *km) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'km'
func (km *km) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (km *km) MonthAbbreviated(month time.Month) string {
	if len(km.monthsAbbreviated) == 0 {
		return ""
	}
	return km.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (km *km) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (km *km) MonthNarrow(month time.Month) string {
	if len(km.monthsNarrow) == 0 {
		return ""
	}
	return km.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (km *km) MonthsNarrow() []string {
	return km.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (km *km) MonthWide(month time.Month) string {
	if len(km.monthsWide) == 0 {
		return ""
	}
	return km.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (km *km) MonthsWide() []string {
	return km.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (km *km) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(km.daysAbbreviated) == 0 {
		return ""
	}
	return km.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (km *km) WeekdaysAbbreviated() []string {
	return km.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (km *km) WeekdayNarrow(weekday time.Weekday) string {
	if len(km.daysNarrow) == 0 {
		return ""
	}
	return km.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (km *km) WeekdaysNarrow() []string {
	return km.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (km *km) WeekdayShort(weekday time.Weekday) string {
	if len(km.daysShort) == 0 {
		return ""
	}
	return km.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (km *km) WeekdaysShort() []string {
	return km.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (km *km) WeekdayWide(weekday time.Weekday) string {
	if len(km.daysWide) == 0 {
		return ""
	}
	return km.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (km *km) WeekdaysWide() []string {
	return km.daysWide
}

// Decimal returns the decimal point of number
func (km *km) Decimal() string {
	return km.decimal
}

// Group returns the group of number
func (km *km) Group() string {
	return km.group
}

// Group returns the minus sign of number
func (km *km) Minus() string {
	return km.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'km' and handles both Whole and Real numbers based on 'v'
func (km *km) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, km.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, km.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, km.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'km' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (km *km) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, km.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, km.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, km.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'km'
func (km *km) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := km.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, km.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, km.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, km.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, km.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'km'
// in accounting notation.
func (km *km) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := km.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, km.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, km.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, km.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, km.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, km.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'km'
func (km *km) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'km'
func (km *km) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, km.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'km'
func (km *km) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, km.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'km'
func (km *km) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, km.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, km.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'km'
func (km *km) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, km.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, km.periodsAbbreviated[0]...)
	} else {
		b = append(b, km.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'km'
func (km *km) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, km.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, km.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, km.periodsAbbreviated[0]...)
	} else {
		b = append(b, km.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'km'
func (km *km) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, km.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, km.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, km.periodsAbbreviated[0]...)
	} else {
		b = append(b, km.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'km'
func (km *km) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, km.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, km.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, km.periodsAbbreviated[0]...)
	} else {
		b = append(b, km.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := km.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
