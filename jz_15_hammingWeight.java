public class Solution {
    // you need to treat n as an unsigned value
    public int hammingWeight(int n) {

        int count = 0;
        
        int flag = 1;
        while (flag != 0) {
            
            if ((n & flag) != 0) {
                count++;
            }
            flag = flag << 1;
        }        
        return count;
    }
}