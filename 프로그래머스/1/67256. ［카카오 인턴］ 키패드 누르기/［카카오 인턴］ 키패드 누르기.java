class Solution {
    public String solution(int[] numbers, String hand) {
        String answer = "";
        int leftNum = 10;
        int rightNum = 12;


        for (int i : numbers) {
            if (i == 1 || i == 4 || i == 7) {
                answer += "L";
                leftNum = i;
            } else if (i == 3 || i == 6 || i == 9) {
                answer += "R";
                rightNum = i;
            } else {
                if (i == 0) {
                    i = 11;
                }
                int leftSpace = Math.abs(i - leftNum) / 3 + Math.abs(i - leftNum) % 3;
                int rightSpace = Math.abs(i - rightNum) / 3 + Math.abs(i - rightNum) % 3;
                if (leftSpace < rightSpace) {
                    answer += "L";
                    leftNum = i;
                } else if (leftSpace > rightSpace) {
                    answer += "R";
                    rightNum = i;
                } else {
                    if (hand.equals("left")) {
                        answer += "L";
                        leftNum = i;
                    } else {
                        answer += "R";
                        rightNum = i;
                    }
                }    
            }
            
        }
        return answer;
    }
}