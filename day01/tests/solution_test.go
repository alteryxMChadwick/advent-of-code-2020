package day01_tests

import (
	"testing"

	code "github.com/capsocrates/advent-of-code/day01"
	"github.com/stretchr/testify/assert"
)

func TestDay01Pair(t *testing.T) {
	t.Run("TestEmpty", func(t *testing.T) {
		left, right := code.GetPair([]int{})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, right)
	})

	t.Run("TestSingleMatchingElement", func(t *testing.T) {
		left, right := code.GetPair([]int{2020})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, right)
	})

	t.Run("TestTwoMatchingElements", func(t *testing.T) {
		left, right := code.GetPair([]int{1000, 1020})

		assert.Equal(t, 1020, left)
		assert.Equal(t, 1000, right)
	})

	t.Run("TestTwoMatchingElementsAgain", func(t *testing.T) {
		left, right := code.GetPair([]int{0, 2020})

		assert.Equal(t, 2020, left)
		assert.Equal(t, 0, right)
	})

	t.Run("TestNoMatchingElement", func(t *testing.T) {
		left, right := code.GetPair([]int{0})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, right)
	})

	t.Run("ExampleSet", func(t *testing.T) {
		exampleNums := []int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		}

		left, right := code.GetPair(exampleNums)

		assert.Equal(t, 1721, left)
		assert.Equal(t, 299, right)

		product := code.GetProduct(left, right)
		assert.Equal(t, 514579, product)
	})

	t.Run("AnswerSet", func(t *testing.T) {
		nums := []int{1749, 1897, 881, 1736, 1161, 1720, 1676, 305, 264, 1904, 1880, 1173, 483, 1978, 1428, 1635, 1386, 1858, 1602, 1916, 1906, 1212, 1730, 1777, 1698, 1845, 1812, 1922, 1729, 1803, 1761, 1901, 1748, 1188, 1964, 1935, 1919, 1810, 1567, 1849, 1417, 1452, 54, 1722, 1784, 1261, 1744, 1594, 1526, 1771, 1762, 1894, 1717, 1716, 51, 1955, 1143, 1741, 1999, 1775, 1944, 1983, 1962, 1198, 1553, 1835, 1867, 1662, 1461, 1811, 1764, 1726, 1927, 1179, 1468, 1948, 1813, 1213, 1905, 1371, 1751, 1215, 1392, 1798, 1823, 1815, 1923, 1942, 1987, 1887, 1838, 1395, 2007, 1479, 1752, 1945, 1621, 1538, 1937, 565, 1969, 1493, 1291, 1438, 1578, 1770, 2005, 1703, 1712, 1943, 2003, 1499, 1903, 1760, 1950, 1990, 1185, 1809, 1337, 1358, 1743, 1707, 1671, 1788, 1785, 1972, 1863, 1690, 1512, 1963, 1825, 1460, 1828, 1902, 1874, 1755, 1951, 1830, 1767, 1787, 1373, 1709, 1514, 1807, 1791, 1724, 1859, 1590, 1976, 1572, 1947, 1913, 1995, 1728, 1624, 1731, 1706, 1782, 1994, 1851, 1843, 1773, 1982, 1685, 2001, 1346, 1200, 1746, 1520, 972, 1834, 1909, 2008, 1733, 1960, 1280, 1879, 1203, 1979, 1133, 1647, 1282, 1684, 860, 1444, 1780, 1989, 1795, 1819, 1797, 1842, 1796, 1457, 1839, 1853, 1711, 1883, 1146, 1734, 1389}

		left, right := code.GetPair(nums)

		assert.Equal(t, 1969, left)
		assert.Equal(t, 51, right)

		product := code.GetProduct(left, right)
		assert.Equal(t, 100419, product)
	})
}

func BenchmarkTestPair(b *testing.B) {
	nums := []int{1749, 1897, 881, 1736, 1161, 1720, 1676, 305, 264, 1904, 1880, 1173, 483, 1978, 1428, 1635, 1386, 1858, 1602, 1916, 1906, 1212, 1730, 1777, 1698, 1845, 1812, 1922, 1729, 1803, 1761, 1901, 1748, 1188, 1964, 1935, 1919, 1810, 1567, 1849, 1417, 1452, 54, 1722, 1784, 1261, 1744, 1594, 1526, 1771, 1762, 1894, 1717, 1716, 51, 1955, 1143, 1741, 1999, 1775, 1944, 1983, 1962, 1198, 1553, 1835, 1867, 1662, 1461, 1811, 1764, 1726, 1927, 1179, 1468, 1948, 1813, 1213, 1905, 1371, 1751, 1215, 1392, 1798, 1823, 1815, 1923, 1942, 1987, 1887, 1838, 1395, 2007, 1479, 1752, 1945, 1621, 1538, 1937, 565, 1969, 1493, 1291, 1438, 1578, 1770, 2005, 1703, 1712, 1943, 2003, 1499, 1903, 1760, 1950, 1990, 1185, 1809, 1337, 1358, 1743, 1707, 1671, 1788, 1785, 1972, 1863, 1690, 1512, 1963, 1825, 1460, 1828, 1902, 1874, 1755, 1951, 1830, 1767, 1787, 1373, 1709, 1514, 1807, 1791, 1724, 1859, 1590, 1976, 1572, 1947, 1913, 1995, 1728, 1624, 1731, 1706, 1782, 1994, 1851, 1843, 1773, 1982, 1685, 2001, 1346, 1200, 1746, 1520, 972, 1834, 1909, 2008, 1733, 1960, 1280, 1879, 1203, 1979, 1133, 1647, 1282, 1684, 860, 1444, 1780, 1989, 1795, 1819, 1797, 1842, 1796, 1457, 1839, 1853, 1711, 1883, 1146, 1734, 1389}
	for i := 0; i < b.N; i++ {
		code.GetPair(nums)
	}
}

func TestDay01Triplet(t *testing.T) {
	t.Run("TestEmpty", func(t *testing.T) {
		left, middle, right := code.GetTriplet([]int{})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, middle)
		assert.Equal(t, 0, right)
	})

	t.Run("TestSingleMatchingElement", func(t *testing.T) {
		left, middle, right := code.GetTriplet([]int{2020})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, middle)
		assert.Equal(t, 0, right)
	})

	t.Run("TestThreeMatchingElements", func(t *testing.T) {
		left, middle, right := code.GetTriplet([]int{500, 1010, 510})

		assert.Equal(t, 1010, left)
		assert.Equal(t, 510, middle)
		assert.Equal(t, 500, right)
	})

	t.Run("TestThreeMatchingElementsAgain", func(t *testing.T) {
		left, middle, right := code.GetTriplet([]int{15, 2000, 5})

		assert.Equal(t, 2000, left)
		assert.Equal(t, 15, middle)
		assert.Equal(t, 5, right)
	})

	t.Run("TestNoMatchingElement", func(t *testing.T) {
		left, middle, right := code.GetTriplet([]int{0, 5, 10, 15, 20})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, middle)
		assert.Equal(t, 0, right)
	})

	t.Run("TestNoMatchingElementAgain", func(t *testing.T) {
		left, middle, right := code.GetTriplet([]int{0, 5, 10, 15, 20})

		assert.Equal(t, 0, left)
		assert.Equal(t, 0, middle)
		assert.Equal(t, 0, right)
	})

	t.Run("ExampleSet", func(t *testing.T) {
		exampleNums := []int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		}

		left, middle, right := code.GetTriplet(exampleNums)

		assert.Equal(t, 979, left)
		assert.Equal(t, 675, middle)
		assert.Equal(t, 366, right)

		product := code.GetTripletProduct(left, middle, right)
		assert.Equal(t, 241861950, product)
	})

	t.Run("AnswerSet", func(t *testing.T) {
		nums := []int{1749, 1897, 881, 1736, 1161, 1720, 1676, 305, 264, 1904, 1880, 1173, 483, 1978, 1428, 1635, 1386, 1858, 1602, 1916, 1906, 1212, 1730, 1777, 1698, 1845, 1812, 1922, 1729, 1803, 1761, 1901, 1748, 1188, 1964, 1935, 1919, 1810, 1567, 1849, 1417, 1452, 54, 1722, 1784, 1261, 1744, 1594, 1526, 1771, 1762, 1894, 1717, 1716, 51, 1955, 1143, 1741, 1999, 1775, 1944, 1983, 1962, 1198, 1553, 1835, 1867, 1662, 1461, 1811, 1764, 1726, 1927, 1179, 1468, 1948, 1813, 1213, 1905, 1371, 1751, 1215, 1392, 1798, 1823, 1815, 1923, 1942, 1987, 1887, 1838, 1395, 2007, 1479, 1752, 1945, 1621, 1538, 1937, 565, 1969, 1493, 1291, 1438, 1578, 1770, 2005, 1703, 1712, 1943, 2003, 1499, 1903, 1760, 1950, 1990, 1185, 1809, 1337, 1358, 1743, 1707, 1671, 1788, 1785, 1972, 1863, 1690, 1512, 1963, 1825, 1460, 1828, 1902, 1874, 1755, 1951, 1830, 1767, 1787, 1373, 1709, 1514, 1807, 1791, 1724, 1859, 1590, 1976, 1572, 1947, 1913, 1995, 1728, 1624, 1731, 1706, 1782, 1994, 1851, 1843, 1773, 1982, 1685, 2001, 1346, 1200, 1746, 1520, 972, 1834, 1909, 2008, 1733, 1960, 1280, 1879, 1203, 1979, 1133, 1647, 1282, 1684, 860, 1444, 1780, 1989, 1795, 1819, 1797, 1842, 1796, 1457, 1839, 1853, 1711, 1883, 1146, 1734, 1389}

		left, middle, right := code.GetTriplet(nums)

		assert.Equal(t, 972, left)
		assert.Equal(t, 565, middle)
		assert.Equal(t, 483, right)

		product := code.GetTripletProduct(left, middle, right)
		assert.Equal(t, 265253940, product)
	})
}

func BenchmarkTestTriplet(b *testing.B) {
	nums := []int{1749, 1897, 881, 1736, 1161, 1720, 1676, 305, 264, 1904, 1880, 1173, 483, 1978, 1428, 1635, 1386, 1858, 1602, 1916, 1906, 1212, 1730, 1777, 1698, 1845, 1812, 1922, 1729, 1803, 1761, 1901, 1748, 1188, 1964, 1935, 1919, 1810, 1567, 1849, 1417, 1452, 54, 1722, 1784, 1261, 1744, 1594, 1526, 1771, 1762, 1894, 1717, 1716, 51, 1955, 1143, 1741, 1999, 1775, 1944, 1983, 1962, 1198, 1553, 1835, 1867, 1662, 1461, 1811, 1764, 1726, 1927, 1179, 1468, 1948, 1813, 1213, 1905, 1371, 1751, 1215, 1392, 1798, 1823, 1815, 1923, 1942, 1987, 1887, 1838, 1395, 2007, 1479, 1752, 1945, 1621, 1538, 1937, 565, 1969, 1493, 1291, 1438, 1578, 1770, 2005, 1703, 1712, 1943, 2003, 1499, 1903, 1760, 1950, 1990, 1185, 1809, 1337, 1358, 1743, 1707, 1671, 1788, 1785, 1972, 1863, 1690, 1512, 1963, 1825, 1460, 1828, 1902, 1874, 1755, 1951, 1830, 1767, 1787, 1373, 1709, 1514, 1807, 1791, 1724, 1859, 1590, 1976, 1572, 1947, 1913, 1995, 1728, 1624, 1731, 1706, 1782, 1994, 1851, 1843, 1773, 1982, 1685, 2001, 1346, 1200, 1746, 1520, 972, 1834, 1909, 2008, 1733, 1960, 1280, 1879, 1203, 1979, 1133, 1647, 1282, 1684, 860, 1444, 1780, 1989, 1795, 1819, 1797, 1842, 1796, 1457, 1839, 1853, 1711, 1883, 1146, 1734, 1389}
	for i := 0; i < b.N; i++ {
		code.GetTriplet(nums)
	}
}
