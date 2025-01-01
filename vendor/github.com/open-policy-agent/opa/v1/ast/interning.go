// Copyright 2024 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package ast

import "strconv"

// NOTE! Great care must be taken **not** to modify the terms returned
// from these functions, as they are shared across all callers.

var (
	booleanTrueTerm  = &Term{Value: Boolean(true)}
	booleanFalseTerm = &Term{Value: Boolean(false)}

	// since this is by far the most common negative number
	minusOneTerm = &Term{Value: Number("-1")}
)

// InternedBooleanTerm returns an interned term with the given boolean value.
func InternedBooleanTerm(b bool) *Term {
	if b {
		return booleanTrueTerm
	}

	return booleanFalseTerm
}

// InternedIntNumberTerm returns a term with the given integer value. The term is
// cached between -1 to 512, and for values outside of that range, this function
// is equivalent to ast.IntNumberTerm.
func InternedIntNumberTerm(i int) *Term {
	if i >= 0 && i < len(intNumberTerms) {
		return intNumberTerms[i]
	}

	if i == -1 {
		return minusOneTerm
	}

	return &Term{Value: Number(strconv.Itoa(i))}
}

// InternedIntFromString returns a term with the given integer value if the string
// maps to an interned term. If the string does not map to an interned term, nil is
// returned.
func InternedIntNumberTermFromString(s string) *Term {
	if term, ok := stringToIntNumberTermMap[s]; ok {
		return term
	}

	return nil
}

// HasInternedIntNumberTerm returns true if the given integer value maps to an interned
// term, otherwise false.
func HasInternedIntNumberTerm(i int) bool {
	return i >= -1 && i < len(intNumberTerms)
}

var stringToIntNumberTermMap = map[string]*Term{
	"-1":  minusOneTerm,
	"0":   intNumberTerms[0],
	"1":   intNumberTerms[1],
	"2":   intNumberTerms[2],
	"3":   intNumberTerms[3],
	"4":   intNumberTerms[4],
	"5":   intNumberTerms[5],
	"6":   intNumberTerms[6],
	"7":   intNumberTerms[7],
	"8":   intNumberTerms[8],
	"9":   intNumberTerms[9],
	"10":  intNumberTerms[10],
	"11":  intNumberTerms[11],
	"12":  intNumberTerms[12],
	"13":  intNumberTerms[13],
	"14":  intNumberTerms[14],
	"15":  intNumberTerms[15],
	"16":  intNumberTerms[16],
	"17":  intNumberTerms[17],
	"18":  intNumberTerms[18],
	"19":  intNumberTerms[19],
	"20":  intNumberTerms[20],
	"21":  intNumberTerms[21],
	"22":  intNumberTerms[22],
	"23":  intNumberTerms[23],
	"24":  intNumberTerms[24],
	"25":  intNumberTerms[25],
	"26":  intNumberTerms[26],
	"27":  intNumberTerms[27],
	"28":  intNumberTerms[28],
	"29":  intNumberTerms[29],
	"30":  intNumberTerms[30],
	"31":  intNumberTerms[31],
	"32":  intNumberTerms[32],
	"33":  intNumberTerms[33],
	"34":  intNumberTerms[34],
	"35":  intNumberTerms[35],
	"36":  intNumberTerms[36],
	"37":  intNumberTerms[37],
	"38":  intNumberTerms[38],
	"39":  intNumberTerms[39],
	"40":  intNumberTerms[40],
	"41":  intNumberTerms[41],
	"42":  intNumberTerms[42],
	"43":  intNumberTerms[43],
	"44":  intNumberTerms[44],
	"45":  intNumberTerms[45],
	"46":  intNumberTerms[46],
	"47":  intNumberTerms[47],
	"48":  intNumberTerms[48],
	"49":  intNumberTerms[49],
	"50":  intNumberTerms[50],
	"51":  intNumberTerms[51],
	"52":  intNumberTerms[52],
	"53":  intNumberTerms[53],
	"54":  intNumberTerms[54],
	"55":  intNumberTerms[55],
	"56":  intNumberTerms[56],
	"57":  intNumberTerms[57],
	"58":  intNumberTerms[58],
	"59":  intNumberTerms[59],
	"60":  intNumberTerms[60],
	"61":  intNumberTerms[61],
	"62":  intNumberTerms[62],
	"63":  intNumberTerms[63],
	"64":  intNumberTerms[64],
	"65":  intNumberTerms[65],
	"66":  intNumberTerms[66],
	"67":  intNumberTerms[67],
	"68":  intNumberTerms[68],
	"69":  intNumberTerms[69],
	"70":  intNumberTerms[70],
	"71":  intNumberTerms[71],
	"72":  intNumberTerms[72],
	"73":  intNumberTerms[73],
	"74":  intNumberTerms[74],
	"75":  intNumberTerms[75],
	"76":  intNumberTerms[76],
	"77":  intNumberTerms[77],
	"78":  intNumberTerms[78],
	"79":  intNumberTerms[79],
	"80":  intNumberTerms[80],
	"81":  intNumberTerms[81],
	"82":  intNumberTerms[82],
	"83":  intNumberTerms[83],
	"84":  intNumberTerms[84],
	"85":  intNumberTerms[85],
	"86":  intNumberTerms[86],
	"87":  intNumberTerms[87],
	"88":  intNumberTerms[88],
	"89":  intNumberTerms[89],
	"90":  intNumberTerms[90],
	"91":  intNumberTerms[91],
	"92":  intNumberTerms[92],
	"93":  intNumberTerms[93],
	"94":  intNumberTerms[94],
	"95":  intNumberTerms[95],
	"96":  intNumberTerms[96],
	"97":  intNumberTerms[97],
	"98":  intNumberTerms[98],
	"99":  intNumberTerms[99],
	"100": intNumberTerms[100],
	"101": intNumberTerms[101],
	"102": intNumberTerms[102],
	"103": intNumberTerms[103],
	"104": intNumberTerms[104],
	"105": intNumberTerms[105],
	"106": intNumberTerms[106],
	"107": intNumberTerms[107],
	"108": intNumberTerms[108],
	"109": intNumberTerms[109],
	"110": intNumberTerms[110],
	"111": intNumberTerms[111],
	"112": intNumberTerms[112],
	"113": intNumberTerms[113],
	"114": intNumberTerms[114],
	"115": intNumberTerms[115],
	"116": intNumberTerms[116],
	"117": intNumberTerms[117],
	"118": intNumberTerms[118],
	"119": intNumberTerms[119],
	"120": intNumberTerms[120],
	"121": intNumberTerms[121],
	"122": intNumberTerms[122],
	"123": intNumberTerms[123],
	"124": intNumberTerms[124],
	"125": intNumberTerms[125],
	"126": intNumberTerms[126],
	"127": intNumberTerms[127],
	"128": intNumberTerms[128],
	"129": intNumberTerms[129],
	"130": intNumberTerms[130],
	"131": intNumberTerms[131],
	"132": intNumberTerms[132],
	"133": intNumberTerms[133],
	"134": intNumberTerms[134],
	"135": intNumberTerms[135],
	"136": intNumberTerms[136],
	"137": intNumberTerms[137],
	"138": intNumberTerms[138],
	"139": intNumberTerms[139],
	"140": intNumberTerms[140],
	"141": intNumberTerms[141],
	"142": intNumberTerms[142],
	"143": intNumberTerms[143],
	"144": intNumberTerms[144],
	"145": intNumberTerms[145],
	"146": intNumberTerms[146],
	"147": intNumberTerms[147],
	"148": intNumberTerms[148],
	"149": intNumberTerms[149],
	"150": intNumberTerms[150],
	"151": intNumberTerms[151],
	"152": intNumberTerms[152],
	"153": intNumberTerms[153],
	"154": intNumberTerms[154],
	"155": intNumberTerms[155],
	"156": intNumberTerms[156],
	"157": intNumberTerms[157],
	"158": intNumberTerms[158],
	"159": intNumberTerms[159],
	"160": intNumberTerms[160],
	"161": intNumberTerms[161],
	"162": intNumberTerms[162],
	"163": intNumberTerms[163],
	"164": intNumberTerms[164],
	"165": intNumberTerms[165],
	"166": intNumberTerms[166],
	"167": intNumberTerms[167],
	"168": intNumberTerms[168],
	"169": intNumberTerms[169],
	"170": intNumberTerms[170],
	"171": intNumberTerms[171],
	"172": intNumberTerms[172],
	"173": intNumberTerms[173],
	"174": intNumberTerms[174],
	"175": intNumberTerms[175],
	"176": intNumberTerms[176],
	"177": intNumberTerms[177],
	"178": intNumberTerms[178],
	"179": intNumberTerms[179],
	"180": intNumberTerms[180],
	"181": intNumberTerms[181],
	"182": intNumberTerms[182],
	"183": intNumberTerms[183],
	"184": intNumberTerms[184],
	"185": intNumberTerms[185],
	"186": intNumberTerms[186],
	"187": intNumberTerms[187],
	"188": intNumberTerms[188],
	"189": intNumberTerms[189],
	"190": intNumberTerms[190],
	"191": intNumberTerms[191],
	"192": intNumberTerms[192],
	"193": intNumberTerms[193],
	"194": intNumberTerms[194],
	"195": intNumberTerms[195],
	"196": intNumberTerms[196],
	"197": intNumberTerms[197],
	"198": intNumberTerms[198],
	"199": intNumberTerms[199],
	"200": intNumberTerms[200],
	"201": intNumberTerms[201],
	"202": intNumberTerms[202],
	"203": intNumberTerms[203],
	"204": intNumberTerms[204],
	"205": intNumberTerms[205],
	"206": intNumberTerms[206],
	"207": intNumberTerms[207],
	"208": intNumberTerms[208],
	"209": intNumberTerms[209],
	"210": intNumberTerms[210],
	"211": intNumberTerms[211],
	"212": intNumberTerms[212],
	"213": intNumberTerms[213],
	"214": intNumberTerms[214],
	"215": intNumberTerms[215],
	"216": intNumberTerms[216],
	"217": intNumberTerms[217],
	"218": intNumberTerms[218],
	"219": intNumberTerms[219],
	"220": intNumberTerms[220],
	"221": intNumberTerms[221],
	"222": intNumberTerms[222],
	"223": intNumberTerms[223],
	"224": intNumberTerms[224],
	"225": intNumberTerms[225],
	"226": intNumberTerms[226],
	"227": intNumberTerms[227],
	"228": intNumberTerms[228],
	"229": intNumberTerms[229],
	"230": intNumberTerms[230],
	"231": intNumberTerms[231],
	"232": intNumberTerms[232],
	"233": intNumberTerms[233],
	"234": intNumberTerms[234],
	"235": intNumberTerms[235],
	"236": intNumberTerms[236],
	"237": intNumberTerms[237],
	"238": intNumberTerms[238],
	"239": intNumberTerms[239],
	"240": intNumberTerms[240],
	"241": intNumberTerms[241],
	"242": intNumberTerms[242],
	"243": intNumberTerms[243],
	"244": intNumberTerms[244],
	"245": intNumberTerms[245],
	"246": intNumberTerms[246],
	"247": intNumberTerms[247],
	"248": intNumberTerms[248],
	"249": intNumberTerms[249],
	"250": intNumberTerms[250],
	"251": intNumberTerms[251],
	"252": intNumberTerms[252],
	"253": intNumberTerms[253],
	"254": intNumberTerms[254],
	"255": intNumberTerms[255],
	"256": intNumberTerms[256],
	"257": intNumberTerms[257],
	"258": intNumberTerms[258],
	"259": intNumberTerms[259],
	"260": intNumberTerms[260],
	"261": intNumberTerms[261],
	"262": intNumberTerms[262],
	"263": intNumberTerms[263],
	"264": intNumberTerms[264],
	"265": intNumberTerms[265],
	"266": intNumberTerms[266],
	"267": intNumberTerms[267],
	"268": intNumberTerms[268],
	"269": intNumberTerms[269],
	"270": intNumberTerms[270],
	"271": intNumberTerms[271],
	"272": intNumberTerms[272],
	"273": intNumberTerms[273],
	"274": intNumberTerms[274],
	"275": intNumberTerms[275],
	"276": intNumberTerms[276],
	"277": intNumberTerms[277],
	"278": intNumberTerms[278],
	"279": intNumberTerms[279],
	"280": intNumberTerms[280],
	"281": intNumberTerms[281],
	"282": intNumberTerms[282],
	"283": intNumberTerms[283],
	"284": intNumberTerms[284],
	"285": intNumberTerms[285],
	"286": intNumberTerms[286],
	"287": intNumberTerms[287],
	"288": intNumberTerms[288],
	"289": intNumberTerms[289],
	"290": intNumberTerms[290],
	"291": intNumberTerms[291],
	"292": intNumberTerms[292],
	"293": intNumberTerms[293],
	"294": intNumberTerms[294],
	"295": intNumberTerms[295],
	"296": intNumberTerms[296],
	"297": intNumberTerms[297],
	"298": intNumberTerms[298],
	"299": intNumberTerms[299],
	"300": intNumberTerms[300],
	"301": intNumberTerms[301],
	"302": intNumberTerms[302],
	"303": intNumberTerms[303],
	"304": intNumberTerms[304],
	"305": intNumberTerms[305],
	"306": intNumberTerms[306],
	"307": intNumberTerms[307],
	"308": intNumberTerms[308],
	"309": intNumberTerms[309],
	"310": intNumberTerms[310],
	"311": intNumberTerms[311],
	"312": intNumberTerms[312],
	"313": intNumberTerms[313],
	"314": intNumberTerms[314],
	"315": intNumberTerms[315],
	"316": intNumberTerms[316],
	"317": intNumberTerms[317],
	"318": intNumberTerms[318],
	"319": intNumberTerms[319],
	"320": intNumberTerms[320],
	"321": intNumberTerms[321],
	"322": intNumberTerms[322],
	"323": intNumberTerms[323],
	"324": intNumberTerms[324],
	"325": intNumberTerms[325],
	"326": intNumberTerms[326],
	"327": intNumberTerms[327],
	"328": intNumberTerms[328],
	"329": intNumberTerms[329],
	"330": intNumberTerms[330],
	"331": intNumberTerms[331],
	"332": intNumberTerms[332],
	"333": intNumberTerms[333],
	"334": intNumberTerms[334],
	"335": intNumberTerms[335],
	"336": intNumberTerms[336],
	"337": intNumberTerms[337],
	"338": intNumberTerms[338],
	"339": intNumberTerms[339],
	"340": intNumberTerms[340],
	"341": intNumberTerms[341],
	"342": intNumberTerms[342],
	"343": intNumberTerms[343],
	"344": intNumberTerms[344],
	"345": intNumberTerms[345],
	"346": intNumberTerms[346],
	"347": intNumberTerms[347],
	"348": intNumberTerms[348],
	"349": intNumberTerms[349],
	"350": intNumberTerms[350],
	"351": intNumberTerms[351],
	"352": intNumberTerms[352],
	"353": intNumberTerms[353],
	"354": intNumberTerms[354],
	"355": intNumberTerms[355],
	"356": intNumberTerms[356],
	"357": intNumberTerms[357],
	"358": intNumberTerms[358],
	"359": intNumberTerms[359],
	"360": intNumberTerms[360],
	"361": intNumberTerms[361],
	"362": intNumberTerms[362],
	"363": intNumberTerms[363],
	"364": intNumberTerms[364],
	"365": intNumberTerms[365],
	"366": intNumberTerms[366],
	"367": intNumberTerms[367],
	"368": intNumberTerms[368],
	"369": intNumberTerms[369],
	"370": intNumberTerms[370],
	"371": intNumberTerms[371],
	"372": intNumberTerms[372],
	"373": intNumberTerms[373],
	"374": intNumberTerms[374],
	"375": intNumberTerms[375],
	"376": intNumberTerms[376],
	"377": intNumberTerms[377],
	"378": intNumberTerms[378],
	"379": intNumberTerms[379],
	"380": intNumberTerms[380],
	"381": intNumberTerms[381],
	"382": intNumberTerms[382],
	"383": intNumberTerms[383],
	"384": intNumberTerms[384],
	"385": intNumberTerms[385],
	"386": intNumberTerms[386],
	"387": intNumberTerms[387],
	"388": intNumberTerms[388],
	"389": intNumberTerms[389],
	"390": intNumberTerms[390],
	"391": intNumberTerms[391],
	"392": intNumberTerms[392],
	"393": intNumberTerms[393],
	"394": intNumberTerms[394],
	"395": intNumberTerms[395],
	"396": intNumberTerms[396],
	"397": intNumberTerms[397],
	"398": intNumberTerms[398],
	"399": intNumberTerms[399],
	"400": intNumberTerms[400],
	"401": intNumberTerms[401],
	"402": intNumberTerms[402],
	"403": intNumberTerms[403],
	"404": intNumberTerms[404],
	"405": intNumberTerms[405],
	"406": intNumberTerms[406],
	"407": intNumberTerms[407],
	"408": intNumberTerms[408],
	"409": intNumberTerms[409],
	"410": intNumberTerms[410],
	"411": intNumberTerms[411],
	"412": intNumberTerms[412],
	"413": intNumberTerms[413],
	"414": intNumberTerms[414],
	"415": intNumberTerms[415],
	"416": intNumberTerms[416],
	"417": intNumberTerms[417],
	"418": intNumberTerms[418],
	"419": intNumberTerms[419],
	"420": intNumberTerms[420],
	"421": intNumberTerms[421],
	"422": intNumberTerms[422],
	"423": intNumberTerms[423],
	"424": intNumberTerms[424],
	"425": intNumberTerms[425],
	"426": intNumberTerms[426],
	"427": intNumberTerms[427],
	"428": intNumberTerms[428],
	"429": intNumberTerms[429],
	"430": intNumberTerms[430],
	"431": intNumberTerms[431],
	"432": intNumberTerms[432],
	"433": intNumberTerms[433],
	"434": intNumberTerms[434],
	"435": intNumberTerms[435],
	"436": intNumberTerms[436],
	"437": intNumberTerms[437],
	"438": intNumberTerms[438],
	"439": intNumberTerms[439],
	"440": intNumberTerms[440],
	"441": intNumberTerms[441],
	"442": intNumberTerms[442],
	"443": intNumberTerms[443],
	"444": intNumberTerms[444],
	"445": intNumberTerms[445],
	"446": intNumberTerms[446],
	"447": intNumberTerms[447],
	"448": intNumberTerms[448],
	"449": intNumberTerms[449],
	"450": intNumberTerms[450],
	"451": intNumberTerms[451],
	"452": intNumberTerms[452],
	"453": intNumberTerms[453],
	"454": intNumberTerms[454],
	"455": intNumberTerms[455],
	"456": intNumberTerms[456],
	"457": intNumberTerms[457],
	"458": intNumberTerms[458],
	"459": intNumberTerms[459],
	"460": intNumberTerms[460],
	"461": intNumberTerms[461],
	"462": intNumberTerms[462],
	"463": intNumberTerms[463],
	"464": intNumberTerms[464],
	"465": intNumberTerms[465],
	"466": intNumberTerms[466],
	"467": intNumberTerms[467],
	"468": intNumberTerms[468],
	"469": intNumberTerms[469],
	"470": intNumberTerms[470],
	"471": intNumberTerms[471],
	"472": intNumberTerms[472],
	"473": intNumberTerms[473],
	"474": intNumberTerms[474],
	"475": intNumberTerms[475],
	"476": intNumberTerms[476],
	"477": intNumberTerms[477],
	"478": intNumberTerms[478],
	"479": intNumberTerms[479],
	"480": intNumberTerms[480],
	"481": intNumberTerms[481],
	"482": intNumberTerms[482],
	"483": intNumberTerms[483],
	"484": intNumberTerms[484],
	"485": intNumberTerms[485],
	"486": intNumberTerms[486],
	"487": intNumberTerms[487],
	"488": intNumberTerms[488],
	"489": intNumberTerms[489],
	"490": intNumberTerms[490],
	"491": intNumberTerms[491],
	"492": intNumberTerms[492],
	"493": intNumberTerms[493],
	"494": intNumberTerms[494],
	"495": intNumberTerms[495],
	"496": intNumberTerms[496],
	"497": intNumberTerms[497],
	"498": intNumberTerms[498],
	"499": intNumberTerms[499],
	"500": intNumberTerms[500],
	"501": intNumberTerms[501],
	"502": intNumberTerms[502],
	"503": intNumberTerms[503],
	"504": intNumberTerms[504],
	"505": intNumberTerms[505],
	"506": intNumberTerms[506],
	"507": intNumberTerms[507],
	"508": intNumberTerms[508],
	"509": intNumberTerms[509],
	"510": intNumberTerms[510],
	"511": intNumberTerms[511],
	"512": intNumberTerms[512],
}

var intNumberTerms = [...]*Term{
	{Value: Number("0")},
	{Value: Number("1")},
	{Value: Number("2")},
	{Value: Number("3")},
	{Value: Number("4")},
	{Value: Number("5")},
	{Value: Number("6")},
	{Value: Number("7")},
	{Value: Number("8")},
	{Value: Number("9")},
	{Value: Number("10")},
	{Value: Number("11")},
	{Value: Number("12")},
	{Value: Number("13")},
	{Value: Number("14")},
	{Value: Number("15")},
	{Value: Number("16")},
	{Value: Number("17")},
	{Value: Number("18")},
	{Value: Number("19")},
	{Value: Number("20")},
	{Value: Number("21")},
	{Value: Number("22")},
	{Value: Number("23")},
	{Value: Number("24")},
	{Value: Number("25")},
	{Value: Number("26")},
	{Value: Number("27")},
	{Value: Number("28")},
	{Value: Number("29")},
	{Value: Number("30")},
	{Value: Number("31")},
	{Value: Number("32")},
	{Value: Number("33")},
	{Value: Number("34")},
	{Value: Number("35")},
	{Value: Number("36")},
	{Value: Number("37")},
	{Value: Number("38")},
	{Value: Number("39")},
	{Value: Number("40")},
	{Value: Number("41")},
	{Value: Number("42")},
	{Value: Number("43")},
	{Value: Number("44")},
	{Value: Number("45")},
	{Value: Number("46")},
	{Value: Number("47")},
	{Value: Number("48")},
	{Value: Number("49")},
	{Value: Number("50")},
	{Value: Number("51")},
	{Value: Number("52")},
	{Value: Number("53")},
	{Value: Number("54")},
	{Value: Number("55")},
	{Value: Number("56")},
	{Value: Number("57")},
	{Value: Number("58")},
	{Value: Number("59")},
	{Value: Number("60")},
	{Value: Number("61")},
	{Value: Number("62")},
	{Value: Number("63")},
	{Value: Number("64")},
	{Value: Number("65")},
	{Value: Number("66")},
	{Value: Number("67")},
	{Value: Number("68")},
	{Value: Number("69")},
	{Value: Number("70")},
	{Value: Number("71")},
	{Value: Number("72")},
	{Value: Number("73")},
	{Value: Number("74")},
	{Value: Number("75")},
	{Value: Number("76")},
	{Value: Number("77")},
	{Value: Number("78")},
	{Value: Number("79")},
	{Value: Number("80")},
	{Value: Number("81")},
	{Value: Number("82")},
	{Value: Number("83")},
	{Value: Number("84")},
	{Value: Number("85")},
	{Value: Number("86")},
	{Value: Number("87")},
	{Value: Number("88")},
	{Value: Number("89")},
	{Value: Number("90")},
	{Value: Number("91")},
	{Value: Number("92")},
	{Value: Number("93")},
	{Value: Number("94")},
	{Value: Number("95")},
	{Value: Number("96")},
	{Value: Number("97")},
	{Value: Number("98")},
	{Value: Number("99")},
	{Value: Number("100")},
	{Value: Number("101")},
	{Value: Number("102")},
	{Value: Number("103")},
	{Value: Number("104")},
	{Value: Number("105")},
	{Value: Number("106")},
	{Value: Number("107")},
	{Value: Number("108")},
	{Value: Number("109")},
	{Value: Number("110")},
	{Value: Number("111")},
	{Value: Number("112")},
	{Value: Number("113")},
	{Value: Number("114")},
	{Value: Number("115")},
	{Value: Number("116")},
	{Value: Number("117")},
	{Value: Number("118")},
	{Value: Number("119")},
	{Value: Number("120")},
	{Value: Number("121")},
	{Value: Number("122")},
	{Value: Number("123")},
	{Value: Number("124")},
	{Value: Number("125")},
	{Value: Number("126")},
	{Value: Number("127")},
	{Value: Number("128")},
	{Value: Number("129")},
	{Value: Number("130")},
	{Value: Number("131")},
	{Value: Number("132")},
	{Value: Number("133")},
	{Value: Number("134")},
	{Value: Number("135")},
	{Value: Number("136")},
	{Value: Number("137")},
	{Value: Number("138")},
	{Value: Number("139")},
	{Value: Number("140")},
	{Value: Number("141")},
	{Value: Number("142")},
	{Value: Number("143")},
	{Value: Number("144")},
	{Value: Number("145")},
	{Value: Number("146")},
	{Value: Number("147")},
	{Value: Number("148")},
	{Value: Number("149")},
	{Value: Number("150")},
	{Value: Number("151")},
	{Value: Number("152")},
	{Value: Number("153")},
	{Value: Number("154")},
	{Value: Number("155")},
	{Value: Number("156")},
	{Value: Number("157")},
	{Value: Number("158")},
	{Value: Number("159")},
	{Value: Number("160")},
	{Value: Number("161")},
	{Value: Number("162")},
	{Value: Number("163")},
	{Value: Number("164")},
	{Value: Number("165")},
	{Value: Number("166")},
	{Value: Number("167")},
	{Value: Number("168")},
	{Value: Number("169")},
	{Value: Number("170")},
	{Value: Number("171")},
	{Value: Number("172")},
	{Value: Number("173")},
	{Value: Number("174")},
	{Value: Number("175")},
	{Value: Number("176")},
	{Value: Number("177")},
	{Value: Number("178")},
	{Value: Number("179")},
	{Value: Number("180")},
	{Value: Number("181")},
	{Value: Number("182")},
	{Value: Number("183")},
	{Value: Number("184")},
	{Value: Number("185")},
	{Value: Number("186")},
	{Value: Number("187")},
	{Value: Number("188")},
	{Value: Number("189")},
	{Value: Number("190")},
	{Value: Number("191")},
	{Value: Number("192")},
	{Value: Number("193")},
	{Value: Number("194")},
	{Value: Number("195")},
	{Value: Number("196")},
	{Value: Number("197")},
	{Value: Number("198")},
	{Value: Number("199")},
	{Value: Number("200")},
	{Value: Number("201")},
	{Value: Number("202")},
	{Value: Number("203")},
	{Value: Number("204")},
	{Value: Number("205")},
	{Value: Number("206")},
	{Value: Number("207")},
	{Value: Number("208")},
	{Value: Number("209")},
	{Value: Number("210")},
	{Value: Number("211")},
	{Value: Number("212")},
	{Value: Number("213")},
	{Value: Number("214")},
	{Value: Number("215")},
	{Value: Number("216")},
	{Value: Number("217")},
	{Value: Number("218")},
	{Value: Number("219")},
	{Value: Number("220")},
	{Value: Number("221")},
	{Value: Number("222")},
	{Value: Number("223")},
	{Value: Number("224")},
	{Value: Number("225")},
	{Value: Number("226")},
	{Value: Number("227")},
	{Value: Number("228")},
	{Value: Number("229")},
	{Value: Number("230")},
	{Value: Number("231")},
	{Value: Number("232")},
	{Value: Number("233")},
	{Value: Number("234")},
	{Value: Number("235")},
	{Value: Number("236")},
	{Value: Number("237")},
	{Value: Number("238")},
	{Value: Number("239")},
	{Value: Number("240")},
	{Value: Number("241")},
	{Value: Number("242")},
	{Value: Number("243")},
	{Value: Number("244")},
	{Value: Number("245")},
	{Value: Number("246")},
	{Value: Number("247")},
	{Value: Number("248")},
	{Value: Number("249")},
	{Value: Number("250")},
	{Value: Number("251")},
	{Value: Number("252")},
	{Value: Number("253")},
	{Value: Number("254")},
	{Value: Number("255")},
	{Value: Number("256")},
	{Value: Number("257")},
	{Value: Number("258")},
	{Value: Number("259")},
	{Value: Number("260")},
	{Value: Number("261")},
	{Value: Number("262")},
	{Value: Number("263")},
	{Value: Number("264")},
	{Value: Number("265")},
	{Value: Number("266")},
	{Value: Number("267")},
	{Value: Number("268")},
	{Value: Number("269")},
	{Value: Number("270")},
	{Value: Number("271")},
	{Value: Number("272")},
	{Value: Number("273")},
	{Value: Number("274")},
	{Value: Number("275")},
	{Value: Number("276")},
	{Value: Number("277")},
	{Value: Number("278")},
	{Value: Number("279")},
	{Value: Number("280")},
	{Value: Number("281")},
	{Value: Number("282")},
	{Value: Number("283")},
	{Value: Number("284")},
	{Value: Number("285")},
	{Value: Number("286")},
	{Value: Number("287")},
	{Value: Number("288")},
	{Value: Number("289")},
	{Value: Number("290")},
	{Value: Number("291")},
	{Value: Number("292")},
	{Value: Number("293")},
	{Value: Number("294")},
	{Value: Number("295")},
	{Value: Number("296")},
	{Value: Number("297")},
	{Value: Number("298")},
	{Value: Number("299")},
	{Value: Number("300")},
	{Value: Number("301")},
	{Value: Number("302")},
	{Value: Number("303")},
	{Value: Number("304")},
	{Value: Number("305")},
	{Value: Number("306")},
	{Value: Number("307")},
	{Value: Number("308")},
	{Value: Number("309")},
	{Value: Number("310")},
	{Value: Number("311")},
	{Value: Number("312")},
	{Value: Number("313")},
	{Value: Number("314")},
	{Value: Number("315")},
	{Value: Number("316")},
	{Value: Number("317")},
	{Value: Number("318")},
	{Value: Number("319")},
	{Value: Number("320")},
	{Value: Number("321")},
	{Value: Number("322")},
	{Value: Number("323")},
	{Value: Number("324")},
	{Value: Number("325")},
	{Value: Number("326")},
	{Value: Number("327")},
	{Value: Number("328")},
	{Value: Number("329")},
	{Value: Number("330")},
	{Value: Number("331")},
	{Value: Number("332")},
	{Value: Number("333")},
	{Value: Number("334")},
	{Value: Number("335")},
	{Value: Number("336")},
	{Value: Number("337")},
	{Value: Number("338")},
	{Value: Number("339")},
	{Value: Number("340")},
	{Value: Number("341")},
	{Value: Number("342")},
	{Value: Number("343")},
	{Value: Number("344")},
	{Value: Number("345")},
	{Value: Number("346")},
	{Value: Number("347")},
	{Value: Number("348")},
	{Value: Number("349")},
	{Value: Number("350")},
	{Value: Number("351")},
	{Value: Number("352")},
	{Value: Number("353")},
	{Value: Number("354")},
	{Value: Number("355")},
	{Value: Number("356")},
	{Value: Number("357")},
	{Value: Number("358")},
	{Value: Number("359")},
	{Value: Number("360")},
	{Value: Number("361")},
	{Value: Number("362")},
	{Value: Number("363")},
	{Value: Number("364")},
	{Value: Number("365")},
	{Value: Number("366")},
	{Value: Number("367")},
	{Value: Number("368")},
	{Value: Number("369")},
	{Value: Number("370")},
	{Value: Number("371")},
	{Value: Number("372")},
	{Value: Number("373")},
	{Value: Number("374")},
	{Value: Number("375")},
	{Value: Number("376")},
	{Value: Number("377")},
	{Value: Number("378")},
	{Value: Number("379")},
	{Value: Number("380")},
	{Value: Number("381")},
	{Value: Number("382")},
	{Value: Number("383")},
	{Value: Number("384")},
	{Value: Number("385")},
	{Value: Number("386")},
	{Value: Number("387")},
	{Value: Number("388")},
	{Value: Number("389")},
	{Value: Number("390")},
	{Value: Number("391")},
	{Value: Number("392")},
	{Value: Number("393")},
	{Value: Number("394")},
	{Value: Number("395")},
	{Value: Number("396")},
	{Value: Number("397")},
	{Value: Number("398")},
	{Value: Number("399")},
	{Value: Number("400")},
	{Value: Number("401")},
	{Value: Number("402")},
	{Value: Number("403")},
	{Value: Number("404")},
	{Value: Number("405")},
	{Value: Number("406")},
	{Value: Number("407")},
	{Value: Number("408")},
	{Value: Number("409")},
	{Value: Number("410")},
	{Value: Number("411")},
	{Value: Number("412")},
	{Value: Number("413")},
	{Value: Number("414")},
	{Value: Number("415")},
	{Value: Number("416")},
	{Value: Number("417")},
	{Value: Number("418")},
	{Value: Number("419")},
	{Value: Number("420")},
	{Value: Number("421")},
	{Value: Number("422")},
	{Value: Number("423")},
	{Value: Number("424")},
	{Value: Number("425")},
	{Value: Number("426")},
	{Value: Number("427")},
	{Value: Number("428")},
	{Value: Number("429")},
	{Value: Number("430")},
	{Value: Number("431")},
	{Value: Number("432")},
	{Value: Number("433")},
	{Value: Number("434")},
	{Value: Number("435")},
	{Value: Number("436")},
	{Value: Number("437")},
	{Value: Number("438")},
	{Value: Number("439")},
	{Value: Number("440")},
	{Value: Number("441")},
	{Value: Number("442")},
	{Value: Number("443")},
	{Value: Number("444")},
	{Value: Number("445")},
	{Value: Number("446")},
	{Value: Number("447")},
	{Value: Number("448")},
	{Value: Number("449")},
	{Value: Number("450")},
	{Value: Number("451")},
	{Value: Number("452")},
	{Value: Number("453")},
	{Value: Number("454")},
	{Value: Number("455")},
	{Value: Number("456")},
	{Value: Number("457")},
	{Value: Number("458")},
	{Value: Number("459")},
	{Value: Number("460")},
	{Value: Number("461")},
	{Value: Number("462")},
	{Value: Number("463")},
	{Value: Number("464")},
	{Value: Number("465")},
	{Value: Number("466")},
	{Value: Number("467")},
	{Value: Number("468")},
	{Value: Number("469")},
	{Value: Number("470")},
	{Value: Number("471")},
	{Value: Number("472")},
	{Value: Number("473")},
	{Value: Number("474")},
	{Value: Number("475")},
	{Value: Number("476")},
	{Value: Number("477")},
	{Value: Number("478")},
	{Value: Number("479")},
	{Value: Number("480")},
	{Value: Number("481")},
	{Value: Number("482")},
	{Value: Number("483")},
	{Value: Number("484")},
	{Value: Number("485")},
	{Value: Number("486")},
	{Value: Number("487")},
	{Value: Number("488")},
	{Value: Number("489")},
	{Value: Number("490")},
	{Value: Number("491")},
	{Value: Number("492")},
	{Value: Number("493")},
	{Value: Number("494")},
	{Value: Number("495")},
	{Value: Number("496")},
	{Value: Number("497")},
	{Value: Number("498")},
	{Value: Number("499")},
	{Value: Number("500")},
	{Value: Number("501")},
	{Value: Number("502")},
	{Value: Number("503")},
	{Value: Number("504")},
	{Value: Number("505")},
	{Value: Number("506")},
	{Value: Number("507")},
	{Value: Number("508")},
	{Value: Number("509")},
	{Value: Number("510")},
	{Value: Number("511")},
	{Value: Number("512")},
}
