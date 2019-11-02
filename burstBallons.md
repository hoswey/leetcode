First, consider a sub-array from indices Left to Right(inclusive).
If we assume the balloon at index Last to be the last balloon to be burst in this sub-array, we would say the coined gained to be-A[left-1]*A[last]*A[right+1].
Also, the total Coin Gained would be this value, plus dp[left][last – 1] + dp[last + 1][right], where dp[i][j] means maximum coin gained for sub-array with indices i, j.
Therefore, for each value of Left and Right, we need find and choose a value of Last with maximum coin gained, and update the dp array.
Our Answer is the value at dp[1][N].


dp[left][right] = A[left-1] * A[last] * A[right+1] + dp[left][last-1] + dp[last+1][right]


A[0] * A[2] * A[1] + dp[1][1] + dp[2][1]


1 5 10 1

length = 1

	left = 1 
	right = 1

	dp[1][1] = A[0]* A[1] * A[2] + dp[1][0] + dp[2][1] = 50


	left = 2
	right = 2

	dp[2][2] = A[1] * A[2] * A[3] + dp[2][1] + d[3][2] = 50


length = 2

left = 1 right = 2

	last = 1

	dp[1][2] = A[0] * A[1] * A[3] + dp[1][0] + dp[2][2] = 55

	last = 2

	dp[1][2] = A[0] * A[2] * A[3] + dp[1][1] + dp[3][2] = 60
	
