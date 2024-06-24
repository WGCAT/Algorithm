import java.util.*;

class Solution {
    public int[] solution(int[] array, int[][] commands) {
        // 
        ArrayList<Integer> answer = new ArrayList<Integer>();
        int index = 0;
        
        for (int[] s : commands) {
            ArrayList<Integer> arraySort = new ArrayList<Integer>();
            ArrayList<Integer> slice = new ArrayList<Integer>();
            
            for (int j = 0; j < s[1] - s[0] + 1; j++) {
                slice.add(j, array[s[0]-1+j]);
            }
                
            for (int i = 0; i<slice.size(); i++) {
                arraySort.add(i, slice.get(i));
            }
                
            Collections.sort(arraySort);
            answer.add(index, arraySort.get(s[2]-1));
            index++;
        }
        int [] result = new int [answer.size()];
        for (int k = 0; k<answer.size(); k++) {
            result[k] = answer.get(k);
        }

        return result;
    }
}