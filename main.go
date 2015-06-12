package main

type square struct {
	n      int
	index  squareIndex
	output [][]int
	filled int
}

type squareIndex struct {
	row    int
	column int
}

func main() {}

func Solve(n int) [][]int {
	s := square{n: n}

	s.fillEmptyOutput()
	if n%4 == 0 {
		return s.doubleEven()
	} else if n%2 == 1 {
		return s.siamese()
	} else {
		return [][]int{}
	}
}

func (s *square) fillEmptyOutput() {
	s.output = makeMatrixOfNSize(s.n)
}
func makeMatrixOfNSize(n int) [][]int {
	matrix := make([][]int, n)
	for x := 0; x < n; x++ {
		matrix[x] = make([]int, n)
	}
	return matrix
}
func (s *square) enterAtIndex(i int) {
	s.output[s.index.row][s.index.column] = i
}
func (s *square) valueAtIndex() int {
	return s.output[s.index.row][s.index.column]
}

// Double Even solution
func (s *square) doubleEven() [][]int {
	guide := [4][4]int{
		[4]int{1, 0, 0, 1},
		[4]int{0, 1, 1, 0},
		[4]int{0, 1, 1, 0},
		[4]int{1, 0, 0, 1},
	}

	s.enterAtIndex(1)
	for i := 1; i < s.n*s.n; i++ {
		s.findNextDoubleEvenStep()
		adjindex := s.index
		adjindex.column = adjindex.column % 4
		adjindex.row = adjindex.row % 4
		guideResult := guide[adjindex.row][adjindex.column]
		if guideResult == 1 {
			s.enterAtIndex(s.getMatrixCellCount() + 1)
		} else {
			s.enterAtIndex(s.n*s.n - s.getMatrixCellCount())
		}
	}
	return s.output
}
func (s *square) getMatrixCellCount() int {
	return s.index.column + (s.n * s.index.row)
}
func (s *square) findNextDoubleEvenStep() {
	sizeFromZero := s.n - 1
	s.index.column += 1
	if s.index.column > sizeFromZero {
		s.index.column = 0
		s.index.row += 1
	}
}

// The siamese solution
func (s *square) siamese() [][]int {
	s.index.row = 0
	s.index.column = s.findSiameseStart()
	for i := 1; i < s.n*s.n; i++ {
		s.enterAtIndex(i)
		s.index.row -= 1
		s.index.column += 1

		if i%s.n == 0 {
			s.index.row += 2
			s.index.column -= 1
		} else {
			if s.index.column == s.n {
				s.index.column -= s.n
			} else if s.index.row < 0 {
				s.index.row += s.n
			}
		}
	}
	s.enterAtIndex(s.n * s.n)
	return s.output
}
func (s *square) findSiameseStart() int {
	return (s.n - 1) / 2
}
