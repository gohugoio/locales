package kxv_Telu

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kxv_Telu struct {
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

// New returns a new instance of translator for the 'kxv_Telu' locale
func New() locales.Translator {
	return &kxv_Telu{
		locale:                 "kxv_Telu",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "pusu", "maha", "pagu", "hire", "bese", "jaṭṭa", "aasaḍi", "srabĩ", "bado", "dasara", "divi", "pande"},
		monthsNarrow:           []string{"", "మా", "గు", "హి", "బె", "ల", "రా", "బా", "బా", "అ", "ది", "పా", "పు"},
		monthsWide:             []string{"", "మాగ", "గుండు", "హిరెఇ", "బెసెకి", "లండి", "రాత", "బాన్దపాణా", "బార్సి", "అస్ర", "దివెడి", "పాండు", "పుసు"},
		daysAbbreviated:        []string{"aadi", "smba", "manga", "puda", "laki", "sukru", "sani"},
		daysNarrow:             []string{"వా", "న", "మా", "వు", "ల", "ను", "సా"},
		daysShort:              []string{"aa", "s", "ma", "pu", "laki", "su", "sa"},
		daysWide:               []string{"వారమి", "నమారా", "మాంగాడా", "వుదారా", "లాకివరా", "నుక్ వరా", "సానివరా"},
		periodsAbbreviated:     []string{"ఎ\u202fఎమ్", "పి\u202fఎమ్"},
		timezones:              map[string]string{"ACDT": "ఆస్ట్రేలియా మధ్యమ పగటి వెలుతురు సమయం", "ACST": "ఆస్ట్రేలియా మధ్యమ ప్రామాణిక సమయం", "ACT": "ACT", "ACWDT": "ఆస్ట్రేలియా మధ్యమ పశ్చిమ పగటి వెలుతురు సమయం", "ACWST": "మధ్యమ ఆస్ట్రేలియన్ పశ్చిమ ప్రామాణిక సమయం", "ADT": "అట్లాంటిక్ పగటి వెలుతురు సమయం", "ADT Arabia": "అరేబియన్ పగటి వెలుతురు సమయం", "AEDT": "ఆస్ట్రేలియన్ తూర్పు పగటి వెలుతురు సమయం", "AEST": "ఆస్ట్రేలియన్ తూర్పు ప్రామాణిక సమయం", "AFT": "ఆఫ్ఘనిస్తాన్ సమయం", "AKDT": "అలాస్కా పగటి వెలుతురు సమయం", "AKST": "అలాస్కా ప్రామాణిక సమయం", "AMST": "అమెజాన్ వేసవి సమయం", "AMST Armenia": "ఆర్మేనియా వేసవి సమయం", "AMT": "అమెజాన్ ప్రామాణిక సమయం", "AMT Armenia": "ఆర్మేనియా ప్రామాణిక సమయం", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ఆర్జెంటీనా వేసవి సమయం", "ART": "అర్జెంటీనా ప్రామాణిక సమయం", "AST": "అట్లాంటిక్ ప్రామాణిక సమయం", "AST Arabia": "అరేబియన్ ప్రామాణిక సమయం", "AWDT": "ఆస్ట్రేలియన్ పశ్చిమ పగటి వెలుతురు సమయం", "AWST": "ఆస్ట్రేలియన్ పశ్చిమ ప్రామాణిక సమయం", "AZST": "అజర్బైజాన్ వేసవి సమయం", "AZT": "అజర్బైజాన్ ప్రామాణిక సమయం", "BDT Bangladesh": "బంగ్లాదేశ్ వేసవి సమయం", "BNT": "బ్రూనే దరుసలామ్ సమయం", "BOT": "బొలీవియా సమయం", "BRST": "బ్రెజిలియా వేసవి సమయం", "BRT": "బ్రెజిలియా ప్రామాణిక సమయం", "BST Bangladesh": "బంగ్లాదేశ్ ప్రామాణిక సమయం", "BT": "భూటాన్ సమయం", "CAST": "CAST", "CAT": "సెంట్రల్ ఆఫ్రికా సమయం", "CCT": "కోకోస్ దీవుల సమయం", "CDT": "మధ్యమ పగటి వెలుతురు సమయం", "CHADT": "చాథమ్ పగటి వెలుతురు సమయం", "CHAST": "చాథమ్ ప్రామాణిక సమయం", "CHUT": "చక్ సమయం", "CKT": "కుక్ దీవుల ప్రామాణిక సమయం", "CKT DST": "కుక్ దీవుల అర్ధ వేసవి సమయం", "CLST": "చిలీ వేసవి సమయం", "CLT": "చిలీ ప్రామాణిక సమయం", "COST": "కొలంబియా వేసవి సమయం", "COT": "కొలంబియా ప్రామాణిక సమయం", "CST": "మధ్యమ ప్రామాణిక సమయం", "CST China": "చైనా ప్రామాణిక సమయం", "CST China DST": "చైనా పగటి వెలుతురు సమయం", "CVST": "కేప్ వెర్డె వేసవి సమయం", "CVT": "కేప్ వెర్డె ప్రామాణిక సమయం", "CXT": "క్రిస్మస్ దీవి సమయం", "ChST": "చామర్రో ప్రామాణిక సమయం", "ChST NMI": "ChST NMI", "CuDT": "క్యూబా పగటి వెలుతురు సమయం", "CuST": "క్యూబా ప్రామాణిక సమయం", "DAVT": "డేవిస్ సమయం", "DDUT": "డ్యూమాంట్-డి’ఉర్విల్లే సమయం", "EASST": "ఈస్టర్ దీవి వేసవి సమయం", "EAST": "ఈస్టర్ దీవి ప్రామాణిక సమయం", "EAT": "తూర్పు ఆఫ్రికా సమయం", "ECT": "ఈక్వడార్ సమయం", "EDT": "తూర్పు పగటి వెలుతురు సమయం", "EGDT": "తూర్పు గ్రీన్\u200cల్యాండ్ వేసవి సమయం", "EGST": "తూర్పు గ్రీన్\u200cల్యాండ్ ప్రామాణిక సమయం", "EST": "తూర్పు ప్రామాణిక సమయం", "FEET": "సుదూర-తూర్పు యూరోపియన్ సమయం", "FJT": "ఫిజీ ప్రామాణిక సమయం", "FJT Summer": "ఫిజీ వేసవి సమయం", "FKST": "ఫాక్\u200cల్యాండ్ దీవుల వేసవి సమయం", "FKT": "ఫాక్\u200cల్యాండ్ దీవుల ప్రామాణిక సమయం", "FNST": "ఫెర్నాండో డి నొరోన్హా వేసవి సమయం", "FNT": "ఫెర్నాండో డి నొరోన్హా ప్రామాణిక సమయం", "GALT": "గాలాపాగోస్ సమయం", "GAMT": "గాంబియర్ సమయం", "GEST": "జార్జియా వేసవి సమయం", "GET": "జార్జియా ప్రామాణిక సమయం", "GFT": "ఫ్రెంచ్ గయానా సమయం", "GIT": "గిల్బర్ట్ దీవుల సమయం", "GMT": "గ్రీన్\u200cవిచ్ సగటు సమయం", "GNSST": "GNSST", "GNST": "GNST", "GST": "దక్షిణ జార్జియా సమయం", "GST Guam": "GST Guam", "GYT": "గయానా సమయం", "HADT": "హవాయ్-అల్యూషియన్ ప్రామాణిక సమయం", "HAST": "హవాయ్-అల్యూషియన్ ప్రామాణిక సమయం", "HKST": "హాంకాంగ్ వేసవి సమయం", "HKT": "హాంకాంగ్ ప్రామాణిక సమయం", "HOVST": "హోవ్డ్ వేసవి సమయం", "HOVT": "హోవ్డ్ ప్రామాణిక సమయం", "ICT": "ఇండోచైనా సమయం", "IDT": "ఇజ్రాయిల్ పగటి వెలుతురు సమయం", "IOT": "హిందూ మహా సముద్ర సమయం", "IRKST": "ఇర్కుట్స్క్ వేసవి సమయం", "IRKT": "ఇర్కుట్స్క్ ప్రామాణిక సమయం", "IRST": "ఇరాన్ ప్రామాణిక సమయం", "IRST DST": "ఇరాన్ పగటి వెలుతురు సమయం", "IST": "భారతదేశ ప్రామాణిక సమయం", "IST Israel": "ఇజ్రాయిల్ ప్రామాణిక సమయం", "JDT": "జపాన్ పగటి వెలుతురు సమయం", "JST": "జపాన్ ప్రామాణిక సమయం", "KOST": "కోస్రాయి సమయం", "KRAST": "క్రాస్నోయార్స్క్ వేసవి సమయం", "KRAT": "క్రాస్నోయార్స్క్ ప్రామాణిక సమయం", "KST": "కొరియన్ ప్రామాణిక సమయం", "KST DST": "కొరియన్ పగటి వెలుతురు సమయం", "LHDT": "లార్డ్ హోవ్ పగటి సమయం", "LHST": "లార్డ్ హోవ్ ప్రామాణిక సమయం", "LINT": "లైన్ దీవుల సమయం", "MAGST": "మగడాన్ వేసవి సమయం", "MAGT": "మగడాన్ ప్రామాణిక సమయం", "MART": "మార్క్వేసాస్ సమయం", "MAWT": "మాసన్ సమయం", "MDT": "MDT", "MESZ": "సెంట్రల్ యూరోపియన్ వేసవి సమయం", "MEZ": "సెంట్రల్ యూరోపియన్ ప్రామాణిక సమయం", "MHT": "మార్షల్ దీవుల సమయం", "MMT": "మయన్మార్ సమయం", "MSD": "మాస్కో వేసవి సమయం", "MST": "MST", "MUST": "మారిషస్ వేసవి సమయం", "MUT": "మారిషస్ ప్రామాణిక సమయం", "MVT": "మాల్దీవుల సమయం", "MYT": "మలేషియా సమయం", "NCT": "న్యూ కాలెడోనియా ప్రామాణిక సమయం", "NDT": "న్యూఫౌండ్\u200cల్యాండ్ పగటి వెలుతురు సమయం", "NDT New Caledonia": "న్యూ కాలెడోనియా వేసవి సమయం", "NFDT": "నార్ఫోక్ దీవి పగటి సమయం", "NFT": "నార్ఫోక్ దీవి ప్రామాణిక సమయం", "NOVST": "నోవోసిబిర్స్క్ వేసవి సమయం", "NOVT": "నోవోసిబిర్క్స్ ప్రామాణిక సమయం", "NPT": "నేపాల్ సమయం", "NRT": "నౌరు సమయం", "NST": "న్యూఫౌండ్\u200cల్యాండ్ ప్రామాణిక సమయం", "NUT": "నియూ సమయం", "NZDT": "న్యూజిల్యాండ్ పగటి వెలుతురు సమయం", "NZST": "న్యూజిల్యాండ్ ప్రామాణిక సమయం", "OESZ": "తూర్పు యూరోపియన్ వేసవి సమయం", "OEZ": "తూర్పు యూరోపియన్ ప్రామాణిక సమయం", "OMSST": "ఓమ్స్క్ వేసవి సమయం", "OMST": "ఓమ్స్క్ ప్రామాణిక సమయం", "PDT": "పసిఫిక్ పగటి వెలుతురు సమయం", "PDTM": "మెక్సికన్ పసిఫిక్ పగటి వెలుతురు సమయం", "PETDT": "PETDT", "PETST": "PETST", "PGT": "పాపువా న్యూ గినియా సమయం", "PHOT": "ఫినిక్స్ దీవుల సమయం", "PKT": "పాకిస్తాన్ ప్రామాణిక సమయం", "PKT DST": "పాకిస్తాన్ వేసవి సమయం", "PMDT": "సెయింట్ పియర్ మరియు మిక్వెలాన్ పగటి వెలుతురు సమయం", "PMST": "సెయింట్ పియెర్ మరియు మిక్వెలాన్ ప్రామాణిక సమయం", "PONT": "పొనేప్ సమయం", "PST": "పసిఫిక్ ప్రామాణిక సమయం", "PST Philippine": "ఫిలిప్పైన్ ప్రామాణిక సమయం", "PST Philippine DST": "ఫిలిప్పైన్ వేసవి సమయం", "PST Pitcairn": "పిట్\u200cకైర్న్ సమయం", "PSTM": "మెక్సికన్ పసిఫిక్ ప్రామాణిక సమయం", "PWT": "పాలావ్ సమయం", "PYST": "పరాగ్వే వేసవి సమయం", "PYT": "పరాగ్వే ప్రామాణిక సమయం", "PYT Korea": "ప్యోంగాంగ్ సమయం", "RET": "రీయూనియన్ సమయం", "ROTT": "రొతేరా సమయం", "SAKST": "సఖాలిన్ వేసవి సమయం", "SAKT": "సఖాలిన్ ప్రామాణిక సమయం", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "దక్షిణ ఆఫ్రికా ప్రామాణిక సమయం", "SBT": "సోలమన్ దీవుల సమయం", "SCT": "సీషెల్స్ సమయం", "SGT": "సింగపూర్ ప్రామాణిక సమయం", "SLST": "SLST", "SRT": "సూరినామ్ సమయం", "SST Samoa": "సమోవా ప్రామాణిక సమయం", "SST Samoa Apia": "ఏపియా ప్రామాణిక సమయం", "SST Samoa Apia DST": "ఏపియా పగటి సమయం", "SST Samoa DST": "సమోవా పగటి వెలుతురు సమయం", "SYOT": "స్యోవా సమయం", "TAAF": "ఫ్రెంచ్ దక్షిణ మరియు అంటార్కిటిక్ సమయం", "TAHT": "తహితి సమయం", "TJT": "తజికిస్తాన్ సమయం", "TKT": "టోకెలావ్ సమయం", "TLT": "తూర్పు తైమూర్ సమయం", "TMST": "తుర్క్\u200cమెనిస్తాన్ వేసవి సమయం", "TMT": "తుర్క్\u200cమెనిస్తాన్ ప్రామాణిక సమయం", "TOST": "టాంగా వేసవి సమయం", "TOT": "టాంగా ప్రామాణిక సమయం", "TVT": "తువాలు సమయం", "TWT": "తైపీ ప్రామాణిక సమయం", "TWT DST": "తైపీ పగటి వెలుతురు సమయం", "ULAST": "ఉలన్ బతోర్ వేసవి సమయం", "ULAT": "ఉలన్ బతోర్ ప్రామాణిక సమయం", "UYST": "ఉరుగ్వే వేసవి సమయం", "UYT": "ఉరుగ్వే ప్రామాణిక సమయం", "UZT": "ఉజ్బెకిస్తాన్ ప్రామాణిక సమయం", "UZT DST": "ఉజ్బెకిస్తాన్ వేసవి సమయం", "VET": "వెనిజులా సమయం", "VLAST": "వ్లాడివోస్టోక్ వేసవి సమయం", "VLAT": "వ్లాడివోస్టోక్ ప్రామాణిక సమయం", "VOLST": "వోల్గోగ్రాడ్ వేసవి సమయం", "VOLT": "వోల్గోగ్రాడ్ ప్రామాణిక సమయం", "VOST": "వోస్టోక్ సమయం", "VUT": "వనౌటు ప్రామాణిక సమయం", "VUT DST": "వనౌటు వేసవి సమయం", "WAKT": "వేక్ దీవి సమయం", "WARST": "పశ్చిమ అర్జెంటీనా వేసవి సమయం", "WART": "పశ్చిమ అర్జెంటీనా ప్రామాణిక సమయం", "WAST": "పశ్చిమ ఆఫ్రికా సమయం", "WAT": "పశ్చిమ ఆఫ్రికా సమయం", "WESZ": "పశ్చిమ యూరోపియన్ వేసవి సమయం", "WEZ": "పశ్చిమ యూరోపియన్ ప్రామాణిక సమయం", "WFT": "వాలీస్ మరియు ఫుటునా సమయం", "WGST": "పశ్చిమ గ్రీన్\u200cల్యాండ్ వేసవి సమయం", "WGT": "పశ్చిమ గ్రీన్\u200cల్యాండ్ ప్రామాణిక సమయం", "WIB": "పశ్చిమ ఇండోనేషియా సమయం", "WIT": "తూర్పు ఇండోనేషియా సమయం", "WITA": "సెంట్రల్ ఇండోనేషియా సమయం", "YAKST": "యాకుట్స్క్ వేసవి సమయం", "YAKT": "యాకుట్స్క్ ప్రామాణిక సమయం", "YEKST": "యెకటెరిన్\u200cబర్గ్ వేసవి సమయం", "YEKT": "యెకటెరిన్\u200cబర్గ్ ప్రామాణిక సమయం", "YST": "YST", "МСК": "మాస్కో ప్రామాణిక సమయం", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "పశ్చిమ కజకిస్తాన్ సమయం", "شىعىش قازاق ەلى": "తూర్పు కజకి\u200cస్తాన్ సమయం", "قازاق ەلى": "కజకి\u200cస్తాన్ సమయం", "قىرعىزستان": "కిర్గిస్తాన్ సమయం", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "పెరూ వేసవి సమయం"},
	}
}

// Locale returns the current translators string locale
func (kxv *kxv_Telu) Locale() string {
	return kxv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kxv_Telu'
func (kxv *kxv_Telu) PluralsCardinal() []locales.PluralRule {
	return kxv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kxv_Telu'
func (kxv *kxv_Telu) PluralsOrdinal() []locales.PluralRule {
	return kxv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kxv_Telu'
func (kxv *kxv_Telu) PluralsRange() []locales.PluralRule {
	return kxv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Telu'
func (kxv *kxv_Telu) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Telu'
func (kxv *kxv_Telu) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kxv_Telu'
func (kxv *kxv_Telu) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kxv *kxv_Telu) MonthAbbreviated(month time.Month) string {
	return kxv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kxv *kxv_Telu) MonthsAbbreviated() []string {
	return kxv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kxv *kxv_Telu) MonthNarrow(month time.Month) string {
	return kxv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kxv *kxv_Telu) MonthsNarrow() []string {
	return kxv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kxv *kxv_Telu) MonthWide(month time.Month) string {
	return kxv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kxv *kxv_Telu) MonthsWide() []string {
	return kxv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kxv *kxv_Telu) WeekdayAbbreviated(weekday time.Weekday) string {
	return kxv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kxv *kxv_Telu) WeekdaysAbbreviated() []string {
	return kxv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kxv *kxv_Telu) WeekdayNarrow(weekday time.Weekday) string {
	return kxv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kxv *kxv_Telu) WeekdaysNarrow() []string {
	return kxv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kxv *kxv_Telu) WeekdayShort(weekday time.Weekday) string {
	return kxv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kxv *kxv_Telu) WeekdaysShort() []string {
	return kxv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kxv *kxv_Telu) WeekdayWide(weekday time.Weekday) string {
	return kxv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kxv *kxv_Telu) WeekdaysWide() []string {
	return kxv.daysWide
}

// Decimal returns the decimal point of number
func (kxv *kxv_Telu) Decimal() string {
	return kxv.decimal
}

// Group returns the group of number
func (kxv *kxv_Telu) Group() string {
	return kxv.group
}

// Group returns the minus sign of number
func (kxv *kxv_Telu) Minus() string {
	return kxv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kxv_Telu' and handles both Whole and Real numbers based on 'v'
func (kxv *kxv_Telu) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kxv.group[0])
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
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kxv_Telu' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kxv *kxv_Telu) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kxv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kxv.group[0])
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

	for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kxv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Telu'
// in accounting notation.
func (kxv *kxv_Telu) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == groupThreshold {
				b = append(b, kxv.group[0])
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

		for j := len(kxv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyNegativePrefix[j])
		}

		b = append(b, kxv.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, kxv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kxv_Telu'
func (kxv *kxv_Telu) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kxv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
