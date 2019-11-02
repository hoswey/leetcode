type Direction int

const (
	Left 	Direction = iota 
	Right	
	Up
	Down
)

func generateMatrix(n int) [][]int {
    


	var cx, cy int

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
    
    if n == 1 {
        matrix[0][0] = 1
        return matrix
    }

	d := Right

	k := 1
	for {

		if matrix[cx][cy] != 0 {
			break
		}

		switch d {
		case Left:
			y := cy
			for ; y >= 0; y-- {

				if matrix[cx][y] != 0 {
					break
				}

				matrix[cx][y] = k
				k++
			}

			d = Up
			cx--
			cy = y + 1
		 
		case Right:
			y := cy
			for ;y < n; y++ {

				if matrix[cx][y] != 0 {
					break
				}

				matrix[cx][y] = k
				k++				
			}

			d = Down
			cx++
			cy = y - 1
		case Up:
			x := cx
			for ; x >= 0; x-- {			

				if matrix[x][cy] != 0 {
					break
				}

				matrix[x][cy] = k
				k++
			} 

			d = Right
			cy++
			cx = x + 1
		case Down:
			x := cx
			for ; x < n; x++ {
				if matrix[x][cy] != 0 {
					break
				}

				matrix[x][cy] = k
				k++
			} 
			d = Left				
			cy--
			cx = x - 1			
		}
	}

	return matrix
}