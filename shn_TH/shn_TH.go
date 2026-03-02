package shn_TH

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type shn_TH struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'shn_TH' locale
func New() locales.Translator {
	return &shn_TH{
		locale:             "shn_TH",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "ၸၼ်ႇ", "ၾႅပ်ႇ", "မၢတ်ႉၶ်ျ", "ဢေႇ", "မေႇ", "ၸုၼ်ႇ", "ၸူႇ", "ဢေႃး", "သႅပ်ႇ", "ဢွၵ်ႇ", "ၼူဝ်ႇ", "တီႇ"},
		monthsNarrow:       []string{"", "ၸ.", "ၾ.", "မ.", "ဢ.", "မ.", "ၸ.", "ၸ.", "ဢ.", "သ.", "ဢ.", "ၼ.", "တ."},
		monthsWide:         []string{"", "ၸၼ်ႇဝႃႇရီႇ", "ၾႅပ်ႇဝႃႇရီႇ", "မၢတ်ႉၶ်ျ", "ဢေႇပရႄႇ", "မေႇ", "ၸုၼ်ႇ", "ၸူႇလၢႆႇ", "ဢေႃးၵၢတ်ႉ", "သႅပ်ႇထႅမ်ႇပႃႇ", "ဢွၵ်ႇထူဝ်ႇပႃႇ", "ၼူဝ်ႇဝႅမ်ႇပႃႇ", "တီႇသႅမ်ႇပႃႇ"},
		daysAbbreviated:    []string{"တိတ်ႉ", "ၸၼ်", "ၵၢၼ်း", "ပုတ်ႉ", "ၽတ်း", "သုၵ်း", "သဝ်"},
		daysNarrow:         []string{"တိ.", "ၸ.", "ၵ.", "ပု.", "ၽ.", "သု.", "သ."},
		daysShort:          []string{"တိတ်ႉ", "ၸၼ်", "ၵၢၼ်း", "ပုတ်ႉ", "ၽတ်း", "သုၵ်း", "သဝ်"},
		daysWide:           []string{"ဝၼ်းဢႃးတိတ်ႉ", "ဝၼ်းၸၼ်", "ဝၼ်းဢင်းၵၢၼ်း", "ဝၼ်းပုတ်ႉ", "ဝၼ်းၽတ်း", "ဝၼ်းသုၵ်း", "ဝၼ်းသဝ်"},
		periodsAbbreviated: []string{"တၸ.", "တလ."},
		periodsNarrow:      []string{"ၸ.", "လ."},
		periodsWide:        []string{"တွၼ်ႈၸဝ်ႉ", "တွၼ်ႈလႃႈ"},
		erasAbbreviated:    []string{"ပီႇၸီႇ", "ဢေႇတီႇ"},
		erasNarrow:         []string{"ပီႇၸီႇ", "ဢေႇတီႇ"},
		erasWide:           []string{"ဢွၼ်ၼႃႈၸဝ်ႈၶရိတ်ႉ", "ပီၶရိတ်ႉ"},
		timezones:          map[string]string{"ACDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢေႃႉသထရေးလီးယႃး", "ACST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢေႇၵႃႇ", "ACT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢေႇၵႃႇ", "ACWDT": "ၶၢဝ်းယၢမ်း ပွတ်းတူၵ်း တွၼ်ႈၵၢင် ၶၢဝ်းမႆႈ ဢေႃႉသထရေးလီးယႃး", "ACWST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပွတ်းတူၵ်း တွၼ်ႈၵၢင် ဢေႃႉသထရေးလီးယႃး", "ADT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢႅတ်ႉလႅၼ်းတိၵ်ႉ", "ADT Arabia": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢႃႇရၢပ်ႈ", "AEDT": "ၶၢဝ်းယၢမ်း ပွတ်းဢွၵ်ႇ ၵၢင်ဝၼ်း ဢေႃႉသထရေးလီးယႃး", "AEST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပွတ်းဢွၵ်ႇ ဢေႃႉသထရေးလီးယႃး", "AFT": "ၶၢဝ်းယၢမ်း ဢႃႇၾၵၢၼ်ႇၼီႇသတၢၼ်ႇ", "AKDT": "ၶၢဝ်းယၢမ်းၵၢင်ဝၼ်း ဢလႅတ်ႉသၵႃႇ", "AKST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢလႅတ်ႉသၵႃႇ", "AMST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢမၸွၼ်ႇ", "AMST Armenia": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢႃႇမေးၼီးယႃး", "AMT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢမၸွၼ်ႇ", "AMT Armenia": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇမေးၼီးယႃး", "ANAST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢႃႇၼႃႇတီႇယႃႇ", "ANAT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇၼႃႇတီႇယႃႇ", "ARST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢႃႇၵျႅၼ်ႇတီးၼႃး", "ART": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇၵျႅၼ်ႇတီးၼႃး", "AST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႅတ်ႉလႅၼ်းတိၵ်ႉ", "AST Arabia": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇရၢပ်ႈ", "AWDT": "ၶၢဝ်းယၢမ်း ပွတ်းတူၵ်း ၵၢင်ဝၼ်း ဢေႃႉသထရေးလီးယႃး", "AWST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပွတ်းတူၵ်း ဢေႃႉသထရေးလီးယႃး", "AZST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢႃႇၸႃႇပၢႆႇၸၼ်ႇ", "AZT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇၸႃႇပၢႆႇၸၼ်ႇ", "BDT Bangladesh": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ပင်းၵလႃးတဵတ်ႈ", "BNT": "ၶၢဝ်းယၢမ်း ပရူႇၼၢႆး", "BOT": "ၶၢဝ်းယၢမ်း ပူဝ်ႇလီးပီးယႃး", "BRST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ပရႃႇၸီးလီးယႃး", "BRT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပရႃႇၸီးလီးယႃး", "BST Bangladesh": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပင်းၵလႃးတဵတ်ႈ", "BT": "ၶၢဝ်းယၢမ်း ၽူႇတၢၼ်ႇ", "CAST": "ၶၢဝ်းယၢမ်း ၶေႇသီႇ", "CAT": "ၶၢဝ်းယၢမ်း ဢႃႇၾရိၵ ပွတ်းၵၢင်", "CCT": "ၶၢဝ်းယၢမ်း မူႇၵုၼ်ၵူဝ်ႇၵူဝ်းသ်", "CDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ပွတ်းၵၢင်", "CHADT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၶျႅတ်ႉထမ်ႇ", "CHAST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶျႅတ်ႉထမ်ႇ", "CHUT": "ၶၢဝ်းယၢမ်း ၶျူၵ်ႉ", "CKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း မူႇၵုၼ်ၶုၵ်ႈ", "CKT DST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ မူႇၵုၼ်ၶုၵ်ႈ", "CLST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၶျီႇလီႇ", "CLT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶျီႇလီႇ", "COST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵူဝ်ႇလမ်ႇပီႇယႃႇ", "COT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵူဝ်ႇလမ်ႇပီႇယႃႇ", "CST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပွတ်းၵၢင်", "CST China": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶႄႇ", "CST China DST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၶႄႇ", "CVST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၶဵပ်ႉဝႄႇတီႇ", "CVT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶဵပ်ႉဝႄႇတီႇ", "CXT": "ၶၢဝ်းယၢမ်း ၵုၼ်ၶရိတ်ႉသမၢတ်ႉ", "ChST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶျမ်ႇမူဝ်ႇရူဝ်ႇ", "ChST NMI": "ၶၢဝ်းယၢမ်း မူႇၵုၼ်မႃႇရီႇယႃႇၼႃႇ ပွတ်းႁွင်ႇ", "CuDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵူးပႃး", "CuST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵူးပႃး", "DAVT": "ၶၢဝ်းယၢမ်း တေးဝိတ်ႉ", "DDUT": "ၶၢဝ်းယၢမ်း တူႇမွၼ်ႉ တဝီႇလ်", "EASST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵုၼ်ဢီႇသတႃႇ", "EAST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵုၼ်ဢီႇသတႃႇ", "EAT": "ၶၢဝ်းယၢမ်း ဢႃႇၾရိၵ ပွတ်းဢွၵ်ႇ", "ECT": "ၶၢဝ်းယၢမ်း ဢေႇၵႂႃႇတေႃႇ", "EDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ပွတ်းဢွၵ်ႇ", "EGDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵရိၼ်းလႅၼ်း ပွတ်းဢွၵ်ႇ", "EGST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵရိၼ်းလႅၼ်း ပွတ်းဢွၵ်ႇ", "EST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပွတ်းဢွၵ်ႇ", "FEET": "ၶၢဝ်းယၢမ်း ယူးရူပ်ႉ ပွတ်းဢွၵ်ႇ ဢၼ်ၵႆတိူဝ်း", "FJT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၾီႇၵျီႇ", "FJT Summer": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၾီႇၵျီႇ", "FKST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း မူႇၵုၼ် ၾွၵ်ႉလႅၼ်ႇ", "FKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း မူႇၵုၼ် ၾွၵ်ႉလႅၼ်ႇ", "FNST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၾႃႇၼႅၼ်ႇတူဝ်ႇ တႄႇ တီႇၼူဝ်ႇရွၼ်ႇႁႃႇ", "FNT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၾႃႇၼႅၼ်ႇတူဝ်ႇ တႄႇ တီႇၼူဝ်ႇရွၼ်ႇႁႃႇ", "GALT": "ၶၢဝ်းယၢမ်း ၵႃႇလႃႇပၵူဝ်ႉသ်", "GAMT": "ၶၢဝ်းယၢမ်း ၵမ်ႇပီႇယႃႇ", "GEST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၵျေႃႇၵျႃႇ", "GET": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵျေႃႇၵျႃႇ", "GFT": "ၶၢဝ်းယၢမ်း ၵုၺ်ႇယႃႇၼႃႇ ၶွင် ၾရၢင်ႇသဵတ်ႈ", "GIT": "ၶၢဝ်းယၢမ်း ၵိလ်ပႅတ်ႉ", "GMT": "ၶၢဝ်းယၢမ်း ၽတ်ႉၽဵင်ႇ ၵရိၼ်းဝိတ်ႉ", "GNSST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၵရိၼ်းလႅၼ်း", "GNST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵရိၼ်းလႅၼ်း", "GST": "ၶၢဝ်းယၢမ်း ၵျေႃႇၵျႃႇ ပွတ်းၸၢၼ်း", "GST Guam": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵႂၢမ်ႇ", "GYT": "ၶၢဝ်းယၢမ်း ၵၢႆႇယႃးၼႃႇ", "HADT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ႁႃႇဝၢႆးဢီး-ဢလူးသျၼ်ႇ", "HAST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ႁႃႇဝၢႆးဢီး-ဢလူးသျၼ်ႇ", "HKST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ႁွင်းၵွင်း", "HKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ႁွင်းၵွင်း", "HOVST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၶူဝ်ဝေတေး", "HOVT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶူဝ်ဝေတေး", "ICT": "ၶၢဝ်းယၢမ်း ဢိၼ်ႇတူဝ်ႇၶႄႇ", "IDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢိတ်ႇသရေး", "IOT": "ၶၢဝ်းယၢမ်း သမုၵ်ႉတရႃႇဢိၼ်းတီးယႃး", "IRKST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢိရၶုတ်ႉသ်ၶ်", "IRKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢိရၶုတ်ႉသ်ၶ်", "IRST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢီႇရၢၼ်း", "IRST DST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢီႇရၢၼ်း", "IST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢိၼ်းတီးယႃး", "IST Israel": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢိတ်ႇသရေး", "JDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵျႃႇပၢၼ်ႇ", "JST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵျႃႇပၢၼ်ႇ", "KOST": "ၶၢဝ်းယၢမ်း ၵွတ်ႉသ်ရေး", "KRAST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၶရတ်ႉၼူဝ်ႇယႃႇသ်ၶ်", "KRAT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶရတ်ႉၼူဝ်ႇယႃႇသ်ၶ်", "KST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵၢဝ်းလီ", "KST DST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵၢဝ်းလီ", "LHDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း လွတ်ႉႁၢဝ်း", "LHST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း လွတ်ႉႁၢဝ်း", "LINT": "ၶၢဝ်းယၢမ်း မူႇၵုၼ်လၢႆး", "MAGST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ မႃႇၵႃႇတၢၼ်ႇ", "MAGT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း မႃႇၵႃႇတၢၼ်ႇ", "MART": "ၶၢဝ်းယၢမ်း မႃၶေးသႅတ်ႉသ်", "MAWT": "ၶၢဝ်းယၢမ်း မေႃသၼ်ႇ", "MDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း သၼ်လွႆ", "MESZ": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ယူးရူပ်ႉ ပွတ်းၵၢင်", "MEZ": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ယူးရူပ်ႉ ပွတ်းၵၢင်", "MHT": "ၶၢဝ်းယၢမ်း မူႇၵုၼ်မႃးသျႄႇ", "MMT": "ၶၢဝ်းယၢမ်း မျၢၼ်ႇမႃႇ", "MSD": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ မွတ်ႉသ်ၵူဝ်ႇ", "MST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သၼ်လွႆ", "MUST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ မေႃးရီႇသႃႇ", "MUT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း မေႃးရီႇသႃႇ", "MVT": "ၶၢဝ်းယၢမ်း မေႃႇတိပ်ႈ", "MYT": "ၶၢဝ်းယၢမ်း မလေးသျႃး", "NCT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၼိဝ်းၶႄႇလီႇတူဝ်းၼီးယႃး", "NDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၼိဝ်းၾွမ်ႇလႅၼ်ႇ", "NDT New Caledonia": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၼိဝ်းၶႄႇလီႇတူဝ်းၼီးယႃး", "NFDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵုၼ်ၼေႃႇၾုၵ်ႉ", "NFT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵုၼ်ၼေႃႇၾုၵ်ႉ", "NOVST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၼူဝ်ႇဝူဝ်ႇသီႇပၢတ်ႉသ်ၶ်", "NOVT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၼူဝ်ႇဝူဝ်ႇသီႇပၢတ်ႉသ်ၶ်", "NPT": "ၶၢဝ်းယၢမ်း ၼေႇပေႃး", "NRT": "ၶၢဝ်းယၢမ်း ၼၢဝ်ရူး", "NST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၼိဝ်းၾွမ်ႇလႅၼ်ႇ", "NUT": "ၶၢဝ်းယၢမ်း ၼီးဝႄႇ", "NZDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၼိဝ်းၸီႇလႅၼ်ႇ", "NZST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၼိဝ်းၸီႇလႅၼ်ႇ", "OESZ": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ယူးရူပ်ႉ ပွတ်းဢွၵ်ႇ", "OEZ": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ယူးရူပ်ႉ ပွတ်းဢွၵ်ႇ", "OMSST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢူမ်းသ်ၶ်", "OMST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢူမ်းသ်ၶ်", "PDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ပၸိၾိၵ်ႉ", "PDTM": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း မႅၵ်ႇသီႇၵူဝ်ႇ ပၸိၾိၵ်ႉ", "PETDT": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၵမ်ႇၶျတ်ႉၵႃႇ", "PETST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵမ်ႇၶျတ်ႉၵႃႇ", "PGT": "ၶၢဝ်းယၢမ်း ပႃးပႂႃႇၼိဝ်းၵီးၼီး", "PHOT": "ၶၢဝ်းယၢမ်း မူႇၵုၼ်ၽီးၼိၵ်ႉ", "PKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပႃႇၵိတ်ႈသတၼ်ႇ", "PKT DST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ပႃႇၵိတ်ႈသတၼ်ႇ", "PMDT": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း သဵင်ႉပီးယႃး လႄႈ မိၵ်ႈၵွႆႇလွၼ်ႇ", "PMST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သဵင်ႉပီးယႃး လႄႈ မိၵ်ႈၵွႆႇလွၼ်ႇ", "PONT": "ၶၢဝ်းယၢမ်း ပူၼ်းပေႇ", "PST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပၸိၾိၵ်ႉ", "PST Philippine": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၾီလိပ်ႈပိၼ်း", "PST Philippine DST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၾီလိပ်ႈပိၼ်း", "PST Pitcairn": "ၶၢဝ်းယၢမ်း ၽိတ်ႉၶႅၼ်ႇ", "PSTM": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း မႅၵ်ႇသီႇၵူဝ်ႇ ပၸိၾိၵ်ႉ", "PWT": "ၶၢဝ်းယၢမ်း ပႃႇလၢဝ်း", "PYST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ပႃႇရႃႇၵူၺ်း", "PYT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ပႃႇရႃႇၵူၺ်း", "PYT Korea": "ၶၢဝ်းယၢမ်း ၵၢဝ်းလီႁွင်ႇ", "RET": "ၶၢဝ်းယၢမ်း ရေႇၼီႇယၼ်ႇ", "ROTT": "ၶၢဝ်းယၢမ်း ရူဝ်ႇထႄႇရႃႇ", "SAKST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ သႃႇၶႃႇလိၼ်ႇ", "SAKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သႃႇၶႃႇလိၼ်ႇ", "SAMST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ သႃႇမႃႇရႃႇ", "SAMT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သႃႇမႃႇရႃႇ", "SAST": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇၾရိၵၸၢၼ်း", "SBT": "ၶၢဝ်းယၢမ်း မူႇၵုၼ်သေႃႇလေႃႇမၼ်ႇ", "SCT": "ၶၢဝ်းယၢမ်း သေးသျႄႇ", "SGT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သိင်ႇၵႃႇပူဝ်ႇ", "SLST": "ၶၢဝ်းယၢမ်း လင်းၵႃ", "SRT": "ၶၢဝ်းယၢမ်း သျူးရီးၼႃႇမႄႇ", "SST Samoa": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သႃႇမူဝ်းဝႃႇ ၶွ\u200bင် ဢမႄႇရိၵၢၼ်ႇ", "SST Samoa Apia": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း သႃႇမူဝ်းဝႃႇ", "SST Samoa Apia DST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း သႃႇမူဝ်းဝႃႇ", "SST Samoa DST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း သႃႇမူဝ်းဝႃႇ ၶွ\u200bင် ဢမႄႇရိၵၢၼ်ႇ", "SYOT": "ၶၢဝ်းယၢမ်း သၢႆးဢူဝ်ႇဝႃႇ", "TAAF": "ၶၢဝ်းယၢမ်း ၾရၢင်ႇသဵတ်ႈ ပွတ်းၸၢၼ်း လႄႈ ဢႅၼ်ႇတၢၵ်ႈတီးၵႃႈ", "TAHT": "ၶၢဝ်းယၢမ်း တႁီႇတီႇ", "TJT": "ၶၢဝ်းယၢမ်း တႃႇၵျီႇၵီႇသတၼ်ႇ", "TKT": "ၶၢဝ်းယၢမ်း ထူဝ်းၵေႇလၢဝ်ႇ", "TLT": "ၶၢဝ်းယၢမ်း တီႇမေႃးလႅတ်ႉသ်တႄး", "TMST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ တၢၵ်ႈမႅၼ်ႇၼီႇသတၼ်ႇ", "TMT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း တၢၵ်ႈမႅၼ်ႇၼီႇသတၼ်ႇ", "TOST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ထွင်းၵႃႇ", "TOT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ထွင်းၵႃႇ", "TVT": "ၶၢဝ်းယၢမ်း ထူးဝႃႇလူႇ", "TWT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ထၢႆႇဝၢၼ်း", "TWT DST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ထၢႆႇဝၢၼ်း", "ULAST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢူလၼ်ပႃတႃ", "ULAT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢူလၼ်ပႃတႃ", "UYST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢုရုၵူၺ်း", "UYT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢုရုၵူၺ်း", "UZT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢူႇၸပႄႉၵိတ်ႇသတၼ်ႇ", "UZT DST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢူႇၸပႄႉၵိတ်ႇသတၼ်ႇ", "VET": "ၶၢဝ်းယၢမ်း ဝႄႇၼေႇၸွႆးလႃး", "VLAST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဝလႃႇတီႇဝူဝ်ႇသတွၵ်ႉ", "VLAT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဝလႃႇတီႇဝူဝ်ႇသတွၵ်ႉ", "VOLST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဝူဝ်ႇၵူဝ်ႇၵရၢတ်ႉ", "VOLT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဝူဝ်ႇၵူဝ်ႇၵရၢတ်ႉ", "VOST": "ၶၢဝ်းယၢမ်း ဝေႃႉသ်တွၵ်ႉ", "VUT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဝႅၼ်ႇၼူးဝႃႇထူႇ", "VUT DST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဝႅၼ်ႇၼူးဝႃႇထူႇ", "WAKT": "ၶၢဝ်းယၢမ်း ၵုၼ်ဝဵၵ်ႉ", "WARST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ဢႃႇၵျႅၼ်ႇတီးၼႃး ပွတ်းတူၵ်း", "WART": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႃႇၵျႅၼ်ႇတီးၼႃး ပွတ်းတူၵ်း", "WAST": "ၶၢဝ်းယၢမ်း ဢႃႇၾရိၵ ပွတ်းတူၵ်း", "WAT": "ၶၢဝ်းယၢမ်း ဢႃႇၾရိၵ ပွတ်းတူၵ်း", "WESZ": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ယူးရူပ်ႉ ပွတ်းတူၵ်း", "WEZ": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ယူးရူပ်ႉ ပွတ်းတူၵ်း", "WFT": "ၶၢဝ်းယၢမ်း ဝႃးလိတ်ႈ လႄႈ ၾူႇတူးၼႃး", "WGST": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ၵရိၼ်းလႅၼ်း ပွတ်းတူၵ်း", "WGT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၵရိၼ်းလႅၼ်း ပွတ်းတူၵ်း", "WIB": "ၶၢဝ်းယၢမ်း ဢိၼ်ႇတူဝ်ႇၼီးသျႃး ပွတ်းတူၵ်း", "WIT": "ၶၢဝ်းယၢမ်း ဢိၼ်ႇတူဝ်ႇၼီးသျႃး ပွတ်းဢွၵ်ႇ", "WITA": "ၶၢဝ်းယၢမ်း ဢိၼ်ႇတူဝ်ႇၼီးသျႃး ပွတ်းၵၢင်", "YAKST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ယၵုတ်ႉသ်ၶ်", "YAKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ယၵုတ်ႉသ်ၶ်", "YEKST": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ယေႇၵႃႇတႄႇရိၼ်ႇပၢၵ်ႉ", "YEKT": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ယေႇၵႃႇတႄႇရိၼ်ႇပၢၵ်ႉ", "YST": "ၶၢဝ်းယၢမ်း ယူးၵွၼ်ႇ", "МСК": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း မွတ်ႉသ်ၵူဝ်ႇ", "اقتاۋ": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႅၵ်ႉတၢဝ်ႉ", "اقتاۋ قالاسى": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢႅၵ်ႉတၢဝ်ႉ", "اقتوبە": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႅၵ်ႉတူဝ်ႇပီႇ", "اقتوبە قالاسى": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢႅၵ်ႉတူဝ်ႇပီႇ", "الماتى": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ဢႄးလ်မႃႇတီႇ", "الماتى قالاسى": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ဢႄးလ်မႃႇတီႇ", "باتىس قازاق ەلى": "ၶၢဝ်းယၢမ်း ၵႃႇၸၢၵ်ႈသတၼ်ႇ ပွတ်းတူၵ်း", "شىعىش قازاق ەلى": "ၶၢဝ်းယၢမ်း ၵႃႇၸၢၵ်ႈသတၼ်ႇ ပွတ်းဢွၵ်ႇ", "قازاق ەلى": "ၶၢဝ်းယၢမ်း ၵႃႇၸၢၵ်ႈသတၼ်ႇ", "قىرعىزستان": "ၶၢဝ်းယၢမ်း ၵႃႇၵိတ်ႈသတၼ်ႇ", "قىزىلوردا": "လၵ်းၸဵင် ၶၢဝ်းယၢမ်း ၶုၺ်ႇၸီးလူဝ်ႇတႃႇ", "قىزىلوردا قالاسى": "ၶၢဝ်းယၢမ်း ၶၢဝ်းမႆႈ ၶုၺ်ႇၸီးလူဝ်ႇတႃႇ", "∅∅∅": "ၶၢဝ်းယၢမ်း ၵၢင်ဝၼ်း ပေႇရူႉ"},
	}
}

// Locale returns the current translators string locale
func (shn *shn_TH) Locale() string {
	return shn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'shn_TH'
func (shn *shn_TH) PluralsCardinal() []locales.PluralRule {
	return shn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'shn_TH'
func (shn *shn_TH) PluralsOrdinal() []locales.PluralRule {
	return shn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'shn_TH'
func (shn *shn_TH) PluralsRange() []locales.PluralRule {
	return shn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'shn_TH'
func (shn *shn_TH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'shn_TH'
func (shn *shn_TH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'shn_TH'
func (shn *shn_TH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (shn *shn_TH) MonthAbbreviated(month time.Month) string {
	return shn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (shn *shn_TH) MonthsAbbreviated() []string {
	return shn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (shn *shn_TH) MonthNarrow(month time.Month) string {
	return shn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (shn *shn_TH) MonthsNarrow() []string {
	return shn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (shn *shn_TH) MonthWide(month time.Month) string {
	return shn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (shn *shn_TH) MonthsWide() []string {
	return shn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (shn *shn_TH) WeekdayAbbreviated(weekday time.Weekday) string {
	return shn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (shn *shn_TH) WeekdaysAbbreviated() []string {
	return shn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (shn *shn_TH) WeekdayNarrow(weekday time.Weekday) string {
	return shn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (shn *shn_TH) WeekdaysNarrow() []string {
	return shn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (shn *shn_TH) WeekdayShort(weekday time.Weekday) string {
	return shn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (shn *shn_TH) WeekdaysShort() []string {
	return shn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (shn *shn_TH) WeekdayWide(weekday time.Weekday) string {
	return shn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (shn *shn_TH) WeekdaysWide() []string {
	return shn.daysWide
}

// Decimal returns the decimal point of number
func (shn *shn_TH) Decimal() string {
	return shn.decimal
}

// Group returns the group of number
func (shn *shn_TH) Group() string {
	return shn.group
}

// Group returns the minus sign of number
func (shn *shn_TH) Minus() string {
	return shn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'shn_TH' and handles both Whole and Real numbers based on 'v'
func (shn *shn_TH) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'shn_TH' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (shn *shn_TH) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'shn_TH'
func (shn *shn_TH) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shn.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'shn_TH'
// in accounting notation.
func (shn *shn_TH) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shn.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, shn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, shn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, shn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x2d, 0x20}...)
	b = append(b, shn.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'shn_TH'
func (shn *shn_TH) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, shn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := shn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
