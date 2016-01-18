class Solution(object):
    def uniquePaths(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        paths = [[1] for i in range(m)]  # initialize the 1st row to be 1   
          
        for i in range(n-1):           # initialize the 1st column to be 1   
            paths[0].append(1)
   
        for i in range(m - 1):  
            for j in range(n - 1):  
                paths[i + 1].append(paths[i][j + 1] + paths[i + 1][j])  

        return paths[m-1][n-1]
