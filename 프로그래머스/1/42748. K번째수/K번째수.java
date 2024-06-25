import java.util.*;

class Solution {
    public int[] solution(int[] array, int[][] commands) { 
        int [] answer = new int [commands.length];
        int index = 0;
        // for each문으로 2차원배열 commands의 값들을 배열 s에 담아 돌림
        for (int[] s : commands) {
            // 잘라낸 배열 array를 담기위한 배열 slice 선언
            int [] slice = new int [s[1] - s[0] + 1];
            // 정렬위한 배열 arraySort 선언
            //int [] arraySort = new int [s[1] - s[0] + 1];
            
            // for문으로 slice에 잘라낸 array 담기
            for (int j = 0; j < s[1] - s[0] + 1; j++) {
                slice[j] = array[s[0]-1+j];
            }
            
            // slice 오름차순 정렬
            Arrays.sort(slice);
            // answer에 slice 인덱스값 담기
            answer[index] = slice[s[2]-1];
            index++;
        }

        return answer;
    }
}