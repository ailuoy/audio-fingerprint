package gotest

import (
	"github.com/audio-fingerprint/fingerprint"
	"testing"
)

func BenchmarkSongFpCompare(b *testing.B) {
	firstSongFP := []int32{-1008033368, -1050021463, -1071037781, 1076507306, 1165632154, 1190797962, 1123758702, 1115161134, 1109918510, 1109930006, 1109925890, 1126850561, 1076555776, 1076543552, 1089110400, 1084981892, 1084981908, 1084919436, 1084918476, -1062233540, -1007896019, -1033049811, -1037277141, -1037293542, -1037091766, -1054211878, -998637414, -990240370, -994434162, -1046859010, -1009253786, -1026031066, -1023905018, -1027874553, 1118544133, 1135124789, 1084797221, 1090113575, 1077593127, 1077595238, 1077528950, 1077864902, 4107206, 20888518, 116297607, 115642263, 115643799, 1189384597, 1315017205, 1316073959, 1313923943, 1313923879, 1607529262, 1574150942, 1574413630, 1574281510, -581656282, -581656273, -48385731, -61950196, -61956596, -61957620, -62027204, -200307156, -183572948, -148998548, -165870851, -163839233, -231997202, -227802962, -202637138, 1908347070, 1354638542, 1352019150, 1419128014, 1419197406, 1150581502, 1200946894, 1184232135, 1183110085, 1317851637, 1322053988, 1296851236, 1280074021, 1276014885, 1276011829, 1280305469, 1288546685, 1556786125, -539321395, -555840628, -553808484, -554925668, -554401380, -564887300, -564891411, -570154771, -548986386, -582602258, -589940762, -589348890, -589529114, -597894314, -602100921, -1675711739, -1675710716, -1671520764, -1671459324, -1646492123, -1651182041, -1718132185, -1680399834, -1680600538, -1680604370, -1681653650, -1681652498, -1681642242, -1706743602, -1706678066, -1974584114, -1966185266, -1949425186, -1886703137, -1920265777, -1932852771, -1940323843, -866647579, -1005059596, -1004984844, -1006033451, -1005883963, -1005847099, -1005663289, -1005679737, 1158581127, 1179552133, 1318557061, 1320530325, 1320464821, 1318441381, 1322578405, -838835737, -838852249, -553163409, -582591201, -582550257, -574162609, -582697649, -582697651, -582179507, -582048499, -1672571124, -1672573428, -1663201748, -1126330835, -1126322643, -1107433939, -1233358289, -1233370513, -1233387282, -1233387346, -1235484498, -145940298, -183623546, -736222074, 1411259782, 1423837830, 1419713230, 1419517646, 1436303055, 1172056015, 1205608919, 98382325, 215823205, 206377845, 1280127767, 1280141079, 1280076583, 1281157927, 1289726774, 1342094150, 1321111494, 1319993798, 1319928007, -828014345, -822755097, -833294107, -833330715, -851155995, -867764619, -859382203, -842604859, -549146939, -553337003, -553347724, -1626495716, -1654741748, -1134643444, -1134590452, -1134721495, -1134719445, -1134563797, -1199870422, -1183302102, -1150810582, -1150548630, -1694757654, -1694757653, -1694690081, -1679944625, -1679941625, -1717690361, -1650499530, -1920114634, -1920176073, -1915981705, -1941170699, -1005910571, -1005975675, -1005975660, -1005901900, -1006033004, 1141452164, 1141705093, 1141852549, 1162791813, 1187961735, 1319092102, -826494506, -827543305, -826502937, -832778137, -836968345, -817045147, -858856587, -863063483, -862996905, -857954321, -870539281, -853766673, -1642164915, -1642156787, -1625379060, -1621189108, -1646424548, -1646424536, -1113741784, -1075977687, -1101499861, -1237843414, -163839510, -700710742, -226754386, -212070210, 1348341902, 1348351118, 1348344974, 274603663, 278886047, 278718031, 280798733, 284999436, 1895604028, 1887218476, 1912449572, -142693868, -163674620, -1237614076, -1237614059, -1238732185, -1238756634, -1234627930, -1213713754, -1264049754, -1264049994, -1800916842, -1799813994, -793047913, -788916075, 1358887109, 1345251396, 1395566660, 1445898308, 372223173, 372216023, 372228278, 372179126, 1445855414, 1601045094, 1555968550, 1420635686, 1420569127, 1152133655, 1152134679, 1207788037, 1191138245, -958447227, -959889259, -968269659, -955554651, -1056285545, -1071985209, -1065693369, -1062564091, -1045848548, -2119385284, -2119389652, -2018726356, -1767066068, -1762879940, -1628597732, -1636990404, -1641250001, -1640135898, -1707355866, -1707470554, -1690685098, -1189469738, -1738858250, -661118746, -662105898, -662043442, -661702450, -863229705, -867407641, -871577129, -820606267, -564754811, -690576763, -692800635, -688095051, -700612443, -214006619, -247550540, -264328556, -264245548, -264382748, -465775060, -411277780, -1493408212, -1229183443, -1224991449, -1224995786, -1300481002, -1854004202, -1870785530, -1799748602, -1799748586, -2070215566, -2070145694, -2071190173, -2071173645, -2075627279, -2063023919, -2062925375, -2062925439, -2050354299, -2054548553, -2071259225, -2066980889, -1001595418, -1004732970, -1004556978, -853627570, -821056434, -833639346, -825252785, 1313846365, -836795291, -834699036, -835813123, -831598115, -555036195, -563363355, -563494427, -583470107, -595729547, -1669340395, -1660967371, -1661092315, -1663263132, -1671586059, -1671446827, -1126175097, -1109528938, -1621336410, -1638117650, -1642328082, -1638128274, -1638391506, -1630003138, -1613164513, -1766129633, -1766125507, -1766235075, -1632082883, -1906805729, -1910946785, -1911667505, -1894878065, -1924520561, -859168867, -859172947, -926343251, -926345795, -933756483, -669236867, -669134049, -937434609, -937562610, -933434850, -933434514, -917739026, -888325890, -900852594, -892465009, -1903820337, -1903819793, -1903857105, -1886592465, -1919109593, -2053335259, -2047047899, -2072475803, -1808234795, -1808163195, -1805996412, -1790267756, -1794394460, -1224112475, -1223063835, -1254369545, -1255512777, -1221891033, -1217844186, -1234629594, -1233449977, -1774570489, -1770380281, -1745215481, -1782972110, -1801584078, -1802697998, -2071194893, -2071173421, -2066980141, -2058591753, -2062736923, -2079444507, -2079456811, -2075524665, -2067069498, -997695530, -997663258, -993465882, -1001867018, -954229690, -968967097, -964774843, -969964428, -970096540, -970018588, -835800620, -831606332, -830553659, -830750219, -837168665, -816209561, -1663442057, -1663310217, -1663437209, -1661411739, -1671897499, -1671823756, -1663308076, -1797656875, -1789067547, -1755558169, -1759754649, -1773385922, -1773394122, -1773395658, -1756618698, -1748033474, -1748106194, -1769012178, -1769206754, -2033447857, -2033644345, -2016799609, -2050345849, -1916132969, -1933041235, -1941425747, -867493987, -867559539, -934598771, -937613443, 1209890621, 1209921295, 1211936526, 1209839370, 1209837946, -888300182, -904945702, -900816946, -892430513, -829713137, -1903627985, -1898389201, -1885740753, -1885695185, -2053471961, -2053471963, -1797750491, -1805877003, -1811118635, -1878024300, -1878085996, -1878085964, -1878067532, -1323384179, -1287699827, -1305737315, -1842596369, -1838404498, -1828901842, -1837225970, -1820383218, -1786824690, -1778512866, -1782708930, -1749155058, -2017656305, -2036846051, -2036834004, -2036834268, -2015621020, -2074345372, -2078539532, -2079447852, -1005705772, -1001509163, -993116475, -992082235, -990034233, -1058224569, -1066629561, -1054046377, -1053976747, -1053837611, -1020283179, -1020283179, -1020415275, -768757003, -1842497803, -1771194649, -1771192729, -1838285209, -1838317850, -1838596106, -1838606410, -1838607946, -2107043658, -2098515786, -2102714202, -2101796702, -2101862174, -2101858078, -2101862398, -2119688185, -2132254697, -2132225739, -2140551916, -1066410732, -1070694636, 1076780804, 1076776196, 1076776197, 4117765, 88003911, 121558342, 121624022, 124646870, 124622294, 258838998, 86873559, 1160754615, 1198505911, 1336917941, 1320140724, 1320140724, -828000859, -828071771, -828071771, -822816603, -837758809, -838835993, -838835738, -867802122, 1288106702, 1288098526, -867910018, -872105169, -603145427, -603084531, -551638260, -564220148, -555842036, -555908596, -1629650387, -1629494739, -1621401041, -1655102929, -1723328978, -1723331282, -1681388418, -1681376002, -1681183522, -1698028338, -1970655162, -1953861546, -1953931146, -1953734401, -1920175875, -1915989796, -1940110124, -870513164, -870447132, -871491868, -871487916, -871422140, -1005769404, -1005905467, -1005958953, -988919673, -955332218, -828448890, -828385658, -826477898, -826477914, -828583001, -824385041, -833822227, -817044499, -867240996, -866200612, -866129923, -870519315, -871581187, -871524019, -603018995, -1675711731, -1671512563, -1646293491, -1646424515, -1617066451, -1634885073, -1631198673, -1638632914, -1642855570, -1643904850, -1773928274, -700182354, -683466594, -733600626, 1409694863, 1409681807, 1414925199, 1423248335, 349465551, 401892815, 116679117, 116732925, 131404653, 81072685, 72159789, 1145909781, 1145908789, 1145858853, 1288498021, 1285483335, 1319017415, 1321111494, 1325237190, 1324122518, 1324712438, 1322631398, -824909586, -825990929, -842702483, -842632900, -841550020, -849477828, -820199556, -551765124, -585324212, -585193140, -580993204, -1654749620, -1672509924, -1671593427, -1671589329, -1669471698, -1669503442, -1669581266, -1652814034, -1652815826, -1680078730, -1681113898, -1681513402, -1681509306, -1681514410, -1618009738, -1886576266, -1886637737, -1916260009, -1916260107, -1933045532, -2074541596, -998702732, -998694572, -998562492, -1004851772, -1004897836, -1005926267, -1005926011, -988903033, -959542905, -962624122, -826502761, -826478153, -822750747, -831233051, -833330203, -824941595, -845683715, -862526497, -858282514, -866861722, -870544026, -602108642, -601903857, -585191667, -44122612, -1117876724, -1084256740, -1084256707, -1084295635, -1105238481, -1103359442, -1103363538, -25427858, -161742610, -170127186, -192470850, 1422271631, 1422797967, 1422802063, 1154295183, 80504719, 80607055, 340755727, 1415546125, 1415417140, 1423747876, -706893276, -673269244, -698380796, -1774616060, -1640246763, -1640361417, -1640365530, -1636101530, -1648877082, -1665654554, -1665650526, -1667755870, -1935866734, -857781115, -1000343419, -1004685116, 1142864068, 1176414336, 1312729217, 1312734338, 1580040326, 1580056726, 1579634870, 1579553270, 1568027238, 1555441190, 1555315239, 1287399989, 1285302837, 1302079077, -807443595, -807517243, -808631931, -976428555, -1001594779, -1004536729, -1004615673, -1052850169, -1052850171, -1006724860, -1010987748, -2017685700, -2017681876, -1766017492, -1766021588, -1766018516, -1761832404, -1640201683, -1640463825, -1641426394, -1624575450, -1658183114, -1673911546, -1136975850, -1661460378, -1732759322, -1734737706, -1734736690, -660985650, -858248986, -870835994, -868496938, -851801401, -553994612, -559181172, -693595508, -693595748, -696667996, -163928924, -182800476, -1273254235, -1273389403, -1273114905, -1606433177, -1602361817, -1308858826, -1313063626, -1308877770, -1317270506, -1870783486, -1871047662, -1862659022, -1866857438, -1866791646, -1866918541, -2135412397, -2135420479, -2141699903, -2129116975, -2128998272, -1056296832, -1056033404, -1068616252, -1060167692, -993255436, -993249323, -993757369, -1001920690, -985208994, -817383554, -825911938, -830102178, 1317383262, 1322102015, 1310559469, 1310557341, 1312457869, 1312465033, 1310367965, 1578807551, 1595506815, 1561953389, -602299027, -597973651, -598039172, -61175980, -1269053740, -1268791564, -1260390684, -1260394779, -1797429659, -1780736409, -1747186138, -1764487642, -1764481738, -1768152026, -1768157130, -1835204586, -1818361770, -2080500650, -2099356554, -2040701865, -2044953529, -2029093561, -2024905273, -2050072123, -1932577899, -859040851, -859302484, -859236948, -934343236, -934339139, -938648067, -870486577, 1277161423, 1277030287, 1275915983, 1296887551, -817037329, -825393169, -825329841, -829532401, -828390081, -1902324433, -2019760849, -2019692249, -2019590025, -2019589899, -2050129435, -1810071195, -1810058380, -1810125116, -1805735284, -1805735236, -1256224083, -1256403281, -1222883649, -1088867618, -1103289506, -1103354818, -1090773970, -1226306514, -1762980834, -1746256882, -1778768889, -2068175865, -2070253561, -2070318825, -2071404253, -2071400093, -2071412239, -2050375215}
	secondSongFP := []int32{-1008033368, -1050021463, -1071037781, 1076507306, 1164583578, 1190797962, 1123758702, 1115161134, 1109918502, 1109930006, 1109925890, 1126850561, 1076555776, 1076527168, 1089110400, 1084981908, 1084981908, 1084984972, 1084918476, -1062495684, -1007896019, -1033049299, -1037277141, -1037293542, -1037091766, -1054211878, -999685986, -990240370, -994434162, -979750146, -1009253786, -1026031066, -1023905018, -1027874553, 1119592709, 1135124789, 1084797221, 1090048039, 1077593127, 1077595495, 1077528950, 1077864902, 4107206, 20888518, 116296583, 116166551, 115643799, 1189384597, 1315017205, 1316073959, 1313923943, 1313923879, 1574040366, 1574150942, 1574414142, 1574281518, -581656282, -585850577, -65162947, -66144500, -61956596, -61961716, -62027204, -200307156, -183572948, -148998548, -166132995, -163839233, -231997202, -227802898, -202637138, 1908347070, 1354638542, 1352019150, 1419128014, 1419197406, 1150581502, 1200946894, 1184232135, 1183110085, 1322045941, 1322053988, 1296851236, 1280074021, 1276014885, 1276011829, 1280305453, 1288546685, 1556786125, -539321395, -555840628, -553808484, -555449956, -554401380, -564887300, -564891411, -570154771, -548986386, -582602258, -589940762, -589348890, -589529114, -597894314, -602100921, -1675711739, -1675710716, -1671520764, -1671590396, -1646492123, -1646987737, -1718132185, -1680399834, -1680600538, -1680604370, -1681653650, -1681652498, -1681642242, -1706743602, -1706678066, -1974584114, -1966185266, -1949425186, -1886703137, -1920265777, -1932852771, -1940323843, -866647580, -1004994060, -1004984844, -1006033452, -1005883963, -1005847099, -1005663289, -1005679737, 1158581127, 1179552135, 1318557063, 1318433173, 1320464821, 1318441381, 1322578405, 1308647911, -838852249, -553163409, -582591137, -582551281, -574162609, -582697649, -582697651, -582179507, -582048499, -1672571124, -1672575476, -1663070676, -1126330835, -1126322643, -1107433939, -1233358291, -1233370513, -1233387282, -1233387346, -1235484498, -144891722, -183623546, -736222074, 1411259782, 1423841926, 1419713230, 1419517646, 1436303055, 1172056015, 1205609431, 98382325, 215823205, 206377845, 1280127767, 1280124695, 1280076583, 1281157927, 1289726774, 1339996998, 1321111494, 1319993798, -827555641, -828079881, -822755097, -833294107, -833330715, -851155995, -867764619, -859382203, -842604859, -549146939, -553337003, -553347724, -552753892, -1654741748, -1134643444, -1135639028, -1134721495, -1134719445, -1134563797, -1199870422, -1183302102, -1149762006, -1149500054, -1694759702, -1694759701, -1694690081, -1679944625, -1679941625, -1717690361, -1650499530, -1920114634, -1920176073, -1915981705, -1941433867, -1005910571, -1005975675, -1005971564, -1005901900, -1006033004, 1141452164, 1141705093, 1141852549, 1162791813, 1187961735, 1319092102, -826494506, -827543305, -826502937, -832778137, -836968345, -833822427, -858856587, -863063483, -862996745, -857954321, -870539793, -853766673, -1642164915, -1642156787, -1625379060, -1621189108, -1646424548, -1650618840, -1113737688, -1075977687, -1235717589, -1237843414, -164101654, -700710742, -226754386, -212070210, 1348341902, 1348351118, 1348344974, 274603663, 278886047, 278718031, 280798733, 284999436, 1895604028, 1887218476, 1912449580, -142693868, -163672572, -1237614076, -1237614059, -1238732185, -1238756634, -1234627930, -1213713754, -1264049754, -1264049994, -1800916842, -1799813994, -793047913, -788916075, 1358887109, 1345251396, 1395566660, 1445898308, 372223173, 372216023, 372228278, 372179126, 1445855670, 1601045094, 1421750822, 1420635686, 1420569127, 1152133655, 1152134679, 1207788037, 1191137733, -958447227, -959889259, -968269659, -955554651, -1056285545, -1071985209, -1065693369, -1062564083, -1045848292, -2119385284, -2119389652, -2018726356, -1767066068, -1762880964, -1628597732, -1636990420, -1641250001, -1640135898, -1707355866, -1707462362, -1690685098, -1189469738, -1738858250, -661118746, -662105898, -662043442, -661702450, -863229705, -867407641, -871577129, -821654843, -564754811, -690576763, -692800635, -688095051, -700612443, -230783835, -247550540, -264328556, -264245612, -264382748, -465775060, -412326356, -1494456788, -1229183451, -1224995545, -1224995786, -1300481002, -1854004218, -1870785530, -1866595322, -1799748586, -2070215582, -2070145694, -2071190173, -2071173647, -2075627311, -2063023935, -2062925375, -2062925439, -2050354299, -2054548555, -2070997081, -2066980889, -1001595418, -1004732970, -1004556978, -853562034, -821056434, -833639346, -825252785, 1313846357, -836795291, -834699036, -835813124, -831589923, -555036195, -563363355, -563494427, -549916187, -595729547, -1669340395, -1660967115, -1661092315, -1663263131, -1671586059, -1671446827, -1126175033, -1109528938, -1621336410, -1638117650, -1638133778, -1638128274, -1630002898, -1613225922, -1613164513, -1766129633, -1766125507, -1766235075, -1632082883, -1906805731, -1910946785, -1911667505, -1894882161, -1924520561, -859168865, -859172947, -926343251, -926345283, -665321027, -669236867, -669125857, -937434609, -937562610, 1214048798, -933434514, -917739026, -888325890, -900852594, -892465009, -1903820337, -1903754257, -1903857105, -1886592465, -1919109593, -2053335259, -2047047899, -2072475803, -1808234795, -1808163195, -1805930876, -1790267756, -1794394460, -1224112475, -1223063835, -1220815113, -1255480009, -1221891033, -1217844186, -1234629594, -1233449977, -1774570489, -1770380281, -1745215481, -1782972110, -1801584078, -1802697998, -2071194893, -2071173421, -2066980141, -2058591753, -2062736923, -2079444507, -2079456811, -2075524665, -2067068986, -997695530, -997663258, -993465882, -1001867018, -954229690, -968967097, -964774843, -969964428, -970092444, -970018588, -835800620, -831606332, -830553659, -830750219, -837168665, -816193177, -1663442057, -1667504521, -1663437209, -1661411739, -1671897499, -1671823756, -1663308092, -1797656875, -1789067547, -1755525401, -1759721881, -1773385922, -1773394122, -1773395658, -1756618690, -1748033474, -1748106194, -1769012178, -1769206754, -2033447857, -2033640249, -2016799609, -2050345849, -1916132969, -1933041235, -1941425747, -867493987, -867559539, -934598771, -937613443, 1209890621, 1209921295, 1211936526, 1209839374, 1209837882, -888300182, -904945702, -900816946, -892430513, -829713137, -1903627985, -1898389201, -1885740753, -1885695185, -2053471961, -2053471963, -1797750491, -1805876747, -1811118635, -1878024300, -1878085996, -1878085964, -1878067532, -1323384179, -1287699827, -1305737315, -1842596369, -1838404498, -1828901842, -1837160434, -1820383218, -1786824690, -1778512866, -1782708930, -1749155058, -2017656305, -2036837859, -2036834004, -2036834268, -2015621020, -2074345372, -2078539532, -2079447852, -1005705772, -1001509163, -993116459, -992082235, -990034233, -1058224569, -1066629561, -1054046377, -1053976747, -1053837611, -1020283179, -1020283179, -1020415275, -768757003, -1842497803, -1771194649, -1771192729, -1838285209, -1838317850, -1838596106, -1838606410, -1838607946, -2107043658, -2098515786, -2102714206, -2101796702, -2101862174, -2101858078, -2101862398, -2119688185, -2132254697, -2132225739, -2140551916, -1066410732, -1070694636, 1076780804, 1076776196, 1076776197, 4117765, 88003911, 121558342, 121624022, 124646870, 124622294, 258838998, 86873559, 1160754615, 1198505911, 1341112245, 1320140724, 1320140724, -828000843, -828071771, -828071771, -822816603, -837758809, -838835993, -834641434, -867802122, 1288106702, 1288098526, -867910018, -872105169, -603145939, -603084531, -551638260, -564220148, -555842036, -555908580, -1629650387, -1629494739, -1621401041, -1655102929, -1723328978, -1723331282, -1681388418, -1681376002, -1681183522, -1698028338, -1970655162, -1953861546, -1953931146, -1953734401, -1920175875, -1915989796, -1940110124, -870513164, -870447132, -871491868, -871487916, -871422140, -1005769404, -1005905467, -1005696809, -988919673, -955332218, -828448890, -828385658, -826477898, -826477914, -828583001, -824385041, -833822227, -817044499, -867240964, -866200612, -866127875, -870535699, -871581187, -871524019, -603018995, -1675710707, -1671512563, -1646359027, -1646424515, -1617066451, -1634885073, -1631198673, -1638632914, -1643904146, -1643904850, -1773928274, -700182354, -683466594, -733600626, 1409694863, 1409681807, 1413876623, 1423264719, 349989839, 401892815, 116679117, 116732797, 131404653, 81596973, 72159789, 1145909781, 1145908789, 1145858853, 1288563557, 1285483335, 1319017415, 1321111494, 1325237190, 1324122518, 1324712438, 1322631398, -824909586, -825990929, -842702483, -842632900, -841550020, -849477828, -820199556, -551765124, -585324212, -585192628, -580993204, -1654749620, -1672509924, -1671593427, -1671589329, -1669471698, -1669503442, -1669579218, -1652814034, -1619261394, -1680078722, -1681113898, -1681513402, -1681509306, -1681514410, -1618009738, -1920130698, -1886637737, -1916260009, -1916259851, -1933045276, -2074541596, -998702732, -998694588, -998562492, -1004851772, -1004897836, -1005926267, -1005926267, -988903033, -959542905, -962624122, -826502761, -826478153, -822750747, -831233051, -833330203, -824941595, -845683715, -862526497, -858282514, -866861722, -870544026, -602108642, -601903857, -585191667, -44122612, -1117876724, -1084256740, -1084258755, -1084295635, -1101044177, -1103359442, -1103363538, -25427858, -161742610, -170127186, -192470850, 1422271631, 1422797967, 1422802063, 1154295183, 80504719, 80607183, 340755727, 1415546125, 1415417140, 1423747876, -706893276, -673269244, -698380796, -1774616060, -1640246763, -1640361417, -1640365530, -1636101530, -1648877082, -1665654554, -1665650526, -1667755870, -1935866750, -857781119, -1000343419, -1004685116, 1142864064, 1176414336, 1312729217, 1581165698, 1580040326, 1580056726, 1579634838, 1579553270, 1576415846, 1555441190, 1555315239, 1287399989, 1285302837, 1302079077, -807443595, -807517243, -808631931, -976428555, -1001594651, -1005585305, -1004615673, -1052850169, -1052850171, -1006724859, -1010987748, -2017685700, -2017681876, -1766021588, -1766021588, -1766018516, -1761832404, -1640201683, -1640463825, -1641426394, -1624575450, -1658183114, -1673911546, -1136975850, -1661460378, -1732759322, -1734737706, -1734736690, -660985650, -858248986, -870835994, -868496938, -851801403, -553994612, -559181172, -693595508, -689401444, -696667996, -163928924, -182800476, -1273254235, -1273389403, -1273114905, -1606433177, -1602361817, -1308858826, -1313063626, -1308877770, -1317266410, -1870783486, -1871047662, -1862659022, -1866857438, -1866791646, -1866918541, -2135420589, -2135420479, -2141699903, -2129116976, -2128998272, -1056296828, -1056033404, -1068616252, -1060167724, -993255436, -993249323, -993757369, -1001920690, -850991266, -817383554, -825903746, -830102178, 1317383262, 1322102015, 1310493933, 1310557341, 1312720009, 1312465033, 1310367961, 1578807551, 1595506815, 1561952365, -602299027, -597973651, -598039172, -61175980, -1269053740, -1269053708, -1260390684, -1260394779, -1797429659, -1780736409, -1747186138, -1764487642, -1764481738, -1768152026, -1768157130, -1835204586, -1818361770, -2080500618, -2099356554, -2040701865, -2044953529, -2029093561, -2024905273, -2050072123, -1932577899, -858975315, -859302484, -859236948, -934343236, -934339139, -938648067, -870486577, 1277161423, 1277030287, 1275915983, 1296887551, -817037329, -825393169, -825329841, -829532401, -828914369, -1902848721, -2019760849, -2019692505, -2019590025, -2053144331, -2049080859, -1810071067, -1810066572, -1810125116, -1805735284, -1805735236, -1256219987, -1256403281, -1222883649, -1088867618, -1103289506, -1103354818, -1090773970, -1763177426, -1762980834, -1763034098, -1778768889, -2068175865, -2070253561, -2070318825, -2071404253, -2071400093, -2071412239, -2050375215}
	for i := 0; i < b.N; i++ {
		_, _, err := fingerprint.Correlate(firstSongFP, secondSongFP)
		if err != nil {
			b.Error(err)
		}
	}
}