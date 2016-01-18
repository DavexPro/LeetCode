public class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int col = obstacleGrid.length;
        int row = obstacleGrid[0].length;
        int flag = 0;
        
        if(obstacleGrid[0][0] == 1)
            return 0;
        
        int[][] paths = new int[col][obstacleGrid[0].length];
        
        for(int i=0;i<col;i++){
            if(obstacleGrid[i][0]==1 || flag == 1){
                flag = 1;
                paths[i][0] = 0;
            }else{
                paths[i][0] = 1;
            }
        }
        
        flag = 0;
        
        for(int i=0;i<row;i++){
            if(obstacleGrid[0][i]==1 || flag == 1){
                flag = 1;
                paths[0][i] = 0;
            }else{
                paths[0][i] = 1;
            }
        }
        
        for(int i=1;i<col;i++){
            for(int j=1;j<row;j++){
                if(obstacleGrid[i][j] == 1)
                    paths[i][j] = 0;
                else
                    paths[i][j] = paths[i-1][j] + paths[i][j-1];
            }
        }
        return paths[col-1][row-1];
    }
}
