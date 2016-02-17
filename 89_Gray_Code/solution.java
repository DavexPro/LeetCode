public class Solution {
    public List<Integer> grayCode(int n) {
        List answerList = new ArrayList();
        for(int i=0;i<Math.pow(2,n);i++) answerList.add(i ^ (i >> 1));
        return answerList;
    }
}
