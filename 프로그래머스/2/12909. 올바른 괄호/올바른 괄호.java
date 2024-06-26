import java.util.*;
class Solution {
    boolean solution(String s) {
        boolean answer = true;
        Stack<Character> stack = new Stack<Character>();
        for (int i = 0; i<s.length(); i++){
            char c = s.charAt(i);
            if (c=='('){
                stack.push(c);
            }else if (c==')'){
                if (stack.empty()) {
                    return false;
                }
                stack.pop();
            }
        }
        if (stack.size()!=0){
            answer = false;
        }
        return answer;
    }
}