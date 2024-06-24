// 배열의 평균값
// numbers의 길이가 짝수이면 배열의 인덱스최소값과 인덱스최대값을 더한 값을 2로 나눈값
// numbers의 길이가 홀수이면 배열의 중간값
class Solution {
    public double solution(int[] numbers) {
        double answer = 0;
            if(numbers.length%2==0){
            answer = (numbers[0] + numbers[numbers.length-1])/2F;
        }else{
            answer = numbers[(numbers.length-1)/2];
        }
        
        return answer;
    }
}