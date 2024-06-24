// 배열에 특정 값 포함 여부 확인

import java.util.*;


class Solution {
    public String solution(String[] participant, String[] completion) {
        String answer = "";
        
        // 배열 오름차순 정렬
        Arrays.sort(participant);
        Arrays.sort(completion);
        
        // 만약 participant가 1명이고 completion가 0명이면
        if (completion.length==0) {
            // participant[0] 반환
            return participant[0];
        }else {
            // for문으로 completion의 길이만큼 돌림
            for (int i = 0; i<completion.length; i++) {
                // participant[i]와 completion[i]의 값이 같지 않다면
                if (participant[i].equals(completion[i]) == false) {
                    // answer에 participant[i]값을 담는다
                    answer = participant[i];
                    break;
                    // participant[i]와 completion[i]의 값이 같다면
                }else {
                    // answer에 participant[completion.length]값을 담는다
                    answer = participant[completion.length];
                } 
            }
        return answer;
        }
        
    }
}
