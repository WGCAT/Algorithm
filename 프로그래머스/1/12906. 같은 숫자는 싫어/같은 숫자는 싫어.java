// 스택을 사용해 중복값 제거

import java.util.*;

public class Solution {
    public int[] solution(int []arr) {
        int[] answer = {};
        Stack<Integer> stack = new Stack<Integer>();
        // arr을 조건식으로 for문 돌림
        for(int i=0; i<arr.length; i++){
            // 첫데이터값 이면
            if(i==0) {
                // 스택의 가장 밑바닥에 값 추가
                stack.add(arr[i]);
                // 루프가 arr의 index i값과 같지 않을때
            } else if(stack.peek() != arr[i]) {
                // arr의 index i값 추가
                stack.add(arr[i]);
            }
        }
        // 크기가 스택의 크기인 배열 선언
        answer = new int[stack.size()];
        
        // 스택은 후입선출 방식 이므로 for문을 거꾸로 돌림
        for(int i=stack.size()-1; i >= 0;i--){
            // 배열에 스택 담기
            answer[i] = stack.pop();
        }
                return answer;
    }
}