package hy_AM

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type hy_AM struct {
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

// New returns a new instance of translator for the 'hy_AM' locale
func New() locales.Translator {
	return &hy_AM{
		locale:                 "hy_AM",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "հնվ", "փտվ", "մրտ", "ապր", "մյս", "հնս", "հլս", "օգս", "սեպ", "հոկ", "նոյ", "դեկ"},
		monthsNarrow:           []string{"", "Հ", "Փ", "Մ", "Ա", "Մ", "Հ", "Հ", "Օ", "Ս", "Հ", "Ն", "Դ"},
		monthsWide:             []string{"", "հունվարի", "փետրվարի", "մարտի", "ապրիլի", "մայիսի", "հունիսի", "հուլիսի", "օգոստոսի", "սեպտեմբերի", "հոկտեմբերի", "նոյեմբերի", "դեկտեմբերի"},
		daysAbbreviated:        []string{"կիր", "երկ", "երք", "չրք", "հնգ", "ուր", "շբթ"},
		daysNarrow:             []string{"Կ", "Ե", "Ե", "Չ", "Հ", "Ո", "Շ"},
		daysShort:              []string{"կր", "եկ", "եք", "չք", "հգ", "ու", "շբ"},
		daysWide:               []string{"կիրակի", "երկուշաբթի", "երեքշաբթի", "չորեքշաբթի", "հինգշաբթի", "ուրբաթ", "շաբաթ"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"ա", "հ"},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"մ.թ.ա.", "մ.թ."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Քրիստոսից առաջ", "Քրիստոսից հետո"},
		timezones:              map[string]string{"ACDT": "Կենտրոնական Ավստրալիայի ամառային ժամանակ", "ACST": "Կենտրոնական Ավստրալիայի ստանդարտ ժամանակ", "ACT": "ACT", "ACWDT": "Կենտրոնական Ավստրալիայի արևմտյան ամառային ժամանակ", "ACWST": "Կենտրոնական Ավստրալիայի արևմտյան ստանդարտ ժամանակ", "ADT": "Ատլանտյան ամառային ժամանակ", "ADT Arabia": "Սաուդյան Արաբիայի ամառային ժամանակ", "AEDT": "Արևելյան Ավստրալիայի ամառային ժամանակ", "AEST": "Արևելյան Ավստրալիայի ստանդարտ ժամանակ", "AFT": "Աֆղանստանի ժամանակ", "AKDT": "Ալյասկայի ամառային ժամանակ", "AKST": "Ալյասկայի ստանդարտ ժամանակ", "AMST": "Ամազոնյան ամառային ժամանակ", "AMST Armenia": "Հայաստանի ամառային ժամանակ", "AMT": "Ամազոնյան ստանդարտ ժամանակ", "AMT Armenia": "Հայաստանի ստանդարտ ժամանակ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Արգենտինայի ամառային ժամանակ", "ART": "Արգենտինայի ստնադարտ ժամանակ", "AST": "Ատլանտյան ստանդարտ ժամանակ", "AST Arabia": "Սաուդյան Արաբիայի ստանդարտ ժամանակ", "AWDT": "Արևմտյան Ավստրալիայի ամառային ժամանակ", "AWST": "Արևմտյան Ավստրալիայի ստանդարտ ժամանակ", "AZST": "Ադրբեջանի ամառային ժամանակ", "AZT": "Ադրբեջանի ստանդարտ ժամանակ", "BDT Bangladesh": "Բանգլադեշի ամառային ժամանակ", "BNT": "Բրունեյի ժամանակ", "BOT": "Բոլիվիայի ժամանակ", "BRST": "Բրազիլիայի ամառային ժամանակ", "BRT": "Բրազիլիայի ստանդարտ ժամանակ", "BST Bangladesh": "Բանգլադեշի ստանդարտ ժամանակ", "BT": "Բութանի ժամանակ", "CAST": "CAST", "CAT": "Կենտրոնական Աֆրիկայի ժամանակ", "CCT": "Կոկոսյան կղզիների ժամանակ", "CDT": "Կենտրոնական Ամերիկայի ամառային ժամանակ", "CHADT": "Չաթեմ կղզու ամառային ժամանակ", "CHAST": "Չաթեմ կղզու ստանդարտ ժամանակ", "CHUT": "Տրուկի ժամանակ", "CKT": "Կուկի կղզիների ստանդարտ ժամանակ", "CKT DST": "Կուկի կղզիների կիսաամառային ժամանակ", "CLST": "Չիլիի ամառային ժամանակ", "CLT": "Չիլիի ստանդարտ ժամանակ", "COST": "Կոլումբիայի ամառային ժամանակ", "COT": "Կոլումբիայի ստանդարտ ժամանակ", "CST": "Կենտրոնական Ամերիկայի ստանդարտ ժամանակ", "CST China": "Չինաստանի ստանդարտ ժամանակ", "CST China DST": "Չինաստանի ամառային ժամանակ", "CVST": "Կաբո Վերդեի ամառային ժամանակ", "CVT": "Կաբո Վերդեի ստանդարտ ժամանակ", "CXT": "Սուրբ Ծննդյան կղզու ժամանակ", "ChST": "Չամոռոյի ժամանակ", "ChST NMI": "ChST NMI", "CuDT": "Կուբայի ամառային ժամանակ", "CuST": "Կուբայի ստանդարտ ժամանակ", "DAVT": "Դեյվիսի ժամանակ", "DDUT": "Դյումոն դ’Յուրվիլի ժամանակ", "EASST": "Զատկի կղզու ամառային ժամանակ", "EAST": "Զատկի կղզու ստանդարտ ժամանակ", "EAT": "Արևելյան Աֆրիկայի ժամանակ", "ECT": "Էկվադորի ժամանակ", "EDT": "Արևելյան Ամերիկայի ամառային ժամանակ", "EGDT": "Արևելյան Գրենլանդիայի ամառային ժամանակ", "EGST": "Արևելյան Գրենլանդիայի ստանդարտ ժամանակ", "EST": "Արևելյան Ամերիկայի ստանդարտ ժամանակ", "FEET": "Մինսկի ժամանակ", "FJT": "Ֆիջիի ստանդարտ ժամանակ", "FJT Summer": "Ֆիջիի ամառային ժամանակ", "FKST": "Ֆոլքլենդյան կղզիների ամառային ժամանակ", "FKT": "Ֆոլքլենդյան կղզիների ստանդարտ ժամանակ", "FNST": "Ֆերնանդու դի Նորոնյայի ամառային ժամանակ", "FNT": "Ֆերնանդու դի Նորոնյայի ստանդարտ ժամանակ", "GALT": "Գալապագոսյան կղզիների ժամանակ", "GAMT": "Գամբյե կղզիների ժամանակ", "GEST": "Վրաստանի ամառային ժամանակ", "GET": "Վրաստանի ստանդարտ ժամանակ", "GFT": "Ֆրանսիական Գվիանայի ժամանակ", "GIT": "Ջիլբերթի կղզիների ժամանակ", "GMT": "Գրինվիչի ժամանակ", "GNSST": "GNSST", "GNST": "GNST", "GST": "Հարավային Ջորջիայի ժամանակ", "GST Guam": "GST Guam", "GYT": "Գայանայի ժամանակ", "HADT": "Հավայան-ալեության ամառային ժամանակ", "HAST": "Հավայան-ալեության ստանդարտ ժամանակ", "HKST": "Հոնկոնգի ամառային ժամանակ", "HKT": "Հոնկոնգի ստանդարտ ժամանակ", "HOVST": "Հովդի ամառային ժամանակ", "HOVT": "Հովդի ստանդարտ ժամանակ", "ICT": "Հնդկաչինական ժամանակ", "IDT": "Իսրայելի ամառային ժամանակ", "IOT": "Հնդկական օվկիանոսի ժամանակ", "IRKST": "Իրկուտսկի ամառային ժամանակ", "IRKT": "Իրկուտսկի ստանդարտ ժամանակ", "IRST": "Իրանի ստանդարտ ժամանակ", "IRST DST": "Իրանի ամառային ժամանակ", "IST": "Հնդկաստանի ստանդարտ ժամանակ", "IST Israel": "Իսրայելի ստանդարտ ժամանակ", "JDT": "Ճապոնիայի ամառային ժամանակ", "JST": "Ճապոնիայի ստանդարտ ժամանակ", "KOST": "Կոսրաեյի ժամանակ", "KRAST": "Կրասնոյարսկի ամառային ժամանակ", "KRAT": "Կրասնոյարսկի ստանդարտ ժամանակ", "KST": "Կորեայի ստանդարտ ժամանակ", "KST DST": "Կորեայի ամառային ժամանակ", "LHDT": "Լորդ Հաուի ամառային ժամանակ", "LHST": "Լորդ Հաուի ստանդարտ ժամանակ", "LINT": "Լայն կղզիների ժամանակ", "MAGST": "Մագադանի ամառային ժամանակ", "MAGT": "Մագադանի ստանդարտ ժամանակ", "MART": "Մարկիզյան կղզիների ժամանակ", "MAWT": "Մոուսոնի ժամանակ", "MDT": "Լեռնային ամառային ժամանակ (ԱՄՆ)", "MESZ": "Կենտրոնական Եվրոպայի ամառային ժամանակ", "MEZ": "Կենտրոնական Եվրոպայի ստանդարտ ժամանակ", "MHT": "Մարշալյան կղզիների ժամանակ", "MMT": "Մյանմայի ժամանակ", "MSD": "Մոսկվայի ամառային ժամանակ", "MST": "Լեռնային ստանդարտ ժամանակ (ԱՄՆ)", "MUST": "Մավրիկիոսի ամառային ժամանակ", "MUT": "Մավրիկիոսի ստանդարտ ժամանակ", "MVT": "Մալդիվների ժամանակ", "MYT": "Մալայզիայի ժամանակ", "NCT": "Նոր Կալեդոնիայի ստանդարտ ժամանակ", "NDT": "Նյուֆաունդլենդի ամառային ժամանակ", "NDT New Caledonia": "Նոր Կալեդոնիայի ամառային ժամանակ", "NFDT": "Նորֆոլկ կղզու ամառային ժամանակ", "NFT": "Նորֆոլկ կղզու ստանդարտ ժամանակ", "NOVST": "Նովոսիբիրսկի ամառային ժամանակ", "NOVT": "Նովոսիբիրսկի ստանդարտ ժամանակ", "NPT": "Նեպալի ժամանակ", "NRT": "Նաուրուի ժամանակ", "NST": "Նյուֆաունդլենդի ստանդարտ ժամանակ", "NUT": "Նիուեյի ժամանակ", "NZDT": "Նոր Զելանդիայի ամառային ժամանակ", "NZST": "Նոր Զելանդիայի ստանդարտ ժամանակ", "OESZ": "Արևելյան Եվրոպայի ամառային ժամանակ", "OEZ": "Արևելյան Եվրոպայի ստանդարտ ժամանակ", "OMSST": "Օմսկի ամառային ժամանակ", "OMST": "Օմսկի ստանդարտ ժամանակ", "PDT": "Խաղաղօվկիանոսյան ամառային ժամանակ", "PDTM": "Մեքսիկայի խաղաղօվկիանոսյան ամառային ժամանակ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Պապուա Նոր Գվինեայի ժամանակ", "PHOT": "Ֆինիքս կղզիների ժամանակ", "PKT": "Պակիստանի ստանդարտ ժամանակ", "PKT DST": "Պակիստանի ամառային ժամանակ", "PMDT": "Սեն Պիեռ և Միքելոնի ամառային ժամանակ", "PMST": "Սեն Պիեռ և Միքելոնի ստանդարտ ժամանակ", "PONT": "Պոնապե կղզու ժամանակ", "PST": "Խաղաղօվկիանոսյան ստանդարտ ժամանակ", "PST Philippine": "Ֆիլիպինների ստանդարտ ժամանակ", "PST Philippine DST": "Ֆիլիպինների ամառային ժամանակ", "PST Pitcairn": "Պիտկեռնի ժամանակ", "PSTM": "Մեքսիկայի խաղաղօվկիանոսյան ստանդարտ ժամանակ", "PWT": "Պալաույի ժամանակ", "PYST": "Պարագվայի ամառային ժամանակ", "PYT": "Պարագվայի ստանդարտ ժամանակ", "PYT Korea": "Փխենյանի ժամանակ", "RET": "Ռեյունիոնի ժամանակ", "ROTT": "Ռոտերայի ժամանակ", "SAKST": "Սախալինի ամառային ժամանակ", "SAKT": "Սախալինի ստանդարտ ժամանակ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Հարավային Աֆրիկայի ժամանակ", "SBT": "Սողոմոնի կղզիների ժամանակ", "SCT": "Սեյշելյան կղզիների ժամանակ", "SGT": "Սինգապուրի ժամանակ", "SLST": "SLST", "SRT": "Սուրինամի ժամանակ", "SST Samoa": "Սամոայի ստանդարտ ժամանակ", "SST Samoa Apia": "Ապիայի ստանդարտ ժամանակ", "SST Samoa Apia DST": "Ապիայի ամառային ժամանակ", "SST Samoa DST": "Սամոայի ամառային ժամանակ", "SYOT": "Սյովայի ժամանակ", "TAAF": "Ֆրանսիական հարավային և անտարկտիդյան ժամանակ", "TAHT": "Թաիթիի ժամանակ", "TJT": "Տաջիկստանի ժամանակ", "TKT": "Տոկելաույի ժամանակ", "TLT": "Արևելյան Թիմորի ժամանակ", "TMST": "Թուրքմենստանի ամառային ժամանակ", "TMT": "Թուրքմենստանի ստանդարտ ժամանակ", "TOST": "Տոնգայի ամառային ժամանակ", "TOT": "Տոնգայի ստանդարտ ժամանակ", "TVT": "Տուվալույի ժամանակ", "TWT": "Թայպեյի ստանդարտ ժամանակ", "TWT DST": "Թայպեյի ամառային ժամանակ", "ULAST": "Ուլան Բատորի ամառային ժամանակ", "ULAT": "Ուլան Բատորի ստանդարտ ժամանակ", "UYST": "Ուրուգվայի ամառային ժամանակ", "UYT": "Ուրուգվայի ստանդարտ ժամանակ", "UZT": "Ուզբեկստանի ստանդարտ ժամանակ", "UZT DST": "Ուզբեկստանի ամառային ժամանակ", "VET": "Վենեսուելայի ժամանակ", "VLAST": "Վլադիվոստոկի ամառային ժամանակ", "VLAT": "Վլադիվոստոկի ստանդարտ ժամանակ", "VOLST": "Վոլգոգրադի ամառային ժամանակ", "VOLT": "Վոլգոգրադի ստանդարտ ժամանակ", "VOST": "Վոստոկի ժամանակ", "VUT": "Վանուատույի ստանդարտ ժամանակ", "VUT DST": "Վանուատույի ամառային ժամանակ", "WAKT": "Ուեյք կղզու ժամանակ", "WARST": "Արևմտյան Արգենտինայի ամառային ժամանակ", "WART": "Արևմտյան Արգենտինայի ստնադարտ ժամանակ", "WAST": "Արևմտյան Աֆրիկայի ժամանակ", "WAT": "Արևմտյան Աֆրիկայի ժամանակ", "WESZ": "Արևմտյան Եվրոպայի ամառային ժամանակ", "WEZ": "Արևմտյան Եվրոպայի ստանդարտ ժամանակ", "WFT": "Ուոլիս և Ֆուտունայի ժամանակ", "WGST": "Արևմտյան Գրենլանդիայի ամառային ժամանակ", "WGT": "Արևմտյան Գրենլանդիայի ստանդարտ ժամանակ", "WIB": "Արևմտյան Ինդոնեզիայի ժամանակ", "WIT": "Արևելյան Ինդոնեզիայի ժամանակ", "WITA": "Կենտրոնական Ինդոնեզիայի ժամանակ", "YAKST": "Յակուտսկի ամառային ժամանակ", "YAKT": "Յակուտսկի ստանդարտ ժամանակ", "YEKST": "Եկատերինբուրգի ամառային ժամանակ", "YEKT": "Եկատերինբուրգի ստանդարտ ժամանակ", "YST": "Յուկոնի ժամանակ", "МСК": "Մոսկվայի ստանդարտ ժամանակ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Արևմտյան Ղազախստանի ժամանակ", "شىعىش قازاق ەلى": "Արևելյան Ղազախստանի ժամանակ", "قازاق ەلى": "Ղազախստանի ժամանակ", "قىرعىزستان": "Ղրղզստանի ժամանակ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Պերուի ամառային ժամանակ"},
	}
}

// Locale returns the current translators string locale
func (hy *hy_AM) Locale() string {
	return hy.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hy_AM'
func (hy *hy_AM) PluralsCardinal() []locales.PluralRule {
	return hy.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hy_AM'
func (hy *hy_AM) PluralsOrdinal() []locales.PluralRule {
	return hy.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'hy_AM'
func (hy *hy_AM) PluralsRange() []locales.PluralRule {
	return hy.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hy_AM'
func (hy *hy_AM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hy_AM'
func (hy *hy_AM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hy_AM'
func (hy *hy_AM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := hy.CardinalPluralRule(num1, v1)
	end := hy.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (hy *hy_AM) MonthAbbreviated(month time.Month) string {
	return hy.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (hy *hy_AM) MonthsAbbreviated() []string {
	return hy.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (hy *hy_AM) MonthNarrow(month time.Month) string {
	return hy.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (hy *hy_AM) MonthsNarrow() []string {
	return hy.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (hy *hy_AM) MonthWide(month time.Month) string {
	return hy.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (hy *hy_AM) MonthsWide() []string {
	return hy.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (hy *hy_AM) WeekdayAbbreviated(weekday time.Weekday) string {
	return hy.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (hy *hy_AM) WeekdaysAbbreviated() []string {
	return hy.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (hy *hy_AM) WeekdayNarrow(weekday time.Weekday) string {
	return hy.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (hy *hy_AM) WeekdaysNarrow() []string {
	return hy.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (hy *hy_AM) WeekdayShort(weekday time.Weekday) string {
	return hy.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (hy *hy_AM) WeekdaysShort() []string {
	return hy.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (hy *hy_AM) WeekdayWide(weekday time.Weekday) string {
	return hy.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (hy *hy_AM) WeekdaysWide() []string {
	return hy.daysWide
}

// Decimal returns the decimal point of number
func (hy *hy_AM) Decimal() string {
	return hy.decimal
}

// Group returns the group of number
func (hy *hy_AM) Group() string {
	return hy.group
}

// Group returns the minus sign of number
func (hy *hy_AM) Minus() string {
	return hy.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hy_AM' and handles both Whole and Real numbers based on 'v'
func (hy *hy_AM) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hy_AM' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hy *hy_AM) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hy_AM'
func (hy *hy_AM) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hy.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(hy.group) - 1; j >= 0; j-- {
					b = append(b, hy.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, hy.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hy_AM'
// in accounting notation.
func (hy *hy_AM) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hy.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(hy.group) - 1; j >= 0; j-- {
					b = append(b, hy.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, hy.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, hy.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, hy.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, hy.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd5, 0xa9, 0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, hy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd5, 0xa9, 0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd5, 0xa9, 0x2e, 0x20}...)
	b = append(b, hy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, hy.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'hy_AM'
func (hy *hy_AM) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hy.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
