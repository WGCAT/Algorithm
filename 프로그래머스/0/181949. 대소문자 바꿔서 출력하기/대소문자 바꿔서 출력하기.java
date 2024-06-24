// 대소문자 바꿔서 출력
// if문으로 문자열의 아스키코드를 비교하여 대소문자 전환 함수를 사용
// A ~ Z : 65 ~ 90, a ~ z : 97 ~ 122
import java.util.Scanner;
import java.util.*;

public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        // 반환값 넣을 배열 생성
        String[] result = new String[a.length()];
        // for each문으로 한글자씩 돌리기 위해 배열에 담는다 spilt로
        String[] str = a.split("");
        int i = 0;
        // for each문으로 str배열의 값을 s에 담아 돌림
        for (String s : str) {
            char ch = s.charAt(0);
            // 대문자라면
            if (ch >= 'A' && ch <= 'Z') {
                // 소문자 전환
                s = Character.toString(ch);
                s = s.toLowerCase();
                result[i] = s;
                i++;
            }else if(ch >= 'a' && ch <= 'z'){
                // 대문자 전환
                s = Character.toString(ch);
                s = s.toUpperCase();
                result[i] = s;
                i++;
            }
        }

        String re = String.join("", result);
        System.out.println(re);
        
    }
}