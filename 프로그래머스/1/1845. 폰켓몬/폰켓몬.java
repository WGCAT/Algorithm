import java.util.*;
class Solution {
    public int solution(int[] nums) {
        int answer = 0;
        int getPon = nums.length/2;
        HashSet<Integer> set = new HashSet<>();
        
        for(int i : nums){
            set.add(i);
        }
        if(getPon < set.size()){
            answer =  getPon;
        }else {
            answer = set.size();
        }
        return answer;
    }
}